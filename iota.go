package main

import "fmt"

func main() {
	const (
		a = iota // 0
		b        //1
		c        //2
		d = "ha" // 独立值，iota += 1
		e
		f = 100
		g
		h = iota
		i
	)
	fmt.Println(a, b, c, d, e, f, g, h, i)
}
