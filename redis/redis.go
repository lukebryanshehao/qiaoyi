package redis

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"github.com/json-iterator/go"
	"io/ioutil"
)
type RedisConf struct {
	RedisHost string
	RedisPort string
	RedisPassword string
}
var conn redis.Conn
func init(){
	v :=RedisConf{}
	b, err := ioutil.ReadFile("config.json")
	if err != nil {
		panic("DB config read err")
	}
	//can only unmarshal into pointer
	err = jsoniter.Unmarshal(b, &v)
	c,err:=redis.Dial("tcp",v.RedisHost+":"+v.RedisPort)
	if err != nil {
		fmt.Println("redis连接失败:::::::", err)
		return
	}
	c.Do("AUTH",v.RedisPassword)
	conn = c
}
//存储手机号验证码
func RedisSetPhone(phone string,code string)bool{
	_,err:=conn.Do("set",phone,code,"EX","180")
	if err!=nil {
		fmt.Println(err)
		return false
	}
	return true
}
//根据手机号获取验证码
func RedisGetPhone(phone string,code string)bool{
	v,err:=redis.String(conn.Do("get",phone))
	if err!=nil {
		fmt.Println(err)
		return false
	}
	if v==code {
		return true
	}else{
		return false
	}
	return true
}