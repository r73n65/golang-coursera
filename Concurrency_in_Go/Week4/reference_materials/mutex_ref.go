package main
import (
	"fmt"
	"sync"
)

var i int = 0
var mut sync.Mutex

func inc() {
	// puts the flag up - shared variable in use
	mut.Lock()

	i = i+ 1
	
	// puts the flag down - done using shared variable 
	mut.Unlock()
}