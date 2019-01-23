package datasource

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/json-iterator/go"
	"io/ioutil"
	"strings"
	"time"
)

var DB *gorm.DB

type dbconfig struct {
	Port       string `json:"Port"`
	StaticPath       string `json:"StaticPath"`
	TokenKey       string `json:"TokenKey"`
	UserName       string `json:"UserName"`
	Password string `json:"Password"`
	Ip string `json:"Ip"`
	DBPort string `json:"DBPort"`
	DBName string `json:"DBName"`
}

var DBconfig = &dbconfig{}

func checkerror(err error)  {
	if err != nil {
		panic(err)
	}
}

func init() {

	b, err := ioutil.ReadFile("config.json")
	if err != nil {
		panic("DB config read err")
	}
	err = jsoniter.Unmarshal(b, DBconfig)
	checkerror(err)

	path := strings.Join([]string{DBconfig.UserName, ":", DBconfig.Password, "@tcp(", DBconfig.Ip, ":", DBconfig.DBPort, ")/", DBconfig.DBName, "?charset=utf8&parseTime=true"}, "")
	DB, err = gorm.Open("mysql", path)
	checkerror(err)
	DB.SingularTable(true)
	DB.DB().SetConnMaxLifetime(1 * time.Second)
	DB.DB().SetMaxIdleConns(20)   //最大打开的连接数
	DB.DB().SetMaxOpenConns(2000) //设置最大闲置个数
	DB.SingularTable(true)	//表生成结尾不带s
	// 启用Logger，显示详细日志
	DB.LogMode(true)
}
