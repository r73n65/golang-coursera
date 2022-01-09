package main
import (
	"data"
)

func main() {
	var p data.Point
	p.InitMe(3, 4)
	p.Scale(2)
	p.PrintMe()
}