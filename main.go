package main

import (
	"fmt"
)

func main() {
	const uri = "pandhu-munjalindra.xyz"

	var (
		name           = "Pandhu"
		age            = 13
		age1           = 25
		age2           = 30
		studentGrade   = 95
		studentPresent = 80
	)

	fmt.Println(greeting(name, age, uri))

	/* fmt.Println("Hello World")
	fmt.Println(len(name)) */

	fmt.Println(calculate(1, 3))
	convertion(213132)
	fmt.Println(convertionString("Pandhu", 4))

	numbers := []int{1, 2, 4, 5, 6, 7, 8, 9, 12}
	oddNumbers := findOddNumbers(numbers)
	fmt.Println("Odd Number in Array:", oddNumbers)

	olderAge := findOlder(age1, age2)
	fmt.Println("Person older is:", olderAge)

	graduation := checkGrade(studentGrade, studentPresent)
	fmt.Println(checkPass(graduation))

	fmt.Println(checkLenName(name))
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
	var find byte = name[index-1]
	var findString string = string(find)
	return fmt.Sprintf("This Index %d from %s is %d, and convert to %s", index, name, find, findString)
}

func findOddNumbers(nums []int) []int {
	var oddNumbers []int

	for _, num := range nums {
		if num%2 != 0 {
			oddNumbers = append(oddNumbers, num)
		}
	}

	return oddNumbers
}

func findOlder(person1 int, person2 int) string {
	if person1 > person2 {
		return "person 1"
	}
	return "person 2"
}

func checkGrade(grade int, present int) bool {
	check := grade > 90
	checkPresent := present > 80

	return check && checkPresent
}

func checkPass(passed bool) string {
	if passed {
		return "Congratulations you passed."
	}

	return "Sorry you're not passing."
}

func checkLenName(name string) string {
	var message string
	if length := len(name); length < 4 {
		message = "name must be more than 4 letters"
	} else {
		message = fmt.Sprint("Your name is ", name)
	}
	return message
}
