package service

import (
	"context"
	"log"

	apiProtos "github.com/Tharsanan1/grpc-agent/grpc/api"
	commonProtos "github.com/Tharsanan1/grpc-agent/grpc/common"
)

type aPIService struct {
	apiProtos.UnimplementedAPIServiceServer
}

func NewApiService() *aPIService {
	return &aPIService{}
}

func (s *aPIService) CreateAPI(ctx context.Context, api *apiProtos.API ) (*commonProtos.Response, error) {
	log.Printf("%q", api);
	// No feature was found, return an unnamed feature
	return &commonProtos.Response{Result : true}, nil
}