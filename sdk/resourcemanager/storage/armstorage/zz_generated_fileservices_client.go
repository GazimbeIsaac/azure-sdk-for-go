//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armstorage

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// FileServicesClient contains the methods for the FileServices group.
// Don't use this type directly, use NewFileServicesClient() instead.
type FileServicesClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewFileServicesClient creates a new instance of FileServicesClient with the specified values.
// subscriptionID - The ID of the target subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewFileServicesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*FileServicesClient, error) {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := cloud.AzurePublicCloud.Services[cloud.ResourceManager].Endpoint
	if c, ok := options.Cloud.Services[cloud.ResourceManager]; ok {
		ep = c.Endpoint
	}
	pl, err := armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	client := &FileServicesClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// GetServiceProperties - Gets the properties of file services in storage accounts, including CORS (Cross-Origin Resource
// Sharing) rules.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group within the user's subscription. The name is case insensitive.
// accountName - The name of the storage account within the specified resource group. Storage account names must be between
// 3 and 24 characters in length and use numbers and lower-case letters only.
// options - FileServicesClientGetServicePropertiesOptions contains the optional parameters for the FileServicesClient.GetServiceProperties
// method.
func (client *FileServicesClient) GetServiceProperties(ctx context.Context, resourceGroupName string, accountName string, options *FileServicesClientGetServicePropertiesOptions) (FileServicesClientGetServicePropertiesResponse, error) {
	req, err := client.getServicePropertiesCreateRequest(ctx, resourceGroupName, accountName, options)
	if err != nil {
		return FileServicesClientGetServicePropertiesResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return FileServicesClientGetServicePropertiesResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return FileServicesClientGetServicePropertiesResponse{}, runtime.NewResponseError(resp)
	}
	return client.getServicePropertiesHandleResponse(resp)
}

// getServicePropertiesCreateRequest creates the GetServiceProperties request.
func (client *FileServicesClient) getServicePropertiesCreateRequest(ctx context.Context, resourceGroupName string, accountName string, options *FileServicesClientGetServicePropertiesOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Storage/storageAccounts/{accountName}/fileServices/{FileServicesName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	urlPath = strings.ReplaceAll(urlPath, "{FileServicesName}", url.PathEscape("default"))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getServicePropertiesHandleResponse handles the GetServiceProperties response.
func (client *FileServicesClient) getServicePropertiesHandleResponse(resp *http.Response) (FileServicesClientGetServicePropertiesResponse, error) {
	result := FileServicesClientGetServicePropertiesResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.FileServiceProperties); err != nil {
		return FileServicesClientGetServicePropertiesResponse{}, err
	}
	return result, nil
}

// List - List all file services in storage accounts
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group within the user's subscription. The name is case insensitive.
// accountName - The name of the storage account within the specified resource group. Storage account names must be between
// 3 and 24 characters in length and use numbers and lower-case letters only.
// options - FileServicesClientListOptions contains the optional parameters for the FileServicesClient.List method.
func (client *FileServicesClient) List(ctx context.Context, resourceGroupName string, accountName string, options *FileServicesClientListOptions) (FileServicesClientListResponse, error) {
	req, err := client.listCreateRequest(ctx, resourceGroupName, accountName, options)
	if err != nil {
		return FileServicesClientListResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return FileServicesClientListResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return FileServicesClientListResponse{}, runtime.NewResponseError(resp)
	}
	return client.listHandleResponse(resp)
}

// listCreateRequest creates the List request.
func (client *FileServicesClient) listCreateRequest(ctx context.Context, resourceGroupName string, accountName string, options *FileServicesClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Storage/storageAccounts/{accountName}/fileServices"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *FileServicesClient) listHandleResponse(resp *http.Response) (FileServicesClientListResponse, error) {
	result := FileServicesClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.FileServiceItems); err != nil {
		return FileServicesClientListResponse{}, err
	}
	return result, nil
}

// SetServiceProperties - Sets the properties of file services in storage accounts, including CORS (Cross-Origin Resource
// Sharing) rules.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group within the user's subscription. The name is case insensitive.
// accountName - The name of the storage account within the specified resource group. Storage account names must be between
// 3 and 24 characters in length and use numbers and lower-case letters only.
// parameters - The properties of file services in storage accounts, including CORS (Cross-Origin Resource Sharing) rules.
// options - FileServicesClientSetServicePropertiesOptions contains the optional parameters for the FileServicesClient.SetServiceProperties
// method.
func (client *FileServicesClient) SetServiceProperties(ctx context.Context, resourceGroupName string, accountName string, parameters FileServiceProperties, options *FileServicesClientSetServicePropertiesOptions) (FileServicesClientSetServicePropertiesResponse, error) {
	req, err := client.setServicePropertiesCreateRequest(ctx, resourceGroupName, accountName, parameters, options)
	if err != nil {
		return FileServicesClientSetServicePropertiesResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return FileServicesClientSetServicePropertiesResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return FileServicesClientSetServicePropertiesResponse{}, runtime.NewResponseError(resp)
	}
	return client.setServicePropertiesHandleResponse(resp)
}

// setServicePropertiesCreateRequest creates the SetServiceProperties request.
func (client *FileServicesClient) setServicePropertiesCreateRequest(ctx context.Context, resourceGroupName string, accountName string, parameters FileServiceProperties, options *FileServicesClientSetServicePropertiesOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Storage/storageAccounts/{accountName}/fileServices/{FileServicesName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	urlPath = strings.ReplaceAll(urlPath, "{FileServicesName}", url.PathEscape("default"))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// setServicePropertiesHandleResponse handles the SetServiceProperties response.
func (client *FileServicesClient) setServicePropertiesHandleResponse(resp *http.Response) (FileServicesClientSetServicePropertiesResponse, error) {
	result := FileServicesClientSetServicePropertiesResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.FileServiceProperties); err != nil {
		return FileServicesClientSetServicePropertiesResponse{}, err
	}
	return result, nil
}
