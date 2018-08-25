package main

import "fmt"

type Books struct {
	title   string
	author  string
	subject string
	id      int
}

func main() {
	var Book1 Books

	Book1.title = "Go 语言"
	Book1.author = "w3cschool"
	Book1.subject = "Go 语言教程"
	Book1.id = 123456

	Book2 := Books{"Go 语言", "w3cschool", "Go 语言教程", 123456}

	fmt.Println("Book1:", Book1)
	fmt.Println("Book2:", Book2)

	printBook(&Book1)
}

func printBook(book *Books) {
	fmt.Println(book)
}
