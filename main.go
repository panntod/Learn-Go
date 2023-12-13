package main

import (
	"fmt"
)

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
	convertion(213132)
	fmt.Println(convertionString("Pandhu", 4))
}

func greeting(name string, age int, uri string) string {
	return fmt.Sprintf("Hello %s, you're %d, See my personal website %s", name, age, uri)
}

func calculate(num1 int, num2 int) string {
	equal := num1 + num2
	result := fmt.Sprintf("Result of %d + %d = %d", num1, num2, equal)
	return result
}

func convertion(num int) int {
	var nilai32 int32 = int32(num)
	fmt.Println("This is a int32", nilai32)

	var nilai64 int64 = int64(num)
	fmt.Println("This is a int64: ", nilai64)

	return num
}

func convertionString(name string, index int) string {
	var find = name[index-1]
	var findString = string(find)
	return fmt.Sprintf("This Index %d from %s is %d, and convert to %s", index, name, find, findString)
}
