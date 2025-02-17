//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armblueprint_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/blueprint/armblueprint"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/blueprint/resource-manager/Microsoft.Blueprint/preview/2018-11-01-preview/examples/managementGroupBPDef/ARMTemplateArtifact_Create.json
func ExampleArtifactsClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armblueprint.NewArtifactsClient(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = client.CreateOrUpdate(ctx,
		"<resource-scope>",
		"<blueprint-name>",
		"<artifact-name>",
		&armblueprint.TemplateArtifact{
			Kind: to.Ptr(armblueprint.ArtifactKindTemplate),
			Properties: &armblueprint.TemplateArtifactProperties{
				Parameters: map[string]*armblueprint.ParameterValue{
					"storageAccountType": {
						Value: "[parameters('storageAccountType')]",
					},
				},
				ResourceGroup: to.Ptr("<resource-group>"),
				Template: map[string]interface{}{
					"contentVersion": "1.0.0.0",
					"outputs": map[string]interface{}{
						"storageAccountName": map[string]interface{}{
							"type":  "string",
							"value": "[variables('storageAccountName')]",
						},
					},
					"parameters": map[string]interface{}{
						"storageAccountType": map[string]interface{}{
							"type": "string",
							"allowedValues": []interface{}{
								"Standard_LRS",
								"Standard_GRS",
								"Standard_ZRS",
								"Premium_LRS",
							},
							"defaultValue": "Standard_LRS",
							"metadata": map[string]interface{}{
								"description": "Storage Account type",
							},
						},
					},
					"resources": []interface{}{
						map[string]interface{}{
							"name":       "[variables('storageAccountName')]",
							"type":       "Microsoft.Storage/storageAccounts",
							"apiVersion": "2016-01-01",
							"kind":       "Storage",
							"location":   "[resourceGroup().location]",
							"properties": map[string]interface{}{},
							"sku": map[string]interface{}{
								"name": "[parameters('storageAccountType')]",
							},
						},
					},
					"variables": map[string]interface{}{
						"storageAccountName": "[concat(uniquestring(resourceGroup().id), 'standardsa')]",
					},
				},
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/blueprint/resource-manager/Microsoft.Blueprint/preview/2018-11-01-preview/examples/managementGroupBPDef/ARMTemplateArtifact_Get.json
func ExampleArtifactsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armblueprint.NewArtifactsClient(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Get(ctx,
		"<resource-scope>",
		"<blueprint-name>",
		"<artifact-name>",
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/blueprint/resource-manager/Microsoft.Blueprint/preview/2018-11-01-preview/examples/managementGroupBPDef/ARMTemplateArtifact_Delete.json
func ExampleArtifactsClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armblueprint.NewArtifactsClient(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Delete(ctx,
		"<resource-scope>",
		"<blueprint-name>",
		"<artifact-name>",
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/blueprint/resource-manager/Microsoft.Blueprint/preview/2018-11-01-preview/examples/managementGroupBPDef/Artifact_List.json
func ExampleArtifactsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armblueprint.NewArtifactsClient(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := client.NewListPager("<resource-scope>",
		"<blueprint-name>",
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
