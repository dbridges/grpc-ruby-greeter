package main

import (
	"context"
	"google.golang.org/grpc"
	pb "grpc-ruby-greeter/greeter"
	"log"
	"net"
	"os"
	"time"
)

type greeterServer struct{}

// Greet gives a nice hello
func (g *greeterServer) Greet(ctx context.Context, req *pb.GreetRequest) (*pb.MessageReply, error) {
	log.Println("Handling Greet request")
	return &pb.MessageReply{Message: "Hello " + req.Name}, nil
}

func serve() {
	log.Println("Starting server...")
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal(err)
	}
	g := greeterServer{}
	server := grpc.NewServer()
	pb.RegisterGreeterServer(server, &g)
	log.Println("Serving on :50051")
	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func greet() {
	conn, err := grpc.Dial(":50051", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	client := pb.NewGreeterClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	res, err := client.Greet(ctx, &pb.GreetRequest{Name: "Dan"})
	if err != nil {
		log.Fatal(err)
	}
	log.Println(res.Message)
}

func main() {
	switch os.Args[1] {
	case "serve":
		serve()
	case "greet":
		greet()
	}
}
