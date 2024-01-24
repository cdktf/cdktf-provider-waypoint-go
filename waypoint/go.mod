// Deprecated: HashiCorp is no longer publishing new versions of the prebuilt provider for waypoint.
// Previously-published versions of this prebuilt provider will still continue to be available as installable Go modules,
// but these will not be compatible with newer versions of CDK for Terraform and are not eligible for commercial support.
// You can continue to use the waypoint provider in your CDK for Terraform projects with newer versions of CDKTF,
// but you will need to generate the bindings locally. See https://cdk.tf/imports for details.
module github.com/cdktf/cdktf-provider-waypoint-go/waypoint/v3

go 1.18

require (
	github.com/aws/jsii-runtime-go v1.94.0
	github.com/hashicorp/terraform-cdk-go/cdktf v0.20.1
	github.com/aws/constructs-go/constructs/v10 v10.3.0
)
