package main
import ( 
	"fmt"
	"bufio"
	"os"
	"strings"
)

func main() {
	fmt.Println("Enter a string: ")

	// read input with spaces
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan() 
	input := scanner.Text()
	
	// convert the string into lowercase
	input = strings.ToLower(input)
	
	if strings.HasPrefix(input, "i") && strings.HasSuffix(input, "n") && strings.Contains(input, "a") {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}
}