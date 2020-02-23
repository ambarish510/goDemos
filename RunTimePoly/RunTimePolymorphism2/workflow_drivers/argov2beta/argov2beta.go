package argov2beta

import (
	"fmt"
	"reflect"
)

type ArgoV2Beta struct {
}
func (t *ArgoV2Beta) TranslatePipeline() string {
	typeObj := reflect.TypeOf(t)
	fmt.Println(typeObj.Name())
	fmt.Println(typeObj.String())

	return "argov2beta"
}

