package repositorys

import (
	"qiaoyi_back/datasource"
	"qiaoyi_back/model"
	"strings"
)

type FileRepository interface {
	GetMultimediaByID(id uint) (multimedia model.Multimedia)
	GetMultimedias(page *model.Page,query model.Multimedia) (multimedias []model.Multimedia)
	GetMultimediasByIds(ids []string) (multimedias []model.Multimedia)
	GetByQuery(query model.Multimedia) (multimedia model.Multimedia)
	DeleteByIds(ids []uint) bool
}
type fileRepository struct{ d FileRepository }

func NewFileRepository() *fileRepository { return nil }

func (d *fileRepository) GetMultimediaByID(id uint) (multimedia model.Multimedia){

	datasource.DB.Where("id = ?",id).First(&multimedia)

	return
}
func (d *fileRepository) GetMultimediasByIds(ids []string) (multimedias []model.Multimedia){

	datasource.DB.Where("id in (?)",ids).Find(&multimedias)

	return
}

func (d *fileRepository) GetMultimedias(page *model.Page,query model.Multimedia) (multimedias []model.Multimedia){

	if page.PageSize == 0 {
		page.PageSize = 10
	}

	db := datasource.DB

	if query.Year != "" {
		db = db.Where("year = ?",query.Year)
	}
	if query.Name != "" {
		db = db.Where("Name like ?","%"+query.Name+"%")
	}
	//"reviewtoknow"/"operatemanual"
	db = db.Where("Remark = 'reviewtoknow' or Remark = 'operatemanual'")
	var count int
	db.Model(&model.Multimedia{}).Count(&count)
	page.PageCount = count
	db.Limit(page.PageSize).Offset(model.PageIndex(page)).Order("id desc").Find(&multimedias)
	for index,multimedia := range multimedias{
		multimedias[index].Name = multimedia.Name[0:strings.Index(multimedia.Name, ".")]
		var user model.User
		datasource.DB.First(&user,multimedia.CreateUserid)
		multimedias[index].CreateUser = user
	}

	return
}

func (d *fileRepository) GetByQuery(query model.Multimedia) (multimedia model.Multimedia){
	db := datasource.DB

	if query.Year != "" {
		db = db.Where("year = ?",query.Year)
	}
	if query.Name != "" {
		db = db.Where("Name like ?","%"+query.Name+"%")
	}
	if query.Remark != "" {
		db = db.Where("Remark = ?",query.Remark)
	}

	db.First(&multimedia)
	return
}

func (d *fileRepository) DeleteByIds(ids []uint) bool{
	flag := true
	tran := datasource.DB.Begin()

	defer func() {
		//恢复程序的控制权
		err := recover()
		if err == nil {
			//提交事务
			tran.Commit()
		} else {
			//回滚
			tran.Rollback()
		}
	}()
	if err := tran.Where("id in (?)",ids).Delete(&model.Multimedia{}).Error; err != nil{
		flag = false
	}
	return flag
}