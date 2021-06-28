// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"github.com/helium12/geek-go/week04/examples/market/internal/biz"
	"github.com/helium12/geek-go/week04/examples/market/internal/conf"
	"github.com/helium12/geek-go/week04/examples/market/internal/data"
	"github.com/helium12/geek-go/week04/examples/market/internal/server"
	"github.com/helium12/geek-go/week04/examples/market/internal/service"
	"go.opentelemetry.io/otel/trace"
)

// initApp init kratos application.
func initApp(*conf.Server, *conf.Data, trace.TracerProvider, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}
