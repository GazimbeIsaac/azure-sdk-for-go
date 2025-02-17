# Release History

## 0.4.0 (2022-04-15)
### Breaking Changes

- Function `*OperationsClient.List` has been removed
- Function `*ProviderInstancesClient.List` has been removed
- Function `*SapMonitorsClient.List` has been removed

### Features Added

- New function `*ProviderInstancesClient.NewListPager(string, string, *ProviderInstancesClientListOptions) *runtime.Pager[ProviderInstancesClientListResponse]`
- New function `*SapMonitorsClient.NewListPager(*SapMonitorsClientListOptions) *runtime.Pager[SapMonitorsClientListResponse]`
- New function `*OperationsClient.NewListPager(*OperationsClientListOptions) *runtime.Pager[OperationsClientListResponse]`


## 0.3.0 (2022-04-11)
### Breaking Changes

- Function `*ProviderInstancesClient.BeginDelete` return value(s) have been changed from `(ProviderInstancesClientDeletePollerResponse, error)` to `(*armruntime.Poller[ProviderInstancesClientDeleteResponse], error)`
- Function `NewOperationsClient` return value(s) have been changed from `(*OperationsClient)` to `(*OperationsClient, error)`
- Function `NewProviderInstancesClient` return value(s) have been changed from `(*ProviderInstancesClient)` to `(*ProviderInstancesClient, error)`
- Function `*SapMonitorsClient.List` return value(s) have been changed from `(*SapMonitorsClientListPager)` to `(*runtime.Pager[SapMonitorsClientListResponse])`
- Function `*ProviderInstancesClient.List` return value(s) have been changed from `(*ProviderInstancesClientListPager)` to `(*runtime.Pager[ProviderInstancesClientListResponse])`
- Function `NewSapMonitorsClient` return value(s) have been changed from `(*SapMonitorsClient)` to `(*SapMonitorsClient, error)`
- Function `*ProviderInstancesClient.BeginCreate` return value(s) have been changed from `(ProviderInstancesClientCreatePollerResponse, error)` to `(*armruntime.Poller[ProviderInstancesClientCreateResponse], error)`
- Function `*SapMonitorsClient.BeginCreate` return value(s) have been changed from `(SapMonitorsClientCreatePollerResponse, error)` to `(*armruntime.Poller[SapMonitorsClientCreateResponse], error)`
- Function `*SapMonitorsClient.BeginDelete` return value(s) have been changed from `(SapMonitorsClientDeletePollerResponse, error)` to `(*armruntime.Poller[SapMonitorsClientDeleteResponse], error)`
- Function `*OperationsClient.List` parameter(s) have been changed from `(context.Context, *OperationsClientListOptions)` to `(*OperationsClientListOptions)`
- Function `*OperationsClient.List` return value(s) have been changed from `(OperationsClientListResponse, error)` to `(*runtime.Pager[OperationsClientListResponse])`
- Function `*ProviderInstancesClientDeletePollerResponse.Resume` has been removed
- Function `*ProviderInstancesClientListPager.Err` has been removed
- Function `*SapMonitorsClientCreatePoller.Poll` has been removed
- Function `*ProviderInstancesClientCreatePoller.FinalResponse` has been removed
- Function `*ProviderInstancesClientDeletePoller.FinalResponse` has been removed
- Function `*SapMonitorsClientDeletePoller.ResumeToken` has been removed
- Function `*SapMonitorsClientDeletePoller.FinalResponse` has been removed
- Function `*SapMonitorsClientListPager.PageResponse` has been removed
- Function `SapMonitorsClientDeletePollerResponse.PollUntilDone` has been removed
- Function `*ProviderInstancesClientDeletePoller.Poll` has been removed
- Function `*SapMonitorsClientDeletePollerResponse.Resume` has been removed
- Function `*ProviderInstancesClientListPager.NextPage` has been removed
- Function `*ProviderInstancesClientCreatePollerResponse.Resume` has been removed
- Function `*SapMonitorsClientListPager.NextPage` has been removed
- Function `*ProviderInstancesClientDeletePoller.Done` has been removed
- Function `*SapMonitorsClientCreatePoller.Done` has been removed
- Function `*SapMonitorsClientCreatePoller.ResumeToken` has been removed
- Function `SapMonitorsClientCreatePollerResponse.PollUntilDone` has been removed
- Function `ProviderInstancesClientDeletePollerResponse.PollUntilDone` has been removed
- Function `*ProviderInstancesClientCreatePoller.Poll` has been removed
- Function `*SapMonitorsClientCreatePoller.FinalResponse` has been removed
- Function `*ProviderInstancesClientDeletePoller.ResumeToken` has been removed
- Function `ProviderInstancesClientCreatePollerResponse.PollUntilDone` has been removed
- Function `*SapMonitorsClientListPager.Err` has been removed
- Function `*SapMonitorsClientDeletePoller.Done` has been removed
- Function `HanaProvisioningStatesEnum.ToPtr` has been removed
- Function `*ProviderInstancesClientCreatePoller.ResumeToken` has been removed
- Function `*SapMonitorsClientCreatePollerResponse.Resume` has been removed
- Function `*ProviderInstancesClientCreatePoller.Done` has been removed
- Function `*SapMonitorsClientDeletePoller.Poll` has been removed
- Function `*ProviderInstancesClientListPager.PageResponse` has been removed
- Struct `OperationsClientListResult` has been removed
- Struct `ProviderInstancesClientCreatePoller` has been removed
- Struct `ProviderInstancesClientCreatePollerResponse` has been removed
- Struct `ProviderInstancesClientCreateResult` has been removed
- Struct `ProviderInstancesClientDeletePoller` has been removed
- Struct `ProviderInstancesClientDeletePollerResponse` has been removed
- Struct `ProviderInstancesClientGetResult` has been removed
- Struct `ProviderInstancesClientListPager` has been removed
- Struct `ProviderInstancesClientListResult` has been removed
- Struct `SapMonitorsClientCreatePoller` has been removed
- Struct `SapMonitorsClientCreatePollerResponse` has been removed
- Struct `SapMonitorsClientCreateResult` has been removed
- Struct `SapMonitorsClientDeletePoller` has been removed
- Struct `SapMonitorsClientDeletePollerResponse` has been removed
- Struct `SapMonitorsClientGetResult` has been removed
- Struct `SapMonitorsClientListPager` has been removed
- Struct `SapMonitorsClientListResult` has been removed
- Struct `SapMonitorsClientUpdateResult` has been removed
- Field `SapMonitorsClientGetResult` of struct `SapMonitorsClientGetResponse` has been removed
- Field `RawResponse` of struct `SapMonitorsClientGetResponse` has been removed
- Field `SapMonitorsClientCreateResult` of struct `SapMonitorsClientCreateResponse` has been removed
- Field `RawResponse` of struct `SapMonitorsClientCreateResponse` has been removed
- Field `SapMonitorsClientUpdateResult` of struct `SapMonitorsClientUpdateResponse` has been removed
- Field `RawResponse` of struct `SapMonitorsClientUpdateResponse` has been removed
- Field `SapMonitorsClientListResult` of struct `SapMonitorsClientListResponse` has been removed
- Field `RawResponse` of struct `SapMonitorsClientListResponse` has been removed
- Field `ProviderInstancesClientCreateResult` of struct `ProviderInstancesClientCreateResponse` has been removed
- Field `RawResponse` of struct `ProviderInstancesClientCreateResponse` has been removed
- Field `ProviderInstancesClientListResult` of struct `ProviderInstancesClientListResponse` has been removed
- Field `RawResponse` of struct `ProviderInstancesClientListResponse` has been removed
- Field `RawResponse` of struct `ProviderInstancesClientDeleteResponse` has been removed
- Field `ProviderInstancesClientGetResult` of struct `ProviderInstancesClientGetResponse` has been removed
- Field `RawResponse` of struct `ProviderInstancesClientGetResponse` has been removed
- Field `OperationsClientListResult` of struct `OperationsClientListResponse` has been removed
- Field `RawResponse` of struct `OperationsClientListResponse` has been removed
- Field `RawResponse` of struct `SapMonitorsClientDeleteResponse` has been removed

