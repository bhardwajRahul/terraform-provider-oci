// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/oracle/terraform-provider-oci/internal/utils"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
)

var (
	IotIotFlowRuntimeActivateRepresentation = map[string]interface{}{
		"iot_flow_runtime_id": acctest.Representation{RepType: acctest.Required, Create: `${var.iot_flow_runtime_id}`},
	}
)

// issue-routing-tag: iot/default
func TestIotIotFlowRuntimeActivateResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestIotIotFlowRuntimeActivateResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	iotFlowruntimeId := utils.GetEnvSettingWithBlankDefault("iot_flow_runtime_id")
	iotFlowruntimeIdVariableStr := fmt.Sprintf("variable \"iot_flow_runtime_id\" { default = \"%s\" }\n", iotFlowruntimeId)

	resourceName := "oci_iot_iot_flow_runtime_activate.test_iot_flow_runtime_activate"

	// Save TF content to Create resource with only required properties. This has to be exactly the same as the config part in the create step in the test.
	acctest.SaveConfigContent(config+iotFlowruntimeIdVariableStr+
		acctest.GenerateResourceFromRepresentationMap("oci_iot_iot_flow_runtime_activate", "test_iot_flow_runtime_activate", acctest.Required, acctest.Create, IotIotFlowRuntimeActivateRepresentation), "iot", "iotFlowRuntimeActivate", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify Create
		{
			Config: config + iotFlowruntimeIdVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_iot_iot_flow_runtime_activate", "test_iot_flow_runtime_activate", acctest.Required, acctest.Create, IotIotFlowRuntimeActivateRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "iot_flow_runtime_id"),
			),
		},
	})
}
