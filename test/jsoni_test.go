package test

import (
	"encoding/json"
	"testing"

	"github.com/binadel/jsonw/jsoni"
)

func TestComplexJSONStructure(t *testing.T) {
	// Create a complex JSON structure with nested objects and arrays
	obj := jsoni.NewObjectWriter(nil)
	obj.Open()

	// Root level fields
	obj.StringField("name", "John Doe")
	obj.IntegerField("age", 30)
	obj.FloatField("height", 5.9)
	obj.BooleanField("active", true)

	// Nested object: address
	address := obj.ObjectField("address")
	address.Open()
	address.StringField("street", "123 Main St")
	address.StringField("city", "New York")
	address.StringField("state", "NY")
	address.IntegerField("zip", 10001)
	address.Close()

	// Nested array: hobbies
	hobbies := obj.ArrayField("hobbies")
	hobbies.Open()
	hobbies.StringValue("reading")
	hobbies.StringValue("coding")
	hobbies.StringValue("hiking")
	hobbies.Close()

	// Nested array: skills with objects
	skills := obj.ArrayField("skills")
	skills.Open()

	// Skill 1: Programming
	skill1 := skills.ObjectValue()
	skill1.Open()
	skill1.StringField("name", "Programming")
	skill1.StringField("level", "expert")
	languages := skill1.ArrayField("languages")
	languages.Open()
	languages.StringValue("Go")
	languages.StringValue("Python")
	languages.StringValue("JavaScript")
	languages.Close()
	skill1.Close()

	// Skill 2: Design
	skill2 := skills.ObjectValue()
	skill2.Open()
	skill2.StringField("name", "Design")
	skill2.StringField("level", "intermediate")
	tools := skill2.ArrayField("tools")
	tools.Open()
	tools.StringValue("Figma")
	tools.StringValue("Photoshop")
	tools.Close()
	skill2.Close()

	skills.Close()

	// Nested object: preferences
	preferences := obj.ObjectField("preferences")
	preferences.Open()
	preferences.BooleanField("notifications", true)
	preferences.StringField("theme", "dark")

	// Nested array in preferences
	preferredLanguages := preferences.ArrayField("languages")
	preferredLanguages.Open()
	preferredLanguages.StringValue("en")
	preferredLanguages.StringValue("es")
	preferredLanguages.Close()

	preferences.Close()

	obj.Close()

	result, err := obj.BuildBytes()
	if err != nil {
		t.Fatalf("BuildBytes failed: %v", err)
	}

	// Parse and validate the complex structure
	var parsed map[string]interface{}
	if err := json.Unmarshal(result, &parsed); err != nil {
		t.Fatalf("Generated JSON is invalid: %v, got: %s", err, string(result))
	}

	// Validate root level fields
	if name, exists := parsed["name"]; !exists || name != "John Doe" {
		t.Error("Name field validation failed")
	}
	if age, exists := parsed["age"]; !exists || age != float64(30) {
		t.Error("Age field validation failed")
	}
	if height, exists := parsed["height"]; !exists || height != 5.9 {
		t.Error("Height field validation failed")
	}
	if active, exists := parsed["active"]; !exists || active != true {
		t.Error("Active field validation failed")
	}

	// Validate address object
	addressData, ok := parsed["address"].(map[string]interface{})
	if !ok {
		t.Fatal("Address field is not an object")
	}
	if street, exists := addressData["street"]; !exists || street != "123 Main St" {
		t.Error("Address street validation failed")
	}
	if city, exists := addressData["city"]; !exists || city != "New York" {
		t.Error("Address city validation failed")
	}
	if state, exists := addressData["state"]; !exists || state != "NY" {
		t.Error("Address state validation failed")
	}
	if zip, exists := addressData["zip"]; !exists || zip != float64(10001) {
		t.Error("Address zip validation failed")
	}

	// Validate hobbies array
	hobbiesData, ok := parsed["hobbies"].([]interface{})
	if !ok {
		t.Fatal("Hobbies field is not an array")
	}
	expectedHobbies := []string{"reading", "coding", "hiking"}
	if len(hobbiesData) != len(expectedHobbies) {
		t.Errorf("Expected %d hobbies, got %d", len(expectedHobbies), len(hobbiesData))
	}
	for i, expected := range expectedHobbies {
		if i >= len(hobbiesData) {
			t.Errorf("Missing hobby at index %d", i)
			continue
		}
		if hobby, ok := hobbiesData[i].(string); !ok || hobby != expected {
			t.Errorf("Hobby at index %d: expected %q, got %v", i, expected, hobbiesData[i])
		}
	}

	// Validate skills array with objects
	skillsData, ok := parsed["skills"].([]interface{})
	if !ok {
		t.Fatal("Skills field is not an array")
	}
	if len(skillsData) != 2 {
		t.Errorf("Expected 2 skills, got %d", len(skillsData))
	}

	// Validate first skill
	skill1Data, ok := skillsData[0].(map[string]interface{})
	if !ok {
		t.Fatal("First skill is not an object")
	}
	if name, exists := skill1Data["name"]; !exists || name != "Programming" {
		t.Error("First skill name validation failed")
	}
	if level, exists := skill1Data["level"]; !exists || level != "expert" {
		t.Error("First skill level validation failed")
	}

	languagesArray, ok := skill1Data["languages"].([]interface{})
	if !ok {
		t.Fatal("First skill languages field is not an array")
	}
	expectedLanguages := []string{"Go", "Python", "JavaScript"}
	if len(languagesArray) != len(expectedLanguages) {
		t.Errorf("Expected %d languages, got %d", len(expectedLanguages), len(languagesArray))
	}
	for i, expected := range expectedLanguages {
		if i >= len(languagesArray) {
			t.Errorf("Missing language at index %d", i)
			continue
		}
		if lang, ok := languagesArray[i].(string); !ok || lang != expected {
			t.Errorf("Language at index %d: expected %q, got %v", i, expected, languagesArray[i])
		}
	}

	// Validate preferences
	preferencesData, ok := parsed["preferences"].(map[string]interface{})
	if !ok {
		t.Fatal("Preferences field is not an object")
	}
	if notifications, exists := preferencesData["notifications"]; !exists || notifications != true {
		t.Error("Preferences notifications validation failed")
	}
	if theme, exists := preferencesData["theme"]; !exists || theme != "dark" {
		t.Error("Preferences theme validation failed")
	}

	preferredLangs, ok := preferencesData["languages"].([]interface{})
	if !ok {
		t.Fatal("Preferences languages field is not an array")
	}
	expectedPrefLangs := []string{"en", "es"}
	if len(preferredLangs) != len(expectedPrefLangs) {
		t.Errorf("Expected %d preferred languages, got %d", len(expectedPrefLangs), len(preferredLangs))
	}
	for i, expected := range expectedPrefLangs {
		if i >= len(preferredLangs) {
			t.Errorf("Missing preferred language at index %d", i)
			continue
		}
		if lang, ok := preferredLangs[i].(string); !ok || lang != expected {
			t.Errorf("Preferred language at index %d: expected %q, got %v", i, expected, preferredLangs[i])
		}
	}
}

