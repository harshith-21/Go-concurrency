/*
	* The optional <- operator specifies rthew channel direction, send or recieve
	* If no diection is given, the channel is bidirectional,
*/

package main

import(
	"fmt"
)

func main() {
	c := incrementor()
	cSum := puller(c)
	for n:= range cSum {
		fmt.Println(n)
	}
}

func incrementor() <-chan int {
	out := make(chan int)
	go func(){
		for i:=0;i<10;i++ {
			out <- i
		}
		close(out)
	}()
	return out
}

func puller(c <-chan int) <-chan int {
	out := make(chan int)
	go func(){
		var sum int
		for n:= range c {
			sum += n
		}
		out <- sum
		close(out)
	}()
	return out
}