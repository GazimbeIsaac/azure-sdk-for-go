package managementgroups

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// Client is the the Azure Management Groups API enables consolidation of multiple
// subscriptions/resources into an organizational hierarchy and centrally
// manage access control, policies, alerting and reporting for those resources.
type Client struct {
	BaseClient
}

// NewClient creates an instance of the Client client.
func NewClient() Client {
	return NewClientWithBaseURI(DefaultBaseURI)
}

// NewClientWithBaseURI creates an instance of the Client client using a custom endpoint.  Use this when interacting
// with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewClientWithBaseURI(baseURI string) Client {
	return Client{NewWithBaseURI(baseURI)}
}

// CreateOrUpdate create or update a management group.
// If a management group is already created and a subsequent create request is issued with different properties, the
// management group properties will be updated.
// Parameters:
// groupID - management Group ID.
// createManagementGroupRequest - management group creation parameters.
// cacheControl - indicates that the request shouldn't utilize any caches.
func (client Client) CreateOrUpdate(ctx context.Context, groupID string, createManagementGroupRequest CreateManagementGroupRequest, cacheControl string) (result CreateOrUpdateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/Client.CreateOrUpdate")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.CreateOrUpdatePreparer(ctx, groupID, createManagementGroupRequest, cacheControl)
	if err != nil {
		err = autorest.NewErrorWithError(err, "managementgroups.Client", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	result, err = client.CreateOrUpdateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "managementgroups.Client", "CreateOrUpdate", result.Response(), "Failure sending request")
		return
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client Client) CreateOrUpdatePreparer(ctx context.Context, groupID string, createManagementGroupRequest CreateManagementGroupRequest, cacheControl string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"groupId": autorest.Encode("path", groupID),
	}

	const APIVersion = "2020-02-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	createManagementGroupRequest.ID = nil
	createManagementGroupRequest.Type = nil
	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/providers/Microsoft.Management/managementGroups/{groupId}", pathParameters),
		autorest.WithJSON(createManagementGroupRequest),
		autorest.WithQueryParameters(queryParameters))
	if len(cacheControl) > 0 {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithHeader("Cache-Control", autorest.String(cacheControl)))
	} else {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithHeader("Cache-Control", autorest.String("no-cache")))
	}
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client Client) CreateOrUpdateSender(req *http.Request) (future CreateOrUpdateFuture, err error) {
	var resp *http.Response
	future.FutureAPI = &azure.Future{}
	resp, err = client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
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
func (client Client) CreateOrUpdateResponder(resp *http.Response) (result SetObject, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result.Value),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete delete management group.
// If a management group contains child resources, the request will fail.
// Parameters:
// groupID - management Group ID.
// cacheControl - indicates that the request shouldn't utilize any caches.
func (client Client) Delete(ctx context.Context, groupID string, cacheControl string) (result DeleteFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/Client.Delete")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeletePreparer(ctx, groupID, cacheControl)
	if err != nil {
		err = autorest.NewErrorWithError(err, "managementgroups.Client", "Delete", nil, "Failure preparing request")
		return
	}

	result, err = client.DeleteSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "managementgroups.Client", "Delete", result.Response(), "Failure sending request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client Client) DeletePreparer(ctx context.Context, groupID string, cacheControl string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"groupId": autorest.Encode("path", groupID),
	}

	const APIVersion = "2020-02-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/providers/Microsoft.Management/managementGroups/{groupId}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	if len(cacheControl) > 0 {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithHeader("Cache-Control", autorest.String(cacheControl)))
	} else {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithHeader("Cache-Control", autorest.String("no-cache")))
	}
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client Client) DeleteSender(req *http.Request) (future DeleteFuture, err error) {
	var resp *http.Response
	future.FutureAPI = &azure.Future{}
	resp, err = client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
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
func (client Client) DeleteResponder(resp *http.Response) (result AzureAsyncOperationResults, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Get get the details of the management group.
// Parameters:
// groupID - management Group ID.
// expand - the $expand=children query string parameter allows clients to request inclusion of children in the
// response payload.  $expand=path includes the path from the root group to the current group.
// recurse - the $recurse=true query string parameter allows clients to request inclusion of entire hierarchy
// in the response payload. Note that  $expand=children must be passed up if $recurse is set to true.
// filter - a filter which allows the exclusion of subscriptions from results (i.e. '$filter=children.childType
// ne Subscription')
// cacheControl - indicates that the request shouldn't utilize any caches.
func (client Client) Get(ctx context.Context, groupID string, expand string, recurse *bool, filter string, cacheControl string) (result ManagementGroup, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/Client.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, groupID, expand, recurse, filter, cacheControl)
	if err != nil {
		err = autorest.NewErrorWithError(err, "managementgroups.Client", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "managementgroups.Client", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "managementgroups.Client", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client Client) GetPreparer(ctx context.Context, groupID string, expand string, recurse *bool, filter string, cacheControl string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"groupId": autorest.Encode("path", groupID),
	}

	const APIVersion = "2020-02-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(string(expand)) > 0 {
		queryParameters["$expand"] = autorest.Encode("query", expand)
	}
	if recurse != nil {
		queryParameters["$recurse"] = autorest.Encode("query", *recurse)
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/providers/Microsoft.Management/managementGroups/{groupId}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	if len(cacheControl) > 0 {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithHeader("Cache-Control", autorest.String(cacheControl)))
	} else {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithHeader("Cache-Control", autorest.String("no-cache")))
	}
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client Client) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client Client) GetResponder(resp *http.Response) (result ManagementGroup, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetDescendants list all entities that descend from a management group.
// Parameters:
// groupID - management Group ID.
// skiptoken - page continuation token is only used if a previous operation returned a partial result.
// If a previous response contains a nextLink element, the value of the nextLink element will include a token
// parameter that specifies a starting point to use for subsequent calls.
// top - number of elements to return when retrieving results. Passing this in will override $skipToken.
func (client Client) GetDescendants(ctx context.Context, groupID string, skiptoken string, top *int32) (result DescendantListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/Client.GetDescendants")
		defer func() {
			sc := -1
			if result.dlr.Response.Response != nil {
				sc = result.dlr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.getDescendantsNextResults
	req, err := client.GetDescendantsPreparer(ctx, groupID, skiptoken, top)
	if err != nil {
		err = autorest.NewErrorWithError(err, "managementgroups.Client", "GetDescendants", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetDescendantsSender(req)
	if err != nil {
		result.dlr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "managementgroups.Client", "GetDescendants", resp, "Failure sending request")
		return
	}

	result.dlr, err = client.GetDescendantsResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "managementgroups.Client", "GetDescendants", resp, "Failure responding to request")
		return
	}
	if result.dlr.hasNextLink() && result.dlr.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// GetDescendantsPreparer prepares the GetDescendants request.
func (client Client) GetDescendantsPreparer(ctx context.Context, groupID string, skiptoken string, top *int32) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"groupId": autorest.Encode("path", groupID),
	}

	const APIVersion = "2020-02-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(skiptoken) > 0 {
		queryParameters["$skiptoken"] = autorest.Encode("query", skiptoken)
	}
	if top != nil {
		queryParameters["$top"] = autorest.Encode("query", *top)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/providers/Microsoft.Management/managementGroups/{groupId}/descendants", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetDescendantsSender sends the GetDescendants request. The method will close the
// http.Response Body if it receives an error.
func (client Client) GetDescendantsSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetDescendantsResponder handles the response to the GetDescendants request. The method always
// closes the http.Response Body.
func (client Client) GetDescendantsResponder(resp *http.Response) (result DescendantListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// getDescendantsNextResults retrieves the next set of results, if any.
func (client Client) getDescendantsNextResults(ctx context.Context, lastResults DescendantListResult) (result DescendantListResult, err error) {
	req, err := lastResults.descendantListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "managementgroups.Client", "getDescendantsNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.GetDescendantsSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "managementgroups.Client", "getDescendantsNextResults", resp, "Failure sending next results request")
	}
	result, err = client.GetDescendantsResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "managementgroups.Client", "getDescendantsNextResults", resp, "Failure responding to next results request")
	}
	return
}

