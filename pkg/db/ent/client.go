// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"log"

	"github.com/NpoolPlatform/notification/pkg/db/ent/migrate"

	"github.com/NpoolPlatform/notification/pkg/db/ent/announcement"
	"github.com/NpoolPlatform/notification/pkg/db/ent/mailbox"
	"github.com/NpoolPlatform/notification/pkg/db/ent/notification"
	"github.com/NpoolPlatform/notification/pkg/db/ent/readstate"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Announcement is the client for interacting with the Announcement builders.
	Announcement *AnnouncementClient
	// MailBox is the client for interacting with the MailBox builders.
	MailBox *MailBoxClient
	// Notification is the client for interacting with the Notification builders.
	Notification *NotificationClient
	// ReadState is the client for interacting with the ReadState builders.
	ReadState *ReadStateClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	cfg := config{log: log.Println, hooks: &hooks{}}
	cfg.options(opts...)
	client := &Client{config: cfg}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.Announcement = NewAnnouncementClient(c.config)
	c.MailBox = NewMailBoxClient(c.config)
	c.Notification = NewNotificationClient(c.config)
	c.ReadState = NewReadStateClient(c.config)
}

// Open opens a database/sql.DB specified by the driver name and
// the data source name, and returns a new client attached to it.
// Optional parameters can be added for configuring the client.
func Open(driverName, dataSourceName string, options ...Option) (*Client, error) {
	switch driverName {
	case dialect.MySQL, dialect.Postgres, dialect.SQLite:
		drv, err := sql.Open(driverName, dataSourceName)
		if err != nil {
			return nil, err
		}
		return NewClient(append(options, Driver(drv))...), nil
	default:
		return nil, fmt.Errorf("unsupported driver: %q", driverName)
	}
}

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:          ctx,
		config:       cfg,
		Announcement: NewAnnouncementClient(cfg),
		MailBox:      NewMailBoxClient(cfg),
		Notification: NewNotificationClient(cfg),
		ReadState:    NewReadStateClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with specified options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
	}
	tx, err := c.driver.(interface {
		BeginTx(context.Context, *sql.TxOptions) (dialect.Tx, error)
	}).BeginTx(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = &txDriver{tx: tx, drv: c.driver}
	return &Tx{
		config:       cfg,
		Announcement: NewAnnouncementClient(cfg),
		MailBox:      NewMailBoxClient(cfg),
		Notification: NewNotificationClient(cfg),
		ReadState:    NewReadStateClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Announcement.
//		Query().
//		Count(ctx)
//
func (c *Client) Debug() *Client {
	if c.debug {
		return c
	}
	cfg := c.config
	cfg.driver = dialect.Debug(c.driver, c.log)
	client := &Client{config: cfg}
	client.init()
	return client
}

// Close closes the database connection and prevents new queries from starting.
func (c *Client) Close() error {
	return c.driver.Close()
}

// Use adds the mutation hooks to all the entity clients.
// In order to add hooks to a specific client, call: `client.Node.Use(...)`.
func (c *Client) Use(hooks ...Hook) {
	c.Announcement.Use(hooks...)
	c.MailBox.Use(hooks...)
	c.Notification.Use(hooks...)
	c.ReadState.Use(hooks...)
}

// AnnouncementClient is a client for the Announcement schema.
type AnnouncementClient struct {
	config
}

// NewAnnouncementClient returns a client for the Announcement from the given config.
func NewAnnouncementClient(c config) *AnnouncementClient {
	return &AnnouncementClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `announcement.Hooks(f(g(h())))`.
func (c *AnnouncementClient) Use(hooks ...Hook) {
	c.hooks.Announcement = append(c.hooks.Announcement, hooks...)
}

// Create returns a create builder for Announcement.
func (c *AnnouncementClient) Create() *AnnouncementCreate {
	mutation := newAnnouncementMutation(c.config, OpCreate)
	return &AnnouncementCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Announcement entities.
func (c *AnnouncementClient) CreateBulk(builders ...*AnnouncementCreate) *AnnouncementCreateBulk {
	return &AnnouncementCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Announcement.
func (c *AnnouncementClient) Update() *AnnouncementUpdate {
	mutation := newAnnouncementMutation(c.config, OpUpdate)
	return &AnnouncementUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *AnnouncementClient) UpdateOne(a *Announcement) *AnnouncementUpdateOne {
	mutation := newAnnouncementMutation(c.config, OpUpdateOne, withAnnouncement(a))
	return &AnnouncementUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *AnnouncementClient) UpdateOneID(id int) *AnnouncementUpdateOne {
	mutation := newAnnouncementMutation(c.config, OpUpdateOne, withAnnouncementID(id))
	return &AnnouncementUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Announcement.
func (c *AnnouncementClient) Delete() *AnnouncementDelete {
	mutation := newAnnouncementMutation(c.config, OpDelete)
	return &AnnouncementDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *AnnouncementClient) DeleteOne(a *Announcement) *AnnouncementDeleteOne {
	return c.DeleteOneID(a.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *AnnouncementClient) DeleteOneID(id int) *AnnouncementDeleteOne {
	builder := c.Delete().Where(announcement.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &AnnouncementDeleteOne{builder}
}

// Query returns a query builder for Announcement.
func (c *AnnouncementClient) Query() *AnnouncementQuery {
	return &AnnouncementQuery{
		config: c.config,
	}
}

// Get returns a Announcement entity by its id.
func (c *AnnouncementClient) Get(ctx context.Context, id int) (*Announcement, error) {
	return c.Query().Where(announcement.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *AnnouncementClient) GetX(ctx context.Context, id int) *Announcement {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *AnnouncementClient) Hooks() []Hook {
	return c.hooks.Announcement
}

// MailBoxClient is a client for the MailBox schema.
type MailBoxClient struct {
	config
}

// NewMailBoxClient returns a client for the MailBox from the given config.
func NewMailBoxClient(c config) *MailBoxClient {
	return &MailBoxClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `mailbox.Hooks(f(g(h())))`.
func (c *MailBoxClient) Use(hooks ...Hook) {
	c.hooks.MailBox = append(c.hooks.MailBox, hooks...)
}

// Create returns a create builder for MailBox.
func (c *MailBoxClient) Create() *MailBoxCreate {
	mutation := newMailBoxMutation(c.config, OpCreate)
	return &MailBoxCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of MailBox entities.
func (c *MailBoxClient) CreateBulk(builders ...*MailBoxCreate) *MailBoxCreateBulk {
	return &MailBoxCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for MailBox.
func (c *MailBoxClient) Update() *MailBoxUpdate {
	mutation := newMailBoxMutation(c.config, OpUpdate)
	return &MailBoxUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *MailBoxClient) UpdateOne(mb *MailBox) *MailBoxUpdateOne {
	mutation := newMailBoxMutation(c.config, OpUpdateOne, withMailBox(mb))
	return &MailBoxUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *MailBoxClient) UpdateOneID(id int) *MailBoxUpdateOne {
	mutation := newMailBoxMutation(c.config, OpUpdateOne, withMailBoxID(id))
	return &MailBoxUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for MailBox.
func (c *MailBoxClient) Delete() *MailBoxDelete {
	mutation := newMailBoxMutation(c.config, OpDelete)
	return &MailBoxDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *MailBoxClient) DeleteOne(mb *MailBox) *MailBoxDeleteOne {
	return c.DeleteOneID(mb.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *MailBoxClient) DeleteOneID(id int) *MailBoxDeleteOne {
	builder := c.Delete().Where(mailbox.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &MailBoxDeleteOne{builder}
}

// Query returns a query builder for MailBox.
func (c *MailBoxClient) Query() *MailBoxQuery {
	return &MailBoxQuery{
		config: c.config,
	}
}

// Get returns a MailBox entity by its id.
func (c *MailBoxClient) Get(ctx context.Context, id int) (*MailBox, error) {
	return c.Query().Where(mailbox.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *MailBoxClient) GetX(ctx context.Context, id int) *MailBox {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *MailBoxClient) Hooks() []Hook {
	return c.hooks.MailBox
}

// NotificationClient is a client for the Notification schema.
type NotificationClient struct {
	config
}

// NewNotificationClient returns a client for the Notification from the given config.
func NewNotificationClient(c config) *NotificationClient {
	return &NotificationClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `notification.Hooks(f(g(h())))`.
func (c *NotificationClient) Use(hooks ...Hook) {
	c.hooks.Notification = append(c.hooks.Notification, hooks...)
}

// Create returns a create builder for Notification.
func (c *NotificationClient) Create() *NotificationCreate {
	mutation := newNotificationMutation(c.config, OpCreate)
	return &NotificationCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Notification entities.
func (c *NotificationClient) CreateBulk(builders ...*NotificationCreate) *NotificationCreateBulk {
	return &NotificationCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Notification.
func (c *NotificationClient) Update() *NotificationUpdate {
	mutation := newNotificationMutation(c.config, OpUpdate)
	return &NotificationUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *NotificationClient) UpdateOne(n *Notification) *NotificationUpdateOne {
	mutation := newNotificationMutation(c.config, OpUpdateOne, withNotification(n))
	return &NotificationUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *NotificationClient) UpdateOneID(id int) *NotificationUpdateOne {
	mutation := newNotificationMutation(c.config, OpUpdateOne, withNotificationID(id))
	return &NotificationUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Notification.
func (c *NotificationClient) Delete() *NotificationDelete {
	mutation := newNotificationMutation(c.config, OpDelete)
	return &NotificationDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *NotificationClient) DeleteOne(n *Notification) *NotificationDeleteOne {
	return c.DeleteOneID(n.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *NotificationClient) DeleteOneID(id int) *NotificationDeleteOne {
	builder := c.Delete().Where(notification.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &NotificationDeleteOne{builder}
}

// Query returns a query builder for Notification.
func (c *NotificationClient) Query() *NotificationQuery {
	return &NotificationQuery{
		config: c.config,
	}
}

// Get returns a Notification entity by its id.
func (c *NotificationClient) Get(ctx context.Context, id int) (*Notification, error) {
	return c.Query().Where(notification.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *NotificationClient) GetX(ctx context.Context, id int) *Notification {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *NotificationClient) Hooks() []Hook {
	return c.hooks.Notification
}

// ReadStateClient is a client for the ReadState schema.
type ReadStateClient struct {
	config
}

// NewReadStateClient returns a client for the ReadState from the given config.
func NewReadStateClient(c config) *ReadStateClient {
	return &ReadStateClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `readstate.Hooks(f(g(h())))`.
func (c *ReadStateClient) Use(hooks ...Hook) {
	c.hooks.ReadState = append(c.hooks.ReadState, hooks...)
}

// Create returns a create builder for ReadState.
func (c *ReadStateClient) Create() *ReadStateCreate {
	mutation := newReadStateMutation(c.config, OpCreate)
	return &ReadStateCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of ReadState entities.
func (c *ReadStateClient) CreateBulk(builders ...*ReadStateCreate) *ReadStateCreateBulk {
	return &ReadStateCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for ReadState.
func (c *ReadStateClient) Update() *ReadStateUpdate {
	mutation := newReadStateMutation(c.config, OpUpdate)
	return &ReadStateUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *ReadStateClient) UpdateOne(rs *ReadState) *ReadStateUpdateOne {
	mutation := newReadStateMutation(c.config, OpUpdateOne, withReadState(rs))
	return &ReadStateUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *ReadStateClient) UpdateOneID(id int) *ReadStateUpdateOne {
	mutation := newReadStateMutation(c.config, OpUpdateOne, withReadStateID(id))
	return &ReadStateUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for ReadState.
func (c *ReadStateClient) Delete() *ReadStateDelete {
	mutation := newReadStateMutation(c.config, OpDelete)
	return &ReadStateDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *ReadStateClient) DeleteOne(rs *ReadState) *ReadStateDeleteOne {
	return c.DeleteOneID(rs.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *ReadStateClient) DeleteOneID(id int) *ReadStateDeleteOne {
	builder := c.Delete().Where(readstate.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &ReadStateDeleteOne{builder}
}

// Query returns a query builder for ReadState.
func (c *ReadStateClient) Query() *ReadStateQuery {
	return &ReadStateQuery{
		config: c.config,
	}
}

// Get returns a ReadState entity by its id.
func (c *ReadStateClient) Get(ctx context.Context, id int) (*ReadState, error) {
	return c.Query().Where(readstate.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *ReadStateClient) GetX(ctx context.Context, id int) *ReadState {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *ReadStateClient) Hooks() []Hook {
	return c.hooks.ReadState
}
