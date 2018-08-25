package main

import "fmt"

func main() {
	var numbers = make([]int, 3, 5)
	fmt.Println(numbers)

	var numbers1 = make([]int, 0, 5)
	numbers1 = append(numbers1, 1, 2, 3, 4, 5, 6, 7, 8, 9)
	fmt.Println(numbers1)
}