// GetDescendantsComplete enumerates all values, automatically crossing page boundaries as required.
func (client Client) GetDescendantsComplete(ctx context.Context, groupID string, skiptoken string, top *int32) (result DescendantListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/Client.GetDescendants")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.GetDescendants(ctx, groupID, skiptoken, top)
	return
}

// List list management groups for the authenticated user.
// Parameters:
// cacheControl - indicates that the request shouldn't utilize any caches.
// skiptoken - page continuation token is only used if a previous operation returned a partial result.
// If a previous response contains a nextLink element, the value of the nextLink element will include a token
// parameter that specifies a starting point to use for subsequent calls.
func (client Client) List(ctx context.Context, cacheControl string, skiptoken string) (result ListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/Client.List")
		defer func() {
			sc := -1
			if result.lr.Response.Response != nil {
				sc = result.lr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx, cacheControl, skiptoken)
	if err != nil {
		err = autorest.NewErrorWithError(err, "managementgroups.Client", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.lr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "managementgroups.Client", "List", resp, "Failure sending request")
		return
	}

	result.lr, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "managementgroups.Client", "List", resp, "Failure responding to request")
		return
	}
	if result.lr.hasNextLink() && result.lr.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListPreparer prepares the List request.
func (client Client) ListPreparer(ctx context.Context, cacheControl string, skiptoken string) (*http.Request, error) {
	const APIVersion = "2020-02-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(skiptoken) > 0 {
		queryParameters["$skiptoken"] = autorest.Encode("query", skiptoken)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/providers/Microsoft.Management/managementGroups"),
		autorest.WithQueryParameters(queryParameters))
	if len(cacheControl) > 0 {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithHeader("Cache-Control", autorest.String(cacheControl)))
	} else {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithHeader("Cache-Control", autorest.String("no-cache")))
	}
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client Client) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client Client) ListResponder(resp *http.Response) (result ListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client Client) listNextResults(ctx context.Context, lastResults ListResult) (result ListResult, err error) {
	req, err := lastResults.listResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "managementgroups.Client", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "managementgroups.Client", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "managementgroups.Client", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client Client) ListComplete(ctx context.Context, cacheControl string, skiptoken string) (result ListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/Client.List")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.List(ctx, cacheControl, skiptoken)
	return
}

