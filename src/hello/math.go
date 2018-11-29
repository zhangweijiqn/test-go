package main

import (
	"fmt"
	"math"
	"strconv"
	"time"
)

func getCurrentTimestamp() int64 {
	t := time.Now()
	timestamp := strconv.FormatInt(t.UTC().UnixNano(), 10)
	timestamp = timestamp[:10]
	timestampInt, _ := strconv.ParseInt(timestamp, 10, 64)
	return timestampInt
}

func GetNow() int64 {
	return time.Now().Unix()
}

func main() {
	data := math.Log10(0)
	fmt.Println(data)
	fmt.Println(getCurrentTimestamp())
	fmt.Println(GetNow())
}
