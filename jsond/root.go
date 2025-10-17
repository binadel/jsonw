package jsond

import (
	"github.com/binadel/jsonw/jsoni"
	"github.com/mailru/easyjson/jwriter"
)

type RootObject struct {
	fields []Field
}

func New(fields ...Field) RootObject {
	return RootObject{
		fields: fields,
	}
}

func (r RootObject) Write() ([]byte, error) {
	w := jwriter.Writer{}
	writer := jsoni.NewObjectWriter(&w)
	writer.Open()
	for _, field := range r.fields {
		field.Write(writer)
	}
	writer.Close()
	return writer.BuildBytes()
}

type RootArray struct {
	values []Value
}

func NewArray(values ...Value) RootArray {
	return RootArray{
		values: values,
	}
}

func (r RootArray) Write() ([]byte, error) {
	w := jwriter.Writer{}
	writer := jsoni.NewArrayWriter(&w)
	writer.Open()
	for _, value := range r.values {
		value.Write(writer)
	}
	writer.Close()
	return writer.BuildBytes()
}
