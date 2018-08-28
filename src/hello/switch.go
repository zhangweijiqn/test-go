package main

import (
	"fmt"
	"runtime"
	"time"
)

func test_switch1() {
	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	switch time.Saturday { //switch 的条件从上到下的执行，当匹配成功的时候停止。
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}
}

func test_switch2() {
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os { //跟 for 一样，`switch` 语句可以在条件之前执行一个简单的语句。
	case "darwin":
		fmt.Println("OS X.") //除非以 fallthrough 语句结束，否则分支会自动终止。
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.", os)
	}
}

func test_switch3() {
	t := time.Now()
	switch { //没有条件的 switch 同 `switch true` 一样, 这一构造使得可以用更清晰的形式来编写长的 if-then-else 链。
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}

func main() {
	test_switch1()
	test_switch2()
	test_switch3()
}
