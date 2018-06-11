package properties

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import "fmt"

// ScalingPlanTagFilter Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-autoscalingplans-scalingplan-tagfilter.html
type ScalingPlanTagFilter struct {
	Key    interface{} `yaml:"Key"`
	Values interface{} `yaml:"Values,omitempty"`
}

// ScalingPlanTagFilter validation
func (resource ScalingPlanTagFilter) Validate() []error {
	errs := []error{}

	if resource.Key == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Key'"))
	}
	return errs
}
