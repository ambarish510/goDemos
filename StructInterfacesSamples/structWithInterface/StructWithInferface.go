package structWithInterface

type PipelineRunRequest struct {
	Id                 int         `json:"id"`
	Comments           string      `json:"comments"` //
	PipelineSpec       interface{} `json:"pipeline_spec"`
	PipelineRunDetails interface{} `json:"pipeline_run_details"` //
	//PipelineRunParams []step.InputParam `json:"pipeline_run_params"`
	Attributes			interface{}		`json:"attr"`
}