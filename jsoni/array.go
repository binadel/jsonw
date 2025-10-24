package jsoni

import "github.com/mailru/easyjson/jwriter"

// ArrayWriter builds a JSON array manually, supporting values of various types,
// including nested objects and arrays.
type ArrayWriter struct {
	writer     *jwriter.Writer
	needsComma bool
}

// NewArrayWriter creates a new ArrayWriter given an optional writer from its parent node.
func NewArrayWriter(writer *jwriter.Writer) *ArrayWriter {
	if writer == nil {
		writer = &jwriter.Writer{}
	}

	return &ArrayWriter{
		writer:     writer,
		needsComma: false,
	}
}

// Open starts the JSON array by writing '['.
func (w *ArrayWriter) Open() {
	w.writer.RawByte(openBracket)

	w.needsComma = false
}

// ObjectValue appends a new object to the array and returns its writer for further modifications.
func (w *ArrayWriter) ObjectValue() *ObjectWriter {
	if w.needsComma {
		w.writer.RawByte(comma)
	}

	w.needsComma = true

	return NewObjectWriter(w.writer)
}

// ArrayValue appends a new nested array and returns its writer for further modifications.
func (w *ArrayWriter) ArrayValue() *ArrayWriter {
	if w.needsComma {
		w.writer.RawByte(comma)
	}

	w.needsComma = true

	return NewArrayWriter(w.writer)
}

// StringValue appends a string value to the array.
func (w *ArrayWriter) StringValue(value string) {
	if w.needsComma {
		w.writer.RawByte(comma)
	}

	w.writer.String(value)

	w.needsComma = true
}

// NumberValue appends a number value to the array.
func (w *ArrayWriter) NumberValue(value string) {
	if w.needsComma {
		w.writer.RawByte(comma)
	}

	w.writer.RawString(value)

	w.needsComma = true
}

// IntegerValue appends an integer value to the array.
func (w *ArrayWriter) IntegerValue(value int64) {
	if w.needsComma {
		w.writer.RawByte(comma)
	}

	w.writer.Int64(value)

	w.needsComma = true
}

// FloatValue appends a float value to the array.
func (w *ArrayWriter) FloatValue(value float64) {
	if w.needsComma {
		w.writer.RawByte(comma)
	}

	w.writer.Float64(value)

	w.needsComma = true
}

// BooleanValue appends a boolean value to the array.
func (w *ArrayWriter) BooleanValue(value bool) {
	if w.needsComma {
		w.writer.RawByte(comma)
	}

	w.writer.Bool(value)

	w.needsComma = true
}

// NullValue appends a JSON null to the array.
func (w *ArrayWriter) NullValue() {
	if w.needsComma {
		w.writer.RawByte(comma)
	}

	w.writer.Raw(nullValue, nil)

	w.needsComma = true
}

// AnyValue appends a value of any type, automatically detecting its JSON representation.
func (w *ArrayWriter) AnyValue(value any) {
	if w.needsComma {
		w.writer.RawByte(comma)
	}

	writeAny(w.writer, value)

	w.needsComma = true
}

// Close finishes the JSON array by writing ']'.
func (w *ArrayWriter) Close() {
	w.writer.RawByte(closeBracket)

	w.needsComma = false
}

// BuildBytes returns the resulting JSON bytes.
func (w *ArrayWriter) BuildBytes() ([]byte, error) {
	return w.writer.BuildBytes()
}
