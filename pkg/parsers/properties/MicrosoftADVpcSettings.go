package properties

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import "fmt"

// MicrosoftADVpcSettings Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-directoryservice-microsoftad-vpcsettings.html
type MicrosoftADVpcSettings struct {
	VpcId     interface{} `yaml:"VpcId"`
	SubnetIds interface{} `yaml:"SubnetIds"`
}

// MicrosoftADVpcSettings validation
func (resource MicrosoftADVpcSettings) Validate() []error {
	errs := []error{}

	if resource.VpcId == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'VpcId'"))
	}
	if resource.SubnetIds == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'SubnetIds'"))
	}
	return errs
}
