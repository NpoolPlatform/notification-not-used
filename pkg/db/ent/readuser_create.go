// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/notification/pkg/db/ent/readuser"
	"github.com/google/uuid"
)

// ReadUserCreate is the builder for creating a ReadUser entity.
type ReadUserCreate struct {
	config
	mutation *ReadUserMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetAppID sets the "app_id" field.
func (ruc *ReadUserCreate) SetAppID(u uuid.UUID) *ReadUserCreate {
	ruc.mutation.SetAppID(u)
	return ruc
}

// SetUserID sets the "user_id" field.
func (ruc *ReadUserCreate) SetUserID(u uuid.UUID) *ReadUserCreate {
	ruc.mutation.SetUserID(u)
	return ruc
}

// SetAnnouncementID sets the "announcement_id" field.
func (ruc *ReadUserCreate) SetAnnouncementID(u uuid.UUID) *ReadUserCreate {
	ruc.mutation.SetAnnouncementID(u)
	return ruc
}

// SetCreateAt sets the "create_at" field.
func (ruc *ReadUserCreate) SetCreateAt(u uint32) *ReadUserCreate {
	ruc.mutation.SetCreateAt(u)
	return ruc
}

// SetNillableCreateAt sets the "create_at" field if the given value is not nil.
func (ruc *ReadUserCreate) SetNillableCreateAt(u *uint32) *ReadUserCreate {
	if u != nil {
		ruc.SetCreateAt(*u)
	}
	return ruc
}

// SetUpdateAt sets the "update_at" field.
func (ruc *ReadUserCreate) SetUpdateAt(u uint32) *ReadUserCreate {
	ruc.mutation.SetUpdateAt(u)
	return ruc
}

// SetNillableUpdateAt sets the "update_at" field if the given value is not nil.
func (ruc *ReadUserCreate) SetNillableUpdateAt(u *uint32) *ReadUserCreate {
	if u != nil {
		ruc.SetUpdateAt(*u)
	}
	return ruc
}

// SetDeleteAt sets the "delete_at" field.
func (ruc *ReadUserCreate) SetDeleteAt(u uint32) *ReadUserCreate {
	ruc.mutation.SetDeleteAt(u)
	return ruc
}

// SetNillableDeleteAt sets the "delete_at" field if the given value is not nil.
func (ruc *ReadUserCreate) SetNillableDeleteAt(u *uint32) *ReadUserCreate {
	if u != nil {
		ruc.SetDeleteAt(*u)
	}
	return ruc
}

// SetID sets the "id" field.
func (ruc *ReadUserCreate) SetID(u uuid.UUID) *ReadUserCreate {
	ruc.mutation.SetID(u)
	return ruc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (ruc *ReadUserCreate) SetNillableID(u *uuid.UUID) *ReadUserCreate {
	if u != nil {
		ruc.SetID(*u)
	}
	return ruc
}

// Mutation returns the ReadUserMutation object of the builder.
func (ruc *ReadUserCreate) Mutation() *ReadUserMutation {
	return ruc.mutation
}

// Save creates the ReadUser in the database.
func (ruc *ReadUserCreate) Save(ctx context.Context) (*ReadUser, error) {
	var (
		err  error
		node *ReadUser
	)
	ruc.defaults()
	if len(ruc.hooks) == 0 {
		if err = ruc.check(); err != nil {
			return nil, err
		}
		node, err = ruc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ReadUserMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ruc.check(); err != nil {
				return nil, err
			}
			ruc.mutation = mutation
			if node, err = ruc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(ruc.hooks) - 1; i >= 0; i-- {
			if ruc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ruc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ruc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (ruc *ReadUserCreate) SaveX(ctx context.Context) *ReadUser {
	v, err := ruc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ruc *ReadUserCreate) Exec(ctx context.Context) error {
	_, err := ruc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ruc *ReadUserCreate) ExecX(ctx context.Context) {
	if err := ruc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ruc *ReadUserCreate) defaults() {
	if _, ok := ruc.mutation.CreateAt(); !ok {
		v := readuser.DefaultCreateAt()
		ruc.mutation.SetCreateAt(v)
	}
	if _, ok := ruc.mutation.UpdateAt(); !ok {
		v := readuser.DefaultUpdateAt()
		ruc.mutation.SetUpdateAt(v)
	}
	if _, ok := ruc.mutation.DeleteAt(); !ok {
		v := readuser.DefaultDeleteAt()
		ruc.mutation.SetDeleteAt(v)
	}
	if _, ok := ruc.mutation.ID(); !ok {
		v := readuser.DefaultID()
		ruc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ruc *ReadUserCreate) check() error {
	if _, ok := ruc.mutation.AppID(); !ok {
		return &ValidationError{Name: "app_id", err: errors.New(`ent: missing required field "ReadUser.app_id"`)}
	}
	if _, ok := ruc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user_id", err: errors.New(`ent: missing required field "ReadUser.user_id"`)}
	}
	if _, ok := ruc.mutation.AnnouncementID(); !ok {
		return &ValidationError{Name: "announcement_id", err: errors.New(`ent: missing required field "ReadUser.announcement_id"`)}
	}
	if _, ok := ruc.mutation.CreateAt(); !ok {
		return &ValidationError{Name: "create_at", err: errors.New(`ent: missing required field "ReadUser.create_at"`)}
	}
	if _, ok := ruc.mutation.UpdateAt(); !ok {
		return &ValidationError{Name: "update_at", err: errors.New(`ent: missing required field "ReadUser.update_at"`)}
	}
	if _, ok := ruc.mutation.DeleteAt(); !ok {
		return &ValidationError{Name: "delete_at", err: errors.New(`ent: missing required field "ReadUser.delete_at"`)}
	}
	return nil
}

func (ruc *ReadUserCreate) sqlSave(ctx context.Context) (*ReadUser, error) {
	_node, _spec := ruc.createSpec()
	if err := sqlgraph.CreateNode(ctx, ruc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
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

func (ruc *ReadUserCreate) createSpec() (*ReadUser, *sqlgraph.CreateSpec) {
	var (
		_node = &ReadUser{config: ruc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: readuser.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: readuser.FieldID,
			},
		}
	)
	_spec.OnConflict = ruc.conflict
	if id, ok := ruc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := ruc.mutation.AppID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: readuser.FieldAppID,
		})
		_node.AppID = value
	}
	if value, ok := ruc.mutation.UserID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: readuser.FieldUserID,
		})
		_node.UserID = value
	}
	if value, ok := ruc.mutation.AnnouncementID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: readuser.FieldAnnouncementID,
		})
		_node.AnnouncementID = value
	}
	if value, ok := ruc.mutation.CreateAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: readuser.FieldCreateAt,
		})
		_node.CreateAt = value
	}
	if value, ok := ruc.mutation.UpdateAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: readuser.FieldUpdateAt,
		})
		_node.UpdateAt = value
	}
	if value, ok := ruc.mutation.DeleteAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: readuser.FieldDeleteAt,
		})
		_node.DeleteAt = value
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.ReadUser.Create().
//		SetAppID(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.ReadUserUpsert) {
//			SetAppID(v+v).
//		}).
//		Exec(ctx)
//
func (ruc *ReadUserCreate) OnConflict(opts ...sql.ConflictOption) *ReadUserUpsertOne {
	ruc.conflict = opts
	return &ReadUserUpsertOne{
		create: ruc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.ReadUser.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (ruc *ReadUserCreate) OnConflictColumns(columns ...string) *ReadUserUpsertOne {
	ruc.conflict = append(ruc.conflict, sql.ConflictColumns(columns...))
	return &ReadUserUpsertOne{
		create: ruc,
	}
}

type (
	// ReadUserUpsertOne is the builder for "upsert"-ing
	//  one ReadUser node.
	ReadUserUpsertOne struct {
		create *ReadUserCreate
	}

	// ReadUserUpsert is the "OnConflict" setter.
	ReadUserUpsert struct {
		*sql.UpdateSet
	}
)

// SetAppID sets the "app_id" field.
func (u *ReadUserUpsert) SetAppID(v uuid.UUID) *ReadUserUpsert {
	u.Set(readuser.FieldAppID, v)
	return u
}

// UpdateAppID sets the "app_id" field to the value that was provided on create.
func (u *ReadUserUpsert) UpdateAppID() *ReadUserUpsert {
	u.SetExcluded(readuser.FieldAppID)
	return u
}

// SetUserID sets the "user_id" field.
func (u *ReadUserUpsert) SetUserID(v uuid.UUID) *ReadUserUpsert {
	u.Set(readuser.FieldUserID, v)
	return u
}

// UpdateUserID sets the "user_id" field to the value that was provided on create.
func (u *ReadUserUpsert) UpdateUserID() *ReadUserUpsert {
	u.SetExcluded(readuser.FieldUserID)
	return u
}

// SetAnnouncementID sets the "announcement_id" field.
func (u *ReadUserUpsert) SetAnnouncementID(v uuid.UUID) *ReadUserUpsert {
	u.Set(readuser.FieldAnnouncementID, v)
	return u
}

// UpdateAnnouncementID sets the "announcement_id" field to the value that was provided on create.
func (u *ReadUserUpsert) UpdateAnnouncementID() *ReadUserUpsert {
	u.SetExcluded(readuser.FieldAnnouncementID)
	return u
}

// SetCreateAt sets the "create_at" field.
func (u *ReadUserUpsert) SetCreateAt(v uint32) *ReadUserUpsert {
	u.Set(readuser.FieldCreateAt, v)
	return u
}

// UpdateCreateAt sets the "create_at" field to the value that was provided on create.
func (u *ReadUserUpsert) UpdateCreateAt() *ReadUserUpsert {
	u.SetExcluded(readuser.FieldCreateAt)
	return u
}

// AddCreateAt adds v to the "create_at" field.
func (u *ReadUserUpsert) AddCreateAt(v uint32) *ReadUserUpsert {
	u.Add(readuser.FieldCreateAt, v)
	return u
}

// SetUpdateAt sets the "update_at" field.
func (u *ReadUserUpsert) SetUpdateAt(v uint32) *ReadUserUpsert {
	u.Set(readuser.FieldUpdateAt, v)
	return u
}

// UpdateUpdateAt sets the "update_at" field to the value that was provided on create.
func (u *ReadUserUpsert) UpdateUpdateAt() *ReadUserUpsert {
	u.SetExcluded(readuser.FieldUpdateAt)
	return u
}

// AddUpdateAt adds v to the "update_at" field.
func (u *ReadUserUpsert) AddUpdateAt(v uint32) *ReadUserUpsert {
	u.Add(readuser.FieldUpdateAt, v)
	return u
}

// SetDeleteAt sets the "delete_at" field.
func (u *ReadUserUpsert) SetDeleteAt(v uint32) *ReadUserUpsert {
	u.Set(readuser.FieldDeleteAt, v)
	return u
}

// UpdateDeleteAt sets the "delete_at" field to the value that was provided on create.
func (u *ReadUserUpsert) UpdateDeleteAt() *ReadUserUpsert {
	u.SetExcluded(readuser.FieldDeleteAt)
	return u
}

// AddDeleteAt adds v to the "delete_at" field.
func (u *ReadUserUpsert) AddDeleteAt(v uint32) *ReadUserUpsert {
	u.Add(readuser.FieldDeleteAt, v)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.ReadUser.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(readuser.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *ReadUserUpsertOne) UpdateNewValues() *ReadUserUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(readuser.FieldID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//  client.ReadUser.Create().
//      OnConflict(sql.ResolveWithIgnore()).
//      Exec(ctx)
//
func (u *ReadUserUpsertOne) Ignore() *ReadUserUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *ReadUserUpsertOne) DoNothing() *ReadUserUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the ReadUserCreate.OnConflict
// documentation for more info.
func (u *ReadUserUpsertOne) Update(set func(*ReadUserUpsert)) *ReadUserUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&ReadUserUpsert{UpdateSet: update})
	}))
	return u
}

