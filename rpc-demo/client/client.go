// Client
package main

import (
	"context"
	"errors"
	"flag"
	"io"
	"log"
	"time"

	pb "github.com/namikmesic/rpc-demo/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	addr = flag.String("addr", "localhost:50052", "the address to connect to")
)

func main() {
	flag.Parse()

	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewTelemetryServiceClient(conn)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	stream, err := client.StreamTelemetryBatch(ctx)
	if err != nil {
		log.Fatalf("Failed to create stream: %v", err)
	}
	defer stream.CloseSend()

	// Goroutine to receive ACKs asynchronously
	go func() {
		for {
			resp := &pb.BatchResponse{}
			err := stream.RecvMsg(resp)
			if errors.Is(err, io.EOF) {
				log.Println("Server closed stream.")
				break
			}
			if err != nil {
				log.Printf("Failed to receive ack: %v", err)
				cancel()
				return
			}
			log.Printf("Received ack: %s", resp.Message)
		}
	}()

	// Send batches periodically
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for i := 0; i < 10; i++ {
		select {
		case <-ticker.C:
			batch := &pb.TelemetryBatch{
				Items: []*pb.TelemetryData{
					{HostId: "host-123", MetricName: "cpu_usage", Value: float32(30 + i), Timestamp: time.Now().Format(time.RFC3339)},
					{HostId: "host-123", MetricName: "memory_usage", Value: float32(40 + i), Timestamp: time.Now().Format(time.RFC3339)},
				},
			}

			// Send the batch
			if err := stream.Send(batch); err != nil {
				log.Fatalf("Failed to send batch: %v", err)
			}
			log.Printf("Sent batch %d", i+1)
		case <-ctx.Done():
			log.Println("Context canceled, stopping batch send.")
			return
		}
	}

	log.Println("Finished sending all batches.")
}
