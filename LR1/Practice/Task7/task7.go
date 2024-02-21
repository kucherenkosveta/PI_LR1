package main

import "fmt"

func main() {
	var a int

	fmt.Print("Введите целое положительное число:")
	fmt.Scan(&a)

	if a < 0 || a > 10000 {
		fmt.Println("Введенное число не соответствует условиям")
		return
	}

	result := (a / 10) % 10
	fmt.Println(result)
}
