package model

import "github.com/jinzhu/gorm"

/*
多媒体
 */
type Multimedia struct {
	gorm.Model
	Name             string //名称
	Path             string //可直接显示的路径
	Year			string	//年度
	Url				string 	`gorm:"-"` //用于接收参数
	DownloadFilePath string //提供下载的路径
	Tablename        string //类型
	Tableid          uint   //ID
	CreateUserid	uint	//创建用户ID
	Remark			string	//备注
	CreateUser		User `gorm:"-"`	//创建用户
}
