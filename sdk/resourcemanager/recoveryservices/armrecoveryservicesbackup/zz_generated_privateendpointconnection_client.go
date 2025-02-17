//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armrecoveryservicesbackup

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

// PrivateEndpointConnectionClient contains the methods for the PrivateEndpointConnection group.
// Don't use this type directly, use NewPrivateEndpointConnectionClient() instead.
type PrivateEndpointConnectionClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewPrivateEndpointConnectionClient creates a new instance of PrivateEndpointConnectionClient with the specified values.
// subscriptionID - The subscription Id.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewPrivateEndpointConnectionClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*PrivateEndpointConnectionClient, error) {
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
	client := &PrivateEndpointConnectionClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// BeginDelete - Delete Private Endpoint requests. This call is made by Backup Admin.
// If the operation fails it returns an *azcore.ResponseError type.
// vaultName - The name of the recovery services vault.
// resourceGroupName - The name of the resource group where the recovery services vault is present.
// privateEndpointConnectionName - The name of the private endpoint connection.
// options - PrivateEndpointConnectionClientBeginDeleteOptions contains the optional parameters for the PrivateEndpointConnectionClient.BeginDelete
// method.
func (client *PrivateEndpointConnectionClient) BeginDelete(ctx context.Context, vaultName string, resourceGroupName string, privateEndpointConnectionName string, options *PrivateEndpointConnectionClientBeginDeleteOptions) (*armruntime.Poller[PrivateEndpointConnectionClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, vaultName, resourceGroupName, privateEndpointConnectionName, options)
		if err != nil {
			return nil, err
		}
		return armruntime.NewPoller[PrivateEndpointConnectionClientDeleteResponse](resp, client.pl, nil)
	} else {
		return armruntime.NewPollerFromResumeToken[PrivateEndpointConnectionClientDeleteResponse](options.ResumeToken, client.pl, nil)
	}
}

// Delete - Delete Private Endpoint requests. This call is made by Backup Admin.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *PrivateEndpointConnectionClient) deleteOperation(ctx context.Context, vaultName string, resourceGroupName string, privateEndpointConnectionName string, options *PrivateEndpointConnectionClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, vaultName, resourceGroupName, privateEndpointConnectionName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *PrivateEndpointConnectionClient) deleteCreateRequest(ctx context.Context, vaultName string, resourceGroupName string, privateEndpointConnectionName string, options *PrivateEndpointConnectionClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{vaultName}/privateEndpointConnections/{privateEndpointConnectionName}"
	if vaultName == "" {
		return nil, errors.New("parameter vaultName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vaultName}", url.PathEscape(vaultName))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if privateEndpointConnectionName == "" {
		return nil, errors.New("parameter privateEndpointConnectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{privateEndpointConnectionName}", url.PathEscape(privateEndpointConnectionName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-12-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// Get - Get Private Endpoint Connection. This call is made by Backup Admin.
// If the operation fails it returns an *azcore.ResponseError type.
// vaultName - The name of the recovery services vault.
// resourceGroupName - The name of the resource group where the recovery services vault is present.
// privateEndpointConnectionName - The name of the private endpoint connection.
// options - PrivateEndpointConnectionClientGetOptions contains the optional parameters for the PrivateEndpointConnectionClient.Get
// method.
func (client *PrivateEndpointConnectionClient) Get(ctx context.Context, vaultName string, resourceGroupName string, privateEndpointConnectionName string, options *PrivateEndpointConnectionClientGetOptions) (PrivateEndpointConnectionClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, vaultName, resourceGroupName, privateEndpointConnectionName, options)
	if err != nil {
		return PrivateEndpointConnectionClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return PrivateEndpointConnectionClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return PrivateEndpointConnectionClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *PrivateEndpointConnectionClient) getCreateRequest(ctx context.Context, vaultName string, resourceGroupName string, privateEndpointConnectionName string, options *PrivateEndpointConnectionClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{vaultName}/privateEndpointConnections/{privateEndpointConnectionName}"
	if vaultName == "" {
		return nil, errors.New("parameter vaultName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vaultName}", url.PathEscape(vaultName))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if privateEndpointConnectionName == "" {
		return nil, errors.New("parameter privateEndpointConnectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{privateEndpointConnectionName}", url.PathEscape(privateEndpointConnectionName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-12-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *PrivateEndpointConnectionClient) getHandleResponse(resp *http.Response) (PrivateEndpointConnectionClientGetResponse, error) {
	result := PrivateEndpointConnectionClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PrivateEndpointConnectionResource); err != nil {
		return PrivateEndpointConnectionClientGetResponse{}, err
	}
	return result, nil
}

// BeginPut - Approve or Reject Private Endpoint requests. This call is made by Backup Admin.
// If the operation fails it returns an *azcore.ResponseError type.
// vaultName - The name of the recovery services vault.
// resourceGroupName - The name of the resource group where the recovery services vault is present.
// privateEndpointConnectionName - The name of the private endpoint connection.
// parameters - Request body for operation
// options - PrivateEndpointConnectionClientBeginPutOptions contains the optional parameters for the PrivateEndpointConnectionClient.BeginPut
// method.
func (client *PrivateEndpointConnectionClient) BeginPut(ctx context.Context, vaultName string, resourceGroupName string, privateEndpointConnectionName string, parameters PrivateEndpointConnectionResource, options *PrivateEndpointConnectionClientBeginPutOptions) (*armruntime.Poller[PrivateEndpointConnectionClientPutResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.put(ctx, vaultName, resourceGroupName, privateEndpointConnectionName, parameters, options)
		if err != nil {
			return nil, err
		}
		return armruntime.NewPoller[PrivateEndpointConnectionClientPutResponse](resp, client.pl, nil)
	} else {
		return armruntime.NewPollerFromResumeToken[PrivateEndpointConnectionClientPutResponse](options.ResumeToken, client.pl, nil)
	}
}

// Put - Approve or Reject Private Endpoint requests. This call is made by Backup Admin.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *PrivateEndpointConnectionClient) put(ctx context.Context, vaultName string, resourceGroupName string, privateEndpointConnectionName string, parameters PrivateEndpointConnectionResource, options *PrivateEndpointConnectionClientBeginPutOptions) (*http.Response, error) {
	req, err := client.putCreateRequest(ctx, vaultName, resourceGroupName, privateEndpointConnectionName, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// putCreateRequest creates the Put request.
func (client *PrivateEndpointConnectionClient) putCreateRequest(ctx context.Context, vaultName string, resourceGroupName string, privateEndpointConnectionName string, parameters PrivateEndpointConnectionResource, options *PrivateEndpointConnectionClientBeginPutOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{vaultName}/privateEndpointConnections/{privateEndpointConnectionName}"
	if vaultName == "" {
		return nil, errors.New("parameter vaultName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vaultName}", url.PathEscape(vaultName))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if privateEndpointConnectionName == "" {
		return nil, errors.New("parameter privateEndpointConnectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{privateEndpointConnectionName}", url.PathEscape(privateEndpointConnectionName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-12-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}
