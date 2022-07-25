package main

import (
	"fmt"

	"github.com/sagar23sj/benchmark-grpc-rest/grpc"

	"github.com/sagar23sj/benchmark-grpc-rest/rest"
)

func main() {
	fmt.Println("Initializing servers")

	var block = make(chan struct{})

	go grpc.GrpcStart()
	go rest.RestStart()
	<-block

}
