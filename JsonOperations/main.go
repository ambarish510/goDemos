package main

import (
    "encoding/json"
    "fmt"
    "github.com/ambsflip/goDemos/models"
    "io/ioutil"
    "os"
    "reflect"
    "strconv"
)

func main() {
    // Open our jsonFile
    jsonFile, err := os.Open("DataFiles/users.json")
    // if we os.Open returns an error then handle it
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println("Successfully Opened users.json")
    // defer the closing of our jsonFile so that we can parse it later on
    defer jsonFile.Close()

    // read our opened file as a byte array.
    byteValue, _ := ioutil.ReadAll(jsonFile)
    // we initialize our Users array
    var users models.Users
    // we unmarshal our byteArray which contains our
    // jsonFile's content into 'users' which we defined above
    json.Unmarshal(byteValue, &users)

    // we iterate through every user within our users array and
    // print out the user Type, their name, and their facebook url
    // as just an example
    for i := 0; i < len(users.Users); i++ {
        fmt.Println("User Type: " + users.Users[i].Type)
        fmt.Println("User Age: " + strconv.Itoa(users.Users[i].Age))
        fmt.Println("User Name: " + users.Users[i].Name)
        fmt.Println("Facebook Url: " + users.Users[i].Social.Facebook)
    }

    var result map[string]interface{}
    json.Unmarshal([]byte(byteValue), &result)
    fmt.Println(result["users"])

    /*
    *
    */
    jsonFile2, err2 := os.Open("DataFiles/driverselect.json")
    if err2 != nil {
        fmt.Println(err2)
    }
    fmt.Println("Successfully Opened driverselect.json")
    defer jsonFile2.Close()
    byteValue2, _ := ioutil.ReadAll(jsonFile2)
    var result2 map[string]interface{}
    json.Unmarshal([]byte(byteValue2), &result2)
    fmt.Println(result2["pipelineids"])
    fmt.Println(reflect.TypeOf(result2["pipelineids"]))
    /*
     *
     */
    jsonFile3, err3 := os.Open("DataFiles/driverconfig.json")
    if err2 != nil {
        fmt.Println(err3)
    }
    fmt.Println("Successfully Opened driverconfig.json")
    defer jsonFile3.Close()
    byteValue3, _ := ioutil.ReadAll(jsonFile3)
    //var driverConfigs DriverConfigs
    //json.Unmarshal(byteValue3, &driverConfigs)
    var drivers []models.Driver
    json.Unmarshal(byteValue3, &drivers)
   // fmt.Println(drivers[1].PipelineIds[0])
    //fmt.Println(reflect.TypeOf(drivers[1].PipelineIds))

    myPipelineId := "101"
    myProjectId := "1001"
    myTenantId := "10001"
    mylabel := "groceries"
    var myDriver string
    for _, driver := range drivers{

        if ArrayContainsString(driver.PipelineIds,myPipelineId) ||
            ArrayContainsString(driver.ProjectIds,myProjectId) ||
            ArrayContainsString(driver.TenantIds,myTenantId) ||
            ArrayContainsString(driver.Labels,mylabel) {
            myDriver = driver.WorkflowDriver
            break
        }
    }
    fmt.Println(myDriver)
}


func ArrayContainsString(arr []string, str string) bool {
    for _, a := range arr {
        if a == str {
            return true
        }
    }
    return false
}

func itemExists(arrayType interface{}, item interface{}) bool {
    arr := reflect.ValueOf(arrayType)

    if arr.Kind() != reflect.Array {
        panic("Invalid data-type")
    }

    for i := 0; i < arr.Len(); i++ {
        if arr.Index(i).Interface() == item {
            return true
        }
    }

    return false
}