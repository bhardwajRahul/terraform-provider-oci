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
	IotIotFlowRuntimeDeactivateRepresentation = map[string]interface{}{
		"iot_flow_runtime_id": acctest.Representation{RepType: acctest.Required, Create: `${var.iot_flow_runtime_id}`},
	}
)

// issue-routing-tag: iot/default
func TestIotIotFlowRuntimeDeactivateResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestIotIotFlowRuntimeDeactivateResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	iotFlowruntimeId := utils.GetEnvSettingWithBlankDefault("iot_flow_runtime_id")
	iotFlowruntimeIdVariableStr := fmt.Sprintf("variable \"iot_flow_runtime_id\" { default = \"%s\" }\n", iotFlowruntimeId)

	resourceName := "oci_iot_iot_flow_runtime_deactivate.test_iot_flow_runtime_deactivate"

	// Save TF content to Create resource with only required properties. This has to be exactly the same as the config part in the create step in the test.
	acctest.SaveConfigContent(config+iotFlowruntimeIdVariableStr+
		acctest.GenerateResourceFromRepresentationMap("oci_iot_iot_flow_runtime_deactivate", "test_iot_flow_runtime_deactivate", acctest.Required, acctest.Create, IotIotFlowRuntimeDeactivateRepresentation), "iot", "iotFlowRuntimeDeactivate", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify Create
		{
			Config: config + iotFlowruntimeIdVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_iot_iot_flow_runtime_deactivate", "test_iot_flow_runtime_deactivate", acctest.Required, acctest.Create, IotIotFlowRuntimeDeactivateRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "iot_flow_runtime_id"),
			),
		},
	})
}
