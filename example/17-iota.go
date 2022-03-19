package main

import "fmt"

type Direction int

const (
	North Direction = iota
	East
	South
	West
)

const (
	a    = iota
	b    = iota
	name = "name"
	c    = iota
	d    = iota
)

const (
	x = iota
	_
	y
	z = "zz"
	k
	p = iota
)

func (d Direction) String() string {
	return [...]string{"North", "East", "South", "West"}[d]
}

func test1() {
	fmt.Println(South)
}

func test2() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
}

func test3() {
	fmt.Println(x, y, z, k, p)
}

func main() {
	test1()
	test2()
	test3()
}
