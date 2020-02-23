package main

import (
	"fmt"
	"github.com/ambsflip/goDemos/StructInterfacesSamples/structsDef"
)

// Generic function to calculate the total area of multiple shapes of different types
func CalculateTotalArea(shapes ...structsDef.Geometry) float64 {
	totalArea := 0.0
	for _, s := range shapes {
		totalArea += s.Area()
	}
	return totalArea
}

func main() {
	totalArea := CalculateTotalArea(structsDef.Circle{2}, structsDef.Rect{4, 5}, structsDef.Circle{10})
	fmt.Println("Total area = ", totalArea)
}