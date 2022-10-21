package server

import (
	"context"
	"log"

	protos "github.com/Tharsanan1/grpc-agent/protos/api"
)

type aPIService struct {
	protos.UnimplementedAPIServiceServer
}

func NewApiService() *aPIService {
	return &aPIService{}
}

func (s *aPIService) CreateAPI(ctx context.Context, api *protos.API ) (*protos.Response, error) {
	log.Printf("%q", api);
	// No feature was found, return an unnamed feature
	return &protos.Response{Result : true}, nil
}