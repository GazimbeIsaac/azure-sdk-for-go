//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armsubscription

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
)

// PolicyClient contains the methods for the SubscriptionPolicy group.
// Don't use this type directly, use NewPolicyClient() instead.
type PolicyClient struct {
	host string
	pl   runtime.Pipeline
}

// NewPolicyClient creates a new instance of PolicyClient with the specified values.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewPolicyClient(credential azcore.TokenCredential, options *arm.ClientOptions) (*PolicyClient, error) {
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
	client := &PolicyClient{
		host: ep,
		pl:   pl,
	}
	return client, nil
}

// AddUpdatePolicyForTenant - Create or Update Subscription tenant policy for user's tenant.
// If the operation fails it returns an *azcore.ResponseError type.
// options - PolicyClientAddUpdatePolicyForTenantOptions contains the optional parameters for the PolicyClient.AddUpdatePolicyForTenant
// method.
func (client *PolicyClient) AddUpdatePolicyForTenant(ctx context.Context, body PutTenantPolicyRequestProperties, options *PolicyClientAddUpdatePolicyForTenantOptions) (PolicyClientAddUpdatePolicyForTenantResponse, error) {
	req, err := client.addUpdatePolicyForTenantCreateRequest(ctx, body, options)
	if err != nil {
		return PolicyClientAddUpdatePolicyForTenantResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return PolicyClientAddUpdatePolicyForTenantResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return PolicyClientAddUpdatePolicyForTenantResponse{}, runtime.NewResponseError(resp)
	}
	return client.addUpdatePolicyForTenantHandleResponse(resp)
}

// addUpdatePolicyForTenantCreateRequest creates the AddUpdatePolicyForTenant request.
func (client *PolicyClient) addUpdatePolicyForTenantCreateRequest(ctx context.Context, body PutTenantPolicyRequestProperties, options *PolicyClientAddUpdatePolicyForTenantOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Subscription/policies/default"
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, body)
}

// addUpdatePolicyForTenantHandleResponse handles the AddUpdatePolicyForTenant response.
func (client *PolicyClient) addUpdatePolicyForTenantHandleResponse(resp *http.Response) (PolicyClientAddUpdatePolicyForTenantResponse, error) {
	result := PolicyClientAddUpdatePolicyForTenantResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.GetTenantPolicyResponse); err != nil {
		return PolicyClientAddUpdatePolicyForTenantResponse{}, err
	}
	return result, nil
}

// GetPolicyForTenant - Get the subscription tenant policy for the user's tenant.
// If the operation fails it returns an *azcore.ResponseError type.
// options - PolicyClientGetPolicyForTenantOptions contains the optional parameters for the PolicyClient.GetPolicyForTenant
// method.
func (client *PolicyClient) GetPolicyForTenant(ctx context.Context, options *PolicyClientGetPolicyForTenantOptions) (PolicyClientGetPolicyForTenantResponse, error) {
	req, err := client.getPolicyForTenantCreateRequest(ctx, options)
	if err != nil {
		return PolicyClientGetPolicyForTenantResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return PolicyClientGetPolicyForTenantResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return PolicyClientGetPolicyForTenantResponse{}, runtime.NewResponseError(resp)
	}
	return client.getPolicyForTenantHandleResponse(resp)
}

// getPolicyForTenantCreateRequest creates the GetPolicyForTenant request.
func (client *PolicyClient) getPolicyForTenantCreateRequest(ctx context.Context, options *PolicyClientGetPolicyForTenantOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Subscription/policies/default"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getPolicyForTenantHandleResponse handles the GetPolicyForTenant response.
func (client *PolicyClient) getPolicyForTenantHandleResponse(resp *http.Response) (PolicyClientGetPolicyForTenantResponse, error) {
	result := PolicyClientGetPolicyForTenantResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.GetTenantPolicyResponse); err != nil {
		return PolicyClientGetPolicyForTenantResponse{}, err
	}
	return result, nil
}

// NewListPolicyForTenantPager - Get the subscription tenant policy for the user's tenant.
// If the operation fails it returns an *azcore.ResponseError type.
// options - PolicyClientListPolicyForTenantOptions contains the optional parameters for the PolicyClient.ListPolicyForTenant
// method.
func (client *PolicyClient) NewListPolicyForTenantPager(options *PolicyClientListPolicyForTenantOptions) *runtime.Pager[PolicyClientListPolicyForTenantResponse] {
	return runtime.NewPager(runtime.PageProcessor[PolicyClientListPolicyForTenantResponse]{
		More: func(page PolicyClientListPolicyForTenantResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *PolicyClientListPolicyForTenantResponse) (PolicyClientListPolicyForTenantResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listPolicyForTenantCreateRequest(ctx, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return PolicyClientListPolicyForTenantResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return PolicyClientListPolicyForTenantResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return PolicyClientListPolicyForTenantResponse{}, runtime.NewResponseError(resp)
			}
			return client.listPolicyForTenantHandleResponse(resp)
		},
	})
}

// listPolicyForTenantCreateRequest creates the ListPolicyForTenant request.
func (client *PolicyClient) listPolicyForTenantCreateRequest(ctx context.Context, options *PolicyClientListPolicyForTenantOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Subscription/policies"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listPolicyForTenantHandleResponse handles the ListPolicyForTenant response.
func (client *PolicyClient) listPolicyForTenantHandleResponse(resp *http.Response) (PolicyClientListPolicyForTenantResponse, error) {
	result := PolicyClientListPolicyForTenantResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.GetTenantPolicyListResponse); err != nil {
		return PolicyClientListPolicyForTenantResponse{}, err
	}
	return result, nil
}
