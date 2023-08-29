// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package project


type ProjectDataSourceGit struct {
	// Indicates signal to be sent to any applications when their config files change.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/waypoint/0.1.0/docs/resources/project#file_change_signal Project#file_change_signal}
	FileChangeSignal *string `field:"optional" json:"fileChangeSignal" yaml:"fileChangeSignal"`
	// Path in git repository when waypoint.hcl file is stored in a sub-directory.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/waypoint/0.1.0/docs/resources/project#git_path Project#git_path}
	GitPath *string `field:"optional" json:"gitPath" yaml:"gitPath"`
	// Interval at which Waypoint should poll git repository for changes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/waypoint/0.1.0/docs/resources/project#git_poll_interval_seconds Project#git_poll_interval_seconds}
	GitPollIntervalSeconds *float64 `field:"optional" json:"gitPollIntervalSeconds" yaml:"gitPollIntervalSeconds"`
	// Git repository ref containing waypoint.hcl file.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/waypoint/0.1.0/docs/resources/project#git_ref Project#git_ref}
	GitRef *string `field:"optional" json:"gitRef" yaml:"gitRef"`
	// Url of git repository storing the waypoint.hcl file.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/waypoint/0.1.0/docs/resources/project#git_url Project#git_url}
	GitUrl *string `field:"optional" json:"gitUrl" yaml:"gitUrl"`
	// Whether Waypoint ignores changes outside path storing waypoint.hcl file.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/waypoint/0.1.0/docs/resources/project#ignore_changes_outside_path Project#ignore_changes_outside_path}
	IgnoreChangesOutsidePath interface{} `field:"optional" json:"ignoreChangesOutsidePath" yaml:"ignoreChangesOutsidePath"`
}

