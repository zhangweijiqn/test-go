package main

/*
	使用sort包对任意类型元素的集合进行排序
*/

import (
	"fmt"
	"sort"
)

//对任意对象进行排序
type Person2 struct {
	name string
	age  int
}

//为*Person添加String()方法，便于输出
func (p *Person2) String() string {
	return fmt.Sprintf("( %s,%d )", p.name, p.age)
}

type PersonList []*Person2

//排序规则：首先按年龄排序（由小到大），年龄相同时按姓名进行排序（按字符串的自然顺序）
// go 中对某个 Type 的对象 obj 排序， 可以使用 sort.Sort(obj) 即可，就是需要对 Type 类型绑定三个方法 ：
// Len() 求长度、Less(i,j) 比较第 i 和 第 j 个元素大小的函数、 Swap(i,j) 交换第 i 和第 j 个元素的函数。

func (list PersonList) Len() int {
	return len(list)
}

func (list PersonList) Less(i, j int) bool {
	if list[i].age < list[j].age {
		return true
	} else if list[i].age > list[j].age {
		return false
	} else {
		return list[i].name < list[j].name
	}
}

func (list PersonList) Swap(i, j int) {
	var temp *Person2 = list[i]
	list[i] = list[j]
	list[j] = temp
}

func main() {
	fmt.Println("------")
	p1 := &Person2{"Tom", 19}
	p2 := &Person2{"Hanks", 19}
	p3 := &Person2{"Amy", 19}
	p4 := &Person2{"Tom", 20}
	p5 := &Person2{"Jogn", 21}
	p6 := &Person2{"Mike", 23}

	pList := PersonList([]*Person2{p1, p2, p3, p4, p5, p6})
	//sort.Sort(pList)
	sort.Stable(pList)
	fmt.Println(pList)

	/*output:
	[( Amy,19 ) ( Hanks,19 ) ( Tom,19 ) ( Tom,20 ) ( Jogn,21 ) ( Mike,23 )] */
}
