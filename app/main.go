package main

import (
  // "fmt"
  // "codingPractice/values"
  // "codingPractice/variables"
  // "codingPractice/constant"
  // "codingPractice/conditions"
  // "codingPractice/array"
  // "codingPractice/slices"
  // "codingPractice/maps"
  // "codingPractice/functionTypes/function"
  // "codingPractice/ranges"
  // "codingPractice/pointers"
  // "codingPractice/stringsAndRunes"
  // "codingPractice/structs"
  // "codingPractice/method"
  // "codingPractice/interfaces"
  "codingPractice/enums"
)

func main(){
  // fmt.Println("hellow main function")
  // fmt.Println("###############VALUES IN GO##############")
  // values.GoValues()
  // fmt.Println("################## VARIABLES IN GO #################")
  // variables.GoVariables()

  // fmt.Println("###############CONSTANTS IN GO#################")
  // constant.GoConstant()

  // conditions.GoIfStatement()
  // conditions.GoSwitchStatement()

  // array.GoArray()
  // slices.GoSlices()
  // maps.GoMaps()
  
  // function.GoFunction()
  // age := function.SetAge("22/1/2003")
  // fmt.Println("You are", age, " years old")
  // total := function.Sum(1, 2, 3, 4)
  // fmt.Println("Total value", total)
  // val, err := function.ReturnMulitpleValues("akalu") //if we only want subset of the return value, we can left either as blank(_, err := func())
  // if err != nil {
  //   fmt.Println("something went wrong", err)
  // }else{
  //   fmt.Println("Correct", val)
  // }

  // sum := function.VariadicFunc(1, 2, 3, 4)
  // fmt.Println("Sum of slice", sum)
  // function.PrintAny(1, 2, true, "hey there")

  // nextInt := function.ClosureFunc()
  // fmt.Println(nextInt())
  // fmt.Println(nextInt())
  // newInt := function.ClosureFunc()
  // fmt.Println(newInt())
  // fmt.Println(nextInt())

  // fib := function.RecursiveFunc(4)
  // fmt.Println("Recursive Function", fib)

  // ranges.GoRange()

  // pointers.GoPointers()

  // stringsAndRunes.GoStrings()

  // teacher := structs.Teacher()
  // driver := structs.Driver()
  // fmt.Println("Memory address of teacher", &teacher, "Memory address of Driver", &driver) // each function constructs its own dedicated instance of struct
  // fmt.Println(teacher.Name) // use . to access fields of struct
  // fmt.Println(driver.Name)

  // method.Caller()

  // interfaces.Caller()

  enums.EnumsCaller()

}
