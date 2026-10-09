package structs


// import "fmt"

type person struct{  // structs has object-attributes structure, used to organize data with different data type.
	Name string  // field name must be capital to be exported. Capitalization is visibility.
	Age int
}

func Teacher() *person{
	var p person
	p.Name = "GirmaTeacher"
	p.Age = 99
	return &p
}

func Driver() *person{
	var pd person
	pd.Name = "JoneDriver"
	pd.Age = 33
	return &pd
}

