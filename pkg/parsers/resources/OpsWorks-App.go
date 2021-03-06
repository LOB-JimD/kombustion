package resources

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import (
	"fmt"
	"github.com/KablamoOSS/kombustion/pkg/parsers/properties"
	"github.com/KablamoOSS/kombustion/types"
	yaml "github.com/KablamoOSS/yaml"
)

// OpsWorksApp Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-app.html
type OpsWorksApp struct {
	Type       string                `yaml:"Type"`
	Properties OpsWorksAppProperties `yaml:"Properties"`
	Condition  interface{}           `yaml:"Condition,omitempty"`
	Metadata   interface{}           `yaml:"Metadata,omitempty"`
	DependsOn  interface{}           `yaml:"DependsOn,omitempty"`
}

// OpsWorksApp Properties
type OpsWorksAppProperties struct {
	Description      interface{}                     `yaml:"Description,omitempty"`
	EnableSsl        interface{}                     `yaml:"EnableSsl,omitempty"`
	Name             interface{}                     `yaml:"Name"`
	Shortname        interface{}                     `yaml:"Shortname,omitempty"`
	StackId          interface{}                     `yaml:"StackId"`
	Type             interface{}                     `yaml:"Type"`
	SslConfiguration *properties.AppSslConfiguration `yaml:"SslConfiguration,omitempty"`
	AppSource        *properties.AppSource           `yaml:"AppSource,omitempty"`
	Attributes       interface{}                     `yaml:"Attributes,omitempty"`
	DataSources      interface{}                     `yaml:"DataSources,omitempty"`
	Domains          interface{}                     `yaml:"Domains,omitempty"`
	Environment      interface{}                     `yaml:"Environment,omitempty"`
}

// NewOpsWorksApp constructor creates a new OpsWorksApp
func NewOpsWorksApp(properties OpsWorksAppProperties, deps ...interface{}) OpsWorksApp {
	return OpsWorksApp{
		Type:       "AWS::OpsWorks::App",
		Properties: properties,
		DependsOn:  deps,
	}
}

// ParseOpsWorksApp parses OpsWorksApp
func ParseOpsWorksApp(
	name string,
	data string,
) (
	source string,
	conditions types.TemplateObject,
	metadata types.TemplateObject,
	mappings types.TemplateObject,
	outputs types.TemplateObject,
	parameters types.TemplateObject,
	resources types.TemplateObject,
	transform types.TemplateObject,
	errors []error,
) {
	source = "kombustion-core-resources"
	var resource OpsWorksApp
	err := yaml.Unmarshal([]byte(data), &resource)

	if err != nil {
		errors = append(errors, err)
		return
	}

	if validateErrs := resource.Properties.Validate(); len(errors) > 0 {
		errors = append(errors, validateErrs...)
		return
	}

	resources = types.TemplateObject{name: resource}

	return
}

// ParseOpsWorksApp validator
func (resource OpsWorksApp) Validate() []error {
	return resource.Properties.Validate()
}

// ParseOpsWorksAppProperties validator
func (resource OpsWorksAppProperties) Validate() []error {
	errors := []error{}
	if resource.Name == nil {
		errors = append(errors, fmt.Errorf("Missing required field 'Name'"))
	}
	if resource.StackId == nil {
		errors = append(errors, fmt.Errorf("Missing required field 'StackId'"))
	}
	if resource.Type == nil {
		errors = append(errors, fmt.Errorf("Missing required field 'Type'"))
	}
	return errors
}