// SetAppID sets the "app_id" field.
func (u *ReadUserUpsertOne) SetAppID(v uuid.UUID) *ReadUserUpsertOne {
	return u.Update(func(s *ReadUserUpsert) {
		s.SetAppID(v)
	})
}

// UpdateAppID sets the "app_id" field to the value that was provided on create.
func (u *ReadUserUpsertOne) UpdateAppID() *ReadUserUpsertOne {
	return u.Update(func(s *ReadUserUpsert) {
		s.UpdateAppID()
	})
}

// SetUserID sets the "user_id" field.
func (u *ReadUserUpsertOne) SetUserID(v uuid.UUID) *ReadUserUpsertOne {
	return u.Update(func(s *ReadUserUpsert) {
		s.SetUserID(v)
	})
}

// UpdateUserID sets the "user_id" field to the value that was provided on create.
func (u *ReadUserUpsertOne) UpdateUserID() *ReadUserUpsertOne {
	return u.Update(func(s *ReadUserUpsert) {
		s.UpdateUserID()
	})
}

// SetAnnouncementID sets the "announcement_id" field.
func (u *ReadUserUpsertOne) SetAnnouncementID(v uuid.UUID) *ReadUserUpsertOne {
	return u.Update(func(s *ReadUserUpsert) {
		s.SetAnnouncementID(v)
	})
}

