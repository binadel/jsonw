package test

import json "github.com/binadel/jsonw/jsondf"

func writeUsersJsondf(users []User) []byte {
	items := make([]json.Value, len(users))
	for i, u := range users {
		tags := make([]json.Value, len(u.Tags))
		for j, tag := range u.Tags {
			tags[j] = json.StringItem(tag)
		}

		addresses := make([]json.Value, len(u.Addresses))
		for j, addr := range u.Addresses {
			addresses[j] = json.ObjectItem(
				json.String("street", addr.Street),
				json.String("city", addr.City),
				json.String("zip", addr.Zip),
				json.String("country", addr.Country),
			)
		}

		items[i] = json.ObjectItem(
			json.Integer("id", u.ID),
			json.String("name", u.Name),
			json.String("email", u.Email),
			json.Boolean("is_active", u.IsActive),
			json.Integer("age", int64(u.Age)),
			json.Float("balance", u.Balance),
			json.Array("tags", tags...),
			json.Object("profile",
				json.String("bio", u.Profile.Bio),
				json.String("avatar_url", u.Profile.AvatarURL),
			),
			json.Array("addresses", addresses...),
		)
	}
	r := json.NewArray(items...)
	b, _ := r.Build()
	return b
}

func writePostsJsondf(posts []Post) []byte {
	items := make([]json.Value, len(posts))
	for i, p := range posts {
		tags := make([]json.Value, len(p.Tags))
		for j, tag := range p.Tags {
			tags[j] = json.StringItem(tag)
		}

		comments := make([]json.Value, len(p.Comments))
		for j, comment := range p.Comments {
			comments[j] = json.ObjectItem(
				json.Integer("id", comment.ID),
				json.Integer("user_id", comment.UserID),
				json.String("message", comment.Message),
			)
		}

		items[i] = json.ObjectItem(
			json.Integer("id", p.ID),
			json.Integer("user_id", p.UserID),
			json.String("title", p.Title),
			json.String("content", p.Content),
			json.Array("tags", tags...),
			json.Integer("likes", int64(p.Likes)),
			json.Array("comments", comments...),
		)
	}
	r := json.NewArray(items...)
	b, _ := r.Build()
	return b
}
