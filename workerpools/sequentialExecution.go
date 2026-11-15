package workerpools

import (
	"fmt"
	"time"
)

func GoWorkerPools(){
	fmt.Println("Worker pools is running..")
}
var result = make(chan int, 5)
func WorkerPoolsCaller(){
   start := time.Now()
   defer func(){
	fmt.Println("Time taken by the program", time.Since(start))
   }()
   JobSender()
   for r := 1; r<=5; r++{
	res := <-result
	fmt.Println("result", res)
   }
}

func JobSender(){
   numJob := 5
   jobs := make(chan int, numJob)
   for j := 1; j<=numJob; j++{
	jobs<-j
   }
   close(jobs)
   JobReciever(1, jobs, result)
}
func JobReciever(id int, jobs <-chan int, result chan<- int){
   for job := range jobs{
	fmt.Println("Worker", id, "Started job", job)
    time.Sleep(1*time.Second)
	fmt.Println("Worker", id, "Finished job", job)
	result<-job*1000
   }
}