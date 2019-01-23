package repositorys

import (
	"qiaoyi_back/datasource"
	"qiaoyi_back/model"
)

type LoginRepository interface {
	Exist(user *model.User) bool
	GetInfo(username string) (model.User)
}

func NewLoginRepository() LoginRepository {
	return &loginMemoryRepository{}
}

type loginMemoryRepository struct {
}

func (r *loginMemoryRepository) Exist(user *model.User) bool {
	exist := true
	if err := datasource.DB.Where("username = ? and password = ?", user.Username,user.Password).First(&user).Error; err != nil {
		exist = false
		//panic(err)
	}
	var area model.Area
	var role model.Role
	datasource.DB.Where("id = ?", user.AreaId).Scan(&area)
	datasource.DB.Where("id = ?", user.RoleId).Scan(&role)
	user.Area = area
	user.Role = role
	return exist
}

func (s *loginMemoryRepository) GetInfo(username string) (user model.User) {
	if err := datasource.DB.Where("username = ?", username).First(&user).Error; err != nil {
		//panic(err)
	}
	var area model.Area
	var role model.Role
	datasource.DB.Where("id = ?", user.AreaId).Scan(&area)
	datasource.DB.Where("id = ?", user.RoleId).Scan(&role)
	user.Area = area
	user.Role = role
	return
}

