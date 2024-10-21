package rpc

import (
	"github.com/cloudwego/kitex/client"
	"log"
	"sheepim-auth-service/kitex_gen/user/userservice"
)

func NewUserClient() userservice.Client {
	userClient, err := userservice.NewClient(
		"hello",
		client.WithHostPorts("0.0.0.0:8888"),
	)
	if err != nil {
		log.Fatal(err)
	}
	return userClient
}
