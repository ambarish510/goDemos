package main


import (
	"fmt"
	"github.com/ambsflip/goDemos/TestSamples/SamplePackage"
)
//The tests are written inside the same package. SamplePackage.
// run the command go test while at the package.
// test suites are identified by files ending woth _test.go
// each functions inside the file that starts with Test_ is considered as testcases.

func main(){
	avg1 := SamplePackage.FindAvg(2.4,3.6)
	fmt.Printf("Average is : %f ", avg1)
}