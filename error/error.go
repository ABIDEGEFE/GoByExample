package error

import (
	"fmt"
	"errors"
)


func GoError(){
	fmt.Println("Go error is running")
}

func f(arg int)(int, error){
	if arg == 42{
		return -1, errors.New("It does work with 42")
	}
	return arg + 3, nil
}

var ErrorOfTea = errors.New("Tea is not supported")
var ErrorOfPower = errors.New("Power is low, could not boil Tea")

func makeTea(n int)error{
	if n == 3{
		return ErrorOfPower
	}else{
		return ErrorOfTea
	}
}

func ErrorCaller(){
	for i := 40; i < 43; i++{
		f(i)
	}
	for j := range 5{
		if err := makeTea(j); err != nil{
			if errors.Is(err, ErrorOfTea){
				fmt.Println("Error:", ErrorOfTea)
			}else if errors.Is(err, ErrorOfPower){
				fmt.Println("Error:", ErrorOfPower)
			}else{
				fmt.Println("Unkown error")
			}
		}
	}

	// workign with custom error
    _, err := f(42)
    if ae, ok := errors.AsType[*argError](err); ok {
        fmt.Println(ae.arg)
        fmt.Println(ae.message)
    } else {
        fmt.Println("err doesn't match argError")
    }

}

// defining custom error
type argError struct{
	arg int
	message string
}

func (e *argError)Error()string{
	return fmt.Sprintf("%d - %s", e.arg, e.message)
}

func fc(arg int)(int, error){
	if arg == 42{
		return -1, &argError{arg, "Can not work with it"}
	}
	return arg+3, nil
}