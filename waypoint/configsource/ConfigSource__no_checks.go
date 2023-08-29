// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package configsource

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_ConfigSource) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (c *jsiiProxy_ConfigSource) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (c *jsiiProxy_ConfigSource) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (c *jsiiProxy_ConfigSource) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (c *jsiiProxy_ConfigSource) validateGetListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (c *jsiiProxy_ConfigSource) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (c *jsiiProxy_ConfigSource) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (c *jsiiProxy_ConfigSource) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (c *jsiiProxy_ConfigSource) validateGetStringAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (c *jsiiProxy_ConfigSource) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (c *jsiiProxy_ConfigSource) validateInterpolationForAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (c *jsiiProxy_ConfigSource) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func validateConfigSource_IsConstructParameters(x interface{}) error {
	return nil
}

func validateConfigSource_IsTerraformElementParameters(x interface{}) error {
	return nil
}

func validateConfigSource_IsTerraformResourceParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_ConfigSource) validateSetApplicationParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ConfigSource) validateSetConfigParameters(val *map[string]*string) error {
	return nil
}

func (j *jsiiProxy_ConfigSource) validateSetConnectionParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ConfigSource) validateSetCountParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ConfigSource) validateSetLifecycleParameters(val *cdktf.TerraformResourceLifecycle) error {
	return nil
}

func (j *jsiiProxy_ConfigSource) validateSetProjectParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ConfigSource) validateSetProvisionersParameters(val *[]interface{}) error {
	return nil
}

func (j *jsiiProxy_ConfigSource) validateSetScopeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ConfigSource) validateSetTypeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ConfigSource) validateSetWorkspaceParameters(val *string) error {
	return nil
}

func validateNewConfigSourceParameters(scope constructs.Construct, id *string, config *ConfigSourceConfig) error {
	return nil
}

