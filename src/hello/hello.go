package main	//go通过package组织代码，一个package由目录下的go组成，目录定义package的作用，每个go都以package声明开始，必须要有main package

import "fmt"	//import必须在package之后， 包含格式化input和output函数

func main() {	//main函数同样是程序入口，func声明函数，{必须和函数在一行
	fmt.Println("Hello, world!\n")
}

/*
	查看go版本： go version
	编译命令：	go build hello.go
	执行： 		./hello
	编译+执行： 	go run hello.go
 */