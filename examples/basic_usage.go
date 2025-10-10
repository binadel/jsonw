package main

import (
	"fmt"
	"log"

	"github.com/binadel/jsonw"
)

func main() {
	// Example 1: Simple object
	fmt.Println("=== Simple Object ===")
	obj := jsonw.NewObjectWriter(nil)
	obj.Open()
	obj.StringField("name", "John Doe")
	obj.IntegerField("age", 30)
	obj.FloatField("height", 5.9)
	obj.BooleanField("active", true)
	obj.Close()

	result, err := obj.BuildBytes()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Simple object: %s\n\n", string(result))

	// Example 2: Using AnyField for convenience
	fmt.Println("=== Using AnyField ===")
	obj2 := jsonw.NewObjectWriter(nil)
	obj2.Open()
	obj2.AnyField("name", "Alice")
	obj2.AnyField("age", 25)
	obj2.AnyField("score", 98.5)
	obj2.AnyField("graduated", true)
	obj2.Close()

	result2, err := obj2.BuildBytes()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("AnyField object: %s\n\n", string(result2))

	// Example 3: Simple array
	fmt.Println("=== Simple Array ===")
	arr := jsonw.NewArrayWriter(nil)
	arr.Open()
	arr.StringValue("apple")
	arr.StringValue("banana")
	arr.StringValue("cherry")
	arr.Close()

	result3, err := arr.BuildBytes()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Simple array: %s\n\n", string(result3))

	// Example 4: Using AnyValue for convenience
	fmt.Println("=== Using AnyValue ===")
	arr2 := jsonw.NewArrayWriter(nil)
	arr2.Open()
	arr2.AnyValue("hello")
	arr2.AnyValue(42)
	arr2.AnyValue(3.14)
	arr2.AnyValue(true)
	arr2.AnyValue(nil)
	arr2.Close()

	result4, err := arr2.BuildBytes()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("AnyValue array: %s\n\n", string(result4))

	// Example 5: Nested structures
	fmt.Println("=== Nested Structures ===")
	complexObj := jsonw.NewObjectWriter(nil)
	complexObj.Open()
	complexObj.StringField("name", "Bob")
	complexObj.IntegerField("id", 123)

	// Nested object
	address := complexObj.ObjectField("address")
	address.Open()
	address.StringField("street", "456 Oak Ave")
	address.StringField("city", "Boston")
	address.StringField("state", "MA")
	address.IntegerField("zip", 02101)
	address.Close()

	// Nested array
	hobbies := complexObj.ArrayField("hobbies")
	hobbies.Open()
	hobbies.StringValue("reading")
	hobbies.StringValue("swimming")
	hobbies.StringValue("cooking")
	hobbies.Close()

	// Array of objects
	skills := complexObj.ArrayField("skills")
	skills.Open()

	skill1 := skills.ObjectValue()
	skill1.Open()
	skill1.StringField("name", "Programming")
	skill1.StringField("level", "expert")
	languages := skill1.ArrayField("languages")
	languages.Open()
	languages.StringValue("Go")
	languages.StringValue("Python")
	languages.Close()
	skill1.Close()

	skill2 := skills.ObjectValue()
	skill2.Open()
	skill2.StringField("name", "Design")
	skill2.StringField("level", "intermediate")
	tools := skill2.ArrayField("tools")
	tools.Open()
	tools.StringValue("Figma")
	tools.StringValue("Sketch")
	tools.Close()
	skill2.Close()

	skills.Close()
	complexObj.Close()

	result5, err := complexObj.BuildBytes()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Complex nested structure: %s\n\n", string(result5))

	// Example 6: API Response pattern
	fmt.Println("=== API Response Pattern ===")
	apiResponse := jsonw.NewObjectWriter(nil)
	apiResponse.Open()
	apiResponse.StringField("status", "success")
	apiResponse.IntegerField("code", 200)
	apiResponse.FloatField("response_time", 0.123)

	// Data array
	data := apiResponse.ArrayField("data")
	data.Open()

	user1 := data.ObjectValue()
	user1.Open()
	user1.IntegerField("id", 1)
	user1.StringField("username", "alice")
	user1.StringField("email", "alice@example.com")
	user1.BooleanField("verified", true)
	user1.Close()

	user2 := data.ObjectValue()
	user2.Open()
	user2.IntegerField("id", 2)
	user2.StringField("username", "bob")
	user2.StringField("email", "bob@example.com")
	user2.BooleanField("verified", false)
	user2.Close()

	data.Close()

	// Pagination
	pagination := apiResponse.ObjectField("pagination")
	pagination.Open()
	pagination.IntegerField("page", 1)
	pagination.IntegerField("per_page", 10)
	pagination.IntegerField("total", 2)
	pagination.IntegerField("total_pages", 1)
	pagination.Close()

	apiResponse.Close()

	result6, err := apiResponse.BuildBytes()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("API Response: %s\n", string(result6))
}
