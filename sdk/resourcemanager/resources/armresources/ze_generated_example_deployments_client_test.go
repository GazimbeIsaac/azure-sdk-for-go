//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armresources_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armresources"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/resources/resource-manager/Microsoft.Resources/stable/2021-04-01/examples/PutDeploymentAtScope.json
func ExampleDeploymentsClient_BeginCreateOrUpdateAtScope() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armresources.NewDeploymentsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreateOrUpdateAtScope(ctx,
		"<scope>",
		"<deployment-name>",
		armresources.Deployment{
			Location: to.Ptr("<location>"),
			Properties: &armresources.DeploymentProperties{
				Mode:       to.Ptr(armresources.DeploymentModeIncremental),
				Parameters: map[string]interface{}{},
				TemplateLink: &armresources.TemplateLink{
					URI: to.Ptr("<uri>"),
				},
			},
			Tags: map[string]*string{
				"tagKey1": to.Ptr("tag-value-1"),
				"tagKey2": to.Ptr("tag-value-2"),
			},
		},
		&armresources.DeploymentsClientBeginCreateOrUpdateAtScopeOptions{ResumeToken: ""})
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

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/resources/resource-manager/Microsoft.Resources/stable/2021-04-01/examples/PutDeploymentAtTenant.json
func ExampleDeploymentsClient_BeginCreateOrUpdateAtTenantScope() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armresources.NewDeploymentsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreateOrUpdateAtTenantScope(ctx,
		"<deployment-name>",
		armresources.ScopedDeployment{
			Location: to.Ptr("<location>"),
			Properties: &armresources.DeploymentProperties{
				Mode:       to.Ptr(armresources.DeploymentModeIncremental),
				Parameters: map[string]interface{}{},
				TemplateLink: &armresources.TemplateLink{
					URI: to.Ptr("<uri>"),
				},
			},
			Tags: map[string]*string{
				"tagKey1": to.Ptr("tag-value-1"),
				"tagKey2": to.Ptr("tag-value-2"),
			},
		},
		&armresources.DeploymentsClientBeginCreateOrUpdateAtTenantScopeOptions{ResumeToken: ""})
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

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/resources/resource-manager/Microsoft.Resources/stable/2021-04-01/examples/PostDeploymentWhatIfOnTenant.json
func ExampleDeploymentsClient_BeginWhatIfAtTenantScope() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armresources.NewDeploymentsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginWhatIfAtTenantScope(ctx,
		"<deployment-name>",
		armresources.ScopedDeploymentWhatIf{
			Location: to.Ptr("<location>"),
			Properties: &armresources.DeploymentWhatIfProperties{
				Mode:         to.Ptr(armresources.DeploymentModeIncremental),
				Parameters:   map[string]interface{}{},
				TemplateLink: &armresources.TemplateLink{},
			},
		},
		&armresources.DeploymentsClientBeginWhatIfAtTenantScopeOptions{ResumeToken: ""})
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

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/resources/resource-manager/Microsoft.Resources/stable/2021-04-01/examples/PutDeploymentAtManagementGroup.json
func ExampleDeploymentsClient_BeginCreateOrUpdateAtManagementGroupScope() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armresources.NewDeploymentsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreateOrUpdateAtManagementGroupScope(ctx,
		"<group-id>",
		"<deployment-name>",
		armresources.ScopedDeployment{
			Location: to.Ptr("<location>"),
			Properties: &armresources.DeploymentProperties{
				Mode:       to.Ptr(armresources.DeploymentModeIncremental),
				Parameters: map[string]interface{}{},
				TemplateLink: &armresources.TemplateLink{
					URI: to.Ptr("<uri>"),
				},
			},
		},
		&armresources.DeploymentsClientBeginCreateOrUpdateAtManagementGroupScopeOptions{ResumeToken: ""})
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

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/resources/resource-manager/Microsoft.Resources/stable/2021-04-01/examples/PostDeploymentWhatIfOnManagementGroup.json
func ExampleDeploymentsClient_BeginWhatIfAtManagementGroupScope() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armresources.NewDeploymentsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginWhatIfAtManagementGroupScope(ctx,
		"<group-id>",
		"<deployment-name>",
		armresources.ScopedDeploymentWhatIf{
			Location: to.Ptr("<location>"),
			Properties: &armresources.DeploymentWhatIfProperties{
				Mode:         to.Ptr(armresources.DeploymentModeIncremental),
				Parameters:   map[string]interface{}{},
				TemplateLink: &armresources.TemplateLink{},
			},
		},
		&armresources.DeploymentsClientBeginWhatIfAtManagementGroupScopeOptions{ResumeToken: ""})
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

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/resources/resource-manager/Microsoft.Resources/stable/2021-04-01/examples/PutDeploymentSubscriptionTemplateSpecsWithId.json
func ExampleDeploymentsClient_BeginCreateOrUpdateAtSubscriptionScope() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armresources.NewDeploymentsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreateOrUpdateAtSubscriptionScope(ctx,
		"<deployment-name>",
		armresources.Deployment{
			Location: to.Ptr("<location>"),
			Properties: &armresources.DeploymentProperties{
				Mode:       to.Ptr(armresources.DeploymentModeIncremental),
				Parameters: map[string]interface{}{},
				TemplateLink: &armresources.TemplateLink{
					ID: to.Ptr("<id>"),
				},
			},
		},
		&armresources.DeploymentsClientBeginCreateOrUpdateAtSubscriptionScopeOptions{ResumeToken: ""})
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

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/resources/resource-manager/Microsoft.Resources/stable/2021-04-01/examples/PostDeploymentWhatIfOnSubscription.json
func ExampleDeploymentsClient_BeginWhatIfAtSubscriptionScope() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armresources.NewDeploymentsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginWhatIfAtSubscriptionScope(ctx,
		"<deployment-name>",
		armresources.DeploymentWhatIf{
			Location: to.Ptr("<location>"),
			Properties: &armresources.DeploymentWhatIfProperties{
				Mode:         to.Ptr(armresources.DeploymentModeIncremental),
				Parameters:   map[string]interface{}{},
				TemplateLink: &armresources.TemplateLink{},
			},
		},
		&armresources.DeploymentsClientBeginWhatIfAtSubscriptionScopeOptions{ResumeToken: ""})
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

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/resources/resource-manager/Microsoft.Resources/stable/2021-04-01/examples/PutDeploymentResourceGroup.json
func ExampleDeploymentsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armresources.NewDeploymentsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<resource-group-name>",
		"<deployment-name>",
		armresources.Deployment{
			Properties: &armresources.DeploymentProperties{
				Mode:       to.Ptr(armresources.DeploymentModeIncremental),
				Parameters: map[string]interface{}{},
				TemplateLink: &armresources.TemplateLink{
					QueryString: to.Ptr("<query-string>"),
					URI:         to.Ptr("<uri>"),
				},
			},
		},
		&armresources.DeploymentsClientBeginCreateOrUpdateOptions{ResumeToken: ""})
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

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/resources/resource-manager/Microsoft.Resources/stable/2021-04-01/examples/PostDeploymentWhatIfOnResourceGroup.json
func ExampleDeploymentsClient_BeginWhatIf() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armresources.NewDeploymentsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginWhatIf(ctx,
		"<resource-group-name>",
		"<deployment-name>",
		armresources.DeploymentWhatIf{
			Properties: &armresources.DeploymentWhatIfProperties{
				Mode:         to.Ptr(armresources.DeploymentModeIncremental),
				Parameters:   map[string]interface{}{},
				TemplateLink: &armresources.TemplateLink{},
			},
		},
		&armresources.DeploymentsClientBeginWhatIfOptions{ResumeToken: ""})
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

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/resources/resource-manager/Microsoft.Resources/stable/2021-04-01/examples/CalculateTemplateHash.json
func ExampleDeploymentsClient_CalculateTemplateHash() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armresources.NewDeploymentsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.CalculateTemplateHash(ctx,
		map[string]interface{}{
			"$schema":        "http://schemas.management.azure.com/deploymentTemplate?api-version=2014-04-01-preview",
			"contentVersion": "1.0.0.0",
			"outputs": map[string]interface{}{
				"string": map[string]interface{}{
					"type":  "string",
					"value": "myvalue",
				},
			},
			"parameters": map[string]interface{}{
				"string": map[string]interface{}{
					"type": "string",
				},
			},
			"resources": []interface{}{},
			"variables": map[string]interface{}{
				"array": []interface{}{
					float64(1),
					float64(2),
					float64(3),
					float64(4),
				},
				"bool": true,
				"int":  float64(42),
				"object": map[string]interface{}{
					"object": map[string]interface{}{
						"location": "West US",
						"vmSize":   "Large",
					},
				},
				"string": "string",
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
