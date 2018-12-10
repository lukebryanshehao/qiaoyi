package repositorys

import (
	"qiaoyi_back/datasource"
	"qiaoyi_back/model"
)

type SystemRepository interface {
	PageQuery(page *model.Page) (int,[]model.Setting)
	Save(setting *model.Setting) (bool,*model.Setting)
	DeleteByKey(key string) (bool)
	GetByKey(key string) (bool,model.Setting)
}

func NewSystemRepository() SystemRepository {
	return &systemMemoryRepository{}
}

type systemMemoryRepository struct {
}

func (r *systemMemoryRepository) PageQuery(page *model.Page) (int,[]model.Setting) {
	settings := []model.Setting{}

	if page.PageSize == 0 {
		page.PageIndex = 0
		page.PageSize = 10
	}
	var allcount int
	datasource.DB.Find(&settings).Count(&allcount)
	datasource.DB.Limit(page.PageSize).Offset(page.PageSize * page.PageIndex).Find(&settings)
	return allcount,settings
}

func (r *systemMemoryRepository) Save(setting *model.Setting) (bool,*model.Setting) {
	flag := true
	if err := datasource.DB.Save(setting).Error; err != nil {
		flag = false
		//panic(err)
	}
	return flag,setting
}

func (r *systemMemoryRepository) DeleteByKey(key string) (bool) {

	flag := true
	if err := datasource.DB.Delete(&model.Setting{}, key).Error; err != nil {
		flag = false
		//panic(err)
	}
	return flag
}

func (r *systemMemoryRepository) GetByKey(key string) (bool,model.Setting) {
	flag := true
	var setting model.Setting
	if err := datasource.DB.First(&setting, key).Error; err != nil {
		flag = false
		//panic(err)
	}
	return flag,setting
}
