package jsonds

import (
	"github.com/binadel/jsonw/jsoni"
	"github.com/mailru/easyjson/jwriter"
)

// RootObject represents a json root object.
type RootObject []Field

// New creates a new root object.
func New(fields ...Field) RootObject {
	return append([]Field{}, fields...)
}

// Build encodes the RootObject into JSON bytes.
func (r RootObject) Build() ([]byte, error) {
	var jw jwriter.Writer
	ow := jsoni.NewObjectWriter(&jw)
	ow.Open()
	for _, f := range r {
		writeField(ow, &f)
	}
	ow.Close()
	return ow.BuildBytes()
}

// RootArray represents a json root array.
type RootArray []Value

// NewArray creates a new root array.
func NewArray(values ...Value) RootArray {
	return append([]Value{}, values...)
}

// Build encodes the RootArray into JSON bytes.
func (r RootArray) Build() ([]byte, error) {
	var jw jwriter.Writer
	aw := jsoni.NewArrayWriter(&jw)
	aw.Open()
	for _, v := range r {
		writeValue(aw, &v)
	}
	aw.Close()
	return aw.BuildBytes()
}

func writeField(w *jsoni.ObjectWriter, f *Field) {
	switch f.kind {
	case kindObject:
		obj := w.ObjectField(f.name)
		obj.Open()
		for i := range f.fields {
			writeField(obj, &f.fields[i])
		}
		obj.Close()
	case kindArray:
		arr := w.ArrayField(f.name)
		arr.Open()
		for i := range f.values {
			writeValue(arr, &f.values[i])
		}
		arr.Close()
	case kindString:
		w.StringField(f.name, f.s)
	case kindNumber:
		w.NumberField(f.name, f.n)
	case kindInteger:
		w.IntegerField(f.name, f.i)
	case kindFloat:
		w.FloatField(f.name, f.f)
	case kindBoolean:
		w.BooleanField(f.name, f.b)
	case kindNull:
		w.NullField(f.name)
	case kindAny:
		w.AnyField(f.name, f.a)
	default:
		panic("invalid field kind")
	}
}

func writeValue(w *jsoni.ArrayWriter, v *Value) {
	switch v.kind {
	case kindObject:
		obj := w.ObjectValue()
		obj.Open()
		for i := range v.fields {
			writeField(obj, &v.fields[i])
		}
		obj.Close()
	case kindArray:
		arr := w.ArrayValue()
		arr.Open()
		for i := range v.values {
			writeValue(arr, &v.values[i])
		}
		arr.Close()
	case kindString:
		w.StringValue(v.s)
	case kindNumber:
		w.NumberValue(v.n)
	case kindInteger:
		w.IntegerValue(v.i)
	case kindFloat:
		w.FloatValue(v.f)
	case kindBoolean:
		w.BooleanValue(v.b)
	case kindNull:
		w.NullValue()
	case kindAny:
		w.AnyValue(v.a)
	default:
		panic("invalid value kind")
	}
}
