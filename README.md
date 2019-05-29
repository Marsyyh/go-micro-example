# go-micro-example

This is example of how to use go-micro for micro service build. Two servers builded, one api service that recieve request from go-micro apigateway, and talk to another user service with protobuf.

## How to run
make sure you have docker installed
```
make build
make run
```

now you can access your localhost:8080/user/user/login?name=john to see the message
