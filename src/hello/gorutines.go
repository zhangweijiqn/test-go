package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	go say("world") //和下面的线程独立，不是顺序执行
	say("hello")
}

/*
	goroutine 是由 Go 运行时环境管理的轻量级线程
	go f()	仅需要在执行的函数前面添加关键词 go
*/
