package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func test1() {
	wg := sync.WaitGroup{}
	for i := 0; i < 5; i++ {
		go func(wg sync.WaitGroup, i int) {
			wg.Add(1)
			fmt.Printf("i:%d\n", i)
			wg.Done()
		}(wg, i)
	}
	wg.Wait()
	fmt.Println("exit")
}

func test2() {
	runtime.GOMAXPROCS(1)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(i)
		}
	}()

	for {
	}
}

func test3() {
	ch := make(chan int, 100)
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
	}()

	go func() {
		for {
			a, ok := <-ch
			if !ok {
				fmt.Println("close")
				return
			}
			fmt.Println("a: ", a)
		}
	}()

	close(ch)
	fmt.Println("ok")
	time.Sleep(time.Second * 10)
}

func test4() {
	var ch chan int

	go func() {
		ch = make(chan int, 1)
		ch <- 1
	}()
	go func(ch chan int) {
		time.Sleep(time.Second)
		<-ch
	}(ch)

	c := time.Tick(1 * time.Second)
	for range c {
		fmt.Printf("#goroutines: %d\n", runtime.NumGoroutine())
	}
}

func test5() {
	var wg sync.WaitGroup
	wg.Add(2)

	var ints = make([]int, 0, 1000)
	go func() {
		for i := 0; i < 1000; i++ {
			ints = append(ints, i)
		}
		wg.Done()
	}()

	go func() {
		for i := 0; i < 1000; i++ {
			ints = append(ints, i)
		}
		wg.Done()
	}()

	wg.Wait()
	fmt.Println(len(ints))
}

type T struct {
	V int
}

func (t *T) Incr(wg *sync.WaitGroup) {
	t.V++
	wg.Done()
}

func (t *T) Print() {
	time.Sleep(1)
	fmt.Println(t.V)
}

func test6() {
	var wg sync.WaitGroup
	wg.Add(10)

	var ts = make([]T, 10)

	for i := 0; i < 10; i++ {
		ts[i] = T{i}
	}
	fmt.Println("ts --> ", ts)

	for _, t := range ts {
		go t.Incr(&wg)
	}
	wg.Wait()

	for _, t := range ts {
		go t.Print()
	}
	time.Sleep(5 * time.Second)
}

func test7() {
	const N = 26
	const GOMAXPROCS = 1
	runtime.GOMAXPROCS(GOMAXPROCS)

	var wg sync.WaitGroup
	wg.Add(2 * N)
	for i := 0; i < N; i++ {
		go func(i int) {
			defer wg.Done()
			runtime.Gosched()
			fmt.Printf("%c", 'a'+i)
		}(i)
		go func(i int) {
			defer wg.Done()
			fmt.Printf("%c", 'A'+i)
		}(i)
	}
	wg.Wait()
}

func main() {
	test1()
	test2()
	test3()
	test4()
	test5()
	test6()
	test7()
}