func TestAPIResponseStructure(t *testing.T) {
	// Create a typical API response structure
	obj := jsoni.NewObjectWriter(nil)
	obj.Open()

	// Response metadata
	obj.StringField("status", "success")
	obj.IntegerField("code", 200)
	obj.FloatField("response_time", 0.123)

	// Data array
	data := obj.ArrayField("data")
	data.Open()

	// User 1
	user1 := data.ObjectValue()
	user1.Open()
	user1.IntegerField("id", 1)
	user1.StringField("username", "alice")
	user1.StringField("email", "alice@example.com")
	user1.BooleanField("verified", true)
	user1.Close()

	// User 2
	user2 := data.ObjectValue()
	user2.Open()
	user2.IntegerField("id", 2)
	user2.StringField("username", "bob")
	user2.StringField("email", "bob@example.com")
	user2.BooleanField("verified", false)
	user2.Close()

	data.Close()

	// Pagination info
	pagination := obj.ObjectField("pagination")
	pagination.Open()
	pagination.IntegerField("page", 1)
	pagination.IntegerField("per_page", 10)
	pagination.IntegerField("total", 2)
	pagination.IntegerField("total_pages", 1)
	pagination.Close()

	obj.Close()

	result, err := obj.BuildBytes()
	if err != nil {
		t.Fatalf("BuildBytes failed: %v", err)
	}

	// Parse and validate
	var parsed map[string]interface{}
	if err := json.Unmarshal(result, &parsed); err != nil {
		t.Fatalf("Generated JSON is invalid: %v, got: %s", err, string(result))
	}

	// Validate response metadata
	if status, exists := parsed["status"]; !exists || status != "success" {
		t.Error("Status field validation failed")
	}
	if code, exists := parsed["code"]; !exists || code != float64(200) {
		t.Error("Code field validation failed")
	}
	if responseTime, exists := parsed["response_time"]; !exists || responseTime != 0.123 {
		t.Error("Response time field validation failed")
	}

	// Validate data array
	dataArray, ok := parsed["data"].([]interface{})
	if !ok {
		t.Fatal("Data field is not an array")
	}
	if len(dataArray) != 2 {
		t.Errorf("Expected 2 users, got %d", len(dataArray))
	}

	// Validate first user
	user1Data, ok := dataArray[0].(map[string]interface{})
	if !ok {
		t.Fatal("First user is not an object")
	}
	if id, exists := user1Data["id"]; !exists || id != float64(1) {
		t.Error("First user ID validation failed")
	}
	if username, exists := user1Data["username"]; !exists || username != "alice" {
		t.Error("First user username validation failed")
	}
	if email, exists := user1Data["email"]; !exists || email != "alice@example.com" {
		t.Error("First user email validation failed")
	}
	if verified, exists := user1Data["verified"]; !exists || verified != true {
		t.Error("First user verified validation failed")
	}

	// Validate pagination
	paginationData, ok := parsed["pagination"].(map[string]interface{})
	if !ok {
		t.Fatal("Pagination field is not an object")
	}
	if page, exists := paginationData["page"]; !exists || page != float64(1) {
		t.Error("Pagination page validation failed")
	}
	if perPage, exists := paginationData["per_page"]; !exists || perPage != float64(10) {
		t.Error("Pagination per_page validation failed")
	}
	if total, exists := paginationData["total"]; !exists || total != float64(2) {
		t.Error("Pagination total validation failed")
	}
	if totalPages, exists := paginationData["total_pages"]; !exists || totalPages != float64(1) {
		t.Error("Pagination total_pages validation failed")
	}
}

