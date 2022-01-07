package main	
import "fmt"

func main() {
	x := [3]int {1, 2, 3}

	for i, v := range x {
		fmt.Printf("ind %d, val %d \n", i, v)
	}
}
