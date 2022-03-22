// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/notification/pkg/db/ent/mailbox"
	"github.com/NpoolPlatform/notification/pkg/db/ent/predicate"
	"github.com/google/uuid"
)

// MailBoxUpdate is the builder for updating MailBox entities.
type MailBoxUpdate struct {
	config
	hooks    []Hook
	mutation *MailBoxMutation
}

// Where appends a list predicates to the MailBoxUpdate builder.
func (mbu *MailBoxUpdate) Where(ps ...predicate.MailBox) *MailBoxUpdate {
	mbu.mutation.Where(ps...)
	return mbu
}

// SetAppID sets the "app_id" field.
func (mbu *MailBoxUpdate) SetAppID(u uuid.UUID) *MailBoxUpdate {
	mbu.mutation.SetAppID(u)
	return mbu
}

// SetFromUserID sets the "from_user_id" field.
func (mbu *MailBoxUpdate) SetFromUserID(u uuid.UUID) *MailBoxUpdate {
	mbu.mutation.SetFromUserID(u)
	return mbu
}

// SetToUserID sets the "to_user_id" field.
func (mbu *MailBoxUpdate) SetToUserID(u uuid.UUID) *MailBoxUpdate {
	mbu.mutation.SetToUserID(u)
	return mbu
}

// SetAlreadyRead sets the "already_read" field.
func (mbu *MailBoxUpdate) SetAlreadyRead(b bool) *MailBoxUpdate {
	mbu.mutation.SetAlreadyRead(b)
	return mbu
}

// SetTitle sets the "title" field.
func (mbu *MailBoxUpdate) SetTitle(s string) *MailBoxUpdate {
	mbu.mutation.SetTitle(s)
	return mbu
}

// SetContent sets the "content" field.
func (mbu *MailBoxUpdate) SetContent(s string) *MailBoxUpdate {
	mbu.mutation.SetContent(s)
	return mbu
}

// SetCreateAt sets the "create_at" field.
func (mbu *MailBoxUpdate) SetCreateAt(u uint32) *MailBoxUpdate {
	mbu.mutation.ResetCreateAt()
	mbu.mutation.SetCreateAt(u)
	return mbu
}

// SetNillableCreateAt sets the "create_at" field if the given value is not nil.
func (mbu *MailBoxUpdate) SetNillableCreateAt(u *uint32) *MailBoxUpdate {
	if u != nil {
		mbu.SetCreateAt(*u)
	}
	return mbu
}

// AddCreateAt adds u to the "create_at" field.
func (mbu *MailBoxUpdate) AddCreateAt(u int32) *MailBoxUpdate {
	mbu.mutation.AddCreateAt(u)
	return mbu
}

// SetUpdateAt sets the "update_at" field.
func (mbu *MailBoxUpdate) SetUpdateAt(u uint32) *MailBoxUpdate {
	mbu.mutation.ResetUpdateAt()
	mbu.mutation.SetUpdateAt(u)
	return mbu
}

// AddUpdateAt adds u to the "update_at" field.
func (mbu *MailBoxUpdate) AddUpdateAt(u int32) *MailBoxUpdate {
	mbu.mutation.AddUpdateAt(u)
	return mbu
}

// SetDeleteAt sets the "delete_at" field.
func (mbu *MailBoxUpdate) SetDeleteAt(u uint32) *MailBoxUpdate {
	mbu.mutation.ResetDeleteAt()
	mbu.mutation.SetDeleteAt(u)
	return mbu
}

// SetNillableDeleteAt sets the "delete_at" field if the given value is not nil.
func (mbu *MailBoxUpdate) SetNillableDeleteAt(u *uint32) *MailBoxUpdate {
	if u != nil {
		mbu.SetDeleteAt(*u)
	}
	return mbu
}

// AddDeleteAt adds u to the "delete_at" field.
func (mbu *MailBoxUpdate) AddDeleteAt(u int32) *MailBoxUpdate {
	mbu.mutation.AddDeleteAt(u)
	return mbu
}

