package atomicCounter

import (
	"fmt"
	"sync"
)

func CounterWithoutAtomic(){
    counter := 0
	var wg sync.WaitGroup

	for i:=1; i<=50; i++{
		wg.Go(func(){
			for range 1000{
				counter++
			}
		})
	}
	wg.Wait()
	fmt.Println("Final value of counter", counter)
}