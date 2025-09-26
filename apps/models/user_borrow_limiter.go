package models

import "time"

type UserBorrowLimiter struct {
	Id          uint      `gorm:"primaryKey;autoIncrement"`
	UserId      uint      `gorm:"index;column:user_id"`
	BorrowCount uint      `gorm:"column:borrow_count;default:0"`
	WindowStart time.Time `gorm:"column:window_start;default:0"`
	WindowEnd   time.Time `gorm:"column:window_end;default:0"`
}
