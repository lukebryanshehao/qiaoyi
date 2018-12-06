package services

import (
	"qiaoyi_back/model"
	"qiaoyi_back/repositorys"
)

type RoleService interface {
	PageQuery(page *model.Page) (*[]model.Role)
	Save(role *model.Role) (bool,*model.Role)
	DeleteByID(id uint) (bool)
}

func NewRoleService(repo repositorys.RoleRepository) RoleService {
	return &roleService{repo: repo,}
}

type roleService struct {
	repo repositorys.RoleRepository
}

func (s *roleService) PageQuery(page *model.Page) (*[]model.Role) {
	return s.repo.PageQuery(page)
}

func (s *roleService) Save(role *model.Role) (bool,*model.Role) {
	return s.repo.Save(role)
}

func (s *roleService) DeleteByID(id uint) (bool) {
	return s.repo.DeleteByID(id)
}
