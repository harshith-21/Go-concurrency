package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	// fmt.Println("Hello thereee !!!!")
	//? no concurr
	// foo()
	// bar()

	// //? concurr
	// go foo()
	// go bar()

	//? concur with wg
	wg.Add(2)
	go foo()
	go bar()
	wg.Wait()
}

func foo() {
	for i:=0; i<10; i++ {
		fmt.Println("foo: ", i)
		time.Sleep(time.Duration(3 * time.Millisecond))
	}
	wg.Done()
}

func bar() {
	for i:=0; i<10; i++ {
		fmt.Println("bar: ", i)
		// time.Sleep(time.Duration(3 * time.Millisecond))
		time.Sleep(time.Duration(10 * time.Millisecond))
	}
	wg.Done()
}
