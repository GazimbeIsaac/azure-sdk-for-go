//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armwebpubsub

import (
	"encoding/json"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"reflect"
)

// MarshalJSON implements the json.Marshaller interface for type ErrorDetail.
func (e ErrorDetail) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "additionalInfo", e.AdditionalInfo)
	populate(objectMap, "code", e.Code)
	populate(objectMap, "details", e.Details)
	populate(objectMap, "message", e.Message)
	populate(objectMap, "target", e.Target)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type EventHandler.
func (e EventHandler) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "auth", e.Auth)
	populate(objectMap, "systemEvents", e.SystemEvents)
	populate(objectMap, "urlTemplate", e.URLTemplate)
	populate(objectMap, "userEventPattern", e.UserEventPattern)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type HubList.
func (h HubList) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "nextLink", h.NextLink)
	populate(objectMap, "value", h.Value)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type HubProperties.
func (h HubProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "anonymousConnectPolicy", h.AnonymousConnectPolicy)
	populate(objectMap, "eventHandlers", h.EventHandlers)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type LiveTraceConfiguration.
func (l LiveTraceConfiguration) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "categories", l.Categories)
	populate(objectMap, "enabled", l.Enabled)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type ManagedIdentity.
func (m ManagedIdentity) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "principalId", m.PrincipalID)
	populate(objectMap, "tenantId", m.TenantID)
	populate(objectMap, "type", m.Type)
	populate(objectMap, "userAssignedIdentities", m.UserAssignedIdentities)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type MetricSpecification.
func (m MetricSpecification) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "aggregationType", m.AggregationType)
	populate(objectMap, "category", m.Category)
	populate(objectMap, "dimensions", m.Dimensions)
	populate(objectMap, "displayDescription", m.DisplayDescription)
	populate(objectMap, "displayName", m.DisplayName)
	populate(objectMap, "fillGapWithZero", m.FillGapWithZero)
	populate(objectMap, "name", m.Name)
	populate(objectMap, "unit", m.Unit)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type NetworkACL.
func (n NetworkACL) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "allow", n.Allow)
	populate(objectMap, "deny", n.Deny)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type NetworkACLs.
func (n NetworkACLs) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "defaultAction", n.DefaultAction)
	populate(objectMap, "privateEndpoints", n.PrivateEndpoints)
	populate(objectMap, "publicNetwork", n.PublicNetwork)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type OperationList.
func (o OperationList) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "nextLink", o.NextLink)
	populate(objectMap, "value", o.Value)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type PrivateEndpointACL.
func (p PrivateEndpointACL) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "allow", p.Allow)
	populate(objectMap, "deny", p.Deny)
	populate(objectMap, "name", p.Name)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type PrivateEndpointConnectionList.
func (p PrivateEndpointConnectionList) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "nextLink", p.NextLink)
	populate(objectMap, "value", p.Value)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type PrivateEndpointConnectionProperties.
func (p PrivateEndpointConnectionProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "groupIds", p.GroupIDs)
	populate(objectMap, "privateEndpoint", p.PrivateEndpoint)
	populate(objectMap, "privateLinkServiceConnectionState", p.PrivateLinkServiceConnectionState)
	populate(objectMap, "provisioningState", p.ProvisioningState)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type PrivateLinkResourceList.
func (p PrivateLinkResourceList) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "nextLink", p.NextLink)
	populate(objectMap, "value", p.Value)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type PrivateLinkResourceProperties.
func (p PrivateLinkResourceProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "groupId", p.GroupID)
	populate(objectMap, "requiredMembers", p.RequiredMembers)
	populate(objectMap, "requiredZoneNames", p.RequiredZoneNames)
	populate(objectMap, "shareablePrivateLinkResourceTypes", p.ShareablePrivateLinkResourceTypes)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type Properties.
