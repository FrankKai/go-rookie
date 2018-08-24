package main

import "fmt"

func main() {

	b := 15
	var a int

	for a := 0; a < 10; a++ {
		fmt.Println(a)
	}

	for a < b {
		a++
		fmt.Println(a)
	}

	numbers := [6]int{1, 2, 3, 5}
	for i, x := range numbers {
		fmt.Printf("第%d位x的值 = %d\n", i, x)
	}
}
