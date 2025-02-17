//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armloadtestservice

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

// LoadTestsClient contains the methods for the LoadTests group.
// Don't use this type directly, use NewLoadTestsClient() instead.
type LoadTestsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewLoadTestsClient creates a new instance of LoadTestsClient with the specified values.
// subscriptionID - The ID of the target subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewLoadTestsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*LoadTestsClient, error) {
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
	client := &LoadTestsClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Create or update LoadTest resource.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// loadTestName - Load Test name.
// loadTestResource - LoadTest resource data
// options - LoadTestsClientBeginCreateOrUpdateOptions contains the optional parameters for the LoadTestsClient.BeginCreateOrUpdate
// method.
func (client *LoadTestsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, loadTestName string, loadTestResource LoadTestResource, options *LoadTestsClientBeginCreateOrUpdateOptions) (*armruntime.Poller[LoadTestsClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, loadTestName, loadTestResource, options)
		if err != nil {
			return nil, err
		}
		return armruntime.NewPoller(resp, client.pl, &armruntime.NewPollerOptions[LoadTestsClientCreateOrUpdateResponse]{
			FinalStateVia: armruntime.FinalStateViaAzureAsyncOp,
		})
	} else {
		return armruntime.NewPollerFromResumeToken[LoadTestsClientCreateOrUpdateResponse](options.ResumeToken, client.pl, nil)
	}
}

