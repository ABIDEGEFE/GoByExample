package array

import "fmt"

func GoArray(){
	var a [4]int
	fmt.Println(len(a)) // if we don't initialize an array, it would be filled with zero-valued.

	a[2] = 100
	// a = append(a, 4000)   we can not modify the size of an array.
	fmt.Println(a)
    
	var b = [4]int{1, 2, 3, 4} // decalring and initalizing an array at one line
	fmt.Println(b)

	var c = [5]bool{true, 3: true, true}
	fmt.Println(c)
}