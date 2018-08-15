package main
import "fmt"

func add1(x int, y int) int {
	return x + y
}

func add2(x, y int) int {						//参数类型相同，可以写成一个
	return x + y
}

func swap(x, y string) (string, string) {		//多个返回值
	return y, x
}

func split(sum int) (x, y int) {				//返回值命名
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	fmt.Println(add1(42, 13))
	fmt.Println(add2(42, 13))

	a, b := swap("hello", "world")
	fmt.Println(a, b)

	fmt.Println(split(17))
}
