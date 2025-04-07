package main

import (
	"fmt"
	"strconv"
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

func addStarsBetweenLetters(s string) string {
	var result []rune
	for i, r := range s {
		if i > 0 {
			result = append(result, '*')
		}
		result = append(result, r)
	}
	return string(result)
}

func squareDigits(n int) int {
	str := strconv.Itoa(n)
	var resultStr string
	for _, ch := range str {
		digit := int(ch - '0')
		squared := digit * digit
		resultStr += strconv.Itoa(squared)
	}
	result, _ := strconv.Atoi(resultStr)
	return result
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
	// Задача 4
	testStrings := []string{
		"Hello",
		"Go-1.2.3",
		"Привет!",
		"😀👍🚀",
		"A",
		"",
	}

	for _, s := range testStrings {
		result := addStarsBetweenLetters(s)
		fmt.Printf("Исходная: '%s'\nРезультат: '%s'\n\n", s, result)
	}

	// Задача 5
	testNumbers := []int{
		9119,
		1001,
		7,
		222,
		987654321,
		0,
	}

	for _, num := range testNumbers {
		result := squareDigits(num)
		fmt.Printf("%d → %d\n", num, result)
	}
}
