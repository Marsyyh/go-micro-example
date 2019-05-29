package main

import (
	"context"
	"encoding/json"
	"log"
	"strings"

	usersvc "carconf-api/proto/user"

	"github.com/micro/go-micro"
	"github.com/micro/go-micro/errors"
	api "github.com/micro/micro/api/proto"
)

type User struct {
	Client usersvc.UserService
}

func (u *User) Login(ctx context.Context, req *api.Request, rsp *api.Response) error {
	log.Print("Received Say.Hello API request")

	name, ok := req.Get["name"]
	if !ok || len(name.Values) == 0 {
		return errors.BadRequest("allride.micro.api.usersvc", "Name cannot be blank")
	}

	response, err := u.Client.Login(ctx, &usersvc.LoginReq{
		Name: strings.Join(name.Values, " "),
	})
	if err != nil {
		return err
	}

	rsp.StatusCode = 200
	b, _ := json.Marshal(map[string]string{
		"message": response.Message,
	})
	rsp.Body = string(b)

	return nil
}

func main() {
	service := micro.NewService(
		micro.Name("allride.micro.api.user"),
	)

	service.Init()

	service.Server().Handle(
		service.Server().NewHandler(
			&User{Client: usersvc.NewUserService("allride.micro.svc.user", service.Client())},
		),
	)

	err := service.Run()
	if err != nil {
		log.Fatal(err)
	}
	// userCli := usersvc.NewUserService("allride.micro.svc.usersvc", service.Client())
	// res, err := userCli.Login(context.TODO(), &usersvc.User{Name: "ares"})
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(res.Message)
}
