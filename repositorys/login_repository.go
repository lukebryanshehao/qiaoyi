package repositorys

import (
	"qiaoyi_back/datasource"
	"qiaoyi_back/model"
)

type LoginRepository interface {
	Exist(user *model.User) (bool,model.User)
	GetInfo(username string) (model.User)
}

func NewLoginRepository() LoginRepository {
	return &loginMemoryRepository{}
}

type loginMemoryRepository struct {
}

func (r *loginMemoryRepository) Exist(user *model.User) (bool,model.User) {
	var user1 model.User
	exist := true
	if err := datasource.DB.Table("users").Where("username = ? and password = ?", user.Username,user.Password).Scan(&user1).Error; err != nil {
		exist = false
		//panic(err)
	}
	var area model.Area
	var role model.Role
	datasource.DB.Table("areas").Where("id = ?", user1.Areaid).Scan(&area)
	datasource.DB.Table("roles").Where("id = ?", user1.Roleid).Scan(&role)
	user1.Role = role
	user1.Area = area
	return exist,user1
}

func (s *loginMemoryRepository) GetInfo(username string) (model.User) {
	var user1 model.User
	if err := datasource.DB.Table("users").Where("username = ?", username).Scan(&user1).Error; err != nil {
		//panic(err)
	}
	var area model.Area
	var role model.Role
	datasource.DB.Table("areas").Where("id = ?", user1.Areaid).Scan(&area)
	datasource.DB.Table("roles").Where("id = ?", user1.Roleid).Scan(&role)
	user1.Role = role
	user1.Area = area
	return user1
}

