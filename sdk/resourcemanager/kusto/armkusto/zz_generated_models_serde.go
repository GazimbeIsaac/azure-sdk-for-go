//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armkusto

import (
	"encoding/json"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"reflect"
)

// MarshalJSON implements the json.Marshaller interface for type AttachedDatabaseConfigurationListResult.
func (a AttachedDatabaseConfigurationListResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "value", a.Value)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type AttachedDatabaseConfigurationProperties.
func (a AttachedDatabaseConfigurationProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "attachedDatabaseNames", a.AttachedDatabaseNames)
	populate(objectMap, "clusterResourceId", a.ClusterResourceID)
	populate(objectMap, "databaseName", a.DatabaseName)
	populate(objectMap, "defaultPrincipalsModificationKind", a.DefaultPrincipalsModificationKind)
	populate(objectMap, "provisioningState", a.ProvisioningState)
	populate(objectMap, "tableLevelSharingProperties", a.TableLevelSharingProperties)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type CloudErrorBody.
func (c CloudErrorBody) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "code", c.Code)
	populate(objectMap, "details", c.Details)
	populate(objectMap, "message", c.Message)
	populate(objectMap, "target", c.Target)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type Cluster.
func (c Cluster) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "etag", c.Etag)
	populate(objectMap, "id", c.ID)
	populate(objectMap, "identity", c.Identity)
	populate(objectMap, "location", c.Location)
	populate(objectMap, "name", c.Name)
	populate(objectMap, "properties", c.Properties)
	populate(objectMap, "sku", c.SKU)
	populate(objectMap, "systemData", c.SystemData)
	populate(objectMap, "tags", c.Tags)
	populate(objectMap, "type", c.Type)
	populate(objectMap, "zones", c.Zones)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type ClusterListResult.
func (c ClusterListResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "value", c.Value)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type ClusterPrincipalAssignmentListResult.
func (c ClusterPrincipalAssignmentListResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "value", c.Value)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type ClusterProperties.
func (c ClusterProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "acceptedAudiences", c.AcceptedAudiences)
	populate(objectMap, "allowedFqdnList", c.AllowedFqdnList)
	populate(objectMap, "allowedIpRangeList", c.AllowedIPRangeList)
	populate(objectMap, "dataIngestionUri", c.DataIngestionURI)
	populate(objectMap, "enableAutoStop", c.EnableAutoStop)
	populate(objectMap, "enableDiskEncryption", c.EnableDiskEncryption)
	populate(objectMap, "enableDoubleEncryption", c.EnableDoubleEncryption)
	populate(objectMap, "enablePurge", c.EnablePurge)
	populate(objectMap, "enableStreamingIngest", c.EnableStreamingIngest)
	populate(objectMap, "engineType", c.EngineType)
	populate(objectMap, "keyVaultProperties", c.KeyVaultProperties)
	populate(objectMap, "languageExtensions", c.LanguageExtensions)
	populate(objectMap, "optimizedAutoscale", c.OptimizedAutoscale)
	populate(objectMap, "privateEndpointConnections", c.PrivateEndpointConnections)
	populate(objectMap, "provisioningState", c.ProvisioningState)
	populate(objectMap, "publicIPType", c.PublicIPType)
	populate(objectMap, "publicNetworkAccess", c.PublicNetworkAccess)
	populate(objectMap, "restrictOutboundNetworkAccess", c.RestrictOutboundNetworkAccess)
	populate(objectMap, "state", c.State)
	populate(objectMap, "stateReason", c.StateReason)
	populate(objectMap, "trustedExternalTenants", c.TrustedExternalTenants)
	populate(objectMap, "uri", c.URI)
	populate(objectMap, "virtualClusterGraduationProperties", c.VirtualClusterGraduationProperties)
	populate(objectMap, "virtualNetworkConfiguration", c.VirtualNetworkConfiguration)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type ClusterUpdate.
