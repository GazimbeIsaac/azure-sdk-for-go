//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armhybridcompute

import (
	"encoding/json"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"reflect"
)

// MarshalJSON implements the json.Marshaller interface for type AgentConfiguration.
func (a AgentConfiguration) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "incomingConnectionsPorts", a.IncomingConnectionsPorts)
	populate(objectMap, "proxyUrl", a.ProxyURL)
	return json.Marshal(objectMap)
}

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

// MarshalJSON implements the json.Marshaller interface for type Machine.
func (m Machine) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "id", m.ID)
	populate(objectMap, "identity", m.Identity)
	populate(objectMap, "location", m.Location)
	populate(objectMap, "name", m.Name)
	populate(objectMap, "properties", m.Properties)
	populate(objectMap, "systemData", m.SystemData)
	populate(objectMap, "tags", m.Tags)
	populate(objectMap, "type", m.Type)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type MachineExtension.
func (m MachineExtension) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "id", m.ID)
	populate(objectMap, "location", m.Location)
	populate(objectMap, "name", m.Name)
	populate(objectMap, "properties", m.Properties)
	populate(objectMap, "systemData", m.SystemData)
	populate(objectMap, "tags", m.Tags)
	populate(objectMap, "type", m.Type)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type MachineExtensionInstanceViewStatus.
func (m MachineExtensionInstanceViewStatus) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "code", m.Code)
	populate(objectMap, "displayStatus", m.DisplayStatus)
	populate(objectMap, "level", m.Level)
	populate(objectMap, "message", m.Message)
	populateTimeRFC3339(objectMap, "time", m.Time)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type MachineExtensionInstanceViewStatus.
func (m *MachineExtensionInstanceViewStatus) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "code":
			err = unpopulate(val, &m.Code)
			delete(rawMsg, key)
		case "displayStatus":
			err = unpopulate(val, &m.DisplayStatus)
			delete(rawMsg, key)
		case "level":
			err = unpopulate(val, &m.Level)
			delete(rawMsg, key)
		case "message":
			err = unpopulate(val, &m.Message)
			delete(rawMsg, key)
		case "time":
			err = unpopulateTimeRFC3339(val, &m.Time)
			delete(rawMsg, key)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type MachineExtensionUpdate.
func (m MachineExtensionUpdate) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "properties", m.Properties)
	populate(objectMap, "tags", m.Tags)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type MachineExtensionUpgrade.
func (m MachineExtensionUpgrade) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "extensionTargets", m.ExtensionTargets)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type MachineExtensionsListResult.
func (m MachineExtensionsListResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "nextLink", m.NextLink)
	populate(objectMap, "value", m.Value)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type MachineListResult.
func (m MachineListResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "nextLink", m.NextLink)
	populate(objectMap, "value", m.Value)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type MachineProperties.
func (m MachineProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "adFqdn", m.AdFqdn)
	populate(objectMap, "agentConfiguration", m.AgentConfiguration)
	populate(objectMap, "agentVersion", m.AgentVersion)
	populate(objectMap, "clientPublicKey", m.ClientPublicKey)
	populate(objectMap, "cloudMetadata", m.CloudMetadata)
	populate(objectMap, "dnsFqdn", m.DNSFqdn)
	populate(objectMap, "detectedProperties", m.DetectedProperties)
	populate(objectMap, "displayName", m.DisplayName)
	populate(objectMap, "domainName", m.DomainName)
	populate(objectMap, "errorDetails", m.ErrorDetails)
	populate(objectMap, "extensions", m.Extensions)
	populateTimeRFC3339(objectMap, "lastStatusChange", m.LastStatusChange)
	populate(objectMap, "locationData", m.LocationData)
	populate(objectMap, "machineFqdn", m.MachineFqdn)
	populate(objectMap, "mssqlDiscovered", m.MssqlDiscovered)
	populate(objectMap, "osName", m.OSName)
	populate(objectMap, "osProfile", m.OSProfile)
	populate(objectMap, "osSku", m.OSSKU)
	populate(objectMap, "osType", m.OSType)
	populate(objectMap, "osVersion", m.OSVersion)
	populate(objectMap, "parentClusterResourceId", m.ParentClusterResourceID)
	populate(objectMap, "privateLinkScopeResourceId", m.PrivateLinkScopeResourceID)
	populate(objectMap, "provisioningState", m.ProvisioningState)
	populate(objectMap, "status", m.Status)
	populate(objectMap, "vmId", m.VMID)
	populate(objectMap, "vmUuid", m.VMUUID)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type MachineProperties.
