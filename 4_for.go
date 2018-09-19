package main

import (
	"fmt"
)

func main() {
	sum := 0
	for i := 0; i < 5; i++ {
		sum += i
	}
	fmt.Println(sum)
	sum = 1
	for sum < 5 { // first and last conditions are optional
		sum += sum
	}
	fmt.Println(sum)

	// for {  // Forever loop

	// }

	i := 0
	for i < 10 { // C's while is spelled as for in Go
		fmt.Println(i)
		i += 1
	}
}
