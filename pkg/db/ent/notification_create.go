// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/notification/pkg/db/ent/notification"
	"github.com/google/uuid"
)

// NotificationCreate is the builder for creating a Notification entity.
type NotificationCreate struct {
	config
	mutation *NotificationMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetAppID sets the "app_id" field.
func (nc *NotificationCreate) SetAppID(u uuid.UUID) *NotificationCreate {
	nc.mutation.SetAppID(u)
	return nc
}

// SetUserID sets the "user_id" field.
func (nc *NotificationCreate) SetUserID(u uuid.UUID) *NotificationCreate {
	nc.mutation.SetUserID(u)
	return nc
}

// SetAlreadyRead sets the "already_read" field.
func (nc *NotificationCreate) SetAlreadyRead(b bool) *NotificationCreate {
	nc.mutation.SetAlreadyRead(b)
	return nc
}

// SetTitle sets the "title" field.
func (nc *NotificationCreate) SetTitle(s string) *NotificationCreate {
	nc.mutation.SetTitle(s)
	return nc
}

// SetContent sets the "content" field.
func (nc *NotificationCreate) SetContent(s string) *NotificationCreate {
	nc.mutation.SetContent(s)
	return nc
}

// SetCreateAt sets the "create_at" field.
func (nc *NotificationCreate) SetCreateAt(u uint32) *NotificationCreate {
	nc.mutation.SetCreateAt(u)
	return nc
}

// SetNillableCreateAt sets the "create_at" field if the given value is not nil.
func (nc *NotificationCreate) SetNillableCreateAt(u *uint32) *NotificationCreate {
	if u != nil {
		nc.SetCreateAt(*u)
	}
	return nc
}

// SetUpdateAt sets the "update_at" field.
func (nc *NotificationCreate) SetUpdateAt(u uint32) *NotificationCreate {
	nc.mutation.SetUpdateAt(u)
	return nc
}

// SetNillableUpdateAt sets the "update_at" field if the given value is not nil.
func (nc *NotificationCreate) SetNillableUpdateAt(u *uint32) *NotificationCreate {
	if u != nil {
		nc.SetUpdateAt(*u)
	}
	return nc
}

// SetDeleteAt sets the "delete_at" field.
func (nc *NotificationCreate) SetDeleteAt(u uint32) *NotificationCreate {
	nc.mutation.SetDeleteAt(u)
	return nc
}

// SetNillableDeleteAt sets the "delete_at" field if the given value is not nil.
func (nc *NotificationCreate) SetNillableDeleteAt(u *uint32) *NotificationCreate {
	if u != nil {
		nc.SetDeleteAt(*u)
	}
	return nc
}

// SetID sets the "id" field.
func (nc *NotificationCreate) SetID(u uuid.UUID) *NotificationCreate {
	nc.mutation.SetID(u)
	return nc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (nc *NotificationCreate) SetNillableID(u *uuid.UUID) *NotificationCreate {
	if u != nil {
		nc.SetID(*u)
	}
	return nc
}

// Mutation returns the NotificationMutation object of the builder.
func (nc *NotificationCreate) Mutation() *NotificationMutation {
	return nc.mutation
}

// Save creates the Notification in the database.
func (nc *NotificationCreate) Save(ctx context.Context) (*Notification, error) {
	var (
		err  error
		node *Notification
	)
	nc.defaults()
	if len(nc.hooks) == 0 {
		if err = nc.check(); err != nil {
			return nil, err
		}
		node, err = nc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*NotificationMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = nc.check(); err != nil {
				return nil, err
			}
			nc.mutation = mutation
			if node, err = nc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(nc.hooks) - 1; i >= 0; i-- {
			if nc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = nc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, nc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Notification)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from NotificationMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (nc *NotificationCreate) SaveX(ctx context.Context) *Notification {
	v, err := nc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (nc *NotificationCreate) Exec(ctx context.Context) error {
	_, err := nc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (nc *NotificationCreate) ExecX(ctx context.Context) {
	if err := nc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (nc *NotificationCreate) defaults() {
	if _, ok := nc.mutation.CreateAt(); !ok {
		v := notification.DefaultCreateAt()
		nc.mutation.SetCreateAt(v)
	}
	if _, ok := nc.mutation.UpdateAt(); !ok {
		v := notification.DefaultUpdateAt()
		nc.mutation.SetUpdateAt(v)
	}
	if _, ok := nc.mutation.DeleteAt(); !ok {
		v := notification.DefaultDeleteAt()
		nc.mutation.SetDeleteAt(v)
	}
	if _, ok := nc.mutation.ID(); !ok {
		v := notification.DefaultID()
		nc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (nc *NotificationCreate) check() error {
	if _, ok := nc.mutation.AppID(); !ok {
		return &ValidationError{Name: "app_id", err: errors.New(`ent: missing required field "Notification.app_id"`)}
	}
	if _, ok := nc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user_id", err: errors.New(`ent: missing required field "Notification.user_id"`)}
	}
	if _, ok := nc.mutation.AlreadyRead(); !ok {
		return &ValidationError{Name: "already_read", err: errors.New(`ent: missing required field "Notification.already_read"`)}
	}
	if _, ok := nc.mutation.Title(); !ok {
		return &ValidationError{Name: "title", err: errors.New(`ent: missing required field "Notification.title"`)}
	}
	if _, ok := nc.mutation.Content(); !ok {
		return &ValidationError{Name: "content", err: errors.New(`ent: missing required field "Notification.content"`)}
	}
	if v, ok := nc.mutation.Content(); ok {
		if err := notification.ContentValidator(v); err != nil {
			return &ValidationError{Name: "content", err: fmt.Errorf(`ent: validator failed for field "Notification.content": %w`, err)}
		}
	}
	if _, ok := nc.mutation.CreateAt(); !ok {
		return &ValidationError{Name: "create_at", err: errors.New(`ent: missing required field "Notification.create_at"`)}
	}
	if _, ok := nc.mutation.UpdateAt(); !ok {
		return &ValidationError{Name: "update_at", err: errors.New(`ent: missing required field "Notification.update_at"`)}
	}
	if _, ok := nc.mutation.DeleteAt(); !ok {
		return &ValidationError{Name: "delete_at", err: errors.New(`ent: missing required field "Notification.delete_at"`)}
	}
	return nil
}

func (nc *NotificationCreate) sqlSave(ctx context.Context) (*Notification, error) {
	_node, _spec := nc.createSpec()
	if err := sqlgraph.CreateNode(ctx, nc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*uuid.UUID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	return _node, nil
}

func (nc *NotificationCreate) createSpec() (*Notification, *sqlgraph.CreateSpec) {
	var (
		_node = &Notification{config: nc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: notification.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: notification.FieldID,
			},
		}
	)
	_spec.OnConflict = nc.conflict
	if id, ok := nc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := nc.mutation.AppID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: notification.FieldAppID,
		})
		_node.AppID = value
	}
	if value, ok := nc.mutation.UserID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: notification.FieldUserID,
		})
		_node.UserID = value
	}
	if value, ok := nc.mutation.AlreadyRead(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: notification.FieldAlreadyRead,
		})
		_node.AlreadyRead = value
	}
	if value, ok := nc.mutation.Title(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: notification.FieldTitle,
		})
		_node.Title = value
	}
	if value, ok := nc.mutation.Content(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: notification.FieldContent,
		})
		_node.Content = value
	}
	if value, ok := nc.mutation.CreateAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: notification.FieldCreateAt,
		})
		_node.CreateAt = value
	}
	if value, ok := nc.mutation.UpdateAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: notification.FieldUpdateAt,
		})
		_node.UpdateAt = value
	}
	if value, ok := nc.mutation.DeleteAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: notification.FieldDeleteAt,
		})
		_node.DeleteAt = value
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Notification.Create().
//		SetAppID(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.NotificationUpsert) {
//			SetAppID(v+v).
//		}).
//		Exec(ctx)
//
func (nc *NotificationCreate) OnConflict(opts ...sql.ConflictOption) *NotificationUpsertOne {
	nc.conflict = opts
	return &NotificationUpsertOne{
		create: nc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Notification.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (nc *NotificationCreate) OnConflictColumns(columns ...string) *NotificationUpsertOne {
	nc.conflict = append(nc.conflict, sql.ConflictColumns(columns...))
	return &NotificationUpsertOne{
		create: nc,
	}
}

type (
	// NotificationUpsertOne is the builder for "upsert"-ing
	//  one Notification node.
	NotificationUpsertOne struct {
		create *NotificationCreate
	}

	// NotificationUpsert is the "OnConflict" setter.
	NotificationUpsert struct {
		*sql.UpdateSet
	}
)

// SetAppID sets the "app_id" field.
func (u *NotificationUpsert) SetAppID(v uuid.UUID) *NotificationUpsert {
	u.Set(notification.FieldAppID, v)
	return u
}

// UpdateAppID sets the "app_id" field to the value that was provided on create.
func (u *NotificationUpsert) UpdateAppID() *NotificationUpsert {
	u.SetExcluded(notification.FieldAppID)
	return u
}

// SetUserID sets the "user_id" field.
func (u *NotificationUpsert) SetUserID(v uuid.UUID) *NotificationUpsert {
	u.Set(notification.FieldUserID, v)
	return u
}

// UpdateUserID sets the "user_id" field to the value that was provided on create.
func (u *NotificationUpsert) UpdateUserID() *NotificationUpsert {
	u.SetExcluded(notification.FieldUserID)
	return u
}

// SetAlreadyRead sets the "already_read" field.
func (u *NotificationUpsert) SetAlreadyRead(v bool) *NotificationUpsert {
	u.Set(notification.FieldAlreadyRead, v)
	return u
}

// UpdateAlreadyRead sets the "already_read" field to the value that was provided on create.
func (u *NotificationUpsert) UpdateAlreadyRead() *NotificationUpsert {
	u.SetExcluded(notification.FieldAlreadyRead)
	return u
}

// SetTitle sets the "title" field.
func (u *NotificationUpsert) SetTitle(v string) *NotificationUpsert {
	u.Set(notification.FieldTitle, v)
	return u
}

// UpdateTitle sets the "title" field to the value that was provided on create.
func (u *NotificationUpsert) UpdateTitle() *NotificationUpsert {
	u.SetExcluded(notification.FieldTitle)
	return u
}

// SetContent sets the "content" field.
func (u *NotificationUpsert) SetContent(v string) *NotificationUpsert {
	u.Set(notification.FieldContent, v)
	return u
}

// UpdateContent sets the "content" field to the value that was provided on create.
func (u *NotificationUpsert) UpdateContent() *NotificationUpsert {
	u.SetExcluded(notification.FieldContent)
	return u
}

// SetCreateAt sets the "create_at" field.
func (u *NotificationUpsert) SetCreateAt(v uint32) *NotificationUpsert {
	u.Set(notification.FieldCreateAt, v)
	return u
}

// UpdateCreateAt sets the "create_at" field to the value that was provided on create.
func (u *NotificationUpsert) UpdateCreateAt() *NotificationUpsert {
	u.SetExcluded(notification.FieldCreateAt)
	return u
}

// AddCreateAt adds v to the "create_at" field.
func (u *NotificationUpsert) AddCreateAt(v uint32) *NotificationUpsert {
	u.Add(notification.FieldCreateAt, v)
	return u
}

// SetUpdateAt sets the "update_at" field.
func (u *NotificationUpsert) SetUpdateAt(v uint32) *NotificationUpsert {
	u.Set(notification.FieldUpdateAt, v)
	return u
}

// UpdateUpdateAt sets the "update_at" field to the value that was provided on create.
func (u *NotificationUpsert) UpdateUpdateAt() *NotificationUpsert {
	u.SetExcluded(notification.FieldUpdateAt)
	return u
}

// AddUpdateAt adds v to the "update_at" field.
func (u *NotificationUpsert) AddUpdateAt(v uint32) *NotificationUpsert {
	u.Add(notification.FieldUpdateAt, v)
	return u
}

// SetDeleteAt sets the "delete_at" field.
func (u *NotificationUpsert) SetDeleteAt(v uint32) *NotificationUpsert {
	u.Set(notification.FieldDeleteAt, v)
	return u
}

// UpdateDeleteAt sets the "delete_at" field to the value that was provided on create.
func (u *NotificationUpsert) UpdateDeleteAt() *NotificationUpsert {
	u.SetExcluded(notification.FieldDeleteAt)
	return u
}

// AddDeleteAt adds v to the "delete_at" field.
func (u *NotificationUpsert) AddDeleteAt(v uint32) *NotificationUpsert {
	u.Add(notification.FieldDeleteAt, v)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.Notification.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(notification.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *NotificationUpsertOne) UpdateNewValues() *NotificationUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(notification.FieldID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//  client.Notification.Create().
//      OnConflict(sql.ResolveWithIgnore()).
//      Exec(ctx)
//
func (u *NotificationUpsertOne) Ignore() *NotificationUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *NotificationUpsertOne) DoNothing() *NotificationUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the NotificationCreate.OnConflict
// documentation for more info.
func (u *NotificationUpsertOne) Update(set func(*NotificationUpsert)) *NotificationUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&NotificationUpsert{UpdateSet: update})
	}))
	return u
}

