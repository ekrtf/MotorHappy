// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/ekrtf/MotorHappy/ent/model"
	"github.com/ekrtf/MotorHappy/ent/predicate"
	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
)

// ModelUpdate is the builder for updating Model entities.
type ModelUpdate struct {
	config
	hooks      []Hook
	mutation   *ModelMutation
	predicates []predicate.Model
}

// Where adds a new predicate for the builder.
func (mu *ModelUpdate) Where(ps ...predicate.Model) *ModelUpdate {
	mu.predicates = append(mu.predicates, ps...)
	return mu
}

// SetName sets the name field.
func (mu *ModelUpdate) SetName(s string) *ModelUpdate {
	mu.mutation.SetName(s)
	return mu
}

// SetNillableName sets the name field if the given value is not nil.
func (mu *ModelUpdate) SetNillableName(s *string) *ModelUpdate {
	if s != nil {
		mu.SetName(*s)
	}
	return mu
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (mu *ModelUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(mu.hooks) == 0 {
		affected, err = mu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ModelMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			mu.mutation = mutation
			affected, err = mu.sqlSave(ctx)
			return affected, err
		})
		for i := len(mu.hooks) - 1; i >= 0; i-- {
			mut = mu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, mu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (mu *ModelUpdate) SaveX(ctx context.Context) int {
	affected, err := mu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (mu *ModelUpdate) Exec(ctx context.Context) error {
	_, err := mu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mu *ModelUpdate) ExecX(ctx context.Context) {
	if err := mu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (mu *ModelUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   model.Table,
			Columns: model.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: model.FieldID,
			},
		},
	}
	if ps := mu.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := mu.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: model.FieldName,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, mu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{model.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// ModelUpdateOne is the builder for updating a single Model entity.
type ModelUpdateOne struct {
	config
	hooks    []Hook
	mutation *ModelMutation
}

// SetName sets the name field.
func (muo *ModelUpdateOne) SetName(s string) *ModelUpdateOne {
	muo.mutation.SetName(s)
	return muo
}

// SetNillableName sets the name field if the given value is not nil.
func (muo *ModelUpdateOne) SetNillableName(s *string) *ModelUpdateOne {
	if s != nil {
		muo.SetName(*s)
	}
	return muo
}

// Save executes the query and returns the updated entity.
func (muo *ModelUpdateOne) Save(ctx context.Context) (*Model, error) {
	var (
		err  error
		node *Model
	)
	if len(muo.hooks) == 0 {
		node, err = muo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ModelMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			muo.mutation = mutation
			node, err = muo.sqlSave(ctx)
			return node, err
		})
		for i := len(muo.hooks) - 1; i >= 0; i-- {
			mut = muo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, muo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (muo *ModelUpdateOne) SaveX(ctx context.Context) *Model {
	m, err := muo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return m
}

// Exec executes the query on the entity.
func (muo *ModelUpdateOne) Exec(ctx context.Context) error {
	_, err := muo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (muo *ModelUpdateOne) ExecX(ctx context.Context) {
	if err := muo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (muo *ModelUpdateOne) sqlSave(ctx context.Context) (m *Model, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   model.Table,
			Columns: model.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: model.FieldID,
			},
		},
	}
	id, ok := muo.mutation.ID()
	if !ok {
		return nil, fmt.Errorf("missing Model.ID for update")
	}
	_spec.Node.ID.Value = id
	if value, ok := muo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: model.FieldName,
		})
	}
	m = &Model{config: muo.config}
	_spec.Assign = m.assignValues
	_spec.ScanValues = m.scanValues()
	if err = sqlgraph.UpdateNode(ctx, muo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{model.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return m, nil
}
