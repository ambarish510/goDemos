package main

import (
	"fmt"
	"github.com/ambsflip/goDemos/EnumEquivalent/EnumStruct"
)

func main() {
	foo(EnumStruct.NO)
	foo(EnumStruct.YES)
	foo(EnumStruct.DONTKNOW)
	foo("abcd")
}

func foo(arg EnumStruct.MyString){
	fmt.Print(arg)
}