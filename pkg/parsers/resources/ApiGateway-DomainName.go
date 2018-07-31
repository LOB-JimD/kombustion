package resources

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import (
	"fmt"
	"github.com/KablamoOSS/kombustion/pkg/parsers/properties"
	"github.com/KablamoOSS/kombustion/types"
	yaml "github.com/KablamoOSS/yaml"
)

// ApiGatewayDomainName Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-domainname.html
type ApiGatewayDomainName struct {
	Type       string                         `yaml:"Type"`
	Properties ApiGatewayDomainNameProperties `yaml:"Properties"`
	Condition  interface{}                    `yaml:"Condition,omitempty"`
	Metadata   interface{}                    `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                    `yaml:"DependsOn,omitempty"`
}

// ApiGatewayDomainName Properties
type ApiGatewayDomainNameProperties struct {
	CertificateArn         interface{}                                 `yaml:"CertificateArn,omitempty"`
	DomainName             interface{}                                 `yaml:"DomainName"`
	RegionalCertificateArn interface{}                                 `yaml:"RegionalCertificateArn,omitempty"`
	EndpointConfiguration  *properties.DomainNameEndpointConfiguration `yaml:"EndpointConfiguration,omitempty"`
}

// NewApiGatewayDomainName constructor creates a new ApiGatewayDomainName
func NewApiGatewayDomainName(properties ApiGatewayDomainNameProperties, deps ...interface{}) ApiGatewayDomainName {
	return ApiGatewayDomainName{
		Type:       "AWS::ApiGateway::DomainName",
		Properties: properties,
		DependsOn:  deps,
	}
}

// ParseApiGatewayDomainName parses ApiGatewayDomainName
func ParseApiGatewayDomainName(
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
	var resource ApiGatewayDomainName
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

// ParseApiGatewayDomainName validator
func (resource ApiGatewayDomainName) Validate() []error {
	return resource.Properties.Validate()
}

// ParseApiGatewayDomainNameProperties validator
func (resource ApiGatewayDomainNameProperties) Validate() []error {
	errors := []error{}
	if resource.DomainName == nil {
		errors = append(errors, fmt.Errorf("Missing required field 'DomainName'"))
	}
	return errors
}