package billing

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

// ProfilesClient is the billing client provides access to billing resources for Azure subscriptions.
type ProfilesClient struct {
	BaseClient
}

// NewProfilesClient creates an instance of the ProfilesClient client.
func NewProfilesClient(subscriptionID string) ProfilesClient {
	return NewProfilesClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewProfilesClientWithBaseURI creates an instance of the ProfilesClient client using a custom endpoint.  Use this
// when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewProfilesClientWithBaseURI(baseURI string, subscriptionID string) ProfilesClient {
	return ProfilesClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate creates or updates a billing profile. The operation is supported for billing accounts with agreement
// type Microsoft Customer Agreement or Microsoft Partner Agreement.
// Parameters:
// billingAccountName - the ID that uniquely identifies a billing account.
// billingProfileName - the ID that uniquely identifies a billing profile.
// parameters - the new or updated billing profile.
func (client ProfilesClient) CreateOrUpdate(ctx context.Context, billingAccountName string, billingProfileName string, parameters Profile) (result ProfilesCreateOrUpdateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ProfilesClient.CreateOrUpdate")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: parameters,
			Constraints: []validation.Constraint{{Target: "parameters.ProfileProperties", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "parameters.ProfileProperties.BillTo", Name: validation.Null, Rule: false,
					Chain: []validation.Constraint{{Target: "parameters.ProfileProperties.BillTo.AddressLine1", Name: validation.Null, Rule: true, Chain: nil},
						{Target: "parameters.ProfileProperties.BillTo.Country", Name: validation.Null, Rule: true, Chain: nil},
					}},
				}}}}}); err != nil {
		return result, validation.NewError("billing.ProfilesClient", "CreateOrUpdate", err.Error())
	}

	req, err := client.CreateOrUpdatePreparer(ctx, billingAccountName, billingProfileName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "billing.ProfilesClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	result, err = client.CreateOrUpdateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "billing.ProfilesClient", "CreateOrUpdate", result.Response(), "Failure sending request")
		return
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client ProfilesClient) CreateOrUpdatePreparer(ctx context.Context, billingAccountName string, billingProfileName string, parameters Profile) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"billingAccountName": autorest.Encode("path", billingAccountName),
		"billingProfileName": autorest.Encode("path", billingProfileName),
	}

	const APIVersion = "2020-05-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/billingProfiles/{billingProfileName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client ProfilesClient) CreateOrUpdateSender(req *http.Request) (future ProfilesCreateOrUpdateFuture, err error) {
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
func (client ProfilesClient) CreateOrUpdateResponder(resp *http.Response) (result Profile, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Get gets a billing profile by its ID. The operation is supported for billing accounts with agreement type Microsoft
// Customer Agreement or Microsoft Partner Agreement.
// Parameters:
// billingAccountName - the ID that uniquely identifies a billing account.
// billingProfileName - the ID that uniquely identifies a billing profile.
// expand - may be used to expand the invoice sections.
func (client ProfilesClient) Get(ctx context.Context, billingAccountName string, billingProfileName string, expand string) (result Profile, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ProfilesClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, billingAccountName, billingProfileName, expand)
	if err != nil {
		err = autorest.NewErrorWithError(err, "billing.ProfilesClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "billing.ProfilesClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "billing.ProfilesClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client ProfilesClient) GetPreparer(ctx context.Context, billingAccountName string, billingProfileName string, expand string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"billingAccountName": autorest.Encode("path", billingAccountName),
		"billingProfileName": autorest.Encode("path", billingProfileName),
	}

	const APIVersion = "2020-05-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(expand) > 0 {
		queryParameters["$expand"] = autorest.Encode("query", expand)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/billingProfiles/{billingProfileName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client ProfilesClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client ProfilesClient) GetResponder(resp *http.Response) (result Profile, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByBillingAccount lists the billing profiles that a user has access to. The operation is supported for billing
// accounts with agreement type Microsoft Customer Agreement or Microsoft Partner Agreement.
// Parameters:
// billingAccountName - the ID that uniquely identifies a billing account.
// expand - may be used to expand the invoice sections.
func (client ProfilesClient) ListByBillingAccount(ctx context.Context, billingAccountName string, expand string) (result ProfileListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ProfilesClient.ListByBillingAccount")
		defer func() {
			sc := -1
			if result.plr.Response.Response != nil {
				sc = result.plr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listByBillingAccountNextResults
	req, err := client.ListByBillingAccountPreparer(ctx, billingAccountName, expand)
	if err != nil {
		err = autorest.NewErrorWithError(err, "billing.ProfilesClient", "ListByBillingAccount", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByBillingAccountSender(req)
	if err != nil {
		result.plr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "billing.ProfilesClient", "ListByBillingAccount", resp, "Failure sending request")
		return
	}

	result.plr, err = client.ListByBillingAccountResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "billing.ProfilesClient", "ListByBillingAccount", resp, "Failure responding to request")
		return
	}
	if result.plr.hasNextLink() && result.plr.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListByBillingAccountPreparer prepares the ListByBillingAccount request.
func (client ProfilesClient) ListByBillingAccountPreparer(ctx context.Context, billingAccountName string, expand string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"billingAccountName": autorest.Encode("path", billingAccountName),
	}

	const APIVersion = "2020-05-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(expand) > 0 {
		queryParameters["$expand"] = autorest.Encode("query", expand)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/billingProfiles", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByBillingAccountSender sends the ListByBillingAccount request. The method will close the
// http.Response Body if it receives an error.
func (client ProfilesClient) ListByBillingAccountSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// ListByBillingAccountResponder handles the response to the ListByBillingAccount request. The method always
// closes the http.Response Body.
func (client ProfilesClient) ListByBillingAccountResponder(resp *http.Response) (result ProfileListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByBillingAccountNextResults retrieves the next set of results, if any.
func (client ProfilesClient) listByBillingAccountNextResults(ctx context.Context, lastResults ProfileListResult) (result ProfileListResult, err error) {
	req, err := lastResults.profileListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "billing.ProfilesClient", "listByBillingAccountNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByBillingAccountSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "billing.ProfilesClient", "listByBillingAccountNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByBillingAccountResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "billing.ProfilesClient", "listByBillingAccountNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByBillingAccountComplete enumerates all values, automatically crossing page boundaries as required.
func (client ProfilesClient) ListByBillingAccountComplete(ctx context.Context, billingAccountName string, expand string) (result ProfileListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ProfilesClient.ListByBillingAccount")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListByBillingAccount(ctx, billingAccountName, expand)
	return
}
