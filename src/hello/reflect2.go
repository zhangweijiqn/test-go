package main

import (
	"fmt"
	"reflect"
)

/*
	根据传入的string返回string对应的真实类型
*/

var typeRegistry = make(map[string]reflect.Type)

func registerType(typedNil interface{}) {
	t := reflect.TypeOf(typedNil).Elem()
	typeRegistry[t.PkgPath()+"."+t.Name()] = t
}

type MyString string
type myString string

func init() {
	registerType((*MyString)(nil))
	registerType((*myString)(nil))
	registerType((*struct1)(nil))
	// ...
}

func makeInstance(name string) interface{} {
	return reflect.New(typeRegistry[name]).Elem().Interface()
}

func main() {
	for k := range typeRegistry {
		fmt.Println(k)
	}
	My := makeInstance("main.MyString")
	my := makeInstance("main.myString")
	s1 := makeInstance("main.struct1")
	fmt.Printf("%T\n", My)
	fmt.Printf("%T\n", my)
	fmt.Printf("%T\n", s1)

	reflect.ValueOf(s1).MethodByName("f")

}
