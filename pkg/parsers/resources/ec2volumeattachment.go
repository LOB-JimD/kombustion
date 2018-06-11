package resources

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import (
	"fmt"
	"github.com/KablamoOSS/kombustion/types"
	yaml "github.com/KablamoOSS/yaml"
	"log"
)

// EC2VolumeAttachment Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-ebs-volumeattachment.html
type EC2VolumeAttachment struct {
	Type       string                        `yaml:"Type"`
	Properties EC2VolumeAttachmentProperties `yaml:"Properties"`
	Condition  interface{}                   `yaml:"Condition,omitempty"`
	Metadata   interface{}                   `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                   `yaml:"DependsOn,omitempty"`
}

// EC2VolumeAttachment Properties
type EC2VolumeAttachmentProperties struct {
	Device     interface{} `yaml:"Device"`
	InstanceId interface{} `yaml:"InstanceId"`
	VolumeId   interface{} `yaml:"VolumeId"`
}

// NewEC2VolumeAttachment constructor creates a new EC2VolumeAttachment
func NewEC2VolumeAttachment(properties EC2VolumeAttachmentProperties, deps ...interface{}) EC2VolumeAttachment {
	return EC2VolumeAttachment{
		Type:       "AWS::EC2::VolumeAttachment",
		Properties: properties,
		DependsOn:  deps,
	}
}

// ParseEC2VolumeAttachment parses EC2VolumeAttachment
func ParseEC2VolumeAttachment(ctx map[string]interface{}, name string, data string) (cf types.TemplateObject, err error) {
	var resource EC2VolumeAttachment
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: EC2VolumeAttachment - ", err)
		}
		return
	}
	cf = types.TemplateObject{name: resource}
	return
}

// ParseEC2VolumeAttachment validator
func (resource EC2VolumeAttachment) Validate() []error {
	return resource.Properties.Validate()
}

// ParseEC2VolumeAttachmentProperties validator
func (resource EC2VolumeAttachmentProperties) Validate() []error {
	errs := []error{}
	if resource.Device == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Device'"))
	}
	if resource.InstanceId == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'InstanceId'"))
	}
	if resource.VolumeId == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'VolumeId'"))
	}
	return errs
}
