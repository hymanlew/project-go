package inoutput

import "fmt"

func main() {

}
func Scan() {
	//要求从控制台接收用户信息，【姓名，年龄，薪水, 是否通过考试 】

	//方式1 fmt.Scanln, 以回车换行为结束符
	var name string
	var age byte
	var sal float32
	var isPass bool
	fmt.Println("请输入姓名 ")
	// 该函数需要接收一个变量的内存地址接收输入的值，否则就接收不到了。
	// 并且如果不是传入变量的内存地址，则此函数也不会阻塞等待输入信息，而是直接跳过去。
	//fmt.Scanln(name)
	fmt.Scanln(&name)
	fmt.Println("请输入年龄 ")
	fmt.Scanln(&age)
	fmt.Println("请输入薪水 ")
	fmt.Scanln(&sal)
	fmt.Println("请输入是否通过考试 ")
	fmt.Scanln(&isPass)
	fmt.Printf("名字是 %v \n 年龄是 %v \n 薪水是 %v \n 是否通过考试 %v \n", name, age, sal, isPass)

	//方式2 fmt.Scanf, 按指定的格式输入
	fmt.Println("请输入你的姓名，年龄，薪水, 是否通过考试， 使用空格隔开")
	fmt.Scanf("%s %d %f %t", &name, &age, &sal, &isPass)
	fmt.Printf("名字是 %v \n年龄是 %v \n 薪水是 %v \n 是否通过考试 %v \n", name, age, sal, isPass)
}
