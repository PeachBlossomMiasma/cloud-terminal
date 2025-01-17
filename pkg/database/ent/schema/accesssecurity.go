package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// AccessSecurity holds the schema definition for the AccessSecurity entity.
type AccessSecurity struct {
	ent.Schema
}

// Annotations of the User.
func (AccessSecurity) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "accessSecuritys"},
	}
}

// Fields of the AccessSecurity.
func (AccessSecurity) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").Unique(),
		field.String("rule"),
		field.String("ip"),
		field.String("source"),
		field.Int64("priority"),
		//field.Time("created_at").Default(time.Now),
		//field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the AccessSecurity.
func (AccessSecurity) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("assets", Asset.Type).Required(),
	}
}

// Mixin xxxx
func (AccessSecurity) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}
