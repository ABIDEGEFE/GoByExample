package gosIterator


import (
    "fmt"
    "iter"
    "slices"
)

type List struct{
    head []int
}

func (lst *List) All() iter.Seq[int]{ // This function returns a closure function which capture yeild as its environment.
    return func(yeild func(a int)bool){
        for i := 0; i < len(lst.head); i++{
            if !yeild(lst.head[i]){
                fmt.Println("I value when it returns", i)
                return
            }
        }
    }
}

func IterCaller(){
    head := []int{10, 20, 30}
    lst := &List{head}
    count := 0
    for e := range lst.All(){
        count++
        fmt.Println("Iterator called for:", count, "time", "with val", e)
    }
    fmt.Println(slices.Collect(lst.All()))
}