package main

import "fmt"

func main() {
	var a int = 10

	var ip *int

	ip = &a
	fmt.Println(a, &a, ip, *ip)
}
