package adder

import (
	"context"
	"fmt"
	"go_service/pkg/api"
)

// GRPCServer ...
type GRPCServer struct {
	api.UnimplementedAdderServer
}

// Add ...
func (s *GRPCServer) Add(ctx context.Context, req *api.AddRequest) (*api.AddResponse, error) {
	fmt.Println(ctx)
	return &api.AddResponse{Result: req.GetX() + req.GetY()}, nil
}
