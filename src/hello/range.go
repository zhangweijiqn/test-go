package main

import "fmt"

func test_range1() {
	var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}
	for i, v := range pow { //index,value
		fmt.Printf("2**%d = %d\n", i, v)
	}
}

func test_range2() {
	pow := make([]int, 10)
	for i := range pow { //i: index
		pow[i] = 1 << uint(i) // 1左移位
		fmt.Printf("%d:%d\n", i, pow[i])
	}
	for _, value := range pow { //range遍历会产生index和value，go不允许有不使用的变量，可以用空标识符_
		fmt.Printf("%d\n", value)
	}
}

func main() {

	test_range1() //

	test_range2() //nil, value

}
