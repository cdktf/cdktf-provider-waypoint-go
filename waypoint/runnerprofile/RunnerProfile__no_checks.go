// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package runnerprofile

// Building without runtime type checking enabled, so all the below just return nil

func (r *jsiiProxy_RunnerProfile) validateAddMoveTargetParameters(moveTarget *string) error {
	return nil
}

func (r *jsiiProxy_RunnerProfile) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (r *jsiiProxy_RunnerProfile) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (r *jsiiProxy_RunnerProfile) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (r *jsiiProxy_RunnerProfile) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (r *jsiiProxy_RunnerProfile) validateGetListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (r *jsiiProxy_RunnerProfile) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (r *jsiiProxy_RunnerProfile) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (r *jsiiProxy_RunnerProfile) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (r *jsiiProxy_RunnerProfile) validateGetStringAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (r *jsiiProxy_RunnerProfile) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (r *jsiiProxy_RunnerProfile) validateImportFromParameters(id *string) error {
	return nil
}

func (r *jsiiProxy_RunnerProfile) validateInterpolationForAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (r *jsiiProxy_RunnerProfile) validateMoveFromIdParameters(id *string) error {
	return nil
}

func (r *jsiiProxy_RunnerProfile) validateMoveToParameters(moveTarget *string, index interface{}) error {
	return nil
}

func (r *jsiiProxy_RunnerProfile) validateMoveToIdParameters(id *string) error {
	return nil
}

func (r *jsiiProxy_RunnerProfile) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func validateRunnerProfile_GenerateConfigForImportParameters(scope constructs.Construct, importToId *string, importFromId *string) error {
	return nil
}

func validateRunnerProfile_IsConstructParameters(x interface{}) error {
	return nil
}

func validateRunnerProfile_IsTerraformElementParameters(x interface{}) error {
	return nil
}

func validateRunnerProfile_IsTerraformResourceParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_RunnerProfile) validateSetConnectionParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_RunnerProfile) validateSetCountParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_RunnerProfile) validateSetDefaultParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_RunnerProfile) validateSetEnvironmentVariablesParameters(val *map[string]*string) error {
	return nil
}

func (j *jsiiProxy_RunnerProfile) validateSetLifecycleParameters(val *cdktf.TerraformResourceLifecycle) error {
	return nil
}

func (j *jsiiProxy_RunnerProfile) validateSetNameParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_RunnerProfile) validateSetOciUrlParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_RunnerProfile) validateSetPluginConfigParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_RunnerProfile) validateSetPluginConfigFormatParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_RunnerProfile) validateSetPluginTypeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_RunnerProfile) validateSetProvisionersParameters(val *[]interface{}) error {
	return nil
}

func (j *jsiiProxy_RunnerProfile) validateSetTargetRunnerIdParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_RunnerProfile) validateSetTargetRunnerLabelsParameters(val *map[string]*string) error {
	return nil
}

func validateNewRunnerProfileParameters(scope constructs.Construct, id *string, config *RunnerProfileConfig) error {
	return nil
}

