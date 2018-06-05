package outputs

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import (
	"github.com/KablamoOSS/kombustion/types"
	yaml "github.com/KablamoOSS/yaml"
)

func ParseEC2Subnet(ctx map[string]interface{}, name string, data string) (cf types.TemplateObject, err error) {

	var resource, output types.TemplateObject
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}

	cf = types.TemplateObject{
		name: types.TemplateObject{
			"Description": name + " Object",
			"Value": map[string]interface{}{
				"Ref": name,
			},
			"Export": map[string]interface{}{
				"Name": map[string]interface{}{
					"Fn::Sub": "${AWS::StackName}-EC2Subnet-" + name,
				},
			},
		},
	}

	output = types.TemplateObject{
		"Description": name + " Object",
		"Value": map[string]interface{}{
			"Fn::GetAtt": []string{name, "AvailabilityZone"},
		},
		"Export": map[string]interface{}{
			"Name": map[string]interface{}{
				"Fn::Sub": "${AWS::StackName}-EC2Subnet-" + name + "-AvailabilityZone",
			},
		},
	}
	if condition, ok := resource["Condition"]; ok {
		output["Condition"] = condition
	}
	cf[name+"AvailabilityZone"] = output

	output = types.TemplateObject{
		"Description": name + " Object",
		"Value": map[string]interface{}{
			"Fn::GetAtt": []string{name, "Ipv6CidrBlocks"},
		},
		"Export": map[string]interface{}{
			"Name": map[string]interface{}{
				"Fn::Sub": "${AWS::StackName}-EC2Subnet-" + name + "-Ipv6CidrBlocks",
			},
		},
	}
	if condition, ok := resource["Condition"]; ok {
		output["Condition"] = condition
	}
	cf[name+"Ipv6CidrBlocks"] = output

	output = types.TemplateObject{
		"Description": name + " Object",
		"Value": map[string]interface{}{
			"Fn::GetAtt": []string{name, "NetworkAclAssociationId"},
		},
		"Export": map[string]interface{}{
			"Name": map[string]interface{}{
				"Fn::Sub": "${AWS::StackName}-EC2Subnet-" + name + "-NetworkAclAssociationId",
			},
		},
	}
	if condition, ok := resource["Condition"]; ok {
		output["Condition"] = condition
	}
	cf[name+"NetworkAclAssociationId"] = output

	output = types.TemplateObject{
		"Description": name + " Object",
		"Value": map[string]interface{}{
			"Fn::GetAtt": []string{name, "VpcId"},
		},
		"Export": map[string]interface{}{
			"Name": map[string]interface{}{
				"Fn::Sub": "${AWS::StackName}-EC2Subnet-" + name + "-VpcId",
			},
		},
	}
	if condition, ok := resource["Condition"]; ok {
		output["Condition"] = condition
	}
	cf[name+"VpcId"] = output

	return
}