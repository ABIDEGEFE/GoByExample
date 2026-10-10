package method


import "fmt"

type rectangle struct{
	height int 
	width int
}

func (r *rectangle)area() int {  // we are giving full access of the struct for this function. It is able to mutate data in the struct
	r.height = 109  // this will change the actual data created in Caller function.
	r.width = 101
   fmt.Printf("r in Area %p\n", r)
   return r.height * r.width
}

func (r rectangle)permiter() int{ // value reciever only recieves a copy of the data with isolated memory. Any change will remain in its memory
	r.height = 10 // This changes its own data. The Caller function would not be affected. 
	r.width = 11
	fmt.Printf("R in Permiter %p\n", r)
    return 2*r.height + 2*r.width
}

func Caller(){
	r := rectangle{8, 9}  // Initial state of r
	fmt.Printf("r in Caller %p\n", &r)
	fmt.Println("Area:", r.area())
	fmt.Println("Permiter", r.permiter())
    fmt.Println("r instance", r) // r changed by area function
}
