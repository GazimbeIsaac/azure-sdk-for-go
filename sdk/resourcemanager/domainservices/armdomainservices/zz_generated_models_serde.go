//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdomainservices

import (
	"encoding/json"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"reflect"
)

// MarshalJSON implements the json.Marshaller interface for type CloudErrorBody.
func (c CloudErrorBody) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "code", c.Code)
	populate(objectMap, "details", c.Details)
	populate(objectMap, "message", c.Message)
	populate(objectMap, "target", c.Target)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type ConfigDiagnostics.
func (c ConfigDiagnostics) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populateTimeRFC1123(objectMap, "lastExecuted", c.LastExecuted)
	populate(objectMap, "validatorResults", c.ValidatorResults)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ConfigDiagnostics.
func (c *ConfigDiagnostics) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "lastExecuted":
			err = unpopulateTimeRFC1123(val, &c.LastExecuted)
			delete(rawMsg, key)
		case "validatorResults":
			err = unpopulate(val, &c.ValidatorResults)
			delete(rawMsg, key)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type ConfigDiagnosticsValidatorResult.
func (c ConfigDiagnosticsValidatorResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "issues", c.Issues)
	populate(objectMap, "replicaSetSubnetDisplayName", c.ReplicaSetSubnetDisplayName)
	populate(objectMap, "status", c.Status)
	populate(objectMap, "validatorId", c.ValidatorID)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type ConfigDiagnosticsValidatorResultIssue.
func (c ConfigDiagnosticsValidatorResultIssue) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "descriptionParams", c.DescriptionParams)
	populate(objectMap, "id", c.ID)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type ContainerAccount.
func (c ContainerAccount) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "accountName", c.AccountName)
	populate(objectMap, "password", c.Password)
	populate(objectMap, "spn", c.Spn)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type DomainService.
func (d DomainService) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "etag", d.Etag)
	populate(objectMap, "id", d.ID)
	populate(objectMap, "location", d.Location)
	populate(objectMap, "name", d.Name)
	populate(objectMap, "properties", d.Properties)
	populate(objectMap, "systemData", d.SystemData)
	populate(objectMap, "tags", d.Tags)
	populate(objectMap, "type", d.Type)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type DomainServiceListResult.
func (d DomainServiceListResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "nextLink", d.NextLink)
	populate(objectMap, "value", d.Value)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type DomainServiceProperties.
func (d DomainServiceProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "configDiagnostics", d.ConfigDiagnostics)
	populate(objectMap, "deploymentId", d.DeploymentID)
	populate(objectMap, "domainConfigurationType", d.DomainConfigurationType)
	populate(objectMap, "domainName", d.DomainName)
	populate(objectMap, "domainSecuritySettings", d.DomainSecuritySettings)
	populate(objectMap, "filteredSync", d.FilteredSync)
	populate(objectMap, "ldapsSettings", d.LdapsSettings)
	populate(objectMap, "migrationProperties", d.MigrationProperties)
	populate(objectMap, "notificationSettings", d.NotificationSettings)
	populate(objectMap, "provisioningState", d.ProvisioningState)
	populate(objectMap, "replicaSets", d.ReplicaSets)
	populate(objectMap, "resourceForestSettings", d.ResourceForestSettings)
	populate(objectMap, "sku", d.SKU)
	populate(objectMap, "syncOwner", d.SyncOwner)
	populate(objectMap, "tenantId", d.TenantID)
	populate(objectMap, "version", d.Version)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type HealthAlert.
func (h HealthAlert) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "id", h.ID)
	populate(objectMap, "issue", h.Issue)
	populateTimeRFC3339(objectMap, "lastDetected", h.LastDetected)
	populate(objectMap, "name", h.Name)
	populateTimeRFC3339(objectMap, "raised", h.Raised)
	populate(objectMap, "resolutionUri", h.ResolutionURI)
	populate(objectMap, "severity", h.Severity)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type HealthAlert.
func (h *HealthAlert) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "id":
			err = unpopulate(val, &h.ID)
			delete(rawMsg, key)
		case "issue":
			err = unpopulate(val, &h.Issue)
			delete(rawMsg, key)
		case "lastDetected":
			err = unpopulateTimeRFC3339(val, &h.LastDetected)
			delete(rawMsg, key)
		case "name":
			err = unpopulate(val, &h.Name)
			delete(rawMsg, key)
		case "raised":
			err = unpopulateTimeRFC3339(val, &h.Raised)
			delete(rawMsg, key)
		case "resolutionUri":
			err = unpopulate(val, &h.ResolutionURI)
			delete(rawMsg, key)
		case "severity":
			err = unpopulate(val, &h.Severity)
			delete(rawMsg, key)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type LdapsSettings.
func (l LdapsSettings) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populateTimeRFC3339(objectMap, "certificateNotAfter", l.CertificateNotAfter)
	populate(objectMap, "certificateThumbprint", l.CertificateThumbprint)
	populate(objectMap, "externalAccess", l.ExternalAccess)
	populate(objectMap, "ldaps", l.Ldaps)
	populate(objectMap, "pfxCertificate", l.PfxCertificate)
	populate(objectMap, "pfxCertificatePassword", l.PfxCertificatePassword)
	populate(objectMap, "publicCertificate", l.PublicCertificate)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type LdapsSettings.
