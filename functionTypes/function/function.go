package function

import (
	"fmt"
	"time"
	"strconv"
)
// a function with no args and return value
func GoFunction(){
	fmt.Println("This is Go's Function.")
}
// a function can accept and process no or many args and return the output.
func SetAge(birthdate string) int{
	year := birthdate[len(birthdate)-4:] // assumign the year is last 4 characteres like dd/mm/yyyy
	birthYear, err := strconv.Atoi(year)
	if err == nil{
        return time.Now().Year() - birthYear
	}
	return 0
}

// if a function is going to accept the same type of many args, we can omit the data type unitl the last arg.
func Sum(num1, num2, num3, num4 int) int {
    return num1 + num2 + num3 + num4
}

// Go allows to return multiple values
func ReturnMulitpleValues(name string) (int, error){
	val, err := strconv.Atoi(name)
	return val, err
}
// There is a function called variadic function we should use when we are uncertain about the number of arguments. It allows us any number of args.
func VariadicFunc(nums ...int) int {
	sum := 0
	for v := range nums{
        sum += nums[v]
	}
	return sum
}
// we can also pass args with different data types
func PrintAny(data ...any){   // equivalent of fmt.Println()
	fmt.Println(data)
}

// A function closes a variable with its state is called closure function.
func ClosureFunc() func() int {
	count := 0
	return func() int{
		count ++
		return count
	}
}

func SampleFunc() int{
	var num1, num2 int = 4, 5
	var sum int = num1 + num2
	return sum
}