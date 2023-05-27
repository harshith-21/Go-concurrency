package main

import (
	"fmt"
	"sync"
)

//? NORMAL

// func main() {
// 	in := gen()
// 	out := factorial(in)
// 	for i := range out {
// 		fmt.Println(i)
// 	}
// }

// func gen() chan int {
// 	out := make(chan int)
// 	go func() {
// 		for i := 0; i < 100; i++ {
// 			for j := 10; j < 20; j++ {
// 				out <- j
// 			}
// 		}
// 		close(out)
// 	}()
// 	return out
// }

// func factorial(c chan int) chan int {
// 	out := make(chan int)
// 	go func() {
// 		for i := range c {
// 			out <- fact(i)
// 		}
// 		close(out)
// 	}()
// 	return out
// }

// func fact(n int) int {
// 	a := 1
// 	for i := 1; i <= n; i++ {
// 		a *= i
// 	}
// 	return a
// }

//?  USING FAN IN AND FAN OUT PATTERN

func main() {
	in := gen()

	//? fan out
	c0 := factorial(in)
	c1 := factorial(in)
	c2 := factorial(in)
	c3 := factorial(in)
	c4 := factorial(in)
	c5 := factorial(in)
	c6 := factorial(in)
	c7 := factorial(in)
	c8 := factorial(in)
	c9 := factorial(in)

	//? fan in
	for n := range merge(c0, c1, c2, c3, c4, c5, c6, c7, c8, c9) {
		fmt.Println(n)
	}
}

func gen() chan int {
	out := make(chan int)
	go func(){
		for i:=0;i<100;i++ {
			for j:=10;j<20;j++ {
				out <- j
			}
		}
		close(out)
	}()
	return out;
}

func factorial(c chan int) chan int{
	out := make(chan int)
	go func(){
		for i := range c {
			out <- fact(i)
		}
		close(out)
	}()
	return out
}

func fact(n int) int {
	a :=1;
	for i:=1;i<=n;i++ {
		a *= i
	}
	return a
}

func merge(cs ...chan int) chan int {
	out := make(chan int)
	var wg sync.WaitGroup

	wg.Add(len(cs))

	for _, c := range cs {
		go func(ch chan int){
			for n:= range ch {
				out <- n
			}
			wg.Done()
		}(c)
	}

	go func(){
		wg.Wait()
		close(out)
	}()

	return out
}