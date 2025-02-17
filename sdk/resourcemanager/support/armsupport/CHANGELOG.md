# Release History

## 0.4.0 (2022-04-18)
### Breaking Changes

- Function `*CommunicationsClient.List` has been removed
- Function `*ProblemClassificationsClient.List` has been removed
- Function `*ServicesClient.List` has been removed
- Function `*OperationsClient.List` has been removed
- Function `*TicketsClient.List` has been removed

### Features Added

- New function `*CommunicationsClient.NewListPager(string, *CommunicationsClientListOptions) *runtime.Pager[CommunicationsClientListResponse]`
- New function `*TicketsClient.NewListPager(*TicketsClientListOptions) *runtime.Pager[TicketsClientListResponse]`
- New function `*OperationsClient.NewListPager(*OperationsClientListOptions) *runtime.Pager[OperationsClientListResponse]`
- New function `*ProblemClassificationsClient.NewListPager(string, *ProblemClassificationsClientListOptions) *runtime.Pager[ProblemClassificationsClientListResponse]`
- New function `*ServicesClient.NewListPager(*ServicesClientListOptions) *runtime.Pager[ServicesClientListResponse]`


## 0.3.0 (2022-04-13)
### Breaking Changes

- Function `*ProblemClassificationsClient.List` parameter(s) have been changed from `(context.Context, string, *ProblemClassificationsClientListOptions)` to `(string, *ProblemClassificationsClientListOptions)`
- Function `*ProblemClassificationsClient.List` return value(s) have been changed from `(ProblemClassificationsClientListResponse, error)` to `(*runtime.Pager[ProblemClassificationsClientListResponse])`
- Function `NewTicketsClient` return value(s) have been changed from `(*TicketsClient)` to `(*TicketsClient, error)`
- Function `NewProblemClassificationsClient` return value(s) have been changed from `(*ProblemClassificationsClient)` to `(*ProblemClassificationsClient, error)`
- Function `*CommunicationsClient.List` return value(s) have been changed from `(*CommunicationsClientListPager)` to `(*runtime.Pager[CommunicationsClientListResponse])`
- Function `*CommunicationsClient.BeginCreate` return value(s) have been changed from `(CommunicationsClientCreatePollerResponse, error)` to `(*armruntime.Poller[CommunicationsClientCreateResponse], error)`
- Function `*ServicesClient.List` parameter(s) have been changed from `(context.Context, *ServicesClientListOptions)` to `(*ServicesClientListOptions)`
- Function `*ServicesClient.List` return value(s) have been changed from `(ServicesClientListResponse, error)` to `(*runtime.Pager[ServicesClientListResponse])`
- Function `*OperationsClient.List` parameter(s) have been changed from `(context.Context, *OperationsClientListOptions)` to `(*OperationsClientListOptions)`
- Function `*OperationsClient.List` return value(s) have been changed from `(OperationsClientListResponse, error)` to `(*runtime.Pager[OperationsClientListResponse])`
- Function `*TicketsClient.List` return value(s) have been changed from `(*TicketsClientListPager)` to `(*runtime.Pager[TicketsClientListResponse])`
- Function `NewCommunicationsClient` return value(s) have been changed from `(*CommunicationsClient)` to `(*CommunicationsClient, error)`
- Function `NewOperationsClient` return value(s) have been changed from `(*OperationsClient)` to `(*OperationsClient, error)`
- Function `NewServicesClient` return value(s) have been changed from `(*ServicesClient)` to `(*ServicesClient, error)`
- Function `*TicketsClient.BeginCreate` return value(s) have been changed from `(TicketsClientCreatePollerResponse, error)` to `(*armruntime.Poller[TicketsClientCreateResponse], error)`
- Function `Type.ToPtr` has been removed
- Function `SeverityLevel.ToPtr` has been removed
- Function `*CommunicationsClientListPager.PageResponse` has been removed
- Function `Status.ToPtr` has been removed
- Function `CommunicationType.ToPtr` has been removed
- Function `*CommunicationsClientCreatePoller.ResumeToken` has been removed
- Function `CommunicationsClientCreatePollerResponse.PollUntilDone` has been removed
- Function `*TicketsClientCreatePollerResponse.Resume` has been removed
- Function `*CommunicationsClientCreatePoller.FinalResponse` has been removed
- Function `*TicketsClientListPager.NextPage` has been removed
- Function `*CommunicationsClientCreatePollerResponse.Resume` has been removed
- Function `*TicketsClientCreatePoller.Done` has been removed
- Function `*TicketsClientCreatePoller.FinalResponse` has been removed
- Function `*TicketsClientListPager.Err` has been removed
- Function `*CommunicationsClientListPager.Err` has been removed
- Function `*TicketsClientListPager.PageResponse` has been removed
- Function `*TicketsClientCreatePoller.ResumeToken` has been removed
- Function `*TicketsClientCreatePoller.Poll` has been removed
- Function `CommunicationDirection.ToPtr` has been removed
- Function `PreferredContactMethod.ToPtr` has been removed
- Function `*CommunicationsClientCreatePoller.Done` has been removed
- Function `*CommunicationsClientListPager.NextPage` has been removed
- Function `*CommunicationsClientCreatePoller.Poll` has been removed
- Function `TicketsClientCreatePollerResponse.PollUntilDone` has been removed
- Struct `CommunicationsClientCheckNameAvailabilityResult` has been removed
- Struct `CommunicationsClientCreatePoller` has been removed
- Struct `CommunicationsClientCreatePollerResponse` has been removed
- Struct `CommunicationsClientCreateResult` has been removed
- Struct `CommunicationsClientGetResult` has been removed
- Struct `CommunicationsClientListPager` has been removed
- Struct `CommunicationsClientListResult` has been removed
- Struct `OperationsClientListResult` has been removed
- Struct `ProblemClassificationsClientGetResult` has been removed
- Struct `ProblemClassificationsClientListResult` has been removed
- Struct `ServicesClientGetResult` has been removed
- Struct `ServicesClientListResult` has been removed
- Struct `TicketsClientCheckNameAvailabilityResult` has been removed
- Struct `TicketsClientCreatePoller` has been removed
- Struct `TicketsClientCreatePollerResponse` has been removed
- Struct `TicketsClientCreateResult` has been removed
- Struct `TicketsClientGetResult` has been removed
- Struct `TicketsClientListPager` has been removed
- Struct `TicketsClientListResult` has been removed
- Struct `TicketsClientUpdateResult` has been removed
- Field `ServicesClientGetResult` of struct `ServicesClientGetResponse` has been removed
- Field `RawResponse` of struct `ServicesClientGetResponse` has been removed
- Field `ProblemClassificationsClientListResult` of struct `ProblemClassificationsClientListResponse` has been removed
- Field `RawResponse` of struct `ProblemClassificationsClientListResponse` has been removed
- Field `ProblemClassificationsClientGetResult` of struct `ProblemClassificationsClientGetResponse` has been removed
- Field `RawResponse` of struct `ProblemClassificationsClientGetResponse` has been removed
- Field `TicketsClientCreateResult` of struct `TicketsClientCreateResponse` has been removed
- Field `RawResponse` of struct `TicketsClientCreateResponse` has been removed
- Field `TicketsClientListResult` of struct `TicketsClientListResponse` has been removed
- Field `RawResponse` of struct `TicketsClientListResponse` has been removed
- Field `CommunicationsClientCheckNameAvailabilityResult` of struct `CommunicationsClientCheckNameAvailabilityResponse` has been removed
- Field `RawResponse` of struct `CommunicationsClientCheckNameAvailabilityResponse` has been removed
- Field `TicketsClientCheckNameAvailabilityResult` of struct `TicketsClientCheckNameAvailabilityResponse` has been removed
- Field `RawResponse` of struct `TicketsClientCheckNameAvailabilityResponse` has been removed
- Field `OperationsClientListResult` of struct `OperationsClientListResponse` has been removed
- Field `RawResponse` of struct `OperationsClientListResponse` has been removed
- Field `CommunicationsClientCreateResult` of struct `CommunicationsClientCreateResponse` has been removed
- Field `RawResponse` of struct `CommunicationsClientCreateResponse` has been removed
- Field `TicketsClientUpdateResult` of struct `TicketsClientUpdateResponse` has been removed
- Field `RawResponse` of struct `TicketsClientUpdateResponse` has been removed
- Field `CommunicationsClientGetResult` of struct `CommunicationsClientGetResponse` has been removed
- Field `RawResponse` of struct `CommunicationsClientGetResponse` has been removed
- Field `CommunicationsClientListResult` of struct `CommunicationsClientListResponse` has been removed
- Field `RawResponse` of struct `CommunicationsClientListResponse` has been removed
- Field `ServicesClientListResult` of struct `ServicesClientListResponse` has been removed
- Field `RawResponse` of struct `ServicesClientListResponse` has been removed
- Field `TicketsClientGetResult` of struct `TicketsClientGetResponse` has been removed
- Field `RawResponse` of struct `TicketsClientGetResponse` has been removed

