package main
import "fmt"

func main() {
	// basic
	sum := 0
	for i := 0; i < 10; i++ {	//Go 只有一种循环结构——`for` 循环。 没有 `( )` ，而 `{ }` 是必须的。
		sum += i
	}
	fmt.Println(sum)

	//可以让前置、后置语句为空
	sum = 1
	for ; sum < 1000; {
		sum += sum
	}
	fmt.Println(sum)

	//省略分号：C 的 while 在 Go 中叫做 `for`
	sum = 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)

	//如果省略了循环条件，循环就不会结束，因此可以用更简洁地形式表达死循环。
	sum = 1
	for {
		sum += sum
		if sum < 1000 {
			break
		}
	}
}

