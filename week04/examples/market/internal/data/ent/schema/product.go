package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
)

// Product holds the schema definition for the Product entity.
type Product struct {
	ent.Schema
}

// Fields of the Post.
func (Product) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id"),
		field.String("name"),
		field.String("desc"),
		field.Time("created_at").
			Default(time.Now).SchemaType(map[string]string{
			dialect.MySQL: "datetime",
		}),
		field.Time("updated_at").
			Default(time.Now).SchemaType(map[string]string{
			dialect.MySQL: "datetime",
		}),
	}
}

// Edges of the Post.
func (Product) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("comments", Comment.Type),
		edge.From("tags", Tag.Type).
			Ref("posts"),
	}
}
