package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Pilih operasi:")
	fmt.Println("1. Print 'Hello, World!'")
	fmt.Println("2. Jumlah bilangan ganjil dari 1 hingga 10")
	fmt.Print("Masukkan pilihan Anda: ")

	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	switch input {
	case "1":
		print.PrintHello()
	case "2":
		result := number.SumOddNumbers()
		fmt.Println("Jumlah bilangan ganjil dari 1 hingga 10 adalah:", result)
	default:
		fmt.Println("Pilihan tidak valid")
	}
}
