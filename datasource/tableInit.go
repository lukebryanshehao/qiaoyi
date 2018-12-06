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

	)
	//if !DB.HasTable(&model.Area{}) {
	//	if err := DB.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").CreateTable(&model.Area{}).Error; err != nil {
	//		panic(err)
	//	}
	//}

}