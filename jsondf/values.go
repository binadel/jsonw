package jsondf

import "github.com/binadel/jsonw/jsoni"

// ObjectItem creates a nested object value.
func ObjectItem(fields ...Field) Value {
	return func(w *jsoni.ArrayWriter) {
		obj := w.ObjectValue()
		obj.Open()
		for _, field := range fields {
			field(&obj)
		}
		obj.Close()
	}
}

// ArrayItem creates a nested array value.
func ArrayItem(values ...Value) Value {
	return func(w *jsoni.ArrayWriter) {
		arr := w.ArrayValue()
		arr.Open()
		for _, value := range values {
			value(&arr)
		}
		arr.Close()
	}
}

// StringItem creates a string value.
func StringItem(value string) Value {
	return func(w *jsoni.ArrayWriter) {
		w.StringValue(value)
	}
}

// NumberItem creates a number value.
func NumberItem(value string) Value {
	return func(w *jsoni.ArrayWriter) {
		w.NumberValue(value)
	}
}

// IntegerItem creates an integer value.
func IntegerItem(value int64) Value {
	return func(w *jsoni.ArrayWriter) {
		w.IntegerValue(value)
	}
}

// FloatItem creates a float value.
func FloatItem(value float64) Value {
	return func(w *jsoni.ArrayWriter) {
		w.FloatValue(value)
	}
}

// BooleanItem creates a boolean value.
func BooleanItem(value bool) Value {
	return func(w *jsoni.ArrayWriter) {
		w.BooleanValue(value)
	}
}

// NullItem creates a null value.
func NullItem() Value {
	return func(w *jsoni.ArrayWriter) {
		w.NullValue()
	}
}

// AnyItem creates a dynamic value. Do not use it.
func AnyItem(value any) Value {
	return func(w *jsoni.ArrayWriter) {
		w.AnyValue(value)
	}
}
