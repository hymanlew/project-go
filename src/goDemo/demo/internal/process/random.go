package process

import (
	"fmt"
	"math/rand"
	"time"
)

func random() {
	//使用 rand 设置一个种子，生成一个随机数
	//time.Now().Unix(): 返回一个从 1970:01:01 0时0分0秒 UTC 时间到现在的秒数
	//rand.Seed(time.Now().Unix()): 设置一个随机数的种子

	//随机的生成 1-100 整数
	//n := rand.Intn(100) + 1
	//fmt.Println(n)

	var count int = 0
	for {
		//设置一个更加随机数的种子，使用纳秒
		rand.Seed(time.Now().UnixNano())
		n := rand.Intn(100) + 1
		fmt.Println("n=", n)
		count++
		if n == 99 {
			break //表示跳出for循环
		}
	}
	fmt.Println("生成 99 一共使用了 ", count)

	//这里演示一下指定标签的形式来使用 break
lable2:
	for i := 0; i < 4; i++ {
		//lable1: // 设置一个标签
		for j := 0; j < 10; j++ {
			if j == 2 {
				//break // break 默认会跳出最近的for循环
				//break lable1
				break lable2 // j=0 j=1
			}
			fmt.Println("j=", j)
		}
	}

}
