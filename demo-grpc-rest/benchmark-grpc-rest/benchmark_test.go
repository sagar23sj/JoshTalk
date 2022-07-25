package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/sagar23sj/benchmark-grpc-rest/sample"
)

func BenchmarkHTTPJSON(b *testing.B) {
	client := &http.Client{}

	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		b.ReportAllocs()
		doPost(client, b)
	}
}

func doPost(client *http.Client, b *testing.B) {
	u := &sample.BenchLarge{
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
	}

	buf := new(bytes.Buffer)
	json.NewEncoder(buf).Encode(u)

	resp, err := client.Post("http://127.0.0.1:8090/hello", "application/json", buf)
	if err != nil {
		b.Fatalf("http request failed: %v", err)
	}

	defer resp.Body.Close()

	// We need to parse response to have a fair comparison as gRPC does it
	var target sample.Response
	decodeErr := json.NewDecoder(resp.Body).Decode(&target)
	if decodeErr != nil {
		b.Fatalf("unable to decode json: %v", decodeErr)
	}

}
