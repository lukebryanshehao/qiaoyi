package model

import "github.com/jinzhu/gorm"

/*
区域
 */
type Area struct {
	gorm.Model
	Name               string `json:"Name"`
	Innercode           string `json:"InnerCode"`
	Level              int `json:"Level"`
	Pid                uint  `json:"PID"`//上级区域ID
	Users []User
}
