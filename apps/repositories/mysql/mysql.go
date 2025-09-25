package mysql

import "gorm.io/gorm"

type mysql struct {
	db *gorm.DB
}

func NewMysql(
	db *gorm.DB,
) *mysql {
	return &mysql{
		db: db,
	}
}
