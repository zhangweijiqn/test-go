package main

import "fmt"

type Vertex struct {					//使用type定义struct
	X int
	Y int
}

func main() {
	v := Vertex{1, 2}
	fmt.Println(v)
	v.X = 4					//通过.访问元素及修改元素值
	fmt.Println(v.X)

	//通过指针来访问，通过指针间接的访问是透明的。
	p := &v
	p.X = 1e9				//通过.访问元素
	fmt.Println(v)

	//
	var (
		v2 = Vertex{X: 1}  // Y:0 被省略,指定X元素进行赋值
		v3 = Vertex{}      // X:0 和 Y:0
		p1  = &Vertex{1, 2} // 类型为 *Vertex
	)
	fmt.Println(p1, v2, v3)
}
