package jsoni

import "github.com/mailru/easyjson/jwriter"

// ObjectWriter provides a JSON object builder used for manually writing objects.
// It supports writing fields of different types, including nested objects and arrays.
type ObjectWriter struct {
	writer     *jwriter.Writer
	needsComma bool
}

// NewObjectWriter creates a new ObjectWriter with an optional writer from its parent node.
func NewObjectWriter(writer *jwriter.Writer) *ObjectWriter {
	if writer == nil {
		writer = &jwriter.Writer{}
	}

	return &ObjectWriter{
		writer:     writer,
		needsComma: false,
	}
}

// Open writes the opening '{' for the object.
func (w *ObjectWriter) Open() {
	w.writer.RawByte(openBrace)

	w.needsComma = false
}

// ObjectField starts a new nested object field with the given name and returns an ObjectWriter.
func (w *ObjectWriter) ObjectField(name string) *ObjectWriter {
	if w.needsComma {
		w.writer.RawByte(comma)
	}

	w.writer.RawByte(quote)
	w.writer.RawString(name)
	w.writer.Raw(quoteColon, nil)

	w.needsComma = true

	return NewObjectWriter(w.writer)
}

// ArrayField starts a new nested array field with the given name and returns an ArrayWriter.
func (w *ObjectWriter) ArrayField(name string) *ArrayWriter {
	if w.needsComma {
		w.writer.RawByte(comma)
	}

	w.writer.RawByte(quote)
	w.writer.RawString(name)
	w.writer.Raw(quoteColon, nil)

	w.needsComma = true

	return NewArrayWriter(w.writer)
}

// StringField writes a string field to the object.
func (w *ObjectWriter) StringField(name, value string) {
	if w.needsComma {
		w.writer.RawByte(comma)
	}

	w.writer.RawByte(quote)
	w.writer.RawString(name)
	w.writer.Raw(quoteColon, nil)
	w.writer.String(value)

	w.needsComma = true
}

// NumberField writes a number field to the object.
func (w *ObjectWriter) NumberField(name, value string) {
	if w.needsComma {
		w.writer.RawByte(comma)
	}

	w.writer.RawByte(quote)
	w.writer.RawString(name)
	w.writer.Raw(quoteColon, nil)
	w.writer.RawString(value)

	w.needsComma = true
}

// IntegerField writes an integer field to the object.
func (w *ObjectWriter) IntegerField(name string, value int64) {
	if w.needsComma {
		w.writer.RawByte(comma)
	}

	w.writer.RawByte(quote)
	w.writer.RawString(name)
	w.writer.Raw(quoteColon, nil)
	w.writer.Int64(value)

	w.needsComma = true
}

// FloatField writes a float field to the object.
func (w *ObjectWriter) FloatField(name string, value float64) {
	if w.needsComma {
		w.writer.RawByte(comma)
	}

	w.writer.RawByte(quote)
	w.writer.RawString(name)
	w.writer.Raw(quoteColon, nil)
	w.writer.Float64(value)

	w.needsComma = true
}

// BooleanField writes a boolean field to the object.
func (w *ObjectWriter) BooleanField(name string, value bool) {
	if w.needsComma {
		w.writer.RawByte(comma)
	}

	w.writer.RawByte(quote)
	w.writer.RawString(name)
	w.writer.Raw(quoteColon, nil)
	w.writer.Bool(value)

	w.needsComma = true
}

// NullField writes a JSON null field to the object.
func (w *ObjectWriter) NullField(name string) {
	if w.needsComma {
		w.writer.RawByte(comma)
	}

	w.writer.RawByte(quote)
	w.writer.RawString(name)
	w.writer.Raw(quoteColonNull, nil)

	w.needsComma = true
}

// AnyField writes a field with any value type to the object.
// It automatically detects the type and uses the appropriate JSON representation.
func (w *ObjectWriter) AnyField(name string, value any) {
	if w.needsComma {
		w.writer.RawByte(comma)
	}

	w.writer.RawByte(quote)
	w.writer.RawString(name)
	w.writer.Raw(quoteColon, nil)
	writeAny(w.writer, value)

	w.needsComma = true
}

// Close writes the closing '}' for the object.
func (w *ObjectWriter) Close() {
	w.writer.RawByte(closeBrace)

	w.needsComma = false
}

// BuildBytes returns the JSON bytes written by the writer.
func (w *ObjectWriter) BuildBytes() ([]byte, error) {
	return w.writer.BuildBytes()
}
