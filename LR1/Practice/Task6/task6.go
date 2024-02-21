package main

import (
	"fmt"
	"math"
)

func main() {
	var a float64

	fmt.Print("Введите целое число:")
	fmt.Scan(&a)

	result := math.Mod(a, 10)
	fmt.Println(result)
}
