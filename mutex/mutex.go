package mutex

import (
	"fmt"
	"sync"
)

func GoMutext(){
	fmt.Println("Go mutex is running..")
}

type Contanier struct{
   mu sync.Mutex
   counter map[string]int
}

func (c *Contanier)inc(name string){
	c.mu.Lock()
	defer c.mu.Unlock()
	c.counter[name]++
}

func (c *Contanier)dec(name string){
	c.mu.Lock()
	defer c.mu.Unlock()
	c.counter[name]--
}
func MutexCaller(){
   var wg sync.WaitGroup
   con := Contanier{
	counter: map[string]int{
		"a": 0,
		"b": 0,
	},
   }
   doIncrement := func(name string, n int){
     for range n{
		con.inc(name)
	 }
   }
   doDecreament := func(name string, n int){
	for range n{
		con.dec(name)
	}
   }
   
   wg.Go(func(){
	fmt.Println("I am decrement")
	doDecreament("a", 3000)
   })
   wg.Go(func(){
	fmt.Println("I am increment 1")
	doIncrement("a", 5000)
   })
   wg.Go(func(){
	fmt.Println("I am increment 2")
	doIncrement("b", 1000)
   })
   wg.Wait()
   fmt.Println("Final value", con.counter)
}