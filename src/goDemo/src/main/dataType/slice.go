package dataType

import "fmt"

func slice()  {
	var intArr = [...]int{1, 22, 33, 66, 99}

	/*
	- 切片是数组的一个引用，因此切片是引用类型，在进行传递时，遵守引用传递的机制。
	- 切片的使用和数组类似，遍历切片、访问切片的元素和求切片长度 len(slice) 都一样。
	- 切片的长度是可变化的，因此切片是一个可以动态变化数组。
	- slice 从底层来说，就是一个 struct 结构体的数据结构

	  type slice struct {
	      ptr  *[2]int
	      len  int
	      cap  int
	  }
	 */
	//声明/定义一个切片
	//1. slice 就是切片名
	//2. intArr[1:3] 表示 slice 引用 intArr 数组起始下标为 1, 最后的下标为3(但是不包含3)的数据
	slice := intArr[1:3]
	fmt.Println("intArr=", intArr)
	fmt.Println("slice 的元素是 =", slice) //22, 33
	fmt.Println("slice 的元素个数 =", len(slice)) //2
	fmt.Println("slice 的容量 =", cap(slice)) //切片的容量是可以动态变化

	fmt.Printf("intArr[1]的地址=%p\n", &intArr[1])
	fmt.Printf("slice[0]的地址=%p slice[0==%v\n", &slice[0], slice[0])
	slice[1] = 34
	fmt.Println()
	fmt.Println("intArr=", intArr)
	fmt.Println("slice 的元素是 =", slice) //  22, 33


	//切片的创建方式
	//1. 直接引用一个已经创建好的数组
	arr := [...]int{1,1,1}
	slice1 := arr[0:3]

	//2. make 创建切片
	//var 切片名 []type = make([]type, len, [cap])
	//- 参数说明: type 是数据类型，len 大小，cap 指定切片容量，可选。如果分配了 cap, 则要求 cap>=len
	//- 如果没有给切片的各个元素赋值，那么就会使用默认值。
	//- 通过 make 方式创建的切片对应的数组是由 make 底层维护，对外不可见，即只能通过 slice 去访问各个元素。
	slice2 := make([]int, 3, 10)
	fmt.Println(arr)
	fmt.Println(slice1)
	fmt.Println(slice2) //输出 len 个数的初始值

	//3. 直接指定具体数组，使用原理类似 make 的方式
	var slice3 []int = []int{1,1,1,1}
	fmt.Println(slice3)

	//4. 切片初始化时，仍然不能越界。范围在 [0-len(arr)] 之间，但是可以动态增长.
	//var slice = arr[0:end]              可以简写 var slice = arr[:end]
	//var slice = arr[start:len(arr)]     可以简写： var slice = arr[start:]
	//var slice = arr[0:len(arr)]         可以简写: var slice = arr[:]
	//cap 是一个内置函数，用于统计切片的容量，即最大可以存放多少个元素。
	//切片定义完后，还不能使用，因为本身是一个空的，需要让其引用到一个数组，或者 make 一个空间供切片来使用
}

func sliceFor()  {
	//使用常规的 for 循环遍历切片
	var arr [5]int = [...]int{10, 20, 30, 40, 50}
	slice := arr[1:4]
	for i := 0; i < len(slice); i++ {
		fmt.Printf("slice[%v]=%v ", i, slice[i])
	}

	//使用for--range 方式遍历切片
	fmt.Println()
	for i, v := range slice {
		fmt.Printf("i=%v v=%v \n", i, v)
	}

	//切片还可以继续切片
	slice2 := slice[1:2] //slice[ 20, 30, 40]    [30]
	slice2[0] = 100  // 因为 arr , slice 和slice2 指向的数据空间是同一个，因此slice2[0]=100，其它的都变化

	fmt.Println("slice2=", slice2)
	fmt.Println("slice=", slice)
	fmt.Println("arr=", arr)
	fmt.Println()

	//用 append 内置函数，可以对切片进行动态追加
	//切片 append 操作的底层原理分析:
	//- 切片 append 操作的本质就是对数组扩容
	//- go 底层会创建一个新的数组 newArr(安装扩容后大小)
	//- 将 slice 原来包含的元素拷贝到新的数组 newArr
	//- slice 重新引用到 newArr
	//- 注意 newArr 是在底层来维护的，外部不可见
	var slice3 []int = []int{100, 200, 300}
	slice3 = append(slice3, 400, 500, 600)
	fmt.Println("slice3", slice3) //100, 200, 300,400, 500, 600

	//通过append将切片slice3追加给slice3
	slice3 = append(slice3, slice3...) // 100, 200, 300,400, 500, 600 100, 200, 300,400, 500, 600
	fmt.Println("slice3", slice3)


	//切片的拷贝操作
	//切片使用copy内置函数完成拷贝，举例说明
	fmt.Println()
	var slice4 []int = []int{1, 2, 3, 4, 5}
	var slice5 = make([]int, 10)
	copy(slice5, slice4)
	fmt.Println("slice4=", slice4)// 1, 2, 3, 4, 5
	fmt.Println("slice5=", slice5) // 1, 2, 3, 4, 5, 0 , 0 ,0,0,0
}