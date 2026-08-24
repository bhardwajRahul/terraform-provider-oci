// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/oracle/oci-go-sdk/v65/common"
	oci_iot "github.com/oracle/oci-go-sdk/v65/iot"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	IotIotFlowRuntimeRequiredOnlyResource = IotIotFlowRuntimeResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_iot_iot_flow_runtime", "test_iot_flow_runtime", acctest.Required, acctest.Create, IotIotFlowRuntimeRepresentation)

	IotIotFlowRuntimeResourceConfig = IotIotFlowRuntimeResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_iot_iot_flow_runtime", "test_iot_flow_runtime", acctest.Optional, acctest.Update, IotIotFlowRuntimeRepresentation)

	IotIotFlowRuntimeSingularDataSourceRepresentation = map[string]interface{}{
		"iot_flow_runtime_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_iot_iot_flow_runtime.test_iot_flow_runtime.id}`},
	}

	IotIotFlowRuntimeDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"id":             acctest.Representation{RepType: acctest.Optional, Create: `${oci_iot_iot_flow_runtime.test_iot_flow_runtime.id}`},
		"iot_domain_id":  acctest.Representation{RepType: acctest.Optional, Create: `${var.iot_domain_id}`},
		"state":          acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: IotIotFlowRuntimeDataSourceFilterRepresentation}}
	IotIotFlowRuntimeDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_iot_iot_flow_runtime.test_iot_flow_runtime.id}`}},
	}
	ignoreIotFlowRuntimeDefinedTagsChangesRepresentation = map[string]interface{}{
		"ignore_changes": acctest.Representation{RepType: acctest.Required, Create: []string{`defined_tags`}},
	}

	IotIotFlowRuntimeRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"iot_domain_id":  acctest.Representation{RepType: acctest.Required, Create: `${var.iot_domain_id}`},
		"defined_tags":   acctest.Representation{RepType: acctest.Optional, Create: `${tomap({"${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}" = "value"})}`, Update: `${tomap({"${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}" = "updatedValue"})}`},
		"description":    acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"freeform_tags":  acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"scale":          acctest.Representation{RepType: acctest.Optional, Create: `MEDIUM`, Update: `HIGHEST`},
		"lifecycle":      acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreIotFlowRuntimeDefinedTagsChangesRepresentation},
	}

	IotIotFlowRuntimeResourceDependencies = DefinedTagsDependencies
)

// issue-routing-tag: iot/default
func TestIotIotFlowRuntimeResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestIotIotFlowRuntimeResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)
	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	iotDomainId := utils.GetEnvSettingWithBlankDefault("iot_domain_ocid")
	iotDomainIdVariableStr := fmt.Sprintf("variable \"iot_domain_id\" { default = \"%s\" }\n", iotDomainId)

	resourceName := "oci_iot_iot_flow_runtime.test_iot_flow_runtime"
	datasourceName := "data.oci_iot_iot_flow_runtimes.test_iot_flow_runtimes"
	singularDatasourceName := "data.oci_iot_iot_flow_runtime.test_iot_flow_runtime"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+iotDomainIdVariableStr+IotIotFlowRuntimeResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_iot_iot_flow_runtime", "test_iot_flow_runtime", acctest.Optional, acctest.Create, IotIotFlowRuntimeRepresentation), "iot", "iotFlowRuntime", t)

	acctest.ResourceTest(t, testAccCheckIotIotFlowRuntimeDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + iotDomainIdVariableStr + IotIotFlowRuntimeResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_iot_iot_flow_runtime", "test_iot_flow_runtime", acctest.Required, acctest.Create, IotIotFlowRuntimeRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "iot_domain_id"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + iotDomainIdVariableStr + IotIotFlowRuntimeResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + iotDomainIdVariableStr + IotIotFlowRuntimeResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_iot_iot_flow_runtime", "test_iot_flow_runtime", acctest.Optional, acctest.Create, IotIotFlowRuntimeRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "iot_domain_id"),
				resource.TestCheckResourceAttr(resourceName, "scale", "MEDIUM"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					if isEnableExportCompartment, _ := strconv.ParseBool(utils.GetEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
						if errExport := resourcediscovery.TestExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
							return errExport
						}
					}
					return err
				},
			),
		},

		// verify Update to the compartment (the compartment will be switched back in the next step)
		{
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + iotDomainIdVariableStr + IotIotFlowRuntimeResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_iot_iot_flow_runtime", "test_iot_flow_runtime", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(IotIotFlowRuntimeRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "iot_domain_id"),
				resource.TestCheckResourceAttr(resourceName, "scale", "MEDIUM"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("resource recreated when it was supposed to be updated")
					}
					return err
				},
			),
		},

		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + iotDomainIdVariableStr + IotIotFlowRuntimeResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_iot_iot_flow_runtime", "test_iot_flow_runtime", acctest.Optional, acctest.Update, IotIotFlowRuntimeRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "iot_domain_id"),
				resource.TestCheckResourceAttr(resourceName, "scale", "HIGHEST"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_iot_iot_flow_runtimes", "test_iot_flow_runtimes", acctest.Optional, acctest.Update, IotIotFlowRuntimeDataSourceRepresentation) +
				compartmentIdVariableStr + iotDomainIdVariableStr + IotIotFlowRuntimeResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_iot_iot_flow_runtime", "test_iot_flow_runtime", acctest.Optional, acctest.Update, IotIotFlowRuntimeRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "iot_domain_id"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),

				resource.TestCheckResourceAttr(datasourceName, "iot_flow_runtime_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "iot_flow_runtime_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_iot_iot_flow_runtime", "test_iot_flow_runtime", acctest.Required, acctest.Create, IotIotFlowRuntimeSingularDataSourceRepresentation) +
				compartmentIdVariableStr + iotDomainIdVariableStr + IotIotFlowRuntimeResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "iot_flow_runtime_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "flow_runtime_host"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "scale", "HIGHEST"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// verify resource import
		{
			Config:                  config + compartmentIdVariableStr + iotDomainIdVariableStr + IotIotFlowRuntimeRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckIotIotFlowRuntimeDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).IotClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_iot_iot_flow_runtime" {
			noResourceFound = false
			request := oci_iot.GetIotFlowRuntimeRequest{}

			tmp := rs.Primary.ID
			request.IotFlowRuntimeId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "iot")

			response, err := client.GetIotFlowRuntime(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_iot.IotFlowRuntimeLifecycleStateDeleted): true,
				}
				if _, ok := deletedLifecycleStates[string(response.LifecycleState)]; !ok {
					//resource lifecycle state is not in expected deleted lifecycle states.
					return fmt.Errorf("resource lifecycle state: %s is not in expected deleted lifecycle states", response.LifecycleState)
				}
				//resource lifecycle state is in expected deleted lifecycle states. continue with next one.
				continue
			}

			//Verify that exception is for '404 not found'.
			if failure, isServiceError := common.IsServiceError(err); !isServiceError || failure.GetHTTPStatusCode() != 404 {
				return err
			}
		}
	}
	if noResourceFound {
		return fmt.Errorf("at least one resource was expected from the state file, but could not be found")
	}

	return nil
}

func init() {
	if acctest.DependencyGraph == nil {
		acctest.InitDependencyGraph()
	}
	if !acctest.InSweeperExcludeList("IotIotFlowRuntime") {
		resource.AddTestSweepers("IotIotFlowRuntime", &resource.Sweeper{
			Name:         "IotIotFlowRuntime",
			Dependencies: acctest.DependencyGraph["iotFlowRuntime"],
			F:            sweepIotIotFlowRuntimeResource,
		})
	}
}

func sweepIotIotFlowRuntimeResource(compartment string) error {
	iotClient := acctest.GetTestClients(&schema.ResourceData{}).IotClient()
	iotFlowRuntimeIds, err := getIotIotFlowRuntimeIds(compartment)
	if err != nil {
		return err
	}
	for _, iotFlowRuntimeId := range iotFlowRuntimeIds {
		if ok := acctest.SweeperDefaultResourceId[iotFlowRuntimeId]; !ok {
			deleteIotFlowRuntimeRequest := oci_iot.DeleteIotFlowRuntimeRequest{}

			deleteIotFlowRuntimeRequest.IotFlowRuntimeId = &iotFlowRuntimeId

			deleteIotFlowRuntimeRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "iot")
			_, error := iotClient.DeleteIotFlowRuntime(context.Background(), deleteIotFlowRuntimeRequest)
			if error != nil {
				fmt.Printf("Error deleting IotFlowRuntime %s %s, It is possible that the resource is already deleted. Please verify manually \n", iotFlowRuntimeId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &iotFlowRuntimeId, IotIotFlowRuntimeSweepWaitCondition, time.Duration(3*time.Minute),
				IotIotFlowRuntimeSweepResponseFetchOperation, "iot", true)
		}
	}
	return nil
}

func getIotIotFlowRuntimeIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "IotFlowRuntimeId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	iotClient := acctest.GetTestClients(&schema.ResourceData{}).IotClient()

	listIotFlowRuntimesRequest := oci_iot.ListIotFlowRuntimesRequest{}
	listIotFlowRuntimesRequest.CompartmentId = &compartmentId
	listIotFlowRuntimesRequest.LifecycleState = oci_iot.IotFlowRuntimeLifecycleStateActive
	listIotFlowRuntimesResponse, err := iotClient.ListIotFlowRuntimes(context.Background(), listIotFlowRuntimesRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting IotFlowRuntime list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, iotFlowRuntime := range listIotFlowRuntimesResponse.Items {
		id := *iotFlowRuntime.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "IotFlowRuntimeId", id)
	}
	return resourceIds, nil
}

func IotIotFlowRuntimeSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if iotFlowRuntimeResponse, ok := response.Response.(oci_iot.GetIotFlowRuntimeResponse); ok {
		return iotFlowRuntimeResponse.LifecycleState != oci_iot.IotFlowRuntimeLifecycleStateDeleted
	}
	return false
}

func IotIotFlowRuntimeSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.IotClient().GetIotFlowRuntime(context.Background(), oci_iot.GetIotFlowRuntimeRequest{
		IotFlowRuntimeId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
