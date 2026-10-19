package enums


import "fmt"

type ServerState int

const (
	stateIdle ServerState = iota // initialize all const values starting from 0 to 3. All would have a type of of ServerState
	stateConnected
	stateRetrying
	stateError
)

func EnumsCaller(){
	// status := transition(stateIdle)
	// fmt.Println(status)

    // nextStatus := transition(status)
	// fmt.Println(nextStatus)
	productStatus := ProductStatus(Pending)
	fmt.Println("Product status:", productStatus)
}

var stateName = map[ServerState]string{
	stateIdle: "idel",
	stateConnected: "connected",
	stateRetrying: "Retrying",
	stateError: "Error",
}

func (ss ServerState) String() string{
	return stateName[ss]
}

func transition(s ServerState) ServerState{
	switch s {
	case stateIdle:
		return stateConnected
	case stateConnected:
		return stateIdle
	case stateError:
		return stateRetrying
	default:
		panic(fmt.Errorf("Unkown status : %s", s))
	}
}

/*
      =======SCENARIO WHICH COULD BE IMPLEMENTED USING enums=======
Imagine we are building a small online-store system. Every order must be in exactly one status:

    - Pending

    - Paid

    - Shipped

    - Delivered

    - Cancelled
*/

type StatusOrder int

const(
	Pending StatusOrder = iota
	Paid
	Shipped
	Delivered
	Cancelled
)

var orderStatus = map[StatusOrder]string{
	Pending: "Pending",
	Paid: "Paid",
	Shipped: "Shipped",
	Delivered: "Delivered",
	Cancelled: "Cancelled",
}

func (so StatusOrder) String() string{
	return orderStatus[so]
}

func ProductStatus(status StatusOrder) StatusOrder{
   switch status{
   case Pending:
	return Paid
   case Paid:
	return Shipped
   case Shipped:
	return Delivered
   default:
	panic(fmt.Errorf("Product status undefined: %s", status))
   }
}