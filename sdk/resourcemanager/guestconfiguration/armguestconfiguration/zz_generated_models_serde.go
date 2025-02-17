//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armguestconfiguration

import (
	"encoding/json"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"reflect"
)

// MarshalJSON implements the json.Marshaller interface for type AssignmentList.
func (a AssignmentList) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "value", a.Value)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type AssignmentProperties.
func (a AssignmentProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "assignmentHash", a.AssignmentHash)
	populate(objectMap, "complianceStatus", a.ComplianceStatus)
	populate(objectMap, "context", a.Context)
	populate(objectMap, "guestConfiguration", a.GuestConfiguration)
	populateTimeRFC3339(objectMap, "lastComplianceStatusChecked", a.LastComplianceStatusChecked)
	populate(objectMap, "latestAssignmentReport", a.LatestAssignmentReport)
	populate(objectMap, "latestReportId", a.LatestReportID)
	populate(objectMap, "parameterHash", a.ParameterHash)
	populate(objectMap, "provisioningState", a.ProvisioningState)
	populate(objectMap, "resourceType", a.ResourceType)
	populate(objectMap, "targetResourceId", a.TargetResourceID)
	populate(objectMap, "vmssVMList", a.VmssVMList)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type AssignmentProperties.
func (a *AssignmentProperties) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "assignmentHash":
			err = unpopulate(val, &a.AssignmentHash)
			delete(rawMsg, key)
		case "complianceStatus":
			err = unpopulate(val, &a.ComplianceStatus)
			delete(rawMsg, key)
		case "context":
			err = unpopulate(val, &a.Context)
			delete(rawMsg, key)
		case "guestConfiguration":
			err = unpopulate(val, &a.GuestConfiguration)
			delete(rawMsg, key)
		case "lastComplianceStatusChecked":
			err = unpopulateTimeRFC3339(val, &a.LastComplianceStatusChecked)
			delete(rawMsg, key)
		case "latestAssignmentReport":
			err = unpopulate(val, &a.LatestAssignmentReport)
			delete(rawMsg, key)
		case "latestReportId":
			err = unpopulate(val, &a.LatestReportID)
			delete(rawMsg, key)
		case "parameterHash":
			err = unpopulate(val, &a.ParameterHash)
			delete(rawMsg, key)
		case "provisioningState":
			err = unpopulate(val, &a.ProvisioningState)
			delete(rawMsg, key)
		case "resourceType":
			err = unpopulate(val, &a.ResourceType)
			delete(rawMsg, key)
		case "targetResourceId":
			err = unpopulate(val, &a.TargetResourceID)
			delete(rawMsg, key)
		case "vmssVMList":
			err = unpopulate(val, &a.VmssVMList)
			delete(rawMsg, key)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type AssignmentReportDetails.
func (a AssignmentReportDetails) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "complianceStatus", a.ComplianceStatus)
	populateTimeRFC3339(objectMap, "endTime", a.EndTime)
	populate(objectMap, "jobId", a.JobID)
	populate(objectMap, "operationType", a.OperationType)
	populate(objectMap, "resources", a.Resources)
	populateTimeRFC3339(objectMap, "startTime", a.StartTime)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type AssignmentReportDetails.
func (a *AssignmentReportDetails) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "complianceStatus":
			err = unpopulate(val, &a.ComplianceStatus)
			delete(rawMsg, key)
		case "endTime":
			err = unpopulateTimeRFC3339(val, &a.EndTime)
			delete(rawMsg, key)
		case "jobId":
			err = unpopulate(val, &a.JobID)
			delete(rawMsg, key)
		case "operationType":
			err = unpopulate(val, &a.OperationType)
			delete(rawMsg, key)
		case "resources":
			err = unpopulate(val, &a.Resources)
			delete(rawMsg, key)
		case "startTime":
			err = unpopulateTimeRFC3339(val, &a.StartTime)
			delete(rawMsg, key)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type AssignmentReportList.
func (a AssignmentReportList) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "value", a.Value)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type AssignmentReportProperties.
func (a AssignmentReportProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "assignment", a.Assignment)
	populate(objectMap, "complianceStatus", a.ComplianceStatus)
	populate(objectMap, "details", a.Details)
	populateTimeRFC3339(objectMap, "endTime", a.EndTime)
	populate(objectMap, "reportId", a.ReportID)
	populateTimeRFC3339(objectMap, "startTime", a.StartTime)
	populate(objectMap, "vm", a.VM)
	populate(objectMap, "vmssResourceId", a.VmssResourceID)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type AssignmentReportProperties.
