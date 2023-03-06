// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/testifysec/archivista/ent/predicate"
	"github.com/testifysec/archivista/ent/statement"
	"github.com/testifysec/archivista/ent/subject"
	"github.com/testifysec/archivista/ent/subjectdigest"
)

// SubjectQuery is the builder for querying Subject entities.
type SubjectQuery struct {
	config
	limit              *int
	offset             *int
	unique             *bool
	order              []OrderFunc
	fields             []string
	predicates         []predicate.Subject
	withSubjectDigests *SubjectDigestQuery
	withStatement      *StatementQuery
	withFKs            bool
	modifiers          []func(*sql.Selector)
	loadTotal          []func(context.Context, []*Subject) error
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the SubjectQuery builder.
func (sq *SubjectQuery) Where(ps ...predicate.Subject) *SubjectQuery {
	sq.predicates = append(sq.predicates, ps...)
	return sq
}

// Limit adds a limit step to the query.
func (sq *SubjectQuery) Limit(limit int) *SubjectQuery {
	sq.limit = &limit
	return sq
}

// Offset adds an offset step to the query.
func (sq *SubjectQuery) Offset(offset int) *SubjectQuery {
	sq.offset = &offset
	return sq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (sq *SubjectQuery) Unique(unique bool) *SubjectQuery {
	sq.unique = &unique
	return sq
}

// Order adds an order step to the query.
func (sq *SubjectQuery) Order(o ...OrderFunc) *SubjectQuery {
	sq.order = append(sq.order, o...)
	return sq
}

// QuerySubjectDigests chains the current query on the "subject_digests" edge.
func (sq *SubjectQuery) QuerySubjectDigests() *SubjectDigestQuery {
	query := &SubjectDigestQuery{config: sq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := sq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := sq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(subject.Table, subject.FieldID, selector),
			sqlgraph.To(subjectdigest.Table, subjectdigest.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, subject.SubjectDigestsTable, subject.SubjectDigestsColumn),
		)
		fromU = sqlgraph.SetNeighbors(sq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryStatement chains the current query on the "statement" edge.
func (sq *SubjectQuery) QueryStatement() *StatementQuery {
	query := &StatementQuery{config: sq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := sq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := sq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(subject.Table, subject.FieldID, selector),
			sqlgraph.To(statement.Table, statement.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, subject.StatementTable, subject.StatementColumn),
		)
		fromU = sqlgraph.SetNeighbors(sq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Subject entity from the query.
// Returns a *NotFoundError when no Subject was found.
func (sq *SubjectQuery) First(ctx context.Context) (*Subject, error) {
	nodes, err := sq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{subject.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (sq *SubjectQuery) FirstX(ctx context.Context) *Subject {
	node, err := sq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Subject ID from the query.
// Returns a *NotFoundError when no Subject ID was found.
func (sq *SubjectQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = sq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{subject.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (sq *SubjectQuery) FirstIDX(ctx context.Context) int {
	id, err := sq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Subject entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Subject entity is found.
// Returns a *NotFoundError when no Subject entities are found.
func (sq *SubjectQuery) Only(ctx context.Context) (*Subject, error) {
	nodes, err := sq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{subject.Label}
	default:
		return nil, &NotSingularError{subject.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (sq *SubjectQuery) OnlyX(ctx context.Context) *Subject {
	node, err := sq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Subject ID in the query.
// Returns a *NotSingularError when more than one Subject ID is found.
// Returns a *NotFoundError when no entities are found.
func (sq *SubjectQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = sq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{subject.Label}
	default:
		err = &NotSingularError{subject.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (sq *SubjectQuery) OnlyIDX(ctx context.Context) int {
	id, err := sq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Subjects.
func (sq *SubjectQuery) All(ctx context.Context) ([]*Subject, error) {
	if err := sq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return sq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (sq *SubjectQuery) AllX(ctx context.Context) []*Subject {
	nodes, err := sq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Subject IDs.
func (sq *SubjectQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := sq.Select(subject.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (sq *SubjectQuery) IDsX(ctx context.Context) []int {
	ids, err := sq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (sq *SubjectQuery) Count(ctx context.Context) (int, error) {
	if err := sq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return sq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (sq *SubjectQuery) CountX(ctx context.Context) int {
	count, err := sq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (sq *SubjectQuery) Exist(ctx context.Context) (bool, error) {
	if err := sq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return sq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (sq *SubjectQuery) ExistX(ctx context.Context) bool {
	exist, err := sq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the SubjectQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (sq *SubjectQuery) Clone() *SubjectQuery {
	if sq == nil {
		return nil
	}
	return &SubjectQuery{
		config:             sq.config,
		limit:              sq.limit,
		offset:             sq.offset,
		order:              append([]OrderFunc{}, sq.order...),
		predicates:         append([]predicate.Subject{}, sq.predicates...),
		withSubjectDigests: sq.withSubjectDigests.Clone(),
		withStatement:      sq.withStatement.Clone(),
		// clone intermediate query.
		sql:    sq.sql.Clone(),
		path:   sq.path,
		unique: sq.unique,
	}
}

// WithSubjectDigests tells the query-builder to eager-load the nodes that are connected to
// the "subject_digests" edge. The optional arguments are used to configure the query builder of the edge.
func (sq *SubjectQuery) WithSubjectDigests(opts ...func(*SubjectDigestQuery)) *SubjectQuery {
	query := &SubjectDigestQuery{config: sq.config}
	for _, opt := range opts {
		opt(query)
	}
	sq.withSubjectDigests = query
	return sq
}

// WithStatement tells the query-builder to eager-load the nodes that are connected to
// the "statement" edge. The optional arguments are used to configure the query builder of the edge.
func (sq *SubjectQuery) WithStatement(opts ...func(*StatementQuery)) *SubjectQuery {
	query := &StatementQuery{config: sq.config}
	for _, opt := range opts {
		opt(query)
	}
	sq.withStatement = query
	return sq
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
//	client.Subject.Query().
//		GroupBy(subject.FieldName).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (sq *SubjectQuery) GroupBy(field string, fields ...string) *SubjectGroupBy {
	grbuild := &SubjectGroupBy{config: sq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := sq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return sq.sqlQuery(ctx), nil
	}
	grbuild.label = subject.Label
	grbuild.flds, grbuild.scan = &grbuild.fields, grbuild.Scan
	return grbuild
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
//	client.Subject.Query().
//		Select(subject.FieldName).
//		Scan(ctx, &v)
func (sq *SubjectQuery) Select(fields ...string) *SubjectSelect {
	sq.fields = append(sq.fields, fields...)
	selbuild := &SubjectSelect{SubjectQuery: sq}
	selbuild.label = subject.Label
	selbuild.flds, selbuild.scan = &sq.fields, selbuild.Scan
	return selbuild
}

func (sq *SubjectQuery) prepareQuery(ctx context.Context) error {
	for _, f := range sq.fields {
		if !subject.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if sq.path != nil {
		prev, err := sq.path(ctx)
		if err != nil {
			return err
		}
		sq.sql = prev
	}
	return nil
}

func (sq *SubjectQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Subject, error) {
	var (
		nodes       = []*Subject{}
		withFKs     = sq.withFKs
		_spec       = sq.querySpec()
		loadedTypes = [2]bool{
			sq.withSubjectDigests != nil,
			sq.withStatement != nil,
		}
	)
	if sq.withStatement != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, subject.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Subject).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Subject{config: sq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(sq.modifiers) > 0 {
		_spec.Modifiers = sq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, sq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := sq.withSubjectDigests; query != nil {
		if err := sq.loadSubjectDigests(ctx, query, nodes,
			func(n *Subject) { n.Edges.SubjectDigests = []*SubjectDigest{} },
			func(n *Subject, e *SubjectDigest) { n.Edges.SubjectDigests = append(n.Edges.SubjectDigests, e) }); err != nil {
			return nil, err
		}
	}
	if query := sq.withStatement; query != nil {
		if err := sq.loadStatement(ctx, query, nodes, nil,
			func(n *Subject, e *Statement) { n.Edges.Statement = e }); err != nil {
			return nil, err
		}
	}
	for i := range sq.loadTotal {
		if err := sq.loadTotal[i](ctx, nodes); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (sq *SubjectQuery) loadSubjectDigests(ctx context.Context, query *SubjectDigestQuery, nodes []*Subject, init func(*Subject), assign func(*Subject, *SubjectDigest)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*Subject)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.SubjectDigest(func(s *sql.Selector) {
		s.Where(sql.InValues(subject.SubjectDigestsColumn, fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.subject_subject_digests
		if fk == nil {
			return fmt.Errorf(`foreign-key "subject_subject_digests" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "subject_subject_digests" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (sq *SubjectQuery) loadStatement(ctx context.Context, query *StatementQuery, nodes []*Subject, init func(*Subject), assign func(*Subject, *Statement)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*Subject)
	for i := range nodes {
		if nodes[i].statement_subjects == nil {
			continue
		}
		fk := *nodes[i].statement_subjects
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	query.Where(statement.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "statement_subjects" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (sq *SubjectQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := sq.querySpec()
	if len(sq.modifiers) > 0 {
		_spec.Modifiers = sq.modifiers
	}
	_spec.Node.Columns = sq.fields
	if len(sq.fields) > 0 {
		_spec.Unique = sq.unique != nil && *sq.unique
	}
	return sqlgraph.CountNodes(ctx, sq.driver, _spec)
}

func (sq *SubjectQuery) sqlExist(ctx context.Context) (bool, error) {
	switch _, err := sq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

func (sq *SubjectQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   subject.Table,
			Columns: subject.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: subject.FieldID,
			},
		},
		From:   sq.sql,
		Unique: true,
	}
	if unique := sq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := sq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, subject.FieldID)
		for i := range fields {
			if fields[i] != subject.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := sq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := sq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := sq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := sq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (sq *SubjectQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(sq.driver.Dialect())
	t1 := builder.Table(subject.Table)
	columns := sq.fields
	if len(columns) == 0 {
		columns = subject.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if sq.sql != nil {
		selector = sq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if sq.unique != nil && *sq.unique {
		selector.Distinct()
	}
	for _, p := range sq.predicates {
		p(selector)
	}
	for _, p := range sq.order {
		p(selector)
	}
	if offset := sq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := sq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// SubjectGroupBy is the group-by builder for Subject entities.
type SubjectGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (sgb *SubjectGroupBy) Aggregate(fns ...AggregateFunc) *SubjectGroupBy {
	sgb.fns = append(sgb.fns, fns...)
	return sgb
}

// Scan applies the group-by query and scans the result into the given value.
func (sgb *SubjectGroupBy) Scan(ctx context.Context, v any) error {
	query, err := sgb.path(ctx)
	if err != nil {
		return err
	}
	sgb.sql = query
	return sgb.sqlScan(ctx, v)
}

func (sgb *SubjectGroupBy) sqlScan(ctx context.Context, v any) error {
	for _, f := range sgb.fields {
		if !subject.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := sgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := sgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (sgb *SubjectGroupBy) sqlQuery() *sql.Selector {
	selector := sgb.sql.Select()
	aggregation := make([]string, 0, len(sgb.fns))
	for _, fn := range sgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(sgb.fields)+len(sgb.fns))
		for _, f := range sgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(sgb.fields...)...)
}

// SubjectSelect is the builder for selecting fields of Subject entities.
type SubjectSelect struct {
	*SubjectQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (ss *SubjectSelect) Scan(ctx context.Context, v any) error {
	if err := ss.prepareQuery(ctx); err != nil {
		return err
	}
	ss.sql = ss.SubjectQuery.sqlQuery(ctx)
	return ss.sqlScan(ctx, v)
}

func (ss *SubjectSelect) sqlScan(ctx context.Context, v any) error {
	rows := &sql.Rows{}
	query, args := ss.sql.Query()
	if err := ss.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
