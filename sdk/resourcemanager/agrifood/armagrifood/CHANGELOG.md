# Release History

## 0.5.0 (2022-04-15)
### Breaking Changes

- Function `*FarmBeatsExtensionsClient.List` has been removed
- Function `*ExtensionsClient.ListByFarmBeats` has been removed
- Function `*FarmBeatsModelsClient.ListBySubscription` has been removed
- Function `*FarmBeatsModelsClient.ListByResourceGroup` has been removed
- Function `*OperationsClient.List` has been removed

### Features Added

- New function `*FarmBeatsModelsClient.NewListByResourceGroupPager(string, *FarmBeatsModelsClientListByResourceGroupOptions) *runtime.Pager[FarmBeatsModelsClientListByResourceGroupResponse]`
- New function `*FarmBeatsExtensionsClient.NewListPager(*FarmBeatsExtensionsClientListOptions) *runtime.Pager[FarmBeatsExtensionsClientListResponse]`
- New function `*OperationsClient.NewListPager(*OperationsClientListOptions) *runtime.Pager[OperationsClientListResponse]`
- New function `*ExtensionsClient.NewListByFarmBeatsPager(string, string, *ExtensionsClientListByFarmBeatsOptions) *runtime.Pager[ExtensionsClientListByFarmBeatsResponse]`
- New function `*FarmBeatsModelsClient.NewListBySubscriptionPager(*FarmBeatsModelsClientListBySubscriptionOptions) *runtime.Pager[FarmBeatsModelsClientListBySubscriptionResponse]`


## 0.4.0 (2022-04-11)
### Breaking Changes

