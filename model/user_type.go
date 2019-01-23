package model

import "github.com/jinzhu/gorm"

//用户类型表
type UserType struct {
	gorm.Model
	Name     string
	Superior int
	Remarks  string
	Role     []Role `gorm:"ForeignKey:Id"`
	User     []User `gorm:"ForeignKey:id"`
}
