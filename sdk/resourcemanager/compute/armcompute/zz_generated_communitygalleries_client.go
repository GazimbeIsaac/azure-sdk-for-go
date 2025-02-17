//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcompute

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

// CommunityGalleriesClient contains the methods for the CommunityGalleries group.
// Don't use this type directly, use NewCommunityGalleriesClient() instead.
type CommunityGalleriesClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewCommunityGalleriesClient creates a new instance of CommunityGalleriesClient with the specified values.
// subscriptionID - Subscription credentials which uniquely identify Microsoft Azure subscription. The subscription ID forms
// part of the URI for every service call.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewCommunityGalleriesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*CommunityGalleriesClient, error) {
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
	client := &CommunityGalleriesClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// Get - Get a community gallery by gallery public name.
// If the operation fails it returns an *azcore.ResponseError type.
// location - Resource location.
// publicGalleryName - The public name of the community gallery.
// options - CommunityGalleriesClientGetOptions contains the optional parameters for the CommunityGalleriesClient.Get method.
func (client *CommunityGalleriesClient) Get(ctx context.Context, location string, publicGalleryName string, options *CommunityGalleriesClientGetOptions) (CommunityGalleriesClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, location, publicGalleryName, options)
	if err != nil {
		return CommunityGalleriesClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return CommunityGalleriesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return CommunityGalleriesClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *CommunityGalleriesClient) getCreateRequest(ctx context.Context, location string, publicGalleryName string, options *CommunityGalleriesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Compute/locations/{location}/communityGalleries/{publicGalleryName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if location == "" {
		return nil, errors.New("parameter location cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{location}", url.PathEscape(location))
	if publicGalleryName == "" {
		return nil, errors.New("parameter publicGalleryName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{publicGalleryName}", url.PathEscape(publicGalleryName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-07-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *CommunityGalleriesClient) getHandleResponse(resp *http.Response) (CommunityGalleriesClientGetResponse, error) {
	result := CommunityGalleriesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.CommunityGallery); err != nil {
		return CommunityGalleriesClientGetResponse{}, err
	}
	return result, nil
}
