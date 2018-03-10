package main

import "fmt"
import "math"

var x int

func add(a int, b int) int {
	return a + b
}

func main() {
	x := 3
	y := 2
	fmt.Println(add(42, 13), x, y)
	fmt.Println(math.Pi)
}
