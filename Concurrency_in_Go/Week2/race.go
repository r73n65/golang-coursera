package main
import (
	"fmt"
	"time"
)

var a int = 0
var b int = 0

func fooOne() {
	var temp int = a
	if temp%10 == 0 {
		b += 2
	}
	fmt.Println(b)
}

func fooTwo() {
	a += 1
	fmt.Println(a)
}

func main() {
	go fooOne()
	go fooTwo() 
	time.Sleep(10 * time.Second)
}

/*

Race Conditions Explanation:
* Outcome of the program depends on the interleaving (non-deterministic ordering)
* Occurs due to communication between tasks through the shared (same) variables

Output of the code showcases race condition:
> go run race.go
1
> go run race.go
2
1
> go run race.go
1
2
> go run race.go
1
2

*/