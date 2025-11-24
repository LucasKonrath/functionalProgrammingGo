package main

import "fmt"

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

var firstClassFunction = func(name string) {
	fmt.Println("Hello, " + name)
}

var closure = func(prefix string) func(name string) {
	return func(name string) {
		fmt.Println(prefix, name)
	}
}

func main() {
	firstClassFunction("World")

	var helloGreeting = closure("Hello ")
	helloGreeting("World")

}
