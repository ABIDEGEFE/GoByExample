package rateLimiter


import (
	"fmt"
	"time"
	"sync"
)

func GoRateLimiter(){
	fmt.Println("Go rate limiter is running.")
}

var smsRequest = make(chan string, 5)

// ========================== Scenario: SENDING SMS DELIVERY NOTIFICATION ============================
// Imagine you are building an online store. 
// After dispatching orders, your application must send 1000 (use 5 for simulation) customers an SMS.
// Assume your SMS provider requires you to space requests at least 1000 milliseconds apart. 
// Sending everything immediately could exceed its limit. 
// Rate limiting controls how quickly your application calls a service to stay within its capacity.
func GoRateLimiterCaller(){
	var wg sync.WaitGroup
	limiter := time.Tick(1000*time.Millisecond)

	customerPhone := []string{"099-333-33", "099-44-44", "099-11-11", "099-999-999", "099-666-88"}
    wg.Go(func(){
		for req := range smsRequest{
			fmt.Println("SMS request is sent for", req)
		}
	})
    for i:=0; i<5; i++{
		<-limiter
        smsRequest <- customerPhone[i]
	}
	close(smsRequest)
	wg.Wait()
}
