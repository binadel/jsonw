package test

import "github.com/binadel/jsonw/jsoni"

func writeUsersJsoni(users []User) []byte {
	writer := jsoni.NewArrayWriter(nil)
	writer.Open()
	for _, u := range users {
		obj := writer.ObjectValue()
		obj.Open()
		obj.IntegerField("id", u.ID)
		obj.StringField("name", u.Name)
		obj.StringField("email", u.Email)
		obj.BooleanField("is_active", u.IsActive)
		obj.IntegerField("age", int64(u.Age))
		obj.FloatField("balance", u.Balance)

		tags := obj.ArrayField("tags")
		tags.Open()
		for _, t := range u.Tags {
			tags.StringValue(t)
		}
		tags.Close()

		profile := obj.ObjectField("profile")
		profile.Open()
		profile.StringField("bio", u.Profile.Bio)
		profile.StringField("avatar_url", u.Profile.AvatarURL)
		profile.Close()

		addrArr := obj.ArrayField("addresses")
		addrArr.Open()
		for _, a := range u.Addresses {
			addrObj := addrArr.ObjectValue()
			addrObj.Open()
			addrObj.StringField("street", a.Street)
			addrObj.StringField("city", a.City)
			addrObj.StringField("zip", a.Zip)
			addrObj.StringField("country", a.Country)
			addrObj.Close()
		}
		addrArr.Close()

		obj.Close()
	}
	writer.Close()
	bytes, _ := writer.BuildBytes()
	return bytes
}

func writePostsJsoni(posts []Post) []byte {
	writer := jsoni.NewArrayWriter(nil)
	writer.Open()
	for _, p := range posts {
		obj := writer.ObjectValue()
		obj.Open()
		obj.IntegerField("id", p.ID)
		obj.IntegerField("user_id", p.UserID)
		obj.StringField("title", p.Title)
		obj.StringField("content", p.Content)
		tags := obj.ArrayField("tags")
		tags.Open()
		for _, t := range p.Tags {
			writer.StringValue(t)
		}
		tags.Close()
		obj.IntegerField("likes", int64(p.Likes))

		comments := obj.ArrayField("comments")
		comments.Open()
		for _, c := range p.Comments {
			commentObj := comments.ObjectValue()
			commentObj.Open()
			commentObj.IntegerField("id", c.ID)
			commentObj.IntegerField("user_id", c.UserID)
			commentObj.StringField("message", c.Message)
			commentObj.Close()
		}
		comments.Close()

		obj.Close()
	}
	writer.Close()
	bytes, _ := writer.BuildBytes()
	return bytes
}
