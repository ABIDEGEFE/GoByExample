package bufferedChannel

import(
	"fmt"
	"time"
)

func GoBufferedChannel(){
	fmt.Println("Buffered channel is running")
}

func BufferedChanelCaller(){
	// messages := make(chan string, 2) //Capability to execute two send channels without knowing the status of the receiver.
	// messages <- "one"
	// messages <- "two"
	// // messages <- "three" // This will create deadlock since the channel has finished its potential to send before calling the receiver.
	// fmt.Println(<-messages)
	// fmt.Println(<-messages)

	upload := make(chan string, 3)
	go UploadImage(upload)
	go ResizeImage(upload)
	time.Sleep(time.Second)
}

// ========================= Scenario: Uploading images, and resizing the image slowly ======================

func UploadImage(ch chan string){
	/*
	- Imagine users can upload several profile photos quickly, but image resizing is slower.
    - Use a buffered channel as a small waiting queue between the uploader and the image-processing worker
	These sends can happen even if the image-resizing goroutine has not started receiving yet.
	The channel temporarily stores the jobs.
	*/
	fmt.Println("Sender")
	ch <- "image1.jpg"
	ch <- "image2.jpg"
	ch <- "image3.jpg"
	// ch <- "image3.jpg" // This will create deadlock if UploadImage and ResizeImage executed syncrounously.
}

func ResizeImage(ch chan string){
	// for image := range ch{
	// 	fmt.Println("Resize", image)
	// }
	fmt.Println("Reciever")
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	// fmt.Println(<-ch)
}