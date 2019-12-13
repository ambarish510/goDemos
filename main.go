package main

import (
	"fmt"
	rtp "github.com/ambsflip/goDemos/RuntimePolymorphism"
	rtp2 "github.com/ambsflip/goDemos/RunTimePolymorphism2"
)
func main() {
	//fmt.Println("Hello world")
	//restAPIcalls.GetAPICall(TimeOut, getURL, headers)
	fmt.Println("Tax Percentage :" , rtp.GetTaxSystem("US"))

	rtp2.RunPipeline()
}
