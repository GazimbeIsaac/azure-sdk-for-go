//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armiotsecurity

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

// DefenderSettingsClient contains the methods for the DefenderSettings group.
// Don't use this type directly, use NewDefenderSettingsClient() instead.
type DefenderSettingsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewDefenderSettingsClient creates a new instance of DefenderSettingsClient with the specified values.
// subscriptionID - The ID of the target subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewDefenderSettingsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*DefenderSettingsClient, error) {
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
	client := &DefenderSettingsClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// CreateOrUpdate - Create or update IoT Defender settings
// If the operation fails it returns an *azcore.ResponseError type.
// defenderSettingsModel - The IoT defender settings model
// options - DefenderSettingsClientCreateOrUpdateOptions contains the optional parameters for the DefenderSettingsClient.CreateOrUpdate
// method.
func (client *DefenderSettingsClient) CreateOrUpdate(ctx context.Context, defenderSettingsModel DefenderSettingsModel, options *DefenderSettingsClientCreateOrUpdateOptions) (DefenderSettingsClientCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, defenderSettingsModel, options)
	if err != nil {
		return DefenderSettingsClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return DefenderSettingsClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return DefenderSettingsClientCreateOrUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *DefenderSettingsClient) createOrUpdateCreateRequest(ctx context.Context, defenderSettingsModel DefenderSettingsModel, options *DefenderSettingsClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.IoTSecurity/defenderSettings/default"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-02-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, defenderSettingsModel)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *DefenderSettingsClient) createOrUpdateHandleResponse(resp *http.Response) (DefenderSettingsClientCreateOrUpdateResponse, error) {
	result := DefenderSettingsClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DefenderSettingsModel); err != nil {
		return DefenderSettingsClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Delete IoT Defender settings
// If the operation fails it returns an *azcore.ResponseError type.
// options - DefenderSettingsClientDeleteOptions contains the optional parameters for the DefenderSettingsClient.Delete method.
func (client *DefenderSettingsClient) Delete(ctx context.Context, options *DefenderSettingsClientDeleteOptions) (DefenderSettingsClientDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, options)
	if err != nil {
		return DefenderSettingsClientDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return DefenderSettingsClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return DefenderSettingsClientDeleteResponse{}, runtime.NewResponseError(resp)
	}
	return DefenderSettingsClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *DefenderSettingsClient) deleteCreateRequest(ctx context.Context, options *DefenderSettingsClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.IoTSecurity/defenderSettings/default"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-02-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// DownloadManagerActivation - Download manager activation data defined for this subscription
// If the operation fails it returns an *azcore.ResponseError type.
// options - DefenderSettingsClientDownloadManagerActivationOptions contains the optional parameters for the DefenderSettingsClient.DownloadManagerActivation
// method.
func (client *DefenderSettingsClient) DownloadManagerActivation(ctx context.Context, options *DefenderSettingsClientDownloadManagerActivationOptions) (DefenderSettingsClientDownloadManagerActivationResponse, error) {
	req, err := client.downloadManagerActivationCreateRequest(ctx, options)
	if err != nil {
		return DefenderSettingsClientDownloadManagerActivationResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return DefenderSettingsClientDownloadManagerActivationResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return DefenderSettingsClientDownloadManagerActivationResponse{}, runtime.NewResponseError(resp)
	}
	return DefenderSettingsClientDownloadManagerActivationResponse{Body: resp.Body}, nil
}

// downloadManagerActivationCreateRequest creates the DownloadManagerActivation request.
func (client *DefenderSettingsClient) downloadManagerActivationCreateRequest(ctx context.Context, options *DefenderSettingsClientDownloadManagerActivationOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.IoTSecurity/defenderSettings/default/downloadManagerActivation"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-02-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	runtime.SkipBodyDownload(req)
	req.Raw().Header.Set("Accept", "application/zip")
	return req, nil
}

// Get - Get IoT Defender Settings
// If the operation fails it returns an *azcore.ResponseError type.
// options - DefenderSettingsClientGetOptions contains the optional parameters for the DefenderSettingsClient.Get method.
func (client *DefenderSettingsClient) Get(ctx context.Context, options *DefenderSettingsClientGetOptions) (DefenderSettingsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, options)
	if err != nil {
		return DefenderSettingsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return DefenderSettingsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return DefenderSettingsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *DefenderSettingsClient) getCreateRequest(ctx context.Context, options *DefenderSettingsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.IoTSecurity/defenderSettings/default"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-02-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *DefenderSettingsClient) getHandleResponse(resp *http.Response) (DefenderSettingsClientGetResponse, error) {
	result := DefenderSettingsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DefenderSettingsModel); err != nil {
		return DefenderSettingsClientGetResponse{}, err
	}
	return result, nil
}

// List - List IoT Defender Settings
// If the operation fails it returns an *azcore.ResponseError type.
// options - DefenderSettingsClientListOptions contains the optional parameters for the DefenderSettingsClient.List method.
func (client *DefenderSettingsClient) List(ctx context.Context, options *DefenderSettingsClientListOptions) (DefenderSettingsClientListResponse, error) {
	req, err := client.listCreateRequest(ctx, options)
	if err != nil {
		return DefenderSettingsClientListResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return DefenderSettingsClientListResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return DefenderSettingsClientListResponse{}, runtime.NewResponseError(resp)
	}
	return client.listHandleResponse(resp)
}

// listCreateRequest creates the List request.
func (client *DefenderSettingsClient) listCreateRequest(ctx context.Context, options *DefenderSettingsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.IoTSecurity/defenderSettings"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-02-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *DefenderSettingsClient) listHandleResponse(resp *http.Response) (DefenderSettingsClientListResponse, error) {
	result := DefenderSettingsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DefenderSettingsList); err != nil {
		return DefenderSettingsClientListResponse{}, err
	}
	return result, nil
}

// PackageDownloads - Information about downloadable packages
// If the operation fails it returns an *azcore.ResponseError type.
// options - DefenderSettingsClientPackageDownloadsOptions contains the optional parameters for the DefenderSettingsClient.PackageDownloads
// method.
func (client *DefenderSettingsClient) PackageDownloads(ctx context.Context, options *DefenderSettingsClientPackageDownloadsOptions) (DefenderSettingsClientPackageDownloadsResponse, error) {
	req, err := client.packageDownloadsCreateRequest(ctx, options)
	if err != nil {
		return DefenderSettingsClientPackageDownloadsResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return DefenderSettingsClientPackageDownloadsResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return DefenderSettingsClientPackageDownloadsResponse{}, runtime.NewResponseError(resp)
	}
	return client.packageDownloadsHandleResponse(resp)
}

// packageDownloadsCreateRequest creates the PackageDownloads request.
func (client *DefenderSettingsClient) packageDownloadsCreateRequest(ctx context.Context, options *DefenderSettingsClientPackageDownloadsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.IoTSecurity/defenderSettings/default/packageDownloads"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-02-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// packageDownloadsHandleResponse handles the PackageDownloads response.
func (client *DefenderSettingsClient) packageDownloadsHandleResponse(resp *http.Response) (DefenderSettingsClientPackageDownloadsResponse, error) {
	result := DefenderSettingsClientPackageDownloadsResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PackageDownloads); err != nil {
		return DefenderSettingsClientPackageDownloadsResponse{}, err
	}
	return result, nil
}
