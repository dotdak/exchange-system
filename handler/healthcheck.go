package handler

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/status"
)

func NewHealthCheck() grpc_health_v1.HealthServer {
	return &HealthCheckImpl{}
}

type HealthCheckImpl struct{}

// Check implements grpc_health_v1.HealthServer.
func (*HealthCheckImpl) Check(context.Context, *grpc_health_v1.HealthCheckRequest) (*grpc_health_v1.HealthCheckResponse, error) {
	return &grpc_health_v1.HealthCheckResponse{Status: grpc_health_v1.HealthCheckResponse_SERVING}, nil
}

// Watch implements grpc_health_v1.HealthServer.
func (*HealthCheckImpl) Watch(*grpc_health_v1.HealthCheckRequest, grpc_health_v1.Health_WatchServer) error {
	return status.Error(codes.Unimplemented, "unimplemented")
}
