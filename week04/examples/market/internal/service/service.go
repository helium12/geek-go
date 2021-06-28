package service

import (
	pb "github.com/helium12/geek-go/week04/examples/market/api/market/v1"
	"github.com/helium12/geek-go/week04/examples/market/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewMarketService)

type MarketService struct {
	pb.UnimplementedMarketServiceServer

	log *log.Helper

	product *biz.ProductUsecase
}
