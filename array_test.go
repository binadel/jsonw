package jsonw

import (
	"encoding/json"
	"testing"

	"github.com/mailru/easyjson/jwriter"
)

func TestArrayWriter_NewArrayWriter(t *testing.T) {
	tests := []struct {
		name   string
		writer *jwriter.Writer
	}{
		{
			name:   "with nil writer",
			writer: nil,
		},
		{
			name:   "with existing writer",
			writer: &jwriter.Writer{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			arr := NewArrayWriter(tt.writer)
			if arr == nil {
				t.Fatal("NewArrayWriter returned nil")
			}
			if arr.writer == nil {
				t.Fatal("ArrayWriter.writer is nil")
			}
			if arr.needsComma != false {
				t.Fatal("ArrayWriter.needsComma should be false initially")
			}
		})
	}
}

func TestArrayWriter_OpenClose(t *testing.T) {
	arr := NewArrayWriter(nil)
	arr.Open()
	arr.Close()

	result, err := arr.BuildBytes()
	if err != nil {
		t.Fatalf("BuildBytes failed: %v", err)
	}

	expected := "[]"
	if string(result) != expected {
		t.Errorf("Expected %q, got %q", expected, string(result))
	}
}

func TestArrayWriter_StringValue(t *testing.T) {
	tests := []struct {
		name     string
		values   []string
		expected []string
	}{
		{
			name:     "single string",
			values:   []string{"hello"},
			expected: []string{"hello"},
		},
		{
			name:     "multiple strings",
			values:   []string{"apple", "banana", "cherry"},
			expected: []string{"apple", "banana", "cherry"},
		},
		{
			name:     "empty strings",
			values:   []string{"", "not empty", ""},
			expected: []string{"", "not empty", ""},
		},
		{
			name:     "strings with quotes",
			values:   []string{`He said "Hello"`, `She said 'Hi'`},
			expected: []string{`He said "Hello"`, `She said 'Hi'`},
		},
		{
			name:     "strings with special chars",
			values:   []string{"line1\nline2", "tab\there", "backslash\\here"},
			expected: []string{"line1\nline2", "tab\there", "backslash\\here"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			arr := NewArrayWriter(nil)
			arr.Open()

			for _, value := range tt.values {
				arr.StringValue(value)
			}

			arr.Close()

			result, err := arr.BuildBytes()
			if err != nil {
				t.Fatalf("BuildBytes failed: %v", err)
			}

			// Parse and compare to ensure valid JSON
			var parsed []interface{}
			if err := json.Unmarshal(result, &parsed); err != nil {
				t.Fatalf("Generated JSON is invalid: %v, got: %s", err, string(result))
			}

			if len(parsed) != len(tt.expected) {
				t.Errorf("Expected %d values, got %d", len(tt.expected), len(parsed))
			}

			for i, expected := range tt.expected {
				if i >= len(parsed) {
					t.Errorf("Missing value at index %d", i)
					continue
				}
				if val, ok := parsed[i].(string); !ok || val != expected {
					t.Errorf("Value at index %d: expected %q, got %v", i, expected, parsed[i])
				}
			}
		})
	}
}

func TestArrayWriter_NumberValue(t *testing.T) {
	tests := []struct {
		name     string
		values   []string
		expected []interface{}
	}{
		{
			name:     "single number",
			values:   []string{"42"},
			expected: []interface{}{float64(42)},
		},
		{
			name:     "multiple numbers",
			values:   []string{"1", "2.5", "-3.14"},
			expected: []interface{}{float64(1), 2.5, -3.14},
		},
		{
			name:     "scientific notation",
			values:   []string{"1e5", "2.5e-3", "1.23e+4"},
			expected: []interface{}{float64(100000), 0.0025, float64(12300)},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			arr := NewArrayWriter(nil)
			arr.Open()

			for _, value := range tt.values {
				arr.NumberValue(value)
			}

			arr.Close()

			result, err := arr.BuildBytes()
			if err != nil {
				t.Fatalf("BuildBytes failed: %v", err)
			}

			// Parse and compare to ensure valid JSON
			var parsed []interface{}
			if err := json.Unmarshal(result, &parsed); err != nil {
				t.Fatalf("Generated JSON is invalid: %v, got: %s", err, string(result))
			}

			if len(parsed) != len(tt.expected) {
				t.Errorf("Expected %d values, got %d", len(tt.expected), len(parsed))
			}

			for i, expected := range tt.expected {
				if i >= len(parsed) {
					t.Errorf("Missing value at index %d", i)
					continue
				}
				if parsed[i] != expected {
					t.Errorf("Value at index %d: expected %v, got %v", i, expected, parsed[i])
				}
			}
		})
	}
}

