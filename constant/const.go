package constant

import "fmt"

const PI = 3.14  // Numeric constants does not have type.
const name string = "constnatName"
func GoConstant(){
   const v = 500000000
   const d = 3e20/v
   fmt.Println("Constant value of PI is", PI)
   fmt.Println("Original value of d is", d)
   fmt.Println("Changed type of v is", int64(d))
   fmt.Println("Constnat name is", name)
}