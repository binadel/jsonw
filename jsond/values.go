package jsond

import "github.com/binadel/jsonw/jsoni"

type objectValue struct {
	fields []Field
}

// ObjectItem represents an array item of type object.
func ObjectItem(fields ...Field) Value {
	return objectValue{fields}
}

func (v objectValue) write(writer *jsoni.ArrayWriter) {
	obj := writer.ObjectValue()
	obj.Open()
	for _, field := range v.fields {
		field.write(obj)
	}
	obj.Close()
}

type arrayValue struct {
	values []Value
}

// ArrayItem represents an array item of type array.
func ArrayItem(values ...Value) Value {
	return arrayValue{values}
}

func (v arrayValue) write(writer *jsoni.ArrayWriter) {
	arr := writer.ArrayValue()
	arr.Open()
	for _, value := range v.values {
		value.write(arr)
	}
	arr.Close()
}

type stringValue struct {
	value string
}

// StringItem represents an array item of type string.
func StringItem(value string) Value {
	return stringValue{value}
}

func (v stringValue) write(writer *jsoni.ArrayWriter) {
	writer.StringValue(v.value)
}

type numberValue struct {
	value string
}

// NumberItem represents an array item of type number.
func NumberItem(value string) Value {
	return numberValue{value}
}

func (v numberValue) write(writer *jsoni.ArrayWriter) {
	writer.NumberValue(v.value)
}

type integerValue struct {
	value int64
}

// IntegerItem represents an array item of type integer.
func IntegerItem(value int64) Value {
	return integerValue{value}
}

func (v integerValue) write(writer *jsoni.ArrayWriter) {
	writer.IntegerValue(v.value)
}

type floatValue struct {
	value float64
}

// FloatItem represents an array item of type float.
func FloatItem(value float64) Value {
	return floatValue{value}
}

func (v floatValue) write(writer *jsoni.ArrayWriter) {
	writer.FloatValue(v.value)
}

type booleanValue struct {
	value bool
}

// BooleanItem represents an array item of type boolean.
func BooleanItem(value bool) Value {
	return booleanValue{value}
}

func (v booleanValue) write(writer *jsoni.ArrayWriter) {
	writer.BooleanValue(v.value)
}

type nullValue struct{}

// NullItem represents an array item of type null.
func NullItem() Value {
	return nullValue{}
}

func (v nullValue) write(writer *jsoni.ArrayWriter) {
	writer.NullValue()
}

type anyValue struct {
	value any
}

// AnyItem represents an array item of type any.
func AnyItem(value any) Value {
	return anyValue{value}
}

func (v anyValue) write(writer *jsoni.ArrayWriter) {
	writer.AnyValue(v.value)
}
