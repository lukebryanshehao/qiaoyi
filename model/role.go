package model

import "github.com/jinzhu/gorm"

//角色表
type Role struct {
	gorm.Model
	Name       string //角色名
	UserTypeId int    //角色对应的用户类型
	Sort       int    //排序
	Remarks    string //备注
	UserType   UserType `gorm:"ForeignKey:UserTypeId"`
	Users       []User   `gorm:"ForeignKey:id"`
}
