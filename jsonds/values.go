package jsonds

// ObjectItem creates a nested object value.
func ObjectItem(fields ...Field) Value {
	return Value{kind: kindObject, fields: append([]Field{}, fields...)}
}

// ArrayItem creates a nested array value.
func ArrayItem(values ...Value) Value {
	return Value{kind: kindArray, values: append([]Value{}, values...)}
}

// StringItem creates a string value.
func StringItem(value string) Value {
	return Value{kind: kindString, s: value}
}

// NumberItem creates a number value.
func NumberItem(value string) Value {
	return Value{kind: kindNumber, n: value}
}

// IntegerItem creates an integer value.
func IntegerItem(value int64) Value {
	return Value{kind: kindInteger, i: value}
}

// FloatItem creates a float value.
func FloatItem(value float64) Value {
	return Value{kind: kindFloat, f: value}
}

// BooleanItem creates a boolean value.
func BooleanItem(value bool) Value {
	return Value{kind: kindBoolean, b: value}
}

// NullItem creates a null value.
func NullItem() Value {
	return Value{kind: kindNull}
}

// AnyItem creates a dynamic value. Do not use it.
func AnyItem(v any) Value {
	return Value{kind: kindAny, a: v}
}
