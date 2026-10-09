package stringsAndRunes


import (
	"fmt"
	// "unicode/utf8"
)

func GoStrings(){
	s := "teshomeuuu99!!!"  // strings are read-only slice of bytes, and they store unicode code of each character(rune)
	                        // designed to support every language, emojis and special characters in the world.
    for i := 0; i<len(s); i++{
		fmt.Println("Char:", s[i]) // UTF-8 byte of each rune.
	}

	for idx, c := range s{  // range over the unicode code of characters and their corresponding index
		fmt.Printf("%#U starts at %d\n", c, idx)
		fmt.Println("Actual value", string(c)) // built-in string function will decode the utf-8 code to its actual value
		ExamineRune(c) // pass rune value
	}
	
}

func ExamineRune(r rune){ // rune data type used to declare a character
	if r == 't'{
		fmt.Println("It is found")
	}
}