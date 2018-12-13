package main

import (
	"context"
	"fmt"
	pb "nodiedGO/reverse"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

func main() {
	opts := []grpc.DialOption{
		grpc.WithInsecure(),
	}
	conn, err := grpc.Dial("127.0.0.1:5300", opts...)
	if err != nil {
		grpclog.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()
	client := pb.NewReverseClient(conn)
	request := &pb.Request{
		Message: "right",
	}

	go func() {
		for {
			response, err := client.Do(context.Background(), request)
			if err != nil {
				grpclog.Fatalf("fail to dial: %v", err)
			}
			if response.Message != "ok" {
				fmt.Println(err)
			}
			fmt.Println(response.Message)
			time.Sleep(time.Duration(1) * time.Second)
		}
	}()
	time.Sleep(time.Duration(5) * time.Second)
}
