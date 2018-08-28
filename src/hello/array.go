package main
import "fmt"

func test_array(){
	var a [2]string					//类型 [n]T 是一个有 n 个类型为 T 的值的数组。
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)
}

func test_slice(){
	p := []int{2, 3, 5, 7, 11, 13}
	fmt.Println("p ==", p)

	for i := 0; i < len(p); i++ {
		fmt.Printf("p[%d] == %d\n", i, p[i])		//通过[]遍历各个slice
	}
	fmt.Println("p[1:4] ==", p[1:4])					//创建一个新的 slice 值指向相同的数组，注意s[lo:lo]是空的

	// 省略下标代表从 0 开始
	fmt.Println("p[:3] ==", p[:3])

	// 省略上标代表到 len(s) 结束
	fmt.Println("p[4:] ==", p[4:])
}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}

func test_make_slice(){
	a := make([]int, 5)				//slice 由函数 make 创建。这会分配一个零长度的数组并且返回一个 slice 指向这个数组,len(a)=5
	printSlice("a", a)
	b := make([]int, 0, 5)			//为了指定容量，可传递第三个参数到 `make`, len(b)=0, cap(b)=5
	printSlice("b", b)
	c := b[:2]
	printSlice("c", c)
	d := c[2:5]
	printSlice("d", d)

	var z []int						//slice 的零值是 `nil`。
	fmt.Println(z, len(z), cap(z))
	if z == nil {
		fmt.Println("nil!")		//一个 nil 的 slice 的长度和容量是 0。
	}

}

func add_element_to_slice(){
	var a []int						//变长数据是可以进行append操作，并且改变长度的
	printSlice("a", a)

	// append works on nil slices.
	a = append(a, 0)			//append 的结果是一个包含原 slice 所有元素加上新添加的元素的 slice。
									//如果 s 的底层数组太小，而不能容纳所有值时，会分配一个更大的数组。 返回的 slice 会指向这个新分配的数组。
	printSlice("a", a)

	// the slice grows as needed.
	a = append(a, 1)
	printSlice("a", a)

	// we can add more than one element at a time.
	a = append(a, 2, 3, 4)
	printSlice("a", a)
}

func main() {
	test_array()

	test_slice()

	test_make_slice()

	add_element_to_slice()
}
