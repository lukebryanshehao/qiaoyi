package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

//用户表
type User struct {
	gorm.Model
	Username      string //登录名
	Password       string //登录密码
	Name           string //真实姓名
	Phone          string //手机号
	Sex				uint	`gorm:"default:1"`  //性别	1男,2女
	DepartmentName string //所属部门
	AreaId			uint	//地区
	RoleId			uint	//角色
	Email          string //邮件
	Birthday       string //生日
	Address       string //联系地址
	UserTypeId     int //所属用户类型ID
	State          int //状态 1.正常,0.待审核,3.禁用,4.未通过
	UserType       UserType
	Role           Role
	Area           Area
	OldPassword       string	`gorm:"-"` //旧密码
	Session       string	`gorm:"-"` //Session
	AccessTime       time.Time	`gorm:"-"` //登陆时间
	LastAccessTime       time.Time	`gorm:"-"` //最后访问时间
	RemoteAddr       string	`gorm:"-"` //IP
	MaxLifeTime       int64	`gorm:"-"` //session过期时间
}