- Function `*OperationsClient.List` return value(s) have been changed from `(*OperationsClientListPager)` to `(*runtime.Pager[OperationsClientListResponse])`
- Function `*FarmBeatsExtensionsClient.List` return value(s) have been changed from `(*FarmBeatsExtensionsClientListPager)` to `(*runtime.Pager[FarmBeatsExtensionsClientListResponse])`
- Function `NewExtensionsClient` return value(s) have been changed from `(*ExtensionsClient)` to `(*ExtensionsClient, error)`
- Function `NewOperationsClient` return value(s) have been changed from `(*OperationsClient)` to `(*OperationsClient, error)`
- Function `*FarmBeatsModelsClient.ListByResourceGroup` return value(s) have been changed from `(*FarmBeatsModelsClientListByResourceGroupPager)` to `(*runtime.Pager[FarmBeatsModelsClientListByResourceGroupResponse])`
- Function `NewFarmBeatsExtensionsClient` return value(s) have been changed from `(*FarmBeatsExtensionsClient)` to `(*FarmBeatsExtensionsClient, error)`
- Function `*FarmBeatsModelsClient.ListBySubscription` return value(s) have been changed from `(*FarmBeatsModelsClientListBySubscriptionPager)` to `(*runtime.Pager[FarmBeatsModelsClientListBySubscriptionResponse])`
- Function `NewLocationsClient` return value(s) have been changed from `(*LocationsClient)` to `(*LocationsClient, error)`
- Function `NewFarmBeatsModelsClient` return value(s) have been changed from `(*FarmBeatsModelsClient)` to `(*FarmBeatsModelsClient, error)`
- Function `*ExtensionsClient.ListByFarmBeats` return value(s) have been changed from `(*ExtensionsClientListByFarmBeatsPager)` to `(*runtime.Pager[ExtensionsClientListByFarmBeatsResponse])`
- Type of `ErrorAdditionalInfo.Info` has been changed from `map[string]interface{}` to `interface{}`
- Function `*OperationsClientListPager.NextPage` has been removed
- Function `*FarmBeatsModelsClientListBySubscriptionPager.Err` has been removed
- Function `*ExtensionsClientListByFarmBeatsPager.NextPage` has been removed
- Function `*FarmBeatsModelsClientListBySubscriptionPager.PageResponse` has been removed
- Function `*FarmBeatsModelsClientListByResourceGroupPager.Err` has been removed
- Function `*FarmBeatsExtensionsClientListPager.NextPage` has been removed
- Function `*FarmBeatsModelsClientListBySubscriptionPager.NextPage` has been removed
- Function `*ExtensionsClientListByFarmBeatsPager.Err` has been removed
- Function `*OperationsClientListPager.PageResponse` has been removed
- Function `*FarmBeatsModelsClientListByResourceGroupPager.PageResponse` has been removed
- Function `*FarmBeatsExtensionsClientListPager.PageResponse` has been removed
- Function `*OperationsClientListPager.Err` has been removed
- Function `*FarmBeatsModelsClientListByResourceGroupPager.NextPage` has been removed
- Function `*FarmBeatsExtensionsClientListPager.Err` has been removed
- Function `*ExtensionsClientListByFarmBeatsPager.PageResponse` has been removed
- Function `ActionType.ToPtr` has been removed
- Function `CreatedByType.ToPtr` has been removed
- Function `ProvisioningState.ToPtr` has been removed
- Function `Origin.ToPtr` has been removed
- Function `CheckNameAvailabilityReason.ToPtr` has been removed
- Struct `ExtensionsClientCreateResult` has been removed
- Struct `ExtensionsClientGetResult` has been removed
- Struct `ExtensionsClientListByFarmBeatsPager` has been removed
- Struct `ExtensionsClientListByFarmBeatsResult` has been removed
- Struct `ExtensionsClientUpdateResult` has been removed
- Struct `FarmBeatsExtensionsClientGetResult` has been removed
- Struct `FarmBeatsExtensionsClientListPager` has been removed
- Struct `FarmBeatsExtensionsClientListResult` has been removed
- Struct `FarmBeatsModelsClientCreateOrUpdateResult` has been removed
- Struct `FarmBeatsModelsClientGetResult` has been removed
- Struct `FarmBeatsModelsClientListByResourceGroupPager` has been removed
- Struct `FarmBeatsModelsClientListByResourceGroupResult` has been removed
- Struct `FarmBeatsModelsClientListBySubscriptionPager` has been removed
- Struct `FarmBeatsModelsClientListBySubscriptionResult` has been removed
- Struct `FarmBeatsModelsClientUpdateResult` has been removed
- Struct `LocationsClientCheckNameAvailabilityResult` has been removed
- Struct `OperationsClientListPager` has been removed
- Struct `OperationsClientListResult` has been removed
- Field `ExtensionsClientGetResult` of struct `ExtensionsClientGetResponse` has been removed
- Field `RawResponse` of struct `ExtensionsClientGetResponse` has been removed
- Field `ExtensionsClientCreateResult` of struct `ExtensionsClientCreateResponse` has been removed
- Field `RawResponse` of struct `ExtensionsClientCreateResponse` has been removed
- Field `OperationsClientListResult` of struct `OperationsClientListResponse` has been removed
- Field `RawResponse` of struct `OperationsClientListResponse` has been removed
- Field `FarmBeatsModelsClientCreateOrUpdateResult` of struct `FarmBeatsModelsClientCreateOrUpdateResponse` has been removed
- Field `RawResponse` of struct `FarmBeatsModelsClientCreateOrUpdateResponse` has been removed
- Field `LocationsClientCheckNameAvailabilityResult` of struct `LocationsClientCheckNameAvailabilityResponse` has been removed
- Field `RawResponse` of struct `LocationsClientCheckNameAvailabilityResponse` has been removed
- Field `RawResponse` of struct `ExtensionsClientDeleteResponse` has been removed
- Field `FarmBeatsModelsClientUpdateResult` of struct `FarmBeatsModelsClientUpdateResponse` has been removed
- Field `RawResponse` of struct `FarmBeatsModelsClientUpdateResponse` has been removed
- Field `ExtensionsClientUpdateResult` of struct `ExtensionsClientUpdateResponse` has been removed
- Field `RawResponse` of struct `ExtensionsClientUpdateResponse` has been removed
- Field `FarmBeatsExtensionsClientListResult` of struct `FarmBeatsExtensionsClientListResponse` has been removed
- Field `RawResponse` of struct `FarmBeatsExtensionsClientListResponse` has been removed
- Field `ExtensionsClientListByFarmBeatsResult` of struct `ExtensionsClientListByFarmBeatsResponse` has been removed
- Field `RawResponse` of struct `ExtensionsClientListByFarmBeatsResponse` has been removed
- Field `FarmBeatsModelsClientGetResult` of struct `FarmBeatsModelsClientGetResponse` has been removed
- Field `RawResponse` of struct `FarmBeatsModelsClientGetResponse` has been removed
- Field `RawResponse` of struct `FarmBeatsModelsClientDeleteResponse` has been removed
- Field `FarmBeatsExtensionsClientGetResult` of struct `FarmBeatsExtensionsClientGetResponse` has been removed
- Field `RawResponse` of struct `FarmBeatsExtensionsClientGetResponse` has been removed
- Field `FarmBeatsModelsClientListBySubscriptionResult` of struct `FarmBeatsModelsClientListBySubscriptionResponse` has been removed
- Field `RawResponse` of struct `FarmBeatsModelsClientListBySubscriptionResponse` has been removed
- Field `FarmBeatsModelsClientListByResourceGroupResult` of struct `FarmBeatsModelsClientListByResourceGroupResponse` has been removed
- Field `RawResponse` of struct `FarmBeatsModelsClientListByResourceGroupResponse` has been removed

