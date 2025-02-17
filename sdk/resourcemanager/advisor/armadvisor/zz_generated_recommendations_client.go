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
	"strconv"
	"strings"
)

// RecommendationsClient contains the methods for the Recommendations group.
// Don't use this type directly, use NewRecommendationsClient() instead.
type RecommendationsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewRecommendationsClient creates a new instance of RecommendationsClient with the specified values.
// subscriptionID - The Azure subscription ID.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewRecommendationsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*RecommendationsClient, error) {
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
	client := &RecommendationsClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// Generate - Initiates the recommendation generation or computation process for a subscription. This operation is asynchronous.
// The generated recommendations are stored in a cache in the Advisor service.
// If the operation fails it returns an *azcore.ResponseError type.
// options - RecommendationsClientGenerateOptions contains the optional parameters for the RecommendationsClient.Generate
// method.
func (client *RecommendationsClient) Generate(ctx context.Context, options *RecommendationsClientGenerateOptions) (RecommendationsClientGenerateResponse, error) {
	req, err := client.generateCreateRequest(ctx, options)
	if err != nil {
		return RecommendationsClientGenerateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return RecommendationsClientGenerateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusAccepted) {
		return RecommendationsClientGenerateResponse{}, runtime.NewResponseError(resp)
	}
	return client.generateHandleResponse(resp)
}

// generateCreateRequest creates the Generate request.
func (client *RecommendationsClient) generateCreateRequest(ctx context.Context, options *RecommendationsClientGenerateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Advisor/generateRecommendations"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-01-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// generateHandleResponse handles the Generate response.
func (client *RecommendationsClient) generateHandleResponse(resp *http.Response) (RecommendationsClientGenerateResponse, error) {
	result := RecommendationsClientGenerateResponse{}
	if val := resp.Header.Get("Location"); val != "" {
		result.Location = &val
	}
	if val := resp.Header.Get("Retry-After"); val != "" {
		result.RetryAfter = &val
	}
	return result, nil
}

// Get - Obtains details of a cached recommendation.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceURI - The fully qualified Azure Resource Manager identifier of the resource to which the recommendation applies.
// recommendationID - The recommendation ID.
// options - RecommendationsClientGetOptions contains the optional parameters for the RecommendationsClient.Get method.
func (client *RecommendationsClient) Get(ctx context.Context, resourceURI string, recommendationID string, options *RecommendationsClientGetOptions) (RecommendationsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceURI, recommendationID, options)
	if err != nil {
		return RecommendationsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return RecommendationsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return RecommendationsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *RecommendationsClient) getCreateRequest(ctx context.Context, resourceURI string, recommendationID string, options *RecommendationsClientGetOptions) (*policy.Request, error) {
	urlPath := "/{resourceUri}/providers/Microsoft.Advisor/recommendations/{recommendationId}"
	if resourceURI == "" {
		return nil, errors.New("parameter resourceURI cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceUri}", url.PathEscape(resourceURI))
	if recommendationID == "" {
		return nil, errors.New("parameter recommendationID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{recommendationId}", url.PathEscape(recommendationID))
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
func (client *RecommendationsClient) getHandleResponse(resp *http.Response) (RecommendationsClientGetResponse, error) {
	result := RecommendationsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ResourceRecommendationBase); err != nil {
		return RecommendationsClientGetResponse{}, err
	}
	return result, nil
}

// GetGenerateStatus - Retrieves the status of the recommendation computation or generation process. Invoke this API after
// calling the generation recommendation. The URI of this API is returned in the Location field of the
// response header.
// If the operation fails it returns an *azcore.ResponseError type.
// operationID - The operation ID, which can be found from the Location field in the generate recommendation response header.
// options - RecommendationsClientGetGenerateStatusOptions contains the optional parameters for the RecommendationsClient.GetGenerateStatus
// method.
func (client *RecommendationsClient) GetGenerateStatus(ctx context.Context, operationID string, options *RecommendationsClientGetGenerateStatusOptions) (RecommendationsClientGetGenerateStatusResponse, error) {
	req, err := client.getGenerateStatusCreateRequest(ctx, operationID, options)
	if err != nil {
		return RecommendationsClientGetGenerateStatusResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return RecommendationsClientGetGenerateStatusResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusAccepted, http.StatusNoContent) {
		return RecommendationsClientGetGenerateStatusResponse{}, runtime.NewResponseError(resp)
	}
	return RecommendationsClientGetGenerateStatusResponse{}, nil
}

// getGenerateStatusCreateRequest creates the GetGenerateStatus request.
func (client *RecommendationsClient) getGenerateStatusCreateRequest(ctx context.Context, operationID string, options *RecommendationsClientGetGenerateStatusOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Advisor/generateRecommendations/{operationId}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	urlPath = strings.ReplaceAll(urlPath, "{operationId}", url.PathEscape(operationID))
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

// NewListPager - Obtains cached recommendations for a subscription. The recommendations are generated or computed by invoking
// generateRecommendations.
// If the operation fails it returns an *azcore.ResponseError type.
// options - RecommendationsClientListOptions contains the optional parameters for the RecommendationsClient.List method.
func (client *RecommendationsClient) NewListPager(options *RecommendationsClientListOptions) *runtime.Pager[RecommendationsClientListResponse] {
	return runtime.NewPager(runtime.PageProcessor[RecommendationsClientListResponse]{
		More: func(page RecommendationsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *RecommendationsClientListResponse) (RecommendationsClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return RecommendationsClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return RecommendationsClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return RecommendationsClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *RecommendationsClient) listCreateRequest(ctx context.Context, options *RecommendationsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Advisor/recommendations"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-01-01")
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	if options != nil && options.SkipToken != nil {
		reqQP.Set("$skipToken", *options.SkipToken)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *RecommendationsClient) listHandleResponse(resp *http.Response) (RecommendationsClientListResponse, error) {
	result := RecommendationsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ResourceRecommendationBaseListResult); err != nil {
		return RecommendationsClientListResponse{}, err
	}
	return result, nil
}
