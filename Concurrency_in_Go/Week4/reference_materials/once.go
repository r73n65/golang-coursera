package main
import (
	"fmt"
	"sync"
)

var on sync.Once

func setup() {
	fmt.Println("Init")
}

func main() {
	// setup() should execute only once
	// "hello" should not print until setup() returns 
	on.Do(setup) 
	fmt.Println("hello")
}