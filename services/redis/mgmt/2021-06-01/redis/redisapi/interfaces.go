package redisapi

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/services/redis/mgmt/2021-06-01/redis"
	"github.com/Azure/go-autorest/autorest"
)

// OperationsClientAPI contains the set of methods on the OperationsClient type.
type OperationsClientAPI interface {
	List(ctx context.Context) (result redis.OperationListResultPage, err error)
	ListComplete(ctx context.Context) (result redis.OperationListResultIterator, err error)
}

var _ OperationsClientAPI = (*redis.OperationsClient)(nil)

// ClientAPI contains the set of methods on the Client type.
type ClientAPI interface {
	CheckNameAvailability(ctx context.Context, parameters redis.CheckNameAvailabilityParameters) (result autorest.Response, err error)
	Create(ctx context.Context, resourceGroupName string, name string, parameters redis.CreateParameters) (result redis.CreateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, name string) (result redis.DeleteFuture, err error)
	ExportData(ctx context.Context, resourceGroupName string, name string, parameters redis.ExportRDBParameters) (result redis.ExportDataFuture, err error)
	ForceReboot(ctx context.Context, resourceGroupName string, name string, parameters redis.RebootParameters) (result redis.ForceRebootResponse, err error)
	Get(ctx context.Context, resourceGroupName string, name string) (result redis.ResourceType, err error)
	ImportData(ctx context.Context, resourceGroupName string, name string, parameters redis.ImportRDBParameters) (result redis.ImportDataFuture, err error)
	ListByResourceGroup(ctx context.Context, resourceGroupName string) (result redis.ListResultPage, err error)
	ListByResourceGroupComplete(ctx context.Context, resourceGroupName string) (result redis.ListResultIterator, err error)
	ListBySubscription(ctx context.Context) (result redis.ListResultPage, err error)
	ListBySubscriptionComplete(ctx context.Context) (result redis.ListResultIterator, err error)
	ListKeys(ctx context.Context, resourceGroupName string, name string) (result redis.AccessKeys, err error)
	ListUpgradeNotifications(ctx context.Context, resourceGroupName string, name string, history float64) (result redis.NotificationListResponsePage, err error)
	ListUpgradeNotificationsComplete(ctx context.Context, resourceGroupName string, name string, history float64) (result redis.NotificationListResponseIterator, err error)
	RegenerateKey(ctx context.Context, resourceGroupName string, name string, parameters redis.RegenerateKeyParameters) (result redis.AccessKeys, err error)
	Update(ctx context.Context, resourceGroupName string, name string, parameters redis.UpdateParameters) (result redis.ResourceType, err error)
}

var _ ClientAPI = (*redis.Client)(nil)

// FirewallRulesClientAPI contains the set of methods on the FirewallRulesClient type.
type FirewallRulesClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, cacheName string, ruleName string, parameters redis.FirewallRule) (result redis.FirewallRule, err error)
	Delete(ctx context.Context, resourceGroupName string, cacheName string, ruleName string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, cacheName string, ruleName string) (result redis.FirewallRule, err error)
	List(ctx context.Context, resourceGroupName string, cacheName string) (result redis.FirewallRuleListResultPage, err error)
	ListComplete(ctx context.Context, resourceGroupName string, cacheName string) (result redis.FirewallRuleListResultIterator, err error)
}

var _ FirewallRulesClientAPI = (*redis.FirewallRulesClient)(nil)

// PatchSchedulesClientAPI contains the set of methods on the PatchSchedulesClient type.
type PatchSchedulesClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, name string, parameters redis.PatchSchedule) (result redis.PatchSchedule, err error)
	Delete(ctx context.Context, resourceGroupName string, name string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, name string) (result redis.PatchSchedule, err error)
	ListByRedisResource(ctx context.Context, resourceGroupName string, cacheName string) (result redis.PatchScheduleListResultPage, err error)
	ListByRedisResourceComplete(ctx context.Context, resourceGroupName string, cacheName string) (result redis.PatchScheduleListResultIterator, err error)
}

var _ PatchSchedulesClientAPI = (*redis.PatchSchedulesClient)(nil)

// LinkedServerClientAPI contains the set of methods on the LinkedServerClient type.
type LinkedServerClientAPI interface {
	Create(ctx context.Context, resourceGroupName string, name string, linkedServerName string, parameters redis.LinkedServerCreateParameters) (result redis.LinkedServerCreateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, name string, linkedServerName string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, name string, linkedServerName string) (result redis.LinkedServerWithProperties, err error)
	List(ctx context.Context, resourceGroupName string, name string) (result redis.LinkedServerWithPropertiesListPage, err error)
	ListComplete(ctx context.Context, resourceGroupName string, name string) (result redis.LinkedServerWithPropertiesListIterator, err error)
}

var _ LinkedServerClientAPI = (*redis.LinkedServerClient)(nil)

// PrivateEndpointConnectionsClientAPI contains the set of methods on the PrivateEndpointConnectionsClient type.
type PrivateEndpointConnectionsClientAPI interface {
	Delete(ctx context.Context, resourceGroupName string, cacheName string, privateEndpointConnectionName string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, cacheName string, privateEndpointConnectionName string) (result redis.PrivateEndpointConnection, err error)
	List(ctx context.Context, resourceGroupName string, cacheName string) (result redis.PrivateEndpointConnectionListResult, err error)
	Put(ctx context.Context, resourceGroupName string, cacheName string, privateEndpointConnectionName string, properties redis.PrivateEndpointConnection) (result redis.PrivateEndpointConnectionsPutFuture, err error)
}

var _ PrivateEndpointConnectionsClientAPI = (*redis.PrivateEndpointConnectionsClient)(nil)

// PrivateLinkResourcesClientAPI contains the set of methods on the PrivateLinkResourcesClient type.
type PrivateLinkResourcesClientAPI interface {
	ListByRedisCache(ctx context.Context, resourceGroupName string, cacheName string) (result redis.PrivateLinkResourceListResult, err error)
}

var _ PrivateLinkResourcesClientAPI = (*redis.PrivateLinkResourcesClient)(nil)

// AsyncOperationStatusClientAPI contains the set of methods on the AsyncOperationStatusClient type.
type AsyncOperationStatusClientAPI interface {
	Get(ctx context.Context, location string, operationID string) (result redis.OperationStatus, err error)
}

var _ AsyncOperationStatusClientAPI = (*redis.AsyncOperationStatusClient)(nil)
