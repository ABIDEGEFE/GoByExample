package selects


import (
	"fmt"
	"time"
)


func GoSelects(){
	fmt.Println("Go select is running")
}

func PrintFirstName(firstname chan<- string){
	time.Sleep(5*time.Second)
	firstname <- "Jhone"
}
func PrintSecondName(secondname chan<-string){
	time.Sleep(20*time.Second)
	secondname <- "Doe"
}

func SelectCaller(){
	firstname := make(chan string)
	secondname := make(chan string)
    
	go PrintFirstName(firstname)
	go PrintSecondName(secondname)

	// fn := <- firstname
	// fmt.Println("Frist name", fn)
	// sn := <- secondname
	// fmt.Println("Second name", sn)

	// select lets a goroutine wait for either channel operation, instead of getting stuck waiting on just one.


	for range 2{
     select{
		case fn := <- firstname:
		   fmt.Println("First Name:", fn)
	    case sn := <- secondname:
		   fmt.Println("Second Name:", sn)
		case <- time.After(3 * time.Second):
			fmt.Println("Time out")
	 }
	}
}
