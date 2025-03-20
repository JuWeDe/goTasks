package main

import (
	"fmt"
)

func reverseNumber(n int) int {
	return (n%10)*100 + (n/10%10)*10 + n/100
}

func isRightTriangle(a, b, c int) string {
	if a*a+b*b == c*c {
		return "Прямоугольный"
	}
	return "Непрямоугольный"
}

func timeFromSeconds(k int) (int, int) {
	hours := k / 3600
	minutes := (k % 3600) / 60
	return hours, minutes
}

func main() {
	number := 843
	a, b, c := 6, 8, 10
	seconds := 13257

	// Задача 1
	reversed := reverseNumber(number)
	fmt.Printf("Перевернутое число: %d\n", reversed)

	// Задача 2
	result := isRightTriangle(a, b, c)
	fmt.Println(result)

	// Задача 3
	hours, minutes := timeFromSeconds(seconds)
	fmt.Printf("It is %d hours %d minutes.\n", hours, minutes)
}
