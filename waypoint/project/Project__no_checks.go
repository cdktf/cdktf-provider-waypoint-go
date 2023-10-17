// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package project

// Building without runtime type checking enabled, so all the below just return nil

func (p *jsiiProxy_Project) validateAddMoveTargetParameters(moveTarget *string) error {
	return nil
}

func (p *jsiiProxy_Project) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (p *jsiiProxy_Project) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (p *jsiiProxy_Project) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (p *jsiiProxy_Project) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (p *jsiiProxy_Project) validateGetListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (p *jsiiProxy_Project) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (p *jsiiProxy_Project) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (p *jsiiProxy_Project) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (p *jsiiProxy_Project) validateGetStringAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (p *jsiiProxy_Project) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (p *jsiiProxy_Project) validateImportFromParameters(id *string) error {
	return nil
}

func (p *jsiiProxy_Project) validateInterpolationForAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (p *jsiiProxy_Project) validateMoveToParameters(moveTarget *string, index interface{}) error {
	return nil
}

func (p *jsiiProxy_Project) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func (p *jsiiProxy_Project) validatePutDataSourceGitParameters(value *ProjectDataSourceGit) error {
	return nil
}

func (p *jsiiProxy_Project) validatePutGitAuthBasicParameters(value *ProjectGitAuthBasic) error {
	return nil
}

func (p *jsiiProxy_Project) validatePutGitAuthSshParameters(value *ProjectGitAuthSsh) error {
	return nil
}

func (p *jsiiProxy_Project) validatePutProjectVariablesParameters(value interface{}) error {
	return nil
}

func validateProject_GenerateConfigForImportParameters(scope constructs.Construct, importToId *string, importFromId *string) error {
	return nil
}

func validateProject_IsConstructParameters(x interface{}) error {
	return nil
}

func validateProject_IsTerraformElementParameters(x interface{}) error {
	return nil
}

func validateProject_IsTerraformResourceParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_Project) validateSetAppStatusPollSecondsParameters(val *float64) error {
	return nil
}

func (j *jsiiProxy_Project) validateSetConnectionParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Project) validateSetCountParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Project) validateSetLifecycleParameters(val *cdktf.TerraformResourceLifecycle) error {
	return nil
}

func (j *jsiiProxy_Project) validateSetProjectNameParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Project) validateSetProvisionersParameters(val *[]interface{}) error {
	return nil
}

func (j *jsiiProxy_Project) validateSetRemoteRunnersEnabledParameters(val interface{}) error {
	return nil
}

func validateNewProjectParameters(scope constructs.Construct, id *string, config *ProjectConfig) error {
	return nil
}