func (c ClusterUpdate) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "id", c.ID)
	populate(objectMap, "identity", c.Identity)
	populate(objectMap, "location", c.Location)
	populate(objectMap, "name", c.Name)
	populate(objectMap, "properties", c.Properties)
	populate(objectMap, "sku", c.SKU)
	populate(objectMap, "tags", c.Tags)
	populate(objectMap, "type", c.Type)
	return json.Marshal(objectMap)
}

// GetDataConnection implements the DataConnectionClassification interface for type DataConnection.
func (d *DataConnection) GetDataConnection() *DataConnection { return d }

// MarshalJSON implements the json.Marshaller interface for type DataConnection.
func (d DataConnection) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "id", d.ID)
	objectMap["kind"] = d.Kind
	populate(objectMap, "location", d.Location)
	populate(objectMap, "name", d.Name)
	populate(objectMap, "type", d.Type)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type DataConnectionListResult.
func (d DataConnectionListResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "value", d.Value)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type DataConnectionListResult.
func (d *DataConnectionListResult) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "value":
			d.Value, err = unmarshalDataConnectionClassificationArray(val)
			delete(rawMsg, key)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type DataConnectionValidation.
func (d DataConnectionValidation) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "dataConnectionName", d.DataConnectionName)
	populate(objectMap, "properties", d.Properties)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type DataConnectionValidation.
func (d *DataConnectionValidation) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "dataConnectionName":
			err = unpopulate(val, &d.DataConnectionName)
			delete(rawMsg, key)
		case "properties":
			d.Properties, err = unmarshalDataConnectionClassification(val)
			delete(rawMsg, key)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type DataConnectionValidationListResult.
func (d DataConnectionValidationListResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "value", d.Value)
	return json.Marshal(objectMap)
}

// GetDatabase implements the DatabaseClassification interface for type Database.
func (d *Database) GetDatabase() *Database { return d }

// MarshalJSON implements the json.Marshaller interface for type Database.
func (d Database) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "id", d.ID)
	objectMap["kind"] = d.Kind
	populate(objectMap, "location", d.Location)
	populate(objectMap, "name", d.Name)
	populate(objectMap, "type", d.Type)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type DatabaseListResult.
func (d DatabaseListResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "value", d.Value)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type DatabaseListResult.
func (d *DatabaseListResult) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "value":
			d.Value, err = unmarshalDatabaseClassificationArray(val)
			delete(rawMsg, key)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type DatabasePrincipalAssignmentListResult.
func (d DatabasePrincipalAssignmentListResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "value", d.Value)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type DatabasePrincipalListRequest.
func (d DatabasePrincipalListRequest) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "value", d.Value)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type DatabasePrincipalListResult.
func (d DatabasePrincipalListResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "value", d.Value)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type DiagnoseVirtualNetworkResult.
func (d DiagnoseVirtualNetworkResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "findings", d.Findings)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type EndpointDependency.
func (e EndpointDependency) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "domainName", e.DomainName)
	populate(objectMap, "endpointDetails", e.EndpointDetails)
	return json.Marshal(objectMap)
}

// GetDataConnection implements the DataConnectionClassification interface for type EventGridDataConnection.
func (e *EventGridDataConnection) GetDataConnection() *DataConnection {
	return &DataConnection{
		Location: e.Location,
		Kind:     e.Kind,
		ID:       e.ID,
		Name:     e.Name,
		Type:     e.Type,
	}
}

// MarshalJSON implements the json.Marshaller interface for type EventGridDataConnection.
func (e EventGridDataConnection) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "id", e.ID)
	objectMap["kind"] = DataConnectionKindEventGrid
	populate(objectMap, "location", e.Location)
	populate(objectMap, "name", e.Name)
	populate(objectMap, "properties", e.Properties)
	populate(objectMap, "type", e.Type)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type EventGridDataConnection.
