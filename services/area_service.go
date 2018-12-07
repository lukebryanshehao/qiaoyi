package services

import (
	"qiaoyi_back/model"
	"qiaoyi_back/repositorys"
)

type AreaService interface {
	Insert(area *model.Area) uint
	DeleteByID(id uint) bool
	Update(area *model.Area) bool
	PageQuery(page *model.Page) (int,[]model.Area)
	GetByID(id int) (model.Area,bool)
}

func NewAreaService(repo repositorys.AreaRepository) AreaService {
	return &areaService{repo: repo,}
}

type areaService struct {
	repo repositorys.AreaRepository
}

func (s *areaService) Insert(area *model.Area) uint {
	return s.repo.Insert(area)
}
func (s *areaService) DeleteByID(id uint) bool {
	return s.repo.DeleteByID(id)
}
func (s *areaService) Update(area *model.Area) bool {
	return s.repo.Update(area)
}
func (s *areaService) PageQuery(page *model.Page) (int,[]model.Area) {
	return s.repo.PageQuery(page)
}
func (s *areaService) GetByID(id int) (model.Area,bool) {
	return s.repo.GetByID(id)
}
