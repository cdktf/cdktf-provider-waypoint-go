// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package authmethod

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AuthMethodConfig struct {
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
	// Client ID of OIDC provider.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/waypoint/0.1.0/docs/resources/auth_method#client_id AuthMethod#client_id}
	ClientId *string `field:"required" json:"clientId" yaml:"clientId"`
	// Client Secret of OIDC provider.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/waypoint/0.1.0/docs/resources/auth_method#client_secret AuthMethod#client_secret}
	ClientSecret *string `field:"required" json:"clientSecret" yaml:"clientSecret"`
	// Discovery URL for OIDC provider.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/waypoint/0.1.0/docs/resources/auth_method#discovery_url AuthMethod#discovery_url}
	DiscoveryUrl *string `field:"required" json:"discoveryUrl" yaml:"discoveryUrl"`
	// The name of the Auth Method.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/waypoint/0.1.0/docs/resources/auth_method#name AuthMethod#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/waypoint/0.1.0/docs/resources/auth_method#accessor_selector AuthMethod#accessor_selector}.
	AccessorSelector *string `field:"optional" json:"accessorSelector" yaml:"accessorSelector"`
	// Allowed URI for auth redirection.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/waypoint/0.1.0/docs/resources/auth_method#allowed_redirect_uris AuthMethod#allowed_redirect_uris}
	AllowedRedirectUris *[]*string `field:"optional" json:"allowedRedirectUris" yaml:"allowedRedirectUris"`
	// The optional audience claims required.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/waypoint/0.1.0/docs/resources/auth_method#auds AuthMethod#auds}
	Auds *[]*string `field:"optional" json:"auds" yaml:"auds"`
	// Mapping of a claim to a variable value for the access selector.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/waypoint/0.1.0/docs/resources/auth_method#claim_mappings AuthMethod#claim_mappings}
	ClaimMappings *map[string]*string `field:"optional" json:"claimMappings" yaml:"claimMappings"`
	// Description of auth method.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/waypoint/0.1.0/docs/resources/auth_method#description AuthMethod#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Optional CA certificate chain to validate the discovery URL. Multiple CA certificates can be specified to support easier rotation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/waypoint/0.1.0/docs/resources/auth_method#discovery_ca_pem AuthMethod#discovery_ca_pem}
	DiscoveryCaPem *[]*string `field:"optional" json:"discoveryCaPem" yaml:"discoveryCaPem"`
	// The display name of the Auth Method.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/waypoint/0.1.0/docs/resources/auth_method#display_name AuthMethod#display_name}
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// Same as claim_mappings but for list values.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/waypoint/0.1.0/docs/resources/auth_method#list_claim_mappings AuthMethod#list_claim_mappings}
	ListClaimMappings *map[string]*string `field:"optional" json:"listClaimMappings" yaml:"listClaimMappings"`
	// The optional claims scope requested.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/waypoint/0.1.0/docs/resources/auth_method#scopes AuthMethod#scopes}
	Scopes *[]*string `field:"optional" json:"scopes" yaml:"scopes"`
	// The signing algorithms supported by the OIDC connect server.
	//
	// If this isn't specified, this will default to RS256 since that should be supported according to the RFC. The string values here should be valid OIDC signing algorithms
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/waypoint/0.1.0/docs/resources/auth_method#signing_algs AuthMethod#signing_algs}
	SigningAlgs *[]*string `field:"optional" json:"signingAlgs" yaml:"signingAlgs"`
}

