package properties

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import "fmt"

// ScalingPlanTargetTrackingConfiguration Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-autoscalingplans-scalingplan-targettrackingconfiguration.html
type ScalingPlanTargetTrackingConfiguration struct {
	DisableScaleIn                       interface{}                                      `yaml:"DisableScaleIn,omitempty"`
	EstimatedInstanceWarmup              interface{}                                      `yaml:"EstimatedInstanceWarmup,omitempty"`
	ScaleInCooldown                      interface{}                                      `yaml:"ScaleInCooldown,omitempty"`
	ScaleOutCooldown                     interface{}                                      `yaml:"ScaleOutCooldown,omitempty"`
	TargetValue                          interface{}                                      `yaml:"TargetValue"`
	PredefinedScalingMetricSpecification *ScalingPlanPredefinedScalingMetricSpecification `yaml:"PredefinedScalingMetricSpecification,omitempty"`
	CustomizedScalingMetricSpecification *ScalingPlanCustomizedScalingMetricSpecification `yaml:"CustomizedScalingMetricSpecification,omitempty"`
}

// ScalingPlanTargetTrackingConfiguration validation
func (resource ScalingPlanTargetTrackingConfiguration) Validate() []error {
	errs := []error{}

	if resource.TargetValue == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'TargetValue'"))
	}
	return errs
}
