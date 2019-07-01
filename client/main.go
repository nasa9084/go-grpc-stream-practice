package main

import (
	"context"
	"io"
	"log"

	rpc "github.com/nasa9084/go-grpc-stream-practice/rpc/idl"
	"google.golang.org/grpc"
)

func main() {
	if err := execute(); err != nil {
		log.Fatal(err)
	}
}

func execute() error {
	conn, err := grpc.Dial("localhost:5000", grpc.WithInsecure())
	if err != nil {
		return err
	}
	defer conn.Close()
	client := rpc.NewStreamClient(conn)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	stream, err := client.Do(ctx, &rpc.Empty{})
	if err != nil {
		return err
	}
	for i := 0; i < 5; i++ {
		pong, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		log.Println(pong.Message)
	}
	return nil
}
