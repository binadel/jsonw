package jsond

import "github.com/binadel/jsonw/jsoni"

type objectField struct {
	name   string
	fields []Field
}

func Object(name string, fields ...Field) Field {
	return objectField{
		name:   name,
		fields: fields,
	}
}

func (f objectField) Write(writer *jsoni.ObjectWriter) {
	obj := writer.ObjectField(f.name)
	obj.Open()
	for _, field := range f.fields {
		field.Write(obj)
	}
	obj.Close()
}

type arrayField struct {
	name   string
	values []Value
}

func Array(name string, values ...Value) Field {
	return arrayField{
		name:   name,
		values: values,
	}
}

func (f arrayField) Write(writer *jsoni.ObjectWriter) {
	arr := writer.ArrayField(f.name)
	arr.Open()
	for _, value := range f.values {
		value.Write(arr)
	}
	arr.Close()
}

type stringField struct {
	name, value string
}

func String(name, value string) Field {
	return stringField{name, value}
}

func (f stringField) Write(writer *jsoni.ObjectWriter) {
	writer.StringField(f.name, f.value)
}

type numberField struct {
	name, value string
}

func Number(name, value string) Field {
	return numberField{name, value}
}

func (f numberField) Write(writer *jsoni.ObjectWriter) {
	writer.NumberField(f.name, f.value)
}

type integerField struct {
	name  string
	value int64
}

func Integer(name string, value int64) Field {
	return integerField{name, value}
}

func (f integerField) Write(writer *jsoni.ObjectWriter) {
	writer.IntegerField(f.name, f.value)
}

type floatField struct {
	name  string
	value float64
}

func Float(name string, value float64) Field {
	return floatField{name, value}
}

func (f floatField) Write(writer *jsoni.ObjectWriter) {
	writer.FloatField(f.name, f.value)
}

type booleanField struct {
	name  string
	value bool
}

func Boolean(name string, value bool) Field {
	return booleanField{name, value}
}

func (f booleanField) Write(writer *jsoni.ObjectWriter) {
	writer.BooleanField(f.name, f.value)
}

type nullField struct {
	name string
}

func Null(name string) Field {
	return nullField{name}
}

func (f nullField) Write(writer *jsoni.ObjectWriter) {
	writer.NullField(f.name)
}

type anyField struct {
	name  string
	value any
}

func Any(name string, value any) Field {
	return anyField{name, value}
}

func (f anyField) Write(writer *jsoni.ObjectWriter) {
	writer.AnyField(f.name, f.value)
}
