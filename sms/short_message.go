package sms

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/json-iterator/go"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"qiaoyi_back/model"
	"qiaoyi_back/redis"
	"time"
	"unsafe"
)

type SmsConf struct {
	SmsAccount  string
	SmsPassword string
	SmsUrl      string
	SysUrl      string
}

var smsConf SmsConf

func init() {
	b, err := ioutil.ReadFile("config.json")
	if err != nil {
		panic("DB config read err")
	}
	err = jsoniter.Unmarshal(b, &smsConf)
	if err != nil {
		panic(err)
	}
	fmt.Println(smsConf)
}

//发送短信验证码
func SendSms(phone string) bool {
	params := make(map[string]interface{})
	//请登录zz.253.com获取API账号、密码以及短信发送的URL
	params["account"] = smsConf.SmsAccount   //创蓝API账号
	params["password"] = smsConf.SmsPassword //创蓝API密码
	params["phone"] = phone                  //手机号码

	//设置您要发送的内容：其中“【】”中括号为运营商签名符号，多签名内容前置添加提交
	var s = fmt.Sprintf("%06v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(10000))
	var code = "【省社科联】您好，您的验证码是" + s + ",有效时间为3分钟"
	params["msg"] = url.QueryEscape(code)
	params["report"] = "true"
	bytesData, err := json.Marshal(params)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	reader := bytes.NewReader(bytesData)
	request, err := http.NewRequest("POST", smsConf.SmsUrl, reader)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	request.Header.Set("Content-Type", "application/json;charset=UTF-8")
	client := http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	//将验证码和手机存入redis
	redis.RedisSetPhone(phone, s)
	str := (*string)(unsafe.Pointer(&respBytes))
	fmt.Println(*str)
	return true
}

//发送短信
func SendShortMessage(user model.User, name string) {
	params := make(map[string]interface{})
	//请登录zz.253.com获取API账号、密码以及短信发送的URL
	params["account"] = smsConf.SmsAccount   //创蓝API账号
	params["password"] = smsConf.SmsPassword //创蓝API密码
	params["phone"] = user.Phone          //手机号码
	//设置您要发送的内容：其中“【】”中括号为运营商签名符号，多签名内容前置添加提交
	var code = "【市委改革办】您好！课题立项系统有新的用户" + name + "注册，请进行审核！"
	params["msg"] = url.QueryEscape(code)
	params["report"] = "true"
	bytesData, err := json.Marshal(params)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	reader := bytes.NewReader(bytesData)
	request, err := http.NewRequest("POST", smsConf.SmsUrl, reader)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	request.Header.Set("Content-Type", "application/json;charset=UTF-8")
	client := http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	str := (*string)(unsafe.Pointer(&respBytes))
	fmt.Println(*str)
}

//发送用户名和密码给用户
func SendMsg(mobile string, loginName string, pwd string) {
	params := make(map[string]interface{})
	//请登录zz.253.com获取API账号、密码以及短信发送的URL
	params["account"] = smsConf.SmsAccount   //创蓝API账号
	params["password"] = smsConf.SmsPassword //创蓝API密码
	params["phone"] = mobile              //手机号码
	//设置您要发送的内容：其中“【】”中括号为运营商签名符号，多签名内容前置添加提交
	var code = "您好！课题立项系统的登录名为:" + loginName + ",密码为:" + pwd
	params["msg"] = url.QueryEscape(code)
	params["report"] = "true"
	bytesData, err := json.Marshal(params)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	reader := bytes.NewReader(bytesData)
	request, err := http.NewRequest("POST", smsConf.SmsUrl, reader)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	request.Header.Set("Content-Type", "application/json;charset=UTF-8")
	client := http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	log.Println(string(respBytes))
}

//发送通知以及用户名和密码给用户
func SendNoticeMsg(noticeMsg string,mobile string) {
	params := make(map[string]interface{})
	//请登录zz.253.com获取API账号、密码以及短信发送的URL
	params["account"] = smsConf.SmsAccount   //创蓝API账号
	params["password"] = smsConf.SmsPassword //创蓝API密码
	params["phone"] = mobile              //手机号码
	//设置您要发送的内容：其中“【】”中括号为运营商签名符号，多签名内容前置添加提交
	var code = noticeMsg
	params["msg"] = url.QueryEscape(code)
	params["report"] = "true"
	bytesData, err := json.Marshal(params)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	reader := bytes.NewReader(bytesData)
	request, err := http.NewRequest("POST", smsConf.SmsUrl, reader)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	request.Header.Set("Content-Type", "application/json;charset=UTF-8")
	client := http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	log.Println(string(respBytes))
}
