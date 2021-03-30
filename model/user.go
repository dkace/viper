package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name      string `gorm:"type:varchar(20);not null;column:name"`
	TelePhone string `gorm:"type:varchar(11);not null;column:telephone"`
	PassWord  string `gorm:"size 255;not null;column:password"`
}
