// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	oci_bds "github.com/oracle/oci-go-sdk/v65/bds"
	"github.com/oracle/oci-go-sdk/v65/common"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	BdsBdsInstanceBdsCapacityReservationConfigurationResourceConfig = BdsBdsInstanceBdsCapacityReservationConfigurationResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_bds_bds_instance_bds_capacity_reservation_configuration", "test_bds_instance_bds_capacity_reservation_configuration", acctest.Optional, acctest.Update, BdsBdsInstanceBdsCapacityReservationConfigurationRepresentation)

	BdsBdsInstanceBdsCapacityReservationConfigurationSingularDataSourceRepresentation = map[string]interface{}{
		"bds_capacity_reservation_configuration_id": acctest.Representation{RepType: acctest.Required, Create: `${regex(".*/bdsCapacityReservationConfigurations/([^/]+)$", oci_bds_bds_instance_bds_capacity_reservation_configuration.test_bds_instance_bds_capacity_reservation_configuration.id)[0]}`},
		"bds_instance_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_bds_bds_instance.test_bds_instance.id}`},
	}

	BdsBdsInstanceBdsCapacityReservationConfigurationDataSourceRepresentation = map[string]interface{}{
		"bds_instance_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_bds_bds_instance.test_bds_instance.id}`},
		"display_name":    acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"state":           acctest.Representation{RepType: acctest.Optional, Create: `INACTIVE`},
		"filter":          acctest.RepresentationGroup{RepType: acctest.Required, Group: BdsBdsInstanceBdsCapacityReservationConfigurationDataSourceFilterRepresentation}}
	BdsBdsInstanceBdsCapacityReservationConfigurationDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `bds_capacity_reservation_id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_bds_bds_instance_bds_capacity_reservation_configuration.test_bds_instance_bds_capacity_reservation_configuration.bds_capacity_reservation_id}`}},
	}

	BdsBdsInstanceBdsCapacityReservationConfigurationRepresentation = map[string]interface{}{
		"bds_capacity_reservation_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_bds_bds_capacity_reservation.test_bds_capacity_reservation.id}`},
		"bds_instance_id":             acctest.Representation{RepType: acctest.Required, Create: `${oci_bds_bds_instance.test_bds_instance.id}`},
		"display_name":                acctest.Representation{RepType: acctest.Required, Create: `displayName`, Update: `displayName2`},
		"activate_trigger":            acctest.Representation{RepType: acctest.Optional, Create: `0`, Update: `1`},
		"deactivate_trigger":          acctest.Representation{RepType: acctest.Optional, Create: `0`, Update: `1`},
	}
	BdsBdsInstanceBdsCapacityReservationConfigurationResourceDependencies = BdsBdsCapacityReservationResourceDependencies +
		BdsBdsInstanceResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_bds_bds_capacity_reservation", "test_bds_capacity_reservation", acctest.Required, acctest.Create, BdsBdsCapacityReservationRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_bds_bds_instance", "test_bds_instance", acctest.Required, acctest.Create, BdsBdsInstanceRepresentation)
)

