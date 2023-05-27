package main

import (
	"fmt"
)

func main() {
	in := gen()

	xc := fanOut(in, 10)

	for n :=  range merge(xc...) {
		fmt.Println(n)
	}
}

func gen() <-chan int {
	out := make(chan int)
	go func() {
		for i:=0;i<10;i++ {
			for j:=3;j<13;j++ {
				out <- j
			}
		}
		close(out)
	}()
	return out
}

