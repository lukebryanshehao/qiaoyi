package model

import "github.com/jinzhu/gorm"

/*
用户权限表
 */
type UserPower struct {
	gorm.Model
	Userid uint	`json:"UserID"`//用户ID
	Resourcetype string	`json:"ResourceType"`//资源类型
	Resourceid	uint	`json:"ResourceID"`//资源ID
}
