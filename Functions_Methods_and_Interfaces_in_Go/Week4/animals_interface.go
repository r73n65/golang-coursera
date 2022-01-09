package main
import (
	"fmt"
	"strings"
)

type Animal interface {
	Eat()
	Move()
	Speak()
}

type Cow struct { name string }
func (c Cow) Eat() {
	fmt.Println("grass")
}

func (c Cow) Move() {
	fmt.Println("walk")
}

func (c Cow) Speak() {
	fmt.Println("moo")	
}

type Bird struct { name string }
func (b Bird) Eat() {
	fmt.Println("worms")
}

func (b Bird) Move() {
	fmt.Println("fly")
}

func (b Bird) Speak() {
	fmt.Println("peep")	
}

type Snake struct { name string }
func (s Snake) Eat() {
	fmt.Println("mice")
}

func (s Snake) Move() {
	fmt.Println("slither")
}

func (s Snake) Speak() {
	fmt.Println("hsss")	
}

func main() {
	// Map all animals
	animal_map := make(map[string]Animal)
	
	// Infinite Loop
	for true {
		// Ask user for input 
		var command, name, action string
		var invalid_flag bool
		fmt.Print("> ")
		fmt.Scan(&command, &name, &action)

		// Convert input to lower case
		command = strings.ToLower(command)
		name = strings.ToLower(name)
		action = strings.ToLower(action)

		var a Animal

		if command == "newanimal" {
			if action == "cow" {
				a = Cow{name: name}
				invalid_flag = false
			} else if action == "bird" {
				a = Bird{name: name}
				invalid_flag = false
			} else if action == "snake" {
				a = Snake{name: name}
				invalid_flag = false
			} else {
				invalid_flag = true
			}

			if !invalid_flag {
				animal_map[name] = a
				fmt.Println("Created it!")
			}

		} else if command == "query" {
			animal_name, ok := animal_map[name]

			if !ok {
				fmt.Println("Animal is not present.")
				break
			}
			
			if action == "eat" {
				animal_name.Eat()
			} else if action == "move" {
				animal_name.Move()
			} else if action == "speak" {
				animal_name.Speak()
			} else {
				invalid_flag = true
			}
		}

		if invalid_flag {
			fmt.Println("Invalid input.")
		}
	}
}
