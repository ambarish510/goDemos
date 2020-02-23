package main

import (
	"fmt"
	"github.com/ambsflip/goDemos/StructInterfacesSamples/structsDef"
)



func measure(g structsDef.Geometry) {
	fmt.Println(g)
	fmt.Println(g.Area())
	fmt.Println(g.Perim())
}

func main() {
	r := structsDef.Rect{Width: 3, Height: 4}
	c := structsDef.Circle{Radius: 5}

	measure(r)
	measure(c)

	//An interface type can hold any value that implements all its methods
	var s structsDef.Geometry = structsDef.Circle{5.0}
	fmt.Printf("Shape Type = %T, Shape Value = %v\n", s, s)
	fmt.Printf("Area = %f, Perimeter = %f\n\n", s.Area(), s.Perim())

	var s1 structsDef.Geometry = structsDef.Rect{4.0, 6.0}
	fmt.Printf("Shape Type = %T, Shape Value = %v\n", s1, s1)
	fmt.Printf("Area = %f, Perimeter = %f\n", s1.Area(), s1.Perim())
}