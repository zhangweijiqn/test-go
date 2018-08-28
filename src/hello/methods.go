package main

import (
	"fmt"
	"math"
)

type VertexNew struct {
	X, Y float64
}

/*
Go 没有类。然而，仍然可以在结构体类型上定义方法。

方法接收者 出现在 func 关键字和方法名之间的参数中 （参数既指明了接收者，也指明了参数类型）

方法可以与命名类型或命名类型的指针关联。

需要使用指针接收者：避免在每个方法调用中拷贝值，方法可以修改接收者指向的值 （直接传递结构体变量的情况下，参数为拷贝对象，并且不修改原值）。
*/

func (v *VertexNew) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

type MyFloat float64

/*
你可以对包中的 任意 类型定义任意方法，而不仅仅是针对结构体。

但是，不能对来自其他包的类型或基础类型定义方法。
*/
func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func main() {
	// test method for struct
	v := &VertexNew{3, 4}
	fmt.Println(v.Abs())

	// test method for variable
	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())

}
