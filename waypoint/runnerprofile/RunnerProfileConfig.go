// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package runnerprofile

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type RunnerProfileConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// The name of the runner profile.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/waypoint/0.1.0/docs/resources/runner_profile#name RunnerProfile#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Plugin type for runner i.e docker / kubernetes / aws-ecs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/waypoint/0.1.0/docs/resources/runner_profile#plugin_type RunnerProfile#plugin_type}
	PluginType *string `field:"required" json:"pluginType" yaml:"pluginType"`
	// Indicates if this runner profile is the default for any new projects. The default is false.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/waypoint/0.1.0/docs/resources/runner_profile#default RunnerProfile#default}
	Default interface{} `field:"optional" json:"default" yaml:"default"`
	// Any environment variables that should be exposed to the on demand runner.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/waypoint/0.1.0/docs/resources/runner_profile#environment_variables RunnerProfile#environment_variables}
	EnvironmentVariables *map[string]*string `field:"optional" json:"environmentVariables" yaml:"environmentVariables"`
	// oci_url is the OCI image that will be used to boot the on demand runner.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/waypoint/0.1.0/docs/resources/runner_profile#oci_url RunnerProfile#oci_url}
	OciUrl *string `field:"optional" json:"ociUrl" yaml:"ociUrl"`
	// Plugin config is the configuration for the plugin that is created.
	//
	// It is usually HCL and is decoded like the other plugins, and is plugin specific.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/waypoint/0.1.0/docs/resources/runner_profile#plugin_config RunnerProfile#plugin_config}
	PluginConfig *string `field:"optional" json:"pluginConfig" yaml:"pluginConfig"`
	// Config format specifies the format of plugin_config. Valid values are HCL or JSON. The default is HCL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/waypoint/0.1.0/docs/resources/runner_profile#plugin_config_format RunnerProfile#plugin_config_format}
	PluginConfigFormat *string `field:"optional" json:"pluginConfigFormat" yaml:"pluginConfigFormat"`
	// The ID of the target runner for this profile.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/waypoint/0.1.0/docs/resources/runner_profile#target_runner_id RunnerProfile#target_runner_id}
	TargetRunnerId *string `field:"optional" json:"targetRunnerId" yaml:"targetRunnerId"`
	// A map of labels on target runners.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/waypoint/0.1.0/docs/resources/runner_profile#target_runner_labels RunnerProfile#target_runner_labels}
	TargetRunnerLabels *map[string]*string `field:"optional" json:"targetRunnerLabels" yaml:"targetRunnerLabels"`
}

