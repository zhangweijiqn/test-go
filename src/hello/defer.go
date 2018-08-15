package main

import "fmt"


func double(x int) (result int) {
	defer func() { fmt.Printf("double(%d) = %d\n", x, result) }()
	//defer语句中的函数会在return语句更新返回值变量后再执行，又因为在函数中定义的匿名函数可以访问该函数包括返回值变量在内的所有变量
	//所以，对匿名函数采用defer机制，可以使其观察函数的返回值。
	x++
	return x + x
}


func main() {

	defer fmt.Println("world")			//defer 语句会延迟函数的执行直到上层函数返回。
	fmt.Println("hello")

	fmt.Println("counting")
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)				//延迟调用的参数会立刻生成，但是在上层函数返回前函数都不会被调用。
											// 延迟到main函数返回，后进先出（栈）的顺序
	}
	fmt.Println("done")

	double(4)
}
