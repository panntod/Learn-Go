package main

import "fmt"

func main() {
	const uri = "pandhu-munjalindra.xyz"

	var (
		name = "Pandhu"
		age  = 13
	)

	/* fmt.Println("Hello World")
	fmt.Println(len(name)) */

	fmt.Println(greeting(name, age, uri))
	fmt.Println(calculate(1, 3))
}

func greeting(name string, age int, uri string) string {
	return fmt.Sprintf("Hello %s, you're %d, See my personal website %s", name, age, uri)
}

func calculate(num1 int, num2 int) string {
	equal := num1 + num2
	result := fmt.Sprintf("Result of %d + %d = %d", num1, num2, equal)
	return result
}
