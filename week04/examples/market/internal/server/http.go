package server

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/middleware/validate"
	"github.com/go-kratos/kratos/v2/transport/http"
	v1 "github.com/helium12/geek-go/week04/examples/market/api/market/v1"
	"github.com/helium12/geek-go/week04/examples/market/internal/conf"
	"github.com/helium12/geek-go/week04/examples/market/internal/service"
	"go.opentelemetry.io/otel/trace"
)

// NewHTTPServer new a HTTP server.
func NewHTTPServer(c *conf.Server, logger log.Logger, tracer trace.TracerProvider, market *service.MarketService) *http.Server {
	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
			tracing.Server(tracing.WithTracerProvider(tracer)),
			logging.Server(logger),
			validate.Validator(),
		),
	}
	if c.Http.Network != "" {
		opts = append(opts, http.Network(c.Http.Network))
	}
	if c.Http.Addr != "" {
		opts = append(opts, http.Address(c.Http.Addr))
	}
	if c.Http.Timeout != nil {
		opts = append(opts, http.Timeout(c.Http.Timeout.AsDuration()))
	}
	srv := http.NewServer(opts...)
	v1.RegisterMarketServiceHTTPServer(srv, market)
	return srv
}
