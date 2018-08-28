package main

import (
	"fmt"
	"math"
)

func add1(x int, y int) int {
	return x + y
}

func add2(x, y int) int { //参数类型相同，可以写成一个
	return x + y
}

func swap(x, y string) (string, string) { //多个返回值
	return y, x
}

func split(sum int) (x, y int) { //返回值命名
	x = sum * 4 / 9
	y = sum - x
	return
}

func test_sqrt() {
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}

	fmt.Println(hypot(3, 4)) //函数也是值
}

func adder() func(int) int { //Go 函数可以是闭包的。闭包是一个函数值，它来自函数体的外部的变量引用。
	sum := 0 //函数可以对这个引用值进行访问和赋值；换句话说这个函数被“绑定”在这个变量上。
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {
	//basic test
	fmt.Println(add1(42, 13))
	fmt.Println(add2(42, 13))

	a, b := swap("hello", "world")
	fmt.Println(a, b)
	fmt.Println(split(17))

	//测试函数也是值
	test_sqrt()

	// 测试闭包
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
}
