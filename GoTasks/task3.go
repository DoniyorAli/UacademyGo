package main

import "fmt"

//*Area of circle inscribed in a square
func  area(r float32) (area float32) {
	radius := r / 2
	areaSquare := r * r
	areaCircle := 3.14 * (radius * radius)
	findShadeArea := areaSquare - areaCircle
	return findShadeArea
	// box := (4 * (r * r) - 3.14 * (r * r)) / 4
	// return box
}

func main() {

	var r float32
	fmt.Scanf("%f", &r)
	fmt.Println("R =", r)
	fmt.Printf("Area: %0.2f\n", area(r))

}