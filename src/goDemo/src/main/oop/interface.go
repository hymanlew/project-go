package oop

import "fmt"

// Usb 定义一个接口
type Usb interface {
	// Start Stop 声明两个没有实现的方法
	Start()
	Stop()
}

type Phone struct{}

// Start 让 Phone 实现 Usb接口的方法，即就是指实现了 Usb接口声明所有方法
func (p Phone) Start() {
	fmt.Println("手机开始工作。。。")
}
func (p Phone) Stop() {
	fmt.Println("手机停止工作。。。")
}

type Camera struct{}

// Start 让 Camera 实现 Usb接口的方法
func (c Camera) Start() {
	fmt.Println("相机开始工作~~~。。。")
}
func (c Camera) Stop() {
	fmt.Println("相机停止工作。。。")
}

type Computer struct{}

// Working 编写一个方法，接收一个Usb接口类型变量
func (c Computer) Working(usb Usb) {

	//通过 usb 接口变量来调用 Start和Stop 方法
	usb.Start()
	usb.Stop()
}

func test1() {
	computer := Computer{}
	phone := Phone{}
	camera := Camera{}

	computer.Working(phone)
	computer.Working(camera)
}
