package main
import (
	"fmt"
	"log"
	"net"

	protos "github.com/Tharsanan1/grpc-agent/protos/api"
	server "github.com/Tharsanan1/grpc-agent/server"

	"google.golang.org/grpc"
)

func main() {

	// create new gRPC server
	s := grpc.NewServer()

	// create new instance of Translation server
	trans := server.NewApiService();

	// register it to the grpc server
	protos.RegisterAPIServiceServer(s, trans)

	// create socket to listen to requests
	tl, err := net.Listen("tcp", "localhost:8765")
	if err != nil {
		log.Fatal(fmt.Println("Error starting tcp listener on port 8765", err))
	}

	// start listening
	s.Serve(tl)
}