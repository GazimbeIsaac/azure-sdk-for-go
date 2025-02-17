//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armsecurity

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

// AccountConnectorsClient contains the methods for the AccountConnectors group.
// Don't use this type directly, use NewAccountConnectorsClient() instead.
type AccountConnectorsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewAccountConnectorsClient creates a new instance of AccountConnectorsClient with the specified values.
// subscriptionID - Azure subscription ID
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewAccountConnectorsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*AccountConnectorsClient, error) {
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
	client := &AccountConnectorsClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// CreateOrUpdate - Create a cloud account connector or update an existing one. Connect to your cloud account. For AWS, use
// either account credentials or role-based authentication. For GCP, use account organization
// credentials.
// If the operation fails it returns an *azcore.ResponseError type.
// connectorName - Name of the cloud account connector
// connectorSetting - Settings for the cloud account connector
// options - AccountConnectorsClientCreateOrUpdateOptions contains the optional parameters for the AccountConnectorsClient.CreateOrUpdate
// method.
func (client *AccountConnectorsClient) CreateOrUpdate(ctx context.Context, connectorName string, connectorSetting ConnectorSetting, options *AccountConnectorsClientCreateOrUpdateOptions) (AccountConnectorsClientCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, connectorName, connectorSetting, options)
	if err != nil {
		return AccountConnectorsClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return AccountConnectorsClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return AccountConnectorsClientCreateOrUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *AccountConnectorsClient) createOrUpdateCreateRequest(ctx context.Context, connectorName string, connectorSetting ConnectorSetting, options *AccountConnectorsClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Security/connectors/{connectorName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if connectorName == "" {
		return nil, errors.New("parameter connectorName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{connectorName}", url.PathEscape(connectorName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-01-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, connectorSetting)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *AccountConnectorsClient) createOrUpdateHandleResponse(resp *http.Response) (AccountConnectorsClientCreateOrUpdateResponse, error) {
	result := AccountConnectorsClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ConnectorSetting); err != nil {
		return AccountConnectorsClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Delete a cloud account connector from a subscription
// If the operation fails it returns an *azcore.ResponseError type.
// connectorName - Name of the cloud account connector
// options - AccountConnectorsClientDeleteOptions contains the optional parameters for the AccountConnectorsClient.Delete
// method.
func (client *AccountConnectorsClient) Delete(ctx context.Context, connectorName string, options *AccountConnectorsClientDeleteOptions) (AccountConnectorsClientDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, connectorName, options)
	if err != nil {
		return AccountConnectorsClientDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return AccountConnectorsClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return AccountConnectorsClientDeleteResponse{}, runtime.NewResponseError(resp)
	}
	return AccountConnectorsClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *AccountConnectorsClient) deleteCreateRequest(ctx context.Context, connectorName string, options *AccountConnectorsClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Security/connectors/{connectorName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if connectorName == "" {
		return nil, errors.New("parameter connectorName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{connectorName}", url.PathEscape(connectorName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-01-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// Get - Details of a specific cloud account connector
// If the operation fails it returns an *azcore.ResponseError type.
// connectorName - Name of the cloud account connector
// options - AccountConnectorsClientGetOptions contains the optional parameters for the AccountConnectorsClient.Get method.
func (client *AccountConnectorsClient) Get(ctx context.Context, connectorName string, options *AccountConnectorsClientGetOptions) (AccountConnectorsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, connectorName, options)
	if err != nil {
		return AccountConnectorsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return AccountConnectorsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return AccountConnectorsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *AccountConnectorsClient) getCreateRequest(ctx context.Context, connectorName string, options *AccountConnectorsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Security/connectors/{connectorName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if connectorName == "" {
		return nil, errors.New("parameter connectorName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{connectorName}", url.PathEscape(connectorName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-01-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *AccountConnectorsClient) getHandleResponse(resp *http.Response) (AccountConnectorsClientGetResponse, error) {
	result := AccountConnectorsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ConnectorSetting); err != nil {
		return AccountConnectorsClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Cloud accounts connectors of a subscription
// If the operation fails it returns an *azcore.ResponseError type.
// options - AccountConnectorsClientListOptions contains the optional parameters for the AccountConnectorsClient.List method.
func (client *AccountConnectorsClient) NewListPager(options *AccountConnectorsClientListOptions) *runtime.Pager[AccountConnectorsClientListResponse] {
	return runtime.NewPager(runtime.PageProcessor[AccountConnectorsClientListResponse]{
		More: func(page AccountConnectorsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *AccountConnectorsClientListResponse) (AccountConnectorsClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return AccountConnectorsClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return AccountConnectorsClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return AccountConnectorsClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *AccountConnectorsClient) listCreateRequest(ctx context.Context, options *AccountConnectorsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Security/connectors"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-01-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *AccountConnectorsClient) listHandleResponse(resp *http.Response) (AccountConnectorsClientListResponse, error) {
	result := AccountConnectorsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ConnectorSettingList); err != nil {
		return AccountConnectorsClientListResponse{}, err
	}
	return result, nil
}
