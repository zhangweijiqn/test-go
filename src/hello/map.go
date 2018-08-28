package main

import "fmt"

type Vertx struct {
	Lat, Long float64
}

var m map[string]Vertx //map 映射键到值

func test_map_basic() {
	m = make(map[string]Vertx) //map 在使用之前必须用 make 而不是 new 来创建；值为 nil 的 map 是空的，并且不能赋值。
	m["Bell Labs"] = Vertx{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])

	var m1 = map[string]Vertx{ //map 的文法跟结构体文法相似，不过必须有键名
		"Bell Labs": Vertx{
			40.68433, -74.39967,
		},
		"Google": Vertx{
			37.42202, -122.08408,
		},
	}
	fmt.Println(m1)

	type Vertex struct {
		Lat, Long float64
	}

	var m2 = map[string]Vertx{
		"Bell Labs": {40.68433, -74.39967}, //如果顶级的类型只有类型名的话，可以在文法的元素中省略键名
		"Google":    {37.42202, -122.08408},
	}
	fmt.Println(m2)
}

func test_map_change() {
	m := make(map[string]int)

	m["Answer"] = 42                       //插入或修改一个元素
	fmt.Println("The value:", m["Answer"]) //获得元素

	m["Answer"] = 48
	fmt.Println("The value:", m["Answer"])

	v, ok := m["Answer"] //通过双赋值检测某个键存在，如果 key 在 m 中，`ok` 为 true 。否则， ok 为 `false`，并且 elem 是 map 的元素类型的零值。 同样的，当从 map 中读取某个不存在的键时，结果是 map 的元素类型的零值。
	fmt.Println("The value:", v, "Present?", ok)

	delete(m, "Answer") //删除元素
	fmt.Println("The value:", m["Answer"])

	v1, ok1 := m["Answer"]
	fmt.Println("The value:", v1, "Present?", ok1)
}

func main() {

	test_map_basic()

	test_map_change()

}
