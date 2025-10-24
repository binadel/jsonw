package test

type User struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	IsActive  bool      `json:"is_active"`
	Age       int       `json:"age"`
	Balance   float64   `json:"balance"`
	Tags      []string  `json:"tags"`
	Profile   Profile   `json:"profile"`
	Addresses []Address `json:"addresses"`
}

type Profile struct {
	Bio       string `json:"bio"`
	AvatarURL string `json:"avatar_url"`
}

type Address struct {
	Street  string `json:"street"`
	City    string `json:"city"`
	Zip     string `json:"zip"`
	Country string `json:"country"`
}

type Post struct {
	ID       int64     `json:"id"`
	UserID   int64     `json:"user_id"`
	Title    string    `json:"title"`
	Content  string    `json:"content"`
	Tags     []string  `json:"tags"`
	Likes    int       `json:"likes"`
	Comments []Comment `json:"comments"`
}

type Comment struct {
	ID      int64  `json:"id"`
	UserID  int64  `json:"user_id"`
	Message string `json:"message"`
}
