package main
import (
	"fmt"
)

type MyInt int // receiver tyoe 

func (mi MyInt) Double() int {
	return int(mi*2)
}

func main(){
	v := MyInt(3)
	fmt.Println(v.Double()) // implicit method argument
}