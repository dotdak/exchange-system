package gateway

import (
	"context"
	"crypto/tls"
	"fmt"
	"io/fs"
	"io/ioutil"
	"mime"
	"net/http"
	"os"
	"strings"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/protobuf/encoding/protojson"

	"github.com/dotdak/exchange-system/pkg/es_errors"
	"github.com/dotdak/exchange-system/pkg/insecure"
	"github.com/dotdak/exchange-system/pkg/utils"
	v1 "github.com/dotdak/exchange-system/proto/v1"
	"github.com/dotdak/exchange-system/third_party"
)

const SWAGGER_PREFIX = "/swagger"

// getOpenAPIHandler serves an OpenAPI UI.
func getOpenAPIHandler() http.Handler {
	mime.AddExtensionType(".svg", "image/svg+xml")
	// Use subdirectory in embedded files
	subFS, err := fs.Sub(third_party.OpenAPI, "OpenAPI")
	if err != nil {
		panic("couldn't create sub filesystem: " + err.Error())
	}

	return http.StripPrefix(SWAGGER_PREFIX, http.FileServer(http.FS(subFS)))
}

// Run runs the gRPC-Gateway, dialling the provided address.
func Run(dialAddr string) error {
	// Adds gRPC internal logs. This is quite verbose, so adjust as desired!
	log := grpclog.NewLoggerV2(os.Stdout, ioutil.Discard, ioutil.Discard)
	grpclog.SetLoggerV2(log)
	ctx := context.Background()
	// Create a client connection to the gRPC Server we just started.
	// This is where the gRPC-Gateway proxies the requests.
	conn, err := grpc.DialContext(
		ctx,
		dialAddr,
		grpc.WithTransportCredentials(credentials.NewClientTLSFromCert(insecure.CertPool, "")),
		grpc.WithBlock(),
		grpc.WithUnaryInterceptor(otelgrpc.UnaryClientInterceptor()),
	)
	if err != nil {
		return fmt.Errorf("failed to dial server: %w", err)
	}

	gwmux := runtime.NewServeMux(
		runtime.WithErrorHandler(func(ctx context.Context, sm *runtime.ServeMux, m runtime.Marshaler, w http.ResponseWriter, r *http.Request, err error) {
			newError := runtime.HTTPStatusError{
				HTTPStatus: 400,
				Err:        err,
			}
			// using default handler to do the rest of heavy lifting of marshaling error and adding headers
			runtime.DefaultHTTPErrorHandler(ctx, sm, m, w, r, &newError)
		}),
		runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.HTTPBodyMarshaler{
			Marshaler: &runtime.JSONPb{
				MarshalOptions: protojson.MarshalOptions{
					EmitUnpopulated: true,
					UseProtoNames:   true,
				},
				UnmarshalOptions: protojson.UnmarshalOptions{
					DiscardUnknown: true,
				},
			},
		}),
		runtime.WithHealthEndpointAt(grpc_health_v1.NewHealthClient(conn), "/health"),
	)
	err = es_errors.AnyError(
		v1.RegisterBuyServiceHandler(ctx, gwmux, conn),
		v1.RegisterWagerServiceHandler(ctx, gwmux, conn),
	)

	if err != nil {
		return fmt.Errorf("failed to register gateway: %w", err)
	}

	oa := getOpenAPIHandler()

	port := utils.Any(os.Getenv("PORT"), "8080")
	gatewayAddr := "0.0.0.0:" + port
	gwServer := &http.Server{
		Addr: gatewayAddr,
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if strings.HasPrefix(r.URL.Path, SWAGGER_PREFIX) {
				// r.URL.Path = strings.TrimSuffix(r.URL.Path, "/") + "/"
				oa.ServeHTTP(w, r)
				return
			}

			gwmux.ServeHTTP(w, r)
		}),
	}
	// Empty parameters mean use the TLS Config specified with the server.
	if strings.ToLower(os.Getenv("SERVE_HTTP")) == "true" {
		log.Info("Serving gRPC-Gateway and OpenAPI Documentation on http://", gatewayAddr)
		return fmt.Errorf("serving gRPC-Gateway server: %w", gwServer.ListenAndServe())
	}

	gwServer.TLSConfig = &tls.Config{
		Certificates: []tls.Certificate{insecure.Cert},
	}
	log.Info("Serving gRPC-Gateway and OpenAPI Documentation on https://", gatewayAddr)
	return fmt.Errorf("serving gRPC-Gateway server: %w", gwServer.ListenAndServeTLS("", ""))
}
