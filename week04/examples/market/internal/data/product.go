package data

import (
	"context"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/helium12/geek-go/week04/examples/market/internal/biz"
)

type productRepo struct {
	data *Data
	log  *log.Helper
}

// NewProductRepo .
func NewProductRepo(data *Data, logger log.Logger) biz.ProductRepo {
	return &productRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (ar *productRepo) ListProduct(ctx context.Context) ([]*biz.Product, error) {
	ps, err := ar.data.db.Product.Query().All(ctx)
	if err != nil {
		return nil, err
	}
	rv := make([]*biz.Product, 0)
	for _, p := range ps {
		rv = append(rv, &biz.Product{
			Id:        p.ID,
			Name:      p.Name,
			Desc:      p.Desc,
			CreatedAt: p.CreatedAt,
			UpdatedAt: p.UpdatedAt,
		})
	}
	return rv, nil
}

func (ar *productRepo) GetProduct(ctx context.Context, id int64) (*biz.Product, error) {
	p, err := ar.data.db.Product.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return &biz.Product{
		Id:        p.ID,
		Name:      p.Name,
		Desc:      p.Desc,
		CreatedAt: p.CreatedAt,
		UpdatedAt: p.UpdatedAt,
	}, nil
}

func (ar *productRepo) CreateProduct(ctx context.Context, product *biz.Product) error {
	_, err := ar.data.db.Product.
		Create().
		SetName(product.Name).
		SetDesc(product.Desc).
		Save(ctx)
	return err
}

func (ar *productRepo) UpdateProduct(ctx context.Context, id int64, product *biz.Product) error {
	p, err := ar.data.db.Product.Get(ctx, id)
	if err != nil {
		return err
	}
	_, err = p.Update().
		SetName(product.Name).
		SetDesc(product.Desc).
		SetUpdatedAt(time.Now()).
		Save(ctx)
	return err
}

func (ar *productRepo) DeleteProduct(ctx context.Context, id int64) error {
	return ar.data.db.Product.DeleteOneID(id).Exec(ctx)
}
