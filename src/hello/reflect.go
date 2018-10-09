package main

import (
	"fmt"
	"reflect"
)

type s interface {
	f()
}
type ss interface {
}

type s1 struct {
}
type s2 struct {
}

func (a1 *s1) f() {
	fmt.Println("s1")
}
func (a1 *s2) f() {
	fmt.Println("s2")
}

func main() {
	t := "s2"
	fmt.Println(reflect.TypeOf(t))
	fmt.Println(reflect.TypeOf(t).Kind())
	fmt.Println(reflect.ValueOf(t))
	tt := reflect.ValueOf(t)
	tt.MethodByName("Print") //调用动态方法
	tt1 := tt.Interface().(s)
	tt1.f()

	/*
		var ttt s
		reflect.ValueOf(ttt).MethodByName("f")
		ttt.f()
	*/
	//var s11 s = s1(t)
	//s11.f()

}
