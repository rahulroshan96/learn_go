package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

func sub(x, y int) int {
	return x - y
}

func swap(x, y string) (string, string) {
	return y, x
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return // naked return, return automatically as x and y and treated as var defined at top of function
}

var c, python, java bool        // type is int last for similar types of variables
var m, n, p = true, false, "no" // var assignment
var i, j int = 100, 200

func main() {
	k := 'a' // short assignment, implicit type declaration. Allowed inside function only

	fmt.Println(add(42, 13))
	fmt.Println(sub(10, 4))
	a, b := swap("2", "3")
	fmt.Println(a, b)
	fmt.Println(split(17))
	fmt.Println(i, j, c, python, java)
}
