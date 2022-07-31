package dataAlgo

import "fmt"

//稀疏数组（五子棋盘案例）： 当一个数组中大部分元素为 0，或者为同一值的数组时，可以使用稀疏数组来保存该数据。
//处理方法：
//1，记录数组一共有几行几列，有多少个不同的值。
//2，把具有不同值的元素的行列及值记录在一个小规模的数组中，从而缩小程序的规模。

//原数组数据：
//0	 0 	 0 	 22  0 	 0 	 15
//0	 11	 0	 0	 0	 17	 0
//0	 0	 0	 6	 0	 0	 0

//转换后的数据，行列都是从 0  开始计算：
//行，列，值
//0，3，22
//0，6，15
//1，1，11
//1，5，17
//2，3，6

type ValNode struct {
	row int
	col int
	val int
}

//需求：五子棋盘，有存盘退出和读盘接续的功能
func array1() {

	//原始传统的记录方法：
	//1. 先创建一个原始数组
	var chessMap [11][11]int
	chessMap[1][2] = 1 //黑子
	chessMap[2][3] = 2 //白子

	//2. 输出看看原始的数组
	for _, v := range chessMap {
		for _, v2 := range v {
			fmt.Printf("%d\t", v2)
		}
		fmt.Println()
	}

	//转成稀疏数组。想-> 算法
	//(1). 遍历 chessMap, 如果我们发现有一个元素的值不为0，创建一个node结构体
	//(2). 将其放入到对应的切片即可
	var sparseArr []ValNode

	//一个标准的稀疏数组，应该还有一个数据的原始二维数组（行和列，默认值)，即用于记录原始的棋盘基本信息
	valNode := ValNode{
		row: 11,
		col: 11,
		val: 0,
	}
	sparseArr = append(sparseArr, valNode)

	//遍历原始数组，生成对应的稀疏数组
	for i, v := range chessMap {
		for j, v := range v {
			if v != 0 {
				valNode := ValNode{
					row: i,
					col: j,
					val: v,
				}
				sparseArr = append(sparseArr, valNode)
			}
		}
	}

	//输出稀疏数组
	fmt.Println("当前的稀疏数组是:::::")
	for i, valNode := range sparseArr {
		fmt.Printf("%d: %d %d %d\n", i, valNode.row, valNode.col, valNode.val)
	}

	//将这个稀疏数组，存盘 d:/chessmap.data，写入文件
	//恢复原始的数组，打开这个 d:/chessmap.data 文件，然后恢复原始数组

	//这里使用稀疏数组恢复
	//先创建一个原始数组，也可以读文件第一行的数据，用于生成整个棋盘
	var chessMap2 [11][11]int

	//遍历 sparseArr [遍历文件每一行]
	for i, valNode := range sparseArr {
		if i != 0 { //跳过第一行记录值
			chessMap2[valNode.row][valNode.col] = valNode.val
		}
	}

	fmt.Println("恢复后的原始数据......")
	for _, v := range chessMap2 {
		for _, v2 := range v {
			fmt.Printf("%d\t", v2)
		}
		fmt.Println()
	}
}
