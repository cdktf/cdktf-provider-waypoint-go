// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package provider

// Building without runtime type checking enabled, so all the below just return nil

func (w *jsiiProxy_WaypointProvider) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (w *jsiiProxy_WaypointProvider) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func validateWaypointProvider_GenerateConfigForImportParameters(scope constructs.Construct, importToId *string, importFromId *string) error {
	return nil
}

func validateWaypointProvider_IsConstructParameters(x interface{}) error {
	return nil
}

func validateWaypointProvider_IsTerraformElementParameters(x interface{}) error {
	return nil
}

func validateWaypointProvider_IsTerraformProviderParameters(x interface{}) error {
	return nil
}

func validateNewWaypointProviderParameters(scope constructs.Construct, id *string, config *WaypointProviderConfig) error {
	return nil
}

