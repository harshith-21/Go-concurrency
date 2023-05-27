package main

import(
	"fmt"
)

// func main() {
// 	for n:= range sq(sq(gen(2, 3))) {
// 		fmt.Println(n)
// 	}
// }

// func gen(nums ...int) chan int {
// 	out := make(chan int)
// 	go func(){
// 		for _, n := range nums {
// 			out <- n
// 		}
// 		close(out)
// 	}()
// 	return out
// }

// func sq(in chan int) chan int {
// 	out := make(chan int)
// 	go func() {
// 		for n := range in {
// 			out <- n*n
// 		}
// 		close(out)
// 	}()
// 	return out
// }


// func main() {
// 	c := factorial(4)
// 	for n := range c {
// 		fmt.Println(n)
// 	}
// }

// func factorial(n int) chan int {
// 	out := make(chan int)
// 	go func() {
// 		total := 1
// 		for i:=n;i>0;i-- {
// 			total *= i
// 		}
// 		out <- total
// 		close(out)
// 	}()
// 	return out
// }

//? doing these 100 calculations parallely


func main() {
	in := gen()
	out := factorial(in)
	for i:= range out {
		fmt.Println(i)
	} 
}

func gen() chan int {
	out := make(chan int)
	go func(){
		for i:=0;i<10;i++ {
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