func TestConfigurationStructure(t *testing.T) {
	// Create a configuration-like structure
	obj := jsoni.NewObjectWriter(nil)
	obj.Open()

	// Database config
	db := obj.ObjectField("database")
	db.Open()
	db.StringField("host", "localhost")
	db.IntegerField("port", 5432)
	db.StringField("name", "myapp")
	db.StringField("user", "admin")
	db.BooleanField("ssl", true)

	// Connection pool settings
	pool := db.ObjectField("pool")
	pool.Open()
	pool.IntegerField("min_connections", 5)
	pool.IntegerField("max_connections", 20)
	pool.FloatField("timeout", 30.0)
	pool.Close()

	db.Close()

	// Redis config
	redis := obj.ObjectField("redis")
	redis.Open()
	redis.StringField("host", "redis.example.com")
	redis.IntegerField("port", 6379)
	redis.IntegerField("db", 0)
	redis.Close()

	// Features array
	features := obj.ArrayField("features")
	features.Open()
	features.StringValue("authentication")
	features.StringValue("authorization")
	features.StringValue("caching")
	features.Close()

	// Environment-specific settings
	environments := obj.ArrayField("environments")
	environments.Open()

	// Development environment
	dev := environments.ObjectValue()
	dev.Open()
	dev.StringField("name", "development")
	dev.BooleanField("debug", true)
	dev.StringField("log_level", "debug")
	dev.Close()

	// Production environment
	prod := environments.ObjectValue()
	prod.Open()
	prod.StringField("name", "production")
	prod.BooleanField("debug", false)
	prod.StringField("log_level", "info")
	prod.Close()

	environments.Close()

	obj.Close()

	result, err := obj.BuildBytes()
	if err != nil {
		t.Fatalf("BuildBytes failed: %v", err)
	}

	// Parse and validate
	var parsed map[string]interface{}
	if err := json.Unmarshal(result, &parsed); err != nil {
		t.Fatalf("Generated JSON is invalid: %v, got: %s", err, string(result))
	}

	// Validate database config
	dbData, ok := parsed["database"].(map[string]interface{})
	if !ok {
		t.Fatal("Database field is not an object")
	}
	if host, exists := dbData["host"]; !exists || host != "localhost" {
		t.Error("Database host validation failed")
	}
	if port, exists := dbData["port"]; !exists || port != float64(5432) {
		t.Error("Database port validation failed")
	}
	if name, exists := dbData["name"]; !exists || name != "myapp" {
		t.Error("Database name validation failed")
	}
	if user, exists := dbData["user"]; !exists || user != "admin" {
		t.Error("Database user validation failed")
	}
	if ssl, exists := dbData["ssl"]; !exists || ssl != true {
		t.Error("Database SSL validation failed")
	}

	// Validate connection pool
	poolData, ok := dbData["pool"].(map[string]interface{})
	if !ok {
		t.Fatal("Database pool field is not an object")
	}
	if minConn, exists := poolData["min_connections"]; !exists || minConn != float64(5) {
		t.Error("Pool min_connections validation failed")
	}
	if maxConn, exists := poolData["max_connections"]; !exists || maxConn != float64(20) {
		t.Error("Pool max_connections validation failed")
	}
	if timeout, exists := poolData["timeout"]; !exists || timeout != 30.0 {
		t.Error("Pool timeout validation failed")
	}

	// Validate Redis config
	redisData, ok := parsed["redis"].(map[string]interface{})
	if !ok {
		t.Fatal("Redis field is not an object")
	}
	if host, exists := redisData["host"]; !exists || host != "redis.example.com" {
		t.Error("Redis host validation failed")
	}
	if port, exists := redisData["port"]; !exists || port != float64(6379) {
		t.Error("Redis port validation failed")
	}
	if db, exists := redisData["db"]; !exists || db != float64(0) {
		t.Error("Redis db validation failed")
	}

	// Validate features array
	featuresData, ok := parsed["features"].([]interface{})
	if !ok {
		t.Fatal("Features field is not an array")
	}
	expectedFeatures := []string{"authentication", "authorization", "caching"}
	if len(featuresData) != len(expectedFeatures) {
		t.Errorf("Expected %d features, got %d", len(expectedFeatures), len(featuresData))
	}
	for i, expected := range expectedFeatures {
		if i >= len(featuresData) {
			t.Errorf("Missing feature at index %d", i)
			continue
		}
		if feature, ok := featuresData[i].(string); !ok || feature != expected {
			t.Errorf("Feature at index %d: expected %q, got %v", i, expected, featuresData[i])
		}
	}

	// Validate environments array
	envsData, ok := parsed["environments"].([]interface{})
	if !ok {
		t.Fatal("Environments field is not an array")
	}
	if len(envsData) != 2 {
		t.Errorf("Expected 2 environments, got %d", len(envsData))
	}

	// Validate development environment
	devData, ok := envsData[0].(map[string]interface{})
	if !ok {
		t.Fatal("Development environment is not an object")
	}
	if name, exists := devData["name"]; !exists || name != "development" {
		t.Error("Development name validation failed")
	}
	if debug, exists := devData["debug"]; !exists || debug != true {
		t.Error("Development debug validation failed")
	}
	if logLevel, exists := devData["log_level"]; !exists || logLevel != "debug" {
		t.Error("Development log_level validation failed")
	}

	// Validate production environment
	prodData, ok := envsData[1].(map[string]interface{})
	if !ok {
		t.Fatal("Production environment is not an object")
	}
	if name, exists := prodData["name"]; !exists || name != "production" {
		t.Error("Production name validation failed")
	}
	if debug, exists := prodData["debug"]; !exists || debug != false {
		t.Error("Production debug validation failed")
	}
	if logLevel, exists := prodData["log_level"]; !exists || logLevel != "info" {
		t.Error("Production log_level validation failed")
	}
}

