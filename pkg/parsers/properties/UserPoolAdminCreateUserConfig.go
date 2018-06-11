package properties

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

// UserPoolAdminCreateUserConfig Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpool-admincreateuserconfig.html
type UserPoolAdminCreateUserConfig struct {
	AllowAdminCreateUserOnly  interface{}                    `yaml:"AllowAdminCreateUserOnly,omitempty"`
	UnusedAccountValidityDays interface{}                    `yaml:"UnusedAccountValidityDays,omitempty"`
	InviteMessageTemplate     *UserPoolInviteMessageTemplate `yaml:"InviteMessageTemplate,omitempty"`
}

// UserPoolAdminCreateUserConfig validation
func (resource UserPoolAdminCreateUserConfig) Validate() []error {
	errs := []error{}

	return errs
}
