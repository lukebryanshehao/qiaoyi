package repositorys

import (
	"qiaoyi_back/datasource"
	"qiaoyi_back/model"
)

type RoleRepository interface {
	PageQuery(page *model.Page) (int,[]model.Role)
	GetByID(id int) (model.Role,bool)
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
		page.PageSize = 10
	}
	var allcount int
	datasource.DB.Find(&roles).Count(&allcount)
	datasource.DB.Limit(page.PageSize).Offset(page.PageSize * (page.PageIndex-1)).Find(&roles)
	for i := 0;i< len(roles);i++  {
		var users []model.User
		datasource.DB.Select("id,name,areaid,roleid").Where("roleid = ?",roles[i].ID).Find(&users)
		roles[i].Users = users
	}
	return allcount,roles
}

func (r *roleMemoryRepository) GetByID(id int) (model.Role,bool) {
	flag := true
	var role model.Role
	if err := datasource.DB.First(&role, id).Error; err != nil {
		flag = false
		//panic(err)
	}
	return role,flag
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
