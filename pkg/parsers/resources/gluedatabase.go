package resources

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import (
	"fmt"
	"github.com/KablamoOSS/kombustion/pkg/parsers/properties"
	"github.com/KablamoOSS/kombustion/types"
	yaml "github.com/KablamoOSS/yaml"
	"log"
)

// GlueDatabase Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-database.html
type GlueDatabase struct {
	Type       string                 `yaml:"Type"`
	Properties GlueDatabaseProperties `yaml:"Properties"`
	Condition  interface{}            `yaml:"Condition,omitempty"`
	Metadata   interface{}            `yaml:"Metadata,omitempty"`
	DependsOn  interface{}            `yaml:"DependsOn,omitempty"`
}

// GlueDatabase Properties
type GlueDatabaseProperties struct {
	CatalogId     interface{}                       `yaml:"CatalogId"`
	DatabaseInput *properties.DatabaseDatabaseInput `yaml:"DatabaseInput"`
}

// NewGlueDatabase constructor creates a new GlueDatabase
func NewGlueDatabase(properties GlueDatabaseProperties, deps ...interface{}) GlueDatabase {
	return GlueDatabase{
		Type:       "AWS::Glue::Database",
		Properties: properties,
		DependsOn:  deps,
	}
}

// ParseGlueDatabase parses GlueDatabase
func ParseGlueDatabase(ctx map[string]interface{}, name string, data string) (cf types.TemplateObject, err error) {
	var resource GlueDatabase
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: GlueDatabase - ", err)
		}
		return
	}
	cf = types.TemplateObject{name: resource}
	return
}

// ParseGlueDatabase validator
func (resource GlueDatabase) Validate() []error {
	return resource.Properties.Validate()
}

// ParseGlueDatabaseProperties validator
func (resource GlueDatabaseProperties) Validate() []error {
	errs := []error{}
	if resource.CatalogId == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'CatalogId'"))
	}
	if resource.DatabaseInput == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'DatabaseInput'"))
	} else {
		errs = append(errs, resource.DatabaseInput.Validate()...)
	}
	return errs
}
