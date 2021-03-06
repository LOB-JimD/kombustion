package properties

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

// SpotFleetEbsBlockDevice Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-spotfleet-spotfleetrequestconfigdata-launchspecifications-blockdevicemappings-ebs.html
type SpotFleetEbsBlockDevice struct {
	DeleteOnTermination interface{} `yaml:"DeleteOnTermination,omitempty"`
	Encrypted           interface{} `yaml:"Encrypted,omitempty"`
	Iops                interface{} `yaml:"Iops,omitempty"`
	SnapshotId          interface{} `yaml:"SnapshotId,omitempty"`
	VolumeSize          interface{} `yaml:"VolumeSize,omitempty"`
	VolumeType          interface{} `yaml:"VolumeType,omitempty"`
}

// SpotFleetEbsBlockDevice validation
func (resource SpotFleetEbsBlockDevice) Validate() []error {
	errors := []error{}

	return errors
}
