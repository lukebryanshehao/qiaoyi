package utils

import (
	"math/rand"
	"strconv"
	"time"
)

//创建数字字符串(不足前面补0)
//CreateNumberStr(4,0) return 0001
//CreateNumberStr(5,1) return 00002
func CreateNumberStr(figures int,startNumber int) (string) {
	numberStr := strconv.Itoa(startNumber+1)
	for(true) {
		if len(strconv.Itoa(startNumber)) < figures {
			numberStr = "0" + numberStr
		}
		if len(numberStr) >= figures {
			break
		}
	}
	return numberStr
}

//创建多位数长度数字字符串
//figures位数
//rang随机数范围
func CreateRandNum(figures int,rang int) (numberStr string) {

	rand.Seed(time.Now().Unix())
	for(true) {
		rnd := rand.Intn(rang)
		numberStr += strconv.Itoa(rnd)

		if len(numberStr) >= figures {
			break
		}
	}
	return
}

