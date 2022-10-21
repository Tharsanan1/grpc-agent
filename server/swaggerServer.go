package server

import (
	"context"
	"log"

	protos "github.com/Tharsanan1/grpc-agent/protos/swagger"
)

type swaggerService struct {
	protos.UnimplementedSwaggerServiceServer
}

func NewSwaggerService() *swaggerService {
	return &swaggerService{}
}

func (s *swaggerService) CreateAPI(ctx context.Context, swagger *protos.Swagger ) (*protos.Response, error) {
	log.Printf("%q", swagger);
	// No feature was found, return an unnamed feature
	return &protos.Response{Result : true}, nil
}