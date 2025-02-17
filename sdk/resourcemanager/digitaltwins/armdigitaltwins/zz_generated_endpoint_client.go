//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdigitaltwins

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

// EndpointClient contains the methods for the DigitalTwinsEndpoint group.
// Don't use this type directly, use NewEndpointClient() instead.
type EndpointClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewEndpointClient creates a new instance of EndpointClient with the specified values.
// subscriptionID - The subscription identifier.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewEndpointClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*EndpointClient, error) {
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
	client := &EndpointClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Create or update DigitalTwinsInstance endpoint.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group that contains the DigitalTwinsInstance.
// resourceName - The name of the DigitalTwinsInstance.
// endpointName - Name of Endpoint Resource.
// endpointDescription - The DigitalTwinsInstance endpoint metadata and security metadata.
// options - EndpointClientBeginCreateOrUpdateOptions contains the optional parameters for the EndpointClient.BeginCreateOrUpdate
// method.
func (client *EndpointClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, resourceName string, endpointName string, endpointDescription EndpointResource, options *EndpointClientBeginCreateOrUpdateOptions) (*armruntime.Poller[EndpointClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, resourceName, endpointName, endpointDescription, options)
		if err != nil {
			return nil, err
		}
		return armruntime.NewPoller[EndpointClientCreateOrUpdateResponse](resp, client.pl, nil)
	} else {
		return armruntime.NewPollerFromResumeToken[EndpointClientCreateOrUpdateResponse](options.ResumeToken, client.pl, nil)
	}
}

// CreateOrUpdate - Create or update DigitalTwinsInstance endpoint.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *EndpointClient) createOrUpdate(ctx context.Context, resourceGroupName string, resourceName string, endpointName string, endpointDescription EndpointResource, options *EndpointClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, resourceName, endpointName, endpointDescription, options)
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

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *EndpointClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, endpointName string, endpointDescription EndpointResource, options *EndpointClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DigitalTwins/digitalTwinsInstances/{resourceName}/endpoints/{endpointName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	if endpointName == "" {
		return nil, errors.New("parameter endpointName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{endpointName}", url.PathEscape(endpointName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-30-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, endpointDescription)
}

// BeginDelete - Delete a DigitalTwinsInstance endpoint.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group that contains the DigitalTwinsInstance.
// resourceName - The name of the DigitalTwinsInstance.
// endpointName - Name of Endpoint Resource.
// options - EndpointClientBeginDeleteOptions contains the optional parameters for the EndpointClient.BeginDelete method.
func (client *EndpointClient) BeginDelete(ctx context.Context, resourceGroupName string, resourceName string, endpointName string, options *EndpointClientBeginDeleteOptions) (*armruntime.Poller[EndpointClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, resourceName, endpointName, options)
		if err != nil {
			return nil, err
		}
		return armruntime.NewPoller[EndpointClientDeleteResponse](resp, client.pl, nil)
	} else {
		return armruntime.NewPollerFromResumeToken[EndpointClientDeleteResponse](options.ResumeToken, client.pl, nil)
	}
}

// Delete - Delete a DigitalTwinsInstance endpoint.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *EndpointClient) deleteOperation(ctx context.Context, resourceGroupName string, resourceName string, endpointName string, options *EndpointClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, resourceName, endpointName, options)
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
func (client *EndpointClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, endpointName string, options *EndpointClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DigitalTwins/digitalTwinsInstances/{resourceName}/endpoints/{endpointName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	if endpointName == "" {
		return nil, errors.New("parameter endpointName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{endpointName}", url.PathEscape(endpointName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-30-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// Get - Get DigitalTwinsInstances Endpoint.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group that contains the DigitalTwinsInstance.
// resourceName - The name of the DigitalTwinsInstance.
// endpointName - Name of Endpoint Resource.
// options - EndpointClientGetOptions contains the optional parameters for the EndpointClient.Get method.
func (client *EndpointClient) Get(ctx context.Context, resourceGroupName string, resourceName string, endpointName string, options *EndpointClientGetOptions) (EndpointClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, resourceName, endpointName, options)
	if err != nil {
		return EndpointClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return EndpointClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return EndpointClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *EndpointClient) getCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, endpointName string, options *EndpointClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DigitalTwins/digitalTwinsInstances/{resourceName}/endpoints/{endpointName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	if endpointName == "" {
		return nil, errors.New("parameter endpointName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{endpointName}", url.PathEscape(endpointName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-30-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *EndpointClient) getHandleResponse(resp *http.Response) (EndpointClientGetResponse, error) {
	result := EndpointClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.EndpointResource); err != nil {
		return EndpointClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Get DigitalTwinsInstance Endpoints.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group that contains the DigitalTwinsInstance.
// resourceName - The name of the DigitalTwinsInstance.
// options - EndpointClientListOptions contains the optional parameters for the EndpointClient.List method.
func (client *EndpointClient) NewListPager(resourceGroupName string, resourceName string, options *EndpointClientListOptions) *runtime.Pager[EndpointClientListResponse] {
	return runtime.NewPager(runtime.PageProcessor[EndpointClientListResponse]{
		More: func(page EndpointClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *EndpointClientListResponse) (EndpointClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, resourceGroupName, resourceName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return EndpointClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return EndpointClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return EndpointClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *EndpointClient) listCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, options *EndpointClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DigitalTwins/digitalTwinsInstances/{resourceName}/endpoints"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-30-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *EndpointClient) listHandleResponse(resp *http.Response) (EndpointClientListResponse, error) {
	result := EndpointClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.EndpointResourceListResult); err != nil {
		return EndpointClientListResponse{}, err
	}
	return result, nil
}
