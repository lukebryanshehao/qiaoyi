package repositorys

import (
	"qiaoyi_back/datasource"
	"qiaoyi_back/model"
)

type RoleRepository interface {
	PageQuery(page *model.Page) (*[]model.Role)
	Save(role *model.Role) (bool,*model.Role)
	DeleteByID(id uint) (bool)
}

func NewRoleRepository() RoleRepository {
	return &roleMemoryRepository{}
}

type roleMemoryRepository struct {
}

func (r *roleMemoryRepository) PageQuery(page *model.Page) (roles *[]model.Role) {
	roles = &[]model.Role{}

	if page.PageSize == 0 {
		page.PageIndex = 0
		page.PageSize = 10
	}

	datasource.DB.Limit(page.PageSize).Offset(page.PageSize * page.PageIndex).Find(roles)
	return roles
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
