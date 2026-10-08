package pointers

import "fmt"


func ZeroPtr(num *int){
	fmt.Println("Memory address", num) // it holds the memory address 
	fmt.Println("Actual Data addressed by", num, "is", *num) // dereferencing is accessing the actual data using its memory address
	ChangeActualData(num) // passing the memory address of num to the function so it would have full access over this data
}

func ChangeActualData(d *int){
	*d = 88
}

func IntVal(val int){
	val = 50 // val has been assigned with its own dedicated memory address
}
func GoPointers(){
	// pointers are created to store the memory address of a data. when we give access for a function to work with 
	// the pointer we are giving the full access fo the data for that function. 
	n := 1
	ZeroPtr(&n)
	fmt.Println("changed by ChangeActualData function", n)
	v := 8
	IntVal(v) // v and val(parameter of IntVal) has been assigned with different memory address.
	          // changing the value of either of the variable could not affect the other one.
    fmt.Println("V value", v) // It is still 8

}