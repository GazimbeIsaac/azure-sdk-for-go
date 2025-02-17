//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armquota

import (
	"encoding/json"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"reflect"
)

// MarshalJSON implements the json.Marshaller interface for type CreateGenericQuotaRequestParameters.
func (c CreateGenericQuotaRequestParameters) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "value", c.Value)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type CurrentQuotaLimitBase.
func (c CurrentQuotaLimitBase) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "id", c.ID)
	populate(objectMap, "name", c.Name)
	populate(objectMap, "properties", c.Properties)
	populate(objectMap, "type", c.Type)
	return json.Marshal(objectMap)
}

// GetLimitJSONObject implements the LimitJSONObjectClassification interface for type LimitJSONObject.
func (l *LimitJSONObject) GetLimitJSONObject() *LimitJSONObject { return l }

// GetLimitJSONObject implements the LimitJSONObjectClassification interface for type LimitObject.
func (l *LimitObject) GetLimitJSONObject() *LimitJSONObject {
	return &LimitJSONObject{
		LimitObjectType: l.LimitObjectType,
	}
}

// MarshalJSON implements the json.Marshaller interface for type LimitObject.
func (l LimitObject) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	objectMap["limitObjectType"] = LimitTypeLimitValue
	populate(objectMap, "limitType", l.LimitType)
	populate(objectMap, "value", l.Value)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type LimitObject.
func (l *LimitObject) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "limitObjectType":
			err = unpopulate(val, &l.LimitObjectType)
			delete(rawMsg, key)
		case "limitType":
			err = unpopulate(val, &l.LimitType)
			delete(rawMsg, key)
		case "value":
			err = unpopulate(val, &l.Value)
			delete(rawMsg, key)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type Limits.
func (l Limits) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "nextLink", l.NextLink)
	populate(objectMap, "value", l.Value)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type LimitsResponse.
func (l LimitsResponse) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "nextLink", l.NextLink)
	populate(objectMap, "value", l.Value)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type OperationList.
func (o OperationList) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "nextLink", o.NextLink)
	populate(objectMap, "value", o.Value)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type Properties.
func (p Properties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "isQuotaApplicable", p.IsQuotaApplicable)
	populate(objectMap, "limit", p.Limit)
	populate(objectMap, "name", p.Name)
	populate(objectMap, "properties", &p.Properties)
	populate(objectMap, "quotaPeriod", p.QuotaPeriod)
	populate(objectMap, "resourceType", p.ResourceType)
	populate(objectMap, "unit", p.Unit)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type Properties.
func (p *Properties) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "isQuotaApplicable":
			err = unpopulate(val, &p.IsQuotaApplicable)
			delete(rawMsg, key)
		case "limit":
			p.Limit, err = unmarshalLimitJSONObjectClassification(val)
			delete(rawMsg, key)
		case "name":
			err = unpopulate(val, &p.Name)
			delete(rawMsg, key)
		case "properties":
			err = unpopulate(val, &p.Properties)
			delete(rawMsg, key)
		case "quotaPeriod":
			err = unpopulate(val, &p.QuotaPeriod)
			delete(rawMsg, key)
		case "resourceType":
			err = unpopulate(val, &p.ResourceType)
			delete(rawMsg, key)
		case "unit":
			err = unpopulate(val, &p.Unit)
			delete(rawMsg, key)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type RequestDetailsList.
func (r RequestDetailsList) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "nextLink", r.NextLink)
	populate(objectMap, "value", r.Value)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type RequestOneResourceProperties.
func (r RequestOneResourceProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "currentValue", r.CurrentValue)
	populate(objectMap, "error", r.Error)
	populate(objectMap, "isQuotaApplicable", r.IsQuotaApplicable)
	populate(objectMap, "limit", r.Limit)
	populate(objectMap, "message", r.Message)
	populate(objectMap, "name", r.Name)
	populate(objectMap, "properties", &r.Properties)
	populate(objectMap, "provisioningState", r.ProvisioningState)
	populate(objectMap, "quotaPeriod", r.QuotaPeriod)
	populateTimeRFC3339(objectMap, "requestSubmitTime", r.RequestSubmitTime)
	populate(objectMap, "resourceType", r.ResourceType)
	populate(objectMap, "unit", r.Unit)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type RequestOneResourceProperties.
