// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/notification/pkg/db/ent/predicate"
	"github.com/NpoolPlatform/notification/pkg/db/ent/readuser"
	"github.com/google/uuid"
)

// ReadUserUpdate is the builder for updating ReadUser entities.
type ReadUserUpdate struct {
	config
	hooks    []Hook
	mutation *ReadUserMutation
}

// Where appends a list predicates to the ReadUserUpdate builder.
func (ruu *ReadUserUpdate) Where(ps ...predicate.ReadUser) *ReadUserUpdate {
	ruu.mutation.Where(ps...)
	return ruu
}

// SetAppID sets the "app_id" field.
func (ruu *ReadUserUpdate) SetAppID(u uuid.UUID) *ReadUserUpdate {
	ruu.mutation.SetAppID(u)
	return ruu
}

// SetUserID sets the "user_id" field.
func (ruu *ReadUserUpdate) SetUserID(u uuid.UUID) *ReadUserUpdate {
	ruu.mutation.SetUserID(u)
	return ruu
}

// SetAnnouncementID sets the "announcement_id" field.
func (ruu *ReadUserUpdate) SetAnnouncementID(u uuid.UUID) *ReadUserUpdate {
	ruu.mutation.SetAnnouncementID(u)
	return ruu
}

// SetCreateAt sets the "create_at" field.
func (ruu *ReadUserUpdate) SetCreateAt(u uint32) *ReadUserUpdate {
	ruu.mutation.ResetCreateAt()
	ruu.mutation.SetCreateAt(u)
	return ruu
}

// SetNillableCreateAt sets the "create_at" field if the given value is not nil.
func (ruu *ReadUserUpdate) SetNillableCreateAt(u *uint32) *ReadUserUpdate {
	if u != nil {
		ruu.SetCreateAt(*u)
	}
	return ruu
}

// AddCreateAt adds u to the "create_at" field.
func (ruu *ReadUserUpdate) AddCreateAt(u int32) *ReadUserUpdate {
	ruu.mutation.AddCreateAt(u)
	return ruu
}

// SetUpdateAt sets the "update_at" field.
func (ruu *ReadUserUpdate) SetUpdateAt(u uint32) *ReadUserUpdate {
	ruu.mutation.ResetUpdateAt()
	ruu.mutation.SetUpdateAt(u)
	return ruu
}

// AddUpdateAt adds u to the "update_at" field.
func (ruu *ReadUserUpdate) AddUpdateAt(u int32) *ReadUserUpdate {
	ruu.mutation.AddUpdateAt(u)
	return ruu
}

// SetDeleteAt sets the "delete_at" field.
func (ruu *ReadUserUpdate) SetDeleteAt(u uint32) *ReadUserUpdate {
	ruu.mutation.ResetDeleteAt()
	ruu.mutation.SetDeleteAt(u)
	return ruu
}

// SetNillableDeleteAt sets the "delete_at" field if the given value is not nil.
func (ruu *ReadUserUpdate) SetNillableDeleteAt(u *uint32) *ReadUserUpdate {
	if u != nil {
		ruu.SetDeleteAt(*u)
	}
	return ruu
}

// AddDeleteAt adds u to the "delete_at" field.
func (ruu *ReadUserUpdate) AddDeleteAt(u int32) *ReadUserUpdate {
	ruu.mutation.AddDeleteAt(u)
	return ruu
}

// Mutation returns the ReadUserMutation object of the builder.
func (ruu *ReadUserUpdate) Mutation() *ReadUserMutation {
	return ruu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ruu *ReadUserUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	ruu.defaults()
	if len(ruu.hooks) == 0 {
		affected, err = ruu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ReadUserMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ruu.mutation = mutation
			affected, err = ruu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ruu.hooks) - 1; i >= 0; i-- {
			if ruu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ruu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ruu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (ruu *ReadUserUpdate) SaveX(ctx context.Context) int {
	affected, err := ruu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ruu *ReadUserUpdate) Exec(ctx context.Context) error {
	_, err := ruu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ruu *ReadUserUpdate) ExecX(ctx context.Context) {
	if err := ruu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ruu *ReadUserUpdate) defaults() {
	if _, ok := ruu.mutation.UpdateAt(); !ok {
		v := readuser.UpdateDefaultUpdateAt()
		ruu.mutation.SetUpdateAt(v)
	}
}

func (ruu *ReadUserUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   readuser.Table,
			Columns: readuser.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: readuser.FieldID,
			},
		},
	}
	if ps := ruu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ruu.mutation.AppID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: readuser.FieldAppID,
		})
	}
	if value, ok := ruu.mutation.UserID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: readuser.FieldUserID,
		})
	}
	if value, ok := ruu.mutation.AnnouncementID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: readuser.FieldAnnouncementID,
		})
	}
	if value, ok := ruu.mutation.CreateAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: readuser.FieldCreateAt,
		})
	}
	if value, ok := ruu.mutation.AddedCreateAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: readuser.FieldCreateAt,
		})
	}
	if value, ok := ruu.mutation.UpdateAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: readuser.FieldUpdateAt,
		})
	}
	if value, ok := ruu.mutation.AddedUpdateAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: readuser.FieldUpdateAt,
		})
	}
	if value, ok := ruu.mutation.DeleteAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: readuser.FieldDeleteAt,
		})
	}
	if value, ok := ruu.mutation.AddedDeleteAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: readuser.FieldDeleteAt,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ruu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{readuser.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// ReadUserUpdateOne is the builder for updating a single ReadUser entity.
type ReadUserUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ReadUserMutation
}

// SetAppID sets the "app_id" field.
func (ruuo *ReadUserUpdateOne) SetAppID(u uuid.UUID) *ReadUserUpdateOne {
	ruuo.mutation.SetAppID(u)
	return ruuo
}

