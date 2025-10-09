package jsonw

import (
	"encoding/json"
	"testing"

	"github.com/mailru/easyjson/jwriter"
)

func TestObjectWriter_NewObjectWriter(t *testing.T) {
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
			obj := NewObjectWriter(tt.writer)
			if obj == nil {
				t.Fatal("NewObjectWriter returned nil")
			}
			if obj.writer == nil {
				t.Fatal("ObjectWriter.writer is nil")
			}
			if obj.needsComma != false {
				t.Fatal("ObjectWriter.needsComma should be false initially")
			}
		})
	}
}

func TestObjectWriter_OpenClose(t *testing.T) {
	obj := NewObjectWriter(nil)
	obj.Open()
	obj.Close()

	result, err := obj.BuildBytes()
	if err != nil {
		t.Fatalf("BuildBytes failed: %v", err)
	}

	expected := "{}"
	if string(result) != expected {
		t.Errorf("Expected %q, got %q", expected, string(result))
	}
}

func TestObjectWriter_StringField(t *testing.T) {
	tests := []struct {
		name     string
		field    string
		value    string
		expected string
	}{
		{
			name:     "simple string",
			field:    "name",
			value:    "John",
			expected: `{"name":"John"}`,
		},
		{
			name:     "empty string",
			field:    "description",
			value:    "",
			expected: `{"description":""}`,
		},
		{
			name:     "string with quotes",
			field:    "message",
			value:    `He said "Hello"`,
			expected: `{"message":"He said \"Hello\""}`,
		},
		{
			name:     "string with special chars",
			field:    "path",
			value:    "/path/to/file\nwith\tnewlines",
			expected: `{"path":"/path/to/file\nwith\tnewlines"}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			obj := NewObjectWriter(nil)
			obj.Open()
			obj.StringField(tt.field, tt.value)
			obj.Close()

			result, err := obj.BuildBytes()
			if err != nil {
				t.Fatalf("BuildBytes failed: %v", err)
			}

			// Parse and compare to ensure valid JSON
			var parsed map[string]interface{}
			if err := json.Unmarshal(result, &parsed); err != nil {
				t.Fatalf("Generated JSON is invalid: %v, got: %s", err, string(result))
			}

			if val, exists := parsed[tt.field]; !exists {
				t.Errorf("Field %q not found in result", tt.field)
			} else if val != tt.value {
				t.Errorf("Expected field value %q, got %q", tt.value, val)
			}
		})
	}
}

func TestObjectWriter_NumberField(t *testing.T) {
	tests := []struct {
		name     string
		field    string
		value    string
		expected interface{}
	}{
		{
			name:     "integer",
			field:    "count",
			value:    "42",
			expected: float64(42), // JSON numbers are float64
		},
		{
			name:     "float",
			field:    "price",
			value:    "19.99",
			expected: 19.99,
		},
		{
			name:     "negative number",
			field:    "balance",
			value:    "-100.50",
			expected: -100.5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			obj := NewObjectWriter(nil)
			obj.Open()
			obj.NumberField(tt.field, tt.value)
			obj.Close()

			result, err := obj.BuildBytes()
			if err != nil {
				t.Fatalf("BuildBytes failed: %v", err)
			}

			// Parse and compare to ensure valid JSON
			var parsed map[string]interface{}
			if err := json.Unmarshal(result, &parsed); err != nil {
				t.Fatalf("Generated JSON is invalid: %v, got: %s", err, string(result))
			}

			if val, exists := parsed[tt.field]; !exists {
				t.Errorf("Field %q not found in result", tt.field)
			} else if val != tt.expected {
				t.Errorf("Expected field value %v, got %v", tt.expected, val)
			}
		})
	}
}

func TestObjectWriter_IntegerField(t *testing.T) {
	tests := []struct {
		name     string
		field    string
		value    int64
		expected string
	}{
		{
			name:     "positive integer",
			field:    "age",
			value:    30,
			expected: `{"age":30}`,
		},
		{
			name:     "zero",
			field:    "count",
			value:    0,
			expected: `{"count":0}`,
		},
		{
			name:     "negative integer",
			field:    "offset",
			value:    -5,
			expected: `{"offset":-5}`,
		},
		{
			name:     "large integer",
			field:    "timestamp",
			value:    1640995200000,
			expected: `{"timestamp":1640995200000}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			obj := NewObjectWriter(nil)
			obj.Open()
			obj.IntegerField(tt.field, tt.value)
			obj.Close()

			result, err := obj.BuildBytes()
			if err != nil {
				t.Fatalf("BuildBytes failed: %v", err)
			}

			// Parse and compare to ensure valid JSON
			var parsed map[string]interface{}
			if err := json.Unmarshal(result, &parsed); err != nil {
				t.Fatalf("Generated JSON is invalid: %v, got: %s", err, string(result))
			}

			if val, exists := parsed[tt.field]; !exists {
				t.Errorf("Field %q not found in result", tt.field)
			} else if floatVal, ok := val.(float64); !ok {
				t.Errorf("Expected numeric value, got %T", val)
			} else if int64(floatVal) != tt.value {
				t.Errorf("Expected field value %d, got %d", tt.value, int64(floatVal))
			}
		})
	}
}

