package model

import (
	"github.com/jinzhu/gorm"
	"qiaoyi_back/utils"
)

type User struct {
	gorm.Model
	Username string	`json:"Username"`//用户名
	Password string `json:"Password"`//登陆密码
	Name string	`json:"Name"`//名称
	Age uint	`json:"Age"`//年龄
	Areaid uint	`json:"AreaID"`//所属地区ID
	Roleid uint	`json:"RoleID"`//所属角色ID
	Checkcode string `gorm:"-"`	//验证码
	Area Area
	Role Role
}

func (u *User) SetPwd(p string) {
	u.Password = utils.Md5(u.Username + p)
}
