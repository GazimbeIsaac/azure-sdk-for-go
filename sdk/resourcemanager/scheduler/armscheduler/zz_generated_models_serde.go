//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armscheduler

import (
	"encoding/json"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"reflect"
)

// GetHTTPAuthentication implements the HTTPAuthenticationClassification interface for type BasicAuthentication.
func (b *BasicAuthentication) GetHTTPAuthentication() *HTTPAuthentication {
	return &HTTPAuthentication{
		Type: b.Type,
	}
}

// MarshalJSON implements the json.Marshaller interface for type BasicAuthentication.
func (b BasicAuthentication) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "password", b.Password)
	objectMap["type"] = HTTPAuthenticationTypeBasic
	populate(objectMap, "username", b.Username)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type BasicAuthentication.
func (b *BasicAuthentication) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "password":
			err = unpopulate(val, &b.Password)
			delete(rawMsg, key)
		case "type":
			err = unpopulate(val, &b.Type)
			delete(rawMsg, key)
		case "username":
			err = unpopulate(val, &b.Username)
			delete(rawMsg, key)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

// GetHTTPAuthentication implements the HTTPAuthenticationClassification interface for type ClientCertAuthentication.
func (c *ClientCertAuthentication) GetHTTPAuthentication() *HTTPAuthentication {
	return &HTTPAuthentication{
		Type: c.Type,
	}
}

// MarshalJSON implements the json.Marshaller interface for type ClientCertAuthentication.
func (c ClientCertAuthentication) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populateTimeRFC3339(objectMap, "certificateExpirationDate", c.CertificateExpirationDate)
	populate(objectMap, "certificateSubjectName", c.CertificateSubjectName)
	populate(objectMap, "certificateThumbprint", c.CertificateThumbprint)
	populate(objectMap, "password", c.Password)
	populate(objectMap, "pfx", c.Pfx)
	objectMap["type"] = HTTPAuthenticationTypeClientCertificate
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ClientCertAuthentication.
func (c *ClientCertAuthentication) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "certificateExpirationDate":
			err = unpopulateTimeRFC3339(val, &c.CertificateExpirationDate)
			delete(rawMsg, key)
		case "certificateSubjectName":
			err = unpopulate(val, &c.CertificateSubjectName)
			delete(rawMsg, key)
		case "certificateThumbprint":
			err = unpopulate(val, &c.CertificateThumbprint)
			delete(rawMsg, key)
		case "password":
			err = unpopulate(val, &c.Password)
			delete(rawMsg, key)
		case "pfx":
			err = unpopulate(val, &c.Pfx)
			delete(rawMsg, key)
		case "type":
			err = unpopulate(val, &c.Type)
			delete(rawMsg, key)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

// GetHTTPAuthentication implements the HTTPAuthenticationClassification interface for type HTTPAuthentication.
func (h *HTTPAuthentication) GetHTTPAuthentication() *HTTPAuthentication { return h }

// MarshalJSON implements the json.Marshaller interface for type HTTPRequest.
func (h HTTPRequest) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "authentication", h.Authentication)
	populate(objectMap, "body", h.Body)
	populate(objectMap, "headers", h.Headers)
	populate(objectMap, "method", h.Method)
	populate(objectMap, "uri", h.URI)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type HTTPRequest.
func (h *HTTPRequest) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "authentication":
			h.Authentication, err = unmarshalHTTPAuthenticationClassification(val)
			delete(rawMsg, key)
		case "body":
			err = unpopulate(val, &h.Body)
			delete(rawMsg, key)
		case "headers":
			err = unpopulate(val, &h.Headers)
			delete(rawMsg, key)
		case "method":
			err = unpopulate(val, &h.Method)
			delete(rawMsg, key)
		case "uri":
			err = unpopulate(val, &h.URI)
			delete(rawMsg, key)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type JobCollectionDefinition.
func (j JobCollectionDefinition) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "id", j.ID)
	populate(objectMap, "location", j.Location)
	populate(objectMap, "name", j.Name)
	populate(objectMap, "properties", j.Properties)
	populate(objectMap, "tags", j.Tags)
	populate(objectMap, "type", j.Type)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type JobCollectionListResult.
func (j JobCollectionListResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "nextLink", j.NextLink)
	populate(objectMap, "value", j.Value)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type JobDefinition.
func (j JobDefinition) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "id", j.ID)
	populate(objectMap, "name", j.Name)
	populate(objectMap, "properties", j.Properties)
	populate(objectMap, "type", j.Type)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type JobHistoryDefinitionProperties.
func (j JobHistoryDefinitionProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "actionName", j.ActionName)
	populateTimeRFC3339(objectMap, "endTime", j.EndTime)
	populateTimeRFC3339(objectMap, "expectedExecutionTime", j.ExpectedExecutionTime)
	populate(objectMap, "message", j.Message)
	populate(objectMap, "repeatCount", j.RepeatCount)
	populate(objectMap, "retryCount", j.RetryCount)
	populateTimeRFC3339(objectMap, "startTime", j.StartTime)
	populate(objectMap, "status", j.Status)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type JobHistoryDefinitionProperties.
