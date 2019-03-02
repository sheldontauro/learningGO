package main 

import(
	"fmt"
	"math/rand"
	"math"
)

func foo() {
	fmt.Println("Welcome to go" , math.Sqrt(5))	
}

func main() {
	fmt.Println(rand.Intn(100))
	foo()	
}