### Features Added

- New anonymous field `ProblemClassificationsListResult` in struct `ProblemClassificationsClientListResponse`
- New anonymous field `CommunicationDetails` in struct `CommunicationsClientGetResponse`
- New field `ResumeToken` in struct `TicketsClientBeginCreateOptions`
- New anonymous field `TicketDetails` in struct `TicketsClientUpdateResponse`
- New anonymous field `TicketDetails` in struct `TicketsClientCreateResponse`
- New anonymous field `Service` in struct `ServicesClientGetResponse`
- New anonymous field `CheckNameAvailabilityOutput` in struct `TicketsClientCheckNameAvailabilityResponse`
- New field `ResumeToken` in struct `CommunicationsClientBeginCreateOptions`
- New anonymous field `CheckNameAvailabilityOutput` in struct `CommunicationsClientCheckNameAvailabilityResponse`
- New anonymous field `ProblemClassification` in struct `ProblemClassificationsClientGetResponse`
- New anonymous field `TicketsListResult` in struct `TicketsClientListResponse`
- New anonymous field `CommunicationDetails` in struct `CommunicationsClientCreateResponse`
- New anonymous field `OperationsListResult` in struct `OperationsClientListResponse`
- New anonymous field `TicketDetails` in struct `TicketsClientGetResponse`
- New anonymous field `CommunicationsListResult` in struct `CommunicationsClientListResponse`
- New anonymous field `ServicesListResult` in struct `ServicesClientListResponse`


