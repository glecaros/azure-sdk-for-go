//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armappservice_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appservice/armappservice/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/de1f3772629b6f4d3ac01548a5f6d719bfb97c9e/specification/web/resource-manager/Microsoft.Web/stable/2023-12-01/examples/WorkflowRuns_List.json
func ExampleWorkflowRunsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappservice.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewWorkflowRunsClient().NewListPager("test-resource-group", "test-name", "test-workflow", &armappservice.WorkflowRunsClientListOptions{Top: nil,
		Filter: nil,
	})
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.WorkflowRunListResult = armappservice.WorkflowRunListResult{
		// 	Value: []*armappservice.WorkflowRun{
		// 		{
		// 			ID: to.Ptr("workflows/test-workflow/runs/08586676746934337772206998657CU22"),
		// 			Name: to.Ptr("08586676746934337772206998657CU22"),
		// 			Type: to.Ptr("workflows/runs"),
		// 			Properties: &armappservice.WorkflowRunProperties{
		// 				Correlation: &armappservice.Correlation{
		// 					ClientTrackingID: to.Ptr("08586676746934337772206998657CU22"),
		// 				},
		// 				EndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-08-10T20:16:32.577Z"); return t}()),
		// 				Outputs: map[string]*armappservice.WorkflowOutputParameter{
		// 				},
		// 				StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-08-10T20:16:32.044Z"); return t}()),
		// 				Status: to.Ptr(armappservice.WorkflowStatusSucceeded),
		// 				Trigger: &armappservice.WorkflowRunTrigger{
		// 					Name: to.Ptr("Recurrence"),
		// 					Code: to.Ptr("OK"),
		// 					Correlation: &armappservice.Correlation{
		// 						ClientTrackingID: to.Ptr("08586676746934337772206998657CU22"),
		// 					},
		// 					EndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-08-10T20:16:32.038Z"); return t}()),
		// 					ScheduledTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-08-10T20:16:31.634Z"); return t}()),
		// 					StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-08-10T20:16:32.038Z"); return t}()),
		// 					Status: to.Ptr(armappservice.WorkflowStatusSucceeded),
		// 				},
		// 				WaitEndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-08-10T20:16:32.044Z"); return t}()),
		// 				Workflow: &armappservice.ResourceReference{
		// 					Name: to.Ptr("08586676754160363885"),
		// 					Type: to.Ptr("workflows/versions"),
		// 					ID: to.Ptr("workflows/test-workflow/versions/08586676754160363885"),
		// 				},
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/de1f3772629b6f4d3ac01548a5f6d719bfb97c9e/specification/web/resource-manager/Microsoft.Web/stable/2023-12-01/examples/WorkflowRuns_Get.json
func ExampleWorkflowRunsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappservice.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewWorkflowRunsClient().Get(ctx, "test-resource-group", "test-name", "test-workflow", "08586676746934337772206998657CU22", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.WorkflowRun = armappservice.WorkflowRun{
	// 	ID: to.Ptr("/workflows/test-workflow/runs/08586676746934337772206998657CU22"),
	// 	Name: to.Ptr("08586676746934337772206998657CU22"),
	// 	Type: to.Ptr("workflows/runs"),
	// 	Properties: &armappservice.WorkflowRunProperties{
	// 		Correlation: &armappservice.Correlation{
	// 			ClientTrackingID: to.Ptr("08586676746934337772206998657CU22"),
	// 		},
	// 		EndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-08-10T20:16:32.577Z"); return t}()),
	// 		Outputs: map[string]*armappservice.WorkflowOutputParameter{
	// 		},
	// 		StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-08-10T20:16:32.044Z"); return t}()),
	// 		Status: to.Ptr(armappservice.WorkflowStatusSucceeded),
	// 		Trigger: &armappservice.WorkflowRunTrigger{
	// 			Name: to.Ptr("Recurrence"),
	// 			Code: to.Ptr("OK"),
	// 			Correlation: &armappservice.Correlation{
	// 				ClientTrackingID: to.Ptr("08586676746934337772206998657CU22"),
	// 			},
	// 			EndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-08-10T20:16:32.038Z"); return t}()),
	// 			ScheduledTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-08-10T20:16:31.634Z"); return t}()),
	// 			StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-08-10T20:16:32.038Z"); return t}()),
	// 			Status: to.Ptr(armappservice.WorkflowStatusSucceeded),
	// 		},
	// 		WaitEndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-08-10T20:16:32.044Z"); return t}()),
	// 		Workflow: &armappservice.ResourceReference{
	// 			Name: to.Ptr("08586676754160363885"),
	// 			Type: to.Ptr("workflows/versions"),
	// 			ID: to.Ptr("/workflows/test-workflow/versions/08586676754160363885"),
	// 		},
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/de1f3772629b6f4d3ac01548a5f6d719bfb97c9e/specification/web/resource-manager/Microsoft.Web/stable/2023-12-01/examples/WorkflowRuns_Cancel.json
func ExampleWorkflowRunsClient_Cancel() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappservice.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewWorkflowRunsClient().Cancel(ctx, "test-resource-group", "test-name", "test-workflow", "08586676746934337772206998657CU22", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}