### Features Added

- New field `ResumeToken` in struct `ProviderInstancesClientBeginDeleteOptions`
- New anonymous field `SapMonitor` in struct `SapMonitorsClientUpdateResponse`
- New anonymous field `ProviderInstance` in struct `ProviderInstancesClientCreateResponse`
- New anonymous field `SapMonitor` in struct `SapMonitorsClientGetResponse`
- New anonymous field `SapMonitor` in struct `SapMonitorsClientCreateResponse`
- New anonymous field `OperationList` in struct `OperationsClientListResponse`
- New anonymous field `SapMonitorListResult` in struct `SapMonitorsClientListResponse`
- New field `ResumeToken` in struct `SapMonitorsClientBeginDeleteOptions`
- New field `ResumeToken` in struct `SapMonitorsClientBeginCreateOptions`
- New anonymous field `ProviderInstanceListResult` in struct `ProviderInstancesClientListResponse`
- New anonymous field `ProviderInstance` in struct `ProviderInstancesClientGetResponse`
- New field `ResumeToken` in struct `ProviderInstancesClientBeginCreateOptions`


## 0.2.1 (2022-02-22)

### Other Changes

- Remove the go_mod_tidy_hack.go file.

## 0.2.0 (2022-01-13)
### Breaking Changes

- Function `*ProviderInstancesClient.BeginCreate` parameter(s) have been changed from `(context.Context, string, string, string, ProviderInstance, *ProviderInstancesBeginCreateOptions)` to `(context.Context, string, string, string, ProviderInstance, *ProviderInstancesClientBeginCreateOptions)`
- Function `*ProviderInstancesClient.BeginCreate` return value(s) have been changed from `(ProviderInstancesCreatePollerResponse, error)` to `(ProviderInstancesClientCreatePollerResponse, error)`
- Function `*ProviderInstancesClient.Get` parameter(s) have been changed from `(context.Context, string, string, string, *ProviderInstancesGetOptions)` to `(context.Context, string, string, string, *ProviderInstancesClientGetOptions)`
- Function `*ProviderInstancesClient.Get` return value(s) have been changed from `(ProviderInstancesGetResponse, error)` to `(ProviderInstancesClientGetResponse, error)`
- Function `*SapMonitorsClient.Update` parameter(s) have been changed from `(context.Context, string, string, Tags, *SapMonitorsUpdateOptions)` to `(context.Context, string, string, Tags, *SapMonitorsClientUpdateOptions)`
- Function `*SapMonitorsClient.Update` return value(s) have been changed from `(SapMonitorsUpdateResponse, error)` to `(SapMonitorsClientUpdateResponse, error)`
- Function `*SapMonitorsClient.Get` parameter(s) have been changed from `(context.Context, string, string, *SapMonitorsGetOptions)` to `(context.Context, string, string, *SapMonitorsClientGetOptions)`
- Function `*SapMonitorsClient.Get` return value(s) have been changed from `(SapMonitorsGetResponse, error)` to `(SapMonitorsClientGetResponse, error)`
- Function `*ProviderInstancesClient.BeginDelete` parameter(s) have been changed from `(context.Context, string, string, string, *ProviderInstancesBeginDeleteOptions)` to `(context.Context, string, string, string, *ProviderInstancesClientBeginDeleteOptions)`
- Function `*ProviderInstancesClient.BeginDelete` return value(s) have been changed from `(ProviderInstancesDeletePollerResponse, error)` to `(ProviderInstancesClientDeletePollerResponse, error)`
- Function `*SapMonitorsClient.BeginCreate` parameter(s) have been changed from `(context.Context, string, string, SapMonitor, *SapMonitorsBeginCreateOptions)` to `(context.Context, string, string, SapMonitor, *SapMonitorsClientBeginCreateOptions)`
- Function `*SapMonitorsClient.BeginCreate` return value(s) have been changed from `(SapMonitorsCreatePollerResponse, error)` to `(SapMonitorsClientCreatePollerResponse, error)`
- Function `*SapMonitorsClient.BeginDelete` parameter(s) have been changed from `(context.Context, string, string, *SapMonitorsBeginDeleteOptions)` to `(context.Context, string, string, *SapMonitorsClientBeginDeleteOptions)`
- Function `*SapMonitorsClient.BeginDelete` return value(s) have been changed from `(SapMonitorsDeletePollerResponse, error)` to `(SapMonitorsClientDeletePollerResponse, error)`
- Function `*SapMonitorsClient.List` parameter(s) have been changed from `(*SapMonitorsListOptions)` to `(*SapMonitorsClientListOptions)`
- Function `*SapMonitorsClient.List` return value(s) have been changed from `(*SapMonitorsListPager)` to `(*SapMonitorsClientListPager)`
- Function `*OperationsClient.List` parameter(s) have been changed from `(context.Context, *OperationsListOptions)` to `(context.Context, *OperationsClientListOptions)`
- Function `*OperationsClient.List` return value(s) have been changed from `(OperationsListResponse, error)` to `(OperationsClientListResponse, error)`
- Function `*ProviderInstancesClient.List` parameter(s) have been changed from `(string, string, *ProviderInstancesListOptions)` to `(string, string, *ProviderInstancesClientListOptions)`
- Function `*ProviderInstancesClient.List` return value(s) have been changed from `(*ProviderInstancesListPager)` to `(*ProviderInstancesClientListPager)`
- Function `*ProviderInstancesDeletePoller.ResumeToken` has been removed
- Function `*ProviderInstancesDeletePoller.Poll` has been removed
- Function `SapMonitorsCreatePollerResponse.PollUntilDone` has been removed
- Function `*ProviderInstancesCreatePoller.Poll` has been removed
- Function `*ProviderInstancesCreatePoller.ResumeToken` has been removed
- Function `*ProviderInstancesListPager.PageResponse` has been removed
- Function `SapMonitorsDeletePollerResponse.PollUntilDone` has been removed
- Function `*SapMonitorsDeletePollerResponse.Resume` has been removed
- Function `*SapMonitorsListPager.PageResponse` has been removed
- Function `*SapMonitorsDeletePoller.Poll` has been removed
- Function `*SapMonitorsListPager.Err` has been removed
- Function `*ProviderInstancesCreatePollerResponse.Resume` has been removed
- Function `*SapMonitorsDeletePoller.Done` has been removed
- Function `ProviderInstance.MarshalJSON` has been removed
- Function `ProviderInstancesCreatePollerResponse.PollUntilDone` has been removed
- Function `*SapMonitorsCreatePollerResponse.Resume` has been removed
- Function `*SapMonitorsCreatePoller.Done` has been removed
- Function `*SapMonitorsListPager.NextPage` has been removed
- Function `*ProviderInstancesCreatePoller.Done` has been removed
- Function `*SapMonitorsDeletePoller.FinalResponse` has been removed
- Function `*ProviderInstancesDeletePoller.FinalResponse` has been removed
- Function `*ProviderInstancesListPager.NextPage` has been removed
- Function `*ProviderInstancesDeletePoller.Done` has been removed
- Function `ErrorResponse.Error` has been removed
- Function `*ProviderInstancesCreatePoller.FinalResponse` has been removed
- Function `Resource.MarshalJSON` has been removed
- Function `*SapMonitorsCreatePoller.Poll` has been removed
- Function `*SapMonitorsDeletePoller.ResumeToken` has been removed
- Function `*SapMonitorsCreatePoller.ResumeToken` has been removed
- Function `*ProviderInstancesDeletePollerResponse.Resume` has been removed
- Function `*SapMonitorsCreatePoller.FinalResponse` has been removed
- Function `ProviderInstancesDeletePollerResponse.PollUntilDone` has been removed
- Function `*ProviderInstancesListPager.Err` has been removed
- Struct `OperationsListOptions` has been removed
- Struct `OperationsListResponse` has been removed
- Struct `OperationsListResult` has been removed
- Struct `ProviderInstancesBeginCreateOptions` has been removed
- Struct `ProviderInstancesBeginDeleteOptions` has been removed
- Struct `ProviderInstancesCreatePoller` has been removed
- Struct `ProviderInstancesCreatePollerResponse` has been removed
- Struct `ProviderInstancesCreateResponse` has been removed
- Struct `ProviderInstancesCreateResult` has been removed
- Struct `ProviderInstancesDeletePoller` has been removed
- Struct `ProviderInstancesDeletePollerResponse` has been removed
- Struct `ProviderInstancesDeleteResponse` has been removed
- Struct `ProviderInstancesGetOptions` has been removed
- Struct `ProviderInstancesGetResponse` has been removed
- Struct `ProviderInstancesGetResult` has been removed
- Struct `ProviderInstancesListOptions` has been removed
- Struct `ProviderInstancesListPager` has been removed
- Struct `ProviderInstancesListResponse` has been removed
- Struct `ProviderInstancesListResult` has been removed
- Struct `SapMonitorsBeginCreateOptions` has been removed
- Struct `SapMonitorsBeginDeleteOptions` has been removed
- Struct `SapMonitorsCreatePoller` has been removed
- Struct `SapMonitorsCreatePollerResponse` has been removed
- Struct `SapMonitorsCreateResponse` has been removed
- Struct `SapMonitorsCreateResult` has been removed
- Struct `SapMonitorsDeletePoller` has been removed
- Struct `SapMonitorsDeletePollerResponse` has been removed
- Struct `SapMonitorsDeleteResponse` has been removed
- Struct `SapMonitorsGetOptions` has been removed
- Struct `SapMonitorsGetResponse` has been removed
- Struct `SapMonitorsGetResult` has been removed
- Struct `SapMonitorsListOptions` has been removed
- Struct `SapMonitorsListPager` has been removed
- Struct `SapMonitorsListResponse` has been removed
- Struct `SapMonitorsListResult` has been removed
- Struct `SapMonitorsUpdateOptions` has been removed
- Struct `SapMonitorsUpdateResponse` has been removed
- Struct `SapMonitorsUpdateResult` has been removed
- Field `InnerError` of struct `ErrorResponse` has been removed
- Field `TrackedResource` of struct `SapMonitor` has been removed
- Field `Resource` of struct `TrackedResource` has been removed
- Field `ProxyResource` of struct `ProviderInstance` has been removed
- Field `Resource` of struct `ProxyResource` has been removed

