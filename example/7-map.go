package main

import "fmt"

func test1() {
	var m = map[string]int{
		"A": 21,
		"B": 22,
		"C": 23,
	}

	counter := 0
	for k, v := range m {
		if counter == 0 {
			delete(m, "A")
		}
		counter++
		fmt.Println(k, v)
	}
	fmt.Println("counter is ", counter)
}

func test2() {
	slice := []int{0, 1, 2, 3}
	m := make(map[int]*int)
	for key, val := range slice {
		m[key] = &val
	}
	for k, v := range m {
		fmt.Println(k, "->", *v)
	}
}

func main() {
	test1()
	test2()
}
