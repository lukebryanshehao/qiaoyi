package services

import (
	"qiaoyi_back/model"
	"qiaoyi_back/repositorys"
)

type LoginService interface {
	Exist(user *model.User) bool
	GetInfo(username string) (model.User)
}

func NewLoginService(repo repositorys.LoginRepository) LoginService {
	return &loginService{repo: repo,}
}

type loginService struct {
	repo repositorys.LoginRepository
}

func (s *loginService) Exist(user *model.User) bool {
	return s.repo.Exist(user)
}

func (s *loginService) GetInfo(username string) (model.User) {
	return s.repo.GetInfo(username)
}