// Mutation returns the MailBoxMutation object of the builder.
func (mbu *MailBoxUpdate) Mutation() *MailBoxMutation {
	return mbu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (mbu *MailBoxUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	mbu.defaults()
	if len(mbu.hooks) == 0 {
		affected, err = mbu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*MailBoxMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			mbu.mutation = mutation
			affected, err = mbu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(mbu.hooks) - 1; i >= 0; i-- {
			if mbu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = mbu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, mbu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (mbu *MailBoxUpdate) SaveX(ctx context.Context) int {
	affected, err := mbu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (mbu *MailBoxUpdate) Exec(ctx context.Context) error {
	_, err := mbu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mbu *MailBoxUpdate) ExecX(ctx context.Context) {
	if err := mbu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (mbu *MailBoxUpdate) defaults() {
	if _, ok := mbu.mutation.UpdateAt(); !ok {
		v := mailbox.UpdateDefaultUpdateAt()
		mbu.mutation.SetUpdateAt(v)
	}
}

func (mbu *MailBoxUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   mailbox.Table,
			Columns: mailbox.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: mailbox.FieldID,
			},
		},
	}
	if ps := mbu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := mbu.mutation.AppID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: mailbox.FieldAppID,
		})
	}
	if value, ok := mbu.mutation.FromUserID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: mailbox.FieldFromUserID,
		})
	}
	if value, ok := mbu.mutation.ToUserID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: mailbox.FieldToUserID,
		})
	}
	if value, ok := mbu.mutation.AlreadyRead(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: mailbox.FieldAlreadyRead,
		})
	}
	if value, ok := mbu.mutation.Title(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: mailbox.FieldTitle,
		})
	}
	if value, ok := mbu.mutation.Content(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: mailbox.FieldContent,
		})
	}
	if value, ok := mbu.mutation.CreateAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: mailbox.FieldCreateAt,
		})
	}
	if value, ok := mbu.mutation.AddedCreateAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: mailbox.FieldCreateAt,
		})
	}
	if value, ok := mbu.mutation.UpdateAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: mailbox.FieldUpdateAt,
		})
	}
	if value, ok := mbu.mutation.AddedUpdateAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: mailbox.FieldUpdateAt,
		})
	}
	if value, ok := mbu.mutation.DeleteAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: mailbox.FieldDeleteAt,
		})
	}
	if value, ok := mbu.mutation.AddedDeleteAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: mailbox.FieldDeleteAt,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, mbu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{mailbox.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// MailBoxUpdateOne is the builder for updating a single MailBox entity.
type MailBoxUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *MailBoxMutation
}

// SetAppID sets the "app_id" field.
func (mbuo *MailBoxUpdateOne) SetAppID(u uuid.UUID) *MailBoxUpdateOne {
	mbuo.mutation.SetAppID(u)
	return mbuo
}

// SetFromUserID sets the "from_user_id" field.
func (mbuo *MailBoxUpdateOne) SetFromUserID(u uuid.UUID) *MailBoxUpdateOne {
	mbuo.mutation.SetFromUserID(u)
	return mbuo
}

// SetToUserID sets the "to_user_id" field.
func (mbuo *MailBoxUpdateOne) SetToUserID(u uuid.UUID) *MailBoxUpdateOne {
	mbuo.mutation.SetToUserID(u)
	return mbuo
}

// SetAlreadyRead sets the "already_read" field.
func (mbuo *MailBoxUpdateOne) SetAlreadyRead(b bool) *MailBoxUpdateOne {
	mbuo.mutation.SetAlreadyRead(b)
	return mbuo
}

// SetTitle sets the "title" field.
func (mbuo *MailBoxUpdateOne) SetTitle(s string) *MailBoxUpdateOne {
	mbuo.mutation.SetTitle(s)
	return mbuo
}

// SetContent sets the "content" field.
func (mbuo *MailBoxUpdateOne) SetContent(s string) *MailBoxUpdateOne {
	mbuo.mutation.SetContent(s)
	return mbuo
}

// SetCreateAt sets the "create_at" field.
func (mbuo *MailBoxUpdateOne) SetCreateAt(u uint32) *MailBoxUpdateOne {
	mbuo.mutation.ResetCreateAt()
	mbuo.mutation.SetCreateAt(u)
	return mbuo
}

// SetNillableCreateAt sets the "create_at" field if the given value is not nil.
func (mbuo *MailBoxUpdateOne) SetNillableCreateAt(u *uint32) *MailBoxUpdateOne {
	if u != nil {
		mbuo.SetCreateAt(*u)
	}
	return mbuo
}

// AddCreateAt adds u to the "create_at" field.
func (mbuo *MailBoxUpdateOne) AddCreateAt(u int32) *MailBoxUpdateOne {
	mbuo.mutation.AddCreateAt(u)
	return mbuo
}

