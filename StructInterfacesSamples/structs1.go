package main

import (
	"fmt"
	"github.com/ambsflip/goDemos/StructInterfacesSamples/structDef1"
)

//ref https://gobyexample.com/structs
//define a struct. initialise it with custom/default values
//use of 'constructors'
//use of methods in go structs




func main() {

	fmt.Println(structDef1.Person{"Bob", 17, false})
	fmt.Println(structDef1.Person{Name: "Fred"}) // explicit member has to be defined if all values are not provided in order
	fmt.Println(&structDef1.Person{Name: "Ann", Age: 40})
	fmt.Println(structDef1.New("Jon"))

	s := structDef1.Person{Name: "Sean", Age: 50, Voter:false}
	fmt.Println(s.Name)
	sp := &s
	fmt.Println(sp.Age)
	sp.Age = 51
	fmt.Println(s.Age)

	fmt.Println("isVoter ? : ", sp.Voter)
	fmt.Println(sp.IsMajor())
	sp.CallByValue()
	fmt.Println("isVoter ? : ", sp.Voter)
	fmt.Println("Name: ",sp.Name)
}