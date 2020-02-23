package main

import (
	"fmt"
	"github.com/ambsflip/goDemos/models"
	yaml "gopkg.in/yaml.v2"
	"io/ioutil"
)

// if the struct is private, then the handling function needs to defined as part of struct itself
type conf struct {
	Hits int64 `yaml:"hits"`
	Time int64 `yaml:"time"`
	Name string `yaml:"name""`
}
var data1 = `
hits: 5
time: 5000000
name: "abcd"
`
func (c *conf) getConf() *conf {

	yamlFile := []byte(data1)
	err := yaml.Unmarshal(yamlFile, c)
	if err != nil {
		fmt.Printf("Unmarshal: %v", err)
	}

	return c
}

func main () {
	//unmarshal handling function defined as part of struct itself
	var c conf
	c.getConf()
	fmt.Println(c)

	//Simple Flat Yaml
	fileName1 := "DataFiles/sample_flat.yaml"
	fileContentBytes1, err := ioutil.ReadFile(fileName1)
	if err != nil {
		fmt.Println(err)
	}
	var newObj1 models.NodesSimple
	err = yaml.Unmarshal(fileContentBytes1, &newObj1)

	if err != nil {
		fmt.Printf("Unmarshal: %v", err)
	}
	fmt.Printf("\n%s\n%s\n%s\n%s\n%s",newObj1.Name,newObj1.DisplayName,newObj1.NodeName,newObj1.ID,newObj1.Phase)
    fmt.Println()


	fileName2 := "DataFiles/sample_kv_pair.yaml"
	fileContentBytes2, err := ioutil.ReadFile(fileName2)
	if err != nil {
		fmt.Println(err)
	}
	var newObj2 models.Config
	err = yaml.Unmarshal(fileContentBytes2, &newObj2)
	if err != nil {
		fmt.Println(err)
	}
	for i :=0 ; i < len(newObj2.GitConfig); i++ {
		fmt.Printf("Owner: %s \t URL: %s \t Token: %s \n",newObj2.GitConfig[i].Owner,newObj2.GitConfig[i].ApiURL,newObj2.GitConfig[i].AccessToken)
	}

	//yaml with Map
	file4 := "DataFiles/sample_kv_pair.yaml"
	fileContentBytes4, err := ioutil.ReadFile(file4)
	if err != nil {
		fmt.Println(err)
	}
	var newObj4 models.GithubConfig
	err = yaml.Unmarshal(fileContentBytes4, &newObj4)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("\n")
	for k,v := range newObj4.GithubOwnerDetailsMap{
		fmt.Printf("Owner: %s \t URL: %s \t Token: %s \n",k,v.ApiURL,v.AccessToken)
	}
	fmt.Println(newObj4.GithubOwnerDetailsMap["ambsflip"].AccessToken,newObj4.GithubOwnerDetailsMap["ambsflip"].ApiURL)


	//yaml with array in value for one key
	fileName3 := "DataFiles/sample_array_in_value.yaml"
	fileContentBytes3, err := ioutil.ReadFile(fileName3)
	if err != nil {
		fmt.Println(err)
	}
	var newObj3 models.StatusSimple
	err = yaml.Unmarshal(fileContentBytes3, &newObj3)
	if err != nil {
		fmt.Printf("Unmarshal: %v", err)
	}
	fmt.Printf("Phase: %s\n",newObj3.Phase)
	for i :=0 ; i < len(newObj3.Nodes); i++ {
		fmt.Println("===============================================")
		fmt.Println("Pod Name",newObj3.Nodes[i].NodeName)
		fmt.Println("Node name",newObj3.Nodes[i].Name)
		fmt.Println("Display Name",newObj3.Nodes[i].DisplayName)
		fmt.Println("Phase ",newObj3.Nodes[i].Phase)
		fmt.Println("node ID",newObj3.Nodes[i].ID)
	}


	fileName5 := "DataFiles/sample_wf_status_array.yaml"
	fileContentBytes5, err := ioutil.ReadFile(fileName5)
	if err != nil {
		fmt.Println(err)
	}
	var newObj5 []models.WorkflowSimple
	err = yaml.Unmarshal(fileContentBytes5, &newObj5)
	if err != nil {
		fmt.Printf("\nUnmarshal: %v", err)
	}
	for j :=0 ; j < len(newObj5); j++ {
		fmt.Println("******************************************************")
		fmt.Printf("\n%s \n%s \n%s \n", newObj5[j].Name, newObj5[j].Namespace, newObj5[j].Status.Phase)
		for i := 0; i < len(newObj5[j].Status.Nodes); i++ {
			fmt.Println("###############################################")
			fmt.Println("Pod Name", newObj5[j].Status.Nodes[i].NodeName)
			fmt.Println("Node name", newObj5[j].Status.Nodes[i].Name)
			fmt.Println("Display Name", newObj5[j].Status.Nodes[i].DisplayName)
			fmt.Println("Phase ", newObj5[j].Status.Nodes[i].Phase)
			fmt.Println("node ID", newObj5[j].Status.Nodes[i].ID)
		}
	}
}

