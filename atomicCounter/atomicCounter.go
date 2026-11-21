package atomicCounter


import (
	"fmt"
	"sync"
	"sync/atomic"
)

func GoAtomicCounter(){
	fmt.Println("Go atomic counter is running")
}


func AtomicCounterCaller(){
	var ops atomic.Uint64
	var wg sync.WaitGroup
	for i:=1; i<=50; i++{
        wg.Go(func(){
			for range 1000{
				ops.Add(1)
			}
	   })
	}

	wg.Wait()
	fmt.Println("Final value of ops", ops.Load())
}