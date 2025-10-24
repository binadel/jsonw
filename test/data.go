package test

import "math/rand"

func generateUsers(n int) []User {
	users := make([]User, n)
	for i := 0; i < n; i++ {
		users[i] = User{
			ID:       int64(i + 1),
			Name:     "User_" + randomString(5),
			Email:    "user" + randomString(3) + "@example.com",
			IsActive: rand.Intn(2) == 1,
			Age:      rand.Intn(50) + 18,
			Balance:  float64(rand.Intn(100000)) / 100.0,
			Tags:     []string{"tag1", "tag2", "tag3"},
			Profile: Profile{
				Bio:       "This is a bio for user " + randomString(4),
				AvatarURL: "https://example.com/avatar/" + randomString(6) + ".png",
			},
			Addresses: []Address{
				{
					Street:  "Street " + randomString(3),
					City:    "City_" + randomString(4),
					Zip:     "12345",
					Country: "Country_" + randomString(3),
				},
				{
					Street:  "Street " + randomString(3),
					City:    "City_" + randomString(4),
					Zip:     "67890",
					Country: "Country_" + randomString(3),
				},
			},
		}
	}
	return users
}

func generatePosts(users []User, postsPerUser int) []Post {
	var posts []Post
	id := int64(1)
	for _, user := range users {
		for j := 0; j < postsPerUser; j++ {
			p := Post{
				ID:      id,
				UserID:  user.ID,
				Title:   "Post Title " + randomString(5),
				Content: "This is the content of post " + randomString(10),
				Tags:    []string{"go", "json", "benchmark"},
				Likes:   rand.Intn(1000),
			}
			numComments := rand.Intn(5)
			for k := 0; k < numComments; k++ {
				comment := Comment{
					ID:      int64(k + 1),
					UserID:  users[rand.Intn(len(users))].ID,
					Message: "Comment message " + randomString(7),
				}
				p.Comments = append(p.Comments, comment)
			}
			posts = append(posts, p)
			id++
		}
	}
	return posts
}

func randomString(n int) string {
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	b := make([]byte, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
