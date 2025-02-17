//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armazurestackhci

import (
	"encoding/json"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"reflect"
)

// MarshalJSON implements the json.Marshaller interface for type ArcSettingList.
func (a ArcSettingList) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "nextLink", a.NextLink)
	populate(objectMap, "value", a.Value)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type ArcSettingProperties.
func (a ArcSettingProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "aggregateState", a.AggregateState)
	populate(objectMap, "arcApplicationClientId", a.ArcApplicationClientID)
	populate(objectMap, "arcApplicationObjectId", a.ArcApplicationObjectID)
	populate(objectMap, "arcApplicationTenantId", a.ArcApplicationTenantID)
	populate(objectMap, "arcInstanceResourceGroup", a.ArcInstanceResourceGroup)
	populate(objectMap, "arcServicePrincipalObjectId", a.ArcServicePrincipalObjectID)
	populate(objectMap, "perNodeDetails", a.PerNodeDetails)
	populate(objectMap, "provisioningState", a.ProvisioningState)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type Cluster.
func (c Cluster) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "id", c.ID)
	populate(objectMap, "location", c.Location)
	populate(objectMap, "name", c.Name)
	populate(objectMap, "properties", c.Properties)
	populate(objectMap, "systemData", c.SystemData)
	populate(objectMap, "tags", c.Tags)
	populate(objectMap, "type", c.Type)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type ClusterList.
func (c ClusterList) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "nextLink", c.NextLink)
	populate(objectMap, "value", c.Value)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type ClusterPatch.
func (c ClusterPatch) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "properties", c.Properties)
	populate(objectMap, "tags", c.Tags)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type ClusterProperties.
func (c ClusterProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "aadApplicationObjectId", c.AADApplicationObjectID)
	populate(objectMap, "aadClientId", c.AADClientID)
	populate(objectMap, "aadServicePrincipalObjectId", c.AADServicePrincipalObjectID)
	populate(objectMap, "aadTenantId", c.AADTenantID)
	populate(objectMap, "billingModel", c.BillingModel)
	populate(objectMap, "cloudId", c.CloudID)
	populate(objectMap, "cloudManagementEndpoint", c.CloudManagementEndpoint)
	populate(objectMap, "desiredProperties", c.DesiredProperties)
	populateTimeRFC3339(objectMap, "lastBillingTimestamp", c.LastBillingTimestamp)
	populateTimeRFC3339(objectMap, "lastSyncTimestamp", c.LastSyncTimestamp)
	populate(objectMap, "provisioningState", c.ProvisioningState)
	populateTimeRFC3339(objectMap, "registrationTimestamp", c.RegistrationTimestamp)
	populate(objectMap, "reportedProperties", c.ReportedProperties)
	populate(objectMap, "serviceEndpoint", c.ServiceEndpoint)
	populate(objectMap, "status", c.Status)
	populate(objectMap, "trialDaysRemaining", c.TrialDaysRemaining)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ClusterProperties.
func (c *ClusterProperties) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "aadApplicationObjectId":
			err = unpopulate(val, &c.AADApplicationObjectID)
			delete(rawMsg, key)
		case "aadClientId":
			err = unpopulate(val, &c.AADClientID)
			delete(rawMsg, key)
		case "aadServicePrincipalObjectId":
			err = unpopulate(val, &c.AADServicePrincipalObjectID)
			delete(rawMsg, key)
		case "aadTenantId":
			err = unpopulate(val, &c.AADTenantID)
			delete(rawMsg, key)
		case "billingModel":
			err = unpopulate(val, &c.BillingModel)
			delete(rawMsg, key)
		case "cloudId":
			err = unpopulate(val, &c.CloudID)
			delete(rawMsg, key)
		case "cloudManagementEndpoint":
			err = unpopulate(val, &c.CloudManagementEndpoint)
			delete(rawMsg, key)
		case "desiredProperties":
			err = unpopulate(val, &c.DesiredProperties)
			delete(rawMsg, key)
		case "lastBillingTimestamp":
			err = unpopulateTimeRFC3339(val, &c.LastBillingTimestamp)
			delete(rawMsg, key)
		case "lastSyncTimestamp":
			err = unpopulateTimeRFC3339(val, &c.LastSyncTimestamp)
			delete(rawMsg, key)
		case "provisioningState":
			err = unpopulate(val, &c.ProvisioningState)
			delete(rawMsg, key)
		case "registrationTimestamp":
			err = unpopulateTimeRFC3339(val, &c.RegistrationTimestamp)
			delete(rawMsg, key)
		case "reportedProperties":
			err = unpopulate(val, &c.ReportedProperties)
			delete(rawMsg, key)
		case "serviceEndpoint":
			err = unpopulate(val, &c.ServiceEndpoint)
			delete(rawMsg, key)
		case "status":
			err = unpopulate(val, &c.Status)
			delete(rawMsg, key)
		case "trialDaysRemaining":
			err = unpopulate(val, &c.TrialDaysRemaining)
			delete(rawMsg, key)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type ClusterReportedProperties.
