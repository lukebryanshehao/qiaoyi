package repositorys

import (
	"fmt"
	"qiaoyi_back/datasource"
	"qiaoyi_back/model"
)

type LoginRepository interface {
	Exist(user *model.User) bool
	GetInfo(username string) (user model.User)
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
	}
	return exist
}

func (s *loginMemoryRepository) GetInfo(username string) (user model.User) {
	if err := datasource.DB.Where("username = ?", username).First(&user).Error; err != nil {
		fmt.Println(err)
	}
	var area model.Area
	var role model.Role
	datasource.DB.Where("id = ?", user.AreaId).First(&area)
	datasource.DB.Where("id = ?", user.RoleId).First(&role)
	user.Area = area
	user.Role = role
	return
}

