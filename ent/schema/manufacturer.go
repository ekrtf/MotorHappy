package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/field"
)

// Manufacturer holds the schema definition for the Manufacturer entity.
type Manufacturer struct {
	ent.Schema
}

// Fields of the Manufacturer.
func (Manufacturer) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			Default("unknown"),
	}
}

// Edges of the Manufacturer.
func (Manufacturer) Edges() []ent.Edge {
	return nil
}
