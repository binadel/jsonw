package jsondf

import "github.com/binadel/jsonw/jsoni"

// Field represents an object field.
type Field func(writer *jsoni.ObjectWriter)

// Value represents an array value.
type Value func(writer *jsoni.ArrayWriter)
