package main

import "unsafe"

const (
	a = "abc"
	b = len(a)
	c = unsafe.Sizeof(a) // unsafe.Sizeof采用任何类型的表达式x，并返回假设变量v的大小（以字节为单位），就像通过var v = x声明v一样。大小不包括x可能引用的任何内存。例如，如果x是切片，则Sizeof返回切片描述符的大小，而不是切片引用的内存大小。
)

func main() {
	println(a, b, c)
}
