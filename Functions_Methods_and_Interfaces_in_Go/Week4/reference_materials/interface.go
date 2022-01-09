package main
import (
	"fmt"
)

// Using interface can hide differences between types

type Speaker interface {
	Speak()
}

type Dog struct {
	name string
}

func (d Dog) Speak() {
	fmt.Println(d.name)
}

func main() {
	var s1 Speaker // Nil Dynamic Type 
	d1 := Dog{"Brian"} // Dynamic Value
	s1 = d1
	s1.Speak()
}