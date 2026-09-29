package maps

import "fmt"


func GoMaps(){
	// Declaring empty map using make
	mp := make(map[string]int)
	fmt.Println("empty map", mp)
    // we can add key and value using the following syntax
	mp["one"] = 1
	mp["two"] = 2
	fmt.Println(mp)
    // we can access values using their corresponding key
	fmt.Println(mp["one"]) // if the key does not exist, zero value of they type will be returned.
	// we can use the 'len' keyword to return the number of key-value pair elements
	fmt.Println(len(mp)) // 2
	// 'delete' built-in helps to remove key-value pair from the data structure. 'clear' built-in will remove all elements 
	delete(mp, "one")
	fmt.Println(mp) // map["two":2]

	num, prs := mp["two"] // the second optional return value helps to determine whether the key existed or not by returning boolean value.
	if prs {
		fmt.Println("The key existed", num)
	} else{
		fmt.Println("The key does not exist")
	}

	// Declaring and initalizing the map at the same line
	mp2 := map[string]bool{"Status":true}
	fmt.Println(mp2)
	mp["two"] = 5
	fmt.Println(mp)
}