func TestObjectWriter_FloatField(t *testing.T) {
	tests := []struct {
		name     string
		field    string
		value    float64
		expected string
	}{
		{
			name:     "positive float",
			field:    "price",
			value:    19.99,
			expected: `{"price":19.99}`,
		},
		{
			name:     "zero float",
			field:    "rate",
			value:    0.0,
			expected: `{"rate":0}`,
		},
		{
			name:     "negative float",
			field:    "temperature",
			value:    -10.5,
			expected: `{"temperature":-10.5}`,
		},
		{
			name:     "scientific notation",
			field:    "distance",
			value:    1.23e-4,
			expected: `{"distance":0.000123}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			obj := NewObjectWriter(nil)
			obj.Open()
			obj.FloatField(tt.field, tt.value)
			obj.Close()

			result, err := obj.BuildBytes()
			if err != nil {
				t.Fatalf("BuildBytes failed: %v", err)
			}

			// Parse and compare to ensure valid JSON
			var parsed map[string]interface{}
			if err := json.Unmarshal(result, &parsed); err != nil {
				t.Fatalf("Generated JSON is invalid: %v, got: %s", err, string(result))
			}

			if val, exists := parsed[tt.field]; !exists {
				t.Errorf("Field %q not found in result", tt.field)
			} else if floatVal, ok := val.(float64); !ok {
				t.Errorf("Expected numeric value, got %T", val)
			} else if floatVal != tt.value {
				t.Errorf("Expected field value %g, got %g", tt.value, floatVal)
			}
		})
	}
}

func TestObjectWriter_BooleanField(t *testing.T) {
	tests := []struct {
		name     string
		field    string
		value    bool
		expected string
	}{
		{
			name:     "true",
			field:    "active",
			value:    true,
			expected: `{"active":true}`,
		},
		{
			name:     "false",
			field:    "enabled",
			value:    false,
			expected: `{"enabled":false}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			obj := NewObjectWriter(nil)
			obj.Open()
			obj.BooleanField(tt.field, tt.value)
			obj.Close()

			result, err := obj.BuildBytes()
			if err != nil {
				t.Fatalf("BuildBytes failed: %v", err)
			}

			// Parse and compare to ensure valid JSON
			var parsed map[string]interface{}
			if err := json.Unmarshal(result, &parsed); err != nil {
				t.Fatalf("Generated JSON is invalid: %v, got: %s", err, string(result))
			}

			if val, exists := parsed[tt.field]; !exists {
				t.Errorf("Field %q not found in result", tt.field)
			} else if boolVal, ok := val.(bool); !ok {
				t.Errorf("Expected boolean value, got %T", val)
			} else if boolVal != tt.value {
				t.Errorf("Expected field value %t, got %t", tt.value, boolVal)
			}
		})
	}
}

