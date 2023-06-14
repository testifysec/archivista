// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/testifysec/archivista/ent/attestation"
	"github.com/testifysec/archivista/ent/attestationcollection"
	"github.com/testifysec/archivista/ent/statement"
)

// AttestationCollectionCreate is the builder for creating a AttestationCollection entity.
type AttestationCollectionCreate struct {
	config
	mutation *AttestationCollectionMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (acc *AttestationCollectionCreate) SetName(s string) *AttestationCollectionCreate {
	acc.mutation.SetName(s)
	return acc
}

// AddAttestationIDs adds the "attestations" edge to the Attestation entity by IDs.
func (acc *AttestationCollectionCreate) AddAttestationIDs(ids ...int) *AttestationCollectionCreate {
	acc.mutation.AddAttestationIDs(ids...)
	return acc
}

// AddAttestations adds the "attestations" edges to the Attestation entity.
func (acc *AttestationCollectionCreate) AddAttestations(a ...*Attestation) *AttestationCollectionCreate {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return acc.AddAttestationIDs(ids...)
}

// SetStatementID sets the "statement" edge to the Statement entity by ID.
func (acc *AttestationCollectionCreate) SetStatementID(id int) *AttestationCollectionCreate {
	acc.mutation.SetStatementID(id)
	return acc
}

// SetStatement sets the "statement" edge to the Statement entity.
func (acc *AttestationCollectionCreate) SetStatement(s *Statement) *AttestationCollectionCreate {
	return acc.SetStatementID(s.ID)
}

// Mutation returns the AttestationCollectionMutation object of the builder.
func (acc *AttestationCollectionCreate) Mutation() *AttestationCollectionMutation {
	return acc.mutation
}

// Save creates the AttestationCollection in the database.
func (acc *AttestationCollectionCreate) Save(ctx context.Context) (*AttestationCollection, error) {
	var (
		err  error
		node *AttestationCollection
	)
	if len(acc.hooks) == 0 {
		if err = acc.check(); err != nil {
			return nil, err
		}
		node, err = acc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AttestationCollectionMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = acc.check(); err != nil {
				return nil, err
			}
			acc.mutation = mutation
			if node, err = acc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(acc.hooks) - 1; i >= 0; i-- {
			if acc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = acc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, acc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*AttestationCollection)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from AttestationCollectionMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (acc *AttestationCollectionCreate) SaveX(ctx context.Context) *AttestationCollection {
	v, err := acc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (acc *AttestationCollectionCreate) Exec(ctx context.Context) error {
	_, err := acc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (acc *AttestationCollectionCreate) ExecX(ctx context.Context) {
	if err := acc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (acc *AttestationCollectionCreate) check() error {
	if _, ok := acc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "AttestationCollection.name"`)}
	}
	if v, ok := acc.mutation.Name(); ok {
		if err := attestationcollection.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "AttestationCollection.name": %w`, err)}
		}
	}
	if _, ok := acc.mutation.StatementID(); !ok {
		return &ValidationError{Name: "statement", err: errors.New(`ent: missing required edge "AttestationCollection.statement"`)}
	}
	return nil
}

func (acc *AttestationCollectionCreate) sqlSave(ctx context.Context) (*AttestationCollection, error) {
	_node, _spec := acc.createSpec()
	if err := sqlgraph.CreateNode(ctx, acc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (acc *AttestationCollectionCreate) createSpec() (*AttestationCollection, *sqlgraph.CreateSpec) {
	var (
		_node = &AttestationCollection{config: acc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: attestationcollection.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: attestationcollection.FieldID,
			},
		}
	)
	if value, ok := acc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: attestationcollection.FieldName,
		})
		_node.Name = value
	}
	if nodes := acc.mutation.AttestationsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   attestationcollection.AttestationsTable,
			Columns: []string{attestationcollection.AttestationsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: attestation.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := acc.mutation.StatementIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   attestationcollection.StatementTable,
			Columns: []string{attestationcollection.StatementColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: statement.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.statement_attestation_collections = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// AttestationCollectionCreateBulk is the builder for creating many AttestationCollection entities in bulk.
type AttestationCollectionCreateBulk struct {
	config
	builders []*AttestationCollectionCreate
}

// Save creates the AttestationCollection entities in the database.
func (accb *AttestationCollectionCreateBulk) Save(ctx context.Context) ([]*AttestationCollection, error) {
	specs := make([]*sqlgraph.CreateSpec, len(accb.builders))
	nodes := make([]*AttestationCollection, len(accb.builders))
	mutators := make([]Mutator, len(accb.builders))
	for i := range accb.builders {
		func(i int, root context.Context) {
			builder := accb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*AttestationCollectionMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, accb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, accb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, accb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (accb *AttestationCollectionCreateBulk) SaveX(ctx context.Context) []*AttestationCollection {
	v, err := accb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (accb *AttestationCollectionCreateBulk) Exec(ctx context.Context) error {
	_, err := accb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (accb *AttestationCollectionCreateBulk) ExecX(ctx context.Context) {
	if err := accb.Exec(ctx); err != nil {
		panic(err)
	}
}
