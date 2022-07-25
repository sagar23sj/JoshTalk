package main

import (
	"bytes"
	"net/http"
	"testing"
)

func BenchmarkRESTSetInfo(b *testing.B) {
	client := &http.Client{}

	buf := bytes.NewBufferString(`
	{
		"name":"test",
		"age":1,
		"height":1
	}
`)

	b.ResetTimer()
	// run http posts against it
	for i := 0; i < b.N; i++ {
		b.ReportAllocs()

		client.Post("https://localhost:4444", "application/json", buf)
	}
}
