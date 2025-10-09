package jsonw

import (
	"encoding/json"

	"github.com/mailru/easyjson/jwriter"
)

func writeAny(writer *jwriter.Writer, value any) {
	switch v := value.(type) {
	case string:
		writer.String(v)
	case int:
		writer.Int(v)
	case int8:
		writer.Int8(v)
	case int16:
		writer.Int16(v)
	case int32:
		writer.Int32(v)
	case int64:
		writer.Int64(v)
	case uint:
		writer.Uint(v)
	case uint8:
		writer.Uint8(v)
	case uint16:
		writer.Uint16(v)
	case uint32:
		writer.Uint32(v)
	case uint64:
		writer.Uint64(v)
	case float32:
		writer.Float32(v)
	case float64:
		writer.Float64(v)
	case bool:
		writer.Bool(v)
	case nil:
		writer.Raw(nullValue, nil)
	default:
		writer.Raw(json.Marshal(value))
	}
}
