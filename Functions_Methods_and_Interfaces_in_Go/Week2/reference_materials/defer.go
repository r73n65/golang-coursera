package main
import (
	"fmt"
)

// Defer will be executed later after all the functions (main) have been called & completed
// Comment either of the examples to run 

// Example 1
func main() {
	defer fmt.Println("Bye!")
	fmt.Println("Hello!")
}

// Example 2
func main() {
	i := 1
	defer fmt.Println(i+1)
	i++
	fmt.Println("Hello!")
}