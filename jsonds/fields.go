package jsonds

// Object creates a nested object field.
func Object(name string, fields ...Field) Field {
	return Field{kind: kindObject, name: name, fields: append([]Field{}, fields...)}
}

// Array creates a nested array field.
func Array(name string, values ...Value) Field {
	return Field{kind: kindArray, name: name, values: append([]Value{}, values...)}
}

// String creates a string field.
func String(name, value string) Field {
	return Field{kind: kindString, name: name, s: value}
}

// Number creates a number field.
func Number(name, value string) Field {
	return Field{kind: kindNumber, name: name, n: value}
}

// Integer creates an integer field.
func Integer(name string, value int64) Field {
	return Field{kind: kindInteger, name: name, i: value}
}

// Float creates a float field.
func Float(name string, value float64) Field {
	return Field{kind: kindFloat, name: name, f: value}
}

// Boolean creates a boolean field.
func Boolean(name string, value bool) Field {
	return Field{kind: kindBoolean, name: name, b: value}
}

// Null creates a null field.
func Null(name string) Field {
	return Field{kind: kindNull, name: name}
}

// Any creates a dynamic field. Do not use it.
func Any(name string, value any) Field {
	return Field{kind: kindAny, name: name, a: value}
}
