package model

import "github.com/jinzhu/gorm"

type ClassifyType struct {
	gorm.Model
	Classifyname               string `json:"ClassifyName"`
	Classifyinnercode           string `json:"ClassifyInnerCode"`
	Level              int `json:"Level"`
	Pid                uint  `json:"PID"`//上级分类ID
}
