// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package authmethod

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AuthMethod) validateAddMoveTargetParameters(moveTarget *string) error {
	return nil
}

func (a *jsiiProxy_AuthMethod) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (a *jsiiProxy_AuthMethod) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (a *jsiiProxy_AuthMethod) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (a *jsiiProxy_AuthMethod) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (a *jsiiProxy_AuthMethod) validateGetListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (a *jsiiProxy_AuthMethod) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (a *jsiiProxy_AuthMethod) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (a *jsiiProxy_AuthMethod) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (a *jsiiProxy_AuthMethod) validateGetStringAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (a *jsiiProxy_AuthMethod) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (a *jsiiProxy_AuthMethod) validateImportFromParameters(id *string) error {
	return nil
}

func (a *jsiiProxy_AuthMethod) validateInterpolationForAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (a *jsiiProxy_AuthMethod) validateMoveFromIdParameters(id *string) error {
	return nil
}

func (a *jsiiProxy_AuthMethod) validateMoveToParameters(moveTarget *string, index interface{}) error {
	return nil
}

func (a *jsiiProxy_AuthMethod) validateMoveToIdParameters(id *string) error {
	return nil
}

func (a *jsiiProxy_AuthMethod) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func validateAuthMethod_GenerateConfigForImportParameters(scope constructs.Construct, importToId *string, importFromId *string) error {
	return nil
}

func validateAuthMethod_IsConstructParameters(x interface{}) error {
	return nil
}

func validateAuthMethod_IsTerraformElementParameters(x interface{}) error {
	return nil
}

func validateAuthMethod_IsTerraformResourceParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_AuthMethod) validateSetAccessorSelectorParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_AuthMethod) validateSetAllowedRedirectUrisParameters(val *[]*string) error {
	return nil
}

func (j *jsiiProxy_AuthMethod) validateSetAudsParameters(val *[]*string) error {
	return nil
}

func (j *jsiiProxy_AuthMethod) validateSetClaimMappingsParameters(val *map[string]*string) error {
	return nil
}

func (j *jsiiProxy_AuthMethod) validateSetClientIdParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_AuthMethod) validateSetClientSecretParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_AuthMethod) validateSetConnectionParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_AuthMethod) validateSetCountParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_AuthMethod) validateSetDescriptionParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_AuthMethod) validateSetDiscoveryCaPemParameters(val *[]*string) error {
	return nil
}

func (j *jsiiProxy_AuthMethod) validateSetDiscoveryUrlParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_AuthMethod) validateSetDisplayNameParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_AuthMethod) validateSetLifecycleParameters(val *cdktf.TerraformResourceLifecycle) error {
	return nil
}

func (j *jsiiProxy_AuthMethod) validateSetListClaimMappingsParameters(val *map[string]*string) error {
	return nil
}

func (j *jsiiProxy_AuthMethod) validateSetNameParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_AuthMethod) validateSetProvisionersParameters(val *[]interface{}) error {
	return nil
}

func (j *jsiiProxy_AuthMethod) validateSetScopesParameters(val *[]*string) error {
	return nil
}

func (j *jsiiProxy_AuthMethod) validateSetSigningAlgsParameters(val *[]*string) error {
	return nil
}

func validateNewAuthMethodParameters(scope constructs.Construct, id *string, config *AuthMethodConfig) error {
	return nil
}

