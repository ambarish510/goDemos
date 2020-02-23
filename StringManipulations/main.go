package main

import (
	"fmt"
	"strings"
)

func main() {
	s := strings.Split("127.0.0.1:5432", ":")
	ip, port := s[0], s[1]
	fmt.Printf("IP %s  PORT %s",ip, port)
	ipSplit := strings.Split(s[0],".")
	fmt.Printf("\nIP splits %s %s %s %s ",ipSplit[0], ipSplit[1],ipSplit[2],ipSplit[3])
	fmt.Printf("\nLength of array ipSplit %d ", len(ipSplit))

	splitWfName := strings.Split("samplepipeline-115", "-")

	fmt.Printf("\nName  %s  Id %s",splitWfName[0], splitWfName[1])

}

