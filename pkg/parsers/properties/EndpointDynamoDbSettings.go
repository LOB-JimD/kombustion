package properties

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

// EndpointDynamoDbSettings Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-endpoint-dynamodbsettings.html
type EndpointDynamoDbSettings struct {
	ServiceAccessRoleArn interface{} `yaml:"ServiceAccessRoleArn,omitempty"`
}

// EndpointDynamoDbSettings validation
func (resource EndpointDynamoDbSettings) Validate() []error {
	errors := []error{}

	return errors
}