func TestAnyFieldIntegration(t *testing.T) {
	// Test AnyField with various types in a complex structure
	obj := jsoni.NewObjectWriter(nil)
	obj.Open()

	// Test AnyField with different types
	obj.AnyField("string_val", "hello")
	obj.AnyField("int_val", 42)
	obj.AnyField("int64_val", int64(12345))
	obj.AnyField("float_val", 3.14)
	obj.AnyField("bool_val", true)
	obj.AnyField("nil_val", nil)
	obj.AnyField("uint_val", uint(100))

	// Nested object using AnyField
	nested := obj.ObjectField("nested")
	nested.Open()
	nested.AnyField("mixed_array", []any{"string", 123, true, nil})
	nested.Close()

	obj.Close()

	result, err := obj.BuildBytes()
	if err != nil {
		t.Fatalf("BuildBytes failed: %v", err)
	}

	// Parse and validate
	var parsed map[string]interface{}
	if err := json.Unmarshal(result, &parsed); err != nil {
		t.Fatalf("Generated JSON is invalid: %v, got: %s", err, string(result))
	}

	// Validate AnyField values
	if val, exists := parsed["string_val"]; !exists || val != "hello" {
		t.Error("AnyField string validation failed")
	}
	if val, exists := parsed["int_val"]; !exists || val != float64(42) {
		t.Error("AnyField int validation failed")
	}
	if val, exists := parsed["int64_val"]; !exists || val != float64(12345) {
		t.Error("AnyField int64 validation failed")
	}
	if val, exists := parsed["float_val"]; !exists || val != 3.14 {
		t.Error("AnyField float validation failed")
	}
	if val, exists := parsed["bool_val"]; !exists || val != true {
		t.Error("AnyField bool validation failed")
	}
	if val, exists := parsed["nil_val"]; !exists || val != nil {
		t.Error("AnyField nil validation failed")
	}
	if val, exists := parsed["uint_val"]; !exists || val != float64(100) {
		t.Error("AnyField uint validation failed")
	}

	// Validate nested object
	nestedData, ok := parsed["nested"].(map[string]interface{})
	if !ok {
		t.Fatal("Nested field is not an object")
	}

	// The mixed array gets properly marshaled as JSON array by writeAny using json.Marshal
	if val, exists := nestedData["mixed_array"]; !exists {
		t.Error("Nested mixed_array field not found")
	} else {
		// The array gets properly marshaled as a JSON array
		if arrayVal, ok := val.([]interface{}); !ok {
			t.Errorf("Expected mixed_array to be marshaled as array, got %T", val)
		} else {
			// Validate the array contents
			expectedArray := []interface{}{"string", float64(123), true, nil}
			if len(arrayVal) != len(expectedArray) {
				t.Errorf("Expected %d array elements, got %d", len(expectedArray), len(arrayVal))
			}
			for i, expected := range expectedArray {
				if i >= len(arrayVal) {
					t.Errorf("Missing array element at index %d", i)
					continue
				}
				if arrayVal[i] != expected {
					t.Errorf("Array element at index %d: expected %v, got %v", i, expected, arrayVal[i])
				}
			}
		}
	}
}
