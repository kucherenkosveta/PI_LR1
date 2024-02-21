package main

import (
	"fmt"
	"math"
)

func main() {
	a := math.Abs(-5.67)              // Возвращает абсолютное значение числа.
	b := math.Ceil(5.67)              // Округляет число вверх до ближайшего целого.
	c := math.Floor(5.67)             // Округляет число вниз до ближайшего целого.
	d := math.Sqrt(16)                // Возвращает квадратный корень числа.
	f := math.Pow(2, 3)               // Возводит число x в степень y .
	sinValue := math.Sin(math.Pi / 2) // Возвращают синус, косинус и тангенс угла в радианах соответственно.
	g := math.Log(10)                 // Возвращает натуральный логарифм числа.
	maxVal := math.Max(3, 7)          // Макс значение
	minVal := math.Min(3, 7)          // Мин значение
	h := math.Mod(10, 3)              // Возвращает остаток от деления x на y.
	i := math.Round(5.67)             // Округляет число к ближайшему целому.
	j := math.Trunc(5.67)             // Отбрасывает дробную часть числа.
	posInf := math.Inf(1)             // Возвращает положительную бесконечность
	negInf := math.Inf(-1)            // Возвращает отрицательную бесконечность
	nan := math.NaN()                 // Возвращает "Not a Number" (NaN).
	k := math.Exp(2)                  // Возвращает экспоненту (e) в степени x.
	l := math.Exp2(3)                 // Возвращает 2 в степени x.
	m := math.Expm1(1)                // Возвращает e в степени x минус 1
	n := math.Log10(100)              // Возвращает десятичный логарифм числа.
	o := math.Log2(8)                 // Возвращает двоичный логарифм числа.
	p := math.Log1p(1)                // Возвращает двоичный логарифм числа х плюс 1
	isNegative := math.Signbit(-5)    // Возвращает true, если x отрицательное или отрицательный нуль.

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(f)
	fmt.Println(sinValue)
	fmt.Println(g)
	fmt.Println(maxVal)
	fmt.Println(minVal)
	fmt.Println(h)
	fmt.Println(i)
	fmt.Println(j)
	fmt.Println(posInf)
	fmt.Println(negInf)
	fmt.Println(nan)
	fmt.Println(k)
	fmt.Println(l)
	fmt.Println(m)
	fmt.Println(n)
	fmt.Println(o)
	fmt.Println(p)
	fmt.Println(isNegative)

}
