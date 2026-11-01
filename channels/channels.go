package channels

import (
	"fmt"
	"time"
)

func GoChannel(){
	fmt.Println("Go channel is running")
}

func ChannelCaller(){
	messages := make(chan Order) //creating a channel through which string value would be passed
    // message := <- messages //recieve channel will block when the channel is empty

	// go func(){
	// 	messages <- "ping"
	// }()
	// go func(){
	// 	msg := <- messages
	// 	fmt.Println("Received", msg)
	// }()
	go RecieveOrder(messages)
	go SendConfirmationEmail(messages)
	
	time.Sleep(time.Second)
}

// =============================== Processing Online Orders ================================
/*
Imagine you are building a small e-commerce backend. Customers place orders quickly, but sending confirmation emails may take time.

Use a channel to let two goroutines work independently:

    1. The order goroutine receives and validates orders. It sends each confirmed order into a channel.

    2. The email goroutine receives orders from that channel and sends confirmation emails.

A Go channel is a typed path for safely passing values between goroutines, and it also coordinates their work.
*/

type Order struct{
	email string
	ID string
}

func RecieveOrder(ch chan Order){
	//Assuming order have been received here
	ch <- Order{email:"abc@gmail.com", ID:"1001"} 
}
func SendConfirmationEmail(ch chan Order){
	order := <- ch
    fmt.Println("Confirmation email", order.email, "With Id", order.ID)
}