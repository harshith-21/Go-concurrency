package main

import(
	"fmt"
)

//* 1
//? wrong bec putting into the channel and taking off should be done parallely
//? eg:  with go routines 

// func main() {
// 	c := make(chan int)
// 	c<-1
// 	fmt.Println(<-c)
// }




//? correct way:

// func main() {
// 	c := make(chan int) 
// 	go func(){
// 		c<-1
// 	}()
// 	fmt.Println(<-c)
// }


//* 2

//? prints only 0 where are other numbers

// func main() {
// 	c := make(chan int)
// 	go func() {
// 		for i:=0;i<10;i++ {
// 			c <- i
// 		}
// 	}()

// 	fmt.Println(<-c)
// }


//? correct way

func main() {
	c := make(chan int)
	go func(){
		for i:=0;i<10;i++ {
			c <- i
		}
		close(c)
	}()
	for i := range c {
		fmt.Println(i)
	}
}

