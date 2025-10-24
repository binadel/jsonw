package jsondi

import "github.com/binadel/jsonw/jsoni"

// Field represents an object field.
type Field interface {
	write(writer *jsoni.ObjectWriter)
}

// Value represents an array value.
type Value interface {
	write(writer *jsoni.ArrayWriter)
}
