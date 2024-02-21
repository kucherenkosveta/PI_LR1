package main

import "fmt"

func main() {
	var a int

	fmt.Print("Введите целое число:")
	fmt.Scan(&a)

	result := a * a
	fmt.Println(result)
}
