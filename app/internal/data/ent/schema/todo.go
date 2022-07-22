package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type Todo struct {
	ent.Schema
}

// Fields of the User.
func (Todo) Fields() []ent.Field {
	return []ent.Field{
		field.String("todo").NotEmpty(),
		field.Time("due_at").Optional(),
		field.Time("notify_at").Optional(),
		field.Time("created_at").
			Immutable().
			Default(time.Now),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now),
	}
}

// Edges of the User.
func (Todo) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("users", User.Type).Ref("todos").Unique(),
	}
}
