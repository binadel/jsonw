package jsoni

import "github.com/mailru/easyjson/jwriter"

// ObjectWriter builds a JSON object manually, supporting fields of various types,
// including nested objects and arrays.
type ObjectWriter struct {
	writer     *jwriter.Writer
	needsComma bool
}

// NewObjectWriter creates a new ObjectWriter given an optional writer from its parent node.
func NewObjectWriter(writer *jwriter.Writer) ObjectWriter {
	if writer == nil {
		writer = &jwriter.Writer{}
	}

	return ObjectWriter{
		writer:     writer,
		needsComma: false,
	}
}

// Open starts the JSON object by writing '{'.
func (w *ObjectWriter) Open() {
	w.writer.RawByte(openBrace)

	w.needsComma = false
}

// ObjectField adds a nested object field and returns its writer for further modifications.
func (w *ObjectWriter) ObjectField(name string) ObjectWriter {
	if w.needsComma {
		w.writer.RawByte(comma)
	}

	w.writer.RawByte(quote)
	w.writer.RawString(name)
	w.writer.Raw(quoteColon, nil)

	w.needsComma = true

	return NewObjectWriter(w.writer)
}

// ArrayField adds a nested array field and returns its writer for further modifications.
func (w *ObjectWriter) ArrayField(name string) ArrayWriter {
	if w.needsComma {
		w.writer.RawByte(comma)
	}

	w.writer.RawByte(quote)
	w.writer.RawString(name)
	w.writer.Raw(quoteColon, nil)

	w.needsComma = true

	return NewArrayWriter(w.writer)
}

// StringField adds a string field to the object.
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

// NumberField adds a number field to the object.
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

// IntegerField adds an integer field to the object.
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

// FloatField adds a float field to the object.
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

// BooleanField adds a boolean field to the object.
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

// NullField adds a JSON null field to the object.
func (w *ObjectWriter) NullField(name string) {
	if w.needsComma {
		w.writer.RawByte(comma)
	}

	w.writer.RawByte(quote)
	w.writer.RawString(name)
	w.writer.Raw(quoteColonNull, nil)

	w.needsComma = true
}

// AnyField adds a field of any type, automatically detecting its JSON representation.
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

// Close finishes the JSON object by writing '}'.
func (w *ObjectWriter) Close() {
	w.writer.RawByte(closeBrace)

	w.needsComma = false
}

// BuildBytes returns the resulting JSON bytes.
func (w *ObjectWriter) BuildBytes() ([]byte, error) {
	return w.writer.BuildBytes()
}
