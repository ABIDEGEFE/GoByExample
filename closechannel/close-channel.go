package closechannel


import (
	"fmt"
	// "time"
)

func GoClosingChannel(){
	fmt.Println("Go closing channel is running..")
}

func ClosingChannelCaller(){
	image := make(chan string)
	go UploadPhoto(image)
	// An empty channel alone cannot tell the resizer whether more photos are coming. 
	// Closing it communicates “no more photos will be sent.”
	for images := range image{
		fmt.Println("Photo resized", images)
	}
	fmt.Println("All images processed")
}
func UploadPhoto(image chan<- string){
	//Assuming the user is only able to upload 3 images
	for i := range 3{
		image <- "image.jpg"
		fmt.Println("Image", i, "Sent")
	}
	close(image) //Only the sender is able to close channel. If the receiver tries to close, run-time panice will appear.
	// image <- "finished"  - sending on closed channel will create an error
}