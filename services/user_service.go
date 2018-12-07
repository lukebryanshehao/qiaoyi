package services

import (
	"qiaoyi_back/model"
	"qiaoyi_back/repositorys"
)

type UserService interface {
	PageQuery(page *model.Page) (int,[]model.User)
	Save(user *model.User) (bool,*model.User)
	DeleteByID(id uint) (bool)
	GetByID(id uint) (bool,model.User)
}

func NewUserService(repo repositorys.UserRepository) UserService {
	return &userService{repo: repo,}
}

type userService struct {
	repo repositorys.UserRepository
}

func (s *userService) PageQuery(page *model.Page) (int,[]model.User) {
	return s.repo.PageQuery(page)
}

func (s *userService) Save(user *model.User) (bool,*model.User) {
	return s.repo.Save(user)
}

func (s *userService) DeleteByID(id uint) (bool) {
	return s.repo.DeleteByID(id)
}

func (s *userService) GetByID(id uint) (bool,model.User) {
	return s.repo.GetByID(id)
}
