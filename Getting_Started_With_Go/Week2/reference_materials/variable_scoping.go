package main
import "fmt"

// Heap
var x = 4

func f() {
	// Stack 
	// var x = 4
	fmt.Printf("%d", x)
}

func g() {
	fmt.Printf("%d", x)
}

func main() {
	f()
	g()
}