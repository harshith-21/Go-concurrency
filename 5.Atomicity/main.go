package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

var wg sync.WaitGroup
// var counter int
var mutex sync.Mutex
var counter int64

func main(){
	wg.Add(2)
	go incrementor("Foo: ")
	go incrementor("bar: ")
	wg.Wait()
	fmt.Println("Final Counter: ", counter)
}

func incrementor(s string){
	for i:= 0;i<20;i++ {
		time.Sleep(time.Duration(rand.Intn(20)) * time.Millisecond)
		//? race:
		// counter++;
		//? no race
		atomic.AddInt64(&counter, 1)
		fmt.Println(s, i, "Counter: ", counter)
	}
	wg.Done()
}