// SetUserID sets the "user_id" field.
func (ruuo *ReadUserUpdateOne) SetUserID(u uuid.UUID) *ReadUserUpdateOne {
	ruuo.mutation.SetUserID(u)
	return ruuo
}

// SetAnnouncementID sets the "announcement_id" field.
func (ruuo *ReadUserUpdateOne) SetAnnouncementID(u uuid.UUID) *ReadUserUpdateOne {
	ruuo.mutation.SetAnnouncementID(u)
	return ruuo
}

// SetCreateAt sets the "create_at" field.
func (ruuo *ReadUserUpdateOne) SetCreateAt(u uint32) *ReadUserUpdateOne {
	ruuo.mutation.ResetCreateAt()
	ruuo.mutation.SetCreateAt(u)
	return ruuo
}

// SetNillableCreateAt sets the "create_at" field if the given value is not nil.
func (ruuo *ReadUserUpdateOne) SetNillableCreateAt(u *uint32) *ReadUserUpdateOne {
	if u != nil {
		ruuo.SetCreateAt(*u)
	}
	return ruuo
}

// AddCreateAt adds u to the "create_at" field.
func (ruuo *ReadUserUpdateOne) AddCreateAt(u int32) *ReadUserUpdateOne {
	ruuo.mutation.AddCreateAt(u)
	return ruuo
}

// SetUpdateAt sets the "update_at" field.
func (ruuo *ReadUserUpdateOne) SetUpdateAt(u uint32) *ReadUserUpdateOne {
	ruuo.mutation.ResetUpdateAt()
	ruuo.mutation.SetUpdateAt(u)
	return ruuo
}

// AddUpdateAt adds u to the "update_at" field.
func (ruuo *ReadUserUpdateOne) AddUpdateAt(u int32) *ReadUserUpdateOne {
	ruuo.mutation.AddUpdateAt(u)
	return ruuo
}

// SetDeleteAt sets the "delete_at" field.
func (ruuo *ReadUserUpdateOne) SetDeleteAt(u uint32) *ReadUserUpdateOne {
	ruuo.mutation.ResetDeleteAt()
	ruuo.mutation.SetDeleteAt(u)
	return ruuo
}

// SetNillableDeleteAt sets the "delete_at" field if the given value is not nil.
func (ruuo *ReadUserUpdateOne) SetNillableDeleteAt(u *uint32) *ReadUserUpdateOne {
	if u != nil {
		ruuo.SetDeleteAt(*u)
	}
	return ruuo
}

// AddDeleteAt adds u to the "delete_at" field.
func (ruuo *ReadUserUpdateOne) AddDeleteAt(u int32) *ReadUserUpdateOne {
	ruuo.mutation.AddDeleteAt(u)
	return ruuo
}

// Mutation returns the ReadUserMutation object of the builder.
func (ruuo *ReadUserUpdateOne) Mutation() *ReadUserMutation {
	return ruuo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ruuo *ReadUserUpdateOne) Select(field string, fields ...string) *ReadUserUpdateOne {
	ruuo.fields = append([]string{field}, fields...)
	return ruuo
}

// Save executes the query and returns the updated ReadUser entity.
func (ruuo *ReadUserUpdateOne) Save(ctx context.Context) (*ReadUser, error) {
	var (
		err  error
		node *ReadUser
	)
	ruuo.defaults()
	if len(ruuo.hooks) == 0 {
		node, err = ruuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ReadUserMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ruuo.mutation = mutation
			node, err = ruuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(ruuo.hooks) - 1; i >= 0; i-- {
			if ruuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ruuo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, ruuo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*ReadUser)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from ReadUserMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (ruuo *ReadUserUpdateOne) SaveX(ctx context.Context) *ReadUser {
	node, err := ruuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ruuo *ReadUserUpdateOne) Exec(ctx context.Context) error {
	_, err := ruuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ruuo *ReadUserUpdateOne) ExecX(ctx context.Context) {
	if err := ruuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ruuo *ReadUserUpdateOne) defaults() {
	if _, ok := ruuo.mutation.UpdateAt(); !ok {
		v := readuser.UpdateDefaultUpdateAt()
		ruuo.mutation.SetUpdateAt(v)
	}
}

func (ruuo *ReadUserUpdateOne) sqlSave(ctx context.Context) (_node *ReadUser, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   readuser.Table,
			Columns: readuser.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: readuser.FieldID,
			},
		},
	}
	id, ok := ruuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "ReadUser.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := ruuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, readuser.FieldID)
		for _, f := range fields {
			if !readuser.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != readuser.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ruuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ruuo.mutation.AppID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: readuser.FieldAppID,
		})
	}
	if value, ok := ruuo.mutation.UserID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: readuser.FieldUserID,
		})
	}
	if value, ok := ruuo.mutation.AnnouncementID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: readuser.FieldAnnouncementID,
		})
	}
	if value, ok := ruuo.mutation.CreateAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: readuser.FieldCreateAt,
		})
	}
	if value, ok := ruuo.mutation.AddedCreateAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: readuser.FieldCreateAt,
		})
	}
	if value, ok := ruuo.mutation.UpdateAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: readuser.FieldUpdateAt,
		})
	}
	if value, ok := ruuo.mutation.AddedUpdateAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: readuser.FieldUpdateAt,
		})
	}
	if value, ok := ruuo.mutation.DeleteAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: readuser.FieldDeleteAt,
		})
	}
	if value, ok := ruuo.mutation.AddedDeleteAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: readuser.FieldDeleteAt,
		})
	}
	_node = &ReadUser{config: ruuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ruuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{readuser.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