// UpdateAnnouncementID sets the "announcement_id" field to the value that was provided on create.
func (u *ReadUserUpsertOne) UpdateAnnouncementID() *ReadUserUpsertOne {
	return u.Update(func(s *ReadUserUpsert) {
		s.UpdateAnnouncementID()
	})
}

// SetCreateAt sets the "create_at" field.
func (u *ReadUserUpsertOne) SetCreateAt(v uint32) *ReadUserUpsertOne {
	return u.Update(func(s *ReadUserUpsert) {
		s.SetCreateAt(v)
	})
}

// AddCreateAt adds v to the "create_at" field.
func (u *ReadUserUpsertOne) AddCreateAt(v uint32) *ReadUserUpsertOne {
	return u.Update(func(s *ReadUserUpsert) {
		s.AddCreateAt(v)
	})
}

// UpdateCreateAt sets the "create_at" field to the value that was provided on create.
func (u *ReadUserUpsertOne) UpdateCreateAt() *ReadUserUpsertOne {
	return u.Update(func(s *ReadUserUpsert) {
		s.UpdateCreateAt()
	})
}

// SetUpdateAt sets the "update_at" field.
func (u *ReadUserUpsertOne) SetUpdateAt(v uint32) *ReadUserUpsertOne {
	return u.Update(func(s *ReadUserUpsert) {
		s.SetUpdateAt(v)
	})
}

