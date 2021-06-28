// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/helium12/geek-go/week04/examples/market/internal/data/ent/comment"
	"github.com/helium12/geek-go/week04/examples/market/internal/data/ent/predicate"
	"github.com/helium12/geek-go/week04/examples/market/internal/data/ent/product"
	"github.com/helium12/geek-go/week04/examples/market/internal/data/ent/tag"
)

// ProductUpdate is the builder for updating Product entities.
type ProductUpdate struct {
	config
	hooks    []Hook
	mutation *ProductMutation
}

// Where adds a new predicate for the ProductUpdate builder.
func (au *ProductUpdate) Where(ps ...predicate.Product) *ProductUpdate {
	au.mutation.predicates = append(au.mutation.predicates, ps...)
	return au
}

// SetName sets the "name" field.
func (au *ProductUpdate) SetName(s string) *ProductUpdate {
	au.mutation.SetName(s)
	return au
}

// SetDesc sets the "desc" field.
func (au *ProductUpdate) SetDesc(s string) *ProductUpdate {
	au.mutation.SetDesc(s)
	return au
}

// SetCreatedAt sets the "created_at" field.
func (au *ProductUpdate) SetCreatedAt(t time.Time) *ProductUpdate {
	au.mutation.SetCreatedAt(t)
	return au
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (au *ProductUpdate) SetNillableCreatedAt(t *time.Time) *ProductUpdate {
	if t != nil {
		au.SetCreatedAt(*t)
	}
	return au
}

// SetUpdatedAt sets the "updated_at" field.
func (au *ProductUpdate) SetUpdatedAt(t time.Time) *ProductUpdate {
	au.mutation.SetUpdatedAt(t)
	return au
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (au *ProductUpdate) SetNillableUpdatedAt(t *time.Time) *ProductUpdate {
	if t != nil {
		au.SetUpdatedAt(*t)
	}
	return au
}

// AddCommentIDs adds the "comments" edge to the Comment entity by IDs.
func (au *ProductUpdate) AddCommentIDs(ids ...int64) *ProductUpdate {
	au.mutation.AddCommentIDs(ids...)
	return au
}

// AddComments adds the "comments" edges to the Comment entity.
func (au *ProductUpdate) AddComments(c ...*Comment) *ProductUpdate {
	ids := make([]int64, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return au.AddCommentIDs(ids...)
}

// AddTagIDs adds the "tags" edge to the Tag entity by IDs.
func (au *ProductUpdate) AddTagIDs(ids ...int64) *ProductUpdate {
	au.mutation.AddTagIDs(ids...)
	return au
}

// AddTags adds the "tags" edges to the Tag entity.
func (au *ProductUpdate) AddTags(t ...*Tag) *ProductUpdate {
	ids := make([]int64, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return au.AddTagIDs(ids...)
}

// Mutation returns the ProductMutation object of the builder.
func (au *ProductUpdate) Mutation() *ProductMutation {
	return au.mutation
}

// ClearComments clears all "comments" edges to the Comment entity.
func (au *ProductUpdate) ClearComments() *ProductUpdate {
	au.mutation.ClearComments()
	return au
}

// RemoveCommentIDs removes the "comments" edge to Comment entities by IDs.
func (au *ProductUpdate) RemoveCommentIDs(ids ...int64) *ProductUpdate {
	au.mutation.RemoveCommentIDs(ids...)
	return au
}

// RemoveComments removes "comments" edges to Comment entities.
func (au *ProductUpdate) RemoveComments(c ...*Comment) *ProductUpdate {
	ids := make([]int64, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return au.RemoveCommentIDs(ids...)
}

// ClearTags clears all "tags" edges to the Tag entity.
func (au *ProductUpdate) ClearTags() *ProductUpdate {
	au.mutation.ClearTags()
	return au
}

// RemoveTagIDs removes the "tags" edge to Tag entities by IDs.
func (au *ProductUpdate) RemoveTagIDs(ids ...int64) *ProductUpdate {
	au.mutation.RemoveTagIDs(ids...)
	return au
}

// RemoveTags removes "tags" edges to Tag entities.
func (au *ProductUpdate) RemoveTags(t ...*Tag) *ProductUpdate {
	ids := make([]int64, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return au.RemoveTagIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (au *ProductUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(au.hooks) == 0 {
		affected, err = au.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ProductMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			au.mutation = mutation
			affected, err = au.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(au.hooks) - 1; i >= 0; i-- {
			mut = au.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, au.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (au *ProductUpdate) SaveX(ctx context.Context) int {
	affected, err := au.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (au *ProductUpdate) Exec(ctx context.Context) error {
	_, err := au.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (au *ProductUpdate) ExecX(ctx context.Context) {
	if err := au.Exec(ctx); err != nil {
		panic(err)
	}
}

func (au *ProductUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   product.Table,
			Columns: product.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: product.FieldID,
			},
		},
	}
	if ps := au.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := au.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: product.FieldName,
		})
	}
	if value, ok := au.mutation.Desc(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: product.FieldDesc,
		})
	}
	if value, ok := au.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: product.FieldCreatedAt,
		})
	}
	if value, ok := au.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: product.FieldUpdatedAt,
		})
	}
	if au.mutation.CommentsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   product.CommentsTable,
			Columns: []string{product.CommentsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: comment.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := au.mutation.RemovedCommentsIDs(); len(nodes) > 0 && !au.mutation.CommentsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   product.CommentsTable,
			Columns: []string{product.CommentsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: comment.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := au.mutation.CommentsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   product.CommentsTable,
			Columns: []string{product.CommentsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: comment.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if au.mutation.TagsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   product.TagsTable,
			Columns: product.TagsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: tag.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := au.mutation.RemovedTagsIDs(); len(nodes) > 0 && !au.mutation.TagsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   product.TagsTable,
			Columns: product.TagsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: tag.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := au.mutation.TagsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   product.TagsTable,
			Columns: product.TagsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: tag.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, au.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{product.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// ProductUpdateOne is the builder for updating a single Product entity.
type ProductUpdateOne struct {
	config
	hooks    []Hook
	mutation *ProductMutation
}

// SetName sets the "name" field.
func (auo *ProductUpdateOne) SetName(s string) *ProductUpdateOne {
	auo.mutation.SetName(s)
	return auo
}

// SetDesc sets the "desc" field.
func (auo *ProductUpdateOne) SetDesc(s string) *ProductUpdateOne {
	auo.mutation.SetDesc(s)
	return auo
}

// SetCreatedAt sets the "created_at" field.
func (auo *ProductUpdateOne) SetCreatedAt(t time.Time) *ProductUpdateOne {
	auo.mutation.SetCreatedAt(t)
	return auo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (auo *ProductUpdateOne) SetNillableCreatedAt(t *time.Time) *ProductUpdateOne {
	if t != nil {
		auo.SetCreatedAt(*t)
	}
	return auo
}

// SetUpdatedAt sets the "updated_at" field.
func (auo *ProductUpdateOne) SetUpdatedAt(t time.Time) *ProductUpdateOne {
	auo.mutation.SetUpdatedAt(t)
	return auo
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (auo *ProductUpdateOne) SetNillableUpdatedAt(t *time.Time) *ProductUpdateOne {
	if t != nil {
		auo.SetUpdatedAt(*t)
	}
	return auo
}

// AddCommentIDs adds the "comments" edge to the Comment entity by IDs.
func (auo *ProductUpdateOne) AddCommentIDs(ids ...int64) *ProductUpdateOne {
	auo.mutation.AddCommentIDs(ids...)
	return auo
}

// AddComments adds the "comments" edges to the Comment entity.
func (auo *ProductUpdateOne) AddComments(c ...*Comment) *ProductUpdateOne {
	ids := make([]int64, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return auo.AddCommentIDs(ids...)
}

// AddTagIDs adds the "tags" edge to the Tag entity by IDs.
func (auo *ProductUpdateOne) AddTagIDs(ids ...int64) *ProductUpdateOne {
	auo.mutation.AddTagIDs(ids...)
	return auo
}

// AddTags adds the "tags" edges to the Tag entity.
func (auo *ProductUpdateOne) AddTags(t ...*Tag) *ProductUpdateOne {
	ids := make([]int64, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return auo.AddTagIDs(ids...)
}

// Mutation returns the ProductMutation object of the builder.
func (auo *ProductUpdateOne) Mutation() *ProductMutation {
	return auo.mutation
}

// ClearComments clears all "comments" edges to the Comment entity.
func (auo *ProductUpdateOne) ClearComments() *ProductUpdateOne {
	auo.mutation.ClearComments()
	return auo
}

// RemoveCommentIDs removes the "comments" edge to Comment entities by IDs.
func (auo *ProductUpdateOne) RemoveCommentIDs(ids ...int64) *ProductUpdateOne {
	auo.mutation.RemoveCommentIDs(ids...)
	return auo
}

// RemoveComments removes "comments" edges to Comment entities.
func (auo *ProductUpdateOne) RemoveComments(c ...*Comment) *ProductUpdateOne {
	ids := make([]int64, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return auo.RemoveCommentIDs(ids...)
}

// ClearTags clears all "tags" edges to the Tag entity.
func (auo *ProductUpdateOne) ClearTags() *ProductUpdateOne {
	auo.mutation.ClearTags()
	return auo
}

// RemoveTagIDs removes the "tags" edge to Tag entities by IDs.
func (auo *ProductUpdateOne) RemoveTagIDs(ids ...int64) *ProductUpdateOne {
	auo.mutation.RemoveTagIDs(ids...)
	return auo
}

// RemoveTags removes "tags" edges to Tag entities.
func (auo *ProductUpdateOne) RemoveTags(t ...*Tag) *ProductUpdateOne {
	ids := make([]int64, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return auo.RemoveTagIDs(ids...)
}

// Save executes the query and returns the updated Product entity.
func (auo *ProductUpdateOne) Save(ctx context.Context) (*Product, error) {
	var (
		err  error
		node *Product
	)
	if len(auo.hooks) == 0 {
		node, err = auo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ProductMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			auo.mutation = mutation
			node, err = auo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(auo.hooks) - 1; i >= 0; i-- {
			mut = auo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, auo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (auo *ProductUpdateOne) SaveX(ctx context.Context) *Product {
	node, err := auo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (auo *ProductUpdateOne) Exec(ctx context.Context) error {
	_, err := auo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (auo *ProductUpdateOne) ExecX(ctx context.Context) {
	if err := auo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (auo *ProductUpdateOne) sqlSave(ctx context.Context) (_node *Product, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   product.Table,
			Columns: product.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: product.FieldID,
			},
		},
	}
	id, ok := auo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Product.ID for update")}
	}
	_spec.Node.ID.Value = id
	if ps := auo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := auo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: product.FieldName,
		})
	}
	if value, ok := auo.mutation.Desc(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: product.FieldDesc,
		})
	}
	if value, ok := auo.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: product.FieldCreatedAt,
		})
	}
	if value, ok := auo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: product.FieldUpdatedAt,
		})
	}
	if auo.mutation.CommentsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   product.CommentsTable,
			Columns: []string{product.CommentsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: comment.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := auo.mutation.RemovedCommentsIDs(); len(nodes) > 0 && !auo.mutation.CommentsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   product.CommentsTable,
			Columns: []string{product.CommentsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: comment.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := auo.mutation.CommentsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   product.CommentsTable,
			Columns: []string{product.CommentsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: comment.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if auo.mutation.TagsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   product.TagsTable,
			Columns: product.TagsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: tag.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := auo.mutation.RemovedTagsIDs(); len(nodes) > 0 && !auo.mutation.TagsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   product.TagsTable,
			Columns: product.TagsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: tag.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := auo.mutation.TagsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   product.TagsTable,
			Columns: product.TagsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: tag.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Product{config: auo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, auo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{product.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}