func TestArrayWriter_IntegerValue(t *testing.T) {
	tests := []struct {
		name     string
		values   []int64
		expected []float64 // JSON numbers are float64
	}{
		{
			name:     "single integer",
			values:   []int64{42},
			expected: []float64{42},
		},
		{
			name:     "multiple integers",
			values:   []int64{1, -5, 0, 1000},
			expected: []float64{1, -5, 0, 1000},
		},
		{
			name:     "large integers",
			values:   []int64{9223372036854775807, -9223372036854775808},
			expected: []float64{9223372036854775807, -9223372036854775808},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			arr := NewArrayWriter(nil)
			arr.Open()

			for _, value := range tt.values {
				arr.IntegerValue(value)
			}

			arr.Close()

			result, err := arr.BuildBytes()
			if err != nil {
				t.Fatalf("BuildBytes failed: %v", err)
			}

			// Parse and compare to ensure valid JSON
			var parsed []interface{}
			if err := json.Unmarshal(result, &parsed); err != nil {
				t.Fatalf("Generated JSON is invalid: %v, got: %s", err, string(result))
			}

			if len(parsed) != len(tt.expected) {
				t.Errorf("Expected %d values, got %d", len(tt.expected), len(parsed))
			}

			for i, expected := range tt.expected {
				if i >= len(parsed) {
					t.Errorf("Missing value at index %d", i)
					continue
				}
				if val, ok := parsed[i].(float64); !ok || val != expected {
					t.Errorf("Value at index %d: expected %g, got %v", i, expected, parsed[i])
				}
			}
		})
	}
}

func TestArrayWriter_FloatValue(t *testing.T) {
	tests := []struct {
		name     string
		values   []float64
		expected []float64
	}{
		{
			name:     "single float",
			values:   []float64{3.14},
			expected: []float64{3.14},
		},
		{
			name:     "multiple floats",
			values:   []float64{1.5, -2.7, 0.0, 100.99},
			expected: []float64{1.5, -2.7, 0.0, 100.99},
		},
		{
			name:     "scientific notation",
			values:   []float64{1.23e-4, 5.67e+8},
			expected: []float64{0.000123, 567000000},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			arr := NewArrayWriter(nil)
			arr.Open()

			for _, value := range tt.values {
				arr.FloatValue(value)
			}

			arr.Close()

			result, err := arr.BuildBytes()
			if err != nil {
				t.Fatalf("BuildBytes failed: %v", err)
			}

			// Parse and compare to ensure valid JSON
			var parsed []interface{}
			if err := json.Unmarshal(result, &parsed); err != nil {
				t.Fatalf("Generated JSON is invalid: %v, got: %s", err, string(result))
			}

			if len(parsed) != len(tt.expected) {
				t.Errorf("Expected %d values, got %d", len(tt.expected), len(parsed))
			}

			for i, expected := range tt.expected {
				if i >= len(parsed) {
					t.Errorf("Missing value at index %d", i)
					continue
				}
				if val, ok := parsed[i].(float64); !ok || val != expected {
					t.Errorf("Value at index %d: expected %g, got %v", i, expected, parsed[i])
				}
			}
		})
	}
}

func TestArrayWriter_BooleanValue(t *testing.T) {
	tests := []struct {
		name     string
		values   []bool
		expected []bool
	}{
		{
			name:     "single boolean",
			values:   []bool{true},
			expected: []bool{true},
		},
		{
			name:     "multiple booleans",
			values:   []bool{true, false, true},
			expected: []bool{true, false, true},
		},
		{
			name:     "all true",
			values:   []bool{true, true, true},
			expected: []bool{true, true, true},
		},
		{
			name:     "all false",
			values:   []bool{false, false, false},
			expected: []bool{false, false, false},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			arr := NewArrayWriter(nil)
			arr.Open()

			for _, value := range tt.values {
				arr.BooleanValue(value)
			}

			arr.Close()

			result, err := arr.BuildBytes()
			if err != nil {
				t.Fatalf("BuildBytes failed: %v", err)
			}

			// Parse and compare to ensure valid JSON
			var parsed []interface{}
			if err := json.Unmarshal(result, &parsed); err != nil {
				t.Fatalf("Generated JSON is invalid: %v, got: %s", err, string(result))
			}

			if len(parsed) != len(tt.expected) {
				t.Errorf("Expected %d values, got %d", len(tt.expected), len(parsed))
			}

			for i, expected := range tt.expected {
				if i >= len(parsed) {
					t.Errorf("Missing value at index %d", i)
					continue
				}
				if val, ok := parsed[i].(bool); !ok || val != expected {
					t.Errorf("Value at index %d: expected %t, got %v", i, expected, parsed[i])
				}
			}
		})
	}
}

