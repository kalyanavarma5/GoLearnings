package main

import (
	"context"
	"google.golang.org/grpc"
	pb "grpc-example/proto"
	"io"
	"log"
	"time"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewTodoServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	stream, err := client.GetTodos(ctx, &pb.TodoRequest{UserId: "user123"})
	if err != nil {
		log.Fatalf("error calling GetTodos: %v", err)
	}

	for {
		todo, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("error receiving todo: %v", err)
		}
		log.Printf("Received Todo: ID=%s, Title=%s, Completed=%v", todo.Id, todo.Title, todo.Completed)
	}
}
