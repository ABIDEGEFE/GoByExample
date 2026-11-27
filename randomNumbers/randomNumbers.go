package randomNumbers

import (
	"fmt"
	"math/rand/v2"
)

func GoRandomNumbers(){
	fmt.Println("Go random number is running..")
}

func RandomNumbersCaller(){
	fmt.Println(rand.IntN(100)) // returns random number n where 0 <= n <= 100
	fmt.Println(rand.Float64()) // retunrs random number f where 0.0 <= f <= 1

	// customize the range of random float numbers
	fmt.Print(rand.Float64()*5, ",")
	fmt.Print(rand.Float64()*5 + 5)
    fmt.Println()

	/*
	Using a fixed, known seed makes a pseudo-random number generator deterministic,
    meaning it will produce the exact same sequence of numbers every time the code runs.
    Computers cannot naturally generate true randomness through pure math. 
	Instead, algorithms like rand.NewPCG take a starting number (the seed) and run a formula that produces a sequence of numbers that statistically looks random (evenly distributed, no obvious pattern). 
	Because the formula is pure math, giving it the same starting numbers (42 and 1024) guarantees the exact same output every time.
    While this seems to defeat the purpose of randomness, deterministic "randomness" is essential in software engineering for several practical reasons.
	*/

	s2 := rand.NewPCG(42, 1024)
	fmt.Println("sssss2", s2)
	r2 := rand.New(s2)
	fmt.Println("rrrrrrr2", r2)
	fmt.Print(r2.IntN(100), ",")
	fmt.Print(r2.IntN(100))
	fmt.Println()

	s3 := rand.NewPCG(42, 1024)
	r3 := rand.New(s3)
	fmt.Print(r3.IntN(100), ",")
	fmt.Println(r3.IntN(100))

	/*
	If we are generating passwords, encryption keys, or security tokens, we shoudl always use crypto/rand instead of math/rand/v2.
	As math/rand/v2 is still mathematically predictable even when auto-seeded
	*/
}