### Features Added

- New anonymous field `FarmBeats` in struct `FarmBeatsModelsClientUpdateResponse`
- New anonymous field `FarmBeatsExtensionListResponse` in struct `FarmBeatsExtensionsClientListResponse`
- New anonymous field `FarmBeatsListResponse` in struct `FarmBeatsModelsClientListByResourceGroupResponse`
- New anonymous field `OperationListResult` in struct `OperationsClientListResponse`
- New anonymous field `FarmBeatsExtension` in struct `FarmBeatsExtensionsClientGetResponse`
- New anonymous field `FarmBeats` in struct `FarmBeatsModelsClientCreateOrUpdateResponse`
- New anonymous field `Extension` in struct `ExtensionsClientCreateResponse`
- New anonymous field `FarmBeats` in struct `FarmBeatsModelsClientGetResponse`
- New anonymous field `CheckNameAvailabilityResponse` in struct `LocationsClientCheckNameAvailabilityResponse`
- New anonymous field `FarmBeatsListResponse` in struct `FarmBeatsModelsClientListBySubscriptionResponse`
- New anonymous field `ExtensionListResponse` in struct `ExtensionsClientListByFarmBeatsResponse`
- New anonymous field `Extension` in struct `ExtensionsClientUpdateResponse`
- New anonymous field `Extension` in struct `ExtensionsClientGetResponse`


## 0.3.1 (2022-02-22)

### Other Changes

- Remove the go_mod_tidy_hack.go file.

## 0.3.0 (2022-01-13)
### Breaking Changes

