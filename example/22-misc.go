package main

import "fmt"

func alwaysFalse() bool {
	return false
}

func test1() {
	switch alwaysFalse(); {
	case true:
		println(true)
	case false:
		println(false)
	}
}

func test2() {
	x := 1
	fmt.Println(x)

	{
		fmt.Println(x)
		i, x := 2, 2
		fmt.Println(i, x)
	}
	fmt.Println(x)
}

func min(a int, b uint) {
	var min = 0
	min = copy(make([]struct{}, a), make([]struct{}, b))
	fmt.Printf("The min of %d and %d is %d\n", a, b, min)

}

func test3() {
	min(1225, 256)
}

func main() {
	test1()
	test2()
	test3()
}
