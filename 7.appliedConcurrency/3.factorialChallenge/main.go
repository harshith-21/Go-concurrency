package main

import (
	"fmt"
)

//? using wait groups

// func main() {
// 	f := factorial(4)
// 	fmt.Println("factorial =", f)
// }

// func factorial(n int) int {
// 	ch := make(chan int)
// 	var wg sync.WaitGroup
// 	wg.Add(1)

// 	go func(){
// 		for i:=1;i<=n;i++ {
// 			ch <- i
// 		}
// 		close(ch)
// 	}()

// 	var ans int = 1
// 	go func() {
// 		for i:=range ch {
// 			ans *= i
// 		}
// 		wg.Done()
// 	}()
// 	wg.Wait()
// 	return ans
// }

//? other way

func main() {
	c := factorial(4)
	for n := range c {
		fmt.Println(n)
	}
}

func factorial(n int) chan int {
	out := make(chan int)
	go func () {
		total := 1
		for i:=n;i>0;i-- {
			total *= i
		}
		out <- total
		close(out)
	}()
	return out;
}