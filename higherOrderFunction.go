package main

func applyFunction(f func(int) int, x int) int {
	return f(x)
}

type Operation func(int, int) int

func main() {
	increment := func(x int) int {
		return x + 1
	}
	println(applyFunction(increment, 5)) // 6
	ops := map[string]Operation{
		"add":  func(a, b int) int { return a + b },
		"sub":  func(a, b int) int { return a - b },
		"mult": func(a, b int) int { return a * b },
		"div":  func(a, b int) int { return a / b },
	}
	result := ops["add"](10, 20)
	println(result)
}
