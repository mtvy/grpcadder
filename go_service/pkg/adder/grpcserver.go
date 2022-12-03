package adder

import (
	"context"
	"go_service/pkg/api"
)

// GRPCServer ...
type GRPCServer struct {
	api.UnimplementedAdderServer
}

// Add ...
func (s *GRPCServer) Add(ctx context.Context, req *api.AddRequest) (*api.AddResponse, error) {
	return &api.AddResponse{Result: req.GetX() + req.GetY()}, nil
}