## 0.2.1 (2022-02-22)

### Other Changes

- Remove the go_mod_tidy_hack.go file.

## 0.2.0 (2022-01-13)
### Breaking Changes

- Function `*OperationsClient.List` parameter(s) have been changed from `(context.Context, *OperationsListOptions)` to `(context.Context, *OperationsClientListOptions)`
- Function `*OperationsClient.List` return value(s) have been changed from `(OperationsListResponse, error)` to `(OperationsClientListResponse, error)`
- Function `*CommunicationsClient.CheckNameAvailability` parameter(s) have been changed from `(context.Context, string, CheckNameAvailabilityInput, *CommunicationsCheckNameAvailabilityOptions)` to `(context.Context, string, CheckNameAvailabilityInput, *CommunicationsClientCheckNameAvailabilityOptions)`
- Function `*CommunicationsClient.CheckNameAvailability` return value(s) have been changed from `(CommunicationsCheckNameAvailabilityResponse, error)` to `(CommunicationsClientCheckNameAvailabilityResponse, error)`
- Function `*CommunicationsClient.Get` parameter(s) have been changed from `(context.Context, string, string, *CommunicationsGetOptions)` to `(context.Context, string, string, *CommunicationsClientGetOptions)`
- Function `*CommunicationsClient.Get` return value(s) have been changed from `(CommunicationsGetResponse, error)` to `(CommunicationsClientGetResponse, error)`
- Function `*ServicesClient.List` parameter(s) have been changed from `(context.Context, *ServicesListOptions)` to `(context.Context, *ServicesClientListOptions)`
- Function `*ServicesClient.List` return value(s) have been changed from `(ServicesListResponse, error)` to `(ServicesClientListResponse, error)`
- Function `*CommunicationsClient.List` parameter(s) have been changed from `(string, *CommunicationsListOptions)` to `(string, *CommunicationsClientListOptions)`
- Function `*CommunicationsClient.List` return value(s) have been changed from `(*CommunicationsListPager)` to `(*CommunicationsClientListPager)`
- Function `*ProblemClassificationsClient.List` parameter(s) have been changed from `(context.Context, string, *ProblemClassificationsListOptions)` to `(context.Context, string, *ProblemClassificationsClientListOptions)`
- Function `*ProblemClassificationsClient.List` return value(s) have been changed from `(ProblemClassificationsListResponse, error)` to `(ProblemClassificationsClientListResponse, error)`
- Function `*ServicesClient.Get` parameter(s) have been changed from `(context.Context, string, *ServicesGetOptions)` to `(context.Context, string, *ServicesClientGetOptions)`
- Function `*ServicesClient.Get` return value(s) have been changed from `(ServicesGetResponse, error)` to `(ServicesClientGetResponse, error)`
- Function `*ProblemClassificationsClient.Get` parameter(s) have been changed from `(context.Context, string, string, *ProblemClassificationsGetOptions)` to `(context.Context, string, string, *ProblemClassificationsClientGetOptions)`
- Function `*ProblemClassificationsClient.Get` return value(s) have been changed from `(ProblemClassificationsGetResponse, error)` to `(ProblemClassificationsClientGetResponse, error)`
- Function `*CommunicationsClient.BeginCreate` parameter(s) have been changed from `(context.Context, string, string, CommunicationDetails, *CommunicationsBeginCreateOptions)` to `(context.Context, string, string, CommunicationDetails, *CommunicationsClientBeginCreateOptions)`
- Function `*CommunicationsClient.BeginCreate` return value(s) have been changed from `(CommunicationsCreatePollerResponse, error)` to `(CommunicationsClientCreatePollerResponse, error)`
- Function `SupportTicketsListResult.MarshalJSON` has been removed
- Function `SupportTicketsCreatePollerResponse.PollUntilDone` has been removed
- Function `*SupportTicketsCreatePoller.FinalResponse` has been removed
- Function `*CommunicationsListPager.PageResponse` has been removed
- Function `SupportTicketDetailsProperties.MarshalJSON` has been removed
- Function `*SupportTicketsCreatePollerResponse.Resume` has been removed
- Function `ExceptionResponse.Error` has been removed
- Function `*SupportTicketsCreatePoller.ResumeToken` has been removed
- Function `*CommunicationsCreatePoller.Poll` has been removed
- Function `*SupportTicketsClient.BeginCreate` has been removed
- Function `NewSupportTicketsClient` has been removed
- Function `*SupportTicketsCreatePoller.Poll` has been removed
- Function `*CommunicationsListPager.Err` has been removed
- Function `*CommunicationsListPager.NextPage` has been removed
- Function `*SupportTicketsListPager.PageResponse` has been removed
- Function `*CommunicationsCreatePoller.FinalResponse` has been removed
- Function `*SupportTicketsListPager.Err` has been removed
- Function `*SupportTicketDetailsProperties.UnmarshalJSON` has been removed
- Function `*SupportTicketsClient.CheckNameAvailability` has been removed
- Function `*CommunicationsCreatePoller.ResumeToken` has been removed
- Function `*CommunicationsCreatePoller.Done` has been removed
- Function `*CommunicationsCreatePollerResponse.Resume` has been removed
- Function `*SupportTicketsListPager.NextPage` has been removed
- Function `CommunicationsCreatePollerResponse.PollUntilDone` has been removed
- Function `*SupportTicketsClient.Get` has been removed
- Function `*SupportTicketsClient.List` has been removed
- Function `*SupportTicketsClient.Update` has been removed
- Function `*SupportTicketsCreatePoller.Done` has been removed
- Struct `CommunicationsBeginCreateOptions` has been removed
- Struct `CommunicationsCheckNameAvailabilityOptions` has been removed
- Struct `CommunicationsCheckNameAvailabilityResponse` has been removed
- Struct `CommunicationsCheckNameAvailabilityResult` has been removed
- Struct `CommunicationsCreatePoller` has been removed
- Struct `CommunicationsCreatePollerResponse` has been removed
- Struct `CommunicationsCreateResponse` has been removed
- Struct `CommunicationsCreateResult` has been removed
- Struct `CommunicationsGetOptions` has been removed
- Struct `CommunicationsGetResponse` has been removed
- Struct `CommunicationsGetResult` has been removed
- Struct `CommunicationsListOptions` has been removed
- Struct `CommunicationsListPager` has been removed
- Struct `CommunicationsListResponse` has been removed
- Struct `CommunicationsListResultEnvelope` has been removed
- Struct `OperationsListOptions` has been removed
- Struct `OperationsListResponse` has been removed
- Struct `OperationsListResultEnvelope` has been removed
- Struct `ProblemClassificationsGetOptions` has been removed
- Struct `ProblemClassificationsGetResponse` has been removed
- Struct `ProblemClassificationsGetResult` has been removed
- Struct `ProblemClassificationsListOptions` has been removed
- Struct `ProblemClassificationsListResponse` has been removed
- Struct `ProblemClassificationsListResultEnvelope` has been removed
- Struct `ServicesGetOptions` has been removed
- Struct `ServicesGetResponse` has been removed
- Struct `ServicesGetResult` has been removed
- Struct `ServicesListOptions` has been removed
- Struct `ServicesListResponse` has been removed
- Struct `ServicesListResultEnvelope` has been removed
- Struct `SupportEngineer` has been removed
- Struct `SupportTicketDetails` has been removed
- Struct `SupportTicketDetailsProperties` has been removed
- Struct `SupportTicketsBeginCreateOptions` has been removed
- Struct `SupportTicketsCheckNameAvailabilityOptions` has been removed
- Struct `SupportTicketsCheckNameAvailabilityResponse` has been removed
- Struct `SupportTicketsCheckNameAvailabilityResult` has been removed
- Struct `SupportTicketsClient` has been removed
- Struct `SupportTicketsCreatePoller` has been removed
- Struct `SupportTicketsCreatePollerResponse` has been removed
- Struct `SupportTicketsCreateResponse` has been removed
- Struct `SupportTicketsCreateResult` has been removed
- Struct `SupportTicketsGetOptions` has been removed
- Struct `SupportTicketsGetResponse` has been removed
- Struct `SupportTicketsGetResult` has been removed
- Struct `SupportTicketsListOptions` has been removed
- Struct `SupportTicketsListPager` has been removed
- Struct `SupportTicketsListResponse` has been removed
- Struct `SupportTicketsListResult` has been removed
- Struct `SupportTicketsListResultEnvelope` has been removed
- Struct `SupportTicketsUpdateOptions` has been removed
- Struct `SupportTicketsUpdateResponse` has been removed
- Struct `SupportTicketsUpdateResult` has been removed
- Field `InnerError` of struct `ExceptionResponse` has been removed

