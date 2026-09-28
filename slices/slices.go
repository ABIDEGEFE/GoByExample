package slices

import "fmt"


func GoSlices(){
	var s = []int{1, 2, 3, 4}  // s is not the actual array, it is the header which is struct containing:
	                           // pointer which points to the underlying fixed size array,
							   // length of the array,
							   // capacity of the array. it is like s = {pointer, len, cap}

							   // if len < cap && append:
							   //    no re-allocation, fast growth of underlying array, change len
							   // else
							   //    - new memory allocation, copy data over new array, change pointer, double capacity
							
	fmt.Println(s == nil)
	fmt.Println(s, len(s), cap(s))  // length and capacity will be the same by default.
	fmt.Printf("Array address: %p\n", s)
	s = append(s, 4)
	fmt.Println("appended", len(s), cap(s))
	fmt.Printf("Array address: %p\n", s)

	// we can use the make key word to control over the capacity of the array
	a := make([]int, 3, 5)   // slice with length 3 and capacity 5
	fmt.Println(a, len(a), cap(a))
	fmt.Printf("Array address: %p\n", a)
	a = append(a, 50)
	fmt.Println(a, len(a), cap(a))
	fmt.Printf("Array address: %p\n", a)

	c := make([]int, 1)
	copy(c, a)   // it is possible to copy the data of the slice into other slice with the same or different lenght.
	fmt.Println(c)
	fmt.Printf("Array address: %p\n", c)

	// we can get some continous part of slices 
	sl := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Printf("Array address of sl: %p\n", &sl[0])
	fmt.Printf("Array address of sl: %p\n", &sl[1])
	fmt.Printf("Array address of sl: %p\n", &sl[4])
	sl1 := sl[2:5] // this will slice starting from index 2 up to 5 but excluding index 5
	fmt.Printf("Array address sl1: %p\n", &sl1[2])
	fmt.Println(sl1)

	// Two dimenssion arrays
	twoD := make([][]int, 3)
	fmt.Println(twoD)  //[[], [], []]
}