package main
import (
	"fmt"
	"strings"
)

type Animal struct {
	food string
	locamotion string
	sound string
}

func (a Animal) Eat() {
	fmt.Println(a.food)
}

func (a Animal) Move() {
	fmt.Println(a.locamotion)
}

func (a Animal) Speak() {
	fmt.Println(a.sound)	
}

func main() {
	// Mapping of all animals
	animalMap := map[string]Animal{
		"cow": Animal{"grass", "walk", "moo"},
		"bird": Animal{"worms", "fly", "peep"},
		"snake": Animal{"mice", "slither", "hsss"},
	}
	
	// Infinite Loop
	for true {
		// Ask user for input 
		var name, action string
		fmt.Print("> ")
		fmt.Scan(&name, &action)

		// Convert input to lower case
		name = strings.ToLower(name)
		action = strings.ToLower(action)

		a, ok := animalMap[name]

		ok = ok && (action == "eat" || action == "move" || action == "speak")

		// Print invalid input for either animal or action
		if !ok {
			fmt.Println("Invalid input.")
			continue
		}

		// Print respective output for valid animal and action
		if action == "eat" {
			a.Eat()
		} else if action == "move" {
			a.Move()
		} else if action == "speak" {
			a.Speak()
		}
	}
}
