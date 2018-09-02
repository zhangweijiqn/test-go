package main

import (
	"fmt"
	"log"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	say("hello")    //后面有可能执行不了，因为并发执行的代码在未执行完之前主程序已exit （除非先将主程序hang住）
	go say("world") //和下面的线程独立，不是顺序执行
	//log.Fatal("this is a fatal log")
	go func() {
		log.Fatal("this is a fatal log!") //log.Fatal会触发程序Exit
	}()

}

/*
	goroutine 是由 Go 运行时环境管理的轻量级线程
	go f()	仅需要在执行的函数前面添加关键词 go
*/
