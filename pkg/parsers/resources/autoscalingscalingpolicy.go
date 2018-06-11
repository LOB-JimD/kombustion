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

// AutoScalingScalingPolicy Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-policy.html
type AutoScalingScalingPolicy struct {
	Type       string                             `yaml:"Type"`
	Properties AutoScalingScalingPolicyProperties `yaml:"Properties"`
	Condition  interface{}                        `yaml:"Condition,omitempty"`
	Metadata   interface{}                        `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                        `yaml:"DependsOn,omitempty"`
}

// AutoScalingScalingPolicy Properties
type AutoScalingScalingPolicyProperties struct {
	AdjustmentType              interface{}                                          `yaml:"AdjustmentType,omitempty"`
	AutoScalingGroupName        interface{}                                          `yaml:"AutoScalingGroupName"`
	Cooldown                    interface{}                                          `yaml:"Cooldown,omitempty"`
	EstimatedInstanceWarmup     interface{}                                          `yaml:"EstimatedInstanceWarmup,omitempty"`
	MetricAggregationType       interface{}                                          `yaml:"MetricAggregationType,omitempty"`
	MinAdjustmentMagnitude      interface{}                                          `yaml:"MinAdjustmentMagnitude,omitempty"`
	PolicyType                  interface{}                                          `yaml:"PolicyType,omitempty"`
	ScalingAdjustment           interface{}                                          `yaml:"ScalingAdjustment,omitempty"`
	TargetTrackingConfiguration *properties.ScalingPolicyTargetTrackingConfiguration `yaml:"TargetTrackingConfiguration,omitempty"`
	StepAdjustments             interface{}                                          `yaml:"StepAdjustments,omitempty"`
}

// NewAutoScalingScalingPolicy constructor creates a new AutoScalingScalingPolicy
func NewAutoScalingScalingPolicy(properties AutoScalingScalingPolicyProperties, deps ...interface{}) AutoScalingScalingPolicy {
	return AutoScalingScalingPolicy{
		Type:       "AWS::AutoScaling::ScalingPolicy",
		Properties: properties,
		DependsOn:  deps,
	}
}

// ParseAutoScalingScalingPolicy parses AutoScalingScalingPolicy
func ParseAutoScalingScalingPolicy(ctx map[string]interface{}, name string, data string) (cf types.TemplateObject, err error) {
	var resource AutoScalingScalingPolicy
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: AutoScalingScalingPolicy - ", err)
		}
		return
	}
	cf = types.TemplateObject{name: resource}
	return
}

// ParseAutoScalingScalingPolicy validator
func (resource AutoScalingScalingPolicy) Validate() []error {
	return resource.Properties.Validate()
}

// ParseAutoScalingScalingPolicyProperties validator
func (resource AutoScalingScalingPolicyProperties) Validate() []error {
	errs := []error{}
	if resource.AutoScalingGroupName == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'AutoScalingGroupName'"))
	}
	return errs
}
