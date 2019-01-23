package model

import "github.com/jinzhu/gorm"

type SystemLog struct {
	gorm.Model
	UserId			uint									//用户ID
	LoginName		string		`gorm:"type:varchar(25)"`	//用户名
	Operate			string		`gorm:"type:varchar(50)"`	//操作
	Method			string		`gorm:"type:varchar(100)"`	//请求方法
	Parameter		string		`gorm:"type:varchar(5000)"`	//参数
	Ip				string		`gorm:"type:varchar(25)"`	//Ip
}