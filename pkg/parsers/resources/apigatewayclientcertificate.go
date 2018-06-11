package resources

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import (
	"github.com/KablamoOSS/kombustion/types"
	yaml "github.com/KablamoOSS/yaml"
	"log"
)

// ApiGatewayClientCertificate Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-clientcertificate.html
type ApiGatewayClientCertificate struct {
	Type       string                                `yaml:"Type"`
	Properties ApiGatewayClientCertificateProperties `yaml:"Properties"`
	Condition  interface{}                           `yaml:"Condition,omitempty"`
	Metadata   interface{}                           `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                           `yaml:"DependsOn,omitempty"`
}

// ApiGatewayClientCertificate Properties
type ApiGatewayClientCertificateProperties struct {
	Description interface{} `yaml:"Description,omitempty"`
}

// NewApiGatewayClientCertificate constructor creates a new ApiGatewayClientCertificate
func NewApiGatewayClientCertificate(properties ApiGatewayClientCertificateProperties, deps ...interface{}) ApiGatewayClientCertificate {
	return ApiGatewayClientCertificate{
		Type:       "AWS::ApiGateway::ClientCertificate",
		Properties: properties,
		DependsOn:  deps,
	}
}

// ParseApiGatewayClientCertificate parses ApiGatewayClientCertificate
func ParseApiGatewayClientCertificate(ctx map[string]interface{}, name string, data string) (cf types.TemplateObject, err error) {
	var resource ApiGatewayClientCertificate
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: ApiGatewayClientCertificate - ", err)
		}
		return
	}
	cf = types.TemplateObject{name: resource}
	return
}

// ParseApiGatewayClientCertificate validator
func (resource ApiGatewayClientCertificate) Validate() []error {
	return resource.Properties.Validate()
}

// ParseApiGatewayClientCertificateProperties validator
func (resource ApiGatewayClientCertificateProperties) Validate() []error {
	errs := []error{}
	return errs
}
