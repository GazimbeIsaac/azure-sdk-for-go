//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmixedreality_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mixedreality/armmixedreality"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/mixedreality/resource-manager/Microsoft.MixedReality/preview/2021-03-01-preview/examples/object-anchors/GetBySubscription.json
func ExampleObjectAnchorsAccountsClient_NewListBySubscriptionPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armmixedreality.NewObjectAnchorsAccountsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := client.NewListBySubscriptionPager(nil)
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

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/mixedreality/resource-manager/Microsoft.MixedReality/preview/2021-03-01-preview/examples/object-anchors/GetByResourceGroup.json
func ExampleObjectAnchorsAccountsClient_NewListByResourceGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armmixedreality.NewObjectAnchorsAccountsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := client.NewListByResourceGroupPager("<resource-group-name>",
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

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/mixedreality/resource-manager/Microsoft.MixedReality/preview/2021-03-01-preview/examples/object-anchors/Delete.json
func ExampleObjectAnchorsAccountsClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armmixedreality.NewObjectAnchorsAccountsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = client.Delete(ctx,
		"<resource-group-name>",
		"<account-name>",
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/mixedreality/resource-manager/Microsoft.MixedReality/preview/2021-03-01-preview/examples/object-anchors/Get.json
func ExampleObjectAnchorsAccountsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armmixedreality.NewObjectAnchorsAccountsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Get(ctx,
		"<resource-group-name>",
		"<account-name>",
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/mixedreality/resource-manager/Microsoft.MixedReality/preview/2021-03-01-preview/examples/object-anchors/Patch.json
func ExampleObjectAnchorsAccountsClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armmixedreality.NewObjectAnchorsAccountsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Update(ctx,
		"<resource-group-name>",
		"<account-name>",
		armmixedreality.ObjectAnchorsAccount{
			Location: to.Ptr("<location>"),
			Tags: map[string]*string{
				"hero":    to.Ptr("romeo"),
				"heroine": to.Ptr("juliet"),
			},
			Identity: &armmixedreality.ObjectAnchorsAccountIdentity{
				Type: to.Ptr("<type>"),
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/mixedreality/resource-manager/Microsoft.MixedReality/preview/2021-03-01-preview/examples/object-anchors/Put.json
func ExampleObjectAnchorsAccountsClient_Create() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armmixedreality.NewObjectAnchorsAccountsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Create(ctx,
		"<resource-group-name>",
		"<account-name>",
		armmixedreality.ObjectAnchorsAccount{
			Location: to.Ptr("<location>"),
			Identity: &armmixedreality.ObjectAnchorsAccountIdentity{
				Type: to.Ptr("<type>"),
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/mixedreality/resource-manager/Microsoft.MixedReality/preview/2021-03-01-preview/examples/object-anchors/ListKeys.json
func ExampleObjectAnchorsAccountsClient_ListKeys() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armmixedreality.NewObjectAnchorsAccountsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.ListKeys(ctx,
		"<resource-group-name>",
		"<account-name>",
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/mixedreality/resource-manager/Microsoft.MixedReality/preview/2021-03-01-preview/examples/object-anchors/RegenerateKey.json
func ExampleObjectAnchorsAccountsClient_RegenerateKeys() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armmixedreality.NewObjectAnchorsAccountsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.RegenerateKeys(ctx,
		"<resource-group-name>",
		"<account-name>",
		armmixedreality.AccountKeyRegenerateRequest{
			Serial: to.Ptr(armmixedreality.SerialPrimary),
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
