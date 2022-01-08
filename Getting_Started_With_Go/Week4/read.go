package main	
import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Person struct {
	fname string
	lname string
}

func main() {
	slice := make([]Person, 0, 1)
	
	// read file name 
	fmt.Println("Enter file name: ")
	fileNameScanner := bufio.NewScanner(os.Stdin)
	fileNameScanner.Scan()
	fileName := fileNameScanner.Text()

	// open file
	content, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
	}

	// read file content
	contentScanner := bufio.NewScanner(content)

	for contentScanner.Scan() {
		arr := strings.Split(contentScanner.Text(), " ")
		slice = append(slice, Person{fname: arr[0], lname: arr[1]})
	}

	content.Close()

	for _, person := range slice {
		fmt.Printf("First Name: %s, Last Name: %s\n", person.fname, person.lname)
	}
}