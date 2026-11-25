package regularexp

import (
	"fmt"
	"bytes"
	"regexp"
)

func GoRegularExp(){
	fmt.Println("Go regular expression is running..")
}

func RegularExpCaller(){
	match, _ := regexp.MatchString("p([a-z]+)ch", "snpeachnn") // return true if it finds a substring match which satisfies the regular expression
	fmt.Println(match)
    // we should compile the regular expression to work with other tasks
	r := regexp.MustCompile("p([a-z]+)ch")
	fmt.Println(r) // Prints out the compiled regular expression
	fmt.Println(r.FindString("abcd unpeachch punch")) // returns the first much of substring
	fmt.Println(r.FindStringIndex("abcd unpeachch punch")) // returns array of start and end index of the first match
	fmt.Println(r.FindStringSubmatch("abcd unpeachch punch")) // returns array of string which satisfies both the whole patter and sub-match
	fmt.Println(r.FindStringSubmatchIndex("abcd unpeachch punch")) // returns array of start and end index of both the whole pattern and sub-match
	fmt.Println(r.FindAllString("abcd unpeachch punch", -1)) // returns array of all strings which satisfy the pattern
	fmt.Println(r.FindAllString("abcd upeachch punch and preach", 2)) // returns the first 2 match

	// we can also provide bytes as a data instead of raw string
    fmt.Println(r.Match([]byte("peach"))) // returns true if it finds the match
	fmt.Println(r.ReplaceAllString("abcd upeachch punch", "mango")) // returns new string which its substrings are replaced by the specified string
	in := []byte("for peach")
	out := (r.ReplaceAllFunc(in, bytes.ToUpper))
	fmt.Println(string(out))
}