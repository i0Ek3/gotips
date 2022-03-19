package main

import "fmt"

func main() {
	var a byte = 0x11
	fmt.Println(a)

	var b uint8 = a
	fmt.Println(b)

	var c uint8 = a + b
	fmt.Println(c)
}