// AddUpdateAt adds v to the "update_at" field.
func (u *ReadUserUpsertOne) AddUpdateAt(v uint32) *ReadUserUpsertOne {
	return u.Update(func(s *ReadUserUpsert) {
		s.AddUpdateAt(v)
	})
}

// UpdateUpdateAt sets the "update_at" field to the value that was provided on create.
func (u *ReadUserUpsertOne) UpdateUpdateAt() *ReadUserUpsertOne {
	return u.Update(func(s *ReadUserUpsert) {
		s.UpdateUpdateAt()
	})
}

// SetDeleteAt sets the "delete_at" field.
func (u *ReadUserUpsertOne) SetDeleteAt(v uint32) *ReadUserUpsertOne {
	return u.Update(func(s *ReadUserUpsert) {
		s.SetDeleteAt(v)
	})
}

// AddDeleteAt adds v to the "delete_at" field.
func (u *ReadUserUpsertOne) AddDeleteAt(v uint32) *ReadUserUpsertOne {
	return u.Update(func(s *ReadUserUpsert) {
		s.AddDeleteAt(v)
	})
}

// UpdateDeleteAt sets the "delete_at" field to the value that was provided on create.
func (u *ReadUserUpsertOne) UpdateDeleteAt() *ReadUserUpsertOne {
	return u.Update(func(s *ReadUserUpsert) {
		s.UpdateDeleteAt()
	})
}

