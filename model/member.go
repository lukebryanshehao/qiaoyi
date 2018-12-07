package model

import "github.com/jinzhu/gorm"

type Member struct {
	gorm.Model
	Username string	`json:"Username"`//用户名
	Password string `json:"Password"`//登陆密码
	Name string	`json:"Name"`//名称
	Age uint	`json:"Age"`//年龄
	Phone string	`json:"Phone"`//电话
	Email string	`json:"Email"`//邮箱
	Address uint	`json:"Address"`//住址
}
