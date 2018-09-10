package main

import (
	"fmt"
	"sort"
)

func testBasic1() {
	// 对于 int 、 float64 和 string 数组或是分片的排序， go 分别提供了 sort.Ints() 、 sort.Float64s() 和 sort.Strings() 函数， 默认都是从小到大排序。
	intList := []int{2, 4, 3, 5, 7, 6, 9, 8, 1, 0}
	float8List := []float64{4.2, 5.9, 12.3, 10.0, 50.4, 99.9, 31.4, 27.81828, 3.14}
	stringList := []string{"a", "c", "b", "d", "f", "i", "z", "x", "w", "y"}

	sort.Ints(intList)
	sort.Float64s(float8List)
	sort.Strings(stringList)

	fmt.Printf("%v\n%v\n%v\n", intList, float8List, stringList)
}

func testBasic2() {
	// go 的 sort 包可以使用 sort.Reverse(slice) 来调换 slice.Interface.Less
	// 也就是比较函数，所以， int 、 float64 和 string 的逆序排序函数可以这么写
	intList := []int{2, 4, 3, 5, 7, 6, 9, 8, 1, 0}
	float8List := []float64{4.2, 5.9, 12.3, 10.0, 50.4, 99.9, 31.4, 27.81828, 3.14}
	stringList := []string{"a", "c", "b", "d", "f", "i", "z", "x", "w", "y"}

	sort.Sort(sort.Reverse(sort.IntSlice(intList)))
	sort.Sort(sort.Reverse(sort.Float64Slice(float8List)))
	sort.Sort(sort.Reverse(sort.StringSlice(stringList)))

	fmt.Printf("%v\n%v\n%v\n", intList, float8List, stringList)
}

func testBasic3() {
	ints := []int{5, 2, 6, 3, 1, 4}

	sort.Ints(ints) // 特殊排序函数，升序
	fmt.Println("after sort by Ints:\t", ints)

	doubles := []float64{2.3, 3.2, 6.7, 10.9, 5.4, 1.8}

	sort.Float64s(doubles)
	fmt.Println("after sort by Float64s:\t", doubles) // [1.8 2.3 3.2 5.4 6.7 10.9]

	strings := []string{"hello", "good", "students", "morning", "people", "world"}
	sort.Strings(strings)
	fmt.Println("after sort by Strings:\t", strings) // [good hello mornig people students world]

	ipos := sort.SearchInts(ints, -1) // int 搜索
	fmt.Printf("pos of 5 is %d th\n", ipos)

	dpos := sort.SearchFloat64s(doubles, 20.1) // float64 搜索
	fmt.Printf("pos of 5.0 is %d th\n", dpos)

	fmt.Printf("doubles is asc ? %v\n", sort.Float64sAreSorted(doubles))

	doubles = []float64{3.5, 4.2, 8.9, 100.98, 20.14, 79.32}
	// sort.Sort(sort.Float64Slice(doubles))    // float64 排序方法 2
	// fmt.Println("after sort by Sort:\t", doubles)    // [3.5 4.2 8.9 20.14 79.32 100.98]
	(sort.Float64Slice(doubles)).Sort()           // float64 排序方法 3
	fmt.Println("after sort by Sort:\t", doubles) // [3.5 4.2 8.9 20.14 79.32 100.98]

	sort.Sort(Reverse{sort.Float64Slice(doubles)})         // float64 逆序排序
	fmt.Println("after sort by Reversed Sort:\t", doubles) // [100.98 79.32 20.14 8.9 4.2 3.5]
}

func main() {
	testBasic1()
	testBasic2()
	testBasic3()
}

// 自定义的 Reverse 类型
type Reverse struct {
	sort.Interface // 这样，Reverse可以接纳任何实现了sort.Interface的对象
}

// Reverse 只是将其中的 Inferface.Less 的顺序对调了一下
func (r Reverse) Less(i, j int) bool {
	return r.Interface.Less(j, i)
}
