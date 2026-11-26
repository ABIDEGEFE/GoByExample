package gotime

import(
	"fmt"
	"time"
)

func GoTime(){
	fmt.Println("Go time is running.")
}

func TimeCaller(){
	var p = fmt.Println

	p(time.Now())
	then := time.Date(2019, 4, 5, 20, 11, 20, 2000, time.UTC)
	p(then)

	p(then.Year())
	p(then.Month())
	p(then.Day())
	p(then.Hour())
	p(then.Minute())
	p(then.Second())
	p(then.Location())
	p(then.Weekday())

	p(then.Before(time.Now()))
	p(then.After(time.Now()))
	p(then.Equal(time.Now()))

	diff := time.Now().Sub(then)

	p(then.Add(diff))
	p(then.Add(-diff))
}