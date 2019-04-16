package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	//Stringer 是一个可以用字符串描述自己的类型。`fmt`包 （还有许多其他包）使用这个来进行输出
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

func main() {
	// test String()
	a := Person{"Arthur Dent", 42}
	z := Person{"Zaphod Beeblebrox", 9001}
	fmt.Println(a, "\n", z) //自动调用String()

	// test String() 2
	addrs := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for n, a := range addrs {
		fmt.Printf("%v: %v\n", n, a)
	}

	aaa := "aaaabbb"
	fmt.Println(strings.Index(aaa, "ab"))
	fmt.Println(strings.Index(aaa, "abc"))

	// test %v
	TestFormat()

	var urlToken int64
	fmt.Println(strconv.FormatInt(urlToken, 10))
}

type Power struct {
	age  int
	high int
	name string
}

func TestFormat() {
	var i Power = Power{age: 10, high: 178, name: "NewMan"}

	fmt.Print("\n\n")
	fmt.Printf("type:%T\n", i)
	fmt.Printf("value:%v\n", i)
	fmt.Printf("value+:%+v\n", i)
	fmt.Printf("value#:%#v\n", i)

	fmt.Println("========interface========")
	var interf interface{} = i
	fmt.Printf("%v\n", interf)
	fmt.Println(interf)
}

type IPAddr [4]byte

func (ip IPAddr) String() string {
	var ip_addr = ""
	for _, v := range ip {
		if ip_addr == "" {
			ip_addr = fmt.Sprintf("%v", v)
		} else {
			ip_addr += fmt.Sprintf(".%v", v)
		}
	}
	return ip_addr
}
