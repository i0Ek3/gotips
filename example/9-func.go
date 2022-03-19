package main

import (
	"fmt"
	"testing"
)

func hello(num ...int) {
	num[0] = 18
}

func Test13(t *testing.T) {
	i := []int{5, 6, 7}
	hello(i...)
	fmt.Println(i[0])
	fmt.Println(i)

}

var f = func(i int) {
	print("x")
}

func test2() {
	f := func(i int) {
		print(i)
		if i > 0 {
			f(i - 1)
		}
	}
	f(10)
}

func test1() {
	t := &testing.T{}
	Test13(t)
}

func main() {
	test1()
	test2()
}
