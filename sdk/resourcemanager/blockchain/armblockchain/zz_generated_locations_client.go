//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armblockchain

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

// LocationsClient contains the methods for the Locations group.
// Don't use this type directly, use NewLocationsClient() instead.
type LocationsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewLocationsClient creates a new instance of LocationsClient with the specified values.
// subscriptionID - Gets the subscription Id which uniquely identifies the Microsoft Azure subscription. The subscription
// ID is part of the URI for every service call.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewLocationsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*LocationsClient, error) {
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
	client := &LocationsClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// CheckNameAvailability - To check whether a resource name is available.
// If the operation fails it returns an *azcore.ResponseError type.
// locationName - Location Name.
// options - LocationsClientCheckNameAvailabilityOptions contains the optional parameters for the LocationsClient.CheckNameAvailability
// method.
func (client *LocationsClient) CheckNameAvailability(ctx context.Context, locationName string, options *LocationsClientCheckNameAvailabilityOptions) (LocationsClientCheckNameAvailabilityResponse, error) {
	req, err := client.checkNameAvailabilityCreateRequest(ctx, locationName, options)
	if err != nil {
		return LocationsClientCheckNameAvailabilityResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return LocationsClientCheckNameAvailabilityResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return LocationsClientCheckNameAvailabilityResponse{}, runtime.NewResponseError(resp)
	}
	return client.checkNameAvailabilityHandleResponse(resp)
}

// checkNameAvailabilityCreateRequest creates the CheckNameAvailability request.
func (client *LocationsClient) checkNameAvailabilityCreateRequest(ctx context.Context, locationName string, options *LocationsClientCheckNameAvailabilityOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Blockchain/locations/{locationName}/checkNameAvailability"
	if locationName == "" {
		return nil, errors.New("parameter locationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{locationName}", url.PathEscape(locationName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	if options != nil && options.NameAvailabilityRequest != nil {
		return req, runtime.MarshalAsJSON(req, *options.NameAvailabilityRequest)
	}
	return req, nil
}

// checkNameAvailabilityHandleResponse handles the CheckNameAvailability response.
func (client *LocationsClient) checkNameAvailabilityHandleResponse(resp *http.Response) (LocationsClientCheckNameAvailabilityResponse, error) {
	result := LocationsClientCheckNameAvailabilityResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.NameAvailability); err != nil {
		return LocationsClientCheckNameAvailabilityResponse{}, err
	}
	return result, nil
}

// ListConsortiums - Lists the available consortiums for a subscription.
// If the operation fails it returns an *azcore.ResponseError type.
// locationName - Location Name.
// options - LocationsClientListConsortiumsOptions contains the optional parameters for the LocationsClient.ListConsortiums
// method.
func (client *LocationsClient) ListConsortiums(ctx context.Context, locationName string, options *LocationsClientListConsortiumsOptions) (LocationsClientListConsortiumsResponse, error) {
	req, err := client.listConsortiumsCreateRequest(ctx, locationName, options)
	if err != nil {
		return LocationsClientListConsortiumsResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return LocationsClientListConsortiumsResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return LocationsClientListConsortiumsResponse{}, runtime.NewResponseError(resp)
	}
	return client.listConsortiumsHandleResponse(resp)
}

// listConsortiumsCreateRequest creates the ListConsortiums request.
func (client *LocationsClient) listConsortiumsCreateRequest(ctx context.Context, locationName string, options *LocationsClientListConsortiumsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Blockchain/locations/{locationName}/listConsortiums"
	if locationName == "" {
		return nil, errors.New("parameter locationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{locationName}", url.PathEscape(locationName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listConsortiumsHandleResponse handles the ListConsortiums response.
func (client *LocationsClient) listConsortiumsHandleResponse(resp *http.Response) (LocationsClientListConsortiumsResponse, error) {
	result := LocationsClientListConsortiumsResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ConsortiumCollection); err != nil {
		return LocationsClientListConsortiumsResponse{}, err
	}
	return result, nil
}
