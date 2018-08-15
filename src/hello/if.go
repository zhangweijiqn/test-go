package main
import (
	"fmt"
	"math"
)

func sqrt(x float64) string {
	if x < 0 {					//if 语句除了没有了 `( )` 之外（甚至强制不能使用它们），而 `{ }` 是必须的
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {				//跟 for 一样，`if` 语句可以在条件之前执行一个简单的语句。
		return v									//由这个语句定义的变量的作用域仅在 if 范围之内
	} else {
		fmt.Printf("%g >= %g\n", v, lim)		// 在 if 的便捷语句定义的变量同样可以在任何对应的 else 块中使用。
	}
	// 这里开始就不能使用 v 了
	return lim
}

func main() {
	fmt.Println(sqrt(2), sqrt(-4))

	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
}