// SetAppID sets the "app_id" field.
func (u *NotificationUpsertOne) SetAppID(v uuid.UUID) *NotificationUpsertOne {
	return u.Update(func(s *NotificationUpsert) {
		s.SetAppID(v)
	})
}

// UpdateAppID sets the "app_id" field to the value that was provided on create.
func (u *NotificationUpsertOne) UpdateAppID() *NotificationUpsertOne {
	return u.Update(func(s *NotificationUpsert) {
		s.UpdateAppID()
	})
}

// SetUserID sets the "user_id" field.
func (u *NotificationUpsertOne) SetUserID(v uuid.UUID) *NotificationUpsertOne {
	return u.Update(func(s *NotificationUpsert) {
		s.SetUserID(v)
	})
}

// UpdateUserID sets the "user_id" field to the value that was provided on create.
func (u *NotificationUpsertOne) UpdateUserID() *NotificationUpsertOne {
	return u.Update(func(s *NotificationUpsert) {
		s.UpdateUserID()
	})
}

// SetAlreadyRead sets the "already_read" field.
func (u *NotificationUpsertOne) SetAlreadyRead(v bool) *NotificationUpsertOne {
	return u.Update(func(s *NotificationUpsert) {
		s.SetAlreadyRead(v)
	})
}

// UpdateAlreadyRead sets the "already_read" field to the value that was provided on create.
func (u *NotificationUpsertOne) UpdateAlreadyRead() *NotificationUpsertOne {
	return u.Update(func(s *NotificationUpsert) {
		s.UpdateAlreadyRead()
	})
}

