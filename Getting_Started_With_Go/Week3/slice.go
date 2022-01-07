package main	
import (
	"fmt"
	"strconv"
	"sort"
)

func main() {
	slice := make([]int, 0, 3)
	var input string
	fmt.Print("Enter a number or X to exit: ")
	fmt.Scanln(&input)

	for input != "X" {
		intInput, _ := strconv.Atoi(input)
		slice = append(slice, intInput)
		sort.Ints(slice)
		fmt.Println(slice)
		fmt.Scanln(&input)
	}
}