func TestObjectWriter_NullField(t *testing.T) {
	obj := NewObjectWriter(nil)
	obj.Open()
	obj.NullField("data")
	obj.Close()

	result, err := obj.BuildBytes()
	if err != nil {
		t.Fatalf("BuildBytes failed: %v", err)
	}

	// Parse and compare to ensure valid JSON
	var parsed map[string]interface{}
	if err := json.Unmarshal(result, &parsed); err != nil {
		t.Fatalf("Generated JSON is invalid: %v, got: %s", err, string(result))
	}

	if val, exists := parsed["data"]; !exists {
		t.Error("Field 'data' not found in result")
	} else if val != nil {
		t.Errorf("Expected null value, got %v", val)
	}
}

func TestObjectWriter_AnyField(t *testing.T) {
	tests := []struct {
		name     string
		field    string
		value    any
		expected interface{}
	}{
		{
			name:     "string",
			field:    "name",
			value:    "Alice",
			expected: "Alice",
		},
		{
			name:     "int",
			field:    "age",
			value:    25,
			expected: float64(25), // JSON numbers are float64
		},
		{
			name:     "int64",
			field:    "id",
			value:    int64(12345),
			expected: float64(12345),
		},
		{
			name:     "float64",
			field:    "score",
			value:    98.5,
			expected: 98.5,
		},
		{
			name:     "bool",
			field:    "active",
			value:    true,
			expected: true,
		},
		{
			name:     "nil",
			field:    "data",
			value:    nil,
			expected: nil,
		},
		{
			name:     "uint",
			field:    "count",
			value:    uint(100),
			expected: float64(100),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			obj := NewObjectWriter(nil)
			obj.Open()
			obj.AnyField(tt.field, tt.value)
			obj.Close()

			result, err := obj.BuildBytes()
			if err != nil {
				t.Fatalf("BuildBytes failed: %v", err)
			}

			// Parse and compare to ensure valid JSON
			var parsed map[string]interface{}
			if err := json.Unmarshal(result, &parsed); err != nil {
				t.Fatalf("Generated JSON is invalid: %v, got: %s", err, string(result))
			}

			if val, exists := parsed[tt.field]; !exists {
				t.Errorf("Field %q not found in result", tt.field)
			} else if val != tt.expected {
				t.Errorf("Expected field value %v (%T), got %v (%T)", tt.expected, tt.expected, val, val)
			}
		})
	}
}

func TestObjectWriter_MultipleFields(t *testing.T) {
	obj := NewObjectWriter(nil)
	obj.Open()
	obj.StringField("name", "John")
	obj.IntegerField("age", 30)
	obj.FloatField("height", 5.9)
	obj.BooleanField("active", true)
	obj.NullField("data")
	obj.Close()

	result, err := obj.BuildBytes()
	if err != nil {
		t.Fatalf("BuildBytes failed: %v", err)
	}

	// Parse and verify all fields
	var parsed map[string]interface{}
	if err := json.Unmarshal(result, &parsed); err != nil {
		t.Fatalf("Generated JSON is invalid: %v, got: %s", err, string(result))
	}

	expected := map[string]interface{}{
		"name":   "John",
		"age":    float64(30),
		"height": 5.9,
		"active": true,
		"data":   nil,
	}

	if len(parsed) != len(expected) {
		t.Errorf("Expected %d fields, got %d", len(expected), len(parsed))
	}

	for key, expectedValue := range expected {
		if val, exists := parsed[key]; !exists {
			t.Errorf("Field %q not found", key)
		} else if val != expectedValue {
			t.Errorf("Field %q: expected %v, got %v", key, expectedValue, val)
		}
	}
}

