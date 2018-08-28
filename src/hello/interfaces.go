package main

import (
	"fmt"
	"math"
	"os"
)

type Abser interface {
	AbsI() float64
}

type MyFloat2 float64

func (f MyFloat2) AbsI() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type VertexI struct {
	X, Y float64
}

func (v *VertexI) AbsI() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *VertexI) AbsIn() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	var a Abser
	f := MyFloat2(-math.Sqrt2)
	v := VertexI{3, 4}

	a = f // a MyFloat 实现了 Abser
	fmt.Println(a.AbsI())

	a = &v // a *Vertex 实现了 Abser

	// 下面一行，v 是一个 Vertex（而不是 *Vertex）
	// 所以没有实现 Abser。
	//a = v, 由于 Abs 只定义在 *Vertex（指针类型） 上， 所以 Vertex（值类型） 不满足 `Abser`。

	fmt.Println(a.AbsI()) //只能调用AbsI，调用不了AbsIn
	fmt.Println(v.AbsIn())

	// test implicitly interface
	var w Writer
	// os.Stdout 实现了 Writer
	w = os.Stdout
	fmt.Fprintf(w, "hello, writer\n")
}

type Reader interface {
	Read(b []byte) (n int, err error)
}

type Writer interface {
	Write(b []byte) (n int, err error)
}

/*
类型通过实现那些方法来实现接口。 没有显式声明的必要；所以也就没有关键字“implements“。

隐式接口解藕了实现接口的包和定义接口的包：互不依赖。

因此，也就无需在每一个实现上增加新的接口名称，这样同时也鼓励了明确的接口定义。
*/
type ReadWriter interface {
	Reader
	Writer
}
