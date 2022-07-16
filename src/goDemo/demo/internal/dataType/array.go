package dataType

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

//在 GO 中，数组是值类型，其变量是直接指向真实数据空间的。而不是指向一个地址
//- 数组的地址可以通过数组名来获取 &intArr
//- 数组的第一个元素的地址，就是数组的首地址
//- 数组的各个元素的地址间隔是依据数组的类型决定，比如 int64 -> 8 int32->4..
//- 其特点是，大小确定，内存连续，可以随机访问
func arry() {

	//思路分析：定义六个变量，分别表示六只鸡的，然后求出和，然后求出平均值。
	hen1 := 3.0
	hen2 := 5.0
	hen3 := 1.0
	hen4 := 3.4
	hen5 := 2.0
	hen6 := 50.0
	totalWeight := hen1 + hen2 + hen3 + hen4 + hen5 + hen6

	//将 totalWeight/6 四舍五入保留到小数点 2 并返回值
	avgWeight := fmt.Sprintf("%.2f", totalWeight/6)
	fmt.Printf("totalWeight=%v avgWeight=%v\n", totalWeight, avgWeight)

	//使用数组的方式来解决问题
	var hens [7]float64
	hens[0] = 3.0
	hens[1] = 5.0
	hens[2] = 1.0
	hens[3] = 3.4
	hens[4] = 2.0
	hens[5] = 50.0
	hens[6] = 150.0

	totalWeight2 := 0.0
	for i := 0; i < len(hens); i++ {
		totalWeight2 += hens[i]
	}
	//求出平均体重
	avgWeight2 := fmt.Sprintf("%.2f", totalWeight2/float64(len(hens)))
	fmt.Printf("totalWeight2=%v avgWeight2=%v", totalWeight2, avgWeight2)

	//数组的声明方式
	var arr = [3]int{1, 2, 3}
	fmt.Println(len(arr))
	var arr2 [3]int = [3]int{1, 2, 3}
	fmt.Println(len(arr2))
	var arr3 = [...]string{"a", "a", "a"}
	fmt.Println(arr3)
	arr4 := [...]string{"a", "a", "a"}
	fmt.Println(arr4)
}

func arrFor() {
	arr := [...]string{"a", "a", "a"}

	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}

	for i, v := range arr {
		fmt.Println(strconv.Itoa(i) + " == " + v)
	}

	for _, v := range arr {
		fmt.Println(v)
	}

	arrData(&arr)
	fmt.Println(arr)
}

//因为数组是值类型，所以修改其值，就需要使用指针修改
//并且长度是数组类型的一部分，在传递函数参数时，就需要声明数组的长度
func arrData(arr *[3]string) {
	(*arr)[1] = "c"
	(*arr)[0] = "d"
}

func work() {
	//创建一个 byte 类型的 26 个元素的数组，分别放置'A'-'Z‘。
	//使用for循环访问所有元素并打印出来。提示：字符数据运算 'A'+1 -> 'B'

	//思路
	//使用for循环，利用字符可以进行运算的特点来赋值 'A'+1 -> 'B'
	var myChars [26]byte
	for i := 0; i < 26; i++ {
		myChars[i] = 'A' + byte(i) // 注意需要将 i => byte
	}
	for i := 0; i < 26; i++ {
		fmt.Printf("%c ", myChars[i])
	}

	//请求出一个数组的最大值，并得到对应的下标
	fmt.Println()
	var intArr = [...]int{1, -1, 9, 90, 11, 9000}
	maxVal := intArr[0]
	maxValIndex := 0

	for i := 1; i < len(intArr); i++ {
		if maxVal < intArr[i] {
			maxVal = intArr[i]
			maxValIndex = i
		}
	}
	fmt.Printf("maxVal=%v maxValIndex=%v\n\n", maxVal, maxValIndex)

	//请求出一个数组的和和平均值。for-range
	var intArr2 = [...]int{1, -1, 9, 90, 12}
	sum := 0
	for _, val := range intArr2 {
		sum += val
	}
	fmt.Printf("sum=%v 平均值=%.2f\n\n", sum, float64(sum)/float64(len(intArr2)))

	//要求：随机生成五个数，并将其反转打印
	//反转打印, 交换的次数是  len / 2, 倒数第一个和第一个元素交换, 倒数第2个和第2个元素交换
	var intArr3 [5]int
	//为了每次生成的随机数不一样，需要给一个seed值
	//rand.Seed(time.Millisecond.Milliseconds())
	rand.Seed(time.Now().UnixNano())
	lengh := len(intArr3)
	for i := 0; i < lengh; i++ {
		intArr3[i] = rand.Intn(100) //  0<=n<100
	}
	fmt.Println("交换前~=", intArr3)

	temp := 0 //做一个临时变量
	for i := 0; i < lengh/2; i++ {
		temp = intArr3[lengh-1-i]
		intArr3[lengh-1-i] = intArr3[i]
		intArr3[i] = temp
	}
	fmt.Println("交换后~=", intArr3)
}

