package main
import "fmt"

// Pass by Reference - Pass a pointer as an argument
// Advantage: Copying Time
// Disadvantage: Data Encapsulation
func foo(y *int) {
	*y = *y + 1
}

func main() {
	x := 2
	foo(&x)
	fmt.Print(x)
}