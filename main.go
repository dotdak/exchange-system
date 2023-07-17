package main

import (
	"context"
	"io/ioutil"
	"net"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/health/grpc_health_v1"

	"github.com/dotdak/exchange-system/gateway"
	"github.com/dotdak/exchange-system/handler"
	"github.com/dotdak/exchange-system/pkg/insecure"
	v1 "github.com/dotdak/exchange-system/proto/v1"
)

func main() {
	// Adds gRPC internal logs. This is quite verbose, so adjust as desired!
	log := grpclog.NewLoggerV2(os.Stdout, ioutil.Discard, ioutil.Discard)
	grpclog.SetLoggerV2(log)

	addr := "0.0.0.0:10000"
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}

	s := grpc.NewServer(
		// TODO: Replace with your own certificate!
		grpc.Creds(credentials.NewServerTLSFromCert(&insecure.Cert)),
	)

	// TODO
	h, err := handler.BuildHandler(context.Background())
	if err != nil {
		panic(err)
	}
	v1.RegisterBuyServiceServer(s, h)
	v1.RegisterWagerServiceServer(s, h)
	grpc_health_v1.RegisterHealthServer(s, handler.NewHealthCheck())

	// Serve gRPC Server
	log.Info("Serving gRPC on https://", addr)
	go func() {
		log.Fatal(s.Serve(lis))
	}()

	err = gateway.Run("dns:///" + addr)
	log.Fatalln(err)
}