func (c ClusterReportedProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "clusterId", c.ClusterID)
	populate(objectMap, "clusterName", c.ClusterName)
	populate(objectMap, "clusterVersion", c.ClusterVersion)
	populate(objectMap, "diagnosticLevel", c.DiagnosticLevel)
	populate(objectMap, "imdsAttestation", c.ImdsAttestation)
	populateTimeRFC3339(objectMap, "lastUpdated", c.LastUpdated)
	populate(objectMap, "nodes", c.Nodes)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ClusterReportedProperties.
func (c *ClusterReportedProperties) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "clusterId":
			err = unpopulate(val, &c.ClusterID)
			delete(rawMsg, key)
		case "clusterName":
			err = unpopulate(val, &c.ClusterName)
			delete(rawMsg, key)
		case "clusterVersion":
			err = unpopulate(val, &c.ClusterVersion)
			delete(rawMsg, key)
		case "diagnosticLevel":
			err = unpopulate(val, &c.DiagnosticLevel)
			delete(rawMsg, key)
		case "imdsAttestation":
			err = unpopulate(val, &c.ImdsAttestation)
			delete(rawMsg, key)
		case "lastUpdated":
			err = unpopulateTimeRFC3339(val, &c.LastUpdated)
			delete(rawMsg, key)
		case "nodes":
			err = unpopulate(val, &c.Nodes)
			delete(rawMsg, key)
		}
		if err != nil {
			return err
		}
	}
	return nil
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

// MarshalJSON implements the json.Marshaller interface for type Extension.
func (e Extension) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "id", e.ID)
	populate(objectMap, "name", e.Name)
	populate(objectMap, "properties", e.Properties)
	populate(objectMap, "systemData", e.SystemData)
	populate(objectMap, "type", e.Type)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type ExtensionList.
func (e ExtensionList) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "nextLink", e.NextLink)
	populate(objectMap, "value", e.Value)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type ExtensionProperties.
func (e ExtensionProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "aggregateState", e.AggregateState)
	populate(objectMap, "extensionParameters", e.ExtensionParameters)
	populate(objectMap, "perNodeExtensionDetails", e.PerNodeExtensionDetails)
	populate(objectMap, "provisioningState", e.ProvisioningState)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type OperationListResult.
func (o OperationListResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "nextLink", o.NextLink)
	populate(objectMap, "value", o.Value)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type PasswordCredential.
func (p PasswordCredential) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populateTimeRFC3339(objectMap, "endDateTime", p.EndDateTime)
	populate(objectMap, "keyId", p.KeyID)
	populate(objectMap, "secretText", p.SecretText)
	populateTimeRFC3339(objectMap, "startDateTime", p.StartDateTime)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type PasswordCredential.
func (p *PasswordCredential) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "endDateTime":
			err = unpopulateTimeRFC3339(val, &p.EndDateTime)
			delete(rawMsg, key)
		case "keyId":
			err = unpopulate(val, &p.KeyID)
			delete(rawMsg, key)
		case "secretText":
			err = unpopulate(val, &p.SecretText)
			delete(rawMsg, key)
		case "startDateTime":
			err = unpopulateTimeRFC3339(val, &p.StartDateTime)
			delete(rawMsg, key)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type RawCertificateData.
func (r RawCertificateData) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "certificates", r.Certificates)
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
