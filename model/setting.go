package model

/*
系统设置
 */
type Setting struct {
	Key	string `gorm:"primary_key;type:varchar(100);not null;unique";json:"Key"`
	Value	string `gorm:"type:varchar(500)";json:"Value"`
}

