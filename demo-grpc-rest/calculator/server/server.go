package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/sagar23sj/benchmark-grpc-rest/sample"
	"google.golang.org/grpc"
)

type server struct{}

func main() {
	lis, err := net.Listen("tcp", ":9090")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	s := grpc.NewServer()
	sample.RegisterSampleServiceServer(s, &server{})
	log.Println("Starting gRPC server")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

func (s *server) Hello(_ context.Context, random *sample.BenchLarge) (pbresp *sample.Response, err error) {
	resp := fmt.Sprintf("\nName: %s, Age: %d, Height: %f,Weight: %f,Alive: %v,Desc: %v,Nickname: %s,Num: %d,Flt: %f,Dbl: %f,Tru: %v,Data: %v",
		random.Name,
		random.Age,
		random.Height,
		random.Weight,
		random.Alive,
		random.Desc,
		random.Nickname,
		random.Num,
		random.Flt,
		random.Dbl,
		random.Tru,
		random.Data,
	)

	pbresp.Message = resp
	return pbresp, nil
}
