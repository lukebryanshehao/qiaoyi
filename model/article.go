package model

import "github.com/jinzhu/gorm"

type Article struct {
	gorm.Model
	Title string
	Subtitle string
	Content string `gorm:"type:text"`
	AuthorID uint	//UserID

}
