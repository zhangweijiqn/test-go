package main

import "fmt"


const (
	Big   = 1 << 100
	Small = Big >> 99
)

func needInt(x int) int { return x*10 + 1 }

func needFloat(x float64) float64 {
	return x * 0.1
}

func main() {
	// 常量定义
	const Pi = 3.14
	const World = "world"
	const Truth = true
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")
	fmt.Println("Go rules?", Truth)

	//未指定类型的常量由上下文来决定其类型
	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
	//fmt.Println(needInt(Big))
}

