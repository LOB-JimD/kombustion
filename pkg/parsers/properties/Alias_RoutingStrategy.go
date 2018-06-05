package properties

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import "fmt"

type Alias_RoutingStrategy struct {
	FleetId interface{} `yaml:"FleetId,omitempty"`
	Message interface{} `yaml:"Message,omitempty"`
	Type    interface{} `yaml:"Type"`
}

func (resource Alias_RoutingStrategy) Validate() []error {
	errs := []error{}

	if resource.Type == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Type'"))
	}
	return errs
}