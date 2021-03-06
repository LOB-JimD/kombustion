package properties

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import "fmt"

// PipelineInputArtifact Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-pipeline-stages-actions-inputartifacts.html
type PipelineInputArtifact struct {
	Name interface{} `yaml:"Name"`
}

// PipelineInputArtifact validation
func (resource PipelineInputArtifact) Validate() []error {
	errors := []error{}

	if resource.Name == nil {
		errors = append(errors, fmt.Errorf("Missing required field 'Name'"))
	}
	return errors
}
