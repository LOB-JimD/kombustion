package properties

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

// EndpointMongoDbSettings Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-endpoint-mongodbsettings.html
type EndpointMongoDbSettings struct {
	AuthMechanism     interface{} `yaml:"AuthMechanism,omitempty"`
	AuthSource        interface{} `yaml:"AuthSource,omitempty"`
	AuthType          interface{} `yaml:"AuthType,omitempty"`
	DatabaseName      interface{} `yaml:"DatabaseName,omitempty"`
	DocsToInvestigate interface{} `yaml:"DocsToInvestigate,omitempty"`
	ExtractDocId      interface{} `yaml:"ExtractDocId,omitempty"`
	NestingLevel      interface{} `yaml:"NestingLevel,omitempty"`
	Password          interface{} `yaml:"Password,omitempty"`
	Port              interface{} `yaml:"Port,omitempty"`
	ServerName        interface{} `yaml:"ServerName,omitempty"`
	Username          interface{} `yaml:"Username,omitempty"`
}

// EndpointMongoDbSettings validation
func (resource EndpointMongoDbSettings) Validate() []error {
	errs := []error{}

	return errs
}