func (r *RequestOneResourceProperties) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "currentValue":
			err = unpopulate(val, &r.CurrentValue)
			delete(rawMsg, key)
		case "error":
			err = unpopulate(val, &r.Error)
			delete(rawMsg, key)
		case "isQuotaApplicable":
			err = unpopulate(val, &r.IsQuotaApplicable)
			delete(rawMsg, key)
		case "limit":
			err = unpopulate(val, &r.Limit)
			delete(rawMsg, key)
		case "message":
			err = unpopulate(val, &r.Message)
			delete(rawMsg, key)
		case "name":
			err = unpopulate(val, &r.Name)
			delete(rawMsg, key)
		case "properties":
			err = unpopulate(val, &r.Properties)
			delete(rawMsg, key)
		case "provisioningState":
			err = unpopulate(val, &r.ProvisioningState)
			delete(rawMsg, key)
		case "quotaPeriod":
			err = unpopulate(val, &r.QuotaPeriod)
			delete(rawMsg, key)
		case "requestSubmitTime":
			err = unpopulateTimeRFC3339(val, &r.RequestSubmitTime)
			delete(rawMsg, key)
		case "resourceType":
			err = unpopulate(val, &r.ResourceType)
			delete(rawMsg, key)
		case "unit":
			err = unpopulate(val, &r.Unit)
			delete(rawMsg, key)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type RequestProperties.
func (r RequestProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "error", r.Error)
	populate(objectMap, "message", r.Message)
	populate(objectMap, "provisioningState", r.ProvisioningState)
	populateTimeRFC3339(objectMap, "requestSubmitTime", r.RequestSubmitTime)
	populate(objectMap, "value", r.Value)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type RequestProperties.
func (r *RequestProperties) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "error":
			err = unpopulate(val, &r.Error)
			delete(rawMsg, key)
		case "message":
			err = unpopulate(val, &r.Message)
			delete(rawMsg, key)
		case "provisioningState":
			err = unpopulate(val, &r.ProvisioningState)
			delete(rawMsg, key)
		case "requestSubmitTime":
			err = unpopulateTimeRFC3339(val, &r.RequestSubmitTime)
			delete(rawMsg, key)
		case "value":
			err = unpopulate(val, &r.Value)
			delete(rawMsg, key)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type ServiceError.
func (s ServiceError) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "code", s.Code)
	populate(objectMap, "details", s.Details)
	populate(objectMap, "message", s.Message)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type SubRequest.
func (s SubRequest) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "limit", s.Limit)
	populate(objectMap, "message", s.Message)
	populate(objectMap, "name", s.Name)
	populate(objectMap, "provisioningState", s.ProvisioningState)
	populate(objectMap, "resourceType", s.ResourceType)
	populate(objectMap, "subRequestId", s.SubRequestID)
	populate(objectMap, "unit", s.Unit)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type SubRequest.
func (s *SubRequest) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "limit":
			s.Limit, err = unmarshalLimitJSONObjectClassification(val)
			delete(rawMsg, key)
		case "message":
			err = unpopulate(val, &s.Message)
			delete(rawMsg, key)
		case "name":
			err = unpopulate(val, &s.Name)
			delete(rawMsg, key)
		case "provisioningState":
			err = unpopulate(val, &s.ProvisioningState)
			delete(rawMsg, key)
		case "resourceType":
			err = unpopulate(val, &s.ResourceType)
			delete(rawMsg, key)
		case "subRequestId":
			err = unpopulate(val, &s.SubRequestID)
			delete(rawMsg, key)
		case "unit":
			err = unpopulate(val, &s.Unit)
			delete(rawMsg, key)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type UsagesLimits.
func (u UsagesLimits) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "nextLink", u.NextLink)
	populate(objectMap, "value", u.Value)
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
