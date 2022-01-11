package main
import (
	"fmt"
	"sync"
)

// sync.WaitGroup solves the issue of time assumptions 
func foo(wg *sync.WaitGroup) {
	fmt.Printf("New routine")
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go foo(&wg)
	wg.Wait()
	fmt.Printf("Main routine")
}
