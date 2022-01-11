package main 
import (
	"sync"
)

// Deadlock: circular dependencies cause all involved goroutines to block
// Deadlock example

var wg sync.WaitGroup

func dostuff(c1 chan int, c2 chan int) {
	<- c1
	c2 <- 1
	wg.Done()
}

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	wg.Add(2)
	go dostuff(ch1, ch2) // goroutine block
	go dostuff(ch2, ch1) // goroutine block
	wg.Wait()
}
