//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armadvisor

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

// RecommendationMetadataClient contains the methods for the RecommendationMetadata group.
// Don't use this type directly, use NewRecommendationMetadataClient() instead.
type RecommendationMetadataClient struct {
	host string
	pl   runtime.Pipeline
}

// NewRecommendationMetadataClient creates a new instance of RecommendationMetadataClient with the specified values.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewRecommendationMetadataClient(credential azcore.TokenCredential, options *arm.ClientOptions) (*RecommendationMetadataClient, error) {
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
	client := &RecommendationMetadataClient{
		host: ep,
		pl:   pl,
	}
	return client, nil
}

// Get - Gets the metadata entity.
// If the operation fails it returns an *azcore.ResponseError type.
// name - Name of metadata entity.
// options - RecommendationMetadataClientGetOptions contains the optional parameters for the RecommendationMetadataClient.Get
// method.
func (client *RecommendationMetadataClient) Get(ctx context.Context, name string, options *RecommendationMetadataClientGetOptions) (RecommendationMetadataClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, name, options)
	if err != nil {
		return RecommendationMetadataClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return RecommendationMetadataClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return RecommendationMetadataClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *RecommendationMetadataClient) getCreateRequest(ctx context.Context, name string, options *RecommendationMetadataClientGetOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Advisor/metadata/{name}"
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-01-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *RecommendationMetadataClient) getHandleResponse(resp *http.Response) (RecommendationMetadataClientGetResponse, error) {
	result := RecommendationMetadataClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.MetadataEntity); err != nil {
		return RecommendationMetadataClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Gets the list of metadata entities.
// If the operation fails it returns an *azcore.ResponseError type.
// options - RecommendationMetadataClientListOptions contains the optional parameters for the RecommendationMetadataClient.List
// method.
func (client *RecommendationMetadataClient) NewListPager(options *RecommendationMetadataClientListOptions) *runtime.Pager[RecommendationMetadataClientListResponse] {
	return runtime.NewPager(runtime.PageProcessor[RecommendationMetadataClientListResponse]{
		More: func(page RecommendationMetadataClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *RecommendationMetadataClientListResponse) (RecommendationMetadataClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return RecommendationMetadataClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return RecommendationMetadataClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return RecommendationMetadataClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *RecommendationMetadataClient) listCreateRequest(ctx context.Context, options *RecommendationMetadataClientListOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Advisor/metadata"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-01-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *RecommendationMetadataClient) listHandleResponse(resp *http.Response) (RecommendationMetadataClientListResponse, error) {
	result := RecommendationMetadataClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.MetadataEntityListResult); err != nil {
		return RecommendationMetadataClientListResponse{}, err
	}
	return result, nil
}
