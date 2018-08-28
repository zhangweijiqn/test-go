package main

import (
	"fmt"
	"time"

	"math"
)

type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	//与 fmt.Stringer 类似，`error` 类型是一个内建接口
	return fmt.Sprintf("at %v, %s",
		e.When, e.What)
}

func run() error {
	return &MyError{ // error接口接收MyError对象，打印的时候会试图调用Error方法
		time.Now(),
		"it didn't work",
	}
}

func main() {
	if err := run(); err != nil { //error 为 nil 时表示成功；非 nil 的 error 表示错误
		fmt.Println(err)
	}

	if result, err := Sqrt(-2); err == nil {
		fmt.Println(result)
	} else {
		fmt.Println(result)
		fmt.Println(err)
	}
}

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number:%v", float64(e))
	// 在 Error 方法内调用 fmt.Sprint(e) 将会让程序陷入死循环。
	// 可以通过先转换 e 来避免这个问题：`fmt.Sprint(float64(e))`。
}

func Sqrt(x float64) (float64, error) {
	if x >= 0 {
		return math.Sqrt(x), nil
	} else {
		return 0, ErrNegativeSqrt(x)
	}

}
