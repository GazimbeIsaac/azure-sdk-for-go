//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armnetwork_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/network/resource-manager/Microsoft.Network/stable/2021-05-01/examples/LoadBalancerDelete.json
func ExampleLoadBalancersClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armnetwork.NewLoadBalancersClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginDelete(ctx,
		"<resource-group-name>",
		"<load-balancer-name>",
		&armnetwork.LoadBalancersClientBeginDeleteOptions{ResumeToken: ""})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/network/resource-manager/Microsoft.Network/stable/2021-05-01/examples/LoadBalancerGet.json
func ExampleLoadBalancersClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armnetwork.NewLoadBalancersClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Get(ctx,
		"<resource-group-name>",
		"<load-balancer-name>",
		&armnetwork.LoadBalancersClientGetOptions{Expand: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/network/resource-manager/Microsoft.Network/stable/2021-05-01/examples/LoadBalancerCreate.json
func ExampleLoadBalancersClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armnetwork.NewLoadBalancersClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<resource-group-name>",
		"<load-balancer-name>",
		armnetwork.LoadBalancer{
			Location: to.Ptr("<location>"),
			Properties: &armnetwork.LoadBalancerPropertiesFormat{
				BackendAddressPools: []*armnetwork.BackendAddressPool{
					{
						Name:       to.Ptr("<name>"),
						Properties: &armnetwork.BackendAddressPoolPropertiesFormat{},
					}},
				FrontendIPConfigurations: []*armnetwork.FrontendIPConfiguration{
					{
						Name: to.Ptr("<name>"),
						Properties: &armnetwork.FrontendIPConfigurationPropertiesFormat{
							Subnet: &armnetwork.Subnet{
								ID: to.Ptr("<id>"),
							},
						},
					}},
				InboundNatPools: []*armnetwork.InboundNatPool{},
				InboundNatRules: []*armnetwork.InboundNatRule{
					{
						Name: to.Ptr("<name>"),
						Properties: &armnetwork.InboundNatRulePropertiesFormat{
							BackendPort:      to.Ptr[int32](3389),
							EnableFloatingIP: to.Ptr(true),
							EnableTCPReset:   to.Ptr(false),
							FrontendIPConfiguration: &armnetwork.SubResource{
								ID: to.Ptr("<id>"),
							},
							FrontendPort:         to.Ptr[int32](3389),
							IdleTimeoutInMinutes: to.Ptr[int32](15),
							Protocol:             to.Ptr(armnetwork.TransportProtocolTCP),
						},
					}},
				LoadBalancingRules: []*armnetwork.LoadBalancingRule{
					{
						Name: to.Ptr("<name>"),
						Properties: &armnetwork.LoadBalancingRulePropertiesFormat{
							BackendAddressPool: &armnetwork.SubResource{
								ID: to.Ptr("<id>"),
							},
							BackendPort:      to.Ptr[int32](80),
							EnableFloatingIP: to.Ptr(true),
							EnableTCPReset:   to.Ptr(false),
							FrontendIPConfiguration: &armnetwork.SubResource{
								ID: to.Ptr("<id>"),
							},
							FrontendPort:         to.Ptr[int32](80),
							IdleTimeoutInMinutes: to.Ptr[int32](15),
							LoadDistribution:     to.Ptr(armnetwork.LoadDistributionDefault),
							Probe: &armnetwork.SubResource{
								ID: to.Ptr("<id>"),
							},
							Protocol: to.Ptr(armnetwork.TransportProtocolTCP),
						},
					}},
				Probes: []*armnetwork.Probe{
					{
						Name: to.Ptr("<name>"),
						Properties: &armnetwork.ProbePropertiesFormat{
							IntervalInSeconds: to.Ptr[int32](15),
							NumberOfProbes:    to.Ptr[int32](2),
							Port:              to.Ptr[int32](80),
							RequestPath:       to.Ptr("<request-path>"),
							Protocol:          to.Ptr(armnetwork.ProbeProtocolHTTP),
						},
					}},
			},
		},
		&armnetwork.LoadBalancersClientBeginCreateOrUpdateOptions{ResumeToken: ""})
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

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/network/resource-manager/Microsoft.Network/stable/2021-05-01/examples/LoadBalancerUpdateTags.json
func ExampleLoadBalancersClient_UpdateTags() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armnetwork.NewLoadBalancersClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.UpdateTags(ctx,
		"<resource-group-name>",
		"<load-balancer-name>",
		armnetwork.TagsObject{
			Tags: map[string]*string{
				"tag1": to.Ptr("value1"),
				"tag2": to.Ptr("value2"),
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/network/resource-manager/Microsoft.Network/stable/2021-05-01/examples/LoadBalancerListAll.json
func ExampleLoadBalancersClient_NewListAllPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armnetwork.NewLoadBalancersClient("<subscription-id>", cred, nil)
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

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/network/resource-manager/Microsoft.Network/stable/2021-05-01/examples/LoadBalancerList.json
func ExampleLoadBalancersClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armnetwork.NewLoadBalancersClient("<subscription-id>", cred, nil)
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

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/network/resource-manager/Microsoft.Network/stable/2021-05-01/examples/QueryInboundNatRulePortMapping.json
func ExampleLoadBalancersClient_BeginListInboundNatRulePortMappings() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armnetwork.NewLoadBalancersClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginListInboundNatRulePortMappings(ctx,
		"<group-name>",
		"<load-balancer-name>",
		"<backend-pool-name>",
		armnetwork.QueryInboundNatRulePortMappingRequest{
			IPAddress: to.Ptr("<ipaddress>"),
		},
		&armnetwork.LoadBalancersClientBeginListInboundNatRulePortMappingsOptions{ResumeToken: ""})
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