func (e *EventGridDataConnection) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "id":
			err = unpopulate(val, &e.ID)
			delete(rawMsg, key)
		case "kind":
			err = unpopulate(val, &e.Kind)
			delete(rawMsg, key)
		case "location":
			err = unpopulate(val, &e.Location)
			delete(rawMsg, key)
		case "name":
			err = unpopulate(val, &e.Name)
			delete(rawMsg, key)
		case "properties":
			err = unpopulate(val, &e.Properties)
			delete(rawMsg, key)
		case "type":
			err = unpopulate(val, &e.Type)
			delete(rawMsg, key)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type EventHubConnectionProperties.
func (e EventHubConnectionProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "compression", e.Compression)
	populate(objectMap, "consumerGroup", e.ConsumerGroup)
	populate(objectMap, "dataFormat", e.DataFormat)
	populate(objectMap, "databaseRouting", e.DatabaseRouting)
	populate(objectMap, "eventHubResourceId", e.EventHubResourceID)
	populate(objectMap, "eventSystemProperties", e.EventSystemProperties)
	populate(objectMap, "managedIdentityObjectId", e.ManagedIdentityObjectID)
	populate(objectMap, "managedIdentityResourceId", e.ManagedIdentityResourceID)
	populate(objectMap, "mappingRuleName", e.MappingRuleName)
	populate(objectMap, "provisioningState", e.ProvisioningState)
	populate(objectMap, "tableName", e.TableName)
	return json.Marshal(objectMap)
}

// GetDataConnection implements the DataConnectionClassification interface for type EventHubDataConnection.
func (e *EventHubDataConnection) GetDataConnection() *DataConnection {
	return &DataConnection{
		Location: e.Location,
		Kind:     e.Kind,
		ID:       e.ID,
		Name:     e.Name,
		Type:     e.Type,
	}
}

// MarshalJSON implements the json.Marshaller interface for type EventHubDataConnection.
func (e EventHubDataConnection) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "id", e.ID)
	objectMap["kind"] = DataConnectionKindEventHub
	populate(objectMap, "location", e.Location)
	populate(objectMap, "name", e.Name)
	populate(objectMap, "properties", e.Properties)
	populate(objectMap, "type", e.Type)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type EventHubDataConnection.
func (e *EventHubDataConnection) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "id":
			err = unpopulate(val, &e.ID)
			delete(rawMsg, key)
		case "kind":
			err = unpopulate(val, &e.Kind)
			delete(rawMsg, key)
		case "location":
			err = unpopulate(val, &e.Location)
			delete(rawMsg, key)
		case "name":
			err = unpopulate(val, &e.Name)
			delete(rawMsg, key)
		case "properties":
			err = unpopulate(val, &e.Properties)
			delete(rawMsg, key)
		case "type":
			err = unpopulate(val, &e.Type)
			delete(rawMsg, key)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type FollowerDatabaseListResult.
func (f FollowerDatabaseListResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "value", f.Value)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type Identity.
func (i Identity) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "principalId", i.PrincipalID)
	populate(objectMap, "tenantId", i.TenantID)
	populate(objectMap, "type", i.Type)
	populate(objectMap, "userAssignedIdentities", i.UserAssignedIdentities)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type IotHubConnectionProperties.
func (i IotHubConnectionProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "consumerGroup", i.ConsumerGroup)
	populate(objectMap, "dataFormat", i.DataFormat)
	populate(objectMap, "databaseRouting", i.DatabaseRouting)
	populate(objectMap, "eventSystemProperties", i.EventSystemProperties)
	populate(objectMap, "iotHubResourceId", i.IotHubResourceID)
	populate(objectMap, "mappingRuleName", i.MappingRuleName)
	populate(objectMap, "provisioningState", i.ProvisioningState)
	populate(objectMap, "sharedAccessPolicyName", i.SharedAccessPolicyName)
	populate(objectMap, "tableName", i.TableName)
	return json.Marshal(objectMap)
}

// GetDataConnection implements the DataConnectionClassification interface for type IotHubDataConnection.
func (i *IotHubDataConnection) GetDataConnection() *DataConnection {
	return &DataConnection{
		Location: i.Location,
		Kind:     i.Kind,
		ID:       i.ID,
		Name:     i.Name,
		Type:     i.Type,
	}
}

