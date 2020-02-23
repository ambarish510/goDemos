package main

import (
	"fmt"
	"github.com/ambsflip/goDemos/models"
	"github.com/argoproj/argo/pkg/apis/workflow/v1alpha1"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"reflect"
	"strconv"
	"strings"
)

/* Read a yaml file and unmarshal into the models.WorkflowSimple
* create instance of the model github.com/argoproj/argo/pkg/apis/workflow/v1alpha1.Workflow
* using the values from models.WorkflowSimple
*
* Slice operations :
*/



const PipelineNameDelimiter = "-"
const EntryStepPrefix = "Entry-step"
const ExitStepPrefix = "Exit-step"

func main (){
	//read the workflow status files recursively
	// create the object for v1alpha1.Workflow
	// invoke for each

	fileName5 := "DataFiles/sample_workflow_status.yaml"
	fileContentBytes5, err := ioutil.ReadFile(fileName5)
	if err != nil {
		fmt.Println(err)
	}
	var newObj5 models.WorkflowSimple
	err = yaml.Unmarshal(fileContentBytes5, &newObj5)
	if err != nil {
		fmt.Printf("\nUnmarshal: %v", err)
	}
	
	var workflow v1alpha1.Workflow
	fmt.Printf("\n%s \n%s \n%s \n", newObj5.Name, newObj5.Namespace, newObj5.Status.Phase)
	workflow.Name = newObj5.Name
	workflow.Namespace = newObj5.Namespace
	workflow.Status.Phase = v1alpha1.NodePhase(newObj5.Status.Phase)

	tempNodes := make(map[string]v1alpha1.NodeStatus)
	for i := 0; i < len(newObj5.Status.Nodes); i++ {
		var tempNodeStatus v1alpha1.NodeStatus
		tempNodeStatus.Name = newObj5.Status.Nodes[i].Name
		tempNodeStatus.ID = newObj5.Status.Nodes[i].ID
		tempNodeStatus.DisplayName = newObj5.Status.Nodes[i].DisplayName
		tempNodeStatus.Phase = v1alpha1.NodePhase(newObj5.Status.Nodes[i].Phase)
		tempNodeStatus.Message = newObj5.Status.Nodes[i].Message
		tempNodeStatus.StartedAt = newObj5.Status.Nodes[i].StartedAt
		tempNodeStatus.FinishedAt = newObj5.Status.Nodes[i].FinishedAt
		//tempNodes = v1alpha1.Nodes{newObj5.Status.Nodes[i].NodeName:tempNodeStatus}
		tempNodes[newObj5.Status.Nodes[i].NodeName] = tempNodeStatus

	}
	workflow.Status.Nodes = tempNodes
	convertArgoStatusToFkStatus(&workflow)

}

func convertArgoStatusToFkStatus(workflow *v1alpha1.Workflow)  {
	pipelineRunDetails := *new(models.PipelineRunDetails)
	splitWfName := strings.Split(workflow.Name, PipelineNameDelimiter)
	pipelineRunDetails.WorkflowName = workflow.Name
	pipelineRunDetails.PipelineName = splitWfName[0]
	pipelineRunDetails.PipelineRunId,_ = strconv.Atoi(splitWfName[1])
	pipelineRunDetails.WorkflowNamespace = workflow.Namespace
	pipelineRunDetails.WorkflowStatus = string(workflow.Status.Phase)


	nodeCount := 1
	for _, v := range workflow.Status.Nodes {
		splitNodeName := strings.Split(v.Name,".")
		if len(splitNodeName) == 3 {
			//node has stage details
			tempStage := *new(models.PipelineRunStage)
			tempStage.StageName = splitNodeName[2]
			tempStage.Status = string(v.Phase)
			tempMap:= make( map[string]string)
			tempMap["pod-name"]= v.ID
			tempStage.Metadata = tempMap
			isStageFound := false
			for i := range pipelineRunDetails.Stages{
				if pipelineRunDetails.Stages[i].StageName == tempStage.StageName {
					isStageFound = true
					tempStage.Steps = pipelineRunDetails.Stages[i].Steps
					tempStage.Entry = pipelineRunDetails.Stages[i].Entry
					tempStage.Exit = pipelineRunDetails.Stages[i].Exit
					pipelineRunDetails.Stages[i] =  tempStage
					break
				}
			}
			if !isStageFound {
				//append a stage
				pipelineRunDetails.Stages = append(pipelineRunDetails.Stages, tempStage)
			}
		}else if len(splitNodeName) == 4 {
			//node has step details
			tempStageName := splitNodeName[2]

			tempStep := *new(models.PipelineRunStep)
			tempStep.StepName = splitNodeName[3]

			tempMap:= make( map[string]string)
			tempMap["pod-name"]= v.ID
			tempStep.Metadata = tempMap

			tempStep.StepStatus = string(v.Phase)
			tempStep.Message = v.Message
			tempStep.StartedAt = v.StartedAt.Time
			tempStep.FinishedAt = v.FinishedAt.Time
			isStageFound := false
			for i := range pipelineRunDetails.Stages{
				if pipelineRunDetails.Stages[i].StageName == tempStageName {
					isStageFound = true
					if strings.HasPrefix(tempStep.StepName, EntryStepPrefix) {
						pipelineRunDetails.Stages[i].Entry = append(pipelineRunDetails.Stages[i].Entry,tempStep)
					}else if strings.HasPrefix(tempStep.StepName, ExitStepPrefix) {
						pipelineRunDetails.Stages[i].Exit = append(pipelineRunDetails.Stages[i].Exit,tempStep)
					}else {
						pipelineRunDetails.Stages[i].Steps = append(pipelineRunDetails.Stages[i].Steps,tempStep)
					}
					break
				}
			}
			if !isStageFound {
				//create a stage and add step
				tempStage := *new(models.PipelineRunStage)
				tempStage.StageName = tempStageName
				if strings.HasPrefix(tempStep.StepName, EntryStepPrefix) {
					tempStage.Entry = append(tempStage.Entry,tempStep)
				}else if strings.HasPrefix(tempStep.StepName, ExitStepPrefix) {
					tempStage.Exit = append(tempStage.Exit,tempStep)
				}else {
					tempStage.Steps = append(tempStage.Steps,tempStep)
				}
				pipelineRunDetails.Stages = append(pipelineRunDetails.Stages, tempStage)
			}
		}
		nodeCount=nodeCount+1
	}
	fmt.Println(pipelineRunDetails)
	stageName,stepName := "flowPocTestStage","TestStep"
	var podName string
	for _, item := range pipelineRunDetails.Stages {
		if item.StageName == stageName {
			if item.Steps == nil || len(item.Steps) == 0 {
				 fmt.Errorf("Nil/Blank Step for the Stage")
			}
			for _, step := range item.Steps {
				if step.StepName == stepName {
					fmt.Println("Type of metadata : %T",step.Metadata)
					v := reflect.ValueOf(step.Metadata)
					fmt.Printf("Type: %v\n", v)

					myMap := step.Metadata.(map[string]string)
					podName = myMap["pod-name"]
					break
				}
			}
		}
		if podName != "" {
			break
		}
	}
}