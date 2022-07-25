package main

import (
	"log"
	"testing"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func BenchmarkGRPCSetInfo(b *testing.B) {

	conn, err := grpc.Dial("localhost:4443", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := NewInfoServerClient(conn)

	b.ResetTimer()
	// run grpc calls against it
	for i := 0; i < b.N; i++ {
		b.ReportAllocs()
		client.SetInfo(context.Background(), &InfoRequest{
			Name:   "test",
			Age:    1,
			Height: 1,
		})
	}

}
