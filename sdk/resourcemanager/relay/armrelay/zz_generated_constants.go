//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armrelay

const (
	moduleName    = "armrelay"
	moduleVersion = "v0.5.0"
)

type AccessRights string

const (
	AccessRightsManage AccessRights = "Manage"
	AccessRightsSend   AccessRights = "Send"
	AccessRightsListen AccessRights = "Listen"
)

// PossibleAccessRightsValues returns the possible values for the AccessRights const type.
func PossibleAccessRightsValues() []AccessRights {
	return []AccessRights{
		AccessRightsManage,
		AccessRightsSend,
		AccessRightsListen,
	}
}

// KeyType - The access key to regenerate.
type KeyType string

const (
	KeyTypePrimaryKey   KeyType = "PrimaryKey"
	KeyTypeSecondaryKey KeyType = "SecondaryKey"
)

// PossibleKeyTypeValues returns the possible values for the KeyType const type.
func PossibleKeyTypeValues() []KeyType {
	return []KeyType{
		KeyTypePrimaryKey,
		KeyTypeSecondaryKey,
	}
}

type ProvisioningStateEnum string

const (
	ProvisioningStateEnumCreated   ProvisioningStateEnum = "Created"
	ProvisioningStateEnumSucceeded ProvisioningStateEnum = "Succeeded"
	ProvisioningStateEnumDeleted   ProvisioningStateEnum = "Deleted"
	ProvisioningStateEnumFailed    ProvisioningStateEnum = "Failed"
	ProvisioningStateEnumUpdating  ProvisioningStateEnum = "Updating"
	ProvisioningStateEnumUnknown   ProvisioningStateEnum = "Unknown"
)

// PossibleProvisioningStateEnumValues returns the possible values for the ProvisioningStateEnum const type.
func PossibleProvisioningStateEnumValues() []ProvisioningStateEnum {
	return []ProvisioningStateEnum{
		ProvisioningStateEnumCreated,
		ProvisioningStateEnumSucceeded,
		ProvisioningStateEnumDeleted,
		ProvisioningStateEnumFailed,
		ProvisioningStateEnumUpdating,
		ProvisioningStateEnumUnknown,
	}
}

// Relaytype - WCF relay type.
type Relaytype string

const (
	RelaytypeNetTCP Relaytype = "NetTcp"
	RelaytypeHTTP   Relaytype = "Http"
)

// PossibleRelaytypeValues returns the possible values for the Relaytype const type.
func PossibleRelaytypeValues() []Relaytype {
	return []Relaytype{
		RelaytypeNetTCP,
		RelaytypeHTTP,
	}
}

// UnavailableReason - Specifies the reason for the unavailability of the service.
type UnavailableReason string

const (
	UnavailableReasonNone                                  UnavailableReason = "None"
	UnavailableReasonInvalidName                           UnavailableReason = "InvalidName"
	UnavailableReasonSubscriptionIsDisabled                UnavailableReason = "SubscriptionIsDisabled"
	UnavailableReasonNameInUse                             UnavailableReason = "NameInUse"
	UnavailableReasonNameInLockdown                        UnavailableReason = "NameInLockdown"
	UnavailableReasonTooManyNamespaceInCurrentSubscription UnavailableReason = "TooManyNamespaceInCurrentSubscription"
)

// PossibleUnavailableReasonValues returns the possible values for the UnavailableReason const type.
func PossibleUnavailableReasonValues() []UnavailableReason {
	return []UnavailableReason{
		UnavailableReasonNone,
		UnavailableReasonInvalidName,
		UnavailableReasonSubscriptionIsDisabled,
		UnavailableReasonNameInUse,
		UnavailableReasonNameInLockdown,
		UnavailableReasonTooManyNamespaceInCurrentSubscription,
	}
}
