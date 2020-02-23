package models


type Config struct {
	GitConfig []Github `yaml:"github"`
}

type Github struct {
	Owner string `yaml:"owner"`
	ApiURL string `yaml:"apiurl"`
	AccessToken string `yaml:"token"`
}

type GitOwnerDetails struct {
	ApiURL string `yaml:"apiurl"`
	AccessToken string `yaml:"token"`
}

type GithubConfig struct {
	GithubOwnerDetailsMap map[string]GitOwnerDetails `yaml:"gihtubconfig"`
}

//type Driver struct {
//	WorkflowDriver string 	`yaml:"workflowdriver" json:"workflowdriver"`
//	PipelineIds   []string 	`yaml:"pipelineids" json:"pipelineids"`
//	ProjectNames   []string `yaml:"projectname" json:"project_name"`
//	TenantNames   []string 	`yaml:"tenantname"  json:"tenant_name"`
//}