func (j *JobHistoryDefinitionProperties) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "actionName":
			err = unpopulate(val, &j.ActionName)
			delete(rawMsg, key)
		case "endTime":
			err = unpopulateTimeRFC3339(val, &j.EndTime)
			delete(rawMsg, key)
		case "expectedExecutionTime":
			err = unpopulateTimeRFC3339(val, &j.ExpectedExecutionTime)
			delete(rawMsg, key)
		case "message":
			err = unpopulate(val, &j.Message)
			delete(rawMsg, key)
		case "repeatCount":
			err = unpopulate(val, &j.RepeatCount)
			delete(rawMsg, key)
		case "retryCount":
			err = unpopulate(val, &j.RetryCount)
			delete(rawMsg, key)
		case "startTime":
			err = unpopulateTimeRFC3339(val, &j.StartTime)
			delete(rawMsg, key)
		case "status":
			err = unpopulate(val, &j.Status)
			delete(rawMsg, key)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type JobHistoryListResult.
func (j JobHistoryListResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "nextLink", j.NextLink)
	populate(objectMap, "value", j.Value)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type JobListResult.
func (j JobListResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "nextLink", j.NextLink)
	populate(objectMap, "value", j.Value)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type JobProperties.
func (j JobProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "action", j.Action)
	populate(objectMap, "recurrence", j.Recurrence)
	populateTimeRFC3339(objectMap, "startTime", j.StartTime)
	populate(objectMap, "state", j.State)
	populate(objectMap, "status", j.Status)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type JobProperties.
func (j *JobProperties) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "action":
			err = unpopulate(val, &j.Action)
			delete(rawMsg, key)
		case "recurrence":
			err = unpopulate(val, &j.Recurrence)
			delete(rawMsg, key)
		case "startTime":
			err = unpopulateTimeRFC3339(val, &j.StartTime)
			delete(rawMsg, key)
		case "state":
			err = unpopulate(val, &j.State)
			delete(rawMsg, key)
		case "status":
			err = unpopulate(val, &j.Status)
			delete(rawMsg, key)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type JobRecurrence.
func (j JobRecurrence) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "count", j.Count)
	populateTimeRFC3339(objectMap, "endTime", j.EndTime)
	populate(objectMap, "frequency", j.Frequency)
	populate(objectMap, "interval", j.Interval)
	populate(objectMap, "schedule", j.Schedule)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type JobRecurrence.
func (j *JobRecurrence) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "count":
			err = unpopulate(val, &j.Count)
			delete(rawMsg, key)
		case "endTime":
			err = unpopulateTimeRFC3339(val, &j.EndTime)
			delete(rawMsg, key)
		case "frequency":
			err = unpopulate(val, &j.Frequency)
			delete(rawMsg, key)
		case "interval":
			err = unpopulate(val, &j.Interval)
			delete(rawMsg, key)
		case "schedule":
			err = unpopulate(val, &j.Schedule)
			delete(rawMsg, key)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type JobRecurrenceSchedule.
func (j JobRecurrenceSchedule) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "hours", j.Hours)
	populate(objectMap, "minutes", j.Minutes)
	populate(objectMap, "monthDays", j.MonthDays)
	populate(objectMap, "monthlyOccurrences", j.MonthlyOccurrences)
	populate(objectMap, "weekDays", j.WeekDays)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type JobStatus.
func (j JobStatus) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "executionCount", j.ExecutionCount)
	populate(objectMap, "failureCount", j.FailureCount)
	populate(objectMap, "faultedCount", j.FaultedCount)
	populateTimeRFC3339(objectMap, "lastExecutionTime", j.LastExecutionTime)
	populateTimeRFC3339(objectMap, "nextExecutionTime", j.NextExecutionTime)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type JobStatus.
func (j *JobStatus) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "executionCount":
			err = unpopulate(val, &j.ExecutionCount)
			delete(rawMsg, key)
		case "failureCount":
			err = unpopulate(val, &j.FailureCount)
			delete(rawMsg, key)
		case "faultedCount":
			err = unpopulate(val, &j.FaultedCount)
			delete(rawMsg, key)
		case "lastExecutionTime":
			err = unpopulateTimeRFC3339(val, &j.LastExecutionTime)
			delete(rawMsg, key)
		case "nextExecutionTime":
			err = unpopulateTimeRFC3339(val, &j.NextExecutionTime)
			delete(rawMsg, key)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

// GetHTTPAuthentication implements the HTTPAuthenticationClassification interface for type OAuthAuthentication.
func (o *OAuthAuthentication) GetHTTPAuthentication() *HTTPAuthentication {
	return &HTTPAuthentication{
		Type: o.Type,
	}
}

