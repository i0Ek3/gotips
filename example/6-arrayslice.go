package main

import (
	"fmt"
)

func test() {
	var x = []int{2: 2, 3, 0: 1}
	fmt.Println(x)
}

func test1() {

	a := [3]int{0, 1, 2}
	fmt.Println(a)

	s := a[1:2]
	fmt.Println(s)

	s[0] = 11
	fmt.Println(s)

	s = append(s, 12)
	fmt.Println(s)

	s = append(s, 13)
	fmt.Println(s)

	s[0] = 21
	fmt.Println(s)

	fmt.Println(a)
}

type T struct {
	n int
}

func test2() {
	ts := [2]T{}
	for i := range ts[:] {
		switch i {
		case 0:
			ts[1].n = 9
		case 1:
			fmt.Print(ts[i].n, " ")
		}
	}
	fmt.Print(ts)
}

func test3() {
	var k = 9
	for k = range []int{} {
	}
	fmt.Println(k)

	for k = 0; k < 3; k++ {
	}
	fmt.Println(k)

	for k = range (*[3]int)(nil) {
	}
	fmt.Println(k)
}

func test4() {
	var a = []int{1, 2, 3, 4, 5}
	var r = make([]int, 0)
	for i, v := range a {
		if i == 0 {
			a = append(a, 6, 7)
		}
		r = append(r, v)
	}
	fmt.Println(r)
}

func test5() {
	var a = []int{1, 2, 3, 4, 5}
	var r [5]int

	for i, v := range a {
		if i == 0 {
			a[1] = 12
			a[2] = 13
		}
		r[i] = v
	}
	fmt.Println("r = ", r)
	fmt.Println("a = ", a)
}

func test6() {
	var a = [5]int{1, 2, 3, 4, 5}
	var r [5]int

	for i, v := range a {
		if i == 0 {
			a[1] = 12
			a[2] = 13
		}
		r[i] = v
	}
	fmt.Println("r = ", r)
	fmt.Println("a = ", a)
}

func test7() {
	s1 := []int{1, 2, 3}
	s2 := s1[1:]
	s2[1] = 4
	fmt.Println(s1)
	s2 = append(s2, 5, 6, 7)
	fmt.Println(s1)
	/* why?
	   golang 中切片底层的数据结构是数组。当使用 s1[1:] 获得切片 s2，和 s1 共享同一个底层数组，
	   这会导致 s2[1] = 4 语句影响 s1。 而 append 操作会导致底层数组扩容，生成新的数组，
	   因此追加数据后的 s2 不会影响 s1。 但是为什么对 s2 赋值后影响的却是 s1 的第三个元素呢？
	   这是因为切片 s2 是从数组的第二个元素开始，s2 索引为 1 的元素 对应的是 s1 索引为 2 的元素。
	*/
}

func test8() {
	s := make([]int, 5)
	s = append(s, 1, 2, 3)
	fmt.Println(s)

	s2 := make([]int, 0)
	s2 = append(s2, 1, 2, 3, 4)
	fmt.Println(s2)
}

func test9() {
	s := []int{1, 2}
	s = append(s, 4, 5, 6, 7)
	fmt.Printf("len = %d, cap = %d", len(s), cap(s))
	// 6 6
}

func main() {
	test()
	test1()
	test2()
	test3()
	test4()
	test5()
	test6()
	test7()
	test8()
	test9()
}
