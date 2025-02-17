//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armservicelinker

// LinkerClientCreateOrUpdateResponse contains the response from method LinkerClient.CreateOrUpdate.
type LinkerClientCreateOrUpdateResponse struct {
	LinkerResource
}

// LinkerClientDeleteResponse contains the response from method LinkerClient.Delete.
type LinkerClientDeleteResponse struct {
	// placeholder for future response values
}

// LinkerClientGetResponse contains the response from method LinkerClient.Get.
type LinkerClientGetResponse struct {
	LinkerResource
}

// LinkerClientListConfigurationsResponse contains the response from method LinkerClient.ListConfigurations.
type LinkerClientListConfigurationsResponse struct {
	SourceConfigurationResult
}

// LinkerClientListResponse contains the response from method LinkerClient.List.
type LinkerClientListResponse struct {
	LinkerList
}

// LinkerClientUpdateResponse contains the response from method LinkerClient.Update.
type LinkerClientUpdateResponse struct {
	LinkerResource
}

// LinkerClientValidateResponse contains the response from method LinkerClient.Validate.
type LinkerClientValidateResponse struct {
	ValidateResult
}

// OperationsClientListResponse contains the response from method OperationsClient.List.
type OperationsClientListResponse struct {
	OperationListResult
}
