package main

import (
  "io"
	"log"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	pb "github.com/bzz/grpc-streaming-poc/streaming"
)

func main() {
	conn, err := grpc.Dial("localhost:6000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to connect: %s", err)
	}
	defer conn.Close()

	client := pb.NewStreamingServiceClient(conn)
	stream, err := client.ParseUAST(context.Background(), &pb.Empty{})
	if err != nil {
		log.Fatalf("failed to call a client: %s", err)
	}

	for {
    fileUast, err := stream.Recv()
    if err == io.EOF {
        break
    }
    if err != nil {
        log.Fatalf("%v.ParseUAST(_) = _, %v", client, err)
    }
    log.Println(fileUast)
	}
	stream.CloseSend()
}
