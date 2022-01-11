package main
import (
	"fmt"
)

// Select statement: to wait on the first data from a set of channels

// Example 1 
func main() {
	c1 := make(chan int)
	c2 := make(chan int)
	select {
		case a = <- c1:
			fmt.Println(a)
		case b = <- c2:
			fmt.Println(b)
	}
}

// Example 2 
// Depends on which case got unlock first
func main() {
	select {
		case a = <- inchan:
			fmt.Println("Received a")
		case outchan <- b:
			fmt.Println("sent b ")
	}
}

// Example 3 
// Select with an abort channel
// Infinite for loop until abort happens
func main() {
	for {
		select {
			case a <- c:
				fmt.Println(a)
			case <- abort:
				return
		}
	}
}

// Example 4
// Default Select
func main() {
	case a = <- c1:
		fmt.Println(a)
	case b = <- c3:
		fmt.Println(b)
	default:
		fmt.Println("nop")
}