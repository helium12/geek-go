package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"time"
)

type Product struct {
	Id        int64
	Name      string
	Desc      string
	CreatedAt time.Time
	UpdatedAt time.Time
	Sold      int64
}

type ProductRepo interface {
	// db
	ListProduct(ctx context.Context) ([]*Product, error)
	GetProduct(ctx context.Context, id int64) (*Product, error)
	CreateProduct(ctx context.Context, product *Product) error
	UpdateProduct(ctx context.Context, id int64, product *Product) error
	DeleteProduct(ctx context.Context, id int64) error

	// redis
	GetProductSold(ctx context.Context, id int64) (rv int64, err error)
	IncProductSold(ctx context.Context, id int64) error
}

type ProductUsecase struct {
	repo ProductRepo
}

func NewProductUsecase(repo ProductRepo, logger log.Logger) *ProductUsecase {
	return &ProductUsecase{repo: repo}
}

func (uc *ProductUsecase) List(ctx context.Context) (ps []*Product, err error) {
	ps, err = uc.repo.ListProduct(ctx)
	if err != nil {
		return
	}
	return
}

func (uc *ProductUsecase) Get(ctx context.Context, id int64) (p *Product, err error) {
	p, err = uc.repo.GetProduct(ctx, id)
	if err != nil {
		return
	}
	err = uc.repo.IncProductSold(ctx, id)
	if err != nil {
		return
	}
	p.Sold, err = uc.repo.GetProductSold(ctx, id)
	if err != nil {
		return
	}
	return
}

func (uc *ProductUsecase) Create(ctx context.Context, product *Product) error {
	return uc.repo.CreateProduct(ctx, product)
}

func (uc *ProductUsecase) Update(ctx context.Context, id int64, product *Product) error {
	return uc.repo.UpdateProduct(ctx, id, product)
}

func (uc *ProductUsecase) Delete(ctx context.Context, id int64) error {
	return uc.repo.DeleteProduct(ctx, id)
}
