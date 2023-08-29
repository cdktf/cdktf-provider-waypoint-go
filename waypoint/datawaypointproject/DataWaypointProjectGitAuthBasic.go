// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datawaypointproject


type DataWaypointProjectGitAuthBasic struct {
	// Git password.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/waypoint/0.1.0/docs/data-sources/project#password DataWaypointProject#password}
	Password *string `field:"required" json:"password" yaml:"password"`
	// Git username.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/waypoint/0.1.0/docs/data-sources/project#username DataWaypointProject#username}
	Username *string `field:"required" json:"username" yaml:"username"`
}

