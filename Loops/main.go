package main

import "fmt"

func main(){

	SimpleForLoopBreakContinue(1,3)
	ForLoopArray()
}
func ForLoopArray() {
	a := [...]float64{67.7, 89.8, 21, 78}
	sum := float64(0)
	for i, v := range a {//range returns both the index and value
		fmt.Printf("%d the element of a is %.2f\n", i, v)
		sum += v
	}
	fmt.Println("\nsum of all elements of a",sum)
}
func SimpleForLoopBreakContinue(breakValue int, continueValue int){
	for i := 1;  i<=10; i++ {
		if i== continueValue {
			continue
		}
		fmt.Printf("Welcome %d times\n",i)
		if i == breakValue {
			break
		}
	}
}