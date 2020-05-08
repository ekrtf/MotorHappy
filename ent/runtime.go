// Code generated by entc, DO NOT EDIT.

package ent

import (
	"github.com/ekrtf/MotorHappy/ent/manufacturer"
	"github.com/ekrtf/MotorHappy/ent/model"
	"github.com/ekrtf/MotorHappy/ent/schema"
)

// The init function reads all schema descriptors with runtime
// code (default values, validators or hooks) and stitches it
// to their package variables.
func init() {
	manufacturerFields := schema.Manufacturer{}.Fields()
	_ = manufacturerFields
	// manufacturerDescName is the schema descriptor for name field.
	manufacturerDescName := manufacturerFields[0].Descriptor()
	// manufacturer.DefaultName holds the default value on creation for the name field.
	manufacturer.DefaultName = manufacturerDescName.Default.(string)
	modelFields := schema.Model{}.Fields()
	_ = modelFields
	// modelDescName is the schema descriptor for name field.
	modelDescName := modelFields[0].Descriptor()
	// model.DefaultName holds the default value on creation for the name field.
	model.DefaultName = modelDescName.Default.(string)
}