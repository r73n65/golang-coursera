package main
import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
	"sort"
	"sync"
)

func sortSlice(wg *sync.WaitGroup, slice []int) {
	sort.Ints(slice)
	fmt.Println("Sorted Sub-Slice: ", slice)
	wg.Done()
}

func merge(from, to []int) []int {
	var slice []int
	var i, j int

	for i < len(from) {
		slice = append(slice, from[i])
		i++
	}
	for j < len(to) {
		slice = append(slice, to[j])
		j++
	}
	for i < len(from) && j < len(to) {
		if from[i] < to[j] {
			slice = append(slice, from[i])
			i++
		} else if from[i] > to[j] {
			slice = append(slice, to[j])
			j++
		} else {
			slice = append(slice, from[i], to[j])
			i,j = i+1,j+1
		}
	}
	return slice
}

func main() {
	parts := 4 
	slice := make([]int, 0)

	// User input - Enter numbers by list 
	fmt.Printf("Enter a list of numbers (with spaces separated): ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	num_list := strings.Fields(scanner.Text())

	for i := 0; i < len(num_list); i++ {
		num, _ := strconv.Atoi(num_list[i])
		slice = append(slice, num)
	}

	var result []int 
	var wg sync.WaitGroup
	length := len(slice) / parts

	// Sort slices
	wg.Add(parts)
	for j := 0; j < len(slice); j += length {
		go sortSlice(&wg, slice[j:j+length])
	}
	wg.Wait()

	// Merge result 
	for k := 0; k < len(slice); k += length {
		result = merge(result, slice[k:k+length])
	}

	sort.Ints(result)
	fmt.Printf("Final Sorted Numbers: %v\n", result)
}