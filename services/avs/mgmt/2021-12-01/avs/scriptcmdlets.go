package avs

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/validation"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// ScriptCmdletsClient is the azure VMware Solution API
type ScriptCmdletsClient struct {
	BaseClient
}

// NewScriptCmdletsClient creates an instance of the ScriptCmdletsClient client.
func NewScriptCmdletsClient(subscriptionID string) ScriptCmdletsClient {
	return NewScriptCmdletsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewScriptCmdletsClientWithBaseURI creates an instance of the ScriptCmdletsClient client using a custom endpoint.
// Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewScriptCmdletsClientWithBaseURI(baseURI string, subscriptionID string) ScriptCmdletsClient {
	return ScriptCmdletsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Get return information about a script cmdlet resource in a specific package on a private cloud
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// privateCloudName - name of the private cloud
// scriptPackageName - name of the script package in the private cloud
// scriptCmdletName - name of the script cmdlet resource in the script package in the private cloud
func (client ScriptCmdletsClient) Get(ctx context.Context, resourceGroupName string, privateCloudName string, scriptPackageName string, scriptCmdletName string) (result ScriptCmdlet, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ScriptCmdletsClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil}}}}); err != nil {
		return result, validation.NewError("avs.ScriptCmdletsClient", "Get", err.Error())
	}

	req, err := client.GetPreparer(ctx, resourceGroupName, privateCloudName, scriptPackageName, scriptCmdletName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "avs.ScriptCmdletsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "avs.ScriptCmdletsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "avs.ScriptCmdletsClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client ScriptCmdletsClient) GetPreparer(ctx context.Context, resourceGroupName string, privateCloudName string, scriptPackageName string, scriptCmdletName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"privateCloudName":  autorest.Encode("path", privateCloudName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"scriptCmdletName":  autorest.Encode("path", scriptCmdletName),
		"scriptPackageName": autorest.Encode("path", scriptPackageName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2021-12-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AVS/privateClouds/{privateCloudName}/scriptPackages/{scriptPackageName}/scriptCmdlets/{scriptCmdletName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client ScriptCmdletsClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client ScriptCmdletsClient) GetResponder(resp *http.Response) (result ScriptCmdlet, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List list script cmdlet resources available for a private cloud to create a script execution resource on a private
// cloud
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// privateCloudName - name of the private cloud
// scriptPackageName - name of the script package in the private cloud
func (client ScriptCmdletsClient) List(ctx context.Context, resourceGroupName string, privateCloudName string, scriptPackageName string) (result ScriptCmdletsListPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ScriptCmdletsClient.List")
		defer func() {
			sc := -1
			if result.scl.Response.Response != nil {
				sc = result.scl.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil}}}}); err != nil {
		return result, validation.NewError("avs.ScriptCmdletsClient", "List", err.Error())
	}

	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx, resourceGroupName, privateCloudName, scriptPackageName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "avs.ScriptCmdletsClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.scl.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "avs.ScriptCmdletsClient", "List", resp, "Failure sending request")
		return
	}

	result.scl, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "avs.ScriptCmdletsClient", "List", resp, "Failure responding to request")
		return
	}
	if result.scl.hasNextLink() && result.scl.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListPreparer prepares the List request.
func (client ScriptCmdletsClient) ListPreparer(ctx context.Context, resourceGroupName string, privateCloudName string, scriptPackageName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"privateCloudName":  autorest.Encode("path", privateCloudName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"scriptPackageName": autorest.Encode("path", scriptPackageName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2021-12-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AVS/privateClouds/{privateCloudName}/scriptPackages/{scriptPackageName}/scriptCmdlets", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client ScriptCmdletsClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client ScriptCmdletsClient) ListResponder(resp *http.Response) (result ScriptCmdletsList, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client ScriptCmdletsClient) listNextResults(ctx context.Context, lastResults ScriptCmdletsList) (result ScriptCmdletsList, err error) {
	req, err := lastResults.scriptCmdletsListPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "avs.ScriptCmdletsClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "avs.ScriptCmdletsClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "avs.ScriptCmdletsClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client ScriptCmdletsClient) ListComplete(ctx context.Context, resourceGroupName string, privateCloudName string, scriptPackageName string) (result ScriptCmdletsListIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ScriptCmdletsClient.List")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.List(ctx, resourceGroupName, privateCloudName, scriptPackageName)
	return
}
