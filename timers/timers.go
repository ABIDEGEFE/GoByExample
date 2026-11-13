package timers


import (
	"fmt"
	"time"
)

func GoTimers(){
	fmt.Println("Go timer is running..")
}

func TimerCaller(){
	time1 := time.NewTimer(time.Second)
	
	// time1.Stop()
	go func(){
		fmt.Println("Before timer")
		// If timer.Stop() get executed first the value will never be sent through the channel.
		// So this channel blocks for ever.
		// This channel only get fired when its time is  reached.
       <-time1.C 
	   fmt.Println("timer under goroutine")
	}()
	// It does not close the channel but it prevents any read from it.
	// It retuns true if the timer is not fired yet, and false if the timer is already fired.
	time1.Stop() 
	time.Sleep(2*time.Second)
}