### Features Added

- New function `*ProviderInstancesClientDeletePollerResponse.Resume(context.Context, *ProviderInstancesClient, string) error`
- New function `*SapMonitorsClientDeletePoller.FinalResponse(context.Context) (SapMonitorsClientDeleteResponse, error)`
- New function `*ProviderInstancesClientDeletePoller.Done() bool`
- New function `SapMonitorsClientDeletePollerResponse.PollUntilDone(context.Context, time.Duration) (SapMonitorsClientDeleteResponse, error)`
- New function `*ProviderInstancesClientCreatePoller.ResumeToken() (string, error)`
- New function `*SapMonitorsClientCreatePoller.Done() bool`
- New function `*SapMonitorsClientListPager.NextPage(context.Context) bool`
- New function `SapMonitorsClientCreatePollerResponse.PollUntilDone(context.Context, time.Duration) (SapMonitorsClientCreateResponse, error)`
- New function `*ProviderInstancesClientListPager.NextPage(context.Context) bool`
- New function `*SapMonitorsClientCreatePollerResponse.Resume(context.Context, *SapMonitorsClient, string) error`
- New function `*SapMonitorsClientCreatePoller.ResumeToken() (string, error)`
- New function `*SapMonitorsClientDeletePoller.Done() bool`
- New function `*ProviderInstancesClientCreatePoller.FinalResponse(context.Context) (ProviderInstancesClientCreateResponse, error)`
- New function `*SapMonitorsClientCreatePoller.Poll(context.Context) (*http.Response, error)`
- New function `*ProviderInstancesClientListPager.Err() error`
- New function `*SapMonitorsClientDeletePollerResponse.Resume(context.Context, *SapMonitorsClient, string) error`
- New function `*ProviderInstancesClientCreatePoller.Poll(context.Context) (*http.Response, error)`
- New function `*ProviderInstancesClientCreatePoller.Done() bool`
- New function `*SapMonitorsClientListPager.Err() error`
- New function `*ProviderInstancesClientDeletePoller.ResumeToken() (string, error)`
- New function `*SapMonitorsClientListPager.PageResponse() SapMonitorsClientListResponse`
- New function `ProviderInstancesClientDeletePollerResponse.PollUntilDone(context.Context, time.Duration) (ProviderInstancesClientDeleteResponse, error)`
- New function `ProviderInstancesClientCreatePollerResponse.PollUntilDone(context.Context, time.Duration) (ProviderInstancesClientCreateResponse, error)`
- New function `*SapMonitorsClientCreatePoller.FinalResponse(context.Context) (SapMonitorsClientCreateResponse, error)`
- New function `*ProviderInstancesClientDeletePoller.FinalResponse(context.Context) (ProviderInstancesClientDeleteResponse, error)`
- New function `*ProviderInstancesClientListPager.PageResponse() ProviderInstancesClientListResponse`
- New function `*ProviderInstancesClientDeletePoller.Poll(context.Context) (*http.Response, error)`
- New function `*ProviderInstancesClientCreatePollerResponse.Resume(context.Context, *ProviderInstancesClient, string) error`
- New function `*SapMonitorsClientDeletePoller.Poll(context.Context) (*http.Response, error)`
- New function `*SapMonitorsClientDeletePoller.ResumeToken() (string, error)`
- New struct `OperationsClientListOptions`
- New struct `OperationsClientListResponse`
- New struct `OperationsClientListResult`
- New struct `ProviderInstancesClientBeginCreateOptions`
- New struct `ProviderInstancesClientBeginDeleteOptions`
- New struct `ProviderInstancesClientCreatePoller`
- New struct `ProviderInstancesClientCreatePollerResponse`
- New struct `ProviderInstancesClientCreateResponse`
- New struct `ProviderInstancesClientCreateResult`
- New struct `ProviderInstancesClientDeletePoller`
- New struct `ProviderInstancesClientDeletePollerResponse`
- New struct `ProviderInstancesClientDeleteResponse`
- New struct `ProviderInstancesClientGetOptions`
- New struct `ProviderInstancesClientGetResponse`
- New struct `ProviderInstancesClientGetResult`
- New struct `ProviderInstancesClientListOptions`
- New struct `ProviderInstancesClientListPager`
- New struct `ProviderInstancesClientListResponse`
- New struct `ProviderInstancesClientListResult`
- New struct `SapMonitorsClientBeginCreateOptions`
- New struct `SapMonitorsClientBeginDeleteOptions`
- New struct `SapMonitorsClientCreatePoller`
- New struct `SapMonitorsClientCreatePollerResponse`
- New struct `SapMonitorsClientCreateResponse`
- New struct `SapMonitorsClientCreateResult`
- New struct `SapMonitorsClientDeletePoller`
- New struct `SapMonitorsClientDeletePollerResponse`
- New struct `SapMonitorsClientDeleteResponse`
- New struct `SapMonitorsClientGetOptions`
- New struct `SapMonitorsClientGetResponse`
- New struct `SapMonitorsClientGetResult`
- New struct `SapMonitorsClientListOptions`
- New struct `SapMonitorsClientListPager`
- New struct `SapMonitorsClientListResponse`
- New struct `SapMonitorsClientListResult`
- New struct `SapMonitorsClientUpdateOptions`
- New struct `SapMonitorsClientUpdateResponse`
- New struct `SapMonitorsClientUpdateResult`
- New field `Type` in struct `ProviderInstance`
- New field `ID` in struct `ProviderInstance`
- New field `Name` in struct `ProviderInstance`
- New field `Error` in struct `ErrorResponse`
- New field `Type` in struct `ProxyResource`
- New field `ID` in struct `ProxyResource`
- New field `Name` in struct `ProxyResource`
- New field `Name` in struct `TrackedResource`
- New field `Type` in struct `TrackedResource`
- New field `ID` in struct `TrackedResource`
- New field `Name` in struct `SapMonitor`
- New field `Type` in struct `SapMonitor`
- New field `Location` in struct `SapMonitor`
- New field `Tags` in struct `SapMonitor`
- New field `ID` in struct `SapMonitor`


## 0.1.0 (2021-12-07)

- Init release.
