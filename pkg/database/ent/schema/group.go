package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/dialect/entsql"
	"github.com/facebook/ent/schema"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
	"time"
)

// Group holds the schema definition for the Group entity.
type Group struct {
	ent.Schema
}

// Annotations of the User.
func (Group) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "groups"},
	}
}

// Fields of the Group.
func (Group) Fields() []ent.Field {
	return []ent.Field{
		//field.String("ID").NotEmpty().Unique(),
		//field.String("ID").MaxLen(30).NotEmpty().Unique().Immutable(),
		field.String("id").Unique(),
		field.String("name").Unique(),
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the Group.
func (Group) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("users", User.Type),
	}
}