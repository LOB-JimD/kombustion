package properties

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import "fmt"

// MetricFilterMetricTransformation Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-logs-metricfilter-metrictransformation.html
type MetricFilterMetricTransformation struct {
	MetricName      interface{} `yaml:"MetricName"`
	MetricNamespace interface{} `yaml:"MetricNamespace"`
	MetricValue     interface{} `yaml:"MetricValue"`
}

// MetricFilterMetricTransformation validation
func (resource MetricFilterMetricTransformation) Validate() []error {
	errs := []error{}

	if resource.MetricName == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'MetricName'"))
	}
	if resource.MetricNamespace == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'MetricNamespace'"))
	}
	if resource.MetricValue == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'MetricValue'"))
	}
	return errs
}