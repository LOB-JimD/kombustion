package properties

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import "fmt"

// DeploymentGroupS3Location Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codedeploy-deploymentgroup-deployment-revision-s3location.html
type DeploymentGroupS3Location struct {
	Bucket     interface{} `yaml:"Bucket"`
	BundleType interface{} `yaml:"BundleType,omitempty"`
	ETag       interface{} `yaml:"ETag,omitempty"`
	Key        interface{} `yaml:"Key"`
	Version    interface{} `yaml:"Version,omitempty"`
}

// DeploymentGroupS3Location validation
func (resource DeploymentGroupS3Location) Validate() []error {
	errs := []error{}

	if resource.Bucket == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Bucket'"))
	}
	if resource.Key == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Key'"))
	}
	return errs
}