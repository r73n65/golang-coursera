package main
import "fmt"

// Pass by Value
// Advantage: Data Encapsulation
// Disadvantage: Copy time
func foo(y int) {
	y = y + 1
}

func main() {
	x := 2
	foo(x)
	fmt.Print(x)
}