// MarshalJSON implements the json.Marshaller interface for type IotHubDataConnection.
func (i IotHubDataConnection) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "id", i.ID)
	objectMap["kind"] = DataConnectionKindIotHub
	populate(objectMap, "location", i.Location)
	populate(objectMap, "name", i.Name)
	populate(objectMap, "properties", i.Properties)
	populate(objectMap, "type", i.Type)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type IotHubDataConnection.
func (i *IotHubDataConnection) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "id":
			err = unpopulate(val, &i.ID)
			delete(rawMsg, key)
		case "kind":
			err = unpopulate(val, &i.Kind)
			delete(rawMsg, key)
		case "location":
			err = unpopulate(val, &i.Location)
			delete(rawMsg, key)
		case "name":
			err = unpopulate(val, &i.Name)
			delete(rawMsg, key)
		case "properties":
			err = unpopulate(val, &i.Properties)
			delete(rawMsg, key)
		case "type":
			err = unpopulate(val, &i.Type)
			delete(rawMsg, key)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type LanguageExtensionsList.
func (l LanguageExtensionsList) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "value", l.Value)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type ListResourceSKUsResult.
func (l ListResourceSKUsResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "value", l.Value)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type ManagedPrivateEndpoint.
func (m ManagedPrivateEndpoint) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "id", m.ID)
	populate(objectMap, "name", m.Name)
	populate(objectMap, "properties", m.Properties)
	populate(objectMap, "systemData", m.SystemData)
	populate(objectMap, "type", m.Type)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type ManagedPrivateEndpointListResult.
func (m ManagedPrivateEndpointListResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "value", m.Value)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type OperationListResult.
func (o OperationListResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "nextLink", o.NextLink)
	populate(objectMap, "value", o.Value)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type OperationResult.
func (o OperationResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populateTimeRFC3339(objectMap, "endTime", o.EndTime)
	populate(objectMap, "error", o.Error)
	populate(objectMap, "id", o.ID)
	populate(objectMap, "name", o.Name)
	populate(objectMap, "percentComplete", o.PercentComplete)
	populate(objectMap, "properties", o.Properties)
	populateTimeRFC3339(objectMap, "startTime", o.StartTime)
	populate(objectMap, "status", o.Status)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type OperationResult.
func (o *OperationResult) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "endTime":
			err = unpopulateTimeRFC3339(val, &o.EndTime)
			delete(rawMsg, key)
		case "error":
			err = unpopulate(val, &o.Error)
			delete(rawMsg, key)
		case "id":
			err = unpopulate(val, &o.ID)
			delete(rawMsg, key)
		case "name":
			err = unpopulate(val, &o.Name)
			delete(rawMsg, key)
		case "percentComplete":
			err = unpopulate(val, &o.PercentComplete)
			delete(rawMsg, key)
		case "properties":
			err = unpopulate(val, &o.Properties)
			delete(rawMsg, key)
		case "startTime":
			err = unpopulateTimeRFC3339(val, &o.StartTime)
			delete(rawMsg, key)
		case "status":
			err = unpopulate(val, &o.Status)
			delete(rawMsg, key)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type OutboundNetworkDependenciesEndpointListResult.
func (o OutboundNetworkDependenciesEndpointListResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "nextLink", o.NextLink)
	populate(objectMap, "value", o.Value)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type OutboundNetworkDependenciesEndpointProperties.
func (o OutboundNetworkDependenciesEndpointProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "category", o.Category)
	populate(objectMap, "endpoints", o.Endpoints)
	populate(objectMap, "provisioningState", o.ProvisioningState)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type PrivateEndpointConnectionListResult.
func (p PrivateEndpointConnectionListResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "value", p.Value)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type PrivateLinkResourceListResult.
func (p PrivateLinkResourceListResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
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

// GetDatabase implements the DatabaseClassification interface for type ReadOnlyFollowingDatabase.
func (r *ReadOnlyFollowingDatabase) GetDatabase() *Database {
	return &Database{
		Location: r.Location,
		Kind:     r.Kind,
		ID:       r.ID,
		Name:     r.Name,
		Type:     r.Type,
	}
}

// MarshalJSON implements the json.Marshaller interface for type ReadOnlyFollowingDatabase.
func (r ReadOnlyFollowingDatabase) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "id", r.ID)
	objectMap["kind"] = KindReadOnlyFollowing
	populate(objectMap, "location", r.Location)
	populate(objectMap, "name", r.Name)
	populate(objectMap, "properties", r.Properties)
	populate(objectMap, "type", r.Type)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ReadOnlyFollowingDatabase.
func (r *ReadOnlyFollowingDatabase) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "id":
			err = unpopulate(val, &r.ID)
			delete(rawMsg, key)
		case "kind":
			err = unpopulate(val, &r.Kind)
			delete(rawMsg, key)
		case "location":
			err = unpopulate(val, &r.Location)
			delete(rawMsg, key)
		case "name":
			err = unpopulate(val, &r.Name)
			delete(rawMsg, key)
		case "properties":
			err = unpopulate(val, &r.Properties)
			delete(rawMsg, key)
		case "type":
			err = unpopulate(val, &r.Type)
			delete(rawMsg, key)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

// GetDatabase implements the DatabaseClassification interface for type ReadWriteDatabase.
func (r *ReadWriteDatabase) GetDatabase() *Database {
	return &Database{
		Location: r.Location,
		Kind:     r.Kind,
		ID:       r.ID,
		Name:     r.Name,
		Type:     r.Type,
	}
}

// MarshalJSON implements the json.Marshaller interface for type ReadWriteDatabase.
func (r ReadWriteDatabase) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "id", r.ID)
	objectMap["kind"] = KindReadWrite
	populate(objectMap, "location", r.Location)
	populate(objectMap, "name", r.Name)
	populate(objectMap, "properties", r.Properties)
	populate(objectMap, "type", r.Type)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ReadWriteDatabase.
