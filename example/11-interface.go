package main

type S struct{}

func f(x interface{}) {

}

func g(x *interface{}) {

}

func test1() {
	s := S{}
	p := &s
	f(s) //A
	g(s) //B
	f(p) //C
	g(p) //D
}

func main() {
	test1()
}