func TestArrayWriter_NullValue(t *testing.T) {
	arr := NewArrayWriter(nil)
	arr.Open()
	arr.NullValue()
	arr.NullValue()
	arr.NullValue()
	arr.Close()

	result, err := arr.BuildBytes()
	if err != nil {
		t.Fatalf("BuildBytes failed: %v", err)
	}

	// Parse and verify null values
	var parsed []interface{}
	if err := json.Unmarshal(result, &parsed); err != nil {
		t.Fatalf("Generated JSON is invalid: %v, got: %s", err, string(result))
	}

	if len(parsed) != 3 {
		t.Errorf("Expected 3 null values, got %d", len(parsed))
	}

	for i, val := range parsed {
		if val != nil {
			t.Errorf("Value at index %d: expected null, got %v", i, val)
		}
	}
}

func TestArrayWriter_AnyValue(t *testing.T) {
	tests := []struct {
		name     string
		values   []any
		expected []interface{}
	}{
		{
			name:     "mixed types",
			values:   []any{"hello", 42, 3.14, true, nil},
			expected: []interface{}{"hello", float64(42), 3.14, true, nil},
		},
		{
			name:     "integers",
			values:   []any{int(1), int8(2), int16(3), int32(4), int64(5)},
			expected: []interface{}{float64(1), float64(2), float64(3), float64(4), float64(5)},
		},
		{
			name:     "floats",
			values:   []any{float32(1.5), float64(2.7)},
			expected: []interface{}{float64(1.5), 2.7},
		},
		{
			name:     "unsigned integers",
			values:   []any{uint(10), uint8(20), uint16(30), uint32(40), uint64(50)},
			expected: []interface{}{float64(10), float64(20), float64(30), float64(40), float64(50)},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			arr := NewArrayWriter(nil)
			arr.Open()

			for _, value := range tt.values {
				arr.AnyValue(value)
			}

			arr.Close()

			result, err := arr.BuildBytes()
			if err != nil {
				t.Fatalf("BuildBytes failed: %v", err)
			}

			// Parse and compare to ensure valid JSON
			var parsed []interface{}
			if err := json.Unmarshal(result, &parsed); err != nil {
				t.Fatalf("Generated JSON is invalid: %v, got: %s", err, string(result))
			}

			if len(parsed) != len(tt.expected) {
				t.Errorf("Expected %d values, got %d", len(tt.expected), len(parsed))
			}

			for i, expected := range tt.expected {
				if i >= len(parsed) {
					t.Errorf("Missing value at index %d", i)
					continue
				}
				if parsed[i] != expected {
					t.Errorf("Value at index %d: expected %v (%T), got %v (%T)",
						i, expected, expected, parsed[i], parsed[i])
				}
			}
		})
	}
}

func TestArrayWriter_MixedTypes(t *testing.T) {
	arr := NewArrayWriter(nil)
	arr.Open()
	arr.StringValue("hello")
	arr.IntegerValue(42)
	arr.FloatValue(3.14)
	arr.BooleanValue(true)
	arr.NullValue()
	arr.Close()

	result, err := arr.BuildBytes()
	if err != nil {
		t.Fatalf("BuildBytes failed: %v", err)
	}

	// Parse and verify mixed types
	var parsed []interface{}
	if err := json.Unmarshal(result, &parsed); err != nil {
		t.Fatalf("Generated JSON is invalid: %v, got: %s", err, string(result))
	}

	expected := []interface{}{
		"hello",
		float64(42),
		3.14,
		true,
		nil,
	}

	if len(parsed) != len(expected) {
		t.Errorf("Expected %d values, got %d", len(expected), len(parsed))
	}

	for i, expectedValue := range expected {
		if i >= len(parsed) {
			t.Errorf("Missing value at index %d", i)
			continue
		}
		if parsed[i] != expectedValue {
			t.Errorf("Value at index %d: expected %v (%T), got %v (%T)",
				i, expectedValue, expectedValue, parsed[i], parsed[i])
		}
	}
}

func TestArrayWriter_NestedObject(t *testing.T) {
	arr := NewArrayWriter(nil)
	arr.Open()
	arr.StringValue("first")

	obj := arr.ObjectValue()
	obj.Open()
	obj.StringField("name", "John")
	obj.IntegerField("age", 30)
	obj.Close()

	arr.StringValue("last")
	arr.Close()

	result, err := arr.BuildBytes()
	if err != nil {
		t.Fatalf("BuildBytes failed: %v", err)
	}

	// Parse and verify nested object
	var parsed []interface{}
	if err := json.Unmarshal(result, &parsed); err != nil {
		t.Fatalf("Generated JSON is invalid: %v, got: %s", err, string(result))
	}

	if len(parsed) != 3 {
		t.Errorf("Expected 3 values, got %d", len(parsed))
	}

	// Check first string
	if val, ok := parsed[0].(string); !ok || val != "first" {
		t.Errorf("First value: expected \"first\", got %v", parsed[0])
	}

	// Check nested object
	objData, ok := parsed[1].(map[string]interface{})
	if !ok {
		t.Fatal("Second value is not an object")
	}

	if name, exists := objData["name"]; !exists || name != "John" {
		t.Error("Nested object name field not found or incorrect")
	}

	if age, exists := objData["age"]; !exists || age != float64(30) {
		t.Error("Nested object age field not found or incorrect")
	}

	// Check last string
	if val, ok := parsed[2].(string); !ok || val != "last" {
		t.Errorf("Last value: expected \"last\", got %v", parsed[2])
	}
}

