package funct

import "fmt"

// BubbleSort 冒泡排序
func BubbleSort(arr *[5]int) {
	fmt.Println("排序前arr=", *arr)

	//临时变量(用于做交换)
	temp := 0
	for i :=0; i < len(*arr) - 1; i++ {
		for j := 0; j < len(*arr) - 1 - i; j++ {
			if (*arr)[j] > (*arr)[j + 1] {
				//交换
				temp = (*arr)[j]
				(*arr)[j] = (*arr)[j + 1]
				(*arr)[j + 1] = temp
			}
		}
	}
	fmt.Println("排序后arr=", (*arr))
}

func sortTest() {
	//冒泡排序
	arr := [5]int{24,69,80,57,13}
	BubbleSort(&arr)
	fmt.Println("main arr=", arr)
}

// 顺序查找
func sortFind()  {
	//有一个数列：白眉鹰王、金毛狮王、紫衫龙王、青翼蝠王
	//猜数游戏：从键盘中任意输入一个名称，判断数列中是否包含此名称【顺序查找】
	names := [4]string{"白眉鹰王", "金毛狮王", "紫衫龙王", "青翼蝠王"}
	var heroName = ""
	fmt.Println("请输入要查找的人名...")
	fmt.Scanln(&heroName)

	//顺序查找:第一种方式
	// for i := 0; i < len(names); i++ {
	// 	if heroName == names[i] {
	// 		fmt.Printf("找到%v , 下标%v \n", heroName, i)
	// 		break
	// 	} else if i == (len(names) - 1) {
	// 		fmt.Printf("没有找到%v \n", heroName)
	// 	}
	// }

	//顺序查找:第2种方式
	index := -1
	for i := 0; i < len(names); i++ {
		if heroName == names[i] {
			index = i //将找到的值对应的下标赋给 index
			break
		}
	}
	if index != -1 {
		fmt.Printf("找到%v , 下标%v \n", heroName, index)
	} else {
		fmt.Println("没有找到", heroName)
	}
}

// BinaryFind 二分查找
/*
1. arr是一个有序数组，并且是从小到大排序
2.  先找到中间的下标 middle = (leftIndex + rightIndex) /2, 然后让中间下标的值和 findVal 进行比较
2.1 如果 arr[middle] > findVal,  就应该向 leftIndex ---- (middle - 1)
2.2 如果 arr[middle] < findVal,  就应该向 middel+1---- rightIndex
2.3 如果 arr[middle] == findVal， 就找到
2.4 上面的 2.1、2.2、2.3 的逻辑会递归执行
if  leftIndex > rightIndex {
   // 找不到..
   return ..
}
*/
func BinaryFind(arr *[6]int, leftIndex int, rightIndex int, findVal int) {

	//判断 leftIndex 是否大于 rightIndex
	if leftIndex > rightIndex {
		fmt.Println("找不到")
		return
	}

	//先找到中间的下标
	middle := (leftIndex + rightIndex) / 2

	if (*arr)[middle] > findVal {
		//说明我们要查找的数，应该在 leftIndex --- middel-1
		BinaryFind(arr, leftIndex, middle - 1, findVal)
	} else if (*arr)[middle] < findVal {
		//说明我们要查找的数，应该在 middel+1 --- rightIndex
		BinaryFind(arr, middle + 1, rightIndex, findVal)
	} else {
		//找到了
		fmt.Printf("找到了，下标为%v \n", middle)
	}
}

func sort2Test() {
	arr := [6]int{1,8, 10, 89, 1000, 1234}
	BinaryFind(&arr, 0, len(arr) - 1, -6)
}