package model

import "github.com/jinzhu/gorm"

/*
角色表
 */
type Role struct {
	gorm.Model
	Rolename string	`json:"RoleName"`//角色名
	Weight uint	`json:"Weight"`//权重
	Remark string	`json:"Remark"`//备注
	Users []User
}
