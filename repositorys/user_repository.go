package repositorys

import (
	"qiaoyi_back/datasource"
	"qiaoyi_back/model"
)

type UserRepository interface {
	PageQuery(page *model.Page) (int,[]model.User)
	Save(user *model.User) (bool,*model.User)
	DeleteByID(id uint) (bool)
	GetByID(id uint) (bool,model.User)
}

func NewUserRepository() UserRepository {
	return &userMemoryRepository{}
}

type userMemoryRepository struct {
}

func (r *userMemoryRepository) PageQuery(page *model.Page) (int,[]model.User) {
	users := []model.User{}

	if page.PageSize == 0 {
		page.PageIndex = 0
		page.PageSize = 10
	}
	var allcount int
	datasource.DB.Find(&users).Count(&allcount)
	datasource.DB.Limit(page.PageSize).Offset(page.PageSize * page.PageIndex).Find(&users)
	for i := 0;i< len(users);i++  {
		areaid := users[i].Areaid
		roleid := users[i].Roleid
		var area model.Area
		var role model.Role
		datasource.DB.Table("areas").Where("id = ?", areaid).Scan(&area)
		datasource.DB.Table("roles").Where("id = ?", roleid).Scan(&role)
		users[i].Role = role
		users[i].Area = area
	}
	return allcount,users
}

func (r *userMemoryRepository) Save(user *model.User) (bool,*model.User) {
	flag := true
	if err := datasource.DB.Save(user).Error; err != nil {
		flag = false
		//panic(err)
	}
	return flag,user
}

func (r *userMemoryRepository) DeleteByID(id uint) (bool) {

	flag := true
	if err := datasource.DB.Delete(&model.User{}, id).Error; err != nil {
		flag = false
		//panic(err)
	}
	return flag
}

func (r *userMemoryRepository) GetByID(id uint) (bool,model.User) {
	flag := true
	var user model.User
	if err := datasource.DB.First(&user, id).Error; err != nil {
		flag = false
		//panic(err)
	}
	var area model.Area
	var role model.Role
	datasource.DB.Table("areas").Where("id = ?", user.Areaid).Scan(&area)
	datasource.DB.Table("roles").Where("id = ?", user.Roleid).Scan(&role)
	user.Role = role
	user.Area = area
	return flag,user
}
