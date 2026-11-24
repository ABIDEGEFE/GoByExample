package panics

import (
	"fmt"
	"os"
	"path/filepath"
)

func GoPanic(){
	fmt.Println("Go panic is running..")
}

func PanicCaller(){
	// 1st defer registered -> runs LAST (LIFO order)
	// When we put multiple defers in a function, we should be aware of their order of calling because it may affect the behaviour of our program.
	// This catches any panic that happens inside PanicCaller or functions it calls.
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()
	path := filepath.Join(os.TempDir(), "file")
	f := createFile(path)
	// 2nd defer registered -> runs FIRST when PanicCaller exits (even during a panic)
	defer closeFile(f)
	writeFile(f)
}

func createFile(path string)*os.File{
	fmt.Println("Creating a file")
	f, err := os.Create(path)
	if err != nil{
		panic(err)
	}
	return f
}
func writeFile(f *os.File){
	fmt.Println("Writing into a file")
	_, err := fmt.Fprintln(f, "Content wrote.")
	if err != nil{
		panic(err)
	}

	// Triggering intentional panic to test recover
	panic("Something went unexpectdely")
}

func closeFile(f *os.File){
	fmt.Println("Closing a file")
    err := f.Close()
	if err != nil{
		panic(err)
	}
}