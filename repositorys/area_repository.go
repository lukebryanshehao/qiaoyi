package repositorys

import (
	_ "github.com/Go-SQL-Driver/MySQL"
	"qiaoyi_back/datasource"
	"qiaoyi_back/model"
	"time"
)

type AreaRepository interface {
	Insert(area *model.Area) uint
	DeleteByID(id uint) bool
	Update(area *model.Area) bool
	PageQuery(page *model.Page) (int,[]model.Area)
	GetByID(id int) (model.Area,bool)
}

func NewAreaRepository() AreaRepository {
	return &areaMemoryRepository{}
}

type areaMemoryRepository struct {
}

func (r *areaMemoryRepository) Insert(area *model.Area) uint {
	area.CreatedAt = time.Now()
	if err := datasource.DB.Create(area).Error; err != nil {
		area.ID = 0
		panic(err)
	}
	return area.ID
}
func (r *areaMemoryRepository) DeleteByID(id uint) bool {
	var area model.Area
	area.ID = id
	flag := true
	//datasource.DB.Delete(&area,id)
	if err := datasource.DB.Delete(&area).Error; err != nil{
		flag = false
		panic(err)
	}

	return flag
}
func (r *areaMemoryRepository) Update(area *model.Area) bool {
	//datasource.DB.Exec("update area set id = ?,Name = ?,UnitCode = ?,Level = ?,PID = ?,SupervisionGroupID = ?,CreatedAt = ? where id = ?")
	datasource.DB.Begin()
	flag := true
	datasource.DB.Save(area)
	//datasource.DB.Model(&area).Where("id = ?",area.ID).Update(map[string]interface{}{"Name":area.Name,"InnerCode":area.InnerCode,"Level":area.Level,"PID":area.PID})
	if err := datasource.DB.Commit().Error; err != nil {//建议使用map修改需要修改的字段
		flag = false
		datasource.DB.Rollback()
		panic(err)
	}
	return flag
}
func (r *areaMemoryRepository) PageQuery(page *model.Page) (int,[]model.Area) {
	var areaArr []model.Area
	if page.PageSize == 0 {
		page.PageIndex = 0
		page.PageSize = 10
	}
	var allcount int
	datasource.DB.Find(&areaArr).Count(&allcount)
	if err := datasource.DB.Table("areas").Limit(page.PageSize).Offset(page.PageSize * page.PageIndex).Scan(&areaArr).Error; err != nil{
		panic(err)
	}
	for i := 0;i< len(areaArr);i++  {
		var users []model.User
		datasource.DB.Where("areaid = ?",areaArr[i].ID).Find(&users)
		areaArr[i].Users = users
	}

	return allcount,areaArr
}
func (r *areaMemoryRepository) GetByID(id int) (model.Area,bool) {
	//db.Raw("SELECT name, age FROM users WHERE name = ?", 3).Scan(&result)
	var area model.Area
	flag := true
	//datasource.DB.Find(area,id)
	if err := datasource.DB.Table("areas").Where("id = ?",id).Scan(&area).Error; err != nil{
		flag = false
		panic(err)
	}

	return area,flag
}