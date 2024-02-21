package main

import "fmt"

func main() {
	var d int

	fmt.Print("Введите целое положительное число:")
	fmt.Scan(&d)

	if d < 0 || d > 360 {
		fmt.Println("Введенное число не соответствует условиям")
		return
	}

	h := d / 30
	m := (d % 30) * 2
	fmt.Println("It is ", h, "hours ", m, "minutes.")
}
