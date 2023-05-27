package main

import(
	"fmt"
	"time"
)
//? only one value stored in channel of certain type and until it is unloaded, it wait like that
func main(){
	c:= make(chan int)
	
	go func(){
		for i:= 0; i<20 ; i++ {
			//? load a number and pauses until its unloaded
			c <- i
		}
	}()

	go func(){
		for{
			//? unloads a number as its running as routine along with above func, another number willl be added
			fmt.Println(<-c)
		}
	}()

	time.Sleep(time.Second)
}