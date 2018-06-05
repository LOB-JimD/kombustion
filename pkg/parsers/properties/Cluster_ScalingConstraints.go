package properties

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import "fmt"

type Cluster_ScalingConstraints struct {
	MaxCapacity interface{} `yaml:"MaxCapacity"`
	MinCapacity interface{} `yaml:"MinCapacity"`
}

func (resource Cluster_ScalingConstraints) Validate() []error {
	errs := []error{}

	if resource.MaxCapacity == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'MaxCapacity'"))
	}
	if resource.MinCapacity == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'MinCapacity'"))
	}
	return errs
}