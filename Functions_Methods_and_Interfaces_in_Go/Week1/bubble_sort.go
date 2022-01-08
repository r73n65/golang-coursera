package main
import (
	"fmt"
)

func Swap(sli []int, i int) {
	sli[i], sli[i+1] = sli[i+1], sli[i]
}

func BubbleSort(sli []int) {
	for i := 0; i < len(sli) - 1; i++ {
		for j:=0; j < len(sli) - i - 1; j++ {
			if sli[j] > sli[j+1] {
				Swap(sli, j)
			}
		}
	} 
}

func main() {
	slice := make([]int, 0, 10)
	var input int
	
	fmt.Println("You will be prompted 10 times to input an integer.")

	for i := 0; i < 10; i++ {
		fmt.Printf("Enter a number %f: ", i+1)
		fmt.Scanln(&input)
		slice = append(slice, input)
	}

	BubbleSort(slice)
	fmt.Println(slice)
}