package main

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"strconv"
)

type ItemType int32

const (
	ItemType_UNKNOWN ItemType = 0
	ItemType_ARTICLE ItemType = 1
	ItemType_VIDEO   ItemType = 2
	ItemType_LIVE    ItemType = 3
)

var ItemType_name = map[int32]string{
	0: "UNKNOWN",
	1: "ARTICLE",
	2: "VIDEO",
	3: "LIVE",
}
var ItemType_value = map[string]int32{
	"UNKNOWN": 0,
	"ARTICLE": 1,
	"VIDEO":   2,
	"LIVE":    3,
}

func (x ItemType) Enum() *ItemType {
	p := new(ItemType)
	*p = x
	return p
}
func (x ItemType) String() string {
	return proto.EnumName(ItemType_name, int32(x))
}
func (x *ItemType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(ItemType_value, data, "ItemType")
	if err != nil {
		return err
	}
	*x = ItemType(value)
	return nil
}

type cacheValue struct {
	itemType *ItemType
}

func main() {
	var i = ItemType_VIDEO
	fmt.Println(i)
	fmt.Println(i.String())
	fmt.Println(i.Enum())
	fmt.Printf("%d\n", i)
	fmt.Println(1 == i)
	fmt.Println(int32(i))
	fmt.Println(strconv.Itoa(int(i)))
	s := fmt.Sprintf("%s[%d][%v][%i]", i, i, i, i)
	fmt.Println(s)

	cV := cacheValue{&i}
	s = fmt.Sprintf("%s[%d][%v][%i]", *cV.itemType, *cV.itemType, *cV.itemType, *cV.itemType)
	fmt.Println(s)
}
