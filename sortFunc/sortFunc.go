package sortFunc

import (
	"fmt"
	"slices"
	"cmp"
)

func GoSortFunc(){
	fmt.Println("Sort function is running..")
}

func SortFuncCaller(){
   fruits := []string{"banana", "mango", "kiwi", "orange", "grapes", "Avocado"}
   slices.SortFunc(fruits,
   func(a, b string)int{
	fmt.Println("Value of a", a, "Value of b", b)
	return cmp.Compare(len(a), len(b))
   },
)
  type person struct{
	name string
	age int
  }

  people := []person{
	{name: "Teshome", age:59},
	{name: "Teskl", age:99},
	{name: "Luku", age:88},
  }
  slices.SortFunc(people,
	func(a, b person)int{
		return cmp.Compare(a.age, b.age)
	},
  )
  fmt.Println("Sorted people by age", people)
  fmt.Println("Sorted fruits", fruits)
}