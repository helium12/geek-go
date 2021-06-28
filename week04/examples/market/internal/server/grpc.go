package server

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/middleware/validate"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	v1 "github.com/helium12/geek-go/week04/examples/market/api/market/v1"
	"github.com/helium12/geek-go/week04/examples/market/internal/conf"
	"github.com/helium12/geek-go/week04/examples/market/internal/service"
	"go.opentelemetry.io/otel/trace"
)

// NewGRPCServer new a gRPC server.
func NewGRPCServer(c *conf.Server, logger log.Logger, tracer trace.TracerProvider, market *service.MarketService) *grpc.Server {
	var opts = []grpc.ServerOption{
		grpc.Middleware(
			recovery.Recovery(),
			tracing.Server(tracing.WithTracerProvider(tracer)),
			logging.Server(logger),
			validate.Validator(),
		),
	}
	if c.Grpc.Network != "" {
		opts = append(opts, grpc.Network(c.Grpc.Network))
	}
	if c.Grpc.Addr != "" {
		opts = append(opts, grpc.Address(c.Grpc.Addr))
	}
	if c.Grpc.Timeout != nil {
		opts = append(opts, grpc.Timeout(c.Grpc.Timeout.AsDuration()))
	}
	srv := grpc.NewServer(opts...)
	v1.RegisterMarketServiceServer(srv, market)
	return srv
}
