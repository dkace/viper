package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name      string `gorm:"type:varchar(20);not null;column:name"`
	PassWord  string `gorm:"size 255;not null;column:password"`
	TelePhone string `gorm:"type:varchar(11);not null;column:telephone"`
}
