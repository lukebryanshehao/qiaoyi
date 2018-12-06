package model

import "github.com/jinzhu/gorm"

/*
角色权限表
 */
type RolePower struct {
	gorm.Model
	Roleid uint	`json:"RoleID"`//角色ID
	Resourcetype string	`json:"ResourceType"`//资源类型
	Resourceid	uint	`json:"ResourceID"`//资源ID
}
