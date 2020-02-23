package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/user"
)

func main(){
	getWorkingDir()
	filename := "FileOperations/input1.yaml"
	_ = readFile(filename)
	_ = readFile2(filename)
}
func readFile(filename string) error{
	dat, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
		return err
	}
	fileContent :=string(dat)
	fmt.Print("File Content : \n", fileContent)
	return nil
}

func readFile2(filename string) error{
	// Open our jsonFile
	jsonFile, err := os.Open(filename)
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println("\nSuccessfully Opened file")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	// read our opened file as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)
	fileContent :=string(byteValue)
	fmt.Print("File Content : \n", fileContent)
	return nil

}

func getWorkingDir(){
	path,_ := os.Getwd()
	fmt.Println("Working Directory",path)
	usr,_ := user.Current()
	fmt.Println( "Home Direcory", usr.HomeDir )
}

//function isFileExists can be used to check the existence of directories also
func isFileExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil { return true, nil }
	if os.IsNotExist(err) { return false, nil }
	return true, err
}

