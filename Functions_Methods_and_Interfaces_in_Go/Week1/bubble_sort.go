package main
import (
	"fmt"
)

func Swap(sli []int, i int) {
	sli[i], sli[i+1] = sli[i+1], sli[i]
}

func BubbleSort(sli []int) {
	for i := 0; i < len(sli)-1; i++ {
		for j:=0; j < len(sli)-i-1; j++ {
			if sli[j] > sli[j+1] {
				Swap(sli, j)
			}
		}
	} 
}

func main() {
	slice := make([]int, 0, 10)
	var input int
	
	fmt.Println("You will be prompted 10 times to input a positive integer.")

	for i := 0; i < 10; i++ {
		fmt.Printf("Enter a number #%d: ", i+1)
		_, err := fmt.Scanln(&input)

		if err != nil {
			fmt.Println(err)
			break
		}

		if input < 0 {
			fmt.Println("Negative integer is entered. Please input a positive integer again: ")
			fmt.Scanln(&input)
		}

		slice = append(slice, input)
	}

	BubbleSort(slice)
	fmt.Println(slice)
}