// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package project


type ProjectGitAuthBasic struct {
	// Git password.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/waypoint/0.1.0/docs/resources/project#password Project#password}
	Password *string `field:"required" json:"password" yaml:"password"`
	// Git username.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/waypoint/0.1.0/docs/resources/project#username Project#username}
	Username *string `field:"required" json:"username" yaml:"username"`
}

