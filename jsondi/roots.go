package jsondi

import (
	"github.com/binadel/jsonw/jsoni"
	"github.com/mailru/easyjson/jwriter"
)

// RootObject represents a json root object.
type RootObject []Field

// New creates a new root object.
func New(fields ...Field) RootObject {
	return fields
}

// Build encodes the RootObject into JSON bytes.
func (r RootObject) Build() ([]byte, error) {
	w := jwriter.Writer{}
	writer := jsoni.NewObjectWriter(&w)
	writer.Open()
	for _, field := range r {
		field.write(&writer)
	}
	writer.Close()
	return writer.BuildBytes()
}

// RootArray represents a json root array.
type RootArray []Value

// NewArray creates a new root array.
func NewArray(values ...Value) RootArray {
	return values
}

// Build encodes the RootArray into JSON bytes.
func (r RootArray) Build() ([]byte, error) {
	w := jwriter.Writer{}
	writer := jsoni.NewArrayWriter(&w)
	writer.Open()
	for _, value := range r {
		value.write(&writer)
	}
	writer.Close()
	return writer.BuildBytes()
}