func TestArrayWriter_NestedArray(t *testing.T) {
	arr := NewArrayWriter(nil)
	arr.Open()
	arr.StringValue("first")

	nested := arr.ArrayValue()
	nested.Open()
	nested.StringValue("a")
	nested.StringValue("b")
	nested.Close()

	arr.StringValue("last")
	arr.Close()

	result, err := arr.BuildBytes()
	if err != nil {
		t.Fatalf("BuildBytes failed: %v", err)
	}

	// Parse and verify nested array
	var parsed []interface{}
	if err := json.Unmarshal(result, &parsed); err != nil {
		t.Fatalf("Generated JSON is invalid: %v, got: %s", err, string(result))
	}

	if len(parsed) != 3 {
		t.Errorf("Expected 3 values, got %d", len(parsed))
	}

	// Check first string
	if val, ok := parsed[0].(string); !ok || val != "first" {
		t.Errorf("First value: expected \"first\", got %v", parsed[0])
	}

	// Check nested array
	nestedArray, ok := parsed[1].([]interface{})
	if !ok {
		t.Fatal("Second value is not an array")
	}

	expectedNested := []string{"a", "b"}
	if len(nestedArray) != len(expectedNested) {
		t.Errorf("Expected %d nested values, got %d", len(expectedNested), len(nestedArray))
	}

	for i, expected := range expectedNested {
		if i >= len(nestedArray) {
			t.Errorf("Missing nested value at index %d", i)
			continue
		}
		if val, ok := nestedArray[i].(string); !ok || val != expected {
			t.Errorf("Nested value at index %d: expected %q, got %v", i, expected, nestedArray[i])
		}
	}

	// Check last string
	if val, ok := parsed[2].(string); !ok || val != "last" {
		t.Errorf("Last value: expected \"last\", got %v", parsed[2])
	}
}

func TestArrayWriter_DeepNesting(t *testing.T) {
	arr := NewArrayWriter(nil)
	arr.Open()

	// Level 1: object
	obj1 := arr.ObjectValue()
	obj1.Open()
	obj1.StringField("type", "user")

	// Level 2: array inside object
	arr2 := obj1.ArrayField("tags")
	arr2.Open()
	arr2.StringValue("admin")
	arr2.StringValue("developer")

	// Level 3: object inside array
	obj3 := arr2.ObjectValue()
	obj3.Open()
	obj3.StringField("level", "senior")
	obj3.Close()

	arr2.Close()
	obj1.Close()

	arr.Close()

	result, err := arr.BuildBytes()
	if err != nil {
		t.Fatalf("BuildBytes failed: %v", err)
	}

	// Parse and verify deep nesting
	var parsed []interface{}
	if err := json.Unmarshal(result, &parsed); err != nil {
		t.Fatalf("Generated JSON is invalid: %v, got: %s", err, string(result))
	}

	if len(parsed) != 1 {
		t.Errorf("Expected 1 value, got %d", len(parsed))
	}

	obj1Data, ok := parsed[0].(map[string]interface{})
	if !ok {
		t.Fatal("First value is not an object")
	}

	if objType, exists := obj1Data["type"]; !exists || objType != "user" {
		t.Error("Object type field not found or incorrect")
	}

	tags, exists := obj1Data["tags"]
	if !exists {
		t.Fatal("Tags field not found")
	}

	tagsArray, ok := tags.([]interface{})
	if !ok {
		t.Fatal("Tags field is not an array")
	}

	if len(tagsArray) != 3 {
		t.Errorf("Expected 3 tags, got %d", len(tagsArray))
	}

	// Check string tags
	if tag1, ok := tagsArray[0].(string); !ok || tag1 != "admin" {
		t.Error("First tag is not \"admin\"")
	}

	if tag2, ok := tagsArray[1].(string); !ok || tag2 != "developer" {
		t.Error("Second tag is not \"developer\"")
	}

	// Check nested object tag
	tag3Obj, ok := tagsArray[2].(map[string]interface{})
	if !ok {
		t.Fatal("Third tag is not an object")
	}

	if level, exists := tag3Obj["level"]; !exists || level != "senior" {
		t.Error("Nested object level field not found or incorrect")
	}
}
