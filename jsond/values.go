package jsond

import "github.com/binadel/jsonw/jsoni"

type objectValue struct {
	fields []Field
}

func ObjectItem(fields ...Field) Value {
	return objectValue{fields}
}

func (v objectValue) Write(writer *jsoni.ArrayWriter) {
	obj := writer.ObjectValue()
	obj.Open()
	for _, field := range v.fields {
		field.Write(obj)
	}
	obj.Close()
}

type arrayValue struct {
	values []Value
}

func ArrayItem(values ...Value) Value {
	return arrayValue{values}
}

func (v arrayValue) Write(writer *jsoni.ArrayWriter) {
	arr := writer.ArrayValue()
	arr.Open()
	for _, value := range v.values {
		value.Write(arr)
	}
	arr.Close()
}

type stringValue struct {
	value string
}

func StringItem(value string) Value {
	return stringValue{value}
}

func (v stringValue) Write(writer *jsoni.ArrayWriter) {
	writer.StringValue(v.value)
}

type numberValue struct {
	value string
}

func NumberItem(value string) Value {
	return numberValue{value}
}

func (v numberValue) Write(writer *jsoni.ArrayWriter) {
	writer.NumberValue(v.value)
}

type integerValue struct {
	value int64
}

func IntegerItem(value int64) Value {
	return integerValue{value}
}

func (v integerValue) Write(writer *jsoni.ArrayWriter) {
	writer.IntegerValue(v.value)
}

type floatValue struct {
	value float64
}

func FloatItem(value float64) Value {
	return floatValue{value}
}

func (v floatValue) Write(writer *jsoni.ArrayWriter) {
	writer.FloatValue(v.value)
}

type booleanValue struct {
	value bool
}

func BooleanItem(value bool) Value {
	return booleanValue{value}
}

func (v booleanValue) Write(writer *jsoni.ArrayWriter) {
	writer.BooleanValue(v.value)
}

type nullValue struct{}

func NullItem() Value {
	return nullValue{}
}

func (v nullValue) Write(writer *jsoni.ArrayWriter) {
	writer.NullValue()
}

type anyValue struct {
	value any
}

func AnyItem(value any) Value {
	return anyValue{value}
}

func (v anyValue) Write(writer *jsoni.ArrayWriter) {
	writer.AnyValue(v.value)
}
