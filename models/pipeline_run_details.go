package models

import "time"

type PipelineRunDetails struct {
	PipelineRunId      	int    `json:"pipeline-run-id"`
	PipelineName		string `json:"pipeline-name"`
	WorkflowName		string `json:"workflow-name"`
	WorkflowNamespace	string `json:"workflow-namespace"`
	WorkflowStatus		string `json:"workflow-status"`
	Stages 				[]PipelineRunStage `json:"stages"`

}

type PipelineRunStage struct{
	StageName string `json:"stage-name"`
	Status string `json:"status"`
	Steps []PipelineRunStep `json:"steps"`
	Entry []PipelineRunStep `json:"entry"`
	Exit []PipelineRunStep `json:"exit"`
	Metadata	interface{}		`json:"metadata"`
}

type PipelineRunStep struct {
	StepName string `json:"step-name"`
	Metadata	interface{}		`json:"metadata"`
	StepStatus string `json:"step-status"`
	StepOutput string `json:"output"`
	StepArtifact string `json:"artifact"`
	Message string `json:"message"`
	StartedAt time.Time `json:"started-at"`
	FinishedAt time.Time `json:"finished-at"`
}
