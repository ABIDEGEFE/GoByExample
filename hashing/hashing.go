package hashing

import (
	"fmt"
	"crypto/sha256"
	"unsafe"
)

func GoHashing(){
	fmt.Println("Go hashing is running.")
}

func HashingCaller(){
	s := "Hash mee"
	fmt.Println("size1: ", unsafe.Sizeof(sha256.Sum256([]byte(s))))
	h := sha256.New()
	fmt.Println("h", h)
	h.Write([]byte(s))
	fmt.Println(h)
	bs := h.Sum(nil)
	fmt.Println("Size of array:", unsafe.Sizeof(bs))
	fmt.Printf("Bs:", bs)
}