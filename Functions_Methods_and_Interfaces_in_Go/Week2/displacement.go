package main
import (
	"fmt"
)

// Given formula: s = 0.5 * a * t^2 + v0 * t + s0

func GenDisplaceFn(a, v, s float64) func(float64) float64 {
	return func(t float64) float64 {
		return 0.5 * a * t * t + v * t + s
	}
}

func main(){
	var a, v, s, t float64

	fmt.Print("Enter acceleration: ")
	fmt.Scanln(&a)

	fmt.Print("Enter initial velocity: ")
	fmt.Scanln(&v)

	fmt.Print("Enter initial displacement: ")
	fmt.Scanln(&s)

	fmt.Print("Enter time taken for displacement in seconds: ")
	fmt.Scanln(&t)

	fn := GenDisplaceFn(a, v, s)

	fmt.Println("Displacement: ", fn(t))
}
	