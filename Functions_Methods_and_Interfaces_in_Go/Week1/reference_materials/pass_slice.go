package main
import "fmt"

// Pass slice instead
// Slices contain a pointer to the array
func foo(sli []int) {
	sli[0] = sli[0] + 1
}

func main() {
	a := []int{1, 2, 3} // size is not declared unlike array
	foo(a)
	fmt.Println(a)
}
