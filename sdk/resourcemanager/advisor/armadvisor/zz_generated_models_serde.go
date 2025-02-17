//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armadvisor

import (
	"encoding/json"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"reflect"
)

// MarshalJSON implements the json.Marshaller interface for type ConfigDataProperties.
func (c ConfigDataProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "digests", c.Digests)
	populate(objectMap, "exclude", c.Exclude)
	populate(objectMap, "lowCpuThreshold", c.LowCPUThreshold)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type ConfigurationListResult.
func (c ConfigurationListResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "nextLink", c.NextLink)
	populate(objectMap, "value", c.Value)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type DigestConfig.
func (d DigestConfig) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "actionGroupResourceId", d.ActionGroupResourceID)
	populate(objectMap, "categories", d.Categories)
	populate(objectMap, "frequency", d.Frequency)
	populate(objectMap, "language", d.Language)
	populate(objectMap, "name", d.Name)
	populate(objectMap, "state", d.State)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type MetadataEntityListResult.
func (m MetadataEntityListResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "nextLink", m.NextLink)
	populate(objectMap, "value", m.Value)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type MetadataEntityProperties.
func (m MetadataEntityProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "applicableScenarios", m.ApplicableScenarios)
	populate(objectMap, "dependsOn", m.DependsOn)
	populate(objectMap, "displayName", m.DisplayName)
	populate(objectMap, "supportedValues", m.SupportedValues)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type OperationEntityListResult.
func (o OperationEntityListResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "nextLink", o.NextLink)
	populate(objectMap, "value", o.Value)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type RecommendationProperties.
func (r RecommendationProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "actions", r.Actions)
	populate(objectMap, "category", r.Category)
	populate(objectMap, "description", r.Description)
	populate(objectMap, "exposedMetadataProperties", r.ExposedMetadataProperties)
	populate(objectMap, "extendedProperties", r.ExtendedProperties)
	populate(objectMap, "impact", r.Impact)
	populate(objectMap, "impactedField", r.ImpactedField)
	populate(objectMap, "impactedValue", r.ImpactedValue)
	populate(objectMap, "label", r.Label)
	populateTimeRFC3339(objectMap, "lastUpdated", r.LastUpdated)
	populate(objectMap, "learnMoreLink", r.LearnMoreLink)
	populate(objectMap, "metadata", r.Metadata)
	populate(objectMap, "potentialBenefits", r.PotentialBenefits)
	populate(objectMap, "recommendationTypeId", r.RecommendationTypeID)
	populate(objectMap, "remediation", r.Remediation)
	populate(objectMap, "resourceMetadata", r.ResourceMetadata)
	populate(objectMap, "risk", r.Risk)
	populate(objectMap, "shortDescription", r.ShortDescription)
	populate(objectMap, "suppressionIds", r.SuppressionIDs)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type RecommendationProperties.
func (r *RecommendationProperties) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "actions":
			err = unpopulate(val, &r.Actions)
			delete(rawMsg, key)
		case "category":
			err = unpopulate(val, &r.Category)
			delete(rawMsg, key)
		case "description":
			err = unpopulate(val, &r.Description)
			delete(rawMsg, key)
		case "exposedMetadataProperties":
			err = unpopulate(val, &r.ExposedMetadataProperties)
			delete(rawMsg, key)
		case "extendedProperties":
			err = unpopulate(val, &r.ExtendedProperties)
			delete(rawMsg, key)
		case "impact":
			err = unpopulate(val, &r.Impact)
			delete(rawMsg, key)
		case "impactedField":
			err = unpopulate(val, &r.ImpactedField)
			delete(rawMsg, key)
		case "impactedValue":
			err = unpopulate(val, &r.ImpactedValue)
			delete(rawMsg, key)
		case "label":
			err = unpopulate(val, &r.Label)
			delete(rawMsg, key)
		case "lastUpdated":
			err = unpopulateTimeRFC3339(val, &r.LastUpdated)
			delete(rawMsg, key)
		case "learnMoreLink":
			err = unpopulate(val, &r.LearnMoreLink)
			delete(rawMsg, key)
		case "metadata":
			err = unpopulate(val, &r.Metadata)
			delete(rawMsg, key)
		case "potentialBenefits":
			err = unpopulate(val, &r.PotentialBenefits)
			delete(rawMsg, key)
		case "recommendationTypeId":
			err = unpopulate(val, &r.RecommendationTypeID)
			delete(rawMsg, key)
		case "remediation":
			err = unpopulate(val, &r.Remediation)
			delete(rawMsg, key)
		case "resourceMetadata":
			err = unpopulate(val, &r.ResourceMetadata)
			delete(rawMsg, key)
		case "risk":
			err = unpopulate(val, &r.Risk)
			delete(rawMsg, key)
		case "shortDescription":
			err = unpopulate(val, &r.ShortDescription)
			delete(rawMsg, key)
		case "suppressionIds":
			err = unpopulate(val, &r.SuppressionIDs)
			delete(rawMsg, key)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type ResourceMetadata.
func (r ResourceMetadata) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "action", r.Action)
	populate(objectMap, "plural", r.Plural)
	populate(objectMap, "resourceId", r.ResourceID)
	populate(objectMap, "singular", r.Singular)
	populate(objectMap, "source", r.Source)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type ResourceRecommendationBaseListResult.
func (r ResourceRecommendationBaseListResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "nextLink", r.NextLink)
	populate(objectMap, "value", r.Value)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type SuppressionContractListResult.
func (s SuppressionContractListResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "nextLink", s.NextLink)
	populate(objectMap, "value", s.Value)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type SuppressionProperties.
func (s SuppressionProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populateTimeRFC3339(objectMap, "expirationTimeStamp", s.ExpirationTimeStamp)
	populate(objectMap, "suppressionId", s.SuppressionID)
	populate(objectMap, "ttl", s.TTL)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type SuppressionProperties.
func (s *SuppressionProperties) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "expirationTimeStamp":
			err = unpopulateTimeRFC3339(val, &s.ExpirationTimeStamp)
			delete(rawMsg, key)
		case "suppressionId":
			err = unpopulate(val, &s.SuppressionID)
			delete(rawMsg, key)
		case "ttl":
			err = unpopulate(val, &s.TTL)
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
