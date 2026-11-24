package stringFunction

import (
	"fmt"
	s "strings"
)

func GoStringFunction(){
	fmt.Println("Go string function is running..")
}
var p = fmt.Println
func StringFunctionCaller(){
   p("Contains", s.Contains("test", "es"))  // returns true if the second argument is found in the first argument
   p("Count", s.Count("abcdabcd", "ab")) // returns the frequency of the second argument in the first argument
   p("Has Prefix", s.HasPrefix("function", "un")) // returns true if the second argument is the prefix of the first one
   p("Has Suffix", s.HasSuffix("Function", "ion")) // returns true if the second argument is the suffix of the first one
   p("Index", s.Index("function", "tio")) // returns zero-indexed value of the second argument in the fist argument.
   p("Join", s.Join([]string{"func", "tion"}, "-")) // returns a joined string from list of strings
   p("Repeat", s.Repeat("abc", 5)) // returns a string of repeated characters or strings
   p("Replace", s.Replace("fooabdcooo", "o", "u", -1)) // returns a string of which its all specified values are replaced by the third argument
   p("Replace", s.Replace("fooabdcooo", "o", "u", 1)) // returns a string of which its first specified value is replaced by the third argument
   p("Split", s.Split("abc-def-gho#abc", "#")) // returns a slice whose element is a string splited from the given string by the second argument
   p("To lower", s.ToLower("ABefgCD")) // retuns a string by changing its all upper case characters in to lower case
   p("To upper", s.ToUpper("ccddAABB")) // returns a string by changing its all lower case letters into upper case
}  