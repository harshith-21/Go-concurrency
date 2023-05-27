package main

import(
	"fmt"
)

func main() {
	c1 := incrementor("foo:") 
	c2 := incrementor("bar:")
	c3 := puller(c1)
	c4 := puller(c2)
	fmt.Println("dunmb sum: ", <-c3 + <- c4)

}

func incrementor(s string) chan int {
	out := make(chan int)
	go func(){
		for i:=0;i<20;i++ {
			out <- i
		}
		close(out)
	}()
	return out
} 

func puller(c chan int) chan int{
	out := make(chan int)
	go func(){
		var sum int
		for i:= range c {
			sum += i
		}
		out <- sum
		close(out)
	}()
	return out
}