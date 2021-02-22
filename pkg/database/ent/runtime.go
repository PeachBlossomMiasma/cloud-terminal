// Code generated by entc, DO NOT EDIT.

package ent

import (
	"time"

	"github.com/willie-lin/cloud-terminal/pkg/database/ent/asset"
	"github.com/willie-lin/cloud-terminal/pkg/database/ent/command"
	"github.com/willie-lin/cloud-terminal/pkg/database/ent/credential"
	"github.com/willie-lin/cloud-terminal/pkg/database/ent/schema"
	"github.com/willie-lin/cloud-terminal/pkg/database/ent/session"
	"github.com/willie-lin/cloud-terminal/pkg/database/ent/user"
	"github.com/willie-lin/cloud-terminal/pkg/database/ent/usergroup"
	"github.com/willie-lin/cloud-terminal/pkg/database/ent/verification"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	assetFields := schema.Asset{}.Fields()
	_ = assetFields
	// assetDescCreatedAt is the schema descriptor for created_at field.
	assetDescCreatedAt := assetFields[13].Descriptor()
	// asset.DefaultCreatedAt holds the default value on creation for the created_at field.
	asset.DefaultCreatedAt = assetDescCreatedAt.Default.(func() time.Time)
	// assetDescUpdatedAt is the schema descriptor for updated_at field.
	assetDescUpdatedAt := assetFields[14].Descriptor()
	// asset.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	asset.DefaultUpdatedAt = assetDescUpdatedAt.Default.(func() time.Time)
	// asset.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	asset.UpdateDefaultUpdatedAt = assetDescUpdatedAt.UpdateDefault.(func() time.Time)
	commandFields := schema.Command{}.Fields()
	_ = commandFields
	// commandDescName is the schema descriptor for name field.
	commandDescName := commandFields[1].Descriptor()
	// command.NameValidator is a validator for the "name" field. It is called by the builders before save.
	command.NameValidator = commandDescName.Validators[0].(func(string) error)
	// commandDescCreatedAt is the schema descriptor for created_at field.
	commandDescCreatedAt := commandFields[3].Descriptor()
	// command.DefaultCreatedAt holds the default value on creation for the created_at field.
	command.DefaultCreatedAt = commandDescCreatedAt.Default.(func() time.Time)
	// commandDescUpdatedAt is the schema descriptor for updated_at field.
	commandDescUpdatedAt := commandFields[4].Descriptor()
	// command.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	command.DefaultUpdatedAt = commandDescUpdatedAt.Default.(func() time.Time)
	// command.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	command.UpdateDefaultUpdatedAt = commandDescUpdatedAt.UpdateDefault.(func() time.Time)
	credentialFields := schema.Credential{}.Fields()
	_ = credentialFields
	// credentialDescCreatedAt is the schema descriptor for created_at field.
	credentialDescCreatedAt := credentialFields[7].Descriptor()
	// credential.DefaultCreatedAt holds the default value on creation for the created_at field.
	credential.DefaultCreatedAt = credentialDescCreatedAt.Default.(func() time.Time)
	// credentialDescUpdatedAt is the schema descriptor for updated_at field.
	credentialDescUpdatedAt := credentialFields[8].Descriptor()
	// credential.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	credential.DefaultUpdatedAt = credentialDescUpdatedAt.Default.(func() time.Time)
	// credential.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	credential.UpdateDefaultUpdatedAt = credentialDescUpdatedAt.UpdateDefault.(func() time.Time)
	sessionFields := schema.Session{}.Fields()
	_ = sessionFields
	// sessionDescConnected is the schema descriptor for connected field.
	sessionDescConnected := sessionFields[18].Descriptor()
	// session.DefaultConnected holds the default value on creation for the connected field.
	session.DefaultConnected = sessionDescConnected.Default.(func() time.Time)
	// sessionDescDisconnected is the schema descriptor for disconnected field.
	sessionDescDisconnected := sessionFields[19].Descriptor()
	// session.DefaultDisconnected holds the default value on creation for the disconnected field.
	session.DefaultDisconnected = sessionDescDisconnected.Default.(func() time.Time)
	// session.UpdateDefaultDisconnected holds the default value on update for the disconnected field.
	session.UpdateDefaultDisconnected = sessionDescDisconnected.UpdateDefault.(func() time.Time)
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescCreatedAt is the schema descriptor for created_at field.
	userDescCreatedAt := userFields[8].Descriptor()
	// user.DefaultCreatedAt holds the default value on creation for the created_at field.
	user.DefaultCreatedAt = userDescCreatedAt.Default.(func() time.Time)
	// userDescUpdatedAt is the schema descriptor for updated_at field.
	userDescUpdatedAt := userFields[9].Descriptor()
	// user.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	user.DefaultUpdatedAt = userDescUpdatedAt.Default.(func() time.Time)
	// user.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	user.UpdateDefaultUpdatedAt = userDescUpdatedAt.UpdateDefault.(func() time.Time)
	usergroupFields := schema.UserGroup{}.Fields()
	_ = usergroupFields
	// usergroupDescCreatedAt is the schema descriptor for created_at field.
	usergroupDescCreatedAt := usergroupFields[2].Descriptor()
	// usergroup.DefaultCreatedAt holds the default value on creation for the created_at field.
	usergroup.DefaultCreatedAt = usergroupDescCreatedAt.Default.(func() time.Time)
	// usergroupDescUpdatedAt is the schema descriptor for updated_at field.
	usergroupDescUpdatedAt := usergroupFields[3].Descriptor()
	// usergroup.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	usergroup.DefaultUpdatedAt = usergroupDescUpdatedAt.Default.(func() time.Time)
	// usergroup.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	usergroup.UpdateDefaultUpdatedAt = usergroupDescUpdatedAt.UpdateDefault.(func() time.Time)
	verificationFields := schema.Verification{}.Fields()
	_ = verificationFields
	// verificationDescLoginTime is the schema descriptor for login_time field.
	verificationDescLoginTime := verificationFields[3].Descriptor()
	// verification.DefaultLoginTime holds the default value on creation for the login_time field.
	verification.DefaultLoginTime = verificationDescLoginTime.Default.(func() time.Time)
	// verificationDescLogoutTime is the schema descriptor for logout_time field.
	verificationDescLogoutTime := verificationFields[4].Descriptor()
	// verification.DefaultLogoutTime holds the default value on creation for the logout_time field.
	verification.DefaultLogoutTime = verificationDescLogoutTime.Default.(func() time.Time)
	// verification.UpdateDefaultLogoutTime holds the default value on update for the logout_time field.
	verification.UpdateDefaultLogoutTime = verificationDescLogoutTime.UpdateDefault.(func() time.Time)
}
