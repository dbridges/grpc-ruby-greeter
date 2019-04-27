package main

import (
	"context"
	"google.golang.org/grpc"
	pb "grpc-ruby-greeter/greeter"
	"log"
	"net"
)

// START OMIT

type greeterServer struct{}

// Greet gives a nice hello
func (g *greeterServer) Greet(ctx context.Context, req *pb.GreetRequest) (*pb.MessageReply, error) {
	log.Println("Handling Greet request")
	return &pb.MessageReply{Message: "Hello " + req.Name}, nil
}

func main() {
	log.Println("Starting server...")
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal(err)
	}
	g := greeterServer{}                 // HL
	server := grpc.NewServer()           // HL
	pb.RegisterGreeterServer(server, &g) // HL
	log.Println("Serving on :50051")
	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

// END OMIT
