package properties

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

// DomainEncryptionAtRestOptions Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticsearch-domain-encryptionatrestoptions.html
type DomainEncryptionAtRestOptions struct {
	Enabled  interface{} `yaml:"Enabled,omitempty"`
	KmsKeyId interface{} `yaml:"KmsKeyId,omitempty"`
}

// DomainEncryptionAtRestOptions validation
func (resource DomainEncryptionAtRestOptions) Validate() []error {
	errors := []error{}

	return errors
}
