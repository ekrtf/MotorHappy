package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/field"
)

// Model holds the schema definition for the Model entity.
type Model struct {
	ent.Schema
}

// Fields of the Model.
func (Model) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			Default("unknown"),
	}
}

// Edges of the Model.
func (Model) Edges() []ent.Edge {
	return nil
}