func (r *ReadWriteDatabase) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "id":
			err = unpopulate(val, &r.ID)
			delete(rawMsg, key)
		case "kind":
			err = unpopulate(val, &r.Kind)
			delete(rawMsg, key)
		case "location":
			err = unpopulate(val, &r.Location)
			delete(rawMsg, key)
		case "name":
			err = unpopulate(val, &r.Name)
			delete(rawMsg, key)
		case "properties":
			err = unpopulate(val, &r.Properties)
			delete(rawMsg, key)
		case "type":
			err = unpopulate(val, &r.Type)
			delete(rawMsg, key)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type SKUDescription.
func (s SKUDescription) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "locationInfo", s.LocationInfo)
	populate(objectMap, "locations", s.Locations)
	populate(objectMap, "name", s.Name)
	populate(objectMap, "resourceType", s.ResourceType)
	populate(objectMap, "restrictions", s.Restrictions)
	populate(objectMap, "tier", s.Tier)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type SKUDescriptionList.
func (s SKUDescriptionList) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "value", s.Value)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type SKULocationInfoItem.
func (s SKULocationInfoItem) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "location", s.Location)
	populate(objectMap, "zones", s.Zones)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type Script.
func (s Script) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "id", s.ID)
	populate(objectMap, "name", s.Name)
	populate(objectMap, "properties", s.Properties)
	populate(objectMap, "systemData", s.SystemData)
	populate(objectMap, "type", s.Type)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type ScriptListResult.
func (s ScriptListResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
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

// MarshalJSON implements the json.Marshaller interface for type TableLevelSharingProperties.
func (t TableLevelSharingProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "externalTablesToExclude", t.ExternalTablesToExclude)
	populate(objectMap, "externalTablesToInclude", t.ExternalTablesToInclude)
	populate(objectMap, "materializedViewsToExclude", t.MaterializedViewsToExclude)
	populate(objectMap, "materializedViewsToInclude", t.MaterializedViewsToInclude)
	populate(objectMap, "tablesToExclude", t.TablesToExclude)
	populate(objectMap, "tablesToInclude", t.TablesToInclude)
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