- Function `*FarmBeatsExtensionsClient.Get` parameter(s) have been changed from `(context.Context, string, *FarmBeatsExtensionsGetOptions)` to `(context.Context, string, *FarmBeatsExtensionsClientGetOptions)`
- Function `*FarmBeatsExtensionsClient.Get` return value(s) have been changed from `(FarmBeatsExtensionsGetResponse, error)` to `(FarmBeatsExtensionsClientGetResponse, error)`
- Function `*FarmBeatsExtensionsClient.List` parameter(s) have been changed from `(*FarmBeatsExtensionsListOptions)` to `(*FarmBeatsExtensionsClientListOptions)`
- Function `*FarmBeatsExtensionsClient.List` return value(s) have been changed from `(*FarmBeatsExtensionsListPager)` to `(*FarmBeatsExtensionsClientListPager)`
- Function `*FarmBeatsModelsClient.CreateOrUpdate` parameter(s) have been changed from `(context.Context, string, string, FarmBeats, *FarmBeatsModelsCreateOrUpdateOptions)` to `(context.Context, string, string, FarmBeats, *FarmBeatsModelsClientCreateOrUpdateOptions)`
- Function `*FarmBeatsModelsClient.CreateOrUpdate` return value(s) have been changed from `(FarmBeatsModelsCreateOrUpdateResponse, error)` to `(FarmBeatsModelsClientCreateOrUpdateResponse, error)`
- Function `*FarmBeatsModelsClient.Delete` parameter(s) have been changed from `(context.Context, string, string, *FarmBeatsModelsDeleteOptions)` to `(context.Context, string, string, *FarmBeatsModelsClientDeleteOptions)`
- Function `*FarmBeatsModelsClient.Delete` return value(s) have been changed from `(FarmBeatsModelsDeleteResponse, error)` to `(FarmBeatsModelsClientDeleteResponse, error)`
- Function `*ExtensionsClient.ListByFarmBeats` parameter(s) have been changed from `(string, string, *ExtensionsListByFarmBeatsOptions)` to `(string, string, *ExtensionsClientListByFarmBeatsOptions)`
- Function `*ExtensionsClient.ListByFarmBeats` return value(s) have been changed from `(*ExtensionsListByFarmBeatsPager)` to `(*ExtensionsClientListByFarmBeatsPager)`
- Function `*FarmBeatsModelsClient.Get` parameter(s) have been changed from `(context.Context, string, string, *FarmBeatsModelsGetOptions)` to `(context.Context, string, string, *FarmBeatsModelsClientGetOptions)`
- Function `*FarmBeatsModelsClient.Get` return value(s) have been changed from `(FarmBeatsModelsGetResponse, error)` to `(FarmBeatsModelsClientGetResponse, error)`
- Function `*OperationsClient.List` parameter(s) have been changed from `(*OperationsListOptions)` to `(*OperationsClientListOptions)`
- Function `*OperationsClient.List` return value(s) have been changed from `(*OperationsListPager)` to `(*OperationsClientListPager)`
- Function `*FarmBeatsModelsClient.ListByResourceGroup` parameter(s) have been changed from `(string, *FarmBeatsModelsListByResourceGroupOptions)` to `(string, *FarmBeatsModelsClientListByResourceGroupOptions)`
- Function `*FarmBeatsModelsClient.ListByResourceGroup` return value(s) have been changed from `(*FarmBeatsModelsListByResourceGroupPager)` to `(*FarmBeatsModelsClientListByResourceGroupPager)`
- Function `*FarmBeatsModelsClient.Update` parameter(s) have been changed from `(context.Context, string, string, FarmBeatsUpdateRequestModel, *FarmBeatsModelsUpdateOptions)` to `(context.Context, string, string, FarmBeatsUpdateRequestModel, *FarmBeatsModelsClientUpdateOptions)`
- Function `*FarmBeatsModelsClient.Update` return value(s) have been changed from `(FarmBeatsModelsUpdateResponse, error)` to `(FarmBeatsModelsClientUpdateResponse, error)`
- Function `*ExtensionsClient.Delete` parameter(s) have been changed from `(context.Context, string, string, string, *ExtensionsDeleteOptions)` to `(context.Context, string, string, string, *ExtensionsClientDeleteOptions)`
- Function `*ExtensionsClient.Delete` return value(s) have been changed from `(ExtensionsDeleteResponse, error)` to `(ExtensionsClientDeleteResponse, error)`
- Function `*ExtensionsClient.Get` parameter(s) have been changed from `(context.Context, string, string, string, *ExtensionsGetOptions)` to `(context.Context, string, string, string, *ExtensionsClientGetOptions)`
- Function `*ExtensionsClient.Get` return value(s) have been changed from `(ExtensionsGetResponse, error)` to `(ExtensionsClientGetResponse, error)`
- Function `*LocationsClient.CheckNameAvailability` parameter(s) have been changed from `(context.Context, CheckNameAvailabilityRequest, *LocationsCheckNameAvailabilityOptions)` to `(context.Context, CheckNameAvailabilityRequest, *LocationsClientCheckNameAvailabilityOptions)`
- Function `*LocationsClient.CheckNameAvailability` return value(s) have been changed from `(LocationsCheckNameAvailabilityResponse, error)` to `(LocationsClientCheckNameAvailabilityResponse, error)`
- Function `*FarmBeatsModelsClient.ListBySubscription` parameter(s) have been changed from `(*FarmBeatsModelsListBySubscriptionOptions)` to `(*FarmBeatsModelsClientListBySubscriptionOptions)`
- Function `*FarmBeatsModelsClient.ListBySubscription` return value(s) have been changed from `(*FarmBeatsModelsListBySubscriptionPager)` to `(*FarmBeatsModelsClientListBySubscriptionPager)`
- Function `*ExtensionsClient.Create` parameter(s) have been changed from `(context.Context, string, string, string, *ExtensionsCreateOptions)` to `(context.Context, string, string, string, *ExtensionsClientCreateOptions)`
- Function `*ExtensionsClient.Create` return value(s) have been changed from `(ExtensionsCreateResponse, error)` to `(ExtensionsClientCreateResponse, error)`
- Function `*ExtensionsClient.Update` parameter(s) have been changed from `(context.Context, string, string, string, *ExtensionsUpdateOptions)` to `(context.Context, string, string, string, *ExtensionsClientUpdateOptions)`
- Function `*ExtensionsClient.Update` return value(s) have been changed from `(ExtensionsUpdateResponse, error)` to `(ExtensionsClientUpdateResponse, error)`
- Function `*ExtensionsListByFarmBeatsPager.NextPage` has been removed
- Function `Extension.MarshalJSON` has been removed
- Function `*OperationsListPager.NextPage` has been removed
- Function `Resource.MarshalJSON` has been removed
- Function `*FarmBeatsExtensionsListPager.PageResponse` has been removed
- Function `*ExtensionsListByFarmBeatsPager.PageResponse` has been removed
- Function `*FarmBeatsExtensionsListPager.Err` has been removed
- Function `FarmBeatsExtension.MarshalJSON` has been removed
- Function `*FarmBeatsExtensionsListPager.NextPage` has been removed
- Function `ErrorResponse.Error` has been removed
- Function `*FarmBeatsModelsListByResourceGroupPager.PageResponse` has been removed
- Function `*FarmBeatsModelsListBySubscriptionPager.Err` has been removed
- Function `*ExtensionsListByFarmBeatsPager.Err` has been removed
- Function `*FarmBeatsModelsListByResourceGroupPager.Err` has been removed
- Function `*FarmBeatsModelsListByResourceGroupPager.NextPage` has been removed
- Function `*OperationsListPager.PageResponse` has been removed
- Function `*FarmBeatsModelsListBySubscriptionPager.NextPage` has been removed
- Function `*FarmBeatsModelsListBySubscriptionPager.PageResponse` has been removed
- Function `*OperationsListPager.Err` has been removed
- Struct `ExtensionsCreateOptions` has been removed
- Struct `ExtensionsCreateResponse` has been removed
- Struct `ExtensionsCreateResult` has been removed
- Struct `ExtensionsDeleteOptions` has been removed
- Struct `ExtensionsDeleteResponse` has been removed
- Struct `ExtensionsGetOptions` has been removed
- Struct `ExtensionsGetResponse` has been removed
- Struct `ExtensionsGetResult` has been removed
- Struct `ExtensionsListByFarmBeatsOptions` has been removed
- Struct `ExtensionsListByFarmBeatsPager` has been removed
- Struct `ExtensionsListByFarmBeatsResponse` has been removed
- Struct `ExtensionsListByFarmBeatsResult` has been removed
- Struct `ExtensionsUpdateOptions` has been removed
- Struct `ExtensionsUpdateResponse` has been removed
- Struct `ExtensionsUpdateResult` has been removed
- Struct `FarmBeatsExtensionsGetOptions` has been removed
- Struct `FarmBeatsExtensionsGetResponse` has been removed
- Struct `FarmBeatsExtensionsGetResult` has been removed
- Struct `FarmBeatsExtensionsListOptions` has been removed
- Struct `FarmBeatsExtensionsListPager` has been removed
- Struct `FarmBeatsExtensionsListResponse` has been removed
- Struct `FarmBeatsExtensionsListResult` has been removed
- Struct `FarmBeatsModelsCreateOrUpdateOptions` has been removed
- Struct `FarmBeatsModelsCreateOrUpdateResponse` has been removed
- Struct `FarmBeatsModelsCreateOrUpdateResult` has been removed
- Struct `FarmBeatsModelsDeleteOptions` has been removed
- Struct `FarmBeatsModelsDeleteResponse` has been removed
- Struct `FarmBeatsModelsGetOptions` has been removed
- Struct `FarmBeatsModelsGetResponse` has been removed
- Struct `FarmBeatsModelsGetResult` has been removed
- Struct `FarmBeatsModelsListByResourceGroupOptions` has been removed
- Struct `FarmBeatsModelsListByResourceGroupPager` has been removed
- Struct `FarmBeatsModelsListByResourceGroupResponse` has been removed
- Struct `FarmBeatsModelsListByResourceGroupResult` has been removed
- Struct `FarmBeatsModelsListBySubscriptionOptions` has been removed
- Struct `FarmBeatsModelsListBySubscriptionPager` has been removed
- Struct `FarmBeatsModelsListBySubscriptionResponse` has been removed
- Struct `FarmBeatsModelsListBySubscriptionResult` has been removed
- Struct `FarmBeatsModelsUpdateOptions` has been removed
- Struct `FarmBeatsModelsUpdateResponse` has been removed
- Struct `FarmBeatsModelsUpdateResult` has been removed
- Struct `LocationsCheckNameAvailabilityOptions` has been removed
- Struct `LocationsCheckNameAvailabilityResponse` has been removed
- Struct `LocationsCheckNameAvailabilityResult` has been removed
- Struct `OperationsListOptions` has been removed
- Struct `OperationsListPager` has been removed
- Struct `OperationsListResponse` has been removed
- Struct `OperationsListResult` has been removed
- Field `ProxyResource` of struct `FarmBeatsExtension` has been removed
- Field `ProxyResource` of struct `Extension` has been removed
- Field `InnerError` of struct `ErrorResponse` has been removed
- Field `Resource` of struct `TrackedResource` has been removed
- Field `TrackedResource` of struct `FarmBeats` has been removed
- Field `Resource` of struct `ProxyResource` has been removed

