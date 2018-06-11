package properties

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

// ConfigRuleScope Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-config-configrule-scope.html
type ConfigRuleScope struct {
	ComplianceResourceId    interface{} `yaml:"ComplianceResourceId,omitempty"`
	TagKey                  interface{} `yaml:"TagKey,omitempty"`
	TagValue                interface{} `yaml:"TagValue,omitempty"`
	ComplianceResourceTypes interface{} `yaml:"ComplianceResourceTypes,omitempty"`
}

// ConfigRuleScope validation
func (resource ConfigRuleScope) Validate() []error {
	errs := []error{}

	return errs
}
