//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcompute_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/compute/resource-manager/Microsoft.Compute/stable/2021-11-01/examples/compute/ListVirtualMachineScaleSetsInASubscriptionByLocation.json
func ExampleVirtualMachineScaleSetsClient_NewListByLocationPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armcompute.NewVirtualMachineScaleSetsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := client.NewListByLocationPager("<location>",
		nil)
	for pager.More() {
		nextResult, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
			return
		}
		for _, v := range nextResult.Value {
			// TODO: use page item
			_ = v
		}
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/compute/resource-manager/Microsoft.Compute/stable/2021-11-01/examples/compute/CreateAScaleSetWithExtensionsSuppressFailuresEnabled.json
func ExampleVirtualMachineScaleSetsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armcompute.NewVirtualMachineScaleSetsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<resource-group-name>",
		"<vm-scale-set-name>",
		armcompute.VirtualMachineScaleSet{
			Location: to.Ptr("<location>"),
			Properties: &armcompute.VirtualMachineScaleSetProperties{
				Overprovision: to.Ptr(true),
				UpgradePolicy: &armcompute.UpgradePolicy{
					Mode: to.Ptr(armcompute.UpgradeModeManual),
				},
				VirtualMachineProfile: &armcompute.VirtualMachineScaleSetVMProfile{
					DiagnosticsProfile: &armcompute.DiagnosticsProfile{
						BootDiagnostics: &armcompute.BootDiagnostics{
							Enabled:    to.Ptr(true),
							StorageURI: to.Ptr("<storage-uri>"),
						},
					},
					ExtensionProfile: &armcompute.VirtualMachineScaleSetExtensionProfile{
						Extensions: []*armcompute.VirtualMachineScaleSetExtension{
							{
								Name: to.Ptr("<name>"),
								Properties: &armcompute.VirtualMachineScaleSetExtensionProperties{
									Type:                    to.Ptr("<type>"),
									AutoUpgradeMinorVersion: to.Ptr(false),
									Publisher:               to.Ptr("<publisher>"),
									Settings:                map[string]interface{}{},
									SuppressFailures:        to.Ptr(true),
									TypeHandlerVersion:      to.Ptr("<type-handler-version>"),
								},
							}},
					},
					NetworkProfile: &armcompute.VirtualMachineScaleSetNetworkProfile{
						NetworkInterfaceConfigurations: []*armcompute.VirtualMachineScaleSetNetworkConfiguration{
							{
								Name: to.Ptr("<name>"),
								Properties: &armcompute.VirtualMachineScaleSetNetworkConfigurationProperties{
									EnableIPForwarding: to.Ptr(true),
									IPConfigurations: []*armcompute.VirtualMachineScaleSetIPConfiguration{
										{
											Name: to.Ptr("<name>"),
											Properties: &armcompute.VirtualMachineScaleSetIPConfigurationProperties{
												Subnet: &armcompute.APIEntityReference{
													ID: to.Ptr("<id>"),
												},
											},
										}},
									Primary: to.Ptr(true),
								},
							}},
					},
					OSProfile: &armcompute.VirtualMachineScaleSetOSProfile{
						AdminPassword:      to.Ptr("<admin-password>"),
						AdminUsername:      to.Ptr("<admin-username>"),
						ComputerNamePrefix: to.Ptr("<computer-name-prefix>"),
					},
					StorageProfile: &armcompute.VirtualMachineScaleSetStorageProfile{
						ImageReference: &armcompute.ImageReference{
							Offer:     to.Ptr("<offer>"),
							Publisher: to.Ptr("<publisher>"),
							SKU:       to.Ptr("<sku>"),
							Version:   to.Ptr("<version>"),
						},
						OSDisk: &armcompute.VirtualMachineScaleSetOSDisk{
							Caching:      to.Ptr(armcompute.CachingTypesReadWrite),
							CreateOption: to.Ptr(armcompute.DiskCreateOptionTypesFromImage),
							ManagedDisk: &armcompute.VirtualMachineScaleSetManagedDiskParameters{
								StorageAccountType: to.Ptr(armcompute.StorageAccountTypesStandardLRS),
							},
						},
					},
				},
			},
			SKU: &armcompute.SKU{
				Name:     to.Ptr("<name>"),
				Capacity: to.Ptr[int64](3),
				Tier:     to.Ptr("<tier>"),
			},
		},
		&armcompute.VirtualMachineScaleSetsClientBeginCreateOrUpdateOptions{ResumeToken: ""})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/compute/resource-manager/Microsoft.Compute/stable/2021-11-01/examples/compute/VirtualMachineScaleSets_Update_MaximumSet_Gen.json
func ExampleVirtualMachineScaleSetsClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armcompute.NewVirtualMachineScaleSetsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginUpdate(ctx,
		"<resource-group-name>",
		"<vm-scale-set-name>",
		armcompute.VirtualMachineScaleSetUpdate{
			Tags: map[string]*string{
				"key246": to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaa"),
			},
			Identity: &armcompute.VirtualMachineScaleSetIdentity{
				Type: to.Ptr(armcompute.ResourceIdentityTypeSystemAssigned),
				UserAssignedIdentities: map[string]*armcompute.VirtualMachineScaleSetIdentityUserAssignedIdentitiesValue{
					"key3951": {},
				},
			},
			Plan: &armcompute.Plan{
				Name:          to.Ptr("<name>"),
				Product:       to.Ptr("<product>"),
				PromotionCode: to.Ptr("<promotion-code>"),
				Publisher:     to.Ptr("<publisher>"),
			},
			Properties: &armcompute.VirtualMachineScaleSetUpdateProperties{
				AdditionalCapabilities: &armcompute.AdditionalCapabilities{
					HibernationEnabled: to.Ptr(true),
					UltraSSDEnabled:    to.Ptr(true),
				},
				AutomaticRepairsPolicy: &armcompute.AutomaticRepairsPolicy{
					Enabled:     to.Ptr(true),
					GracePeriod: to.Ptr("<grace-period>"),
				},
				DoNotRunExtensionsOnOverprovisionedVMs: to.Ptr(true),
				Overprovision:                          to.Ptr(true),
				ProximityPlacementGroup: &armcompute.SubResource{
					ID: to.Ptr("<id>"),
				},
				ScaleInPolicy: &armcompute.ScaleInPolicy{
					ForceDeletion: to.Ptr(true),
					Rules: []*armcompute.VirtualMachineScaleSetScaleInRules{
						to.Ptr(armcompute.VirtualMachineScaleSetScaleInRulesOldestVM)},
				},
				SinglePlacementGroup: to.Ptr(true),
				UpgradePolicy: &armcompute.UpgradePolicy{
					AutomaticOSUpgradePolicy: &armcompute.AutomaticOSUpgradePolicy{
						DisableAutomaticRollback: to.Ptr(true),
						EnableAutomaticOSUpgrade: to.Ptr(true),
					},
					Mode: to.Ptr(armcompute.UpgradeModeManual),
					RollingUpgradePolicy: &armcompute.RollingUpgradePolicy{
						EnableCrossZoneUpgrade:              to.Ptr(true),
						MaxBatchInstancePercent:             to.Ptr[int32](49),
						MaxUnhealthyInstancePercent:         to.Ptr[int32](81),
						MaxUnhealthyUpgradedInstancePercent: to.Ptr[int32](98),
						PauseTimeBetweenBatches:             to.Ptr("<pause-time-between-batches>"),
						PrioritizeUnhealthyInstances:        to.Ptr(true),
					},
				},
				VirtualMachineProfile: &armcompute.VirtualMachineScaleSetUpdateVMProfile{
					BillingProfile: &armcompute.BillingProfile{
						MaxPrice: to.Ptr[float64](-1),
					},
					DiagnosticsProfile: &armcompute.DiagnosticsProfile{
						BootDiagnostics: &armcompute.BootDiagnostics{
							Enabled:    to.Ptr(true),
							StorageURI: to.Ptr("<storage-uri>"),
						},
					},
					ExtensionProfile: &armcompute.VirtualMachineScaleSetExtensionProfile{
						ExtensionsTimeBudget: to.Ptr("<extensions-time-budget>"),
						Extensions: []*armcompute.VirtualMachineScaleSetExtension{
							{
								Name: to.Ptr("<name>"),
								Properties: &armcompute.VirtualMachineScaleSetExtensionProperties{
									Type:                    to.Ptr("<type>"),
									AutoUpgradeMinorVersion: to.Ptr(true),
									EnableAutomaticUpgrade:  to.Ptr(true),
									ForceUpdateTag:          to.Ptr("<force-update-tag>"),
									ProtectedSettings:       map[string]interface{}{},
									ProvisionAfterExtensions: []*string{
										to.Ptr("aa")},
									Publisher:          to.Ptr("<publisher>"),
									Settings:           map[string]interface{}{},
									SuppressFailures:   to.Ptr(true),
									TypeHandlerVersion: to.Ptr("<type-handler-version>"),
								},
							}},
					},
					LicenseType: to.Ptr("<license-type>"),
					NetworkProfile: &armcompute.VirtualMachineScaleSetUpdateNetworkProfile{
						HealthProbe: &armcompute.APIEntityReference{
							ID: to.Ptr("<id>"),
						},
						NetworkAPIVersion: to.Ptr(armcompute.NetworkAPIVersionTwoThousandTwenty1101),
						NetworkInterfaceConfigurations: []*armcompute.VirtualMachineScaleSetUpdateNetworkConfiguration{
							{
								ID:   to.Ptr("<id>"),
								Name: to.Ptr("<name>"),
								Properties: &armcompute.VirtualMachineScaleSetUpdateNetworkConfigurationProperties{
									DeleteOption: to.Ptr(armcompute.DeleteOptionsDelete),
									DNSSettings: &armcompute.VirtualMachineScaleSetNetworkConfigurationDNSSettings{
										DNSServers: []*string{},
									},
									EnableAcceleratedNetworking: to.Ptr(true),
									EnableFpga:                  to.Ptr(true),
									EnableIPForwarding:          to.Ptr(true),
									IPConfigurations: []*armcompute.VirtualMachineScaleSetUpdateIPConfiguration{
										{
											ID:   to.Ptr("<id>"),
											Name: to.Ptr("<name>"),
											Properties: &armcompute.VirtualMachineScaleSetUpdateIPConfigurationProperties{
												ApplicationGatewayBackendAddressPools: []*armcompute.SubResource{
													{
														ID: to.Ptr("<id>"),
													}},
												ApplicationSecurityGroups: []*armcompute.SubResource{
													{
														ID: to.Ptr("<id>"),
													}},
												LoadBalancerBackendAddressPools: []*armcompute.SubResource{
													{
														ID: to.Ptr("<id>"),
													}},
												LoadBalancerInboundNatPools: []*armcompute.SubResource{
													{
														ID: to.Ptr("<id>"),
													}},
												Primary:                 to.Ptr(true),
												PrivateIPAddressVersion: to.Ptr(armcompute.IPVersionIPv4),
												PublicIPAddressConfiguration: &armcompute.VirtualMachineScaleSetUpdatePublicIPAddressConfiguration{
													Name: to.Ptr("<name>"),
													Properties: &armcompute.VirtualMachineScaleSetUpdatePublicIPAddressConfigurationProperties{
														DeleteOption: to.Ptr(armcompute.DeleteOptionsDelete),
														DNSSettings: &armcompute.VirtualMachineScaleSetPublicIPAddressConfigurationDNSSettings{
															DomainNameLabel: to.Ptr("<domain-name-label>"),
														},
														IdleTimeoutInMinutes: to.Ptr[int32](3),
													},
												},
												Subnet: &armcompute.APIEntityReference{
													ID: to.Ptr("<id>"),
												},
											},
										}},
									NetworkSecurityGroup: &armcompute.SubResource{
										ID: to.Ptr("<id>"),
									},
									Primary: to.Ptr(true),
								},
							}},
					},
					OSProfile: &armcompute.VirtualMachineScaleSetUpdateOSProfile{
						CustomData: to.Ptr("<custom-data>"),
						LinuxConfiguration: &armcompute.LinuxConfiguration{
							DisablePasswordAuthentication: to.Ptr(true),
							PatchSettings: &armcompute.LinuxPatchSettings{
								AssessmentMode: to.Ptr(armcompute.LinuxPatchAssessmentModeImageDefault),
								PatchMode:      to.Ptr(armcompute.LinuxVMGuestPatchModeImageDefault),
							},
							ProvisionVMAgent: to.Ptr(true),
							SSH: &armcompute.SSHConfiguration{
								PublicKeys: []*armcompute.SSHPublicKey{
									{
										Path:    to.Ptr("<path>"),
										KeyData: to.Ptr("<key-data>"),
									}},
							},
						},
						Secrets: []*armcompute.VaultSecretGroup{
							{
								SourceVault: &armcompute.SubResource{
									ID: to.Ptr("<id>"),
								},
								VaultCertificates: []*armcompute.VaultCertificate{
									{
										CertificateStore: to.Ptr("<certificate-store>"),
										CertificateURL:   to.Ptr("<certificate-url>"),
									}},
							}},
						WindowsConfiguration: &armcompute.WindowsConfiguration{
							AdditionalUnattendContent: []*armcompute.AdditionalUnattendContent{
								{
									ComponentName: to.Ptr("<component-name>"),
									Content:       to.Ptr("<content>"),
									PassName:      to.Ptr("<pass-name>"),
									SettingName:   to.Ptr(armcompute.SettingNamesAutoLogon),
								}},
							EnableAutomaticUpdates: to.Ptr(true),
							PatchSettings: &armcompute.PatchSettings{
								AssessmentMode:    to.Ptr(armcompute.WindowsPatchAssessmentModeImageDefault),
								EnableHotpatching: to.Ptr(true),
								PatchMode:         to.Ptr(armcompute.WindowsVMGuestPatchModeAutomaticByOS),
							},
							ProvisionVMAgent: to.Ptr(true),
							TimeZone:         to.Ptr("<time-zone>"),
							WinRM: &armcompute.WinRMConfiguration{
								Listeners: []*armcompute.WinRMListener{
									{
										CertificateURL: to.Ptr("<certificate-url>"),
										Protocol:       to.Ptr(armcompute.ProtocolTypesHTTP),
									}},
							},
						},
					},
					ScheduledEventsProfile: &armcompute.ScheduledEventsProfile{
						TerminateNotificationProfile: &armcompute.TerminateNotificationProfile{
							Enable:           to.Ptr(true),
							NotBeforeTimeout: to.Ptr("<not-before-timeout>"),
						},
					},
					SecurityProfile: &armcompute.SecurityProfile{
						EncryptionAtHost: to.Ptr(true),
						SecurityType:     to.Ptr(armcompute.SecurityTypesTrustedLaunch),
						UefiSettings: &armcompute.UefiSettings{
							SecureBootEnabled: to.Ptr(true),
							VTpmEnabled:       to.Ptr(true),
						},
					},
					StorageProfile: &armcompute.VirtualMachineScaleSetUpdateStorageProfile{
						DataDisks: []*armcompute.VirtualMachineScaleSetDataDisk{
							{
								Name:              to.Ptr("<name>"),
								Caching:           to.Ptr(armcompute.CachingTypesNone),
								CreateOption:      to.Ptr(armcompute.DiskCreateOptionTypesEmpty),
								DiskIOPSReadWrite: to.Ptr[int64](28),
								DiskMBpsReadWrite: to.Ptr[int64](15),
								DiskSizeGB:        to.Ptr[int32](1023),
								Lun:               to.Ptr[int32](26),
								ManagedDisk: &armcompute.VirtualMachineScaleSetManagedDiskParameters{
									DiskEncryptionSet: &armcompute.DiskEncryptionSetParameters{
										ID: to.Ptr("<id>"),
									},
									StorageAccountType: to.Ptr(armcompute.StorageAccountTypesStandardLRS),
								},
								WriteAcceleratorEnabled: to.Ptr(true),
							}},
						ImageReference: &armcompute.ImageReference{
							ID:                   to.Ptr("<id>"),
							Offer:                to.Ptr("<offer>"),
							Publisher:            to.Ptr("<publisher>"),
							SharedGalleryImageID: to.Ptr("<shared-gallery-image-id>"),
							SKU:                  to.Ptr("<sku>"),
							Version:              to.Ptr("<version>"),
						},
						OSDisk: &armcompute.VirtualMachineScaleSetUpdateOSDisk{
							Caching:    to.Ptr(armcompute.CachingTypesReadWrite),
							DiskSizeGB: to.Ptr[int32](6),
							Image: &armcompute.VirtualHardDisk{
								URI: to.Ptr("<uri>"),
							},
							ManagedDisk: &armcompute.VirtualMachineScaleSetManagedDiskParameters{
								DiskEncryptionSet: &armcompute.DiskEncryptionSetParameters{
									ID: to.Ptr("<id>"),
								},
								StorageAccountType: to.Ptr(armcompute.StorageAccountTypesStandardLRS),
							},
							VhdContainers: []*string{
								to.Ptr("aa")},
							WriteAcceleratorEnabled: to.Ptr(true),
						},
					},
					UserData: to.Ptr("<user-data>"),
				},
			},
			SKU: &armcompute.SKU{
				Name:     to.Ptr("<name>"),
				Capacity: to.Ptr[int64](7),
				Tier:     to.Ptr("<tier>"),
			},
		},
		&armcompute.VirtualMachineScaleSetsClientBeginUpdateOptions{ResumeToken: ""})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/compute/resource-manager/Microsoft.Compute/stable/2021-11-01/examples/compute/ForceDeleteVirtualMachineScaleSets.json
func ExampleVirtualMachineScaleSetsClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armcompute.NewVirtualMachineScaleSetsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginDelete(ctx,
		"<resource-group-name>",
		"<vm-scale-set-name>",
		&armcompute.VirtualMachineScaleSetsClientBeginDeleteOptions{ForceDeletion: to.Ptr(true),
			ResumeToken: "",
		})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/compute/resource-manager/Microsoft.Compute/stable/2021-11-01/examples/compute/GetVirtualMachineScaleSet.json
func ExampleVirtualMachineScaleSetsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armcompute.NewVirtualMachineScaleSetsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Get(ctx,
		"<resource-group-name>",
		"<vm-scale-set-name>",
		&armcompute.VirtualMachineScaleSetsClientGetOptions{Expand: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/compute/resource-manager/Microsoft.Compute/stable/2021-11-01/examples/compute/VirtualMachineScaleSets_Deallocate_MaximumSet_Gen.json
func ExampleVirtualMachineScaleSetsClient_BeginDeallocate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armcompute.NewVirtualMachineScaleSetsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginDeallocate(ctx,
		"<resource-group-name>",
		"<vm-scale-set-name>",
		&armcompute.VirtualMachineScaleSetsClientBeginDeallocateOptions{VMInstanceIDs: &armcompute.VirtualMachineScaleSetVMInstanceIDs{
			InstanceIDs: []*string{
				to.Ptr("aaaaaaaaaaaaaaaaa")},
		},
			ResumeToken: "",
		})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/compute/resource-manager/Microsoft.Compute/stable/2021-11-01/examples/compute/VirtualMachineScaleSets_DeleteInstances_MaximumSet_Gen.json
func ExampleVirtualMachineScaleSetsClient_BeginDeleteInstances() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armcompute.NewVirtualMachineScaleSetsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginDeleteInstances(ctx,
		"<resource-group-name>",
		"<vm-scale-set-name>",
		armcompute.VirtualMachineScaleSetVMInstanceRequiredIDs{
			InstanceIDs: []*string{
				to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaaa")},
		},
		&armcompute.VirtualMachineScaleSetsClientBeginDeleteInstancesOptions{ForceDeletion: to.Ptr(true),
			ResumeToken: "",
		})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/compute/resource-manager/Microsoft.Compute/stable/2021-11-01/examples/compute/VirtualMachineScaleSets_GetInstanceView_MaximumSet_Gen.json
func ExampleVirtualMachineScaleSetsClient_GetInstanceView() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armcompute.NewVirtualMachineScaleSetsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.GetInstanceView(ctx,
		"<resource-group-name>",
		"<vm-scale-set-name>",
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/compute/resource-manager/Microsoft.Compute/stable/2021-11-01/examples/compute/VirtualMachineScaleSets_List_MaximumSet_Gen.json
func ExampleVirtualMachineScaleSetsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armcompute.NewVirtualMachineScaleSetsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := client.NewListPager("<resource-group-name>",
		nil)
	for pager.More() {
		nextResult, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
			return
		}
		for _, v := range nextResult.Value {
			// TODO: use page item
			_ = v
		}
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/compute/resource-manager/Microsoft.Compute/stable/2021-11-01/examples/compute/VirtualMachineScaleSets_ListAll_MaximumSet_Gen.json
func ExampleVirtualMachineScaleSetsClient_NewListAllPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armcompute.NewVirtualMachineScaleSetsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := client.NewListAllPager(nil)
	for pager.More() {
		nextResult, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
			return
		}
		for _, v := range nextResult.Value {
			// TODO: use page item
			_ = v
		}
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/compute/resource-manager/Microsoft.Compute/stable/2021-11-01/examples/compute/VirtualMachineScaleSets_GetOSUpgradeHistory_MaximumSet_Gen.json
func ExampleVirtualMachineScaleSetsClient_NewGetOSUpgradeHistoryPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armcompute.NewVirtualMachineScaleSetsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := client.NewGetOSUpgradeHistoryPager("<resource-group-name>",
		"<vm-scale-set-name>",
		nil)
	for pager.More() {
		nextResult, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
			return
		}
		for _, v := range nextResult.Value {
			// TODO: use page item
			_ = v
		}
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/compute/resource-manager/Microsoft.Compute/stable/2021-11-01/examples/compute/VirtualMachineScaleSets_PowerOff_MaximumSet_Gen.json
func ExampleVirtualMachineScaleSetsClient_BeginPowerOff() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armcompute.NewVirtualMachineScaleSetsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginPowerOff(ctx,
		"<resource-group-name>",
		"<vm-scale-set-name>",
		&armcompute.VirtualMachineScaleSetsClientBeginPowerOffOptions{SkipShutdown: to.Ptr(true),
			VMInstanceIDs: &armcompute.VirtualMachineScaleSetVMInstanceIDs{
				InstanceIDs: []*string{
					to.Ptr("aaaaaaaaaaaaaaaaa")},
			},
			ResumeToken: "",
		})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/compute/resource-manager/Microsoft.Compute/stable/2021-11-01/examples/compute/VirtualMachineScaleSets_Restart_MaximumSet_Gen.json
func ExampleVirtualMachineScaleSetsClient_BeginRestart() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armcompute.NewVirtualMachineScaleSetsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginRestart(ctx,
		"<resource-group-name>",
		"<vm-scale-set-name>",
		&armcompute.VirtualMachineScaleSetsClientBeginRestartOptions{VMInstanceIDs: &armcompute.VirtualMachineScaleSetVMInstanceIDs{
			InstanceIDs: []*string{
				to.Ptr("aaaaaaaaaaaaaaaaa")},
		},
			ResumeToken: "",
		})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/compute/resource-manager/Microsoft.Compute/stable/2021-11-01/examples/compute/VirtualMachineScaleSets_Start_MaximumSet_Gen.json
func ExampleVirtualMachineScaleSetsClient_BeginStart() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armcompute.NewVirtualMachineScaleSetsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginStart(ctx,
		"<resource-group-name>",
		"<vm-scale-set-name>",
		&armcompute.VirtualMachineScaleSetsClientBeginStartOptions{VMInstanceIDs: &armcompute.VirtualMachineScaleSetVMInstanceIDs{
			InstanceIDs: []*string{
				to.Ptr("aaaaaaaaaaaaaaaaa")},
		},
			ResumeToken: "",
		})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/compute/resource-manager/Microsoft.Compute/stable/2021-11-01/examples/compute/VirtualMachineScaleSets_Redeploy_MaximumSet_Gen.json
func ExampleVirtualMachineScaleSetsClient_BeginRedeploy() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armcompute.NewVirtualMachineScaleSetsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginRedeploy(ctx,
		"<resource-group-name>",
		"<vm-scale-set-name>",
		&armcompute.VirtualMachineScaleSetsClientBeginRedeployOptions{VMInstanceIDs: &armcompute.VirtualMachineScaleSetVMInstanceIDs{
			InstanceIDs: []*string{
				to.Ptr("aaaaaaaaaaaaaaaaa")},
		},
			ResumeToken: "",
		})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/compute/resource-manager/Microsoft.Compute/stable/2021-11-01/examples/compute/VirtualMachineScaleSets_PerformMaintenance_MaximumSet_Gen.json
func ExampleVirtualMachineScaleSetsClient_BeginPerformMaintenance() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armcompute.NewVirtualMachineScaleSetsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginPerformMaintenance(ctx,
		"<resource-group-name>",
		"<vm-scale-set-name>",
		&armcompute.VirtualMachineScaleSetsClientBeginPerformMaintenanceOptions{VMInstanceIDs: &armcompute.VirtualMachineScaleSetVMInstanceIDs{
			InstanceIDs: []*string{
				to.Ptr("aaaaaaaaaaaaaaaaa")},
		},
			ResumeToken: "",
		})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/compute/resource-manager/Microsoft.Compute/stable/2021-11-01/examples/compute/VirtualMachineScaleSets_UpdateInstances_MaximumSet_Gen.json
func ExampleVirtualMachineScaleSetsClient_BeginUpdateInstances() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armcompute.NewVirtualMachineScaleSetsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginUpdateInstances(ctx,
		"<resource-group-name>",
		"<vm-scale-set-name>",
		armcompute.VirtualMachineScaleSetVMInstanceRequiredIDs{
			InstanceIDs: []*string{
				to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaaa")},
		},
		&armcompute.VirtualMachineScaleSetsClientBeginUpdateInstancesOptions{ResumeToken: ""})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/compute/resource-manager/Microsoft.Compute/stable/2021-11-01/examples/compute/VirtualMachineScaleSets_Reimage_MaximumSet_Gen.json
func ExampleVirtualMachineScaleSetsClient_BeginReimage() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armcompute.NewVirtualMachineScaleSetsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginReimage(ctx,
		"<resource-group-name>",
		"<vm-scale-set-name>",
		&armcompute.VirtualMachineScaleSetsClientBeginReimageOptions{VMScaleSetReimageInput: &armcompute.VirtualMachineScaleSetReimageParameters{
			TempDisk: to.Ptr(true),
			InstanceIDs: []*string{
				to.Ptr("aaaaaaaaaa")},
		},
			ResumeToken: "",
		})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/compute/resource-manager/Microsoft.Compute/stable/2021-11-01/examples/compute/VirtualMachineScaleSets_ReimageAll_MaximumSet_Gen.json
func ExampleVirtualMachineScaleSetsClient_BeginReimageAll() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armcompute.NewVirtualMachineScaleSetsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginReimageAll(ctx,
		"<resource-group-name>",
		"<vm-scale-set-name>",
		&armcompute.VirtualMachineScaleSetsClientBeginReimageAllOptions{VMInstanceIDs: &armcompute.VirtualMachineScaleSetVMInstanceIDs{
			InstanceIDs: []*string{
				to.Ptr("aaaaaaaaaaaaaaaaa")},
		},
			ResumeToken: "",
		})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/compute/resource-manager/Microsoft.Compute/stable/2021-11-01/examples/compute/VirtualMachineScaleSets_ForceRecoveryServiceFabricPlatformUpdateDomainWalk_MaximumSet_Gen.json
func ExampleVirtualMachineScaleSetsClient_ForceRecoveryServiceFabricPlatformUpdateDomainWalk() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armcompute.NewVirtualMachineScaleSetsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.ForceRecoveryServiceFabricPlatformUpdateDomainWalk(ctx,
		"<resource-group-name>",
		"<vm-scale-set-name>",
		30,
		&armcompute.VirtualMachineScaleSetsClientForceRecoveryServiceFabricPlatformUpdateDomainWalkOptions{Zone: nil,
			PlacementGroupID: nil,
		})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/compute/resource-manager/Microsoft.Compute/stable/2021-11-01/examples/compute/VirtualMachineScaleSets_ConvertToSinglePlacementGroup_MaximumSet_Gen.json
func ExampleVirtualMachineScaleSetsClient_ConvertToSinglePlacementGroup() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armcompute.NewVirtualMachineScaleSetsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = client.ConvertToSinglePlacementGroup(ctx,
		"<resource-group-name>",
		"<vm-scale-set-name>",
		armcompute.VMScaleSetConvertToSinglePlacementGroupInput{
			ActivePlacementGroupID: to.Ptr("<active-placement-group-id>"),
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/compute/resource-manager/Microsoft.Compute/stable/2021-11-01/examples/compute/VirtualMachineScaleSets_SetOrchestrationServiceState_MaximumSet_Gen.json
func ExampleVirtualMachineScaleSetsClient_BeginSetOrchestrationServiceState() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armcompute.NewVirtualMachineScaleSetsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginSetOrchestrationServiceState(ctx,
		"<resource-group-name>",
		"<vm-scale-set-name>",
		armcompute.OrchestrationServiceStateInput{
			Action:      to.Ptr(armcompute.OrchestrationServiceStateActionResume),
			ServiceName: to.Ptr(armcompute.OrchestrationServiceNamesAutomaticRepairs),
		},
		&armcompute.VirtualMachineScaleSetsClientBeginSetOrchestrationServiceStateOptions{ResumeToken: ""})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}
