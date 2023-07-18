package project

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ProjectConfig struct {
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
	// Configuration of Git repository where waypoint.hcl file is stored.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/waypoint/0.1.0/docs/resources/project#data_source_git Project#data_source_git}
	DataSourceGit *ProjectDataSourceGit `field:"required" json:"dataSourceGit" yaml:"dataSourceGit"`
	// The name of the Waypoint project.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/waypoint/0.1.0/docs/resources/project#project_name Project#project_name}
	ProjectName *string `field:"required" json:"projectName" yaml:"projectName"`
	// Application status poll interval in seconds.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/waypoint/0.1.0/docs/resources/project#app_status_poll_seconds Project#app_status_poll_seconds}
	AppStatusPollSeconds *float64 `field:"optional" json:"appStatusPollSeconds" yaml:"appStatusPollSeconds"`
	// Basic authentication details for Git consisting of `username` and `password`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/waypoint/0.1.0/docs/resources/project#git_auth_basic Project#git_auth_basic}
	GitAuthBasic *ProjectGitAuthBasic `field:"optional" json:"gitAuthBasic" yaml:"gitAuthBasic"`
	// SSH authentication details for Git.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/waypoint/0.1.0/docs/resources/project#git_auth_ssh Project#git_auth_ssh}
	GitAuthSsh *ProjectGitAuthSsh `field:"optional" json:"gitAuthSsh" yaml:"gitAuthSsh"`
	// List of variables in Key/value pairs associated with the Waypoint Project.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/waypoint/0.1.0/docs/resources/project#project_variables Project#project_variables}
	ProjectVariables interface{} `field:"optional" json:"projectVariables" yaml:"projectVariables"`
	// Enable remote runners for project.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/waypoint/0.1.0/docs/resources/project#remote_runners_enabled Project#remote_runners_enabled}
	RemoteRunnersEnabled interface{} `field:"optional" json:"remoteRunnersEnabled" yaml:"remoteRunnersEnabled"`
}

