//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v5"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/f4c6c8697c59f966db0d1e36b62df3af3bca9065/specification/network/resource-manager/Microsoft.Network/stable/2023-11-01/examples/InboundNatRuleList.json
func ExampleInboundNatRulesClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewInboundNatRulesClient().NewListPager("testrg", "lb1", nil)
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
		// page.InboundNatRuleListResult = armnetwork.InboundNatRuleListResult{
		// 	Value: []*armnetwork.InboundNatRule{
		// 		{
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/testrg/providers/Microsoft.Network/loadBalancers/lb1/inboundNatRules/natRule1.1"),
		// 			Name: to.Ptr("natRule1.1"),
		// 			Properties: &armnetwork.InboundNatRulePropertiesFormat{
		// 				BackendIPConfiguration: &armnetwork.InterfaceIPConfiguration{
		// 					ID: to.Ptr("/subscriptions/subid/resourceGroups/testrg/providers/Microsoft.Compute/virtualMachineScaleSets/vmss1/virtualMachines/1/networkInterfaces/nic1/ipConfigurations/ip1"),
		// 				},
		// 				BackendPort: to.Ptr[int32](3389),
		// 				EnableFloatingIP: to.Ptr(false),
		// 				EnableTCPReset: to.Ptr(true),
		// 				FrontendIPConfiguration: &armnetwork.SubResource{
		// 					ID: to.Ptr("/subscriptions/subid/resourceGroups/testrg/providers/Microsoft.Network/loadBalancers/lb1/frontendIPConfigurations/ip1"),
		// 				},
		// 				FrontendPort: to.Ptr[int32](3390),
		// 				IdleTimeoutInMinutes: to.Ptr[int32](4),
		// 				ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 				Protocol: to.Ptr(armnetwork.TransportProtocolTCP),
		// 			},
		// 		},
		// 		{
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/testrg/providers/Microsoft.Network/loadBalancers/lb1/inboundNatRules/natRule1.3"),
		// 			Name: to.Ptr("natRule1.3"),
		// 			Properties: &armnetwork.InboundNatRulePropertiesFormat{
		// 				BackendIPConfiguration: &armnetwork.InterfaceIPConfiguration{
		// 					ID: to.Ptr("/subscriptions/subid/resourceGroups/testrg/providers/Microsoft.Compute/virtualMachineScaleSets/vmss1/virtualMachines/3/networkInterfaces/nic1/ipConfigurations/ip1"),
		// 				},
		// 				BackendPort: to.Ptr[int32](3389),
		// 				EnableFloatingIP: to.Ptr(false),
		// 				EnableTCPReset: to.Ptr(true),
		// 				FrontendIPConfiguration: &armnetwork.SubResource{
		// 					ID: to.Ptr("/subscriptions/subid/resourceGroups/testrg/providers/Microsoft.Network/loadBalancers/lb1/frontendIPConfigurations/ip1"),
		// 				},
		// 				FrontendPort: to.Ptr[int32](3392),
		// 				IdleTimeoutInMinutes: to.Ptr[int32](4),
		// 				ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 				Protocol: to.Ptr(armnetwork.TransportProtocolTCP),
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/f4c6c8697c59f966db0d1e36b62df3af3bca9065/specification/network/resource-manager/Microsoft.Network/stable/2023-11-01/examples/InboundNatRuleDelete.json
func ExampleInboundNatRulesClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewInboundNatRulesClient().BeginDelete(ctx, "testrg", "lb1", "natRule1.1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/f4c6c8697c59f966db0d1e36b62df3af3bca9065/specification/network/resource-manager/Microsoft.Network/stable/2023-11-01/examples/InboundNatRuleGet.json
func ExampleInboundNatRulesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewInboundNatRulesClient().Get(ctx, "testrg", "lb1", "natRule1.1", &armnetwork.InboundNatRulesClientGetOptions{Expand: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.InboundNatRule = armnetwork.InboundNatRule{
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/testrg/providers/Microsoft.Network/loadBalancers/lb1/inboundNatRules/natRule1.1"),
	// 	Name: to.Ptr("natRule1.1"),
	// 	Properties: &armnetwork.InboundNatRulePropertiesFormat{
	// 		BackendIPConfiguration: &armnetwork.InterfaceIPConfiguration{
	// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/testrg/providers/Microsoft.Compute/virtualMachineScaleSets/vmss1/virtualMachines/1/networkInterfaces/nic1/ipConfigurations/ip1"),
	// 		},
	// 		BackendPort: to.Ptr[int32](3389),
	// 		EnableFloatingIP: to.Ptr(false),
	// 		EnableTCPReset: to.Ptr(true),
	// 		FrontendIPConfiguration: &armnetwork.SubResource{
	// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/testrg/providers/Microsoft.Network/loadBalancers/lb1/frontendIPConfigurations/ip1"),
	// 		},
	// 		FrontendPort: to.Ptr[int32](3390),
	// 		IdleTimeoutInMinutes: to.Ptr[int32](4),
	// 		ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 		Protocol: to.Ptr(armnetwork.TransportProtocolTCP),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/f4c6c8697c59f966db0d1e36b62df3af3bca9065/specification/network/resource-manager/Microsoft.Network/stable/2023-11-01/examples/InboundNatRuleCreate.json
func ExampleInboundNatRulesClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewInboundNatRulesClient().BeginCreateOrUpdate(ctx, "testrg", "lb1", "natRule1.1", armnetwork.InboundNatRule{
		Properties: &armnetwork.InboundNatRulePropertiesFormat{
			BackendPort:      to.Ptr[int32](3389),
			EnableFloatingIP: to.Ptr(false),
			EnableTCPReset:   to.Ptr(false),
			FrontendIPConfiguration: &armnetwork.SubResource{
				ID: to.Ptr("/subscriptions/subid/resourceGroups/testrg/providers/Microsoft.Network/loadBalancers/lb1/frontendIPConfigurations/ip1"),
			},
			FrontendPort:         to.Ptr[int32](3390),
			IdleTimeoutInMinutes: to.Ptr[int32](4),
			Protocol:             to.Ptr(armnetwork.TransportProtocolTCP),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.InboundNatRule = armnetwork.InboundNatRule{
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/testrg/providers/Microsoft.Network/loadBalancers/lb1/inboundNatRules/natRule1.1"),
	// 	Name: to.Ptr("natRule1.1"),
	// 	Properties: &armnetwork.InboundNatRulePropertiesFormat{
	// 		BackendIPConfiguration: &armnetwork.InterfaceIPConfiguration{
	// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/testrg/providers/Microsoft.Compute/virtualMachineScaleSets/vmss1/virtualMachines/1/networkInterfaces/nic1/ipConfigurations/ip1"),
	// 		},
	// 		BackendPort: to.Ptr[int32](3389),
	// 		EnableFloatingIP: to.Ptr(false),
	// 		EnableTCPReset: to.Ptr(false),
	// 		FrontendIPConfiguration: &armnetwork.SubResource{
	// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/testrg/providers/Microsoft.Network/loadBalancers/lb1/frontendIPConfigurations/ip1"),
	// 		},
	// 		FrontendPort: to.Ptr[int32](3390),
	// 		IdleTimeoutInMinutes: to.Ptr[int32](4),
	// 		ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 		Protocol: to.Ptr(armnetwork.TransportProtocolTCP),
	// 	},
	// }
}
