package powerbidedicatedapi

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/services/powerbidedicated/mgmt/2021-01-01/powerbidedicated"
	"github.com/Azure/go-autorest/autorest"
)

// CapacitiesClientAPI contains the set of methods on the CapacitiesClient type.
type CapacitiesClientAPI interface {
	CheckNameAvailability(ctx context.Context, location string, capacityParameters powerbidedicated.CheckCapacityNameAvailabilityParameters) (result powerbidedicated.CheckCapacityNameAvailabilityResult, err error)
	Create(ctx context.Context, resourceGroupName string, dedicatedCapacityName string, capacityParameters powerbidedicated.DedicatedCapacity) (result powerbidedicated.CapacitiesCreateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, dedicatedCapacityName string) (result powerbidedicated.CapacitiesDeleteFuture, err error)
	GetDetails(ctx context.Context, resourceGroupName string, dedicatedCapacityName string) (result powerbidedicated.DedicatedCapacity, err error)
	List(ctx context.Context) (result powerbidedicated.DedicatedCapacities, err error)
	ListByResourceGroup(ctx context.Context, resourceGroupName string) (result powerbidedicated.DedicatedCapacities, err error)
	ListSkus(ctx context.Context) (result powerbidedicated.SkuEnumerationForNewResourceResult, err error)
	ListSkusForCapacity(ctx context.Context, resourceGroupName string, dedicatedCapacityName string) (result powerbidedicated.SkuEnumerationForExistingResourceResult, err error)
	Resume(ctx context.Context, resourceGroupName string, dedicatedCapacityName string) (result powerbidedicated.CapacitiesResumeFuture, err error)
	Suspend(ctx context.Context, resourceGroupName string, dedicatedCapacityName string) (result powerbidedicated.CapacitiesSuspendFuture, err error)
	Update(ctx context.Context, resourceGroupName string, dedicatedCapacityName string, capacityUpdateParameters powerbidedicated.DedicatedCapacityUpdateParameters) (result powerbidedicated.CapacitiesUpdateFuture, err error)
}

var _ CapacitiesClientAPI = (*powerbidedicated.CapacitiesClient)(nil)

// OperationsClientAPI contains the set of methods on the OperationsClient type.
type OperationsClientAPI interface {
	List(ctx context.Context) (result powerbidedicated.OperationListResultPage, err error)
	ListComplete(ctx context.Context) (result powerbidedicated.OperationListResultIterator, err error)
}

var _ OperationsClientAPI = (*powerbidedicated.OperationsClient)(nil)

// AutoScaleVCoresClientAPI contains the set of methods on the AutoScaleVCoresClient type.
type AutoScaleVCoresClientAPI interface {
	Create(ctx context.Context, resourceGroupName string, vcoreName string, vCoreParameters powerbidedicated.AutoScaleVCore) (result powerbidedicated.AutoScaleVCore, err error)
	Delete(ctx context.Context, resourceGroupName string, vcoreName string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, vcoreName string) (result powerbidedicated.AutoScaleVCore, err error)
	ListByResourceGroup(ctx context.Context, resourceGroupName string) (result powerbidedicated.AutoScaleVCoreListResult, err error)
	ListBySubscription(ctx context.Context) (result powerbidedicated.AutoScaleVCoreListResult, err error)
	Update(ctx context.Context, resourceGroupName string, vcoreName string, vCoreUpdateParameters powerbidedicated.AutoScaleVCoreUpdateParameters) (result powerbidedicated.AutoScaleVCore, err error)
}

var _ AutoScaleVCoresClientAPI = (*powerbidedicated.AutoScaleVCoresClient)(nil)
