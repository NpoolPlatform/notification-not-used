// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/notification/pkg/db/ent/predicate"
	"github.com/NpoolPlatform/notification/pkg/db/ent/readstate"
)

// ReadStateDelete is the builder for deleting a ReadState entity.
type ReadStateDelete struct {
	config
	hooks    []Hook
	mutation *ReadStateMutation
}

// Where appends a list predicates to the ReadStateDelete builder.
func (rsd *ReadStateDelete) Where(ps ...predicate.ReadState) *ReadStateDelete {
	rsd.mutation.Where(ps...)
	return rsd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (rsd *ReadStateDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(rsd.hooks) == 0 {
		affected, err = rsd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ReadStateMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			rsd.mutation = mutation
			affected, err = rsd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(rsd.hooks) - 1; i >= 0; i-- {
			if rsd.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = rsd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, rsd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (rsd *ReadStateDelete) ExecX(ctx context.Context) int {
	n, err := rsd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (rsd *ReadStateDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: readstate.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: readstate.FieldID,
			},
		},
	}
	if ps := rsd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, rsd.driver, _spec)
}

// ReadStateDeleteOne is the builder for deleting a single ReadState entity.
type ReadStateDeleteOne struct {
	rsd *ReadStateDelete
}

// Exec executes the deletion query.
func (rsdo *ReadStateDeleteOne) Exec(ctx context.Context) error {
	n, err := rsdo.rsd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{readstate.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (rsdo *ReadStateDeleteOne) ExecX(ctx context.Context) {
	rsdo.rsd.ExecX(ctx)
}