// issue-routing-tag: bds/default
func TestBdsBdsInstanceBdsCapacityReservationConfigurationResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestBdsBdsInstanceBdsCapacityReservationConfigurationResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_bds_bds_instance_bds_capacity_reservation_configuration.test_bds_instance_bds_capacity_reservation_configuration"
	datasourceName := "data.oci_bds_bds_instance_bds_capacity_reservation_configurations.test_bds_instance_bds_capacity_reservation_configurations"
	singularDatasourceName := "data.oci_bds_bds_instance_bds_capacity_reservation_configuration.test_bds_instance_bds_capacity_reservation_configuration"

	var resId, resId2 string
	acctest.SaveConfigContent(config+compartmentIdVariableStr+BdsBdsInstanceBdsCapacityReservationConfigurationResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_bds_bds_instance_bds_capacity_reservation_configuration", "test_bds_instance_bds_capacity_reservation_configuration", acctest.Required, acctest.Create, BdsBdsInstanceBdsCapacityReservationConfigurationRepresentation), "bds", "bdsInstanceBdsCapacityReservationConfiguration", t)

	acctest.ResourceTest(t, testAccCheckBdsBdsInstanceBdsCapacityReservationConfigurationDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + BdsBdsInstanceBdsCapacityReservationConfigurationResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_bds_bds_instance_bds_capacity_reservation_configuration", "test_bds_instance_bds_capacity_reservation_configuration", acctest.Required, acctest.Create, BdsBdsInstanceBdsCapacityReservationConfigurationRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "bds_capacity_reservation_id"),
				resource.TestCheckResourceAttrSet(resourceName, "bds_instance_id"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),

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

		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + BdsBdsInstanceBdsCapacityReservationConfigurationResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_bds_bds_instance_bds_capacity_reservation_configuration", "test_bds_instance_bds_capacity_reservation_configuration", acctest.Optional, acctest.Update, BdsBdsInstanceBdsCapacityReservationConfigurationRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "bds_capacity_reservation_id"),
				resource.TestCheckResourceAttrSet(resourceName, "bds_instance_id"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
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
				acctest.GenerateDataSourceFromRepresentationMap("oci_bds_bds_instance_bds_capacity_reservation_configurations", "test_bds_instance_bds_capacity_reservation_configurations", acctest.Required, acctest.Create, BdsBdsInstanceBdsCapacityReservationConfigurationDataSourceRepresentation) +
				compartmentIdVariableStr + BdsBdsInstanceBdsCapacityReservationConfigurationResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_bds_bds_instance_bds_capacity_reservation_configuration", "test_bds_instance_bds_capacity_reservation_configuration", acctest.Optional, acctest.Update, BdsBdsInstanceBdsCapacityReservationConfigurationRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "bds_instance_id"),
				resource.TestCheckResourceAttr(datasourceName, "bds_capacity_reservation_configuration_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "bds_capacity_reservation_configuration_collection.0.items.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "bds_capacity_reservation_configuration_collection.0.items.0.display_name", "displayName2"),
				resource.TestCheckResourceAttr(datasourceName, "bds_capacity_reservation_configuration_collection.0.items.0.state", "INACTIVE"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_bds_bds_instance_bds_capacity_reservation_configuration", "test_bds_instance_bds_capacity_reservation_configuration", acctest.Required, acctest.Create, BdsBdsInstanceBdsCapacityReservationConfigurationSingularDataSourceRepresentation) +
				compartmentIdVariableStr + BdsBdsInstanceBdsCapacityReservationConfigurationResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "bds_capacity_reservation_configuration_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "bds_instance_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
				testAccSetBdsCapacityReservationConfigurationIdForDelete(resourceName),
			),
		},
		// verify resource import
		{
			Config:                  config + compartmentIdVariableStr + BdsBdsInstanceBdsCapacityReservationConfigurationResourceConfig,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{"activate_trigger", "deactivate_trigger"},
			ResourceName:            resourceName,
		},
	})
}

func testAccSetBdsCapacityReservationConfigurationIdForDelete(resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("resource %s not found in state", resourceName)
		}
		id := rs.Primary.ID
		if idx := strings.LastIndex(id, "/"); idx >= 0 {
			id = id[idx+1:]
		}
		rs.Primary.Attributes["bds_capacity_reservation_id"] = id
		return nil
	}
}

func testAccCheckBdsBdsInstanceBdsCapacityReservationConfigurationDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).BdsClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_bds_bds_instance_bds_capacity_reservation_configuration" {
			noResourceFound = false
			request := oci_bds.GetBdsCapacityReservationConfigurationRequest{}

			if value := rs.Primary.ID; value != "" {
				if idx := strings.LastIndex(value, "/"); idx >= 0 {
					value = value[idx+1:]
				}
				request.BdsCapacityReservationConfigurationId = &value
			}

			if value, ok := rs.Primary.Attributes["bds_instance_id"]; ok {
				request.BdsInstanceId = &value
			}

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "bds")

			response, err := client.GetBdsCapacityReservationConfiguration(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_bds.BdsCapacityReservationConfigurationLifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("BdsBdsInstanceBdsCapacityReservationConfiguration") {
		resource.AddTestSweepers("BdsBdsInstanceBdsCapacityReservationConfiguration", &resource.Sweeper{
			Name:         "BdsBdsInstanceBdsCapacityReservationConfiguration",
			Dependencies: acctest.DependencyGraph["bdsInstanceBdsCapacityReservationConfiguration"],
			F:            sweepBdsBdsInstanceBdsCapacityReservationConfigurationResource,
		})
	}
}

