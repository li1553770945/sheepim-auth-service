package rpc

import (
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	etcd "github.com/kitex-contrib/registry-etcd"
	"github.com/li1553770945/sheepim-auth-service/biz/infra/config"
	"github.com/li1553770945/sheepim-user-service/kitex_gen/user/userservice"
	"log"
)

func NewUserClient(config *config.Config) userservice.Client {
	r, err := etcd.NewEtcdResolver(config.EtcdConfig.Endpoint)
	userClient, err := userservice.NewClient(
		config.RpcConfig.UserServiceName,
		client.WithResolver(r),
		client.WithSuite(tracing.NewClientSuite()),
		client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: config.ServerConfig.ServiceName}),
	)
	if err != nil {
		log.Fatal(err)
	}
	return userClient
}