//二维数组
func doubleArray() {
	//二维数组的演示案例
	/*
		0 0 0 0 0 0
		0 0 1 0 0 0
		0 2 0 3 0 0
		0 0 0 0 0 0
	*/
	var arr [4][6]int
	arr[1][2] = 1
	arr[2][1] = 2
	arr[2][3] = 3

	//遍历二维数组，按照要求输出图形
	for i := 0; i < 4; i++ {
		for j := 0; j < 6; j++ {
			fmt.Print(arr[i][j], " ")
		}
		fmt.Println()
	}
	fmt.Println()

	var arr2 [2][3]int //以这个为例来分析arr2在内存的布局!!
	arr2[1][1] = 10
	fmt.Println(arr2)

	fmt.Printf("arr2[0]的地址%p\n", &arr2[0])
	fmt.Printf("arr2[1]的地址%p\n", &arr2[1])
	fmt.Printf("arr2[0][0]的地址%p\n", &arr2[0][0])
	fmt.Printf("arr2[1][0]的地址%p\n", &arr2[1][0])
	fmt.Println()

	arr3 := [2][3]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println("arr3=", arr3)
}

func doubleArrayFor() {
	var arr3 = [2][3]int{{1, 2, 3}, {4, 5, 6}}

	//for循环来遍历
	for i := 0; i < len(arr3); i++ {
		for j := 0; j < len(arr3[i]); j++ {
			fmt.Printf("%v\t", arr3[i][j])
		}
		fmt.Println()
	}

	//for-range来遍历二维数组
	for i, v := range arr3 {
		for j, v2 := range v {
			fmt.Printf("arr3[%v][%v]=%v \t", i, j, v2)
		}
		fmt.Println()
	}
}

func work3() {
	/*
		定义二维数组，用于保存三个班，每个班五名同学成绩，
		并求出每个班级平均分、以及所有班级平均分
	*/
	var scores [3][5]float64

	//2.循环的输入成绩
	for i := 0; i < len(scores); i++ {
		for j := 0; j < len(scores[i]); j++ {
			fmt.Printf("请输入第 %d 班的第 %d 个学生的成绩\n", i+1, j+1)
			fmt.Scanln(&scores[i][j])
		}
	}
	fmt.Println(scores)

	//3.遍历输出成绩后的二维数组，统计平均分
	// 定义一个变量，用于累计所有班级的总分
	totalSum := 0.0

	for i := 0; i < len(scores); i++ {
		//定义一个变量，用于累计各个班级的总分
		sum := 0.0
		for j := 0; j < len(scores[i]); j++ {
			sum += scores[i][j]
		}
		totalSum += sum

		fmt.Printf("第%d班级的总分为%v , 平均分%v\n", i+1, sum, sum/float64(len(scores[i])))
	}
	fmt.Printf("所有班级的总分为%v , 所有班级平均分%v\n", totalSum, totalSum/15)
}
