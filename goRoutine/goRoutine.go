package goRoutine


import (
	"fmt"
	"time"
)

//assuming GoRoutine's task may take serveral hours (5 hr)
//since it is not goroutine, it will block other threads and make the program to freeze
func GoRoutine(){
	fmt.Println("Goroutine is running")
	// time.Sleep(5000 * time.Second)
}

func f(from string){
	time.Sleep(1000*time.Second)
	for i := range 3{
		fmt.Println(from, ":", i)
	}
}

func GoroutineCaller(){
	//this is blocking function which run syncronously
	GoRoutine()
	//The following two goroutines are able to be executed concurrently,a nd run asyncronously.
	//goroutines allows the CPU to execute mutliple threads concurrently to enhance the performance of the program
	//goroutines has light weight which are around 2 KB initially, grows or shrinks dynamically.
	go f("goroutine")
	go func(msg string){
        fmt.Println(msg)
	}("hello")

	//we should keep the caller function to wait for all goroutines until they return or exit.
	//if the caller finishes its execution before goroutines inside it, all goroutines will be killed regardless of 
	//their status.
	time.Sleep(2000*time.Second)
}