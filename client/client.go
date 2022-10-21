package main

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc"
	apiProtos "github.com/Tharsanan1/grpc-agent/grpc/api"
	swaggerProtos "github.com/Tharsanan1/grpc-agent/grpc/swagger"
)

func main() {
	var opts []grpc.DialOption
    opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	conn, err := grpc.Dial("localhost:8765", opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()
	client := apiProtos.NewAPIServiceClient(conn);

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	
	response, err :=  client.CreateAPI(ctx, &apiProtos.API{
		ApiUUID: "test",
		ApiVersion: "1.0.0",
	})

	if err != nil {
		log.Fatalf("client.ListFeatures failed: %v", err)
	}

	log.Printf("%q",response)
	swaggerClient := swaggerProtos.NewSwaggerServiceClient(conn);
	response1, err1 :=  swaggerClient.CreateSwagger(ctx, &swaggerProtos.Swagger{
		ApiUUID: "test",
		Swagger: "1.0.0",
	})

	if err1 != nil {
		log.Fatalf("client.ListFeatures failed: %v", err1)
	}

	log.Printf("%q",response1)
}
