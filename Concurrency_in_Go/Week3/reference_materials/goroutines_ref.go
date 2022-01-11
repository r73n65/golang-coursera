package main
import (
	"fmt"
)

func foo() {
	return a = 3
}

// one goroutine created auto in main()
// when main goroutine is complete, other goroutines will force to exit as well
func main() {
	a = 1 
	foo()
	a = 2
}

// other goroutines are created using the go keyword
func gofunc() {
	a = 1
	go foo()
	a = 2
}
