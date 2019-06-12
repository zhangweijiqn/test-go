package main

import (
	"fmt"
	"time"
)

func sum(a []int, c chan int) {
	sum := 0
	for _, v := range a {
		sum += v
	}
	c <- sum // 将和送入 c
}

func test() {
	// test channel
	a := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int) //chan
	go sum(a[:len(a)/2], c)
	go sum(a[len(a)/2:], c)
	x, y := <-c, <-c // 从 c 中获取

	fmt.Println(x, y, x+y)

	// test channel缓冲及阻塞
	c = make(chan int, 2)
	c <- 1
	c <- 2
	fmt.Println(<-c)
	fmt.Println(<-c)
	c <- 1
	c <- 2
	//c <- 3 //超过了缓冲大小，会报错

	// test close channel
	c = make(chan int, 10)
	go fibonacci(cap(c), c)
	for i := range c { //循环 `for i := range c` 会不断从 channel 接收值，直到它被关闭。
		fmt.Println(i)
	}
}

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
	/*
		发送者可以 close 一个 channel 来表示再没有值会被发送了。接收者可以通过赋值语句的第二参数来测试 channel 是否被关闭：当没有值可以接收并且 channel 已经被关闭，那么经过
		v, ok := <-ch, 之后 ok 会被设置为 `false`。

		注意： 只有发送者才能关闭 channel，而不是接收者。向一个已经关闭的 channel 发送数据会引起 panic。
		还要注意： channel 与文件不同；通常情况下无需关闭它们。只有在需要告诉接收者没有更多的数据的时候才有必要进行关闭，例如中断一个 `range`。
	*/
}

/*
	channel 是有类型的管道，可以用 channel 操作符 <- 对其发送或者接收值。

	ch <- v    // 将 v 送入 channel ch。
	v := <-ch  // 从 ch 接收，并且赋值给 v。
	（“箭头”就是数据流的方向。）

	和 map 与 slice 一样，channel 使用前必须创建：
	ch := make(chan int)

	默认情况下，在另一端准备好之前，发送和接收都会阻塞。这使得 goroutine 可以在没有明确的锁或竞态变量的情况下进行同步。

*/

func worker(jobs <-chan int) {
	for j := range jobs { //会一直监听channel中消息
		fmt.Println("started  job", j)
		time.Sleep(3 * time.Second)
		//fmt.Println("finished job", j)
	}
}

func testRange() {
	var ch = make(chan string, 2)
	go func() {
		for i := 0; i < 5; i++ {
			time.Sleep(500 * time.Millisecond)
			ch <- "hello"
		}
		close(ch)
	}()

	for recv := range ch { //注意go func内要调用close，否则range会阻塞
		fmt.Println(recv)
	}
}

func main() {
	//reference : https://gobyexample.com/worker-pools

	jobs := make(chan int, 20) //缓冲数，设置为并发的2倍
	//
	for j := 1; j <= 10; j++ { //work数量，并发数
		go worker(jobs)
	}
	//time.Sleep(time.Second * 1)
	//fmt.Printf("len = %v", len(jobs))

	for j := 1; j <= 50; j++ {
		fmt.Printf("i=%d\n", j)
		jobs <- j
	}

	println("dddd")

	for len(jobs) > 0 {
		time.Sleep(1 * time.Second)
	}

	println("done")
	time.Sleep(time.Second * 10) //保证主进程还在执行，否则gorutine进程会挂
	close(jobs)

	testRange()
}
