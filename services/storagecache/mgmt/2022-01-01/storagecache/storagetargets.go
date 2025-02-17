package storagecache

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

// StorageTargetsClient is the a Storage Cache provides scalable caching service for NAS clients, serving data from
// either NFSv3 or Blob at-rest storage (referred to as "Storage Targets"). These operations allow you to manage
// Caches.
type StorageTargetsClient struct {
	BaseClient
}

// NewStorageTargetsClient creates an instance of the StorageTargetsClient client.
func NewStorageTargetsClient(subscriptionID string) StorageTargetsClient {
	return NewStorageTargetsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewStorageTargetsClientWithBaseURI creates an instance of the StorageTargetsClient client using a custom endpoint.
// Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewStorageTargetsClientWithBaseURI(baseURI string, subscriptionID string) StorageTargetsClient {
	return StorageTargetsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate create or update a Storage Target. This operation is allowed at any time, but if the Cache is down or
// unhealthy, the actual creation/modification of the Storage Target may be delayed until the Cache is healthy again.
// Parameters:
// resourceGroupName - target resource group.
// cacheName - name of Cache. Length of name must not be greater than 80 and chars must be from the
// [-0-9a-zA-Z_] char class.
// storageTargetName - name of Storage Target.
// storagetarget - object containing the definition of a Storage Target.
func (client StorageTargetsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, cacheName string, storageTargetName string, storagetarget *StorageTarget) (result StorageTargetsCreateOrUpdateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/StorageTargetsClient.CreateOrUpdate")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: cacheName,
			Constraints: []validation.Constraint{{Target: "cacheName", Name: validation.Pattern, Rule: `^[-0-9a-zA-Z_]{1,80}$`, Chain: nil}}},
		{TargetValue: storageTargetName,
			Constraints: []validation.Constraint{{Target: "storageTargetName", Name: validation.Pattern, Rule: `^[-0-9a-zA-Z_]{1,80}$`, Chain: nil}}},
		{TargetValue: storagetarget,
			Constraints: []validation.Constraint{{Target: "storagetarget", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "storagetarget.StorageTargetProperties", Name: validation.Null, Rule: false,
					Chain: []validation.Constraint{{Target: "storagetarget.StorageTargetProperties.Nfs3", Name: validation.Null, Rule: false,
						Chain: []validation.Constraint{{Target: "storagetarget.StorageTargetProperties.Nfs3.Target", Name: validation.Null, Rule: false,
							Chain: []validation.Constraint{{Target: "storagetarget.StorageTargetProperties.Nfs3.Target", Name: validation.Pattern, Rule: `^[-.,0-9a-zA-Z]+$`, Chain: nil}}},
						}},
					}},
				}}}}}); err != nil {
		return result, validation.NewError("storagecache.StorageTargetsClient", "CreateOrUpdate", err.Error())
	}

	req, err := client.CreateOrUpdatePreparer(ctx, resourceGroupName, cacheName, storageTargetName, storagetarget)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storagecache.StorageTargetsClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	result, err = client.CreateOrUpdateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storagecache.StorageTargetsClient", "CreateOrUpdate", result.Response(), "Failure sending request")
		return
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client StorageTargetsClient) CreateOrUpdatePreparer(ctx context.Context, resourceGroupName string, cacheName string, storageTargetName string, storagetarget *StorageTarget) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"cacheName":         autorest.Encode("path", cacheName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"storageTargetName": autorest.Encode("path", storageTargetName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2022-01-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.StorageCache/caches/{cacheName}/storageTargets/{storageTargetName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	if storagetarget != nil {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithJSON(storagetarget))
	}
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client StorageTargetsClient) CreateOrUpdateSender(req *http.Request) (future StorageTargetsCreateOrUpdateFuture, err error) {
	var resp *http.Response
	future.FutureAPI = &azure.Future{}
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = future.result
	return
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client StorageTargetsClient) CreateOrUpdateResponder(resp *http.Response) (result StorageTarget, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete removes a Storage Target from a Cache. This operation is allowed at any time, but if the Cache is down or
// unhealthy, the actual removal of the Storage Target may be delayed until the Cache is healthy again. Note that if
// the Cache has data to flush to the Storage Target, the data will be flushed before the Storage Target will be
// deleted.
// Parameters:
// resourceGroupName - target resource group.
// cacheName - name of Cache. Length of name must not be greater than 80 and chars must be from the
// [-0-9a-zA-Z_] char class.
// storageTargetName - name of Storage Target.
// force - boolean value requesting the force delete operation for a storage target. Force delete discards
// unwritten-data in the cache instead of flushing it to back-end storage.
func (client StorageTargetsClient) Delete(ctx context.Context, resourceGroupName string, cacheName string, storageTargetName string, force string) (result StorageTargetsDeleteFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/StorageTargetsClient.Delete")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: cacheName,
			Constraints: []validation.Constraint{{Target: "cacheName", Name: validation.Pattern, Rule: `^[-0-9a-zA-Z_]{1,80}$`, Chain: nil}}},
		{TargetValue: storageTargetName,
			Constraints: []validation.Constraint{{Target: "storageTargetName", Name: validation.Pattern, Rule: `^[-0-9a-zA-Z_]{1,80}$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("storagecache.StorageTargetsClient", "Delete", err.Error())
	}

	req, err := client.DeletePreparer(ctx, resourceGroupName, cacheName, storageTargetName, force)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storagecache.StorageTargetsClient", "Delete", nil, "Failure preparing request")
		return
	}

	result, err = client.DeleteSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storagecache.StorageTargetsClient", "Delete", result.Response(), "Failure sending request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client StorageTargetsClient) DeletePreparer(ctx context.Context, resourceGroupName string, cacheName string, storageTargetName string, force string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"cacheName":         autorest.Encode("path", cacheName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"storageTargetName": autorest.Encode("path", storageTargetName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2022-01-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(force) > 0 {
		queryParameters["force"] = autorest.Encode("query", force)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.StorageCache/caches/{cacheName}/storageTargets/{storageTargetName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client StorageTargetsClient) DeleteSender(req *http.Request) (future StorageTargetsDeleteFuture, err error) {
	var resp *http.Response
	future.FutureAPI = &azure.Future{}
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = future.result
	return
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client StorageTargetsClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// DNSRefresh tells a storage target to refresh its DNS information.
// Parameters:
// resourceGroupName - target resource group.
// cacheName - name of Cache. Length of name must not be greater than 80 and chars must be from the
// [-0-9a-zA-Z_] char class.
// storageTargetName - name of Storage Target.
func (client StorageTargetsClient) DNSRefresh(ctx context.Context, resourceGroupName string, cacheName string, storageTargetName string) (result StorageTargetsDNSRefreshFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/StorageTargetsClient.DNSRefresh")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: cacheName,
			Constraints: []validation.Constraint{{Target: "cacheName", Name: validation.Pattern, Rule: `^[-0-9a-zA-Z_]{1,80}$`, Chain: nil}}},
		{TargetValue: storageTargetName,
			Constraints: []validation.Constraint{{Target: "storageTargetName", Name: validation.Pattern, Rule: `^[-0-9a-zA-Z_]{1,80}$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("storagecache.StorageTargetsClient", "DNSRefresh", err.Error())
	}

	req, err := client.DNSRefreshPreparer(ctx, resourceGroupName, cacheName, storageTargetName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storagecache.StorageTargetsClient", "DNSRefresh", nil, "Failure preparing request")
		return
	}

	result, err = client.DNSRefreshSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storagecache.StorageTargetsClient", "DNSRefresh", result.Response(), "Failure sending request")
		return
	}

	return
}

// DNSRefreshPreparer prepares the DNSRefresh request.
func (client StorageTargetsClient) DNSRefreshPreparer(ctx context.Context, resourceGroupName string, cacheName string, storageTargetName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"cacheName":         autorest.Encode("path", cacheName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"storageTargetName": autorest.Encode("path", storageTargetName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2022-01-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.StorageCache/caches/{cacheName}/storageTargets/{storageTargetName}/dnsRefresh", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DNSRefreshSender sends the DNSRefresh request. The method will close the
// http.Response Body if it receives an error.
func (client StorageTargetsClient) DNSRefreshSender(req *http.Request) (future StorageTargetsDNSRefreshFuture, err error) {
	var resp *http.Response
	future.FutureAPI = &azure.Future{}
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = future.result
	return
}

// DNSRefreshResponder handles the response to the DNSRefresh request. The method always
// closes the http.Response Body.
func (client StorageTargetsClient) DNSRefreshResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get returns a Storage Target from a Cache.
// Parameters:
// resourceGroupName - target resource group.
// cacheName - name of Cache. Length of name must not be greater than 80 and chars must be from the
// [-0-9a-zA-Z_] char class.
// storageTargetName - name of Storage Target.
func (client StorageTargetsClient) Get(ctx context.Context, resourceGroupName string, cacheName string, storageTargetName string) (result StorageTarget, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/StorageTargetsClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: cacheName,
			Constraints: []validation.Constraint{{Target: "cacheName", Name: validation.Pattern, Rule: `^[-0-9a-zA-Z_]{1,80}$`, Chain: nil}}},
		{TargetValue: storageTargetName,
			Constraints: []validation.Constraint{{Target: "storageTargetName", Name: validation.Pattern, Rule: `^[-0-9a-zA-Z_]{1,80}$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("storagecache.StorageTargetsClient", "Get", err.Error())
	}

	req, err := client.GetPreparer(ctx, resourceGroupName, cacheName, storageTargetName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storagecache.StorageTargetsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "storagecache.StorageTargetsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storagecache.StorageTargetsClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client StorageTargetsClient) GetPreparer(ctx context.Context, resourceGroupName string, cacheName string, storageTargetName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"cacheName":         autorest.Encode("path", cacheName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"storageTargetName": autorest.Encode("path", storageTargetName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2022-01-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.StorageCache/caches/{cacheName}/storageTargets/{storageTargetName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client StorageTargetsClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client StorageTargetsClient) GetResponder(resp *http.Response) (result StorageTarget, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByCache returns a list of Storage Targets for the specified Cache.
// Parameters:
// resourceGroupName - target resource group.
// cacheName - name of Cache. Length of name must not be greater than 80 and chars must be from the
// [-0-9a-zA-Z_] char class.
func (client StorageTargetsClient) ListByCache(ctx context.Context, resourceGroupName string, cacheName string) (result StorageTargetsResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/StorageTargetsClient.ListByCache")
		defer func() {
			sc := -1
			if result.str.Response.Response != nil {
				sc = result.str.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: cacheName,
			Constraints: []validation.Constraint{{Target: "cacheName", Name: validation.Pattern, Rule: `^[-0-9a-zA-Z_]{1,80}$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("storagecache.StorageTargetsClient", "ListByCache", err.Error())
	}

	result.fn = client.listByCacheNextResults
	req, err := client.ListByCachePreparer(ctx, resourceGroupName, cacheName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storagecache.StorageTargetsClient", "ListByCache", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByCacheSender(req)
	if err != nil {
		result.str.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "storagecache.StorageTargetsClient", "ListByCache", resp, "Failure sending request")
		return
	}

	result.str, err = client.ListByCacheResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storagecache.StorageTargetsClient", "ListByCache", resp, "Failure responding to request")
		return
	}
	if result.str.hasNextLink() && result.str.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListByCachePreparer prepares the ListByCache request.
func (client StorageTargetsClient) ListByCachePreparer(ctx context.Context, resourceGroupName string, cacheName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"cacheName":         autorest.Encode("path", cacheName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2022-01-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.StorageCache/caches/{cacheName}/storageTargets", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByCacheSender sends the ListByCache request. The method will close the
// http.Response Body if it receives an error.
func (client StorageTargetsClient) ListByCacheSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListByCacheResponder handles the response to the ListByCache request. The method always
// closes the http.Response Body.
func (client StorageTargetsClient) ListByCacheResponder(resp *http.Response) (result StorageTargetsResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByCacheNextResults retrieves the next set of results, if any.
func (client StorageTargetsClient) listByCacheNextResults(ctx context.Context, lastResults StorageTargetsResult) (result StorageTargetsResult, err error) {
	req, err := lastResults.storageTargetsResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "storagecache.StorageTargetsClient", "listByCacheNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByCacheSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "storagecache.StorageTargetsClient", "listByCacheNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByCacheResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storagecache.StorageTargetsClient", "listByCacheNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByCacheComplete enumerates all values, automatically crossing page boundaries as required.
func (client StorageTargetsClient) ListByCacheComplete(ctx context.Context, resourceGroupName string, cacheName string) (result StorageTargetsResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/StorageTargetsClient.ListByCache")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListByCache(ctx, resourceGroupName, cacheName)
	return
}
