// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package authmethod

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-waypoint-go/waypoint/v2/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-waypoint-go/waypoint/v2/authmethod/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/waypoint/0.1.0/docs/resources/auth_method waypoint_auth_method}.
type AuthMethod interface {
	cdktf.TerraformResource
	AccessorSelector() *string
	SetAccessorSelector(val *string)
	AccessorSelectorInput() *string
	AllowedRedirectUris() *[]*string
	SetAllowedRedirectUris(val *[]*string)
	AllowedRedirectUrisInput() *[]*string
	Auds() *[]*string
	SetAuds(val *[]*string)
	AudsInput() *[]*string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	ClaimMappings() *map[string]*string
	SetClaimMappings(val *map[string]*string)
	ClaimMappingsInput() *map[string]*string
	ClientId() *string
	SetClientId(val *string)
	ClientIdInput() *string
	ClientSecret() *string
	SetClientSecret(val *string)
	ClientSecretInput() *string
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	DiscoveryCaPem() *[]*string
	SetDiscoveryCaPem(val *[]*string)
	DiscoveryCaPemInput() *[]*string
	DiscoveryUrl() *string
	SetDiscoveryUrl(val *string)
	DiscoveryUrlInput() *string
	DisplayName() *string
	SetDisplayName(val *string)
	DisplayNameInput() *string
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	ListClaimMappings() *map[string]*string
	SetListClaimMappings(val *map[string]*string)
	ListClaimMappingsInput() *map[string]*string
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	// Experimental.
	RawOverrides() interface{}
	Scopes() *[]*string
	SetScopes(val *[]*string)
	ScopesInput() *[]*string
	SigningAlgs() *[]*string
	SetSigningAlgs(val *[]*string)
	SigningAlgsInput() *[]*string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	// Adds a user defined moveTarget string to this resource to be later used in .moveTo(moveTarget) to resolve the location of the move.
	// Experimental.
	AddMoveTarget(moveTarget *string)
	// Experimental.
	AddOverride(path *string, value interface{})
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	// Experimental.
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	// Experimental.
	GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool
	// Experimental.
	GetListAttribute(terraformAttribute *string) *[]*string
	// Experimental.
	GetNumberAttribute(terraformAttribute *string) *float64
	// Experimental.
	GetNumberListAttribute(terraformAttribute *string) *[]*float64
	// Experimental.
	GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64
	// Experimental.
	GetStringAttribute(terraformAttribute *string) *string
	// Experimental.
	GetStringMapAttribute(terraformAttribute *string) *map[string]*string
	// Experimental.
	ImportFrom(id *string, provider cdktf.TerraformProvider)
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	// Moves this resource to the target resource given by moveTarget.
	// Experimental.
	MoveTo(moveTarget *string, index interface{})
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	ResetAccessorSelector()
	ResetAllowedRedirectUris()
	ResetAuds()
	ResetClaimMappings()
	ResetDescription()
	ResetDiscoveryCaPem()
	ResetDisplayName()
	ResetListClaimMappings()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetScopes()
	ResetSigningAlgs()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for AuthMethod
type jsiiProxy_AuthMethod struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_AuthMethod) AccessorSelector() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accessorSelector",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuthMethod) AccessorSelectorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accessorSelectorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuthMethod) AllowedRedirectUris() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedRedirectUris",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuthMethod) AllowedRedirectUrisInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedRedirectUrisInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuthMethod) Auds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"auds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuthMethod) AudsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"audsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuthMethod) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuthMethod) ClaimMappings() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"claimMappings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuthMethod) ClaimMappingsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"claimMappingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuthMethod) ClientId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuthMethod) ClientIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuthMethod) ClientSecret() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientSecret",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuthMethod) ClientSecretInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientSecretInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuthMethod) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuthMethod) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuthMethod) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuthMethod) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuthMethod) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuthMethod) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuthMethod) DiscoveryCaPem() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"discoveryCaPem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuthMethod) DiscoveryCaPemInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"discoveryCaPemInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuthMethod) DiscoveryUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"discoveryUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuthMethod) DiscoveryUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"discoveryUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuthMethod) DisplayName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuthMethod) DisplayNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuthMethod) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuthMethod) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuthMethod) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuthMethod) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuthMethod) ListClaimMappings() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"listClaimMappings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuthMethod) ListClaimMappingsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"listClaimMappingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuthMethod) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuthMethod) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuthMethod) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuthMethod) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuthMethod) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuthMethod) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuthMethod) Scopes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"scopes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuthMethod) ScopesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"scopesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuthMethod) SigningAlgs() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"signingAlgs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuthMethod) SigningAlgsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"signingAlgsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuthMethod) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuthMethod) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuthMethod) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/waypoint/0.1.0/docs/resources/auth_method waypoint_auth_method} Resource.
func NewAuthMethod(scope constructs.Construct, id *string, config *AuthMethodConfig) AuthMethod {
	_init_.Initialize()

	if err := validateNewAuthMethodParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_AuthMethod{}

	_jsii_.Create(
		"@cdktf/provider-waypoint.authMethod.AuthMethod",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/waypoint/0.1.0/docs/resources/auth_method waypoint_auth_method} Resource.
func NewAuthMethod_Override(a AuthMethod, scope constructs.Construct, id *string, config *AuthMethodConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-waypoint.authMethod.AuthMethod",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_AuthMethod)SetAccessorSelector(val *string) {
	if err := j.validateSetAccessorSelectorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accessorSelector",
		val,
	)
}

