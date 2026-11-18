package waitgroup


import (
	"fmt"
	"sync"
	"time"
)

func GoWaitGroup(){
	fmt.Println("Go wait group is running..")
}

func Worker(i int){
	fmt.Println("Worker", i, "starting")
	time.Sleep(time.Second)
	fmt.Println("Worker", i, "done")
}

func Printer(name string){
	fmt.Println("Printer started")
	time.Sleep(5*time.Second)
	fmt.Println("Printer finished")
}

func WaitGroupCaller(){
    var wg sync.WaitGroup

	wg.Go(func(){
		Printer("Marta") // Printer is not a goroutine, it is a function executed under the launched goroutine with wg.Go
	})

	wg.Go(func(){
		Worker(1) // Worker is not a goroutine, it is a function executed under the launched goroutine with wg.Go
	})
	wg.Wait() // This will block the caller to wait until all goroutines launched under wg finish their execution
}