func (p Properties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "disableAadAuth", p.DisableAADAuth)
	populate(objectMap, "disableLocalAuth", p.DisableLocalAuth)
	populate(objectMap, "externalIP", p.ExternalIP)
	populate(objectMap, "hostName", p.HostName)
	populate(objectMap, "hostNamePrefix", p.HostNamePrefix)
	populate(objectMap, "liveTraceConfiguration", p.LiveTraceConfiguration)
	populate(objectMap, "networkACLs", p.NetworkACLs)
	populate(objectMap, "privateEndpointConnections", p.PrivateEndpointConnections)
	populate(objectMap, "provisioningState", p.ProvisioningState)
	populate(objectMap, "publicNetworkAccess", p.PublicNetworkAccess)
	populate(objectMap, "publicPort", p.PublicPort)
	populate(objectMap, "resourceLogConfiguration", p.ResourceLogConfiguration)
	populate(objectMap, "serverPort", p.ServerPort)
	populate(objectMap, "sharedPrivateLinkResources", p.SharedPrivateLinkResources)
	populate(objectMap, "tls", p.TLS)
	populate(objectMap, "version", p.Version)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type ResourceInfo.
func (r ResourceInfo) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "id", r.ID)
	populate(objectMap, "identity", r.Identity)
	populate(objectMap, "location", r.Location)
	populate(objectMap, "name", r.Name)
	populate(objectMap, "properties", r.Properties)
	populate(objectMap, "sku", r.SKU)
	populate(objectMap, "systemData", r.SystemData)
	populate(objectMap, "tags", r.Tags)
	populate(objectMap, "type", r.Type)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type ResourceInfoList.
func (r ResourceInfoList) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "nextLink", r.NextLink)
	populate(objectMap, "value", r.Value)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type ResourceLogConfiguration.
func (r ResourceLogConfiguration) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "categories", r.Categories)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type SKUCapacity.
func (s SKUCapacity) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "allowedValues", s.AllowedValues)
	populate(objectMap, "default", s.Default)
	populate(objectMap, "maximum", s.Maximum)
	populate(objectMap, "minimum", s.Minimum)
	populate(objectMap, "scaleType", s.ScaleType)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type SKUList.
func (s SKUList) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "nextLink", s.NextLink)
	populate(objectMap, "value", s.Value)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type ServiceSpecification.
func (s ServiceSpecification) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "logSpecifications", s.LogSpecifications)
	populate(objectMap, "metricSpecifications", s.MetricSpecifications)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type SharedPrivateLinkResourceList.
func (s SharedPrivateLinkResourceList) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "nextLink", s.NextLink)
	populate(objectMap, "value", s.Value)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type SignalRServiceUsageList.
func (s SignalRServiceUsageList) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "nextLink", s.NextLink)
	populate(objectMap, "value", s.Value)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type SystemData.
func (s SystemData) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populateTimeRFC3339(objectMap, "createdAt", s.CreatedAt)
	populate(objectMap, "createdBy", s.CreatedBy)
	populate(objectMap, "createdByType", s.CreatedByType)
	populateTimeRFC3339(objectMap, "lastModifiedAt", s.LastModifiedAt)
	populate(objectMap, "lastModifiedBy", s.LastModifiedBy)
	populate(objectMap, "lastModifiedByType", s.LastModifiedByType)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type SystemData.
func (s *SystemData) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "createdAt":
			err = unpopulateTimeRFC3339(val, &s.CreatedAt)
			delete(rawMsg, key)
		case "createdBy":
			err = unpopulate(val, &s.CreatedBy)
			delete(rawMsg, key)
		case "createdByType":
			err = unpopulate(val, &s.CreatedByType)
			delete(rawMsg, key)
		case "lastModifiedAt":
			err = unpopulateTimeRFC3339(val, &s.LastModifiedAt)
			delete(rawMsg, key)
		case "lastModifiedBy":
			err = unpopulate(val, &s.LastModifiedBy)
			delete(rawMsg, key)
		case "lastModifiedByType":
			err = unpopulate(val, &s.LastModifiedByType)
			delete(rawMsg, key)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type TrackedResource.
func (t TrackedResource) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "id", t.ID)
	populate(objectMap, "location", t.Location)
	populate(objectMap, "name", t.Name)
	populate(objectMap, "tags", t.Tags)
	populate(objectMap, "type", t.Type)
	return json.Marshal(objectMap)
}

func populate(m map[string]interface{}, k string, v interface{}) {
	if v == nil {
		return
	} else if azcore.IsNullValue(v) {
		m[k] = nil
	} else if !reflect.ValueOf(v).IsNil() {
		m[k] = v
	}
}

func unpopulate(data json.RawMessage, v interface{}) error {
	if data == nil {
		return nil
	}
	return json.Unmarshal(data, v)
}
