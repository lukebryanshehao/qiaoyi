package services

import (
	"qiaoyi_back/model"
	"qiaoyi_back/repositorys"
)

type SystemService interface {
	PageQuery(page *model.Page) (int,[]model.Setting)
	Save(setting *model.Setting) (bool,*model.Setting)
	DeleteByKey(key string) (bool)
	GetByKey(key string) (bool,model.Setting)
}

func NewSystemService(repo repositorys.SystemRepository) SystemService {
	return &systemService{repo: repo,}
}

type systemService struct {
	repo repositorys.SystemRepository
}

func (s *systemService) PageQuery(page *model.Page) (int,[]model.Setting) {
	return s.repo.PageQuery(page)
}
func (s *systemService) Save(setting *model.Setting) (bool,*model.Setting) {
	return s.repo.Save(setting)
}
func (s *systemService) DeleteByKey(key string) (bool) {
	return s.repo.DeleteByKey(key)
}
func (s *systemService) GetByKey(key string) (bool,model.Setting) {
	return s.repo.GetByKey(key)
}
