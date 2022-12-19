// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package entv2

import (
	"context"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/entc/integration/migrate/entv2/customtype"
	"entgo.io/ent/schema/field"
)

// CustomTypeCreate is the builder for creating a CustomType entity.
type CustomTypeCreate struct {
	config
	mutation *CustomTypeMutation
	hooks    []Hook
}

// SetCustom sets the "custom" field.
func (ctc *CustomTypeCreate) SetCustom(s string) *CustomTypeCreate {
	ctc.mutation.SetCustom(s)
	return ctc
}

// SetNillableCustom sets the "custom" field if the given value is not nil.
func (ctc *CustomTypeCreate) SetNillableCustom(s *string) *CustomTypeCreate {
	if s != nil {
		ctc.SetCustom(*s)
	}
	return ctc
}

// SetTz0 sets the "tz0" field.
func (ctc *CustomTypeCreate) SetTz0(t time.Time) *CustomTypeCreate {
	ctc.mutation.SetTz0(t)
	return ctc
}

// SetNillableTz0 sets the "tz0" field if the given value is not nil.
func (ctc *CustomTypeCreate) SetNillableTz0(t *time.Time) *CustomTypeCreate {
	if t != nil {
		ctc.SetTz0(*t)
	}
	return ctc
}

// SetTz3 sets the "tz3" field.
func (ctc *CustomTypeCreate) SetTz3(t time.Time) *CustomTypeCreate {
	ctc.mutation.SetTz3(t)
	return ctc
}

// SetNillableTz3 sets the "tz3" field if the given value is not nil.
func (ctc *CustomTypeCreate) SetNillableTz3(t *time.Time) *CustomTypeCreate {
	if t != nil {
		ctc.SetTz3(*t)
	}
	return ctc
}

// Mutation returns the CustomTypeMutation object of the builder.
func (ctc *CustomTypeCreate) Mutation() *CustomTypeMutation {
	return ctc.mutation
}

// Save creates the CustomType in the database.
func (ctc *CustomTypeCreate) Save(ctx context.Context) (*CustomType, error) {
	return withHooks[*CustomType, CustomTypeMutation](ctx, ctc.sqlSave, ctc.mutation, ctc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (ctc *CustomTypeCreate) SaveX(ctx context.Context) *CustomType {
	v, err := ctc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ctc *CustomTypeCreate) Exec(ctx context.Context) error {
	_, err := ctc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ctc *CustomTypeCreate) ExecX(ctx context.Context) {
	if err := ctc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ctc *CustomTypeCreate) check() error {
	return nil
}

func (ctc *CustomTypeCreate) sqlSave(ctx context.Context) (*CustomType, error) {
	if err := ctc.check(); err != nil {
		return nil, err
	}
	_node, _spec := ctc.createSpec()
	if err := sqlgraph.CreateNode(ctx, ctc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	ctc.mutation.id = &_node.ID
	ctc.mutation.done = true
	return _node, nil
}

func (ctc *CustomTypeCreate) createSpec() (*CustomType, *sqlgraph.CreateSpec) {
	var (
		_node = &CustomType{config: ctc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: customtype.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: customtype.FieldID,
			},
		}
	)
	if value, ok := ctc.mutation.Custom(); ok {
		_spec.SetField(customtype.FieldCustom, field.TypeString, value)
		_node.Custom = value
	}
	if value, ok := ctc.mutation.Tz0(); ok {
		_spec.SetField(customtype.FieldTz0, field.TypeTime, value)
		_node.Tz0 = value
	}
	if value, ok := ctc.mutation.Tz3(); ok {
		_spec.SetField(customtype.FieldTz3, field.TypeTime, value)
		_node.Tz3 = value
	}
	return _node, _spec
}

// CustomTypeCreateBulk is the builder for creating many CustomType entities in bulk.
type CustomTypeCreateBulk struct {
	config
	builders []*CustomTypeCreate
}

// Save creates the CustomType entities in the database.
func (ctcb *CustomTypeCreateBulk) Save(ctx context.Context) ([]*CustomType, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ctcb.builders))
	nodes := make([]*CustomType, len(ctcb.builders))
	mutators := make([]Mutator, len(ctcb.builders))
	for i := range ctcb.builders {
		func(i int, root context.Context) {
			builder := ctcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*CustomTypeMutation)
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
					_, err = mutators[i+1].Mutate(root, ctcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ctcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, ctcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ctcb *CustomTypeCreateBulk) SaveX(ctx context.Context) []*CustomType {
	v, err := ctcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ctcb *CustomTypeCreateBulk) Exec(ctx context.Context) error {
	_, err := ctcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ctcb *CustomTypeCreateBulk) ExecX(ctx context.Context) {
	if err := ctcb.Exec(ctx); err != nil {
		panic(err)
	}
}
