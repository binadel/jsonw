package test

import (
	"encoding/json"
	"testing"

	"github.com/binadel/jsonw/pkg"
)

func BenchmarkObjectWriter_SimpleObject(b *testing.B) {
	for i := 0; i < b.N; i++ {
		obj := jsonw.NewObjectWriter(nil)
		obj.Open()
		obj.StringField("name", "John")
		obj.IntegerField("age", 30)
		obj.BooleanField("active", true)
		obj.Close()

		_, err := obj.BuildBytes()
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkObjectWriter_AnyField(b *testing.B) {
	for i := 0; i < b.N; i++ {
		obj := jsonw.NewObjectWriter(nil)
		obj.Open()
		obj.AnyField("name", "John")
		obj.AnyField("age", 30)
		obj.AnyField("active", true)
		obj.Close()

		_, err := obj.BuildBytes()
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkObjectWriter_NestedStructure(b *testing.B) {
	for i := 0; i < b.N; i++ {
		obj := jsonw.NewObjectWriter(nil)
		obj.Open()
		obj.StringField("name", "John")

		address := obj.ObjectField("address")
		address.Open()
		address.StringField("street", "123 Main St")
		address.StringField("city", "NYC")
		address.Close()

		hobbies := obj.ArrayField("hobbies")
		hobbies.Open()
		hobbies.StringValue("reading")
		hobbies.StringValue("coding")
		hobbies.Close()

		obj.Close()

		_, err := obj.BuildBytes()
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkArrayWriter_SimpleArray(b *testing.B) {
	for i := 0; i < b.N; i++ {
		arr := jsonw.NewArrayWriter(nil)
		arr.Open()
		arr.StringValue("apple")
		arr.StringValue("banana")
		arr.StringValue("cherry")
		arr.Close()

		_, err := arr.BuildBytes()
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkArrayWriter_MixedTypes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		arr := jsonw.NewArrayWriter(nil)
		arr.Open()
		arr.StringValue("hello")
		arr.IntegerValue(42)
		arr.FloatValue(3.14)
		arr.BooleanValue(true)
		arr.NullValue()
		arr.Close()

		_, err := arr.BuildBytes()
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkArrayWriter_AnyValue(b *testing.B) {
	for i := 0; i < b.N; i++ {
		arr := jsonw.NewArrayWriter(nil)
		arr.Open()
		arr.AnyValue("hello")
		arr.AnyValue(42)
		arr.AnyValue(3.14)
		arr.AnyValue(true)
		arr.AnyValue(nil)
		arr.Close()

		_, err := arr.BuildBytes()
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkObjectWriter_LargeObject(b *testing.B) {
	for i := 0; i < b.N; i++ {
		obj := jsonw.NewObjectWriter(nil)
		obj.Open()

		// Add many fields
		for j := 0; j < 100; j++ {
			obj.StringField("field_"+string(rune(j)), "value_"+string(rune(j)))
			obj.IntegerField("int_"+string(rune(j)), int64(j))
			obj.BooleanField("bool_"+string(rune(j)), j%2 == 0)
		}

		obj.Close()

		_, err := obj.BuildBytes()
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkArrayWriter_LargeArray(b *testing.B) {
	for i := 0; i < b.N; i++ {
		arr := jsonw.NewArrayWriter(nil)
		arr.Open()

		// Add many values
		for j := 0; j < 1000; j++ {
			arr.StringValue("value_" + string(rune(j)))
		}

		arr.Close()

		_, err := arr.BuildBytes()
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkObjectWriter_DeepNesting(b *testing.B) {
	for i := 0; i < b.N; i++ {
		obj := jsonw.NewObjectWriter(nil)
		obj.Open()

		// Create 5 levels of nesting
		current := obj
		for level := 0; level < 5; level++ {
			nested := current.ObjectField("level_" + string(rune(level)))
			nested.Open()
			nested.StringField("data", "value")
			current = nested
		}

		// Close all levels
		for level := 0; level < 5; level++ {
			current.Close()
		}
		obj.Close()

		_, err := obj.BuildBytes()
		if err != nil {
			b.Fatal(err)
		}
	}
}

// Compare with standard library
func BenchmarkStandardJSON_Marshal(b *testing.B) {
	data := map[string]interface{}{
		"name":   "John",
		"age":    30,
		"active": true,
		"address": map[string]interface{}{
			"street": "123 Main St",
			"city":   "NYC",
		},
		"hobbies": []string{"reading", "coding"},
	}

	for i := 0; i < b.N; i++ {
		_, err := json.Marshal(data)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkJSONW_vs_Standard(b *testing.B) {
	b.Run("jsonw", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			obj := jsonw.NewObjectWriter(nil)
			obj.Open()
			obj.StringField("name", "John")
			obj.IntegerField("age", 30)
			obj.BooleanField("active", true)

			address := obj.ObjectField("address")
			address.Open()
			address.StringField("street", "123 Main St")
			address.StringField("city", "NYC")
			address.Close()

			hobbies := obj.ArrayField("hobbies")
			hobbies.Open()
			hobbies.StringValue("reading")
			hobbies.StringValue("coding")
			hobbies.Close()

			obj.Close()

			_, err := obj.BuildBytes()
			if err != nil {
				b.Fatal(err)
			}
		}
	})

	b.Run("standard", func(b *testing.B) {
		data := map[string]interface{}{
			"name":   "John",
			"age":    30,
			"active": true,
			"address": map[string]interface{}{
				"street": "123 Main St",
				"city":   "NYC",
			},
			"hobbies": []string{"reading", "coding"},
		}

		for i := 0; i < b.N; i++ {
			_, err := json.Marshal(data)
			if err != nil {
				b.Fatal(err)
			}
		}
	})
}
