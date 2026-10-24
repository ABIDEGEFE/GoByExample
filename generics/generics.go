package generics

import "fmt"

func SlicesIndex[S []E, E comparable](s S, v E)int{
	for i := range s{
		if v == s[i]{
			return i
		}
	}
	return -1
}

type List[T any] struct { // The type of element would any type. 
    head, tail *element[T] // head and tail are a pointer of the first and the last element.
}

type element[T any] struct {
    next *element[T] // each element has a pointer to the next element, [10|next]---->[13|next]
    val  T // each element has its own value.
}

func (lst *List[T]) Push(v T) {
	fmt.Println("list:", lst.tail, "\n", "head:", lst.head)
    if lst.tail == nil {
        lst.head = &element[T]{val: v}
        lst.tail = lst.head
    } else {
        lst.tail.next = &element[T]{val: v}
        lst.tail = lst.tail.next
    }
}

func (lst *List[T]) AllElements() []T {
    var elems []T
    for e := lst.head; e != nil; e = e.next {
        elems = append(elems, e.val)
    }
    return elems
}

func GenericCaller(){
	s := []string{"a", "B", "c"}
	fmt.Println(SlicesIndex[[]string](s, "c"))
    lst := List[int]{}
    lst.Push(10)
    lst.Push(13)
    lst.Push(23)
	lst.Push(11)
	lst.Push(12)
	lst.Push(88)
    fmt.Println("list:", lst.AllElements())
}