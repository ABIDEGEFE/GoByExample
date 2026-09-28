package values


import "fmt"


func GoValues(){
  name := "Abebe"
  var fatherName = "Kebede"

  fmt.Println("Full name :" + name + fatherName)

  fmt.Println("INTEGERS.......FLOATS")
  age := 34
  fmt.Println("He is", age, "years old.")

  fmt.Println("BOOLEANS.........TRUE/FALSE")
  fmt.Println(fatherName == "Kebede" && name == "Abebe", "expected outcome is True")
  fmt.Println(true || false)
  fmt.Println("Expected outcome is True")
}
