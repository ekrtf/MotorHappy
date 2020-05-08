// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"github.com/ekrtf/MotorHappy/ent/manufacturer"
	"github.com/facebookincubator/ent/dialect/sql"
)

// Manufacturer is the model entity for the Manufacturer schema.
type Manufacturer struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Manufacturer) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},  // id
		&sql.NullString{}, // name
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Manufacturer fields.
func (m *Manufacturer) assignValues(values ...interface{}) error {
	if m, n := len(values), len(manufacturer.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	m.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field name", values[0])
	} else if value.Valid {
		m.Name = value.String
	}
	return nil
}

// Update returns a builder for updating this Manufacturer.
// Note that, you need to call Manufacturer.Unwrap() before calling this method, if this Manufacturer
// was returned from a transaction, and the transaction was committed or rolled back.
func (m *Manufacturer) Update() *ManufacturerUpdateOne {
	return (&ManufacturerClient{config: m.config}).UpdateOne(m)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (m *Manufacturer) Unwrap() *Manufacturer {
	tx, ok := m.config.driver.(*txDriver)
	if !ok {
		panic("ent: Manufacturer is not a transactional entity")
	}
	m.config.driver = tx.drv
	return m
}

// String implements the fmt.Stringer.
func (m *Manufacturer) String() string {
	var builder strings.Builder
	builder.WriteString("Manufacturer(")
	builder.WriteString(fmt.Sprintf("id=%v", m.ID))
	builder.WriteString(", name=")
	builder.WriteString(m.Name)
	builder.WriteByte(')')
	return builder.String()
}

// Manufacturers is a parsable slice of Manufacturer.
type Manufacturers []*Manufacturer

func (m Manufacturers) config(cfg config) {
	for _i := range m {
		m[_i].config = cfg
	}
}
