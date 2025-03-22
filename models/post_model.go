package models

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	Title string `gorm:"type:varchar(100)"`
	Body  string
}
