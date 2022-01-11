package main
import (
	"fmt"
	"sync"
)

// Global variables
var i int = 0
var wg sync.WaitGroup

// Increment function
func inc() {
	i = i+1
	wg.Done()
}

func main() {
	wg.Add(2)
	go inc()
	go inc()
	wg.Wait()
	fmt.Println(i)
}
