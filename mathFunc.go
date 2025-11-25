package main

import _ "fmt"

type MathFunc func(int, int) int

func calculate(a, b int, operation MathFunc) int {
	return operation(a, b)
}
func main() {
	var add = func(a, b int) int {
		return a + b
	}
	var result = calculate(1, 2, add)
	println(result)
}
