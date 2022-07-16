package bit

import "fmt"

func calc() {
	// 二进制的原码、反码、补码
	/**
	对于有符号的数据而言：
	- 二进制的最高位是符号位（0 为正数，1 为负数）。
	- 正数的原码、反码、补码都一样。
	- 负数的反码 = 其原码符号位不变，其它位取反（0->1，1->0）。
	- 负数的补码 = 其反码 + 1。
	- 0 的反码补码都是 0。
	- 计算机在运算时，都是以补码的方式来运算的。因为在计算机中只能做加法不能做减法。如下：

	1 + 1 = 1 + 1
	1 - 1 = 1 + (-1)
	如此在计算机中，正好可以将正数与负数做计算，因此最后计算机就采用了补码进行计算
	*/

	//位运算的演示（与、或、异或（一个为 0 一个为 1 则为 1 否则为 0））
	//或：两个有一个为 1 则为 1，否则为 0
	//异或：一个为 0 一个为 1 则为 1 否则为 0
	/**
	2 的补码  0000 0010
	3 的补码  0000 0011
	2&3      0000 0010  => 2

	2 的补码  0000 0010
	3 的补码  0000 0011
	2|3      0000 0011 => 3

	2 的补码  0000 0010
	3 的补码  0000 0011
	2^3      0000 0001 =>1

	-2^2
	-2 的原码  1000 0010 =》反码 1111 1101 => 补码  1111 1110
	-2 的补码  1111 1110
	2  的补码  0000 0010
	计算结果   1111 1100，表明它是个负数，并且结果都是补码（因为计算机是按数据的补码计算的），因此需要将补码转换为原码。

	补码 1111 1100 -1 ==》反码 1111 1011 ==》取反得原码 1000 0100 ==》 -4
	*/
	fmt.Println(2 & 3)  // 2
	fmt.Println(2 | 3)  // 3
	fmt.Println(2 ^ 3)  // 3
	fmt.Println(-2 ^ 2) //-4

	//移位运算
	//- 右移运算 >>：低位溢出，符号位不变，并用符号位补溢出的高位。
	//- 左移运算 <<：符号位不变，低位补 0。
	a := 1 >> 2 //0
	c := 1 << 2 //4
	fmt.Println("a=", a, "c=", c)
}
