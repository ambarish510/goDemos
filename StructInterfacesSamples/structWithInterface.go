package main

import (
	"encoding/json"
	"fmt"
	"github.com/ambsflip/goDemos/StructInterfacesSamples/structWithInterface"
	"github.com/ambsflip/goDemos/models"
	"io/ioutil"
	"os"
)

func main(){


	temp := structWithInterface.PipelineRunRequest{
		Id:       10,
		Comments: "samplecomment",
	}
	fmt.Println(temp)

	jsonFile, err := os.Open("DataFiles/users.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened users.json")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var users models.Users
	json.Unmarshal(byteValue, &users)
	temp.PipelineRunDetails=users

	tempMap:= make( map[string]string)
	tempMap["driver"] = "argov1"

	temp.Attributes =tempMap
	fmt.Println(temp)
}