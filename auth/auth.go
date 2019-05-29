package main

import (
	pb "carconf-api/proto/user"
	"context"
	"fmt"

	"github.com/micro/go-micro"
)

type service struct{}

func (this *service) Login(ctx context.Context, req *pb.User, res *pb.Response) error {
	res.Message = "hello " + req.Name
	res.Success = true
	return nil
}

func main() {
	server := micro.NewService(
		micro.Name("go.micro.api.user"),
	)

	server.Init()

	pb.RegisterLoginServiceHandler(server.Server(), new(service))

	err := server.Run()
	if err != nil {
		fmt.Println(err)
	}
}
