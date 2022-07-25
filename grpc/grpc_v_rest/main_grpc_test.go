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
		_, err := client.SetInfo(context.Background(), &InfoRequest{
			Name:     "Tester",
			Age:      20,
			Height:   5.8,
			Weight:   180.7,
			Alive:    true,
			Desc:     []byte("Lets benchmark some json and protobuf"),
			Nickname: "Another name",
			Num:      2314,
			Flt:      123451231.1234,
			Data: []byte(`If you’ve ever heard of ProtoBuf you may be thinking that
				the results of this benchmarking experiment will be obvious, JSON < ProtoBuf.
				My interest was in how much they actually differ in practice.
				How do they compare on a couple of different metrics, specifically serialization
				and de-serialization speeds, and the memory footprint of encoding the data.
				I was also curious about how the different serialization methods would
				behave under small, medium, and large chunks of data.`),
		})

		if err != nil {
			b.Fatalf("http request failed: %v", err)
		}
	}

}
