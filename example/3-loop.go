package main

import "fmt"

func test1() {
	x := []int{0, 1, 2}
	y := [3]*int{}
	for i, v := range x {
		defer func() {
			print(v)
		}()
		y[i] = &v
	}
	print(*y[0], *y[1], *y[2])
}

func test2() {
	isMatch := func(i int) bool {
		switch i {
		case 1:
		case 2:
			return true
		}
		return false
	}
	fmt.Println(isMatch(1))
	fmt.Println(isMatch(2))
}

func test3() {
	i := 0
	f := func() int {
		i++
		return i
	}
	c := make(chan int, 1)
	c <- f()
	select {
	case c <- f():
	default:
		fmt.Println(i)
	}
}

func main() {
	test1()
	test2()
	test3()
}
