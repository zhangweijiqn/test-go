package main

import (
	"fmt"
	"time"
)

func main() {
	curDay := time.Now()
	fmt.Println(curDay)
	baseTime := time.Date(curDay.Year(), 10, 25, 0, 0, 0, 0, time.UTC)
	date := baseTime.Add(1*7*24*time.Hour + 24*time.Hour + 66355*time.Second)
	fmt.Println(date.Format("2006-01-02 15:04:05"))
	fmt.Println(baseTime.AddDate(0, 0, -2))
}
