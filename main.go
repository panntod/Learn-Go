package main

import "fmt"

func main() {
	name := "Pandhu"
	age := 13

	fmt.Println("Hello World")
	fmt.Println(len(name))
	fmt.Println(greeting(name, age))
}

func greeting(name string, age int) string {
	return fmt.Sprintf("Hello %s, you're %d", name, age)
}
