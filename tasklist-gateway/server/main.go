// Copyright 2018 budougumi0617 All Rights Reserved.

// It implements the route guide service whose definition can be found in proto/task_list.proto
package main

import (
	"context"
	"errors"
	"log"
	"net"

	tlpb "github.com/budougumi0617/sandbox-grpc/tasklist-gateway/proto"
	google_protobuf "github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":50051"
)

// server is used to implement sa
type server struct{}

func (s *server) GetTask(ctx context.Context, req *tlpb.GetTaskRequest) (*tlpb.Task, error) {
	log.Println("GetTask in gPRC server")
	var task = &tlpb.Task{
		Id:     10,
		Title:  "Implement gRPC server",
		Detail: "Implement interface",
	}
	if req.Id == task.Id {
		return task, nil
	}
	return nil, errors.New("Not find Task")
}
func (s *server) ListTasks(_ *google_protobuf.Empty, stream tlpb.TaskManager_ListTasksServer) error {
	log.Println("ListTasks in gPRC server")
	tasks := []*tlpb.Task{
		&tlpb.Task{
			Id:     20,
			Title:  "List Tasks 1",
			Detail: "Implement interface",
		},
		&tlpb.Task{
			Id:     21,
			Title:  "List Tasks 2",
			Detail: "Implement interface",
		},
		&tlpb.Task{
			Id:     22,
			Title:  "List Tasks 3",
			Detail: "Implement interface",
		},
	}
	for _, task := range tasks {
		if err := stream.Send(task); err != nil {
			return err
		}
	}
	return nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	tlpb.RegisterTaskManagerServer(s, &server{})
	// Register reflection service on gRPC server.
	reflection.Register(s)
	log.Printf("gRPC Server started: localhost%s\n", port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
