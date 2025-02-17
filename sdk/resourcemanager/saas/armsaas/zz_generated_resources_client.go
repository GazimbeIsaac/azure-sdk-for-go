//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armsaas

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

// ResourcesClient contains the methods for the SaasResources group.
// Don't use this type directly, use NewResourcesClient() instead.
type ResourcesClient struct {
	host string
	pl   runtime.Pipeline
}

// NewResourcesClient creates a new instance of ResourcesClient with the specified values.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewResourcesClient(credential azcore.TokenCredential, options *arm.ClientOptions) (*ResourcesClient, error) {
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
	client := &ResourcesClient{
		host: ep,
		pl:   pl,
	}
	return client, nil
}

// NewListPager - Get All Resources
// If the operation fails it returns an *azcore.ResponseError type.
// options - ResourcesClientListOptions contains the optional parameters for the ResourcesClient.List method.
func (client *ResourcesClient) NewListPager(options *ResourcesClientListOptions) *runtime.Pager[ResourcesClientListResponse] {
	return runtime.NewPager(runtime.PageProcessor[ResourcesClientListResponse]{
		More: func(page ResourcesClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ResourcesClientListResponse) (ResourcesClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return ResourcesClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return ResourcesClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ResourcesClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *ResourcesClient) listCreateRequest(ctx context.Context, options *ResourcesClientListOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.SaaS/saasresources"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-03-01-beta")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *ResourcesClient) listHandleResponse(resp *http.Response) (ResourcesClientListResponse, error) {
	result := ResourcesClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ResourceResponseWithContinuation); err != nil {
		return ResourcesClientListResponse{}, err
	}
	return result, nil
}

// ListAccessToken - Gets the ISV access token for a SaaS resource.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceID - The Saas resource ID. This is a GUID-formatted string (e.g. 00000000-0000-0000-0000-000000000000)
// options - ResourcesClientListAccessTokenOptions contains the optional parameters for the ResourcesClient.ListAccessToken
// method.
func (client *ResourcesClient) ListAccessToken(ctx context.Context, resourceID string, options *ResourcesClientListAccessTokenOptions) (ResourcesClientListAccessTokenResponse, error) {
	req, err := client.listAccessTokenCreateRequest(ctx, resourceID, options)
	if err != nil {
		return ResourcesClientListAccessTokenResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ResourcesClientListAccessTokenResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ResourcesClientListAccessTokenResponse{}, runtime.NewResponseError(resp)
	}
	return client.listAccessTokenHandleResponse(resp)
}

// listAccessTokenCreateRequest creates the ListAccessToken request.
func (client *ResourcesClient) listAccessTokenCreateRequest(ctx context.Context, resourceID string, options *ResourcesClientListAccessTokenOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.SaaS/saasresources/{resourceId}/listAccessToken"
	if resourceID == "" {
		return nil, errors.New("parameter resourceID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceId}", url.PathEscape(resourceID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-03-01-beta")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listAccessTokenHandleResponse handles the ListAccessToken response.
func (client *ResourcesClient) listAccessTokenHandleResponse(resp *http.Response) (ResourcesClientListAccessTokenResponse, error) {
	result := ResourcesClientListAccessTokenResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AccessTokenResult); err != nil {
		return ResourcesClientListAccessTokenResponse{}, err
	}
	return result, nil
}
