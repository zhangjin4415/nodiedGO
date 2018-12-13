package main

import (
	"fmt"
	"net"
	pb "nodiedGO/reverse"
	"os/exec"
	"time"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

func main() {

	listener, err := net.Listen("tcp", ":5300")
	if err != nil {
		grpclog.Fatalf("failed to listen: %v", err)
	}
	opts := []grpc.ServerOption{}
	grpcServer := grpc.NewServer(opts...)

	srv := &server{}
	srv.numTime = 0
	go func() {
		for {
			if srv.numTime > 3 {
				fmt.Println("client is died! to run one...")
				srv.numTime = 0
				cmd := exec.Command("client")
				cmd.Run()
			}
			srv.numTime++
			time.Sleep(time.Duration(2) * time.Second)
		}
	}()

	pb.RegisterReverseServer(grpcServer, srv)
	grpcServer.Serve(listener)
}

type server struct {
	numTime int
}

func (s *server) Do(c context.Context, request *pb.Request) (response *pb.Response, err error) {
	fmt.Println("accept meg: ", request.Message)
	s.IsRun()
	output := "ok"
	response = &pb.Response{
		Message: output,
	}
	return response, nil
}

func (s *server) IsRun() {
	s.numTime = 0
}
