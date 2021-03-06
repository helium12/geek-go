// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/helium12/geek-go/week04/examples/market/internal/data/ent/predicate"
	"github.com/helium12/geek-go/week04/examples/market/internal/data/ent/product"
)

// ProductDelete is the builder for deleting a Product entity.
type ProductDelete struct {
	config
	hooks    []Hook
	mutation *ProductMutation
}

// Where adds a new predicate to the ProductDelete builder.
func (ad *ProductDelete) Where(ps ...predicate.Product) *ProductDelete {
	ad.mutation.predicates = append(ad.mutation.predicates, ps...)
	return ad
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (ad *ProductDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(ad.hooks) == 0 {
		affected, err = ad.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ProductMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ad.mutation = mutation
			affected, err = ad.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ad.hooks) - 1; i >= 0; i-- {
			mut = ad.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ad.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (ad *ProductDelete) ExecX(ctx context.Context) int {
	n, err := ad.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (ad *ProductDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: product.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: product.FieldID,
			},
		},
	}
	if ps := ad.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, ad.driver, _spec)
}

// ProductDeleteOne is the builder for deleting a single Product entity.
type ProductDeleteOne struct {
	ad *ProductDelete
}

// Exec executes the deletion query.
func (ado *ProductDeleteOne) Exec(ctx context.Context) error {
	n, err := ado.ad.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{product.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (ado *ProductDeleteOne) ExecX(ctx context.Context) {
	ado.ad.ExecX(ctx)
}
