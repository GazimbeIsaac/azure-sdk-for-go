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

// MemberOperationResultsClient contains the methods for the BlockchainMemberOperationResults group.
// Don't use this type directly, use NewMemberOperationResultsClient() instead.
type MemberOperationResultsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewMemberOperationResultsClient creates a new instance of MemberOperationResultsClient with the specified values.
// subscriptionID - Gets the subscription Id which uniquely identifies the Microsoft Azure subscription. The subscription
// ID is part of the URI for every service call.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewMemberOperationResultsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*MemberOperationResultsClient, error) {
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
	client := &MemberOperationResultsClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// Get - Get Async operation result.
// If the operation fails it returns an *azcore.ResponseError type.
// locationName - Location name.
// operationID - Operation Id.
// options - MemberOperationResultsClientGetOptions contains the optional parameters for the MemberOperationResultsClient.Get
// method.
func (client *MemberOperationResultsClient) Get(ctx context.Context, locationName string, operationID string, options *MemberOperationResultsClientGetOptions) (MemberOperationResultsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, locationName, operationID, options)
	if err != nil {
		return MemberOperationResultsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return MemberOperationResultsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return MemberOperationResultsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *MemberOperationResultsClient) getCreateRequest(ctx context.Context, locationName string, operationID string, options *MemberOperationResultsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Blockchain/locations/{locationName}/blockchainMemberOperationResults/{operationId}"
	if locationName == "" {
		return nil, errors.New("parameter locationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{locationName}", url.PathEscape(locationName))
	if operationID == "" {
		return nil, errors.New("parameter operationID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{operationId}", url.PathEscape(operationID))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *MemberOperationResultsClient) getHandleResponse(resp *http.Response) (MemberOperationResultsClientGetResponse, error) {
	result := MemberOperationResultsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.OperationResult); err != nil {
		return MemberOperationResultsClientGetResponse{}, err
	}
	return result, nil
}
