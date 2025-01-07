package main

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	pb "github.com/namikmesic/playground/rpc-demo/libs/proto"
	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50052, "port number")
)

type server struct {
	pb.UnimplementedTelemetryServiceServer
}

func (s *server) StreamTelemetryBatch(stream pb.TelemetryService_StreamTelemetryBatchServer) error {
	for {
		batch := &pb.TelemetryBatch{}
		err := stream.RecvMsg(batch)
		if errors.Is(err, io.EOF) {
			log.Println("Stream received EOF. Closing stream.")
			if err := stream.SendMsg(&pb.BatchResponse{
				StatusCode: 0,
				Message:    fmt.Sprintf("Batch of %d items processed", len(batch.Items)),
			}); err != nil {
				log.Printf("Failed to send ack: %v", err)
				return err
			}
			return nil
		}
		if err != nil {
			log.Printf("Failed to receive message: %v", err)
			return err
		}

		// Process each batch item
		for _, item := range batch.Items {
			log.Printf("Received item from host %s with metric %s and value %f at %s",
				item.HostId, item.MetricName, item.Value, item.Timestamp)
		}
		log.Printf("Received batch with %d items", len(batch.Items))

		// Send ACK after processing each batch
		if err := stream.SendMsg(&pb.BatchResponse{
			StatusCode: 0,
			Message:    fmt.Sprintf("Batch of %d items processed", len(batch.Items)),
		}); err != nil {
			log.Printf("Failed to send ack: %v", err)
			return err
		}
	}
}

func main() {
	flag.Parse()

	address := fmt.Sprintf(":%v", *port)
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterTelemetryServiceServer(s, &server{})

	// Graceful shutdown handling
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-quit
		log.Println("Shutting down gRPC server...")
		s.GracefulStop()
		log.Println("gRPC server stopped.")
	}()

	log.Printf("Starting gRPC server on %s", address)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
