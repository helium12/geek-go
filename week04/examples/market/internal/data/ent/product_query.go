// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"errors"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/helium12/geek-go/week04/examples/market/internal/data/ent/comment"
	"github.com/helium12/geek-go/week04/examples/market/internal/data/ent/predicate"
	"github.com/helium12/geek-go/week04/examples/market/internal/data/ent/product"
	"github.com/helium12/geek-go/week04/examples/market/internal/data/ent/tag"
)

// ProductQuery is the builder for querying Product entities.
type ProductQuery struct {
	config
	limit      *int
	offset     *int
	order      []OrderFunc
	fields     []string
	predicates []predicate.Product
	// eager-loading edges.
	withComments *CommentQuery
	withTags     *TagQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the ProductQuery builder.
func (aq *ProductQuery) Where(ps ...predicate.Product) *ProductQuery {
	aq.predicates = append(aq.predicates, ps...)
	return aq
}

// Limit adds a limit step to the query.
func (aq *ProductQuery) Limit(limit int) *ProductQuery {
	aq.limit = &limit
	return aq
}

// Offset adds an offset step to the query.
func (aq *ProductQuery) Offset(offset int) *ProductQuery {
	aq.offset = &offset
	return aq
}

// Order adds an order step to the query.
func (aq *ProductQuery) Order(o ...OrderFunc) *ProductQuery {
	aq.order = append(aq.order, o...)
	return aq
}

// QueryComments chains the current query on the "comments" edge.
func (aq *ProductQuery) QueryComments() *CommentQuery {
	query := &CommentQuery{config: aq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := aq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := aq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(product.Table, product.FieldID, selector),
			sqlgraph.To(comment.Table, comment.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, product.CommentsTable, product.CommentsColumn),
		)
		fromU = sqlgraph.SetNeighbors(aq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryTags chains the current query on the "tags" edge.
func (aq *ProductQuery) QueryTags() *TagQuery {
	query := &TagQuery{config: aq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := aq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := aq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(product.Table, product.FieldID, selector),
			sqlgraph.To(tag.Table, tag.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, product.TagsTable, product.TagsPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(aq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Product entity from the query.
// Returns a *NotFoundError when no Product was found.
func (aq *ProductQuery) First(ctx context.Context) (*Product, error) {
	nodes, err := aq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{product.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (aq *ProductQuery) FirstX(ctx context.Context) *Product {
	node, err := aq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Product ID from the query.
// Returns a *NotFoundError when no Product ID was found.
func (aq *ProductQuery) FirstID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = aq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{product.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (aq *ProductQuery) FirstIDX(ctx context.Context) int64 {
	id, err := aq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Product entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when exactly one Product entity is not found.
// Returns a *NotFoundError when no Product entities are found.
func (aq *ProductQuery) Only(ctx context.Context) (*Product, error) {
	nodes, err := aq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{product.Label}
	default:
		return nil, &NotSingularError{product.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (aq *ProductQuery) OnlyX(ctx context.Context) *Product {
	node, err := aq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Product ID in the query.
// Returns a *NotSingularError when exactly one Product ID is not found.
// Returns a *NotFoundError when no entities are found.
func (aq *ProductQuery) OnlyID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = aq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{product.Label}
	default:
		err = &NotSingularError{product.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (aq *ProductQuery) OnlyIDX(ctx context.Context) int64 {
	id, err := aq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Products.
func (aq *ProductQuery) All(ctx context.Context) ([]*Product, error) {
	if err := aq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return aq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (aq *ProductQuery) AllX(ctx context.Context) []*Product {
	nodes, err := aq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Product IDs.
func (aq *ProductQuery) IDs(ctx context.Context) ([]int64, error) {
	var ids []int64
	if err := aq.Select(product.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (aq *ProductQuery) IDsX(ctx context.Context) []int64 {
	ids, err := aq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (aq *ProductQuery) Count(ctx context.Context) (int, error) {
	if err := aq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return aq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (aq *ProductQuery) CountX(ctx context.Context) int {
	count, err := aq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (aq *ProductQuery) Exist(ctx context.Context) (bool, error) {
	if err := aq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return aq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (aq *ProductQuery) ExistX(ctx context.Context) bool {
	exist, err := aq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the ProductQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (aq *ProductQuery) Clone() *ProductQuery {
	if aq == nil {
		return nil
	}
	return &ProductQuery{
		config:       aq.config,
		limit:        aq.limit,
		offset:       aq.offset,
		order:        append([]OrderFunc{}, aq.order...),
		predicates:   append([]predicate.Product{}, aq.predicates...),
		withComments: aq.withComments.Clone(),
		withTags:     aq.withTags.Clone(),
		// clone intermediate query.
		sql:  aq.sql.Clone(),
		path: aq.path,
	}
}

// WithComments tells the query-builder to eager-load the nodes that are connected to
// the "comments" edge. The optional arguments are used to configure the query builder of the edge.
func (aq *ProductQuery) WithComments(opts ...func(*CommentQuery)) *ProductQuery {
	query := &CommentQuery{config: aq.config}
	for _, opt := range opts {
		opt(query)
	}
	aq.withComments = query
	return aq
}

// WithTags tells the query-builder to eager-load the nodes that are connected to
// the "tags" edge. The optional arguments are used to configure the query builder of the edge.
func (aq *ProductQuery) WithTags(opts ...func(*TagQuery)) *ProductQuery {
	query := &TagQuery{config: aq.config}
	for _, opt := range opts {
		opt(query)
	}
	aq.withTags = query
	return aq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Product.Query().
//		GroupBy(product.FieldName).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (aq *ProductQuery) GroupBy(field string, fields ...string) *ProductGroupBy {
	group := &ProductGroupBy{config: aq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := aq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return aq.sqlQuery(ctx), nil
	}
	return group
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//	}
//
//	client.Product.Query().
//		Select(product.FieldName).
//		Scan(ctx, &v)
//
func (aq *ProductQuery) Select(field string, fields ...string) *ProductSelect {
	aq.fields = append([]string{field}, fields...)
	return &ProductSelect{ProductQuery: aq}
}

func (aq *ProductQuery) prepareQuery(ctx context.Context) error {
	for _, f := range aq.fields {
		if !product.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if aq.path != nil {
		prev, err := aq.path(ctx)
		if err != nil {
			return err
		}
		aq.sql = prev
	}
	return nil
}

func (aq *ProductQuery) sqlAll(ctx context.Context) ([]*Product, error) {
	var (
		nodes       = []*Product{}
		_spec       = aq.querySpec()
		loadedTypes = [2]bool{
			aq.withComments != nil,
			aq.withTags != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &Product{config: aq.config}
		nodes = append(nodes, node)
		return node.scanValues(columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("ent: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if err := sqlgraph.QueryNodes(ctx, aq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := aq.withComments; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[int64]*Product)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
			nodes[i].Edges.Comments = []*Comment{}
		}
		query.withFKs = true
		query.Where(predicate.Comment(func(s *sql.Selector) {
			s.Where(sql.InValues(product.CommentsColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.product_comments
			if fk == nil {
				return nil, fmt.Errorf(`foreign-key "product_comments" is nil for node %v`, n.ID)
			}
			node, ok := nodeids[*fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "product_comments" returned %v for node %v`, *fk, n.ID)
			}
			node.Edges.Comments = append(node.Edges.Comments, n)
		}
	}

	if query := aq.withTags; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		ids := make(map[int64]*Product, len(nodes))
		for _, node := range nodes {
			ids[node.ID] = node
			fks = append(fks, node.ID)
			node.Edges.Tags = []*Tag{}
		}
		var (
			edgeids []int64
			edges   = make(map[int64][]*Product)
		)
		_spec := &sqlgraph.EdgeQuerySpec{
			Edge: &sqlgraph.EdgeSpec{
				Inverse: true,
				Table:   product.TagsTable,
				Columns: product.TagsPrimaryKey,
			},
			Predicate: func(s *sql.Selector) {
				s.Where(sql.InValues(product.TagsPrimaryKey[1], fks...))
			},

			ScanValues: func() [2]interface{} {
				return [2]interface{}{&sql.NullInt64{}, &sql.NullInt64{}}
			},
			Assign: func(out, in interface{}) error {
				eout, ok := out.(*sql.NullInt64)
				if !ok || eout == nil {
					return fmt.Errorf("unexpected id value for edge-out")
				}
				ein, ok := in.(*sql.NullInt64)
				if !ok || ein == nil {
					return fmt.Errorf("unexpected id value for edge-in")
				}
				outValue := eout.Int64
				inValue := ein.Int64
				node, ok := ids[outValue]
				if !ok {
					return fmt.Errorf("unexpected node id in edges: %v", outValue)
				}
				edgeids = append(edgeids, inValue)
				edges[inValue] = append(edges[inValue], node)
				return nil
			},
		}
		if err := sqlgraph.QueryEdges(ctx, aq.driver, _spec); err != nil {
			return nil, fmt.Errorf(`query edges "tags": %v`, err)
		}
		query.Where(tag.IDIn(edgeids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := edges[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected "tags" node returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Tags = append(nodes[i].Edges.Tags, n)
			}
		}
	}

	return nodes, nil
}

func (aq *ProductQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := aq.querySpec()
	return sqlgraph.CountNodes(ctx, aq.driver, _spec)
}

func (aq *ProductQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := aq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %v", err)
	}
	return n > 0, nil
}

func (aq *ProductQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   product.Table,
			Columns: product.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: product.FieldID,
			},
		},
		From:   aq.sql,
		Unique: true,
	}
	if fields := aq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, product.FieldID)
		for i := range fields {
			if fields[i] != product.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := aq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := aq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := aq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := aq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector, product.ValidColumn)
			}
		}
	}
	return _spec
}

func (aq *ProductQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(aq.driver.Dialect())
	t1 := builder.Table(product.Table)
	selector := builder.Select(t1.Columns(product.Columns...)...).From(t1)
	if aq.sql != nil {
		selector = aq.sql
		selector.Select(selector.Columns(product.Columns...)...)
	}
	for _, p := range aq.predicates {
		p(selector)
	}
	for _, p := range aq.order {
		p(selector, product.ValidColumn)
	}
	if offset := aq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := aq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ProductGroupBy is the group-by builder for Product entities.
type ProductGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (agb *ProductGroupBy) Aggregate(fns ...AggregateFunc) *ProductGroupBy {
	agb.fns = append(agb.fns, fns...)
	return agb
}

// Scan applies the group-by query and scans the result into the given value.
func (agb *ProductGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := agb.path(ctx)
	if err != nil {
		return err
	}
	agb.sql = query
	return agb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (agb *ProductGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := agb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (agb *ProductGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(agb.fields) > 1 {
		return nil, errors.New("ent: ProductGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := agb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (agb *ProductGroupBy) StringsX(ctx context.Context) []string {
	v, err := agb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (agb *ProductGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = agb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{product.Label}
	default:
		err = fmt.Errorf("ent: ProductGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (agb *ProductGroupBy) StringX(ctx context.Context) string {
	v, err := agb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (agb *ProductGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(agb.fields) > 1 {
		return nil, errors.New("ent: ProductGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := agb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (agb *ProductGroupBy) IntsX(ctx context.Context) []int {
	v, err := agb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (agb *ProductGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = agb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{product.Label}
	default:
		err = fmt.Errorf("ent: ProductGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (agb *ProductGroupBy) IntX(ctx context.Context) int {
	v, err := agb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (agb *ProductGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(agb.fields) > 1 {
		return nil, errors.New("ent: ProductGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := agb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (agb *ProductGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := agb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (agb *ProductGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = agb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{product.Label}
	default:
		err = fmt.Errorf("ent: ProductGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (agb *ProductGroupBy) Float64X(ctx context.Context) float64 {
	v, err := agb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (agb *ProductGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(agb.fields) > 1 {
		return nil, errors.New("ent: ProductGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := agb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (agb *ProductGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := agb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (agb *ProductGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = agb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{product.Label}
	default:
		err = fmt.Errorf("ent: ProductGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (agb *ProductGroupBy) BoolX(ctx context.Context) bool {
	v, err := agb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (agb *ProductGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range agb.fields {
		if !product.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := agb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := agb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (agb *ProductGroupBy) sqlQuery() *sql.Selector {
	selector := agb.sql
	columns := make([]string, 0, len(agb.fields)+len(agb.fns))
	columns = append(columns, agb.fields...)
	for _, fn := range agb.fns {
		columns = append(columns, fn(selector, product.ValidColumn))
	}
	return selector.Select(columns...).GroupBy(agb.fields...)
}

// ProductSelect is the builder for selecting fields of Product entities.
type ProductSelect struct {
	*ProductQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (as *ProductSelect) Scan(ctx context.Context, v interface{}) error {
	if err := as.prepareQuery(ctx); err != nil {
		return err
	}
	as.sql = as.ProductQuery.sqlQuery(ctx)
	return as.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (as *ProductSelect) ScanX(ctx context.Context, v interface{}) {
	if err := as.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (as *ProductSelect) Strings(ctx context.Context) ([]string, error) {
	if len(as.fields) > 1 {
		return nil, errors.New("ent: ProductSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := as.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (as *ProductSelect) StringsX(ctx context.Context) []string {
	v, err := as.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (as *ProductSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = as.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{product.Label}
	default:
		err = fmt.Errorf("ent: ProductSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (as *ProductSelect) StringX(ctx context.Context) string {
	v, err := as.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (as *ProductSelect) Ints(ctx context.Context) ([]int, error) {
	if len(as.fields) > 1 {
		return nil, errors.New("ent: ProductSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := as.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (as *ProductSelect) IntsX(ctx context.Context) []int {
	v, err := as.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (as *ProductSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = as.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{product.Label}
	default:
		err = fmt.Errorf("ent: ProductSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (as *ProductSelect) IntX(ctx context.Context) int {
	v, err := as.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (as *ProductSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(as.fields) > 1 {
		return nil, errors.New("ent: ProductSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := as.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (as *ProductSelect) Float64sX(ctx context.Context) []float64 {
	v, err := as.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (as *ProductSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = as.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{product.Label}
	default:
		err = fmt.Errorf("ent: ProductSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (as *ProductSelect) Float64X(ctx context.Context) float64 {
	v, err := as.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (as *ProductSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(as.fields) > 1 {
		return nil, errors.New("ent: ProductSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := as.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (as *ProductSelect) BoolsX(ctx context.Context) []bool {
	v, err := as.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (as *ProductSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = as.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{product.Label}
	default:
		err = fmt.Errorf("ent: ProductSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (as *ProductSelect) BoolX(ctx context.Context) bool {
	v, err := as.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (as *ProductSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := as.sqlQuery().Query()
	if err := as.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (as *ProductSelect) sqlQuery() sql.Querier {
	selector := as.sql
	selector.Select(selector.Columns(as.fields...)...)
	return selector
}
