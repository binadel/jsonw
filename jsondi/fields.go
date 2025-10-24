package jsondi

import "github.com/binadel/jsonw/jsoni"

type objectField struct {
	name   string
	fields []Field
}

// Object creates a nested object field.
func Object(name string, fields ...Field) Field {
	return objectField{name, fields}
}

func (f objectField) write(writer *jsoni.ObjectWriter) {
	obj := writer.ObjectField(f.name)
	obj.Open()
	for _, field := range f.fields {
		field.write(obj)
	}
	obj.Close()
}

type arrayField struct {
	name   string
	values []Value
}

// Array creates a nested array field.
func Array(name string, values ...Value) Field {
	return arrayField{name, values}
}

func (f arrayField) write(writer *jsoni.ObjectWriter) {
	arr := writer.ArrayField(f.name)
	arr.Open()
	for _, value := range f.values {
		value.write(arr)
	}
	arr.Close()
}

type stringField struct {
	name, value string
}

// String creates a string field.
func String(name, value string) Field {
	return stringField{name, value}
}

func (f stringField) write(writer *jsoni.ObjectWriter) {
	writer.StringField(f.name, f.value)
}

type numberField struct {
	name, value string
}

// Number creates a number field.
func Number(name, value string) Field {
	return numberField{name, value}
}

func (f numberField) write(writer *jsoni.ObjectWriter) {
	writer.NumberField(f.name, f.value)
}

type integerField struct {
	name  string
	value int64
}

// Integer creates an integer field.
func Integer(name string, value int64) Field {
	return integerField{name, value}
}

func (f integerField) write(writer *jsoni.ObjectWriter) {
	writer.IntegerField(f.name, f.value)
}

type floatField struct {
	name  string
	value float64
}

// Float creates a float field.
func Float(name string, value float64) Field {
	return floatField{name, value}
}

func (f floatField) write(writer *jsoni.ObjectWriter) {
	writer.FloatField(f.name, f.value)
}

type booleanField struct {
	name  string
	value bool
}

// Boolean creates a boolean field.
func Boolean(name string, value bool) Field {
	return booleanField{name, value}
}

func (f booleanField) write(writer *jsoni.ObjectWriter) {
	writer.BooleanField(f.name, f.value)
}

type nullField struct {
	name string
}

// Null creates a null field.
func Null(name string) Field {
	return nullField{name}
}

func (f nullField) write(writer *jsoni.ObjectWriter) {
	writer.NullField(f.name)
}

type anyField struct {
	name  string
	value any
}

// Any creates a dynamic field. Do not use it.
func Any(name string, value any) Field {
	return anyField{name, value}
}

func (f anyField) write(writer *jsoni.ObjectWriter) {
	writer.AnyField(f.name, f.value)
}
