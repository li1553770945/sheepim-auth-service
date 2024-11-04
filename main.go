package main

import (
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	"github.com/kitex-contrib/obs-opentelemetry/provider"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	etcd "github.com/kitex-contrib/registry-etcd"
	"net"
	"os"
	configInfra "sheepim-auth-service/biz/infra/config"
	"sheepim-auth-service/biz/infra/container"
	"sheepim-auth-service/biz/infra/log"
	"sheepim-auth-service/biz/infra/trace"
	"sheepim-auth-service/kitex_gen/auth/authservice"
)

func main() {

	env := os.Getenv("ENV")
	if env == "" {
		env = "development"
	}
	config := configInfra.GetConfig(env)
	log.InitLog()
	p := trace.InitTrace(config)

	container.InitGlobalContainer(config)
	App := container.GetGlobalContainer()

	defer func(p provider.OtelProvider, ctx context.Context) {
		err := p.Shutdown(ctx)
		if err != nil {
			klog.Fatalf("server stopped with error:", err)
		}
	}(p, context.Background())

	addr, err := net.ResolveTCPAddr("tcp", App.Config.ServerConfig.ListenAddress)
	if err != nil {
		panic("设置监听地址出错")
	}

	r, err := etcd.NewEtcdRegistry([]string{App.Config.EtcdConfig.Endpoint}) // r should not be reused.
	if err != nil {
		klog.Fatal(err)
	}
	serviceName := App.Config.ServerConfig.ServiceName
	svr := authservice.NewServer(
		new(AuthServiceImpl),
		server.WithSuite(tracing.NewServerSuite()),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: serviceName}),
		server.WithRegistry(r),
		server.WithServiceAddr(addr),
	)
	if err := svr.Run(); err != nil {
		klog.Fatalf("server stopped with error:", err)
	}
}
