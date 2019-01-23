package datasource

import (
	"qiaoyi_back/model"
)

func Createtable() {
	DB.AutoMigrate(
		&model.Area{},
		&model.Article{},
		&model.User{},
		&model.Role{},
		&model.Area{},
		&model.Setting{},
		&model.SystemLog{}, //系统日志表
		&model.UserType{}, //用户类型
		&model.UserPower{}, //用户权限表
		&model.Setting{}, //设置表
		&model.RolePower{}, //角色权限表
		&model.RoleMenu{}, //角色应菜单表
		&model.Member{}, //前台登陆用户表
		&model.ClassifyType{}, //分类表

	)
	//if !DB.HasTable(&model.Area{}) {
	//	if err := DB.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").CreateTable(&model.Area{}).Error; err != nil {
	//		panic(err)
	//	}
	//}

}