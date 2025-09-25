package domain

type UserContext struct {
	Id       int    `json:"id"`
	FullName string `json:"fullname"`
	Email    string `json:"email"`
	Exp      int64  `json:"exp"`
}
