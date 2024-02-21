package main

import "fmt"

func main() {
	var number int
	var res int

	fmt.Print("Введите число: ")
	fmt.Scan(&number)

	res = number * 2
	res = res + 100

	fmt.Println(res)
}