func (l *LdapsSettings) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "certificateNotAfter":
			err = unpopulateTimeRFC3339(val, &l.CertificateNotAfter)
			delete(rawMsg, key)
		case "certificateThumbprint":
			err = unpopulate(val, &l.CertificateThumbprint)
			delete(rawMsg, key)
		case "externalAccess":
			err = unpopulate(val, &l.ExternalAccess)
			delete(rawMsg, key)
		case "ldaps":
			err = unpopulate(val, &l.Ldaps)
			delete(rawMsg, key)
		case "pfxCertificate":
			err = unpopulate(val, &l.PfxCertificate)
			delete(rawMsg, key)
		case "pfxCertificatePassword":
			err = unpopulate(val, &l.PfxCertificatePassword)
			delete(rawMsg, key)
		case "publicCertificate":
			err = unpopulate(val, &l.PublicCertificate)
			delete(rawMsg, key)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type NotificationSettings.
func (n NotificationSettings) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "additionalRecipients", n.AdditionalRecipients)
	populate(objectMap, "notifyDcAdmins", n.NotifyDcAdmins)
	populate(objectMap, "notifyGlobalAdmins", n.NotifyGlobalAdmins)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type OperationEntityListResult.
func (o OperationEntityListResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "nextLink", o.NextLink)
	populate(objectMap, "value", o.Value)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type OuContainer.
func (o OuContainer) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "etag", o.Etag)
	populate(objectMap, "id", o.ID)
	populate(objectMap, "location", o.Location)
	populate(objectMap, "name", o.Name)
	populate(objectMap, "properties", o.Properties)
	populate(objectMap, "systemData", o.SystemData)
	populate(objectMap, "tags", o.Tags)
	populate(objectMap, "type", o.Type)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type OuContainerListResult.
func (o OuContainerListResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "nextLink", o.NextLink)
	populate(objectMap, "value", o.Value)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type OuContainerProperties.
func (o OuContainerProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "accounts", o.Accounts)
	populate(objectMap, "containerId", o.ContainerID)
	populate(objectMap, "deploymentId", o.DeploymentID)
	populate(objectMap, "distinguishedName", o.DistinguishedName)
	populate(objectMap, "domainName", o.DomainName)
	populate(objectMap, "provisioningState", o.ProvisioningState)
	populate(objectMap, "serviceStatus", o.ServiceStatus)
	populate(objectMap, "tenantId", o.TenantID)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type ReplicaSet.
func (r ReplicaSet) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "domainControllerIpAddress", r.DomainControllerIPAddress)
	populate(objectMap, "externalAccessIpAddress", r.ExternalAccessIPAddress)
	populate(objectMap, "healthAlerts", r.HealthAlerts)
	populateTimeRFC1123(objectMap, "healthLastEvaluated", r.HealthLastEvaluated)
	populate(objectMap, "healthMonitors", r.HealthMonitors)
	populate(objectMap, "location", r.Location)
	populate(objectMap, "replicaSetId", r.ReplicaSetID)
	populate(objectMap, "serviceStatus", r.ServiceStatus)
	populate(objectMap, "subnetId", r.SubnetID)
	populate(objectMap, "vnetSiteId", r.VnetSiteID)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ReplicaSet.
func (r *ReplicaSet) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "domainControllerIpAddress":
			err = unpopulate(val, &r.DomainControllerIPAddress)
			delete(rawMsg, key)
		case "externalAccessIpAddress":
			err = unpopulate(val, &r.ExternalAccessIPAddress)
			delete(rawMsg, key)
		case "healthAlerts":
			err = unpopulate(val, &r.HealthAlerts)
			delete(rawMsg, key)
		case "healthLastEvaluated":
			err = unpopulateTimeRFC1123(val, &r.HealthLastEvaluated)
			delete(rawMsg, key)
		case "healthMonitors":
			err = unpopulate(val, &r.HealthMonitors)
			delete(rawMsg, key)
		case "location":
			err = unpopulate(val, &r.Location)
			delete(rawMsg, key)
		case "replicaSetId":
			err = unpopulate(val, &r.ReplicaSetID)
			delete(rawMsg, key)
		case "serviceStatus":
			err = unpopulate(val, &r.ServiceStatus)
			delete(rawMsg, key)
		case "subnetId":
			err = unpopulate(val, &r.SubnetID)
			delete(rawMsg, key)
		case "vnetSiteId":
			err = unpopulate(val, &r.VnetSiteID)
			delete(rawMsg, key)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type Resource.
func (r Resource) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "etag", r.Etag)
	populate(objectMap, "id", r.ID)
	populate(objectMap, "location", r.Location)
	populate(objectMap, "name", r.Name)
	populate(objectMap, "systemData", r.SystemData)
	populate(objectMap, "tags", r.Tags)
	populate(objectMap, "type", r.Type)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type ResourceForestSettings.
func (r ResourceForestSettings) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "resourceForest", r.ResourceForest)
	populate(objectMap, "settings", r.Settings)
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
