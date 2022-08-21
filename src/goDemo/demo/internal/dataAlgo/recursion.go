package dataAlgo

import "fmt"

//递归就是函数/方法自己调自己，每次调用时是传入不同的变量。
//递归用于解决的问题：8皇后问题，汉诺塔，阶乘，迷宫，球和篮子的放置问题。

func recursion(n int) {
	if n > 2 {
		n--
		recursion(n)
	}
	fmt.Println("n=", n)
}

func demo() {
	n := 4
	recursion(n)
}

//findWay 递归案例：老鼠找路
func findWay(myMap *[8][9]string, i int, j int) bool {

	//分析出什么情况下，就算找到出口，第 7 行第 8 列
	if myMap[6][7] == "@" {
		return true

		//说明要继续找，i=行，j=列，0 表示这个点是可以探测
	} else if myMap[i][j] == "-" {

		//假设这个点是可以通, 但是需要探测上下左右是否可走通
		//换一个策略 下右上左
		myMap[i][j] = "@"
		if findWay(myMap, i+1, j) { //下
			return true
		} else if findWay(myMap, i, j+1) { //右
			return true
		} else if findWay(myMap, i-1, j) { //上
			return true
		} else if findWay(myMap, i, j-1) { //左
			return true
		} else { //死路不通，设为 3 做标记
			myMap[i][j] = "x"
			return false
		}

		//说明这个点不能探测，为走过的路，或是墙
	} else {
		return false
	}
}

func RecursionDemo() {
	//创建一个二维数组，模拟迷宫，规则如下：
	//1. 如果元素的值为1, 就是墙
	//2. 如果元素的值为0, 是没有走过的点
	//3. 如果元素的值为2, 是一个通路，是走的路线
	//4. 如果元素的值为3，是走过的点，但是走不通

	//迷宫地图 map（8 行 7 列），为了保证是同一张地图，要使用引用
	//i,j，表示对地图的哪个点进行测试
	var myMap [8][9]string

	//设置地图边界，把地图的最上和最下，最左和最右设置为墙
	for i := 0; i < 8; i++ {
		for j := 0; j < 9; j++ {
			if i == 0 || j == 0 || j == 8 || i == 7 {
				myMap[i][j] = "*"
			} else {
				myMap[i][j] = "-"
			}
			if i == 1 && j == 1 {
				myMap[i][j] = "@"
			}
		}
	}

	//设置地图中的障碍物
	myMap[3][1] = "*"
	myMap[3][2] = "*"
	myMap[3][3] = "*"

	//输出地图
	for i := 0; i < 8; i++ {
		for j := 0; j < 9; j++ {
			fmt.Print(myMap[i][j], " ")
		}
		fmt.Println()
	}

	//使用测试
	findWay(&myMap, 2, 1)
	fmt.Println("探测完毕....")

	//输出地图
	for i := 0; i < 8; i++ {
		for j := 0; j < 9; j++ {
			fmt.Print(myMap[i][j], " ")
		}
		fmt.Println()
	}
}
