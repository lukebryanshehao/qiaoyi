/**
*@Author: haoxiongxiao
*@Date: 2018/12/14
*@Description: CREATE GO FILE redis
 */
package test

import (
	"testing"
	"log"
	"fmt"
)

func TestRedisGetPhone(t *testing.T) {
	ii := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	jj := []int{11, 12, 13, 14, 15, 16, 17, 18, 1, 2}

ij:
	for _, i := range ii {

		for _, j := range jj {
			if i == j {
				log.Println(i)
				break ij
			}
		}
	}

	s := "010010001"
	fmt.Println(s[:2])
	fmt.Println(s[:4])

}
