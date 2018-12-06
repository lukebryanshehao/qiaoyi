package datasource

import (
	_ "github.com/Go-SQL-Driver/MySQL"
	"github.com/jinzhu/gorm"
	"github.com/json-iterator/go"
	"io/ioutil"
	"strings"
)

var DB *gorm.DB

type dbconfig struct {
	Port       string `json:"Port"`
	StaticPath       string `json:"StaticPath"`
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
	DB.DB().SetMaxIdleConns(10)
	DB.DB().SetMaxOpenConns(100)
}
