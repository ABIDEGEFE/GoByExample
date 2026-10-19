package embededStruct

import "fmt"


func GoEmbededStruct(){
	fmt.Println("Embeded struct is running.")
}

type base struct{
	num int
}

type container struct{
	base  // embeding base struc into container struct so that the instance of container can access base's fields. 
	str string
}

func (b base) describe(){
	fmt.Println("Number of the base", b.num)
}

func EmbededCaller(){
	co := container{
		base: base{
			num: 10,
		},
		str: "some number",
	}
    
	fmt.Println("Accessing base feild through container instance", co.num)
	fmt.Println("Alternative access", co.base.num)
	co.describe()

    manager := Manager{  // instance of the outer struct is able to access fields and methods of the embeded struct.
		Employee: Employee{
			Name: "Jhone",
			Email: "jhone@gmail.com",
		},
		Team: "Finance",
	}

	type organization interface{
		EmployeeContact() string
	}
	var org organization = manager
	contactInfo := org.EmployeeContact()
	fmt.Println(contactInfo)

}

/*
      =================== Scenario which requires strut embeding =============
	  Imagine a company application that manages different employee types. Every employee needs common identity data and behavior—such as a name, email, and a method that returns their contact details.
	  But a manager also needs manager-specific data, such as the team they lead. Managers are Employees but the need additional features.
*/

type Employee struct{
	Name string
	Email string
}

type Manager struct{
	Employee // no need to duplicate fields of Employee here
	Team string // Additional feature of Manager
}

func (e Employee) EmployeeContact()string{
	return e.Name + "<" + e.Email + ">"
}