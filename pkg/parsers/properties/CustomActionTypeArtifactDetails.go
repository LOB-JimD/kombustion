package properties

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import "fmt"

// CustomActionTypeArtifactDetails Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-customactiontype-artifactdetails.html
type CustomActionTypeArtifactDetails struct {
	MaximumCount interface{} `yaml:"MaximumCount"`
	MinimumCount interface{} `yaml:"MinimumCount"`
}

// CustomActionTypeArtifactDetails validation
func (resource CustomActionTypeArtifactDetails) Validate() []error {
	errs := []error{}

	if resource.MaximumCount == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'MaximumCount'"))
	}
	if resource.MinimumCount == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'MinimumCount'"))
	}
	return errs
}
