package main

import (
	_ "context"
	"google.golang.org/grpc"
	"log"
	"net"
	"time"
	pb. "grpc-example/proto;proto"

)

type todoServer struct {
	pb.UnimplementedTodoServiceServer
}

func (s *todoServer) GetTodos(req *pb.TodoRequest, stream pb.TodoService_GetTodosServer) error {
	todos := []*pb.Todo{
		{Id: "1", Title: "Buy groceries", Completed: false},
		{Id: "2", Title: "Clean the house", Completed: true},
		{Id: "3", Title: "Walk the dog", Completed: false},
	}

	for _, todo := range todos {
		if err := stream.Send(todo); err != nil {
			return err
		}
		time.Sleep(500 * time.Millisecond) // Simulate delay
	}

	return nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterTodoServiceServer(grpcServer, &todoServer{})

	log.Println("Starting gRPC server on port 50051...")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
