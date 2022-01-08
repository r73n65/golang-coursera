package main	
import (
	"fmt"
	"encoding/json"
)

func main() {
	var inputName string
	var inputAddress string 

	details := make(map[string]string)
	fmt.Println("Enter your name: ")
	fmt.Scanln(&inputName)
	
	fmt.Println("Enter your address: ")
	fmt.Scanln(&inputAddress)

	details["name"] = inputName
	details["address"] = inputAddress

	json, err := json.Marshal(details)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(json))
}

