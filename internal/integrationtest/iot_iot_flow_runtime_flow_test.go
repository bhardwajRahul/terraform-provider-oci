// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"

	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	IotIotFlowRuntimeFlowResourceConfig = acctest.GenerateResourceFromRepresentationMap("oci_iot_iot_flow_runtime_flow", "test_iot_flow_runtime_flow", acctest.Optional, acctest.Update, IotIotFlowRuntimeFlowRepresentation)

	flow = `${jsonencode({
            "flows" : [
   				{
                  "id": "simple-flow-tab",
                  "type": "tab",
                  "label": "Simple Flow"
                }
			]
          })}`

	IotIotFlowRuntimeFlowSingularDataSourceRepresentation = map[string]interface{}{
		"iot_flow_runtime_id": acctest.Representation{RepType: acctest.Required, Create: `${var.iot_flow_runtime_id}`},
	}

	IotIotFlowRuntimeFlowRepresentation = map[string]interface{}{
		"flows_document":      acctest.Representation{RepType: acctest.Required, Create: flow},
		"iot_flow_runtime_id": acctest.Representation{RepType: acctest.Required, Create: `${var.iot_flow_runtime_id}`},
	}
)

// issue-routing-tag: iot/default
func TestIotIotFlowRuntimeFlowResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestIotIotFlowRuntimeFlowResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	iotFlowruntimeId := utils.GetEnvSettingWithBlankDefault("iot_flow_runtime_id")
	iotFlowruntimeIdVariableStr := fmt.Sprintf("variable \"iot_flow_runtime_id\" { default = \"%s\" }\n", iotFlowruntimeId)
	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_id")

	resourceName := "oci_iot_iot_flow_runtime_flow.test_iot_flow_runtime_flow"

	singularDatasourceName := "data.oci_iot_iot_flow_runtime_flow.test_iot_flow_runtime_flow"

	var resId, resId2 string
	// Save TF content to Create resource with only required properties. This has to be exactly the same as the config part in the create step in the test.
	acctest.SaveConfigContent(config+iotFlowruntimeIdVariableStr+
		acctest.GenerateResourceFromRepresentationMap("oci_iot_iot_flow_runtime_flow", "test_iot_flow_runtime_flow", acctest.Required, acctest.Create, IotIotFlowRuntimeFlowRepresentation), "iot", "iotFlowRuntimeFlow", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify Create
		{
			Config: config + iotFlowruntimeIdVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_iot_iot_flow_runtime_flow", "test_iot_flow_runtime_flow", acctest.Required, acctest.Create, IotIotFlowRuntimeFlowRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "flows_document"),
				resource.TestCheckResourceAttrSet(resourceName, "iot_flow_runtime_id"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// verify updates to updatable parameters
		{
			Config: config + iotFlowruntimeIdVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_iot_iot_flow_runtime_flow", "test_iot_flow_runtime_flow", acctest.Optional, acctest.Update, IotIotFlowRuntimeFlowRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "flows_document"),
				resource.TestCheckResourceAttrSet(resourceName, "iot_flow_runtime_id"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					if isEnableExportCompartment, _ := strconv.ParseBool(utils.GetEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
						if errExport := resourcediscovery.TestExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
							return errExport
						}
					}
					return err
				},
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_iot_iot_flow_runtime_flow", "test_iot_flow_runtime_flow", acctest.Required, acctest.Create, IotIotFlowRuntimeFlowSingularDataSourceRepresentation) +
				iotFlowruntimeIdVariableStr + IotIotFlowRuntimeFlowResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "iot_flow_runtime_id"),
			),
		},
		// verify resource import
		{
			Config:            config + iotFlowruntimeIdVariableStr + IotIotFlowRuntimeFlowResourceConfig,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateVerifyIgnore: []string{
				"flows_document",
			},
			ResourceName: resourceName,
		},
	})
}