func (j *jsiiProxy_AuthMethod)SetAllowedRedirectUris(val *[]*string) {
	if err := j.validateSetAllowedRedirectUrisParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowedRedirectUris",
		val,
	)
}

func (j *jsiiProxy_AuthMethod)SetAuds(val *[]*string) {
	if err := j.validateSetAudsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"auds",
		val,
	)
}

func (j *jsiiProxy_AuthMethod)SetClaimMappings(val *map[string]*string) {
	if err := j.validateSetClaimMappingsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"claimMappings",
		val,
	)
}

func (j *jsiiProxy_AuthMethod)SetClientId(val *string) {
	if err := j.validateSetClientIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientId",
		val,
	)
}

func (j *jsiiProxy_AuthMethod)SetClientSecret(val *string) {
	if err := j.validateSetClientSecretParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientSecret",
		val,
	)
}

func (j *jsiiProxy_AuthMethod)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_AuthMethod)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_AuthMethod)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_AuthMethod)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_AuthMethod)SetDiscoveryCaPem(val *[]*string) {
	if err := j.validateSetDiscoveryCaPemParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"discoveryCaPem",
		val,
	)
}

func (j *jsiiProxy_AuthMethod)SetDiscoveryUrl(val *string) {
	if err := j.validateSetDiscoveryUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"discoveryUrl",
		val,
	)
}

func (j *jsiiProxy_AuthMethod)SetDisplayName(val *string) {
	if err := j.validateSetDisplayNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"displayName",
		val,
	)
}

func (j *jsiiProxy_AuthMethod)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_AuthMethod)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_AuthMethod)SetListClaimMappings(val *map[string]*string) {
	if err := j.validateSetListClaimMappingsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"listClaimMappings",
		val,
	)
}

func (j *jsiiProxy_AuthMethod)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_AuthMethod)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_AuthMethod)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_AuthMethod)SetScopes(val *[]*string) {
	if err := j.validateSetScopesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scopes",
		val,
	)
}

func (j *jsiiProxy_AuthMethod)SetSigningAlgs(val *[]*string) {
	if err := j.validateSetSigningAlgsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"signingAlgs",
		val,
	)
}

// Generates CDKTF code for importing a AuthMethod resource upon running "cdktf plan <stack-name>".
func AuthMethod_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateAuthMethod_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-waypoint.authMethod.AuthMethod",
		"generateConfigForImport",
		[]interface{}{scope, importToId, importFromId, provider},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func AuthMethod_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAuthMethod_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-waypoint.authMethod.AuthMethod",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AuthMethod_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAuthMethod_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-waypoint.authMethod.AuthMethod",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AuthMethod_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAuthMethod_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-waypoint.authMethod.AuthMethod",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func AuthMethod_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-waypoint.authMethod.AuthMethod",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (a *jsiiProxy_AuthMethod) AddMoveTarget(moveTarget *string) {
	if err := a.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (a *jsiiProxy_AuthMethod) AddOverride(path *string, value interface{}) {
	if err := a.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (a *jsiiProxy_AuthMethod) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := a.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AuthMethod) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := a.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AuthMethod) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := a.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AuthMethod) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := a.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AuthMethod) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := a.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AuthMethod) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := a.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AuthMethod) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := a.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AuthMethod) GetStringAttribute(terraformAttribute *string) *string {
	if err := a.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AuthMethod) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := a.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AuthMethod) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := a.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (a *jsiiProxy_AuthMethod) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := a.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AuthMethod) MoveTo(moveTarget *string, index interface{}) {
	if err := a.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (a *jsiiProxy_AuthMethod) OverrideLogicalId(newLogicalId *string) {
	if err := a.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_AuthMethod) ResetAccessorSelector() {
	_jsii_.InvokeVoid(
		a,
		"resetAccessorSelector",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AuthMethod) ResetAllowedRedirectUris() {
	_jsii_.InvokeVoid(
		a,
		"resetAllowedRedirectUris",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AuthMethod) ResetAuds() {
	_jsii_.InvokeVoid(
		a,
		"resetAuds",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AuthMethod) ResetClaimMappings() {
	_jsii_.InvokeVoid(
		a,
		"resetClaimMappings",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AuthMethod) ResetDescription() {
	_jsii_.InvokeVoid(
		a,
		"resetDescription",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AuthMethod) ResetDiscoveryCaPem() {
	_jsii_.InvokeVoid(
		a,
		"resetDiscoveryCaPem",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AuthMethod) ResetDisplayName() {
	_jsii_.InvokeVoid(
		a,
		"resetDisplayName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AuthMethod) ResetListClaimMappings() {
	_jsii_.InvokeVoid(
		a,
		"resetListClaimMappings",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AuthMethod) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AuthMethod) ResetScopes() {
	_jsii_.InvokeVoid(
		a,
		"resetScopes",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AuthMethod) ResetSigningAlgs() {
	_jsii_.InvokeVoid(
		a,
		"resetSigningAlgs",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AuthMethod) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AuthMethod) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AuthMethod) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AuthMethod) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

