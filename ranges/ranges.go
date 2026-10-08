package ranges


import "fmt"

func GoRange(){
	// range iterate over elements of slices
	nums := []int{1, 2, 3, 4}
	for n, v := range nums{ // range helps us to iterate over elements in slice accessing both the index and the actual value
		fmt.Println("vlaue at index", n,  "is :", v)
	}

	// range iterates over maps accessing both keys and values
	alphabet := make(map[string]string)
	alphabet["A"] = "a"
	alphabet["B"] = "b"
	alphabet["C"] = "c"
    for k, v := range alphabet{
		fmt.Println("Key :", k, "Value :", v)
	}

	// range iterates over strings accessing index of character and thier unicode value. we can use built-in string function
	// to get their equivalent character
	name := "Amin"
	for _, c := range name{
		fmt.Println("Name", string(c))
	}
}