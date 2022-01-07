package main	
import "fmt"

func main() {
	var x float32
	fmt.Println("Enter a number: ")
	_, err := fmt.Scanf("%f", &x)

	if err != nil {
		fmt.Println("Error: ", err)
	} else {
		fmt.Println("Result: ", int(x))
	}
}