func (m *MachineProperties) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "adFqdn":
			err = unpopulate(val, &m.AdFqdn)
			delete(rawMsg, key)
		case "agentConfiguration":
			err = unpopulate(val, &m.AgentConfiguration)
			delete(rawMsg, key)
		case "agentVersion":
			err = unpopulate(val, &m.AgentVersion)
			delete(rawMsg, key)
		case "clientPublicKey":
			err = unpopulate(val, &m.ClientPublicKey)
			delete(rawMsg, key)
		case "cloudMetadata":
			err = unpopulate(val, &m.CloudMetadata)
			delete(rawMsg, key)
		case "dnsFqdn":
			err = unpopulate(val, &m.DNSFqdn)
			delete(rawMsg, key)
		case "detectedProperties":
			err = unpopulate(val, &m.DetectedProperties)
			delete(rawMsg, key)
		case "displayName":
			err = unpopulate(val, &m.DisplayName)
			delete(rawMsg, key)
		case "domainName":
			err = unpopulate(val, &m.DomainName)
			delete(rawMsg, key)
		case "errorDetails":
			err = unpopulate(val, &m.ErrorDetails)
			delete(rawMsg, key)
		case "extensions":
			err = unpopulate(val, &m.Extensions)
			delete(rawMsg, key)
		case "lastStatusChange":
			err = unpopulateTimeRFC3339(val, &m.LastStatusChange)
			delete(rawMsg, key)
		case "locationData":
			err = unpopulate(val, &m.LocationData)
			delete(rawMsg, key)
		case "machineFqdn":
			err = unpopulate(val, &m.MachineFqdn)
			delete(rawMsg, key)
		case "mssqlDiscovered":
			err = unpopulate(val, &m.MssqlDiscovered)
			delete(rawMsg, key)
		case "osName":
			err = unpopulate(val, &m.OSName)
			delete(rawMsg, key)
		case "osProfile":
			err = unpopulate(val, &m.OSProfile)
			delete(rawMsg, key)
		case "osSku":
			err = unpopulate(val, &m.OSSKU)
			delete(rawMsg, key)
		case "osType":
			err = unpopulate(val, &m.OSType)
			delete(rawMsg, key)
		case "osVersion":
			err = unpopulate(val, &m.OSVersion)
			delete(rawMsg, key)
		case "parentClusterResourceId":
			err = unpopulate(val, &m.ParentClusterResourceID)
			delete(rawMsg, key)
		case "privateLinkScopeResourceId":
			err = unpopulate(val, &m.PrivateLinkScopeResourceID)
			delete(rawMsg, key)
		case "provisioningState":
			err = unpopulate(val, &m.ProvisioningState)
			delete(rawMsg, key)
		case "status":
			err = unpopulate(val, &m.Status)
			delete(rawMsg, key)
		case "vmId":
			err = unpopulate(val, &m.VMID)
			delete(rawMsg, key)
		case "vmUuid":
			err = unpopulate(val, &m.VMUUID)
			delete(rawMsg, key)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type MachineUpdate.
func (m MachineUpdate) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "identity", m.Identity)
	populate(objectMap, "properties", m.Properties)
	populate(objectMap, "tags", m.Tags)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type OperationListResult.
func (o OperationListResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "value", o.Value)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type PrivateEndpointConnectionListResult.
func (p PrivateEndpointConnectionListResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "nextLink", p.NextLink)
	populate(objectMap, "value", p.Value)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type PrivateLinkResourceListResult.
func (p PrivateLinkResourceListResult) MarshalJSON() ([]byte, error) {
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
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type PrivateLinkScope.
func (p PrivateLinkScope) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "id", p.ID)
	populate(objectMap, "location", p.Location)
	populate(objectMap, "name", p.Name)
	populate(objectMap, "properties", p.Properties)
	populate(objectMap, "systemData", p.SystemData)
	populate(objectMap, "tags", p.Tags)
	populate(objectMap, "type", p.Type)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type PrivateLinkScopeListResult.
func (p PrivateLinkScopeListResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "nextLink", p.NextLink)
	populate(objectMap, "value", p.Value)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type PrivateLinkScopeProperties.
func (p PrivateLinkScopeProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "privateEndpointConnections", p.PrivateEndpointConnections)
	populate(objectMap, "privateLinkScopeId", p.PrivateLinkScopeID)
	populate(objectMap, "provisioningState", p.ProvisioningState)
	populate(objectMap, "publicNetworkAccess", p.PublicNetworkAccess)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type PrivateLinkScopeValidationDetails.
func (p PrivateLinkScopeValidationDetails) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "connectionDetails", p.ConnectionDetails)
	populate(objectMap, "id", p.ID)
	populate(objectMap, "publicNetworkAccess", p.PublicNetworkAccess)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type PrivateLinkScopesResource.
func (p PrivateLinkScopesResource) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "id", p.ID)
	populate(objectMap, "location", p.Location)
	populate(objectMap, "name", p.Name)
	populate(objectMap, "tags", p.Tags)
	populate(objectMap, "type", p.Type)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type ResourceUpdate.
func (r ResourceUpdate) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "tags", r.Tags)
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

// MarshalJSON implements the json.Marshaller interface for type TagsResource.
func (t TagsResource) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "tags", t.Tags)
	return json.Marshal(objectMap)
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
