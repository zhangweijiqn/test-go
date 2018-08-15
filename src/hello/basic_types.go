package main

import (
	"fmt"
	"math/cmplx"
	"math"
)

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 1)
)

func main() {
	// 不同类型的变量
	const f = "%T(%v)\n"			// %T:type %v:value
	fmt.Printf(f, ToBe, ToBe)
	fmt.Printf(f, MaxInt, MaxInt)
	fmt.Printf(f, z, z)


	// 变量初始化
	var i int						//变量在定义时没有明确的初始化时会赋值为_零值_。
	var f2 float64
	var b bool
	var s string
	fmt.Printf("%v %v %v %q\n", i, f2, b, s)

	//类型转换
	var x, y int = 3, 4
	var f3 float64 = math.Sqrt(float64(x*x + y*y))
	var z int = int(f3)				//类型转换使用 Type(v)，go中类型必须显示转换
	/* 或者直接使用如下方式：
	i := 42
	f := float64(i)
	u := uint(f)
	*/
	fmt.Println(x, y, z)

	// 自动类型
	v := 42.9 						//当右边包含了未指名类型的数字常量时，新的变量就可能是 int 、 float64 或 `complex128`。 这取决于常量的精度
	fmt.Printf("v is of type %T\n", v)
}


/*
Go 的基本类型有Basic types

	bool

	string

	int  int8  int16  int32  int64
	uint uint8 uint16 uint32 uint64 uintptr

	byte // uint8 的别名

	rune // int32 的别名
		 // 代表一个Unicode码

	float32 float64

	complex64 complex128

初始化零值是：

	数值类型为 `0`，
	布尔类型为 `false`，
	字符串为 `""`（空字符串）。

 */