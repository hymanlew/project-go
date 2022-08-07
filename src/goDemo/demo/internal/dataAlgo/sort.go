package dataAlgo

import (
	"fmt"
	"time"
)

//排序是将一组数据，依指定的顺序进行排列的过程，常见的排序有：
//冒泡排序：内存排序法，通过对序列从前向后（从 0 下标开始），依次比较之后相邻的元素，并进行交换进行排序。
//选择排序：内存排序法，是从数据中按指定的规则选出某个元素（如最小的数），然后和序列中按排序的元素进行交换位置（如第一个数交换），然后如此循环交换来进行排序。
//插入排序：内存排序法，是对列表中的元素以插入的方式找到其适当位置，分成有序和无序两个表，然后依次从无序表中取一个元素放到有序表中进行对比，以实现排序。
//快速排序：内存排序法，是对冒泡排序的改进，是先将数据分成两部分，一部分比另一部分小，然后分别将其进行排序，然后再分割两部分，依此循环进行直至实现有序。

//BubbleSort 1，冒泡排序
func BubbleSort(arr *[15]int) {
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
func SelectSort(arr *[15]int) {

	//标准的数据操作方式
	//(*arr)[1] = 600 其等价于 arr[1] = 600，两者都可以实现数据的修改

	//双层循环，依次筛选出第次循环得到的最大值
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
}

//InsertSort 3，插入排序
func InsertSort(arr *[15]int) {

	//第一次，给第二个元素找到合适的位置并插入
	for i := 1; i < len(arr); i++ {
		insertVal := arr[i]
		insertIndex := i - 1

		//从指定坐标开始比较，即从后往前，从大到小排序比较
		for insertIndex >= 0 && arr[insertIndex] < insertVal {
			arr[insertIndex+1] = arr[insertIndex]
			insertIndex--
		}

		//插入，并判断只有发生数据交换后，才需要插入新数据，否则就是不需要交换，因为本身就是有序的
		if insertIndex+1 != i {
			arr[insertIndex+1] = insertVal
		}
		//fmt.Printf("第%d次插入后 %v\n",i, *arr)
	}
}

//sortFind 4，顺序查找
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

//QuickSort 5，快速排序，二分排序
func QuickSort(left int, right int, array *[15]int) {

	//从小到大排序
	//left 表示数组左边下标，right 表示数组右边的下标，array 表示要排序的数组
	l := left
	r := right

	// middle 表示中轴，支点
	middle := array[(left+right)/2]
	temp := 0

	//for 循环将比 middle 小的数放左边，将比 middle 大的数放到右边
	for l < r {

		//从 middle 的左边找到大于等于 middle 的值
		for array[l] < middle {
			l++
		}

		//从 middle 的右边边找到小于等于 middle 的值
		for array[r] > middle {
			r--
		}

		// 1 >= r 表明本次分解任务完成, break
		if l >= r {
			break
		}

		//将中轴两边的大数和小数，进行交换
		temp = array[l]
		array[l] = array[r]
		array[r] = temp

		//优化，增加等于中轴值的数据的处理，坐标再移动一位
		if array[l] == middle {
			r--
		}
		if array[r] == middle {
			l++
		}
	}

	// 如果  1== r, 再移动下
	if l == r {
		l++
		r--
	}

	// 向左递归
	if left < r {
		QuickSort(left, r, array)
	}
	// 向右递归
	if right > l {
		QuickSort(l, right, array)
	}
}

//BinaryFind 6，二分查找
func BinaryFind(arr *[15]int, leftIndex int, rightIndex int, findVal int) {

	//实现思路：
	//1，先找到中间下标 middle = (leftIndex + rightIndex) /2, 然后让中间下标的值和 findVal 进行比较
	//2，如果 arr[middle] > findVal, 就应该设置查询范围为 leftIndex ---- (middle - 1)
	//3，如果 arr[middle] < findVal, 就应该设置查询范围为 middel+1---- rightIndex
	//4，如果 arr[middle] == findVal，就代表找到了
	//5，将上面的 2、3 步骤的逻辑递归执行，直至找到

	//判断 leftIndex 是否大于 rightIndex
	if leftIndex > rightIndex {
		fmt.Println("找不到")
		return
	}

	//先找到中间的下标
	middle := (leftIndex + rightIndex) / 2

	if (*arr)[middle] > findVal {
		//说明要查找的数，应该在 leftIndex --- middel-1 之间
		BinaryFind(arr, leftIndex, middle-1, findVal)
	} else if (*arr)[middle] < findVal {
		//说明要查找的数，应该在 middel+1 --- rightIndex 之间
		BinaryFind(arr, middle+1, rightIndex, findVal)
	} else {
		fmt.Printf("找到了，下标为%v \n", middle)
	}
}

func sortTest() {
	arr := [15]int{24, 69, 80, 57, 13}

	//冒泡排序：是双层循环，即第个元素都要循环一次所有的数据。所以会很慢
	BubbleSort(&arr)
	fmt.Println("冒泡排序 arr=", arr)

	//选择排序：是从左到右，是第一个元素都要与后面所有的数据比对一次才能得到结果
	start := time.Now().Unix()
	SelectSort(&arr)
	end := time.Now().Unix()
	fmt.Printf("选择排序耗时=%d秒，arr=%v", end-start, arr)

	//插入排序：是从右到左，是后面的元素比对所有左边的数据，但因为左边的数据已经是排过序了，所以不是比对的所有。所以比选择排序要快
	InsertSort(&arr)
	fmt.Printf("插入排序耗时%d秒，arr=%v", end-start, arr)

	//快速排序：是二分排序方法，即在区分大数和小数的集合时，就对两个数据进行了交换（而不是其他算法的一个个比对）。而且之后每次数据
	//处理就是原数据的一半（一次交换两个数据），即是减少了数据比较和交换的次数，所以会很快。
	//但其缺点是，每次数据递归处理时都是开了一个新栈，这就表示要吃一部分内存，即此算法比较占资源。
	QuickSort(0, len(arr)-1, &arr)
	fmt.Printf("快速排序法耗时%d秒, arr=%v", end-start, arr)

	//二分查找
	BinaryFind(&arr, 0, len(arr)-1, -6)
}
