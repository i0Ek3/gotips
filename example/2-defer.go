package main

import "fmt"

type Slice []int

func NewSlice() Slice {
	return make(Slice, 0)
}

func (s *Slice) Add(elem int) *Slice {
	*s = append(*s, elem)
	fmt.Print(elem)
	return s
}

func test1() {
	s := NewSlice()
	defer func() {
		s.Add(1).Add(2)
	}()
	s.Add(3)
}

func DeferTest1(i int) (r int) {
	r = i
	defer func() {
		r += 3
	}()
	return r
}

func DeferTest2(i int) (r int) {
	defer func() {
		r += i
	}()
	return 2
}

func test2() {
	println(DeferTest1(1))
	println(DeferTest2(1))
}

func F(n int) func() int {
	return func() int {
		n++
		return n
	}
}

func test3() {
	f := F(5)
	defer func() {
		fmt.Println(f())
	}()
	defer fmt.Println(f())
	i := f()
	fmt.Println(i)
}

func test4() {
	defer func() {
		fmt.Print(recover())
	}()

	defer func() {
		defer func() {
			fmt.Print(recover())
		}()
		panic(1)
	}()

	defer recover()
	panic(2)
}

func f(n int) (r int) {
	defer func() {
		r += n
		recover()
	}()

	var f func()
	defer f()

	f = func() {
		r += 2
	}
	return n + 1
}

func test5() {
	fmt.Println(f(3))
}

func test6() {
	a := 1
	b := 2
	defer calc("1", a, calc("10", a, b))
	a = 0
	defer calc("2", a, calc("20", a, b))
	b = 1

	/* why
	程序执行到 main() 函数的第三行代码的时候，会先执行 calc() 函数的 b 参数，即：calc(“10”,a,b)，输出：10 1 2 3，得到值 3，因为 defer 定义的函数是延迟函数，故 calc(“1”,1,3) 会被延迟执行；
	程序执行到第五行的时候，同样先执行 calc(“20”,a,b) 输出：20 0 2 2 得到值 2，同样将 calc(“2”,0,2) 延迟执行；
	程序执行到末尾的时候，按照栈先进后出的方式依次执行：calc(“2”,0,2)，calc(“1”,1,3)，则就依次输出：2 0 2 2，1 1 3 4。
	*/

}

func calc(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}

type Person struct {
	age int
}

func test7() {
	person := &Person{28}
	// 1.
	defer fmt.Println(person.age)
	// 2.
	defer func(p *Person) {
		fmt.Println(p.age)
	}(person)
	// 3.
	defer func() {
		fmt.Println(person.age)
	}()
	person = &Person{29}
}

func f1() (r int) {
	defer func() {
		r++
	}()
	return 0
}
func f2() (r int) {
	t := 5
	defer func() {
		t = t + 5
	}()
	return t
}
func f3() (r int) {
	defer func(r int) {
		r = r + 5
	}(r)
	return 1
}

func test8() {
	fmt.Println(f1())
	fmt.Println(f2())
	fmt.Println(f3())
}

func increaseA() int {
	var i int
	defer func() {
		i++
	}()
	return i
}
func increaseB() (r int) {
	defer func() {
		r++
	}()
	return r
}
func test9() {
	fmt.Println(increaseA())
	fmt.Println(increaseB())
}

func main() {
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
