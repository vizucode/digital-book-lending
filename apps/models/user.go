package models

import (
	"time"
)

type Users struct {
	Id           uint      `gorm:"column:id;primaryKey;autoIncrement"`
	Name         string    `gorm:"column:name"`
	Email        string    `gorm:"column:email"`
	PasswordHash string    `gorm:"column:password_hash"`
	CreatedAt    time.Time `gorm:"column:created_at"`
	UpdatedAt    time.Time `gorm:"column:updated_at"`
}
