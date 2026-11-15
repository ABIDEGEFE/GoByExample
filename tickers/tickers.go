package tickers

import (
	"fmt"
	"time"
)

func GoTickers(){
	fmt.Println("Go ticker is running")
}

func TickerCaller(){
	// ticker := time.NewTicker(500*time.Millisecond)
	done := make(chan bool, 2)
	// go func(){
    //    for{
	// 	select{
	// 	case t :=<- ticker.C:   //This will receive values with specified specific interval until ticker.Stop() executed
	// 		fmt.Println("Ticker at", t)
	// 	case <-done:
	// 		return
	// 	}
	//    }
	// }()
	// time.Sleep(1600*time.Millisecond)
	// ticker.Stop() // This will close the ticker channel, and the channel will stop reciving values.
	// done <- true
	// fmt.Println("Ticker stopped")

	imageResizingTicker := time.NewTicker(5*time.Second)
	processedImage := 0
	go PhotoResizingService(processedImage, imageResizingTicker.C, done)
	time.Sleep(15*time.Second)
	imageResizingTicker.Stop()
	done <- true
	time.Sleep(time.Second)
}
// ======================================= Scenario ====================================

//Imagine our photo-resizing service runs continuously. 
//You want it to report the number of processed images every 5 seconds, rather than report only once.
func PhotoResizingService(processedImage int, tickerChan <-chan time.Time, done <-chan bool){
	// Assuming the program should report processed images within 15 seconds.
   for{
	select{
	case t := <-tickerChan:
		processedImage += 10
		fmt.Println("Processed Image:",processedImage, "at", t)
    case <-done:
		fmt.Println("Photo resizing service is done")
	return
	}
   }
}