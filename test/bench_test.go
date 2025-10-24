package test

import (
	"encoding/json"
	"testing"

	"github.com/mailru/easyjson/jwriter"
)

var users = generateUsers(1000)
var posts = generatePosts(users, 5)

// ----------------- Benchmark: encoding/json -----------------
func BenchmarkEncodingJSON_Users(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = json.Marshal(users)
	}
}

func BenchmarkEncodingJSON_Posts(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = json.Marshal(posts)
	}
}

// ----------------- Benchmark: easyjson ----------------------
func BenchmarkEasyJSON_Users(b *testing.B) {
	for i := 0; i < b.N; i++ {
		w := &jwriter.Writer{}
		for _, u := range users {
			u.MarshalEasyJSON(w)
		}
		_, _ = w.BuildBytes()
	}
}

func BenchmarkEasyJSON_Posts(b *testing.B) {
	for i := 0; i < b.N; i++ {
		w := &jwriter.Writer{}
		for _, p := range posts {
			p.MarshalEasyJSON(w)
		}
		_, _ = w.BuildBytes()
	}
}

// ----------------- Benchmark: jsoni -------------------------
func BenchmarkJsoniWriter_Users(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = writeUsersJsoni(users)
	}
}

func BenchmarkJsoniWriter_Posts(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = writePostsJsoni(posts)
	}
}

// ----------------- Benchmark: jsondi ------------------------
func BenchmarkJsondiWriter_Users(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = writeUsersJsondi(users)
	}
}

func BenchmarkJsondiWriter_Posts(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = writePostsJsondi(posts)
	}
}

// ----------------- Benchmark: jsondf ------------------------
func BenchmarkJsondfWriter_Users(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = writeUsersJsondf(users)
	}
}

func BenchmarkJsondfWriter_Posts(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = writePostsJsondf(posts)
	}
}

// ----------------- Benchmark: jsonds ------------------------
func BenchmarkJsondsWriter_Users(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = writeUsersJsonds(users)
	}
}

func BenchmarkJsondsWriter_Posts(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = writePostsJsonds(posts)
	}
}
