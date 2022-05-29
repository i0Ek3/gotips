package main

import (
	"fmt"
	"runtime"
)

var o = fmt.Println

func test1() {
	c := make(chan int, 1)
	for range [3]struct{}{} {
		// 0 1 2 in [3]struct{}{}
		select {
		// ready to read
		case <-c:
			o(2)
			c = nil
		// ready to write
		case c <- 1:
			o(3)
		// now c is nil
		default:
			o(1)
		}
	}
}

func test2() {
	var ch chan int
	select {
	case v, ok := <-ch:
		o(v, ok)
	default:
		o("default")
	}
}

func test3() {
	runtime.GOMAXPROCS(1)

	int_chan := make(chan int, 1)
	string_chan := make(chan string, 1)

	int_chan <- 1
	string_chan <- "hello"

	select {
	case value := <-int_chan:
		fmt.Println(value)
	case value := <-string_chan:
		panic(value)
	}
}

func test4() {
    ch := make(chan int)
    for i := range ch {
        ch <- i
    }
    close(ch)
}

func main() {
	test1()
	test2()
	test3()
    test4()
}
