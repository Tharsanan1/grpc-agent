package service

import (
	"context"
	"log"
	swaggerProtos "github.com/Tharsanan1/grpc-agent/grpc/swagger"
	commonProtos "github.com/Tharsanan1/grpc-agent/grpc/common"
)

type swaggerService struct {
	swaggerProtos.UnimplementedSwaggerServiceServer
}

func NewSwaggerService() *swaggerService {
	return &swaggerService{}
}

func (s *swaggerService) CreateSwagger(ctx context.Context, swagger *swaggerProtos.Swagger ) (*commonProtos.Response, error) {
	log.Printf("%q", swagger);
	// No feature was found, return an unnamed feature
	return &commonProtos.Response{Result : true}, nil
}