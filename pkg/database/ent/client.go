// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"log"

	"github.com/willie-lin/cloud-terminal/pkg/database/ent/migrate"

	"github.com/willie-lin/cloud-terminal/pkg/database/ent/asset"
	"github.com/willie-lin/cloud-terminal/pkg/database/ent/command"
	"github.com/willie-lin/cloud-terminal/pkg/database/ent/credential"
	"github.com/willie-lin/cloud-terminal/pkg/database/ent/property"
	"github.com/willie-lin/cloud-terminal/pkg/database/ent/resourcesharer"
	"github.com/willie-lin/cloud-terminal/pkg/database/ent/session"
	"github.com/willie-lin/cloud-terminal/pkg/database/ent/user"
	"github.com/willie-lin/cloud-terminal/pkg/database/ent/usergroup"
	"github.com/willie-lin/cloud-terminal/pkg/database/ent/verification"

	"github.com/facebook/ent/dialect"
	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Asset is the client for interacting with the Asset builders.
	Asset *AssetClient
	// Command is the client for interacting with the Command builders.
	Command *CommandClient
	// Credential is the client for interacting with the Credential builders.
	Credential *CredentialClient
	// Property is the client for interacting with the Property builders.
	Property *PropertyClient
	// ResourceSharer is the client for interacting with the ResourceSharer builders.
	ResourceSharer *ResourceSharerClient
	// Session is the client for interacting with the Session builders.
	Session *SessionClient
	// User is the client for interacting with the User builders.
	User *UserClient
	// UserGroup is the client for interacting with the UserGroup builders.
	UserGroup *UserGroupClient
	// Verification is the client for interacting with the Verification builders.
	Verification *VerificationClient
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
	c.Asset = NewAssetClient(c.config)
	c.Command = NewCommandClient(c.config)
	c.Credential = NewCredentialClient(c.config)
	c.Property = NewPropertyClient(c.config)
	c.ResourceSharer = NewResourceSharerClient(c.config)
	c.Session = NewSessionClient(c.config)
	c.User = NewUserClient(c.config)
	c.UserGroup = NewUserGroupClient(c.config)
	c.Verification = NewVerificationClient(c.config)
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
		return nil, fmt.Errorf("ent: starting a transaction: %v", err)
	}
	cfg := config{driver: tx, log: c.log, debug: c.debug, hooks: c.hooks}
	return &Tx{
		ctx:            ctx,
		config:         cfg,
		Asset:          NewAssetClient(cfg),
		Command:        NewCommandClient(cfg),
		Credential:     NewCredentialClient(cfg),
		Property:       NewPropertyClient(cfg),
		ResourceSharer: NewResourceSharerClient(cfg),
		Session:        NewSessionClient(cfg),
		User:           NewUserClient(cfg),
		UserGroup:      NewUserGroupClient(cfg),
		Verification:   NewVerificationClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with specified options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
	}
	tx, err := c.driver.(*sql.Driver).BeginTx(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %v", err)
	}
	cfg := config{driver: &txDriver{tx: tx, drv: c.driver}, log: c.log, debug: c.debug, hooks: c.hooks}
	return &Tx{
		config:         cfg,
		Asset:          NewAssetClient(cfg),
		Command:        NewCommandClient(cfg),
		Credential:     NewCredentialClient(cfg),
		Property:       NewPropertyClient(cfg),
		ResourceSharer: NewResourceSharerClient(cfg),
		Session:        NewSessionClient(cfg),
		User:           NewUserClient(cfg),
		UserGroup:      NewUserGroupClient(cfg),
		Verification:   NewVerificationClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Asset.
//		Query().
//		Count(ctx)
//
func (c *Client) Debug() *Client {
	if c.debug {
		return c
	}
	cfg := config{driver: dialect.Debug(c.driver, c.log), log: c.log, debug: true, hooks: c.hooks}
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
	c.Asset.Use(hooks...)
	c.Command.Use(hooks...)
	c.Credential.Use(hooks...)
	c.Property.Use(hooks...)
	c.ResourceSharer.Use(hooks...)
	c.Session.Use(hooks...)
	c.User.Use(hooks...)
	c.UserGroup.Use(hooks...)
	c.Verification.Use(hooks...)
}

// AssetClient is a client for the Asset schema.
type AssetClient struct {
	config
}

// NewAssetClient returns a client for the Asset from the given config.
func NewAssetClient(c config) *AssetClient {
	return &AssetClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `asset.Hooks(f(g(h())))`.
func (c *AssetClient) Use(hooks ...Hook) {
	c.hooks.Asset = append(c.hooks.Asset, hooks...)
}

// Create returns a create builder for Asset.
func (c *AssetClient) Create() *AssetCreate {
	mutation := newAssetMutation(c.config, OpCreate)
	return &AssetCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Asset entities.
func (c *AssetClient) CreateBulk(builders ...*AssetCreate) *AssetCreateBulk {
	return &AssetCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Asset.
func (c *AssetClient) Update() *AssetUpdate {
	mutation := newAssetMutation(c.config, OpUpdate)
	return &AssetUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *AssetClient) UpdateOne(a *Asset) *AssetUpdateOne {
	mutation := newAssetMutation(c.config, OpUpdateOne, withAsset(a))
	return &AssetUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *AssetClient) UpdateOneID(id string) *AssetUpdateOne {
	mutation := newAssetMutation(c.config, OpUpdateOne, withAssetID(id))
	return &AssetUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Asset.
func (c *AssetClient) Delete() *AssetDelete {
	mutation := newAssetMutation(c.config, OpDelete)
	return &AssetDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *AssetClient) DeleteOne(a *Asset) *AssetDeleteOne {
	return c.DeleteOneID(a.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *AssetClient) DeleteOneID(id string) *AssetDeleteOne {
	builder := c.Delete().Where(asset.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &AssetDeleteOne{builder}
}

// Query returns a query builder for Asset.
func (c *AssetClient) Query() *AssetQuery {
	return &AssetQuery{config: c.config}
}

// Get returns a Asset entity by its id.
func (c *AssetClient) Get(ctx context.Context, id string) (*Asset, error) {
	return c.Query().Where(asset.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *AssetClient) GetX(ctx context.Context, id string) *Asset {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QuerySessions queries the sessions edge of a Asset.
func (c *AssetClient) QuerySessions(a *Asset) *SessionQuery {
	query := &SessionQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := a.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(asset.Table, asset.FieldID, id),
			sqlgraph.To(session.Table, session.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, asset.SessionsTable, asset.SessionsColumn),
		)
		fromV = sqlgraph.Neighbors(a.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *AssetClient) Hooks() []Hook {
	return c.hooks.Asset
}

// CommandClient is a client for the Command schema.
type CommandClient struct {
	config
}

// NewCommandClient returns a client for the Command from the given config.
func NewCommandClient(c config) *CommandClient {
	return &CommandClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `command.Hooks(f(g(h())))`.
func (c *CommandClient) Use(hooks ...Hook) {
	c.hooks.Command = append(c.hooks.Command, hooks...)
}

// Create returns a create builder for Command.
func (c *CommandClient) Create() *CommandCreate {
	mutation := newCommandMutation(c.config, OpCreate)
	return &CommandCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Command entities.
func (c *CommandClient) CreateBulk(builders ...*CommandCreate) *CommandCreateBulk {
	return &CommandCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Command.
func (c *CommandClient) Update() *CommandUpdate {
	mutation := newCommandMutation(c.config, OpUpdate)
	return &CommandUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *CommandClient) UpdateOne(co *Command) *CommandUpdateOne {
	mutation := newCommandMutation(c.config, OpUpdateOne, withCommand(co))
	return &CommandUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *CommandClient) UpdateOneID(id string) *CommandUpdateOne {
	mutation := newCommandMutation(c.config, OpUpdateOne, withCommandID(id))
	return &CommandUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Command.
func (c *CommandClient) Delete() *CommandDelete {
	mutation := newCommandMutation(c.config, OpDelete)
	return &CommandDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *CommandClient) DeleteOne(co *Command) *CommandDeleteOne {
	return c.DeleteOneID(co.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *CommandClient) DeleteOneID(id string) *CommandDeleteOne {
	builder := c.Delete().Where(command.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &CommandDeleteOne{builder}
}

// Query returns a query builder for Command.
func (c *CommandClient) Query() *CommandQuery {
	return &CommandQuery{config: c.config}
}

// Get returns a Command entity by its id.
func (c *CommandClient) Get(ctx context.Context, id string) (*Command, error) {
	return c.Query().Where(command.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *CommandClient) GetX(ctx context.Context, id string) *Command {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *CommandClient) Hooks() []Hook {
	return c.hooks.Command
}

// CredentialClient is a client for the Credential schema.
type CredentialClient struct {
	config
}

// NewCredentialClient returns a client for the Credential from the given config.
func NewCredentialClient(c config) *CredentialClient {
	return &CredentialClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `credential.Hooks(f(g(h())))`.
func (c *CredentialClient) Use(hooks ...Hook) {
	c.hooks.Credential = append(c.hooks.Credential, hooks...)
}

// Create returns a create builder for Credential.
func (c *CredentialClient) Create() *CredentialCreate {
	mutation := newCredentialMutation(c.config, OpCreate)
	return &CredentialCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Credential entities.
func (c *CredentialClient) CreateBulk(builders ...*CredentialCreate) *CredentialCreateBulk {
	return &CredentialCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Credential.
func (c *CredentialClient) Update() *CredentialUpdate {
	mutation := newCredentialMutation(c.config, OpUpdate)
	return &CredentialUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *CredentialClient) UpdateOne(cr *Credential) *CredentialUpdateOne {
	mutation := newCredentialMutation(c.config, OpUpdateOne, withCredential(cr))
	return &CredentialUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *CredentialClient) UpdateOneID(id string) *CredentialUpdateOne {
	mutation := newCredentialMutation(c.config, OpUpdateOne, withCredentialID(id))
	return &CredentialUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Credential.
func (c *CredentialClient) Delete() *CredentialDelete {
	mutation := newCredentialMutation(c.config, OpDelete)
	return &CredentialDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *CredentialClient) DeleteOne(cr *Credential) *CredentialDeleteOne {
	return c.DeleteOneID(cr.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *CredentialClient) DeleteOneID(id string) *CredentialDeleteOne {
	builder := c.Delete().Where(credential.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &CredentialDeleteOne{builder}
}

// Query returns a query builder for Credential.
func (c *CredentialClient) Query() *CredentialQuery {
	return &CredentialQuery{config: c.config}
}

// Get returns a Credential entity by its id.
func (c *CredentialClient) Get(ctx context.Context, id string) (*Credential, error) {
	return c.Query().Where(credential.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *CredentialClient) GetX(ctx context.Context, id string) *Credential {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *CredentialClient) Hooks() []Hook {
	return c.hooks.Credential
}

// PropertyClient is a client for the Property schema.
type PropertyClient struct {
	config
}

// NewPropertyClient returns a client for the Property from the given config.
func NewPropertyClient(c config) *PropertyClient {
	return &PropertyClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `property.Hooks(f(g(h())))`.
func (c *PropertyClient) Use(hooks ...Hook) {
	c.hooks.Property = append(c.hooks.Property, hooks...)
}

// Create returns a create builder for Property.
func (c *PropertyClient) Create() *PropertyCreate {
	mutation := newPropertyMutation(c.config, OpCreate)
	return &PropertyCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Property entities.
func (c *PropertyClient) CreateBulk(builders ...*PropertyCreate) *PropertyCreateBulk {
	return &PropertyCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Property.
func (c *PropertyClient) Update() *PropertyUpdate {
	mutation := newPropertyMutation(c.config, OpUpdate)
	return &PropertyUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *PropertyClient) UpdateOne(pr *Property) *PropertyUpdateOne {
	mutation := newPropertyMutation(c.config, OpUpdateOne, withProperty(pr))
	return &PropertyUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *PropertyClient) UpdateOneID(id int) *PropertyUpdateOne {
	mutation := newPropertyMutation(c.config, OpUpdateOne, withPropertyID(id))
	return &PropertyUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Property.
func (c *PropertyClient) Delete() *PropertyDelete {
	mutation := newPropertyMutation(c.config, OpDelete)
	return &PropertyDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *PropertyClient) DeleteOne(pr *Property) *PropertyDeleteOne {
	return c.DeleteOneID(pr.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *PropertyClient) DeleteOneID(id int) *PropertyDeleteOne {
	builder := c.Delete().Where(property.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &PropertyDeleteOne{builder}
}

// Query returns a query builder for Property.
func (c *PropertyClient) Query() *PropertyQuery {
	return &PropertyQuery{config: c.config}
}

// Get returns a Property entity by its id.
func (c *PropertyClient) Get(ctx context.Context, id int) (*Property, error) {
	return c.Query().Where(property.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *PropertyClient) GetX(ctx context.Context, id int) *Property {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *PropertyClient) Hooks() []Hook {
	return c.hooks.Property
}

// ResourceSharerClient is a client for the ResourceSharer schema.
type ResourceSharerClient struct {
	config
}

// NewResourceSharerClient returns a client for the ResourceSharer from the given config.
func NewResourceSharerClient(c config) *ResourceSharerClient {
	return &ResourceSharerClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `resourcesharer.Hooks(f(g(h())))`.
func (c *ResourceSharerClient) Use(hooks ...Hook) {
	c.hooks.ResourceSharer = append(c.hooks.ResourceSharer, hooks...)
}

// Create returns a create builder for ResourceSharer.
func (c *ResourceSharerClient) Create() *ResourceSharerCreate {
	mutation := newResourceSharerMutation(c.config, OpCreate)
	return &ResourceSharerCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of ResourceSharer entities.
func (c *ResourceSharerClient) CreateBulk(builders ...*ResourceSharerCreate) *ResourceSharerCreateBulk {
	return &ResourceSharerCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for ResourceSharer.
func (c *ResourceSharerClient) Update() *ResourceSharerUpdate {
	mutation := newResourceSharerMutation(c.config, OpUpdate)
	return &ResourceSharerUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *ResourceSharerClient) UpdateOne(rs *ResourceSharer) *ResourceSharerUpdateOne {
	mutation := newResourceSharerMutation(c.config, OpUpdateOne, withResourceSharer(rs))
	return &ResourceSharerUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *ResourceSharerClient) UpdateOneID(id string) *ResourceSharerUpdateOne {
	mutation := newResourceSharerMutation(c.config, OpUpdateOne, withResourceSharerID(id))
	return &ResourceSharerUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for ResourceSharer.
func (c *ResourceSharerClient) Delete() *ResourceSharerDelete {
	mutation := newResourceSharerMutation(c.config, OpDelete)
	return &ResourceSharerDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *ResourceSharerClient) DeleteOne(rs *ResourceSharer) *ResourceSharerDeleteOne {
	return c.DeleteOneID(rs.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *ResourceSharerClient) DeleteOneID(id string) *ResourceSharerDeleteOne {
	builder := c.Delete().Where(resourcesharer.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &ResourceSharerDeleteOne{builder}
}

// Query returns a query builder for ResourceSharer.
func (c *ResourceSharerClient) Query() *ResourceSharerQuery {
	return &ResourceSharerQuery{config: c.config}
}

// Get returns a ResourceSharer entity by its id.
func (c *ResourceSharerClient) Get(ctx context.Context, id string) (*ResourceSharer, error) {
	return c.Query().Where(resourcesharer.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *ResourceSharerClient) GetX(ctx context.Context, id string) *ResourceSharer {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *ResourceSharerClient) Hooks() []Hook {
	return c.hooks.ResourceSharer
}

// SessionClient is a client for the Session schema.
type SessionClient struct {
	config
}

// NewSessionClient returns a client for the Session from the given config.
func NewSessionClient(c config) *SessionClient {
	return &SessionClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `session.Hooks(f(g(h())))`.
func (c *SessionClient) Use(hooks ...Hook) {
	c.hooks.Session = append(c.hooks.Session, hooks...)
}

// Create returns a create builder for Session.
func (c *SessionClient) Create() *SessionCreate {
	mutation := newSessionMutation(c.config, OpCreate)
	return &SessionCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Session entities.
func (c *SessionClient) CreateBulk(builders ...*SessionCreate) *SessionCreateBulk {
	return &SessionCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Session.
func (c *SessionClient) Update() *SessionUpdate {
	mutation := newSessionMutation(c.config, OpUpdate)
	return &SessionUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *SessionClient) UpdateOne(s *Session) *SessionUpdateOne {
	mutation := newSessionMutation(c.config, OpUpdateOne, withSession(s))
	return &SessionUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *SessionClient) UpdateOneID(id string) *SessionUpdateOne {
	mutation := newSessionMutation(c.config, OpUpdateOne, withSessionID(id))
	return &SessionUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Session.
func (c *SessionClient) Delete() *SessionDelete {
	mutation := newSessionMutation(c.config, OpDelete)
	return &SessionDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *SessionClient) DeleteOne(s *Session) *SessionDeleteOne {
	return c.DeleteOneID(s.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *SessionClient) DeleteOneID(id string) *SessionDeleteOne {
	builder := c.Delete().Where(session.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &SessionDeleteOne{builder}
}

// Query returns a query builder for Session.
func (c *SessionClient) Query() *SessionQuery {
	return &SessionQuery{config: c.config}
}

// Get returns a Session entity by its id.
func (c *SessionClient) Get(ctx context.Context, id string) (*Session, error) {
	return c.Query().Where(session.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *SessionClient) GetX(ctx context.Context, id string) *Session {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryAssets queries the assets edge of a Session.
func (c *SessionClient) QueryAssets(s *Session) *AssetQuery {
	query := &AssetQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := s.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(session.Table, session.FieldID, id),
			sqlgraph.To(asset.Table, asset.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, session.AssetsTable, session.AssetsColumn),
		)
		fromV = sqlgraph.Neighbors(s.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *SessionClient) Hooks() []Hook {
	return c.hooks.Session
}

// UserClient is a client for the User schema.
type UserClient struct {
	config
}

// NewUserClient returns a client for the User from the given config.
func NewUserClient(c config) *UserClient {
	return &UserClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `user.Hooks(f(g(h())))`.
func (c *UserClient) Use(hooks ...Hook) {
	c.hooks.User = append(c.hooks.User, hooks...)
}

// Create returns a create builder for User.
func (c *UserClient) Create() *UserCreate {
	mutation := newUserMutation(c.config, OpCreate)
	return &UserCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of User entities.
func (c *UserClient) CreateBulk(builders ...*UserCreate) *UserCreateBulk {
	return &UserCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for User.
func (c *UserClient) Update() *UserUpdate {
	mutation := newUserMutation(c.config, OpUpdate)
	return &UserUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *UserClient) UpdateOne(u *User) *UserUpdateOne {
	mutation := newUserMutation(c.config, OpUpdateOne, withUser(u))
	return &UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *UserClient) UpdateOneID(id string) *UserUpdateOne {
	mutation := newUserMutation(c.config, OpUpdateOne, withUserID(id))
	return &UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for User.
func (c *UserClient) Delete() *UserDelete {
	mutation := newUserMutation(c.config, OpDelete)
	return &UserDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *UserClient) DeleteOne(u *User) *UserDeleteOne {
	return c.DeleteOneID(u.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *UserClient) DeleteOneID(id string) *UserDeleteOne {
	builder := c.Delete().Where(user.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &UserDeleteOne{builder}
}

// Query returns a query builder for User.
func (c *UserClient) Query() *UserQuery {
	return &UserQuery{config: c.config}
}

// Get returns a User entity by its id.
func (c *UserClient) Get(ctx context.Context, id string) (*User, error) {
	return c.Query().Where(user.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *UserClient) GetX(ctx context.Context, id string) *User {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryUserGroups queries the user_groups edge of a User.
func (c *UserClient) QueryUserGroups(u *User) *UserGroupQuery {
	query := &UserGroupQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := u.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, id),
			sqlgraph.To(usergroup.Table, usergroup.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, user.UserGroupsTable, user.UserGroupsPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(u.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *UserClient) Hooks() []Hook {
	return c.hooks.User
}

// UserGroupClient is a client for the UserGroup schema.
type UserGroupClient struct {
	config
}

// NewUserGroupClient returns a client for the UserGroup from the given config.
func NewUserGroupClient(c config) *UserGroupClient {
	return &UserGroupClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `usergroup.Hooks(f(g(h())))`.
func (c *UserGroupClient) Use(hooks ...Hook) {
	c.hooks.UserGroup = append(c.hooks.UserGroup, hooks...)
}

// Create returns a create builder for UserGroup.
func (c *UserGroupClient) Create() *UserGroupCreate {
	mutation := newUserGroupMutation(c.config, OpCreate)
	return &UserGroupCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of UserGroup entities.
func (c *UserGroupClient) CreateBulk(builders ...*UserGroupCreate) *UserGroupCreateBulk {
	return &UserGroupCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for UserGroup.
func (c *UserGroupClient) Update() *UserGroupUpdate {
	mutation := newUserGroupMutation(c.config, OpUpdate)
	return &UserGroupUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *UserGroupClient) UpdateOne(ug *UserGroup) *UserGroupUpdateOne {
	mutation := newUserGroupMutation(c.config, OpUpdateOne, withUserGroup(ug))
	return &UserGroupUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *UserGroupClient) UpdateOneID(id string) *UserGroupUpdateOne {
	mutation := newUserGroupMutation(c.config, OpUpdateOne, withUserGroupID(id))
	return &UserGroupUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for UserGroup.
func (c *UserGroupClient) Delete() *UserGroupDelete {
	mutation := newUserGroupMutation(c.config, OpDelete)
	return &UserGroupDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *UserGroupClient) DeleteOne(ug *UserGroup) *UserGroupDeleteOne {
	return c.DeleteOneID(ug.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *UserGroupClient) DeleteOneID(id string) *UserGroupDeleteOne {
	builder := c.Delete().Where(usergroup.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &UserGroupDeleteOne{builder}
}

// Query returns a query builder for UserGroup.
func (c *UserGroupClient) Query() *UserGroupQuery {
	return &UserGroupQuery{config: c.config}
}

// Get returns a UserGroup entity by its id.
func (c *UserGroupClient) Get(ctx context.Context, id string) (*UserGroup, error) {
	return c.Query().Where(usergroup.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *UserGroupClient) GetX(ctx context.Context, id string) *UserGroup {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryUsers queries the users edge of a UserGroup.
func (c *UserGroupClient) QueryUsers(ug *UserGroup) *UserQuery {
	query := &UserQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := ug.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(usergroup.Table, usergroup.FieldID, id),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, usergroup.UsersTable, usergroup.UsersPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(ug.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *UserGroupClient) Hooks() []Hook {
	return c.hooks.UserGroup
}

// VerificationClient is a client for the Verification schema.
type VerificationClient struct {
	config
}

// NewVerificationClient returns a client for the Verification from the given config.
func NewVerificationClient(c config) *VerificationClient {
	return &VerificationClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `verification.Hooks(f(g(h())))`.
func (c *VerificationClient) Use(hooks ...Hook) {
	c.hooks.Verification = append(c.hooks.Verification, hooks...)
}

// Create returns a create builder for Verification.
func (c *VerificationClient) Create() *VerificationCreate {
	mutation := newVerificationMutation(c.config, OpCreate)
	return &VerificationCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Verification entities.
func (c *VerificationClient) CreateBulk(builders ...*VerificationCreate) *VerificationCreateBulk {
	return &VerificationCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Verification.
func (c *VerificationClient) Update() *VerificationUpdate {
	mutation := newVerificationMutation(c.config, OpUpdate)
	return &VerificationUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *VerificationClient) UpdateOne(v *Verification) *VerificationUpdateOne {
	mutation := newVerificationMutation(c.config, OpUpdateOne, withVerification(v))
	return &VerificationUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *VerificationClient) UpdateOneID(id string) *VerificationUpdateOne {
	mutation := newVerificationMutation(c.config, OpUpdateOne, withVerificationID(id))
	return &VerificationUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Verification.
func (c *VerificationClient) Delete() *VerificationDelete {
	mutation := newVerificationMutation(c.config, OpDelete)
	return &VerificationDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *VerificationClient) DeleteOne(v *Verification) *VerificationDeleteOne {
	return c.DeleteOneID(v.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *VerificationClient) DeleteOneID(id string) *VerificationDeleteOne {
	builder := c.Delete().Where(verification.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &VerificationDeleteOne{builder}
}

// Query returns a query builder for Verification.
func (c *VerificationClient) Query() *VerificationQuery {
	return &VerificationQuery{config: c.config}
}

// Get returns a Verification entity by its id.
func (c *VerificationClient) Get(ctx context.Context, id string) (*Verification, error) {
	return c.Query().Where(verification.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *VerificationClient) GetX(ctx context.Context, id string) *Verification {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryUsers queries the users edge of a Verification.
func (c *VerificationClient) QueryUsers(v *Verification) *UserQuery {
	query := &UserQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := v.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(verification.Table, verification.FieldID, id),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, verification.UsersTable, verification.UsersColumn),
		)
		fromV = sqlgraph.Neighbors(v.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *VerificationClient) Hooks() []Hook {
	return c.hooks.Verification
}
