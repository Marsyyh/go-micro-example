package main

import (
	"context"
	"fmt"

	usersrv "carconf-api/proto/user"

	"github.com/micro/go-micro"
)

func main() {
	service := micro.NewService(
		micro.Name("go.micro.user.client"),
	)

	service.Init()

	userCli := usersrv.NewLoginService("go.micro.srv.user", service.Client())
	res, err := userCli.Login(context.TODO(), &usersrv.User{Name: "ares"})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res.Message)
}
