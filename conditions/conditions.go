package conditions

import "fmt"

func GoIfStatement(){
	num := 10
	if num < 20 {
		fmt.Println("It is less than 20")
	} else if num % 2 == 0 {
		fmt.Println("It is divisible by 2")
	} else if num % 5 == 0 {
		fmt.Println("It is divisible by 5")
	} else {
		fmt.Println("What is going on.")
	}

}

func GoSwitchStatement(){
	num := 40
	switch{  // without expression, switch used as if/else condition
	case num < 50:
		fmt.Println("It is 30")
		// break
	case num < 45:
		fmt.Println("It is 50")
		// break
	case num > 50:
		fmt.Println("It is 40")
		break
	case num > 80:
		fmt.Println("It is 20")
		break
	default:
		fmt.Println("It is not number")
	}

	WhatAmI := func(val any) {
		switch val.(type) {
		case int:
			fmt.Println("I am Integer")
			break
		case bool:
			fmt.Println("I am boolean")
			break
		case string:
			fmt.Println("I am string")
			break
		default:
			fmt.Println("I am default")
		}
	}

	WhatAmI("string")
	WhatAmI(5)
	WhatAmI(true)
}