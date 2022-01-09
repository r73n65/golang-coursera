package data
import (
	"fmt"
)

// Public: An identifier is exported if it begins with a capital letter.
// Private: If an identifier starts with a lower case letter, it can only be accessed from within the package.

type Point struct {
	x float64
	y float64
}

func (p *Point) InitMe(xn, yn float64) float64 {
	p.x = xn
	p.y = yn
}

func (p *Point) Scale(v float64) float64 {
	p.x = p.x * v
	p.y = p.y * v
}

func (p *Point) PrintMe() {
	fmt.Println(p.x, p.y)
}