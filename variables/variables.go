package variables

import "fmt"

func GoVariables(){

	// Initializing a variable with out a value is considered as zero value
	var age int
	fmt.Println(age)

	// we can declare and initialize mulitple variables at once using 'var'
	var name, fatherName = "Jhone", "Kemal"
	fmt.Println(name, fatherName)

	// := is shorthand represenation for declaring and initializing a variable
	name2 := "Thomas" // instead of var name string = "Thomas"
	fmt.Println(name2)

}