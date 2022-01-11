package main
import (
	"fmt"
	"time"
)

func main() {
	// Other routine
	go fmt.Printf("New routine")

	// Add a delay in main routine to give the new routine a chance to complete
	// Not a good choice to add a delay as time assumptions may be wrong 
	// Recommend: Synchronisation WaitGroup (refer to synchronisation.go)
	time.Sleep(100 * time.Millisecond)

	// Main routine
 	fmt.Printf("Main routine")
}