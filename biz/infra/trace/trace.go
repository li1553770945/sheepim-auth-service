package trace

import (
	"github.com/hertz-contrib/obs-opentelemetry/provider"
	"sheepim-auth-service/biz/infra/config"
)

func InitTrace(config *config.Config) provider.OtelProvider {
	p := provider.NewOpenTelemetryProvider(
		provider.WithServiceName(config.ServerConfig.ServiceName),
		provider.WithExportEndpoint(config.OpenTelemetryConfig.Endpoint),
		provider.WithInsecure(),
	)
	return p

}
