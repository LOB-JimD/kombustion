package properties

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import "fmt"

// ClusterVolumeSpecification Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-cluster-volumespecification.html
type ClusterVolumeSpecification struct {
	Iops       interface{} `yaml:"Iops,omitempty"`
	SizeInGB   interface{} `yaml:"SizeInGB"`
	VolumeType interface{} `yaml:"VolumeType"`
}

// ClusterVolumeSpecification validation
func (resource ClusterVolumeSpecification) Validate() []error {
	errors := []error{}

	if resource.SizeInGB == nil {
		errors = append(errors, fmt.Errorf("Missing required field 'SizeInGB'"))
	}
	if resource.VolumeType == nil {
		errors = append(errors, fmt.Errorf("Missing required field 'VolumeType'"))
	}
	return errors
}
