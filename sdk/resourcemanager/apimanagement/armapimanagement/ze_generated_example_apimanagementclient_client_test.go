//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armapimanagement_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/apimanagement/armapimanagement"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementPerformConnectivityCheckHttpConnect.json
func ExampleClient_BeginPerformConnectivityCheckAsync() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armapimanagement.NewClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginPerformConnectivityCheckAsync(ctx,
		"<resource-group-name>",
		"<service-name>",
		armapimanagement.ConnectivityCheckRequest{
			Destination: &armapimanagement.ConnectivityCheckRequestDestination{
				Address: to.Ptr("<address>"),
				Port:    to.Ptr[int64](3306),
			},
			ProtocolConfiguration: &armapimanagement.ConnectivityCheckRequestProtocolConfiguration{
				HTTPConfiguration: &armapimanagement.ConnectivityCheckRequestProtocolConfigurationHTTPConfiguration{
					Method: to.Ptr(armapimanagement.MethodGET),
					Headers: []*armapimanagement.HTTPHeader{
						{
							Name:  to.Ptr("<name>"),
							Value: to.Ptr("<value>"),
						}},
					ValidStatusCodes: []*int64{
						to.Ptr[int64](200),
						to.Ptr[int64](204)},
				},
			},
			Source: &armapimanagement.ConnectivityCheckRequestSource{
				Region: to.Ptr("<region>"),
			},
			Protocol: to.Ptr(armapimanagement.ConnectivityCheckProtocolHTTPS),
		},
		&armapimanagement.ClientBeginPerformConnectivityCheckAsyncOptions{ResumeToken: ""})
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
