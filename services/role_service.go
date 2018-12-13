package services

import (
	"qiaoyi_back/model"
	"qiaoyi_back/repositorys"
)

type RoleService interface {
	PageQuery(page *model.Page) (int,[]model.Role)
	GetByID(id int) (model.Role,bool)
	Save(role *model.Role) (bool,*model.Role)
	DeleteByIDs(ids []uint) (bool)
}

func NewRoleService(repo repositorys.RoleRepository) RoleService {
	return &roleService{repo: repo,}
}

type roleService struct {
	repo repositorys.RoleRepository
}

func (s *roleService) PageQuery(page *model.Page) (int,[]model.Role) {
	return s.repo.PageQuery(page)
}

func (s *roleService) GetByID(id int) (model.Role,bool) {
	return s.repo.GetByID(id)
}

func (s *roleService) Save(role *model.Role) (bool,*model.Role) {
	return s.repo.Save(role)
}

func (s *roleService) DeleteByIDs(ids []uint) (bool) {
	return s.repo.DeleteByIDs(ids)
}