### Features Added

- New function `*ExtensionsClientListByFarmBeatsPager.NextPage(context.Context) bool`
- New function `*FarmBeatsExtensionsClientListPager.NextPage(context.Context) bool`
- New function `*ExtensionsClientListByFarmBeatsPager.Err() error`
- New function `*ExtensionsClientListByFarmBeatsPager.PageResponse() ExtensionsClientListByFarmBeatsResponse`
- New function `*FarmBeatsModelsClientListBySubscriptionPager.Err() error`
- New function `*FarmBeatsModelsClientListByResourceGroupPager.PageResponse() FarmBeatsModelsClientListByResourceGroupResponse`
- New function `*FarmBeatsModelsClientListBySubscriptionPager.NextPage(context.Context) bool`
- New function `*OperationsClientListPager.Err() error`
- New function `*FarmBeatsExtensionsClientListPager.Err() error`
- New function `*FarmBeatsModelsClientListByResourceGroupPager.Err() error`
- New function `*OperationsClientListPager.NextPage(context.Context) bool`
- New function `*OperationsClientListPager.PageResponse() OperationsClientListResponse`
- New function `*FarmBeatsExtensionsClientListPager.PageResponse() FarmBeatsExtensionsClientListResponse`
- New function `*FarmBeatsModelsClientListByResourceGroupPager.NextPage(context.Context) bool`
- New function `*FarmBeatsModelsClientListBySubscriptionPager.PageResponse() FarmBeatsModelsClientListBySubscriptionResponse`
- New struct `ExtensionsClientCreateOptions`
- New struct `ExtensionsClientCreateResponse`
- New struct `ExtensionsClientCreateResult`
- New struct `ExtensionsClientDeleteOptions`
- New struct `ExtensionsClientDeleteResponse`
- New struct `ExtensionsClientGetOptions`
- New struct `ExtensionsClientGetResponse`
- New struct `ExtensionsClientGetResult`
- New struct `ExtensionsClientListByFarmBeatsOptions`
- New struct `ExtensionsClientListByFarmBeatsPager`
- New struct `ExtensionsClientListByFarmBeatsResponse`
- New struct `ExtensionsClientListByFarmBeatsResult`
- New struct `ExtensionsClientUpdateOptions`
- New struct `ExtensionsClientUpdateResponse`
- New struct `ExtensionsClientUpdateResult`
- New struct `FarmBeatsExtensionsClientGetOptions`
- New struct `FarmBeatsExtensionsClientGetResponse`
- New struct `FarmBeatsExtensionsClientGetResult`
- New struct `FarmBeatsExtensionsClientListOptions`
- New struct `FarmBeatsExtensionsClientListPager`
- New struct `FarmBeatsExtensionsClientListResponse`
- New struct `FarmBeatsExtensionsClientListResult`
- New struct `FarmBeatsModelsClientCreateOrUpdateOptions`
- New struct `FarmBeatsModelsClientCreateOrUpdateResponse`
- New struct `FarmBeatsModelsClientCreateOrUpdateResult`
- New struct `FarmBeatsModelsClientDeleteOptions`
- New struct `FarmBeatsModelsClientDeleteResponse`
- New struct `FarmBeatsModelsClientGetOptions`
- New struct `FarmBeatsModelsClientGetResponse`
- New struct `FarmBeatsModelsClientGetResult`
- New struct `FarmBeatsModelsClientListByResourceGroupOptions`
- New struct `FarmBeatsModelsClientListByResourceGroupPager`
- New struct `FarmBeatsModelsClientListByResourceGroupResponse`
- New struct `FarmBeatsModelsClientListByResourceGroupResult`
- New struct `FarmBeatsModelsClientListBySubscriptionOptions`
- New struct `FarmBeatsModelsClientListBySubscriptionPager`
- New struct `FarmBeatsModelsClientListBySubscriptionResponse`
- New struct `FarmBeatsModelsClientListBySubscriptionResult`
- New struct `FarmBeatsModelsClientUpdateOptions`
- New struct `FarmBeatsModelsClientUpdateResponse`
- New struct `FarmBeatsModelsClientUpdateResult`
- New struct `LocationsClientCheckNameAvailabilityOptions`
- New struct `LocationsClientCheckNameAvailabilityResponse`
- New struct `LocationsClientCheckNameAvailabilityResult`
- New struct `OperationsClientListOptions`
- New struct `OperationsClientListPager`
- New struct `OperationsClientListResponse`
- New struct `OperationsClientListResult`
- New field `ID` in struct `TrackedResource`
- New field `Name` in struct `TrackedResource`
- New field `Type` in struct `TrackedResource`
- New field `Type` in struct `FarmBeatsExtension`
- New field `ID` in struct `FarmBeatsExtension`
- New field `Name` in struct `FarmBeatsExtension`
- New field `Error` in struct `ErrorResponse`
- New field `ID` in struct `ProxyResource`
- New field `Name` in struct `ProxyResource`
- New field `Type` in struct `ProxyResource`
- New field `ID` in struct `Extension`
- New field `Name` in struct `Extension`
- New field `Type` in struct `Extension`
- New field `Name` in struct `FarmBeats`
- New field `Type` in struct `FarmBeats`
- New field `Location` in struct `FarmBeats`
- New field `Tags` in struct `FarmBeats`
- New field `ID` in struct `FarmBeats`


## 0.2.0 (2021-10-29)

### Breaking Changes

- `arm.Connection` has been removed in `github.com/Azure/azure-sdk-for-go/sdk/azcore/v0.20.0`
- The parameters of `NewXXXClient` has been changed from `(con *arm.Connection, subscriptionID string)` to `(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions)`

## 0.1.0 (2021-10-08)
- To better align with the Azure SDK guidelines (https://azure.github.io/azure-sdk/general_introduction.html), we have decided to change the module path to "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/agrifood/armagrifood". Therefore, we are deprecating the old module path (which is "github.com/Azure/azure-sdk-for-go/sdk/agrifood/armagrifood") to avoid confusion.