// CreateOrUpdate - Create or update LoadTest resource.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *LoadTestsClient) createOrUpdate(ctx context.Context, resourceGroupName string, loadTestName string, loadTestResource LoadTestResource, options *LoadTestsClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, loadTestName, loadTestResource, options)
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
func (client *LoadTestsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, loadTestName string, loadTestResource LoadTestResource, options *LoadTestsClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.LoadTestService/loadTests/{loadTestName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if loadTestName == "" {
		return nil, errors.New("parameter loadTestName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{loadTestName}", url.PathEscape(loadTestName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-04-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, loadTestResource)
}

// BeginDelete - Delete a LoadTest resource.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// loadTestName - Load Test name.
// options - LoadTestsClientBeginDeleteOptions contains the optional parameters for the LoadTestsClient.BeginDelete method.
func (client *LoadTestsClient) BeginDelete(ctx context.Context, resourceGroupName string, loadTestName string, options *LoadTestsClientBeginDeleteOptions) (*armruntime.Poller[LoadTestsClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, loadTestName, options)
		if err != nil {
			return nil, err
		}
		return armruntime.NewPoller(resp, client.pl, &armruntime.NewPollerOptions[LoadTestsClientDeleteResponse]{
			FinalStateVia: armruntime.FinalStateViaLocation,
		})
	} else {
		return armruntime.NewPollerFromResumeToken[LoadTestsClientDeleteResponse](options.ResumeToken, client.pl, nil)
	}
}

// Delete - Delete a LoadTest resource.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *LoadTestsClient) deleteOperation(ctx context.Context, resourceGroupName string, loadTestName string, options *LoadTestsClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, loadTestName, options)
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
func (client *LoadTestsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, loadTestName string, options *LoadTestsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.LoadTestService/loadTests/{loadTestName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if loadTestName == "" {
		return nil, errors.New("parameter loadTestName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{loadTestName}", url.PathEscape(loadTestName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-04-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// Get - Get a LoadTest resource.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// loadTestName - Load Test name.
// options - LoadTestsClientGetOptions contains the optional parameters for the LoadTestsClient.Get method.
func (client *LoadTestsClient) Get(ctx context.Context, resourceGroupName string, loadTestName string, options *LoadTestsClientGetOptions) (LoadTestsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, loadTestName, options)
	if err != nil {
		return LoadTestsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return LoadTestsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return LoadTestsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *LoadTestsClient) getCreateRequest(ctx context.Context, resourceGroupName string, loadTestName string, options *LoadTestsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.LoadTestService/loadTests/{loadTestName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if loadTestName == "" {
		return nil, errors.New("parameter loadTestName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{loadTestName}", url.PathEscape(loadTestName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-04-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *LoadTestsClient) getHandleResponse(resp *http.Response) (LoadTestsClientGetResponse, error) {
	result := LoadTestsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.LoadTestResource); err != nil {
		return LoadTestsClientGetResponse{}, err
	}
	return result, nil
}

// NewListByResourceGroupPager - Lists loadtest resources in a resource group.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// options - LoadTestsClientListByResourceGroupOptions contains the optional parameters for the LoadTestsClient.ListByResourceGroup
// method.
func (client *LoadTestsClient) NewListByResourceGroupPager(resourceGroupName string, options *LoadTestsClientListByResourceGroupOptions) *runtime.Pager[LoadTestsClientListByResourceGroupResponse] {
	return runtime.NewPager(runtime.PageProcessor[LoadTestsClientListByResourceGroupResponse]{
		More: func(page LoadTestsClientListByResourceGroupResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *LoadTestsClientListByResourceGroupResponse) (LoadTestsClientListByResourceGroupResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return LoadTestsClientListByResourceGroupResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return LoadTestsClientListByResourceGroupResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return LoadTestsClientListByResourceGroupResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByResourceGroupHandleResponse(resp)
		},
	})
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *LoadTestsClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *LoadTestsClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.LoadTestService/loadTests"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-04-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *LoadTestsClient) listByResourceGroupHandleResponse(resp *http.Response) (LoadTestsClientListByResourceGroupResponse, error) {
	result := LoadTestsClientListByResourceGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.LoadTestResourcePageList); err != nil {
		return LoadTestsClientListByResourceGroupResponse{}, err
	}
	return result, nil
}

// NewListBySubscriptionPager - Lists loadtests resources in a subscription.
// If the operation fails it returns an *azcore.ResponseError type.
// options - LoadTestsClientListBySubscriptionOptions contains the optional parameters for the LoadTestsClient.ListBySubscription
// method.
func (client *LoadTestsClient) NewListBySubscriptionPager(options *LoadTestsClientListBySubscriptionOptions) *runtime.Pager[LoadTestsClientListBySubscriptionResponse] {
	return runtime.NewPager(runtime.PageProcessor[LoadTestsClientListBySubscriptionResponse]{
		More: func(page LoadTestsClientListBySubscriptionResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *LoadTestsClientListBySubscriptionResponse) (LoadTestsClientListBySubscriptionResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listBySubscriptionCreateRequest(ctx, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return LoadTestsClientListBySubscriptionResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return LoadTestsClientListBySubscriptionResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return LoadTestsClientListBySubscriptionResponse{}, runtime.NewResponseError(resp)
			}
			return client.listBySubscriptionHandleResponse(resp)
		},
	})
}

// listBySubscriptionCreateRequest creates the ListBySubscription request.
func (client *LoadTestsClient) listBySubscriptionCreateRequest(ctx context.Context, options *LoadTestsClientListBySubscriptionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.LoadTestService/loadTests"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-04-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listBySubscriptionHandleResponse handles the ListBySubscription response.
func (client *LoadTestsClient) listBySubscriptionHandleResponse(resp *http.Response) (LoadTestsClientListBySubscriptionResponse, error) {
	result := LoadTestsClientListBySubscriptionResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.LoadTestResourcePageList); err != nil {
		return LoadTestsClientListBySubscriptionResponse{}, err
	}
	return result, nil
}

// BeginUpdate - Update a loadtest resource.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// loadTestName - Load Test name.
// loadTestResourcePatchRequestBody - LoadTest resource update data
// options - LoadTestsClientBeginUpdateOptions contains the optional parameters for the LoadTestsClient.BeginUpdate method.
func (client *LoadTestsClient) BeginUpdate(ctx context.Context, resourceGroupName string, loadTestName string, loadTestResourcePatchRequestBody LoadTestResourcePatchRequestBody, options *LoadTestsClientBeginUpdateOptions) (*armruntime.Poller[LoadTestsClientUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.update(ctx, resourceGroupName, loadTestName, loadTestResourcePatchRequestBody, options)
		if err != nil {
			return nil, err
		}
		return armruntime.NewPoller(resp, client.pl, &armruntime.NewPollerOptions[LoadTestsClientUpdateResponse]{
			FinalStateVia: armruntime.FinalStateViaAzureAsyncOp,
		})
	} else {
		return armruntime.NewPollerFromResumeToken[LoadTestsClientUpdateResponse](options.ResumeToken, client.pl, nil)
	}
}

// Update - Update a loadtest resource.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *LoadTestsClient) update(ctx context.Context, resourceGroupName string, loadTestName string, loadTestResourcePatchRequestBody LoadTestResourcePatchRequestBody, options *LoadTestsClientBeginUpdateOptions) (*http.Response, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, loadTestName, loadTestResourcePatchRequestBody, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// updateCreateRequest creates the Update request.
func (client *LoadTestsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, loadTestName string, loadTestResourcePatchRequestBody LoadTestResourcePatchRequestBody, options *LoadTestsClientBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.LoadTestService/loadTests/{loadTestName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if loadTestName == "" {
		return nil, errors.New("parameter loadTestName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{loadTestName}", url.PathEscape(loadTestName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-04-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, loadTestResourcePatchRequestBody)
}
