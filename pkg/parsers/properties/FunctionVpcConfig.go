package properties

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import "fmt"

// FunctionVpcConfig Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-function-vpcconfig.html
type FunctionVpcConfig struct {
	SecurityGroupIds interface{} `yaml:"SecurityGroupIds"`
	SubnetIds        interface{} `yaml:"SubnetIds"`
}

// FunctionVpcConfig validation
func (resource FunctionVpcConfig) Validate() []error {
	errs := []error{}

	if resource.SecurityGroupIds == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'SecurityGroupIds'"))
	}
	if resource.SubnetIds == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'SubnetIds'"))
	}
	return errs
}
