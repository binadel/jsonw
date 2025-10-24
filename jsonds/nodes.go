package jsonds

// NodeKind indicates the concrete kind of field or value.
type NodeKind uint8

const (
	kindObject NodeKind = iota
	kindArray
	kindString
	kindNumber
	kindInteger
	kindFloat
	kindBoolean
	kindNull
	kindAny
)

// Field represents an object field.
type Field struct {
	kind   NodeKind
	name   string
	fields []Field // for object
	values []Value // for array
	s      string  // for string
	n      string  // for number
	i      int64   // for integer
	f      float64 // for float
	b      bool    // for bool
	a      any     // for any
}

// Value represents an array value.
type Value struct {
	kind   NodeKind
	fields []Field // for object
	values []Value // for array
	s      string  // for string
	n      string  // for number
	i      int64   // for integer
	f      float64 // for float
	b      bool    // for bool
	a      any     // for any
}