func sweepBdsBdsInstanceBdsCapacityReservationConfigurationResource(compartment string) error {
	bdsClient := acctest.GetTestClients(&schema.ResourceData{}).BdsClient()
	bdsInstanceBdsCapacityReservationConfigurationIds, err := getBdsBdsInstanceBdsCapacityReservationConfigurationIds(compartment)
	if err != nil {
		return err
	}
	for _, bdsInstanceBdsCapacityReservationConfigurationId := range bdsInstanceBdsCapacityReservationConfigurationIds {
		if ok := acctest.SweeperDefaultResourceId[bdsInstanceBdsCapacityReservationConfigurationId]; !ok {
			deleteBdsCapacityReservationConfigurationRequest := oci_bds.DeleteBdsCapacityReservationConfigurationRequest{}

			deleteBdsCapacityReservationConfigurationRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "bds")
			_, error := bdsClient.DeleteBdsCapacityReservationConfiguration(context.Background(), deleteBdsCapacityReservationConfigurationRequest)
			if error != nil {
				fmt.Printf("Error deleting BdsInstanceBdsCapacityReservationConfiguration %s %s, It is possible that the resource is already deleted. Please verify manually \n", bdsInstanceBdsCapacityReservationConfigurationId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &bdsInstanceBdsCapacityReservationConfigurationId, BdsBdsInstanceBdsCapacityReservationConfigurationSweepWaitCondition, time.Duration(3*time.Minute),
				BdsBdsInstanceBdsCapacityReservationConfigurationSweepResponseFetchOperation, "bds", true)
		}
	}
	return nil
}

func getBdsBdsInstanceBdsCapacityReservationConfigurationIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "BdsInstanceBdsCapacityReservationConfigurationId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	bdsClient := acctest.GetTestClients(&schema.ResourceData{}).BdsClient()

	listBdsCapacityReservationConfigurationsRequest := oci_bds.ListBdsCapacityReservationConfigurationsRequest{}

	bdsInstanceIds, error := getBdsInstanceIds(compartment)
	if error != nil {
		return resourceIds, fmt.Errorf("Error getting bdsInstanceId required for BdsInstanceBdsCapacityReservationConfiguration resource requests \n")
	}
	for _, bdsInstanceId := range bdsInstanceIds {
		listBdsCapacityReservationConfigurationsRequest.BdsInstanceId = &bdsInstanceId

		listBdsCapacityReservationConfigurationsRequest.LifecycleState = oci_bds.BdsCapacityReservationConfigurationLifecycleStateActive
		listBdsCapacityReservationConfigurationsResponse, err := bdsClient.ListBdsCapacityReservationConfigurations(context.Background(), listBdsCapacityReservationConfigurationsRequest)

		if err != nil {
			return resourceIds, fmt.Errorf("Error getting BdsInstanceBdsCapacityReservationConfiguration list for compartment id : %s , %s \n", compartmentId, err)
		}
		for _, bdsInstanceBdsCapacityReservationConfiguration := range listBdsCapacityReservationConfigurationsResponse.Items {
			id := *bdsInstanceBdsCapacityReservationConfiguration.Id
			resourceIds = append(resourceIds, id)
			acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "BdsInstanceBdsCapacityReservationConfigurationId", id)
		}

	}
	return resourceIds, nil
}

func BdsBdsInstanceBdsCapacityReservationConfigurationSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if bdsInstanceBdsCapacityReservationConfigurationResponse, ok := response.Response.(oci_bds.GetBdsCapacityReservationConfigurationResponse); ok {
		return bdsInstanceBdsCapacityReservationConfigurationResponse.LifecycleState != oci_bds.BdsCapacityReservationConfigurationLifecycleStateDeleted
	}
	return false
}

func BdsBdsInstanceBdsCapacityReservationConfigurationSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.BdsClient().GetBdsCapacityReservationConfiguration(context.Background(), oci_bds.GetBdsCapacityReservationConfigurationRequest{RequestMetadata: common.RequestMetadata{
		RetryPolicy: retryPolicy,
	},
	})
	return err
}