func (a *AssignmentReportProperties) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "assignment":
			err = unpopulate(val, &a.Assignment)
			delete(rawMsg, key)
		case "complianceStatus":
			err = unpopulate(val, &a.ComplianceStatus)
			delete(rawMsg, key)
		case "details":
			err = unpopulate(val, &a.Details)
			delete(rawMsg, key)
		case "endTime":
			err = unpopulateTimeRFC3339(val, &a.EndTime)
			delete(rawMsg, key)
		case "reportId":
			err = unpopulate(val, &a.ReportID)
			delete(rawMsg, key)
		case "startTime":
			err = unpopulateTimeRFC3339(val, &a.StartTime)
			delete(rawMsg, key)
		case "vm":
			err = unpopulate(val, &a.VM)
			delete(rawMsg, key)
		case "vmssResourceId":
			err = unpopulate(val, &a.VmssResourceID)
			delete(rawMsg, key)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type AssignmentReportResource.
func (a AssignmentReportResource) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "complianceStatus", a.ComplianceStatus)
	populate(objectMap, "properties", &a.Properties)
	populate(objectMap, "reasons", a.Reasons)
	populate(objectMap, "resourceId", a.ResourceID)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type CommonAssignmentReport.
func (c CommonAssignmentReport) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "assignment", c.Assignment)
	populate(objectMap, "complianceStatus", c.ComplianceStatus)
	populateTimeRFC3339(objectMap, "endTime", c.EndTime)
	populate(objectMap, "id", c.ID)
	populate(objectMap, "operationType", c.OperationType)
	populate(objectMap, "reportId", c.ReportID)
	populate(objectMap, "resources", c.Resources)
	populateTimeRFC3339(objectMap, "startTime", c.StartTime)
	populate(objectMap, "vm", c.VM)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type CommonAssignmentReport.
func (c *CommonAssignmentReport) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "assignment":
			err = unpopulate(val, &c.Assignment)
			delete(rawMsg, key)
		case "complianceStatus":
			err = unpopulate(val, &c.ComplianceStatus)
			delete(rawMsg, key)
		case "endTime":
			err = unpopulateTimeRFC3339(val, &c.EndTime)
			delete(rawMsg, key)
		case "id":
			err = unpopulate(val, &c.ID)
			delete(rawMsg, key)
		case "operationType":
			err = unpopulate(val, &c.OperationType)
			delete(rawMsg, key)
		case "reportId":
			err = unpopulate(val, &c.ReportID)
			delete(rawMsg, key)
		case "resources":
			err = unpopulate(val, &c.Resources)
			delete(rawMsg, key)
		case "startTime":
			err = unpopulateTimeRFC3339(val, &c.StartTime)
			delete(rawMsg, key)
		case "vm":
			err = unpopulate(val, &c.VM)
			delete(rawMsg, key)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type Navigation.
func (n Navigation) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "assignmentType", n.AssignmentType)
	populate(objectMap, "configurationParameter", n.ConfigurationParameter)
	populate(objectMap, "configurationProtectedParameter", n.ConfigurationProtectedParameter)
	populate(objectMap, "configurationSetting", n.ConfigurationSetting)
	populate(objectMap, "contentHash", n.ContentHash)
	populate(objectMap, "contentType", n.ContentType)
	populate(objectMap, "contentUri", n.ContentURI)
	populate(objectMap, "kind", n.Kind)
	populate(objectMap, "name", n.Name)
	populate(objectMap, "version", n.Version)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type OperationList.
func (o OperationList) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "value", o.Value)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type VMSSVMInfo.
func (v VMSSVMInfo) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "complianceStatus", v.ComplianceStatus)
	populateTimeRFC3339(objectMap, "lastComplianceChecked", v.LastComplianceChecked)
	populate(objectMap, "latestReportId", v.LatestReportID)
	populate(objectMap, "vmId", v.VMID)
	populate(objectMap, "vmResourceId", v.VMResourceID)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type VMSSVMInfo.
func (v *VMSSVMInfo) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "complianceStatus":
			err = unpopulate(val, &v.ComplianceStatus)
			delete(rawMsg, key)
		case "lastComplianceChecked":
			err = unpopulateTimeRFC3339(val, &v.LastComplianceChecked)
			delete(rawMsg, key)
		case "latestReportId":
			err = unpopulate(val, &v.LatestReportID)
			delete(rawMsg, key)
		case "vmId":
			err = unpopulate(val, &v.VMID)
			delete(rawMsg, key)
		case "vmResourceId":
			err = unpopulate(val, &v.VMResourceID)
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
