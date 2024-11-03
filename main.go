// Copyright 2021 CloudWeGo Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package main

import (
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	"github.com/kitex-contrib/obs-opentelemetry/provider"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	etcd "github.com/kitex-contrib/registry-etcd"
	"log"
	"net"
	"os"
	"sheepim-auth-service/biz/infra/container"
	"sheepim-auth-service/kitex_gen/auth/authservice"
)

func main() {

	env := os.Getenv("ENV")
	if env == "" {
		env = "development"
	}

	App := container.GetContainer(env)

	serviceName := "sheepim-user-service"
	p := provider.NewOpenTelemetryProvider(
		provider.WithServiceName(serviceName),
		provider.WithExportEndpoint(App.Config.OpenTelemetryConfig.Endpoint),
		provider.WithInsecure(),
	)

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
		log.Fatal(err)
	}

	svr := authservice.NewServer(
		new(AuthServiceImpl),
		server.WithSuite(tracing.NewServerSuite()),
		// Please keep the same as provider.WithServiceName
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: serviceName}),
		server.WithRegistry(r),
		server.WithServiceAddr(addr),
	)
	if err := svr.Run(); err != nil {
		klog.Fatalf("server stopped with error:", err)
	}
}
