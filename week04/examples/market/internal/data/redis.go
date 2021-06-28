package data

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
)

func soldKey(id int64) string {
	return fmt.Sprintf("sold:%d", id)
}

func (ar *productRepo) GetProductSold(ctx context.Context, id int64) (rv int64, err error) {
	get := ar.data.rdb.Get(ctx, soldKey(id))
	rv, err = get.Int64()
	if err == redis.Nil {
		return 0, nil
	}
	return
}

func (ar *productRepo) IncProductSold(ctx context.Context, id int64) error {
	_, err := ar.data.rdb.Incr(ctx, soldKey(id)).Result()
	return err
}
