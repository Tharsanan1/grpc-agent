package main
import (
	"fmt"
	"log"
	"net"

	apiProtos "github.com/Tharsanan1/grpc-agent/grpc/api"
	swaggerProtos "github.com/Tharsanan1/grpc-agent/grpc/swagger"
	service "github.com/Tharsanan1/grpc-agent/service"

	"google.golang.org/grpc"
)

func main() {

	s := grpc.NewServer()

	apiService := service.NewApiService();
	apiProtos.RegisterAPIServiceServer(s, apiService)

	swaggerService := service.NewSwaggerService();
	swaggerProtos.RegisterSwaggerServiceServer(s, swaggerService)

	tl, err := net.Listen("tcp", "localhost:8765")
	if err != nil {
		log.Fatal(fmt.Println("Error starting tcp listener on port 8765", err))
	}

	// start listening
	s.Serve(tl)
}