// MarshalJSON implements the json.Marshaller interface for type OAuthAuthentication.
func (o OAuthAuthentication) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "audience", o.Audience)
	populate(objectMap, "clientId", o.ClientID)
	populate(objectMap, "secret", o.Secret)
	populate(objectMap, "tenant", o.Tenant)
	objectMap["type"] = HTTPAuthenticationTypeActiveDirectoryOAuth
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type OAuthAuthentication.
func (o *OAuthAuthentication) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "audience":
			err = unpopulate(val, &o.Audience)
			delete(rawMsg, key)
		case "clientId":
			err = unpopulate(val, &o.ClientID)
			delete(rawMsg, key)
		case "secret":
			err = unpopulate(val, &o.Secret)
			delete(rawMsg, key)
		case "tenant":
			err = unpopulate(val, &o.Tenant)
			delete(rawMsg, key)
		case "type":
			err = unpopulate(val, &o.Type)
			delete(rawMsg, key)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type ServiceBusBrokeredMessageProperties.
func (s ServiceBusBrokeredMessageProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "contentType", s.ContentType)
	populate(objectMap, "correlationId", s.CorrelationID)
	populate(objectMap, "forcePersistence", s.ForcePersistence)
	populate(objectMap, "label", s.Label)
	populate(objectMap, "messageId", s.MessageID)
	populate(objectMap, "partitionKey", s.PartitionKey)
	populate(objectMap, "replyTo", s.ReplyTo)
	populate(objectMap, "replyToSessionId", s.ReplyToSessionID)
	populateTimeRFC3339(objectMap, "scheduledEnqueueTimeUtc", s.ScheduledEnqueueTimeUTC)
	populate(objectMap, "sessionId", s.SessionID)
	populate(objectMap, "timeToLive", s.TimeToLive)
	populate(objectMap, "to", s.To)
	populate(objectMap, "viaPartitionKey", s.ViaPartitionKey)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ServiceBusBrokeredMessageProperties.
func (s *ServiceBusBrokeredMessageProperties) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "contentType":
			err = unpopulate(val, &s.ContentType)
			delete(rawMsg, key)
		case "correlationId":
			err = unpopulate(val, &s.CorrelationID)
			delete(rawMsg, key)
		case "forcePersistence":
			err = unpopulate(val, &s.ForcePersistence)
			delete(rawMsg, key)
		case "label":
			err = unpopulate(val, &s.Label)
			delete(rawMsg, key)
		case "messageId":
			err = unpopulate(val, &s.MessageID)
			delete(rawMsg, key)
		case "partitionKey":
			err = unpopulate(val, &s.PartitionKey)
			delete(rawMsg, key)
		case "replyTo":
			err = unpopulate(val, &s.ReplyTo)
			delete(rawMsg, key)
		case "replyToSessionId":
			err = unpopulate(val, &s.ReplyToSessionID)
			delete(rawMsg, key)
		case "scheduledEnqueueTimeUtc":
			err = unpopulateTimeRFC3339(val, &s.ScheduledEnqueueTimeUTC)
			delete(rawMsg, key)
		case "sessionId":
			err = unpopulate(val, &s.SessionID)
			delete(rawMsg, key)
		case "timeToLive":
			err = unpopulate(val, &s.TimeToLive)
			delete(rawMsg, key)
		case "to":
			err = unpopulate(val, &s.To)
			delete(rawMsg, key)
		case "viaPartitionKey":
			err = unpopulate(val, &s.ViaPartitionKey)
			delete(rawMsg, key)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type ServiceBusMessage.
func (s ServiceBusMessage) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "authentication", s.Authentication)
	populate(objectMap, "brokeredMessageProperties", s.BrokeredMessageProperties)
	populate(objectMap, "customMessageProperties", s.CustomMessageProperties)
	populate(objectMap, "message", s.Message)
	populate(objectMap, "namespace", s.Namespace)
	populate(objectMap, "transportType", s.TransportType)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type ServiceBusQueueMessage.
func (s ServiceBusQueueMessage) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "authentication", s.Authentication)
	populate(objectMap, "brokeredMessageProperties", s.BrokeredMessageProperties)
	populate(objectMap, "customMessageProperties", s.CustomMessageProperties)
	populate(objectMap, "message", s.Message)
	populate(objectMap, "namespace", s.Namespace)
	populate(objectMap, "queueName", s.QueueName)
	populate(objectMap, "transportType", s.TransportType)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type ServiceBusTopicMessage.
func (s ServiceBusTopicMessage) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "authentication", s.Authentication)
	populate(objectMap, "brokeredMessageProperties", s.BrokeredMessageProperties)
	populate(objectMap, "customMessageProperties", s.CustomMessageProperties)
	populate(objectMap, "message", s.Message)
	populate(objectMap, "namespace", s.Namespace)
	populate(objectMap, "topicPath", s.TopicPath)
	populate(objectMap, "transportType", s.TransportType)
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
