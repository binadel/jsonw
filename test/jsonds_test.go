package test

import (
	js "encoding/json"
	"testing"

	json "github.com/binadel/jsonw/jsonds"
)

func TestJsonds(t *testing.T) {
	users := generateUsers(100)
	usersJson, _ := js.Marshal(users)
	usersJsonds := writeUsersJsonds(users)

	if string(usersJsonds) != string(usersJson) {
		t.Error("Users json are different")
	}

	posts := generatePosts(users, 50)
	postsJson, _ := js.Marshal(posts)
	postsJsonds := writePostsJsonds(posts)

	if string(postsJsonds) != string(postsJson) {
		t.Error("Posts json are different")
	}
}

func writeUsersJsonds(users []User) []byte {
	items := make([]json.Value, len(users))
	for i, u := range users {
		var tagsField json.Field
		if u.Tags != nil {
			tags := make([]json.Value, len(u.Tags))
			for j, tag := range u.Tags {
				tags[j] = json.StringItem(tag)
			}
			tagsField = json.Array("tags", tags...)
		} else {
			tagsField = json.Null("tags")
		}

		var addressesField json.Field
		if u.Addresses != nil {
			addresses := make([]json.Value, len(u.Addresses))
			for j, addr := range u.Addresses {
				addresses[j] = json.ObjectItem(
					json.String("street", addr.Street),
					json.String("city", addr.City),
					json.String("zip", addr.Zip),
					json.String("country", addr.Country),
				)
			}
			addressesField = json.Array("addresses", addresses...)
		} else {
			addressesField = json.Null("addresses")
		}

		items[i] = json.ObjectItem(
			json.Integer("id", u.ID),
			json.String("name", u.Name),
			json.String("email", u.Email),
			json.Boolean("is_active", u.IsActive),
			json.Integer("age", int64(u.Age)),
			json.Float("balance", u.Balance),
			tagsField,
			json.Object("profile",
				json.String("bio", u.Profile.Bio),
				json.String("avatar_url", u.Profile.AvatarURL),
			),
			addressesField,
		)
	}
	r := json.NewArray(items...)
	b, _ := r.Build()
	return b
}

func writePostsJsonds(posts []Post) []byte {
	items := make([]json.Value, len(posts))
	for i, p := range posts {
		var tagsField json.Field
		if p.Tags != nil {
			tags := make([]json.Value, len(p.Tags))
			for j, tag := range p.Tags {
				tags[j] = json.StringItem(tag)
			}
			tagsField = json.Array("tags", tags...)
		} else {
			tagsField = json.Null("tags")
		}

		var commentsField json.Field
		if p.Comments != nil {
			comments := make([]json.Value, len(p.Comments))
			for j, comment := range p.Comments {
				comments[j] = json.ObjectItem(
					json.Integer("id", comment.ID),
					json.Integer("user_id", comment.UserID),
					json.String("message", comment.Message),
				)
			}
			commentsField = json.Array("comments", comments...)
		} else {
			commentsField = json.Null("comments")
		}

		items[i] = json.ObjectItem(
			json.Integer("id", p.ID),
			json.Integer("user_id", p.UserID),
			json.String("title", p.Title),
			json.String("content", p.Content),
			tagsField,
			json.Integer("likes", int64(p.Likes)),
			commentsField,
		)
	}
	r := json.NewArray(items...)
	b, _ := r.Build()
	return b
}
