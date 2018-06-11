package resources

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import (
	"fmt"
	"github.com/KablamoOSS/kombustion/types"
	yaml "github.com/KablamoOSS/yaml"
	"log"
)

// RDSDBSecurityGroupIngress Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-security-group-ingress.html
type RDSDBSecurityGroupIngress struct {
	Type       string                              `yaml:"Type"`
	Properties RDSDBSecurityGroupIngressProperties `yaml:"Properties"`
	Condition  interface{}                         `yaml:"Condition,omitempty"`
	Metadata   interface{}                         `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                         `yaml:"DependsOn,omitempty"`
}

// RDSDBSecurityGroupIngress Properties
type RDSDBSecurityGroupIngressProperties struct {
	CIDRIP                  interface{} `yaml:"CIDRIP,omitempty"`
	DBSecurityGroupName     interface{} `yaml:"DBSecurityGroupName"`
	EC2SecurityGroupId      interface{} `yaml:"EC2SecurityGroupId,omitempty"`
	EC2SecurityGroupName    interface{} `yaml:"EC2SecurityGroupName,omitempty"`
	EC2SecurityGroupOwnerId interface{} `yaml:"EC2SecurityGroupOwnerId,omitempty"`
}

// NewRDSDBSecurityGroupIngress constructor creates a new RDSDBSecurityGroupIngress
func NewRDSDBSecurityGroupIngress(properties RDSDBSecurityGroupIngressProperties, deps ...interface{}) RDSDBSecurityGroupIngress {
	return RDSDBSecurityGroupIngress{
		Type:       "AWS::RDS::DBSecurityGroupIngress",
		Properties: properties,
		DependsOn:  deps,
	}
}

// ParseRDSDBSecurityGroupIngress parses RDSDBSecurityGroupIngress
func ParseRDSDBSecurityGroupIngress(ctx map[string]interface{}, name string, data string) (cf types.TemplateObject, err error) {
	var resource RDSDBSecurityGroupIngress
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: RDSDBSecurityGroupIngress - ", err)
		}
		return
	}
	cf = types.TemplateObject{name: resource}
	return
}

// ParseRDSDBSecurityGroupIngress validator
func (resource RDSDBSecurityGroupIngress) Validate() []error {
	return resource.Properties.Validate()
}

// ParseRDSDBSecurityGroupIngressProperties validator
func (resource RDSDBSecurityGroupIngressProperties) Validate() []error {
	errs := []error{}
	if resource.DBSecurityGroupName == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'DBSecurityGroupName'"))
	}
	return errs
}