// Update update a management group.
// Parameters:
// groupID - management Group ID.
// patchGroupRequest - management group patch parameters.
// cacheControl - indicates that the request shouldn't utilize any caches.
func (client Client) Update(ctx context.Context, groupID string, patchGroupRequest PatchManagementGroupRequest, cacheControl string) (result ManagementGroup, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/Client.Update")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.UpdatePreparer(ctx, groupID, patchGroupRequest, cacheControl)
	if err != nil {
		err = autorest.NewErrorWithError(err, "managementgroups.Client", "Update", nil, "Failure preparing request")
		return
	}

	resp, err := client.UpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "managementgroups.Client", "Update", resp, "Failure sending request")
		return
	}

	result, err = client.UpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "managementgroups.Client", "Update", resp, "Failure responding to request")
		return
	}

	return
}

// UpdatePreparer prepares the Update request.
func (client Client) UpdatePreparer(ctx context.Context, groupID string, patchGroupRequest PatchManagementGroupRequest, cacheControl string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"groupId": autorest.Encode("path", groupID),
	}

	const APIVersion = "2020-02-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPatch(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/providers/Microsoft.Management/managementGroups/{groupId}", pathParameters),
		autorest.WithJSON(patchGroupRequest),
		autorest.WithQueryParameters(queryParameters))
	if len(cacheControl) > 0 {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithHeader("Cache-Control", autorest.String(cacheControl)))
	} else {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithHeader("Cache-Control", autorest.String("no-cache")))
	}
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// UpdateSender sends the Update request. The method will close the
// http.Response Body if it receives an error.
func (client Client) UpdateSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// UpdateResponder handles the response to the Update request. The method always
// closes the http.Response Body.
func (client Client) UpdateResponder(resp *http.Response) (result ManagementGroup, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
