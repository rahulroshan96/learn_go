package main // Name of current package

import (
	"fmt" // package
	"math"
	"time" // import time, will import all files that has statement package time in
)

func main() {
	fmt.Println("Welcome to the playground!")
	fmt.Println(math.Pi) // Pi is exported from math.go, because the name starts with Capital letter
	fmt.Println("The time is", time.Now())
}
