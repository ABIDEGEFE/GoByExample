package statefulGoroutines

import (
	"fmt"
	"math/rand"
	"sync/atomic"
	"sync"
	// "time"
)

func GoStatefulGoroutines(){
	fmt.Println("Go statefulGoroutines are running..")
}

type readOp struct{
	key int
	resp chan int
}
type writeOp struct{
	key int
	val int
	resp chan bool
}

func StatefulGoroutines(){
   var wg sync.WaitGroup
   reads := make(chan readOp)
   writes := make(chan writeOp)

   var breakCount atomic.Int64

   wg.Go(func(){
 	state := make(map[int]int)
	for{
		// fmt.Println("##############################", breakCount.Load())
		if breakCount.Load() == 110{
			return
		}
		select{
		case read :=<- reads:
			read.resp <- state[read.key]
		case write :=<- writes:
			state[write.key] = write.val
			write.resp <- true
		}
	}
   })

   

   

   for range 100{
	wg.Go(func(){
		for range 10{
			read := readOp{
              key: rand.Intn(5),
			  resp: make(chan int),
			}
			reads <- read
			<-read.resp
			// recievedKey :=<- read.resp
			// fmt.Println("You have recieved key", <-read.resp)
			// breakCount<-true
		}
        breakCount.Add(1)
	})
   }

   for range 10{
	wg.Go(func(){
		for range 10{
			write := writeOp{
				key: rand.Intn(5),
				val: rand.Intn(5),
				resp: make(chan bool),
			}
			writes <- write
			<-write.resp
			// recievedWriteResponse :=<- write.resp
			// fmt.Println("You have received write response", <-write.resp)
			// breakCount<-true
		}
		breakCount.Add(1)
	})
   }
   wg.Wait()
   fmt.Println("State of the program has finished succesfully", breakCount.Load())
}