package main

import (
	"fmt"
	"strings"
)

func main()  {

}
/*
Named Return Parameters are generally termed as the named parameters. Golang allows giving the names to the return or
result parameters of the functions in the function signature or definition. Or you can say it is the explicit naming
of the return variables in the function definition.
Basically, it eliminates the requirement of mentioning the variables name with the return statement.
By using named return parameters or named parameters one can only use return keyword at the end of the function
to return the result to the caller.
This concept is generally used when a function has to return multiple values.
So for the user comfort and to enhance the code readability,

ref: https://www.geeksforgeeks.org/named-return-parameters-in-golang/

 */

func functionWithNamedReturnValues(input_a string) (upper_a string,lower_a string){
	fmt.Println("input is : ", input_a)
	if input_a != "" {
		upper_a = strings.ToUpper(input_a)
		lower_a = strings.ToLower(input_a)
		return
	}
	return "EMPTY INPUT","empty input"

}