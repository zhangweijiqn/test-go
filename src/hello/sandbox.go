package main

import (
	"fmt"
	"time"
	"math/rand"
	"math"
)

func main() {
	fmt.Println("Welcome to the playground!")

	fmt.Println("The time is", time.Now())

	fmt.Println("My favorite number is", rand.Intn(10))

	fmt.Printf("Now you have %g problems.\n", math.Nextafter(2, 3))	//Printf

	fmt.Println(math.Pi)	//export
}
