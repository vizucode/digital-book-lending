package models

type Book struct {
	Id       uint   `json:"id" gorm:"primaryKey;autoIncrement"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	Isbn     string `json:"isbn"`
	Quantity int    `json:"quantity"`
	Category string `json:"category"`
}
