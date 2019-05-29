package main

import (
	usersvc "carconf-api/proto/user"
	"context"
	"fmt"
	"log"
	"github.com/micro/go-micro"
)

type User struct{}

func (u *User) Login(ctx context.Context, req *usersvc.LoginReq, rsp *usersvc.LoginRsp) error {
	log.Print("Received Say.Hello request")

	rsp.Message = "hello " + req.Name
	rsp.Success = true
	return nil
}

func main() {
	server := micro.NewService(
		micro.Name("allride.micro.svc.user"),
	)

	server.Init()

	usersvc.RegisterUserHandler(server.Server(), new(User))

	err := server.Run()
	if err != nil {
		fmt.Println(err)
	}
}
