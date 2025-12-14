package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/Rushi2398/learn-gRPC/proto"
)

func callSayHelloBidirectionalStream(client pb.GreetServiceClient, names *pb.NamesList) {
	log.Printf("Bidirectional Streaming Started")
	stream, err := client.SayHelloBidirectionalStreaming(context.Background())
	if err != nil {
		log.Fatalf("Could not send names: %v", err)
	}

	waitch := make(chan struct{})
	go func() {
		for {
			message, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("Error while receiving: %v", err)
			}
			log.Println(message)
		}
		close(waitch)
	}()

	for _, name := range names.Names {
		req := &pb.HelloRequest{
			Name: name,
		}
		if err := stream.Send(req); err != nil {
			log.Fatalf("Error while sending: %v", err)
		}
		time.Sleep(2 * time.Second)
	}
	stream.CloseSend()
	<-waitch
	log.Printf("Bidirectional Streaming finished")
}
