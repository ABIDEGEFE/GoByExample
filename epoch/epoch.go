package epoch

import (
	"fmt"
	"time"
)

func GoEpoch(){
	fmt.Println("Go epoch is running.")
}

func EpochCaller(){
	now := time.Now()
	fmt.Println(now.Unix())    // epoch is fixed amount of time measured from January 1, 1970 which is when unix system started counting time.
	fmt.Println(now.UnixMilli())
	fmt.Println(now.UnixNano())

	fmt.Println(time.Unix(now.Unix(), 0))
}