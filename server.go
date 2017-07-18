package main

import (
	"log"
	"net"

	"google.golang.org/grpc"

	pb "github.com/bzz/grpc-streaming-poc/streaming"
)

type streamingServer struct {
	files []string
}

func (s *streamingServer) ParseUAST(req *pb.Empty, stream pb.StreamingService_ParseUASTServer) error {
	for _, file := range s.files {
		log.Println("Parsing UAST for ", file)
		if err := stream.Send(&pb.FileUAST{File: file, Uast: "yyy"}); err != nil {
			return err
		}
	}
	return nil
}

func main() {
	grpcServer := grpc.NewServer()
	pb.RegisterStreamingServiceServer(grpcServer, &streamingServer{files: []string{"a.py", "b.java", "c/d.cpp"}})

	l, err := net.Listen("tcp", ":6000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	log.Println("Listening on tcp://localhost:6000")
	grpcServer.Serve(l)
}