func TestObjectWriter_NestedObject(t *testing.T) {
	obj := NewObjectWriter(nil)
	obj.Open()
	obj.StringField("name", "John")

	nested := obj.ObjectField("address")
	nested.Open()
	nested.StringField("street", "123 Main St")
	nested.StringField("city", "NYC")
	nested.Close()

	obj.Close()

	result, err := obj.BuildBytes()
	if err != nil {
		t.Fatalf("BuildBytes failed: %v", err)
	}

	// Parse and verify nested structure
	var parsed map[string]interface{}
	if err := json.Unmarshal(result, &parsed); err != nil {
		t.Fatalf("Generated JSON is invalid: %v, got: %s", err, string(result))
	}

	if name, exists := parsed["name"]; !exists || name != "John" {
		t.Error("Root name field not found or incorrect")
	}

	address, exists := parsed["address"]
	if !exists {
		t.Fatal("Address field not found")
	}

	addressMap, ok := address.(map[string]interface{})
	if !ok {
		t.Fatal("Address field is not an object")
	}

	if street, exists := addressMap["street"]; !exists || street != "123 Main St" {
		t.Error("Street field not found or incorrect")
	}

	if city, exists := addressMap["city"]; !exists || city != "NYC" {
		t.Error("City field not found or incorrect")
	}
}

func TestObjectWriter_NestedArray(t *testing.T) {
	obj := NewObjectWriter(nil)
	obj.Open()
	obj.StringField("name", "John")

	array := obj.ArrayField("hobbies")
	array.Open()
	array.StringValue("reading")
	array.StringValue("coding")
	array.Close()

	obj.Close()

	result, err := obj.BuildBytes()
	if err != nil {
		t.Fatalf("BuildBytes failed: %v", err)
	}

	// Parse and verify nested array
	var parsed map[string]interface{}
	if err := json.Unmarshal(result, &parsed); err != nil {
		t.Fatalf("Generated JSON is invalid: %v, got: %s", err, string(result))
	}

	if name, exists := parsed["name"]; !exists || name != "John" {
		t.Error("Root name field not found or incorrect")
	}

	hobbies, exists := parsed["hobbies"]
	if !exists {
		t.Fatal("Hobbies field not found")
	}

	hobbiesArray, ok := hobbies.([]interface{})
	if !ok {
		t.Fatal("Hobbies field is not an array")
	}

	expectedHobbies := []string{"reading", "coding"}
	if len(hobbiesArray) != len(expectedHobbies) {
		t.Errorf("Expected %d hobbies, got %d", len(expectedHobbies), len(hobbiesArray))
	}

	for i, expected := range expectedHobbies {
		if i >= len(hobbiesArray) {
			t.Errorf("Missing hobby at index %d", i)
			continue
		}
		if hobby, ok := hobbiesArray[i].(string); !ok || hobby != expected {
			t.Errorf("Hobby at index %d: expected %q, got %v", i, expected, hobbiesArray[i])
		}
	}
}

func TestObjectWriter_EmptyFields(t *testing.T) {
	obj := NewObjectWriter(nil)
	obj.Open()
	obj.StringField("", "empty key")
	obj.StringField("empty value", "")
	obj.Close()

	result, err := obj.BuildBytes()
	if err != nil {
		t.Fatalf("BuildBytes failed: %v", err)
	}

	// Should still be valid JSON
	var parsed map[string]interface{}
	if err := json.Unmarshal(result, &parsed); err != nil {
		t.Fatalf("Generated JSON is invalid: %v, got: %s", err, string(result))
	}

	// Verify empty key and empty value are handled
	if val, exists := parsed[""]; !exists || val != "empty key" {
		t.Error("Empty key field not handled correctly")
	}

	if val, exists := parsed["empty value"]; !exists || val != "" {
		t.Error("Empty value field not handled correctly")
	}
}
