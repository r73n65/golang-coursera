package main
import (
	"fmt"
)

// Purpose of channels: Transfer data between goroutines

func main() {
	// create a channel
	c := make(chan int)

	// send data on a channel
	c <- 3

	//receive data from a channel
	x := <- c
}
