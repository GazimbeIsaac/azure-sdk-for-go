//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armresourcehealth

const (
	moduleName    = "armresourcehealth"
	moduleVersion = "v0.6.0"
)

// AvailabilityStateValues - Impacted resource status of the resource.
type AvailabilityStateValues string

const (
	AvailabilityStateValuesAvailable   AvailabilityStateValues = "Available"
	AvailabilityStateValuesDegraded    AvailabilityStateValues = "Degraded"
	AvailabilityStateValuesUnavailable AvailabilityStateValues = "Unavailable"
	AvailabilityStateValuesUnknown     AvailabilityStateValues = "Unknown"
)

// PossibleAvailabilityStateValuesValues returns the possible values for the AvailabilityStateValues const type.
func PossibleAvailabilityStateValuesValues() []AvailabilityStateValues {
	return []AvailabilityStateValues{
		AvailabilityStateValuesAvailable,
		AvailabilityStateValuesDegraded,
		AvailabilityStateValuesUnavailable,
		AvailabilityStateValuesUnknown,
	}
}

// ReasonChronicityTypes - Chronicity of the availability transition.
type ReasonChronicityTypes string

const (
	ReasonChronicityTypesPersistent ReasonChronicityTypes = "Persistent"
	ReasonChronicityTypesTransient  ReasonChronicityTypes = "Transient"
)

// PossibleReasonChronicityTypesValues returns the possible values for the ReasonChronicityTypes const type.
func PossibleReasonChronicityTypesValues() []ReasonChronicityTypes {
	return []ReasonChronicityTypes{
		ReasonChronicityTypesPersistent,
		ReasonChronicityTypesTransient,
	}
}

// ReasonTypeValues - When the resource's availabilityState is Unavailable, it describes where the health impacting event
// was originated.
type ReasonTypeValues string

const (
	ReasonTypeValuesPlanned       ReasonTypeValues = "Planned"
	ReasonTypeValuesUnplanned     ReasonTypeValues = "Unplanned"
	ReasonTypeValuesUserInitiated ReasonTypeValues = "UserInitiated"
)

// PossibleReasonTypeValuesValues returns the possible values for the ReasonTypeValues const type.
func PossibleReasonTypeValuesValues() []ReasonTypeValues {
	return []ReasonTypeValues{
		ReasonTypeValuesPlanned,
		ReasonTypeValuesUnplanned,
		ReasonTypeValuesUserInitiated,
	}
}
