package main

import "fmt"

func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	a, b := swap("凯", "高")
	fmt.Println(a, b)
}
