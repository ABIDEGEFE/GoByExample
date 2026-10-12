package interfaces

import (
	"fmt"
	"math"
)

// Interfaces are collection of common(universal) methods executed by instances of multiple structs
// we may have struct1. struct2, struct3. All may want to execute a method called area()
// we may be forced to decide which struct(data) should get excuted at specific time based on conditions
// when condiiton1 met struct1 would be utilized, condition2 met struct2 would be utilized. This might be decided at run time.
// interfaces helps us to check the struct type and execute a logic accordingly.
type geometry interface {
    area() float64
    perim() float64
}

// structs are supposed to have all methods declared in the interface to implement the interface.
type rect struct {
	height, width float64
}

type circle struct {
	radius float64
}

func (r rect)area()float64{
    return r.height * r.width   
}
func (r rect)perim()float64{
	return 2*r.height + 2*r.width
}

func (c circle)area()float64{
	return math.Pi * c.radius * c.radius
}
func (c circle)perim()float64{
	return 2*math.Pi*c.radius
}
// structs can have additional methods which are not declared in the interface.
// But they are supposed to have all methods declared in the interfaces.
func (c circle)volume(){
	fmt.Println("This is volume of circle")
}

func measure(g geometry){
	if c, ok := g.(circle); ok{
		fmt.Println("Area of circle", c.area())
	}
	if r, ok := g.(rect); ok{
		fmt.Println("Area of rectangle", r.perim())
	}
}

func Caller(){
	var c = circle{8}
	var r = rect{10, 10}
    c.volume()
	// rectArea := r.area()
	// fmt.Println("Area of rectangle", rectArea)
	// perimArea := r.perim()
	// fmt.Println("Permiter of rectangle", perimArea)

	// circleArea := c.area()
	// fmt.Println("Area of circle", circleArea)
	// circlePerim := c.perim()
	// fmt.Println("Perimter of Circle", circlePerim)
	
	measure(r)
	measure(c)
}