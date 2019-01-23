package model

import "github.com/jinzhu/gorm"

//角色应菜单表
type RoleMenu struct {
	gorm.Model
	RoleId int
	MenuId int
}
