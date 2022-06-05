package process

import "fmt"

func work()  {
	//打印1~100之间所有是9的倍数的整数的个数及总和
	var max uint64 = 100
	var count uint64 = 0
	var sum uint64 = 0
	var i uint64 = 1
	for ; i <= max; i++ {
		if i % 9 == 0 {
			count++
			sum += i
		}
	}
	fmt.Printf("count=%v sum=%v\n", count, sum)

	fmt.Println("--------------------------------")

	//完成下面的表达式输出 ，6是可变的
	var n int = 60
	for i := 0; i <= n; i++ {
		fmt.Printf("%v + %v = %v \n", i, n-i, n)
	}
}

func star()  {
	//使用 for 循环编写一个程序，可以接收一个整数, 表示层数，并打印出金字塔
	//打印空心金字塔

	//  *
	// * *
	//*****

	var totalLevel int = 20
	for i := 1; i <= totalLevel; i++ {

		//在打印 * 前先打印空格
		for k := 1; k <= totalLevel - i; k++ {
			fmt.Print(" ")
		}

		//j 表示每层打印多少*
		for j :=1; j <= 2 * i - 1; j++ {
			if j == 1 || j == 2 * i - 1 || i == totalLevel {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}


	//打印出九九乘法表
	var num int = 9
	for i := 1; i <= num; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%v * %v = %v \t", j, i, j * i)
		}
		fmt.Println()
	}
}

func star2()  {
	//打印实心金字塔

	//	*
	// ***
	//*****
	// 2 * level -1

	var num = 1
	for num > 0 {
		num--
		var level = 1
		fmt.Printf("输入金字塔层级。。。\n")
		fmt.Scanf("%d\n", &level)
		fmt.Printf("\n打印金字塔。。。\n")

		var total = 2 * level - 1
		var star = 1
		var space = 0
		for i:=1; i<=level; i++ {

			star = i * 2 -1
			space = (total - star) / 2
			for start:=0; start<space; start++ {
				fmt.Print(" ")
			}

			for start:=0; start<star; start++ {
				fmt.Print("*")
			}

			for start:=0; start<space; start++ {
				fmt.Print(" ")
			}
			fmt.Println()
		}

		//当使用 Scanf 方法时，要特别注意要过滤掉换行符。否则下次使用该方法读取数据时，会直接读取到上次缓存的换行符。
		fmt.Printf("\n\n是否继续。。。（1 继续 0 结束）\n")
		var check bool
		fmt.Scanf("%t\n", &check)
		fmt.Printf("输入结果 = %v \n", check)
		if check {
			num++
		}else {
			break
		}
		fmt.Printf("\n")
	}
	fmt.Printf("\n\n游戏结束。。。")
}
