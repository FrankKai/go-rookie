package main

import "fmt"

func main() {
	a := 1
	b := 122

	var result int

	result = max(a, b)

	fmt.Println(result)
}

func max(num1, num2 int) int {
	var result int

	if num1 > num2 {
		result = num1
	} else {
		result = num2
	}
	return result
}
