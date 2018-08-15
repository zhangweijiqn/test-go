package main

import "fmt"

var c, python, java bool	//函数外定义变量，默认都会被初始化
var i, j int = 1, 2			//变量定义可以包含初始值，每个变量对应一个。 如果初始化是使用表达式，则可以省略类型；变量从初始值中获得类型。

func main() {
	var i int
	fmt.Println(i, c, python, java)

	//初始化
	k := 3									//在函数中，`:=` 简洁赋值语句在明确类型的地方，可以用于替代 var 定义。
	c, python, java := true, false, "no!"	//函数外的每个语句都必须以关键字开始（`var`、`func`、等等），`:=` 结构不能使用在函数外。
	fmt.Println(i, j, k, c, python, java)

}
