package models

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

type WorkflowSimple struct {
	Name      string `yaml:"name"`
	Namespace string `yaml:"namespace"`
	Status    StatusSimple `yaml:"status"`
}

type StatusSimple struct {
	Phase string        `yaml:"phase"`
	Nodes []NodesSimple `yaml:"nodes"`
}

type NodesSimple struct {
	NodeName    string `yaml:"nodename"`
	Name        string `yaml:"name"`
	DisplayName string `yaml:"displayname"`
	ID          string `yaml:"id"`
	Phase       string `yaml:"phase"`
	Message		string `yaml:"message"`
	StartedAt	metav1.Time `yaml:"startedat"`
	FinishedAt  metav1.Time `yaml:"finishedat"`
}