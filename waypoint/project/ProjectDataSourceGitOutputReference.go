// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package project

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-waypoint-go/waypoint/v2/jsii"

	"github.com/cdktf/cdktf-provider-waypoint-go/waypoint/v2/project/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ProjectDataSourceGitOutputReference interface {
	cdktf.ComplexObject
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	FileChangeSignal() *string
	SetFileChangeSignal(val *string)
	FileChangeSignalInput() *string
	// Experimental.
	Fqn() *string
	GitPath() *string
	SetGitPath(val *string)
	GitPathInput() *string
	GitPollIntervalSeconds() *float64
	SetGitPollIntervalSeconds(val *float64)
	GitPollIntervalSecondsInput() *float64
	GitRef() *string
	SetGitRef(val *string)
	GitRefInput() *string
	GitUrl() *string
	SetGitUrl(val *string)
	GitUrlInput() *string
	IgnoreChangesOutsidePath() interface{}
	SetIgnoreChangesOutsidePath(val interface{})
	IgnoreChangesOutsidePathInput() interface{}
	InternalValue() interface{}
	SetInternalValue(val interface{})
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	// Experimental.
	ComputeFqn() *string
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
	InterpolationAsList() cdktf.IResolvable
	// Experimental.
	InterpolationForAttribute(property *string) cdktf.IResolvable
	ResetFileChangeSignal()
	ResetGitPath()
	ResetGitPollIntervalSeconds()
	ResetGitRef()
	ResetGitUrl()
	ResetIgnoreChangesOutsidePath()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ProjectDataSourceGitOutputReference
type jsiiProxy_ProjectDataSourceGitOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ProjectDataSourceGitOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectDataSourceGitOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectDataSourceGitOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectDataSourceGitOutputReference) FileChangeSignal() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fileChangeSignal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectDataSourceGitOutputReference) FileChangeSignalInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fileChangeSignalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectDataSourceGitOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectDataSourceGitOutputReference) GitPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gitPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectDataSourceGitOutputReference) GitPathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gitPathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectDataSourceGitOutputReference) GitPollIntervalSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"gitPollIntervalSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectDataSourceGitOutputReference) GitPollIntervalSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"gitPollIntervalSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectDataSourceGitOutputReference) GitRef() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gitRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectDataSourceGitOutputReference) GitRefInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gitRefInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectDataSourceGitOutputReference) GitUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gitUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectDataSourceGitOutputReference) GitUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gitUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectDataSourceGitOutputReference) IgnoreChangesOutsidePath() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ignoreChangesOutsidePath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectDataSourceGitOutputReference) IgnoreChangesOutsidePathInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ignoreChangesOutsidePathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectDataSourceGitOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectDataSourceGitOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectDataSourceGitOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewProjectDataSourceGitOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ProjectDataSourceGitOutputReference {
	_init_.Initialize()

	if err := validateNewProjectDataSourceGitOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ProjectDataSourceGitOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-waypoint.project.ProjectDataSourceGitOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewProjectDataSourceGitOutputReference_Override(p ProjectDataSourceGitOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-waypoint.project.ProjectDataSourceGitOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		p,
	)
}

func (j *jsiiProxy_ProjectDataSourceGitOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ProjectDataSourceGitOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ProjectDataSourceGitOutputReference)SetFileChangeSignal(val *string) {
	if err := j.validateSetFileChangeSignalParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fileChangeSignal",
		val,
	)
}

func (j *jsiiProxy_ProjectDataSourceGitOutputReference)SetGitPath(val *string) {
	if err := j.validateSetGitPathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"gitPath",
		val,
	)
}

func (j *jsiiProxy_ProjectDataSourceGitOutputReference)SetGitPollIntervalSeconds(val *float64) {
	if err := j.validateSetGitPollIntervalSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"gitPollIntervalSeconds",
		val,
	)
}

func (j *jsiiProxy_ProjectDataSourceGitOutputReference)SetGitRef(val *string) {
	if err := j.validateSetGitRefParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"gitRef",
		val,
	)
}

func (j *jsiiProxy_ProjectDataSourceGitOutputReference)SetGitUrl(val *string) {
	if err := j.validateSetGitUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"gitUrl",
		val,
	)
}

func (j *jsiiProxy_ProjectDataSourceGitOutputReference)SetIgnoreChangesOutsidePath(val interface{}) {
	if err := j.validateSetIgnoreChangesOutsidePathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ignoreChangesOutsidePath",
		val,
	)
}

func (j *jsiiProxy_ProjectDataSourceGitOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ProjectDataSourceGitOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ProjectDataSourceGitOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (p *jsiiProxy_ProjectDataSourceGitOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_ProjectDataSourceGitOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := p.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		p,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_ProjectDataSourceGitOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := p.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_ProjectDataSourceGitOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := p.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		p,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_ProjectDataSourceGitOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := p.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		p,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_ProjectDataSourceGitOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := p.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		p,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_ProjectDataSourceGitOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := p.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		p,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_ProjectDataSourceGitOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := p.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		p,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_ProjectDataSourceGitOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := p.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		p,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_ProjectDataSourceGitOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := p.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		p,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_ProjectDataSourceGitOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_ProjectDataSourceGitOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := p.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_ProjectDataSourceGitOutputReference) ResetFileChangeSignal() {
	_jsii_.InvokeVoid(
		p,
		"resetFileChangeSignal",
		nil, // no parameters
	)
}

func (p *jsiiProxy_ProjectDataSourceGitOutputReference) ResetGitPath() {
	_jsii_.InvokeVoid(
		p,
		"resetGitPath",
		nil, // no parameters
	)
}

func (p *jsiiProxy_ProjectDataSourceGitOutputReference) ResetGitPollIntervalSeconds() {
	_jsii_.InvokeVoid(
		p,
		"resetGitPollIntervalSeconds",
		nil, // no parameters
	)
}

func (p *jsiiProxy_ProjectDataSourceGitOutputReference) ResetGitRef() {
	_jsii_.InvokeVoid(
		p,
		"resetGitRef",
		nil, // no parameters
	)
}

func (p *jsiiProxy_ProjectDataSourceGitOutputReference) ResetGitUrl() {
	_jsii_.InvokeVoid(
		p,
		"resetGitUrl",
		nil, // no parameters
	)
}

func (p *jsiiProxy_ProjectDataSourceGitOutputReference) ResetIgnoreChangesOutsidePath() {
	_jsii_.InvokeVoid(
		p,
		"resetIgnoreChangesOutsidePath",
		nil, // no parameters
	)
}

func (p *jsiiProxy_ProjectDataSourceGitOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := p.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		p,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_ProjectDataSourceGitOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

