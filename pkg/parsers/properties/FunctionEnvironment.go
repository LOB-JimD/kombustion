package properties

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

// FunctionEnvironment Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-function-environment.html
type FunctionEnvironment struct {
	Variables interface{} `yaml:"Variables,omitempty"`
}

// FunctionEnvironment validation
func (resource FunctionEnvironment) Validate() []error {
	errors := []error{}

	return errors
}
