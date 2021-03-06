package properties

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import "fmt"

// ReceiptRuleAddHeaderAction Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-receiptrule-addheaderaction.html
type ReceiptRuleAddHeaderAction struct {
	HeaderName  interface{} `yaml:"HeaderName"`
	HeaderValue interface{} `yaml:"HeaderValue"`
}

// ReceiptRuleAddHeaderAction validation
func (resource ReceiptRuleAddHeaderAction) Validate() []error {
	errors := []error{}

	if resource.HeaderName == nil {
		errors = append(errors, fmt.Errorf("Missing required field 'HeaderName'"))
	}
	if resource.HeaderValue == nil {
		errors = append(errors, fmt.Errorf("Missing required field 'HeaderValue'"))
	}
	return errors
}
