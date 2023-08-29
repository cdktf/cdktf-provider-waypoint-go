// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package runnerprofile

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-waypoint-go/waypoint/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-waypoint-go/waypoint/runnerprofile/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/waypoint/0.1.0/docs/resources/runner_profile waypoint_runner_profile}.
type RunnerProfile interface {
	cdktf.TerraformResource
	// Experimental.
	CdktfStack() cdktf.TerraformStack
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
	Default() interface{}
	SetDefault(val interface{})
	DefaultInput() interface{}
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	EnvironmentVariables() *map[string]*string
	SetEnvironmentVariables(val *map[string]*string)
	EnvironmentVariablesInput() *map[string]*string
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Id() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	OciUrl() *string
	SetOciUrl(val *string)
	OciUrlInput() *string
	PluginConfig() *string
	SetPluginConfig(val *string)
	PluginConfigFormat() *string
	SetPluginConfigFormat(val *string)
	PluginConfigFormatInput() *string
	PluginConfigInput() *string
	PluginType() *string
	SetPluginType(val *string)
	PluginTypeInput() *string
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
	TargetRunnerId() *string
	SetTargetRunnerId(val *string)
	TargetRunnerIdInput() *string
	TargetRunnerLabels() *map[string]*string
	SetTargetRunnerLabels(val *map[string]*string)
	TargetRunnerLabelsInput() *map[string]*string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
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
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	ResetDefault()
	ResetEnvironmentVariables()
	ResetOciUrl()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPluginConfig()
	ResetPluginConfigFormat()
	ResetTargetRunnerId()
	ResetTargetRunnerLabels()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for RunnerProfile
type jsiiProxy_RunnerProfile struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_RunnerProfile) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RunnerProfile) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RunnerProfile) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RunnerProfile) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RunnerProfile) Default() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"default",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RunnerProfile) DefaultInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"defaultInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RunnerProfile) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RunnerProfile) EnvironmentVariables() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"environmentVariables",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RunnerProfile) EnvironmentVariablesInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"environmentVariablesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RunnerProfile) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RunnerProfile) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RunnerProfile) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RunnerProfile) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RunnerProfile) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RunnerProfile) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RunnerProfile) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RunnerProfile) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RunnerProfile) OciUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ociUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RunnerProfile) OciUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ociUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RunnerProfile) PluginConfig() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pluginConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RunnerProfile) PluginConfigFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pluginConfigFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RunnerProfile) PluginConfigFormatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pluginConfigFormatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RunnerProfile) PluginConfigInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pluginConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RunnerProfile) PluginType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pluginType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RunnerProfile) PluginTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pluginTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RunnerProfile) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RunnerProfile) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RunnerProfile) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RunnerProfile) TargetRunnerId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetRunnerId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RunnerProfile) TargetRunnerIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetRunnerIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RunnerProfile) TargetRunnerLabels() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"targetRunnerLabels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RunnerProfile) TargetRunnerLabelsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"targetRunnerLabelsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RunnerProfile) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RunnerProfile) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RunnerProfile) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/waypoint/0.1.0/docs/resources/runner_profile waypoint_runner_profile} Resource.
func NewRunnerProfile(scope constructs.Construct, id *string, config *RunnerProfileConfig) RunnerProfile {
	_init_.Initialize()

	if err := validateNewRunnerProfileParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_RunnerProfile{}

	_jsii_.Create(
		"@cdktf/provider-waypoint.runnerProfile.RunnerProfile",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/waypoint/0.1.0/docs/resources/runner_profile waypoint_runner_profile} Resource.
func NewRunnerProfile_Override(r RunnerProfile, scope constructs.Construct, id *string, config *RunnerProfileConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-waypoint.runnerProfile.RunnerProfile",
		[]interface{}{scope, id, config},
		r,
	)
}

func (j *jsiiProxy_RunnerProfile)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_RunnerProfile)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_RunnerProfile)SetDefault(val interface{}) {
	if err := j.validateSetDefaultParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"default",
		val,
	)
}

func (j *jsiiProxy_RunnerProfile)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_RunnerProfile)SetEnvironmentVariables(val *map[string]*string) {
	if err := j.validateSetEnvironmentVariablesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"environmentVariables",
		val,
	)
}

func (j *jsiiProxy_RunnerProfile)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_RunnerProfile)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_RunnerProfile)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_RunnerProfile)SetOciUrl(val *string) {
	if err := j.validateSetOciUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ociUrl",
		val,
	)
}

func (j *jsiiProxy_RunnerProfile)SetPluginConfig(val *string) {
	if err := j.validateSetPluginConfigParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pluginConfig",
		val,
	)
}

func (j *jsiiProxy_RunnerProfile)SetPluginConfigFormat(val *string) {
	if err := j.validateSetPluginConfigFormatParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pluginConfigFormat",
		val,
	)
}

func (j *jsiiProxy_RunnerProfile)SetPluginType(val *string) {
	if err := j.validateSetPluginTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pluginType",
		val,
	)
}

func (j *jsiiProxy_RunnerProfile)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_RunnerProfile)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_RunnerProfile)SetTargetRunnerId(val *string) {
	if err := j.validateSetTargetRunnerIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targetRunnerId",
		val,
	)
}

func (j *jsiiProxy_RunnerProfile)SetTargetRunnerLabels(val *map[string]*string) {
	if err := j.validateSetTargetRunnerLabelsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targetRunnerLabels",
		val,
	)
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
func RunnerProfile_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateRunnerProfile_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-waypoint.runnerProfile.RunnerProfile",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func RunnerProfile_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateRunnerProfile_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-waypoint.runnerProfile.RunnerProfile",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func RunnerProfile_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateRunnerProfile_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-waypoint.runnerProfile.RunnerProfile",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func RunnerProfile_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-waypoint.runnerProfile.RunnerProfile",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (r *jsiiProxy_RunnerProfile) AddOverride(path *string, value interface{}) {
	if err := r.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (r *jsiiProxy_RunnerProfile) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := r.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		r,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RunnerProfile) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := r.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		r,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RunnerProfile) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := r.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		r,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RunnerProfile) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := r.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		r,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RunnerProfile) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := r.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		r,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RunnerProfile) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := r.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		r,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RunnerProfile) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := r.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		r,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RunnerProfile) GetStringAttribute(terraformAttribute *string) *string {
	if err := r.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		r,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RunnerProfile) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := r.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		r,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RunnerProfile) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := r.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		r,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RunnerProfile) OverrideLogicalId(newLogicalId *string) {
	if err := r.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (r *jsiiProxy_RunnerProfile) ResetDefault() {
	_jsii_.InvokeVoid(
		r,
		"resetDefault",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RunnerProfile) ResetEnvironmentVariables() {
	_jsii_.InvokeVoid(
		r,
		"resetEnvironmentVariables",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RunnerProfile) ResetOciUrl() {
	_jsii_.InvokeVoid(
		r,
		"resetOciUrl",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RunnerProfile) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		r,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RunnerProfile) ResetPluginConfig() {
	_jsii_.InvokeVoid(
		r,
		"resetPluginConfig",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RunnerProfile) ResetPluginConfigFormat() {
	_jsii_.InvokeVoid(
		r,
		"resetPluginConfigFormat",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RunnerProfile) ResetTargetRunnerId() {
	_jsii_.InvokeVoid(
		r,
		"resetTargetRunnerId",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RunnerProfile) ResetTargetRunnerLabels() {
	_jsii_.InvokeVoid(
		r,
		"resetTargetRunnerLabels",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RunnerProfile) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		r,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RunnerProfile) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		r,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RunnerProfile) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RunnerProfile) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		r,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

