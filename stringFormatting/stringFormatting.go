package stringFormatting

import (
	"fmt"
	"os"
)

func GoStringFormatting(){
	fmt.Println("Go string formatting running..")
}

type point struct{
	x, y int
}

var p = fmt.Printf
func StringFormattingCaller(){
	po := point{1, 2}
	p("Struct instance: %v\n", po)
	p("Struct instance with field name: %+v\n", po)
	p("Struct code snippet: %#v\n", po)
	p("Struct type of value: %T\n", po)
	p("Formatting boolean: %t\n", true)
	p("Formatting integers: %d\n", 123)
	p("Formatting integers to binary: %b\n", 5)
	p("ASCII value: %c\n", 97)
	p("Hexadecimal value: %x\n", 11)
	p("Formatting float: %f\n", 123.99)
	p("Scientific notation: %e\n", 123000.0)
	p("String formatting: %s\n", "\"string\"")
	p("Double quote string formatting: %q\n", "\"string\"")
	p("Hexadecimal value of string: %x\n", "This is new string")
	p("Representation of a pointer: %p\n", &po)
	p("Integer formatting with specified width: |%6d|%6d|\n", 55, 44)
    p("width4: |%6s|%6s|\n", "foo", "b")
	p("width4: |%-6s|%-6s|\n", "foo", "b")
	bin := fmt.Sprintf("Returned binary: %b\n", 5)
	fmt.Println(bin)
	fmt.Fprintf(os.Stderr, "Write to standard error: %s\n", "error")
}