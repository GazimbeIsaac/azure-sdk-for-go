//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcostmanagement_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/costmanagement/armcostmanagement"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/cost-management/resource-manager/Microsoft.CostManagement/stable/2021-10-01/examples/ExportsGetByBillingAccount.json
func ExampleExportsClient_List() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armcostmanagement.NewExportsClient(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.List(ctx,
		"<scope>",
		&armcostmanagement.ExportsClientListOptions{Expand: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/cost-management/resource-manager/Microsoft.CostManagement/stable/2021-10-01/examples/ExportGetByBillingAccount.json
func ExampleExportsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armcostmanagement.NewExportsClient(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Get(ctx,
		"<scope>",
		"<export-name>",
		&armcostmanagement.ExportsClientGetOptions{Expand: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/cost-management/resource-manager/Microsoft.CostManagement/stable/2021-10-01/examples/ExportCreateOrUpdateByBillingAccount.json
func ExampleExportsClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armcostmanagement.NewExportsClient(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.CreateOrUpdate(ctx,
		"<scope>",
		"<export-name>",
		armcostmanagement.Export{
			Properties: &armcostmanagement.ExportProperties{
				Format: to.Ptr(armcostmanagement.FormatTypeCSV),
				Definition: &armcostmanagement.ExportDefinition{
					Type: to.Ptr(armcostmanagement.ExportTypeActualCost),
					DataSet: &armcostmanagement.ExportDataset{
						Configuration: &armcostmanagement.ExportDatasetConfiguration{
							Columns: []*string{
								to.Ptr("Date"),
								to.Ptr("MeterId"),
								to.Ptr("ResourceId"),
								to.Ptr("ResourceLocation"),
								to.Ptr("Quantity")},
						},
						Granularity: to.Ptr(armcostmanagement.GranularityTypeDaily),
					},
					Timeframe: to.Ptr(armcostmanagement.TimeframeTypeMonthToDate),
				},
				DeliveryInfo: &armcostmanagement.ExportDeliveryInfo{
					Destination: &armcostmanagement.ExportDeliveryDestination{
						Container:      to.Ptr("<container>"),
						ResourceID:     to.Ptr("<resource-id>"),
						RootFolderPath: to.Ptr("<root-folder-path>"),
					},
				},
				Schedule: &armcostmanagement.ExportSchedule{
					Recurrence: to.Ptr(armcostmanagement.RecurrenceTypeWeekly),
					RecurrencePeriod: &armcostmanagement.ExportRecurrencePeriod{
						From: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-06-01T00:00:00Z"); return t }()),
						To:   to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-10-31T00:00:00Z"); return t }()),
					},
					Status: to.Ptr(armcostmanagement.StatusTypeActive),
				},
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/cost-management/resource-manager/Microsoft.CostManagement/stable/2021-10-01/examples/ExportDeleteByBillingAccount.json
func ExampleExportsClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armcostmanagement.NewExportsClient(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = client.Delete(ctx,
		"<scope>",
		"<export-name>",
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/cost-management/resource-manager/Microsoft.CostManagement/stable/2021-10-01/examples/ExportRunByBillingAccount.json
func ExampleExportsClient_Execute() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armcostmanagement.NewExportsClient(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = client.Execute(ctx,
		"<scope>",
		"<export-name>",
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/cost-management/resource-manager/Microsoft.CostManagement/stable/2021-10-01/examples/ExportRunHistoryGetByBillingAccount.json
func ExampleExportsClient_GetExecutionHistory() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armcostmanagement.NewExportsClient(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.GetExecutionHistory(ctx,
		"<scope>",
		"<export-name>",
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