// SetTitle sets the "title" field.
func (u *NotificationUpsertOne) SetTitle(v string) *NotificationUpsertOne {
	return u.Update(func(s *NotificationUpsert) {
		s.SetTitle(v)
	})
}

// UpdateTitle sets the "title" field to the value that was provided on create.
func (u *NotificationUpsertOne) UpdateTitle() *NotificationUpsertOne {
	return u.Update(func(s *NotificationUpsert) {
		s.UpdateTitle()
	})
}

// SetContent sets the "content" field.
func (u *NotificationUpsertOne) SetContent(v string) *NotificationUpsertOne {
	return u.Update(func(s *NotificationUpsert) {
		s.SetContent(v)
	})
}

// UpdateContent sets the "content" field to the value that was provided on create.
func (u *NotificationUpsertOne) UpdateContent() *NotificationUpsertOne {
	return u.Update(func(s *NotificationUpsert) {
		s.UpdateContent()
	})
}

// SetCreateAt sets the "create_at" field.
func (u *NotificationUpsertOne) SetCreateAt(v uint32) *NotificationUpsertOne {
	return u.Update(func(s *NotificationUpsert) {
		s.SetCreateAt(v)
	})
}

// AddCreateAt adds v to the "create_at" field.
func (u *NotificationUpsertOne) AddCreateAt(v uint32) *NotificationUpsertOne {
	return u.Update(func(s *NotificationUpsert) {
		s.AddCreateAt(v)
	})
}