// Exec executes the query.
func (u *ReadUserUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for ReadUserCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *ReadUserUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *ReadUserUpsertOne) ID(ctx context.Context) (id uuid.UUID, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("ent: ReadUserUpsertOne.ID is not supported by MySQL driver. Use ReadUserUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *ReadUserUpsertOne) IDX(ctx context.Context) uuid.UUID {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// ReadUserCreateBulk is the builder for creating many ReadUser entities in bulk.
type ReadUserCreateBulk struct {
	config
	builders []*ReadUserCreate
	conflict []sql.ConflictOption
}

// Save creates the ReadUser entities in the database.
func (rucb *ReadUserCreateBulk) Save(ctx context.Context) ([]*ReadUser, error) {
	specs := make([]*sqlgraph.CreateSpec, len(rucb.builders))
	nodes := make([]*ReadUser, len(rucb.builders))
	mutators := make([]Mutator, len(rucb.builders))
	for i := range rucb.builders {
		func(i int, root context.Context) {
			builder := rucb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ReadUserMutation)
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
					_, err = mutators[i+1].Mutate(root, rucb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = rucb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, rucb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{err.Error(), err}
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
		if _, err := mutators[0].Mutate(ctx, rucb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (rucb *ReadUserCreateBulk) SaveX(ctx context.Context) []*ReadUser {
	v, err := rucb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (rucb *ReadUserCreateBulk) Exec(ctx context.Context) error {
	_, err := rucb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rucb *ReadUserCreateBulk) ExecX(ctx context.Context) {
	if err := rucb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.ReadUser.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.ReadUserUpsert) {
//			SetAppID(v+v).
//		}).
//		Exec(ctx)
//
func (rucb *ReadUserCreateBulk) OnConflict(opts ...sql.ConflictOption) *ReadUserUpsertBulk {
	rucb.conflict = opts
	return &ReadUserUpsertBulk{
		create: rucb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.ReadUser.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (rucb *ReadUserCreateBulk) OnConflictColumns(columns ...string) *ReadUserUpsertBulk {
	rucb.conflict = append(rucb.conflict, sql.ConflictColumns(columns...))
	return &ReadUserUpsertBulk{
		create: rucb,
	}
}

// ReadUserUpsertBulk is the builder for "upsert"-ing
// a bulk of ReadUser nodes.
type ReadUserUpsertBulk struct {
	create *ReadUserCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.ReadUser.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(readuser.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *ReadUserUpsertBulk) UpdateNewValues() *ReadUserUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(readuser.FieldID)
				return
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.ReadUser.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
//
func (u *ReadUserUpsertBulk) Ignore() *ReadUserUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *ReadUserUpsertBulk) DoNothing() *ReadUserUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the ReadUserCreateBulk.OnConflict
// documentation for more info.
func (u *ReadUserUpsertBulk) Update(set func(*ReadUserUpsert)) *ReadUserUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&ReadUserUpsert{UpdateSet: update})
	}))
	return u
}

// SetAppID sets the "app_id" field.
func (u *ReadUserUpsertBulk) SetAppID(v uuid.UUID) *ReadUserUpsertBulk {
	return u.Update(func(s *ReadUserUpsert) {
		s.SetAppID(v)
	})
}

// UpdateAppID sets the "app_id" field to the value that was provided on create.
func (u *ReadUserUpsertBulk) UpdateAppID() *ReadUserUpsertBulk {
	return u.Update(func(s *ReadUserUpsert) {
		s.UpdateAppID()
	})
}

// SetUserID sets the "user_id" field.
func (u *ReadUserUpsertBulk) SetUserID(v uuid.UUID) *ReadUserUpsertBulk {
	return u.Update(func(s *ReadUserUpsert) {
		s.SetUserID(v)
	})
}

// UpdateUserID sets the "user_id" field to the value that was provided on create.
func (u *ReadUserUpsertBulk) UpdateUserID() *ReadUserUpsertBulk {
	return u.Update(func(s *ReadUserUpsert) {
		s.UpdateUserID()
	})
}

// SetAnnouncementID sets the "announcement_id" field.
func (u *ReadUserUpsertBulk) SetAnnouncementID(v uuid.UUID) *ReadUserUpsertBulk {
	return u.Update(func(s *ReadUserUpsert) {
		s.SetAnnouncementID(v)
	})
}

// UpdateAnnouncementID sets the "announcement_id" field to the value that was provided on create.
func (u *ReadUserUpsertBulk) UpdateAnnouncementID() *ReadUserUpsertBulk {
	return u.Update(func(s *ReadUserUpsert) {
		s.UpdateAnnouncementID()
	})
}

// SetCreateAt sets the "create_at" field.
func (u *ReadUserUpsertBulk) SetCreateAt(v uint32) *ReadUserUpsertBulk {
	return u.Update(func(s *ReadUserUpsert) {
		s.SetCreateAt(v)
	})
}

// AddCreateAt adds v to the "create_at" field.
func (u *ReadUserUpsertBulk) AddCreateAt(v uint32) *ReadUserUpsertBulk {
	return u.Update(func(s *ReadUserUpsert) {
		s.AddCreateAt(v)
	})
}

// UpdateCreateAt sets the "create_at" field to the value that was provided on create.
func (u *ReadUserUpsertBulk) UpdateCreateAt() *ReadUserUpsertBulk {
	return u.Update(func(s *ReadUserUpsert) {
		s.UpdateCreateAt()
	})
}

// SetUpdateAt sets the "update_at" field.
func (u *ReadUserUpsertBulk) SetUpdateAt(v uint32) *ReadUserUpsertBulk {
	return u.Update(func(s *ReadUserUpsert) {
		s.SetUpdateAt(v)
	})
}

// AddUpdateAt adds v to the "update_at" field.
func (u *ReadUserUpsertBulk) AddUpdateAt(v uint32) *ReadUserUpsertBulk {
	return u.Update(func(s *ReadUserUpsert) {
		s.AddUpdateAt(v)
	})
}

// UpdateUpdateAt sets the "update_at" field to the value that was provided on create.
func (u *ReadUserUpsertBulk) UpdateUpdateAt() *ReadUserUpsertBulk {
	return u.Update(func(s *ReadUserUpsert) {
		s.UpdateUpdateAt()
	})
}

// SetDeleteAt sets the "delete_at" field.
func (u *ReadUserUpsertBulk) SetDeleteAt(v uint32) *ReadUserUpsertBulk {
	return u.Update(func(s *ReadUserUpsert) {
		s.SetDeleteAt(v)
	})
}

// AddDeleteAt adds v to the "delete_at" field.
func (u *ReadUserUpsertBulk) AddDeleteAt(v uint32) *ReadUserUpsertBulk {
	return u.Update(func(s *ReadUserUpsert) {
		s.AddDeleteAt(v)
	})
}

// UpdateDeleteAt sets the "delete_at" field to the value that was provided on create.
func (u *ReadUserUpsertBulk) UpdateDeleteAt() *ReadUserUpsertBulk {
	return u.Update(func(s *ReadUserUpsert) {
		s.UpdateDeleteAt()
	})
}

// Exec executes the query.
func (u *ReadUserUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the ReadUserCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for ReadUserCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *ReadUserUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
