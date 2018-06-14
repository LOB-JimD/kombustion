package properties

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import "fmt"

// DeliveryStreamProcessorParameter Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-processorparameter.html
type DeliveryStreamProcessorParameter struct {
	ParameterName  interface{} `yaml:"ParameterName"`
	ParameterValue interface{} `yaml:"ParameterValue"`
}

// DeliveryStreamProcessorParameter validation
func (resource DeliveryStreamProcessorParameter) Validate() []error {
	errs := []error{}

	if resource.ParameterName == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'ParameterName'"))
	}
	if resource.ParameterValue == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'ParameterValue'"))
	}
	return errs
}