// UpdateCreateAt sets the "create_at" field to the value that was provided on create.
func (u *NotificationUpsertOne) UpdateCreateAt() *NotificationUpsertOne {
	return u.Update(func(s *NotificationUpsert) {
		s.UpdateCreateAt()
	})
}

// SetUpdateAt sets the "update_at" field.
func (u *NotificationUpsertOne) SetUpdateAt(v uint32) *NotificationUpsertOne {
	return u.Update(func(s *NotificationUpsert) {
		s.SetUpdateAt(v)
	})
}

// AddUpdateAt adds v to the "update_at" field.
func (u *NotificationUpsertOne) AddUpdateAt(v uint32) *NotificationUpsertOne {
	return u.Update(func(s *NotificationUpsert) {
		s.AddUpdateAt(v)
	})
}

// UpdateUpdateAt sets the "update_at" field to the value that was provided on create.
func (u *NotificationUpsertOne) UpdateUpdateAt() *NotificationUpsertOne {
	return u.Update(func(s *NotificationUpsert) {
		s.UpdateUpdateAt()
	})
}

// SetDeleteAt sets the "delete_at" field.
func (u *NotificationUpsertOne) SetDeleteAt(v uint32) *NotificationUpsertOne {
	return u.Update(func(s *NotificationUpsert) {
		s.SetDeleteAt(v)
	})
}

