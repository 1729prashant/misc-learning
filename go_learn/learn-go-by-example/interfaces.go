package main

import (
	"fmt"
	"math"
)

type geometry interface {
	area() float64
	perim() float64
}

type rect2 struct {
	width, height float64
}

type circle struct {
	radius float64
}

func (r rect2) area() float64 {
	return r.width * r.height
}

func (r rect2) perim() float64 {
	return 2*r.width + 2*r.height
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

// If a variable has an interface type, then we can call methods that
// are in the named interface. Here’s a generic measure function taking
// advantage of this to work on any geometry.
func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

// Sometimes it’s useful to know the runtime type of an interface value.
// One option is using a type assertion as shown here; another is a type switch.
func detectCircle(g geometry) {
	if c, ok := g.(circle); ok {
		fmt.Println("circle with radius", c.radius)
	}
}

func main() {
	r := rect2{width: 3, height: 4}
	c := circle{radius: 5}

	measure(r)
	measure(c)
	detectCircle(r)
	detectCircle(c)
}