### Features Added

- New function `CommunicationsClientCreatePollerResponse.PollUntilDone(context.Context, time.Duration) (CommunicationsClientCreateResponse, error)`
- New function `NewTicketsClient(string, azcore.TokenCredential, *arm.ClientOptions) *TicketsClient`
- New function `*TicketsClient.CheckNameAvailability(context.Context, CheckNameAvailabilityInput, *TicketsClientCheckNameAvailabilityOptions) (TicketsClientCheckNameAvailabilityResponse, error)`
- New function `*CommunicationsClientListPager.Err() error`
- New function `*CommunicationsClientCreatePoller.Poll(context.Context) (*http.Response, error)`
- New function `*CommunicationsClientListPager.PageResponse() CommunicationsClientListResponse`
- New function `*TicketsClientListPager.PageResponse() TicketsClientListResponse`
- New function `*CommunicationsClientCreatePoller.Done() bool`
- New function `*TicketsClientCreatePollerResponse.Resume(context.Context, *TicketsClient, string) error`
- New function `*TicketDetailsProperties.UnmarshalJSON([]byte) error`
- New function `*TicketsClient.Update(context.Context, string, UpdateSupportTicket, *TicketsClientUpdateOptions) (TicketsClientUpdateResponse, error)`
- New function `*TicketsClient.BeginCreate(context.Context, string, TicketDetails, *TicketsClientBeginCreateOptions) (TicketsClientCreatePollerResponse, error)`
- New function `*TicketsClientCreatePoller.ResumeToken() (string, error)`
- New function `TicketDetailsProperties.MarshalJSON() ([]byte, error)`
- New function `TicketsListResult.MarshalJSON() ([]byte, error)`
- New function `*TicketsClientCreatePoller.Poll(context.Context) (*http.Response, error)`
- New function `TicketsClientCreatePollerResponse.PollUntilDone(context.Context, time.Duration) (TicketsClientCreateResponse, error)`
- New function `*TicketsClient.List(*TicketsClientListOptions) *TicketsClientListPager`
- New function `*CommunicationsClientListPager.NextPage(context.Context) bool`
- New function `*TicketsClient.Get(context.Context, string, *TicketsClientGetOptions) (TicketsClientGetResponse, error)`
- New function `*TicketsClientListPager.Err() error`
- New function `*CommunicationsClientCreatePoller.FinalResponse(context.Context) (CommunicationsClientCreateResponse, error)`
- New function `*TicketsClientCreatePoller.FinalResponse(context.Context) (TicketsClientCreateResponse, error)`
- New function `*CommunicationsClientCreatePollerResponse.Resume(context.Context, *CommunicationsClient, string) error`
- New function `*TicketsClientListPager.NextPage(context.Context) bool`
- New function `*CommunicationsClientCreatePoller.ResumeToken() (string, error)`
- New function `*TicketsClientCreatePoller.Done() bool`
- New struct `CommunicationsClientBeginCreateOptions`
- New struct `CommunicationsClientCheckNameAvailabilityOptions`
- New struct `CommunicationsClientCheckNameAvailabilityResponse`
- New struct `CommunicationsClientCheckNameAvailabilityResult`
- New struct `CommunicationsClientCreatePoller`
- New struct `CommunicationsClientCreatePollerResponse`
- New struct `CommunicationsClientCreateResponse`
- New struct `CommunicationsClientCreateResult`
- New struct `CommunicationsClientGetOptions`
- New struct `CommunicationsClientGetResponse`
- New struct `CommunicationsClientGetResult`
- New struct `CommunicationsClientListOptions`
- New struct `CommunicationsClientListPager`
- New struct `CommunicationsClientListResponse`
- New struct `CommunicationsClientListResult`
- New struct `Engineer`
- New struct `OperationsClientListOptions`
- New struct `OperationsClientListResponse`
- New struct `OperationsClientListResult`
- New struct `ProblemClassificationsClientGetOptions`
- New struct `ProblemClassificationsClientGetResponse`
- New struct `ProblemClassificationsClientGetResult`
- New struct `ProblemClassificationsClientListOptions`
- New struct `ProblemClassificationsClientListResponse`
- New struct `ProblemClassificationsClientListResult`
- New struct `ServicesClientGetOptions`
- New struct `ServicesClientGetResponse`
- New struct `ServicesClientGetResult`
- New struct `ServicesClientListOptions`
- New struct `ServicesClientListResponse`
- New struct `ServicesClientListResult`
- New struct `TicketDetails`
- New struct `TicketDetailsProperties`
- New struct `TicketsClient`
- New struct `TicketsClientBeginCreateOptions`
- New struct `TicketsClientCheckNameAvailabilityOptions`
- New struct `TicketsClientCheckNameAvailabilityResponse`
- New struct `TicketsClientCheckNameAvailabilityResult`
- New struct `TicketsClientCreatePoller`
- New struct `TicketsClientCreatePollerResponse`
- New struct `TicketsClientCreateResponse`
- New struct `TicketsClientCreateResult`
- New struct `TicketsClientGetOptions`
- New struct `TicketsClientGetResponse`
- New struct `TicketsClientGetResult`
- New struct `TicketsClientListOptions`
- New struct `TicketsClientListPager`
- New struct `TicketsClientListResponse`
- New struct `TicketsClientListResult`
- New struct `TicketsClientUpdateOptions`
- New struct `TicketsClientUpdateResponse`
- New struct `TicketsClientUpdateResult`
- New struct `TicketsListResult`
- New field `Error` in struct `ExceptionResponse`


## 0.1.0 (2021-12-22)

- Init release.
