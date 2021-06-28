package service

import (
	"context"

	"go.opentelemetry.io/otel"

	pb "github.com/helium12/geek-go/week04/examples/market/api/market/v1"
	"github.com/helium12/geek-go/week04/examples/market/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

func NewMarketService(product *biz.ProductUsecase, logger log.Logger) *MarketService {
	return &MarketService{
		product: product,
		log:     log.NewHelper(logger),
	}
}

func (s *MarketService) CreateProduct(ctx context.Context, req *pb.CreateProductRequest) (*pb.CreateProductReply, error) {
	s.log.Infof("input data %v", req)
	err := s.product.Create(ctx, &biz.Product{
		Name: req.Name,
		Desc: req.Desc,
	})
	return &pb.CreateProductReply{}, err
}

func (s *MarketService) UpdateProduct(ctx context.Context, req *pb.UpdateProductRequest) (*pb.UpdateProductReply, error) {
	s.log.Infof("input data %v", req)
	err := s.product.Update(ctx, req.Id, &biz.Product{
		Name: req.Name,
		Desc: req.Desc,
	})
	return &pb.UpdateProductReply{}, err
}

func (s *MarketService) DeleteProduct(ctx context.Context, req *pb.DeleteProductRequest) (*pb.DeleteProductReply, error) {
	s.log.Infof("input data %v", req)
	err := s.product.Delete(ctx, req.Id)
	return &pb.DeleteProductReply{}, err
}

func (s *MarketService) GetProduct(ctx context.Context, req *pb.GetProductRequest) (*pb.GetProductReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "GetProduct")
	defer span.End()
	p, err := s.product.Get(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &pb.GetProductReply{Product: &pb.Product{Id: p.Id, Name: p.Name, Desc: p.Desc, Sold: p.Sold}}, nil
}

func (s *MarketService) ListProduct(ctx context.Context, req *pb.ListProductRequest) (*pb.ListProductReply, error) {
	ps, err := s.product.List(ctx)
	reply := &pb.ListProductReply{}
	for _, p := range ps {
		reply.Results = append(reply.Results, &pb.Product{
			Id:   p.Id,
			Name: p.Name,
			Desc: p.Desc,
		})
	}
	return reply, err
}
