package jsondf

import "github.com/binadel/jsonw/jsoni"

// Object creates a nested object field.
func Object(name string, fields ...Field) Field {
	return func(writer *jsoni.ObjectWriter) {
		obj := writer.ObjectField(name)
		obj.Open()
		for _, field := range fields {
			field(obj)
		}
		obj.Close()
	}
}

// Array creates a nested array field.
func Array(name string, values ...Value) Field {
	return func(writer *jsoni.ObjectWriter) {
		arr := writer.ArrayField(name)
		arr.Open()
		for _, value := range values {
			value(arr)
		}
		arr.Close()
	}
}

// String creates a string field.
func String(name, value string) Field {
	return func(writer *jsoni.ObjectWriter) {
		writer.StringField(name, value)
	}
}

// Number creates a number field.
func Number(name string, value string) Field {
	return func(writer *jsoni.ObjectWriter) {
		writer.NumberField(name, value)
	}
}

// Integer creates an integer field.
func Integer(name string, value int64) Field {
	return func(writer *jsoni.ObjectWriter) {
		writer.IntegerField(name, value)
	}
}

// Float creates a float field.
func Float(name string, value float64) Field {
	return func(writer *jsoni.ObjectWriter) {
		writer.FloatField(name, value)
	}
}

// Boolean creates a boolean field.
func Boolean(name string, value bool) Field {
	return func(writer *jsoni.ObjectWriter) {
		writer.BooleanField(name, value)
	}
}

// Null creates a null field.
func Null(name string) Field {
	return func(writer *jsoni.ObjectWriter) {
		writer.NullField(name)
	}
}

// Any creates a dynamic field. Do not use it.
func Any(name string, value any) Field {
	return func(writer *jsoni.ObjectWriter) {
		writer.AnyField(name, value)
	}
}