// SetUpdateAt sets the "update_at" field.
func (mbuo *MailBoxUpdateOne) SetUpdateAt(u uint32) *MailBoxUpdateOne {
	mbuo.mutation.ResetUpdateAt()
	mbuo.mutation.SetUpdateAt(u)
	return mbuo
}

// AddUpdateAt adds u to the "update_at" field.
func (mbuo *MailBoxUpdateOne) AddUpdateAt(u int32) *MailBoxUpdateOne {
	mbuo.mutation.AddUpdateAt(u)
	return mbuo
}

// SetDeleteAt sets the "delete_at" field.
func (mbuo *MailBoxUpdateOne) SetDeleteAt(u uint32) *MailBoxUpdateOne {
	mbuo.mutation.ResetDeleteAt()
	mbuo.mutation.SetDeleteAt(u)
	return mbuo
}

// SetNillableDeleteAt sets the "delete_at" field if the given value is not nil.
func (mbuo *MailBoxUpdateOne) SetNillableDeleteAt(u *uint32) *MailBoxUpdateOne {
	if u != nil {
		mbuo.SetDeleteAt(*u)
	}
	return mbuo
}

// AddDeleteAt adds u to the "delete_at" field.
func (mbuo *MailBoxUpdateOne) AddDeleteAt(u int32) *MailBoxUpdateOne {
	mbuo.mutation.AddDeleteAt(u)
	return mbuo
}

// Mutation returns the MailBoxMutation object of the builder.
func (mbuo *MailBoxUpdateOne) Mutation() *MailBoxMutation {
	return mbuo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (mbuo *MailBoxUpdateOne) Select(field string, fields ...string) *MailBoxUpdateOne {
	mbuo.fields = append([]string{field}, fields...)
	return mbuo
}

// Save executes the query and returns the updated MailBox entity.
func (mbuo *MailBoxUpdateOne) Save(ctx context.Context) (*MailBox, error) {
	var (
		err  error
		node *MailBox
	)
	mbuo.defaults()
	if len(mbuo.hooks) == 0 {
		node, err = mbuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*MailBoxMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			mbuo.mutation = mutation
			node, err = mbuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(mbuo.hooks) - 1; i >= 0; i-- {
			if mbuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = mbuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, mbuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (mbuo *MailBoxUpdateOne) SaveX(ctx context.Context) *MailBox {
	node, err := mbuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (mbuo *MailBoxUpdateOne) Exec(ctx context.Context) error {
	_, err := mbuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mbuo *MailBoxUpdateOne) ExecX(ctx context.Context) {
	if err := mbuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (mbuo *MailBoxUpdateOne) defaults() {
	if _, ok := mbuo.mutation.UpdateAt(); !ok {
		v := mailbox.UpdateDefaultUpdateAt()
		mbuo.mutation.SetUpdateAt(v)
	}
}

func (mbuo *MailBoxUpdateOne) sqlSave(ctx context.Context) (_node *MailBox, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   mailbox.Table,
			Columns: mailbox.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: mailbox.FieldID,
			},
		},
	}
	id, ok := mbuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "MailBox.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := mbuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, mailbox.FieldID)
		for _, f := range fields {
			if !mailbox.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != mailbox.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := mbuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := mbuo.mutation.AppID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: mailbox.FieldAppID,
		})
	}
	if value, ok := mbuo.mutation.FromUserID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: mailbox.FieldFromUserID,
		})
	}
	if value, ok := mbuo.mutation.ToUserID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: mailbox.FieldToUserID,
		})
	}
	if value, ok := mbuo.mutation.AlreadyRead(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: mailbox.FieldAlreadyRead,
		})
	}
	if value, ok := mbuo.mutation.Title(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: mailbox.FieldTitle,
		})
	}
	if value, ok := mbuo.mutation.Content(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: mailbox.FieldContent,
		})
	}
	if value, ok := mbuo.mutation.CreateAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: mailbox.FieldCreateAt,
		})
	}
	if value, ok := mbuo.mutation.AddedCreateAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: mailbox.FieldCreateAt,
		})
	}
	if value, ok := mbuo.mutation.UpdateAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: mailbox.FieldUpdateAt,
		})
	}
	if value, ok := mbuo.mutation.AddedUpdateAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: mailbox.FieldUpdateAt,
		})
	}
	if value, ok := mbuo.mutation.DeleteAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: mailbox.FieldDeleteAt,
		})
	}
	if value, ok := mbuo.mutation.AddedDeleteAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: mailbox.FieldDeleteAt,
		})
	}
	_node = &MailBox{config: mbuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, mbuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{mailbox.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
