package main

import (
	"errors"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	"golang.org/x/net/context"
)

func MainGRPC(addr string) {

	fmt.Println("Starting GRPC SErver......")
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()

	RegisterInfoServerServer(s, &server{})
	s.Serve(lis)

}

type server struct{}

// SetInfo - implements our InfoServer
func (s *server) SetInfo(ctx context.Context, in *InfoRequest) (*InfoReply, error) {
	if err := validate(in); err != nil {
		return &InfoReply{
			Message: "error",
		}, err
	}

	resp := fmt.Sprintf("\nName: %s, Age: %d, Height: %f,Weight: %f,Alive: %v,Desc: %v,Nickname: %s,Num: %d,Flt: %f,Dbl: %f,Tru: %v,Data: %v",
		in.Name,
		in.Age,
		in.Height,
		in.Weight,
		in.Alive,
		in.Desc,
		in.Nickname,
		in.Num,
		in.Flt,
		in.Dbl,
		in.Tru,
		in.Data,
	)

	return &InfoReply{Message: resp}, nil
}

// Validate - implement validatable
func (ir *InfoRequest) Validate() error {
	var err validationErrors
	if ir.Name == "" {
		err = append(err, errors.New("Name must be present"))
	}
	if ir.Age <= 0 {
		err = append(err, errors.New("Age must be real"))
	}
	if ir.Height <= 0 {
		err = append(err, errors.New("Height must be real"))
	}
	if len(err) == 0 {
		return nil
	}
	return err
}
