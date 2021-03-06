package resources

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import (
	"fmt"
	"github.com/KablamoOSS/kombustion/types"
	yaml "github.com/KablamoOSS/yaml"
)

// ApiGatewayRequestValidator Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-requestvalidator.html
type ApiGatewayRequestValidator struct {
	Type       string                               `yaml:"Type"`
	Properties ApiGatewayRequestValidatorProperties `yaml:"Properties"`
	Condition  interface{}                          `yaml:"Condition,omitempty"`
	Metadata   interface{}                          `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                          `yaml:"DependsOn,omitempty"`
}

// ApiGatewayRequestValidator Properties
type ApiGatewayRequestValidatorProperties struct {
	Name                      interface{} `yaml:"Name,omitempty"`
	RestApiId                 interface{} `yaml:"RestApiId"`
	ValidateRequestBody       interface{} `yaml:"ValidateRequestBody,omitempty"`
	ValidateRequestParameters interface{} `yaml:"ValidateRequestParameters,omitempty"`
}

// NewApiGatewayRequestValidator constructor creates a new ApiGatewayRequestValidator
func NewApiGatewayRequestValidator(properties ApiGatewayRequestValidatorProperties, deps ...interface{}) ApiGatewayRequestValidator {
	return ApiGatewayRequestValidator{
		Type:       "AWS::ApiGateway::RequestValidator",
		Properties: properties,
		DependsOn:  deps,
	}
}

// ParseApiGatewayRequestValidator parses ApiGatewayRequestValidator
func ParseApiGatewayRequestValidator(
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
	var resource ApiGatewayRequestValidator
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

// ParseApiGatewayRequestValidator validator
func (resource ApiGatewayRequestValidator) Validate() []error {
	return resource.Properties.Validate()
}

// ParseApiGatewayRequestValidatorProperties validator
func (resource ApiGatewayRequestValidatorProperties) Validate() []error {
	errors := []error{}
	if resource.RestApiId == nil {
		errors = append(errors, fmt.Errorf("Missing required field 'RestApiId'"))
	}
	return errors
}
