package project


type ProjectGitAuthSsh struct {
	// Private key to authenticate to Git.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/waypoint/0.1.0/docs/resources/project#ssh_private_key Project#ssh_private_key}
	SshPrivateKey *string `field:"required" json:"sshPrivateKey" yaml:"sshPrivateKey"`
	// Git user associated with private key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/waypoint/0.1.0/docs/resources/project#git_user Project#git_user}
	GitUser *string `field:"optional" json:"gitUser" yaml:"gitUser"`
	// Passphrase to use with private key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/waypoint/0.1.0/docs/resources/project#passphrase Project#passphrase}
	Passphrase *string `field:"optional" json:"passphrase" yaml:"passphrase"`
}

