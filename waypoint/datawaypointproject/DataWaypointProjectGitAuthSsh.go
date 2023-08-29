// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datawaypointproject


type DataWaypointProjectGitAuthSsh struct {
	// Private key to authenticate to Git.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/waypoint/0.1.0/docs/data-sources/project#ssh_private_key DataWaypointProject#ssh_private_key}
	SshPrivateKey *string `field:"required" json:"sshPrivateKey" yaml:"sshPrivateKey"`
}

