package timeouts

import (
	"fmt"
	"time"
)

func GoTimeOuts(){
	fmt.Println("Go timeout is running")
}

func TimeoutCaller(){
	ch1 := make(chan string)
	
	// Assuming this is a function which has a task which might take longer time than expected. Ex: Database access
	go func(){
		time.Sleep(20*time.Second)
		ch1 <- "hello"
	}()

	// select helps the program to work with ready channel at a time.
	select {
	case val := <- ch1:
		fmt.Println("Data is received", val)
	// if the above recieve channel is not ready within 2 seconds, no need to wait for it.
	// Other part of the program should be executed instead of waiting for tasks with unexpected behaviour.
	case <-time.After(2*time.Second): 
		fmt.Println("Time out")
	}
	fmt.Println("Starting execution of other instructions")
	// val := <- ch1
	// fmt.Println("Data is recieved", val)

}