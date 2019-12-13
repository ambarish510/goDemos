package RunTimePolymorphism2

import (
	"fmt"
	//"github.com/ambsflip/goDemos/RunTimePolymorphism2/workflow_drivers"
	"github.com/ambsflip/goDemos/RunTimePolymorphism2/workflow_drivers/argov1alpha"
	"github.com/ambsflip/goDemos/RunTimePolymorphism2/workflow_drivers/argov2beta"
)

type WorkFlowDriver interface {
	TranslatePipeline() string
}

func RunPipeline (){
	var driver WorkFlowDriver
	driver = getWorkflowDriver(2)
	translatedWorkflow := driver.TranslatePipeline()
	fmt.Println(translatedWorkflow)
}


func getWorkflowDriver(pipelineId int) WorkFlowDriver {

	var t3 WorkFlowDriver
	if pipelineId == 1 {
		t3 = &argov1alpha.ArgoV1Alpha{}
	}else {
		t3 = &argov2beta.ArgoV2Beta{}
	}
	return t3
}
