package dataAlgo

import (
	"fmt"
	"math/rand"
	"time"
)

//排序是将一组数据，依指定的顺序进行排列的过程，常见的排序有：
//冒泡排序：内存排序法，通过对序列从前向后（从 0 下标开始），依次比较之后相邻的元素，并进行交换进行排序。
//选择排序：内存排序法，是从数据中按指定的规则选出某个元素（如最小的数），然后和序列中按排序的元素进行交换位置（如第一个数交换），然后如此循环交换来进行排序。
//，插入排序，快速排序

//BubbleSort 1，冒泡排序
func BubbleSort(arr *[5]int) {
	fmt.Println("排序前arr=", *arr)

	//临时变量(用于做交换)
	temp := 0
	for i := 0; i < len(*arr)-1; i++ {
		for j := 0; j < len(*arr)-1-i; j++ {
			if (*arr)[j] > (*arr)[j+1] {
				//交换
				temp = (*arr)[j]
				(*arr)[j] = (*arr)[j+1]
				(*arr)[j+1] = temp
			}
		}
	}
	fmt.Println("排序后arr=", (*arr))
}

//SelectSort 2，选择排序
func SelectSort(arr *[80000]int) {

	//标准的访问方式
	//(*arr)[1] = 600 等价于 arr[1] = 900
	//arr[1] = 900
	//1. 先完成将第一个最大值和 arr[0] => 先易后难

	for j := 0; j < len(arr)-1; j++ {

		//假设 arr[0] 是最大值
		max := arr[j]
		maxIndex := j

		//遍历后面的元素，并与最大值进行比较
		for i := j + 1; i < len(arr); i++ {
			if max < arr[i] {
				max = arr[i]
				maxIndex = i
			}
		}

		//如果最大值下标发生了变化，则进行交换
		if maxIndex != j {
			arr[j], arr[maxIndex] = arr[maxIndex], arr[j]
		}

		//fmt.Printf("第%d次 %v\n  ", j+1 ,*arr)
	}

	/*
		max = arr[1]
		maxIndex = 1
		//2. 遍历后面 2---[len(arr) -1] 比较
		for i := 1 + 1; i < len(arr); i++ {
			if max < arr[i] { //找到真正的最大值
				max = arr[i]
				maxIndex = i
			}
		}
		//交换
		if maxIndex != 1 {
			arr[1], arr[maxIndex] = arr[maxIndex], arr[1]
		}

		fmt.Println("第2次 ", *arr)



		max = arr[2]
		maxIndex = 2
		//2. 遍历后面 3---[len(arr) -1] 比较
		for i := 2 + 1; i < len(arr); i++ {
			if max < arr[i] { //找到真正的最大值
				max = arr[i]
				maxIndex = i
			}
		}
		//交换
		if maxIndex != 2 {
			arr[2], arr[maxIndex] = arr[maxIndex], arr[2]
		}

		fmt.Println("第3次 ", *arr)

		max = arr[3]
		maxIndex = 3
		//2. 遍历后面 4---[len(arr) -1] 比较
		for i := 3 + 1; i < len(arr); i++ {
			if max < arr[i] { //找到真正的最大值
				max = arr[i]
				maxIndex = i
			}
		}
		//交换
		if maxIndex != 3 {
			arr[3], arr[maxIndex] = arr[maxIndex], arr[3]
		}

		fmt.Println("第4次 ", *arr)*/
}

//sortFind 6，顺序查找
func sortFind() {
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

func sortTest() {
	//冒泡排序
	arr := [5]int{24, 69, 80, 57, 13}
	BubbleSort(&arr)
	fmt.Println("main arr=", arr)
}

func main() {
	//定义一个数组 , 从大到小
	//arr := [6]int{10, 34, 19, 100, 80, 789}

	var arr [80000]int
	for i := 0; i < 80000; i++ {
		arr[i] = rand.Intn(900000)
	}

	//fmt.Println(arr)
	start := time.Now().Unix()
	SelectSort(&arr)
	end := time.Now().Unix()
	fmt.Printf("选择排序耗时=%d秒", end-start)
	fmt.Println("main函数")
	//fmt.Println(arr)
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
		BinaryFind(arr, leftIndex, middle-1, findVal)
	} else if (*arr)[middle] < findVal {
		//说明我们要查找的数，应该在 middel+1 --- rightIndex
		BinaryFind(arr, middle+1, rightIndex, findVal)
	} else {
		//找到了
		fmt.Printf("找到了，下标为%v \n", middle)
	}
}

func sort2Test() {
	arr := [6]int{1, 8, 10, 89, 1000, 1234}
	BinaryFind(&arr, 0, len(arr)-1, -6)
}