// AddDeleteAt adds v to the "delete_at" field.
func (u *NotificationUpsertOne) AddDeleteAt(v uint32) *NotificationUpsertOne {
	return u.Update(func(s *NotificationUpsert) {
		s.AddDeleteAt(v)
	})
}

// UpdateDeleteAt sets the "delete_at" field to the value that was provided on create.
func (u *NotificationUpsertOne) UpdateDeleteAt() *NotificationUpsertOne {
	return u.Update(func(s *NotificationUpsert) {
		s.UpdateDeleteAt()
	})
}

// Exec executes the query.
func (u *NotificationUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for NotificationCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *NotificationUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *NotificationUpsertOne) ID(ctx context.Context) (id uuid.UUID, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("ent: NotificationUpsertOne.ID is not supported by MySQL driver. Use NotificationUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *NotificationUpsertOne) IDX(ctx context.Context) uuid.UUID {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// NotificationCreateBulk is the builder for creating many Notification entities in bulk.
type NotificationCreateBulk struct {
	config
	builders []*NotificationCreate
	conflict []sql.ConflictOption
}

// Save creates the Notification entities in the database.
func (ncb *NotificationCreateBulk) Save(ctx context.Context) ([]*Notification, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ncb.builders))
	nodes := make([]*Notification, len(ncb.builders))
	mutators := make([]Mutator, len(ncb.builders))
	for i := range ncb.builders {
		func(i int, root context.Context) {
			builder := ncb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*NotificationMutation)
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
					_, err = mutators[i+1].Mutate(root, ncb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = ncb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ncb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
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
		if _, err := mutators[0].Mutate(ctx, ncb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ncb *NotificationCreateBulk) SaveX(ctx context.Context) []*Notification {
	v, err := ncb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ncb *NotificationCreateBulk) Exec(ctx context.Context) error {
	_, err := ncb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ncb *NotificationCreateBulk) ExecX(ctx context.Context) {
	if err := ncb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Notification.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.NotificationUpsert) {
//			SetAppID(v+v).
//		}).
//		Exec(ctx)
//
func (ncb *NotificationCreateBulk) OnConflict(opts ...sql.ConflictOption) *NotificationUpsertBulk {
	ncb.conflict = opts
	return &NotificationUpsertBulk{
		create: ncb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Notification.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (ncb *NotificationCreateBulk) OnConflictColumns(columns ...string) *NotificationUpsertBulk {
	ncb.conflict = append(ncb.conflict, sql.ConflictColumns(columns...))
	return &NotificationUpsertBulk{
		create: ncb,
	}
}

// NotificationUpsertBulk is the builder for "upsert"-ing
// a bulk of Notification nodes.
type NotificationUpsertBulk struct {
	create *NotificationCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Notification.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(notification.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *NotificationUpsertBulk) UpdateNewValues() *NotificationUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(notification.FieldID)
				return
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Notification.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
//
func (u *NotificationUpsertBulk) Ignore() *NotificationUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *NotificationUpsertBulk) DoNothing() *NotificationUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the NotificationCreateBulk.OnConflict
// documentation for more info.
func (u *NotificationUpsertBulk) Update(set func(*NotificationUpsert)) *NotificationUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&NotificationUpsert{UpdateSet: update})
	}))
	return u
}

// SetAppID sets the "app_id" field.
func (u *NotificationUpsertBulk) SetAppID(v uuid.UUID) *NotificationUpsertBulk {
	return u.Update(func(s *NotificationUpsert) {
		s.SetAppID(v)
	})
}

// UpdateAppID sets the "app_id" field to the value that was provided on create.
func (u *NotificationUpsertBulk) UpdateAppID() *NotificationUpsertBulk {
	return u.Update(func(s *NotificationUpsert) {
		s.UpdateAppID()
	})
}

// SetUserID sets the "user_id" field.
func (u *NotificationUpsertBulk) SetUserID(v uuid.UUID) *NotificationUpsertBulk {
	return u.Update(func(s *NotificationUpsert) {
		s.SetUserID(v)
	})
}

// UpdateUserID sets the "user_id" field to the value that was provided on create.
func (u *NotificationUpsertBulk) UpdateUserID() *NotificationUpsertBulk {
	return u.Update(func(s *NotificationUpsert) {
		s.UpdateUserID()
	})
}

// SetAlreadyRead sets the "already_read" field.
func (u *NotificationUpsertBulk) SetAlreadyRead(v bool) *NotificationUpsertBulk {
	return u.Update(func(s *NotificationUpsert) {
		s.SetAlreadyRead(v)
	})
}

// UpdateAlreadyRead sets the "already_read" field to the value that was provided on create.
func (u *NotificationUpsertBulk) UpdateAlreadyRead() *NotificationUpsertBulk {
	return u.Update(func(s *NotificationUpsert) {
		s.UpdateAlreadyRead()
	})
}

// SetTitle sets the "title" field.
func (u *NotificationUpsertBulk) SetTitle(v string) *NotificationUpsertBulk {
	return u.Update(func(s *NotificationUpsert) {
		s.SetTitle(v)
	})
}

// UpdateTitle sets the "title" field to the value that was provided on create.
func (u *NotificationUpsertBulk) UpdateTitle() *NotificationUpsertBulk {
	return u.Update(func(s *NotificationUpsert) {
		s.UpdateTitle()
	})
}

// SetContent sets the "content" field.
func (u *NotificationUpsertBulk) SetContent(v string) *NotificationUpsertBulk {
	return u.Update(func(s *NotificationUpsert) {
		s.SetContent(v)
	})
}

// UpdateContent sets the "content" field to the value that was provided on create.
func (u *NotificationUpsertBulk) UpdateContent() *NotificationUpsertBulk {
	return u.Update(func(s *NotificationUpsert) {
		s.UpdateContent()
	})
}

// SetCreateAt sets the "create_at" field.
func (u *NotificationUpsertBulk) SetCreateAt(v uint32) *NotificationUpsertBulk {
	return u.Update(func(s *NotificationUpsert) {
		s.SetCreateAt(v)
	})
}

// AddCreateAt adds v to the "create_at" field.
func (u *NotificationUpsertBulk) AddCreateAt(v uint32) *NotificationUpsertBulk {
	return u.Update(func(s *NotificationUpsert) {
		s.AddCreateAt(v)
	})
}

// UpdateCreateAt sets the "create_at" field to the value that was provided on create.
func (u *NotificationUpsertBulk) UpdateCreateAt() *NotificationUpsertBulk {
	return u.Update(func(s *NotificationUpsert) {
		s.UpdateCreateAt()
	})
}

// SetUpdateAt sets the "update_at" field.
func (u *NotificationUpsertBulk) SetUpdateAt(v uint32) *NotificationUpsertBulk {
	return u.Update(func(s *NotificationUpsert) {
		s.SetUpdateAt(v)
	})
}

// AddUpdateAt adds v to the "update_at" field.
func (u *NotificationUpsertBulk) AddUpdateAt(v uint32) *NotificationUpsertBulk {
	return u.Update(func(s *NotificationUpsert) {
		s.AddUpdateAt(v)
	})
}

// UpdateUpdateAt sets the "update_at" field to the value that was provided on create.
func (u *NotificationUpsertBulk) UpdateUpdateAt() *NotificationUpsertBulk {
	return u.Update(func(s *NotificationUpsert) {
		s.UpdateUpdateAt()
	})
}

// SetDeleteAt sets the "delete_at" field.
func (u *NotificationUpsertBulk) SetDeleteAt(v uint32) *NotificationUpsertBulk {
	return u.Update(func(s *NotificationUpsert) {
		s.SetDeleteAt(v)
	})
}

// AddDeleteAt adds v to the "delete_at" field.
func (u *NotificationUpsertBulk) AddDeleteAt(v uint32) *NotificationUpsertBulk {
	return u.Update(func(s *NotificationUpsert) {
		s.AddDeleteAt(v)
	})
}

// UpdateDeleteAt sets the "delete_at" field to the value that was provided on create.
func (u *NotificationUpsertBulk) UpdateDeleteAt() *NotificationUpsertBulk {
	return u.Update(func(s *NotificationUpsert) {
		s.UpdateDeleteAt()
	})
}

// Exec executes the query.
func (u *NotificationUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the NotificationCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for NotificationCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *NotificationUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
