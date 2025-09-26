package models

import "time"

type LendingRecord struct {
	Id         uint       `json:"id" gorm:"primaryKey;autoIncrement"`
	BookId     uint       `json:"book_id"`
	UserId     uint       `json:"user_id"`
	BorrowDate time.Time  `json:"borrow_date"`
	ReturnDate *time.Time `json:"return_date"`
}
