package main

func sumOddNumbers() int {
	sum := 0

	for i := 1; i <= 10; i++ {
		if i%2 != 0 {
			sum += i
		}
	}

	return sum
}
