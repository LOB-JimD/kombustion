package properties

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

// EnvironmentTier Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-beanstalk-environment-tier.html
type EnvironmentTier struct {
	Name    interface{} `yaml:"Name,omitempty"`
	Type    interface{} `yaml:"Type,omitempty"`
	Version interface{} `yaml:"Version,omitempty"`
}

// EnvironmentTier validation
func (resource EnvironmentTier) Validate() []error {
	errs := []error{}

	return errs
}
