package repositorys

import (
	"qiaoyi_back/datasource"
	"qiaoyi_back/model"
)

type RoleRepository interface {
	PageQuery(page *model.Page) (int,[]model.Role)
	Save(role *model.Role) (bool,*model.Role)
	DeleteByID(id uint) (bool)
}

func NewRoleRepository() RoleRepository {
	return &roleMemoryRepository{}
}

type roleMemoryRepository struct {
}

func (r *roleMemoryRepository) PageQuery(page *model.Page) (int,[]model.Role) {
	roles := []model.Role{}

	if page.PageSize == 0 {
		page.PageIndex = 0
		page.PageSize = 10
	}
	var allcount int
	datasource.DB.Find(&roles).Count(&allcount)
	datasource.DB.Limit(page.PageSize).Offset(page.PageSize * page.PageIndex).Find(&roles)
	for i := 0;i< len(roles);i++  {
		var users []model.User
		datasource.DB.Where("roleid = ?",roles[i].ID).Find(&users)
		roles[i].Users = users
	}
	return allcount,roles
}

func (r *roleMemoryRepository) Save(role *model.Role) (bool,*model.Role) {
	flag := true
	if err := datasource.DB.Save(role).Error; err != nil {
		flag = false
		//panic(err)
	}
	return flag,role
}
func (r *roleMemoryRepository) DeleteByID(id uint) (bool) {

	flag := true
	if err := datasource.DB.Delete(&model.Role{}, id).Error; err != nil {
		flag = false
		//panic(err)
	}
	return flag
}
