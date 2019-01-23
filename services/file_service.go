package services

import (
	"qiaoyi_back/model"
	"qiaoyi_back/repositorys"
)

type FileService interface {
	GetMultimediaByID(id uint) (multimedia model.Multimedia)
	GetMultimedias(page *model.Page,query model.Multimedia) (multimedias []model.Multimedia)
	GetMultimediasByIds(ids []string) (multimedias []model.Multimedia)
	GetByQuery(query model.Multimedia) (multimedia model.Multimedia)
	DeleteByIds(ids []uint) bool
}

func NewFileService(repo repositorys.FileRepository) FileService {
	return &fileService{repo: repo,}
}

type fileService struct {
	repo repositorys.FileRepository
}

func (t *fileService) GetMultimediaByID(id uint) (multimedia model.Multimedia)  {
	return t.repo.GetMultimediaByID(id)
}
func (t *fileService) GetMultimediasByIds(ids []string) (multimedias []model.Multimedia)  {
	return t.repo.GetMultimediasByIds(ids)
}

func (t *fileService) GetMultimedias(page *model.Page,query model.Multimedia) (multimedias []model.Multimedia){
	return t.repo.GetMultimedias(page,query)
}

func (t *fileService) GetByQuery(query model.Multimedia) (multimedia model.Multimedia){
	return t.repo.GetByQuery(query)
}
func (t *fileService) DeleteByIds(ids []uint) bool{
	return t.repo.DeleteByIds(ids)
}