# exchange-system

Grpc-gateway powered REST exchange system in Go.

## Running

Before running, please fill in essential key-value in `.env`, then make `start.sh` be executable `chmod +x ./start.sh`

Using `start.sh` scripts to run:

1. Run service connecting to existing database:

```
./start.sh -s
```

2. Run service with its own database:

```
./start.sh -ls
```

An OpenAPI UI is served on https://0.0.0.0:8080/swagger/#/.

### Running the standalone server

If you want to use a separate gRPC server, for example one written in Java or C++, you can run the
standalone web server instead:

```
$ go run ./cmd/standalone/ --server-address dns:///0.0.0.0:10000
```

## Development

After cloning the repo, there are a couple of initial steps;

1. Install the generate dependencies with `make install`.
   This will install `buf`, `protoc-gen-go`, `protoc-gen-go-grpc`, `protoc-gen-grpc-gateway`, `protoc-gen-validate`,
   `protoc-gen-openapiv2` and `wire` which are necessary for us to generate the Go and swagger files.

1. Generate the files with `make generate`.

1. Running `main.go` starts a web server on https://0.0.0.0:8080/. You can configure
   the port used with the `$PORT` environment variable, and to serve on HTTP set
   `$SERVE_HTTP=true`.

```
$ SERVE_HTTP=true go run main.go
```
