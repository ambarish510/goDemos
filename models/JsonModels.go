package models


// Users struct which contains an array of users
type Users struct {
	Users []User `json:"users"`
}

// User struct which contains a name, a type and a list of social links
type User struct {
	Name   string `json:"name"`
	Type   string `json:"type"`
	Age    int    `json:"Age"`
	Social Social `json:"social"`
}

// Social struct which contains a list of links
type Social struct {
	Facebook string `json:"facebook"`
	Twitter  string `json:"twitter"`
}

type Driver struct {
	WorkflowDriver string `json:workflowdriver`
	PipelineIds   []string `json:"pipelineids"`
	ProjectIds   []string `json:"projectids"`
	TenantIds   []string `json:"tenantids"`
	Labels   []string `json:"labels"`
}
type DriverConfigs struct {
	drivers []Driver
}
