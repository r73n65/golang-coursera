package main
import (
	"fmt"
	"strings"
	"bufio"
	"os"
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
		var invalid_flag bool

		// Ask user for input 
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Print("> ")
		scanner.Scan()
		fields := strings.Fields(strings.ToLower(scanner.Text()))
		
		if len(fields) != 3 {
			fmt.Println("Error. Additional input has been entered.")
			continue
		}

		var a Animal
		command := fields[0]
		name := fields[1]
		action := fields[2]

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
