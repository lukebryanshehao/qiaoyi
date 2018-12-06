package model

type Page struct {
	PageSize int `gorm:"-";json:"PageSize"`   // 忽略这个字段
	PageIndex int `gorm:"-";json:"PageIndex"`   // 忽略这个字段
}
