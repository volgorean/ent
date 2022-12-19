// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/entc/integration/edgeschema/ent/predicate"
	"entgo.io/ent/entc/integration/edgeschema/ent/tweettag"
	"entgo.io/ent/schema/field"
)

// TweetTagDelete is the builder for deleting a TweetTag entity.
type TweetTagDelete struct {
	config
	hooks    []Hook
	mutation *TweetTagMutation
}

// Where appends a list predicates to the TweetTagDelete builder.
func (ttd *TweetTagDelete) Where(ps ...predicate.TweetTag) *TweetTagDelete {
	ttd.mutation.Where(ps...)
	return ttd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (ttd *TweetTagDelete) Exec(ctx context.Context) (int, error) {
	return withHooks[int, TweetTagMutation](ctx, ttd.sqlExec, ttd.mutation, ttd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (ttd *TweetTagDelete) ExecX(ctx context.Context) int {
	n, err := ttd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (ttd *TweetTagDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: tweettag.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: tweettag.FieldID,
			},
		},
	}
	if ps := ttd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, ttd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	ttd.mutation.done = true
	return affected, err
}

// TweetTagDeleteOne is the builder for deleting a single TweetTag entity.
type TweetTagDeleteOne struct {
	ttd *TweetTagDelete
}

// Exec executes the deletion query.
func (ttdo *TweetTagDeleteOne) Exec(ctx context.Context) error {
	n, err := ttdo.ttd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{tweettag.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (ttdo *TweetTagDeleteOne) ExecX(ctx context.Context) {
	ttdo.ttd.ExecX(ctx)
}
