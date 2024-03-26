package data

type User struct {
	ID         int           `json:"id"`
	Username   string        `json:"username"`
	FullName   string        `json:"fullname"`
	Password   string        `json:"-"`
	Role       string        `json:"role"`
	Addressses []UserAddress `json:"addresses"`
}

type UserAddress struct {
	ID       int    `json:"id"`
	Address  string `json:"address"`
	IsActive bool   `json:"is_active"`
}
