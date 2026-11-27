package numberParsing

import (
	"fmt"
	"strconv"
)

func GoNumberParsing(){
	fmt.Println("Go number parsing is running.")
}

func NumberParsingCaller(){
	f, _ := strconv.ParseFloat("-44", 64)  // the second argument determines how much precision the number rounds to.
	fmt.Println(f)
	// 0 determines the numbering system is base 0. If we change it to 16 we will get hex value of 1823 which would be stored in int64
	i, _ := strconv.ParseInt("1823", 0, 64)
	fmt.Println(i)
	x, _ := strconv.ParseInt("0x1c8", 0, 64) // recognizing hex-formatted value
	fmt.Println(x)
	u, _ := strconv.ParseUint("45", 0, 64) // it is for unsigned integer. Since it does reserve a bit for sign (+, -) it holds twice the positive range of ParseInt
	fmt.Println(u)
	a, _ := strconv.Atoi("88") // parse with base 10
	fmt.Println(a)
}