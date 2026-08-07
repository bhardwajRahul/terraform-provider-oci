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
	oci_bds "github.com/oracle/oci-go-sdk/v65/bds"
	"github.com/oracle/oci-go-sdk/v65/common"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	BdsBdsCapacityReservationRequiredOnlyResource = BdsBdsCapacityReservationResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_bds_bds_capacity_reservation", "test_bds_capacity_reservation", acctest.Required, acctest.Create, BdsBdsCapacityReservationRepresentation)

	BdsBdsCapacityReservationResourceConfig = BdsBdsCapacityReservationResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_bds_bds_capacity_reservation", "test_bds_capacity_reservation", acctest.Optional, acctest.Update, BdsBdsCapacityReservationRepresentation)

	BdsBdsCapacityReservationSingularDataSourceRepresentation = map[string]interface{}{
		"bds_capacity_reservation_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_bds_bds_capacity_reservation.test_bds_capacity_reservation.id}`},
	}

	BdsBdsCapacityReservationDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"state":          acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: BdsBdsCapacityReservationDataSourceFilterRepresentation}}
	BdsBdsCapacityReservationDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_bds_bds_capacity_reservation.test_bds_capacity_reservation.id}`}},
	}

	BdsBdsCapacityReservationRepresentation = map[string]interface{}{
		"compartment_id":                acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"compute_capacity_reservations": acctest.RepresentationGroup{RepType: acctest.Required, Group: BdsBdsCapacityReservationComputeCapacityReservationsRepresentation},
		"display_name":                  acctest.Representation{RepType: acctest.Required, Create: `displayName`, Update: `displayName2`},
		// "defined_tags":                  acctest.Representation{RepType: acctest.Optional, Create: `${tomap({"${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}" = "value"})}`, Update: `${tomap({"${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}" = "updatedValue"})}`},
		"freeform_tags": acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"freeformTags": "freeformTags"}, Update: map[string]string{"freeformTags2": "freeformTags2"}},
	}
	BdsBdsCapacityReservationComputeCapacityReservationsRepresentation = map[string]interface{}{
		"domain1reservation_id": acctest.Representation{RepType: acctest.Required, Create: `${var.domain1reservation_id}`},
		"domain2reservation_id": acctest.Representation{RepType: acctest.Required, Create: `${var.domain2reservation_id}`},
		"domain3reservation_id": acctest.Representation{RepType: acctest.Required, Create: `${var.domain3reservation_id}`},
	}

	BdsBdsCapacityReservationResourceDependencies = `
variable "domain1reservation_id" {}
variable "domain2reservation_id" {}
variable "domain3reservation_id" {}
`
	// + DefinedTagsDependencies
)

// issue-routing-tag: bds/default
func TestBdsBdsCapacityReservationResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestBdsBdsCapacityReservationResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_bds_bds_capacity_reservation.test_bds_capacity_reservation"
	datasourceName := "data.oci_bds_bds_capacity_reservations.test_bds_capacity_reservations"
	singularDatasourceName := "data.oci_bds_bds_capacity_reservation.test_bds_capacity_reservation"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+BdsBdsCapacityReservationResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_bds_bds_capacity_reservation", "test_bds_capacity_reservation", acctest.Optional, acctest.Create, BdsBdsCapacityReservationRepresentation), "bds", "bdsCapacityReservation", t)

	acctest.ResourceTest(t, testAccCheckBdsBdsCapacityReservationDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + BdsBdsCapacityReservationResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_bds_bds_capacity_reservation", "test_bds_capacity_reservation", acctest.Required, acctest.Create, BdsBdsCapacityReservationRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "compute_capacity_reservations.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + BdsBdsCapacityReservationResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + BdsBdsCapacityReservationResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_bds_bds_capacity_reservation", "test_bds_capacity_reservation", acctest.Optional, acctest.Create, BdsBdsCapacityReservationRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "compute_capacity_reservations.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "compute_capacity_reservations.0.domain1reservation_id"),
				resource.TestCheckResourceAttrSet(resourceName, "compute_capacity_reservations.0.domain2reservation_id"),
				resource.TestCheckResourceAttrSet(resourceName, "compute_capacity_reservations.0.domain3reservation_id"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),

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
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + BdsBdsCapacityReservationResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_bds_bds_capacity_reservation", "test_bds_capacity_reservation", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(BdsBdsCapacityReservationRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "compute_capacity_reservations.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "compute_capacity_reservations.0.domain1reservation_id"),
				resource.TestCheckResourceAttrSet(resourceName, "compute_capacity_reservations.0.domain2reservation_id"),
				resource.TestCheckResourceAttrSet(resourceName, "compute_capacity_reservations.0.domain3reservation_id"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),

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
			Config: config + compartmentIdVariableStr + BdsBdsCapacityReservationResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_bds_bds_capacity_reservation", "test_bds_capacity_reservation", acctest.Optional, acctest.Update, BdsBdsCapacityReservationRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "compute_capacity_reservations.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "compute_capacity_reservations.0.domain1reservation_id"),
				resource.TestCheckResourceAttrSet(resourceName, "compute_capacity_reservations.0.domain2reservation_id"),
				resource.TestCheckResourceAttrSet(resourceName, "compute_capacity_reservations.0.domain3reservation_id"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_bds_bds_capacity_reservations", "test_bds_capacity_reservations", acctest.Optional, acctest.Update, BdsBdsCapacityReservationDataSourceRepresentation) +
				compartmentIdVariableStr + BdsBdsCapacityReservationResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_bds_bds_capacity_reservation", "test_bds_capacity_reservation", acctest.Optional, acctest.Update, BdsBdsCapacityReservationRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),
				resource.TestCheckResourceAttr(datasourceName, "bds_capacity_reservation_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "bds_capacity_reservation_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_bds_bds_capacity_reservation", "test_bds_capacity_reservation", acctest.Required, acctest.Create, BdsBdsCapacityReservationSingularDataSourceRepresentation) +
				compartmentIdVariableStr + BdsBdsCapacityReservationResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "bds_capacity_reservation_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "compute_capacity_reservations.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// verify resource import
		{
			Config:                  config + BdsBdsCapacityReservationRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckBdsBdsCapacityReservationDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).BdsClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_bds_bds_capacity_reservation" {
			noResourceFound = false
			request := oci_bds.GetBdsCapacityReservationRequest{}

			tmp := rs.Primary.ID
			request.BdsCapacityReservationId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "bds")

			response, err := client.GetBdsCapacityReservation(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_bds.BdsCapacityReservationLifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("BdsBdsCapacityReservation") {
		resource.AddTestSweepers("BdsBdsCapacityReservation", &resource.Sweeper{
			Name:         "BdsBdsCapacityReservation",
			Dependencies: acctest.DependencyGraph["bdsCapacityReservation"],
			F:            sweepBdsBdsCapacityReservationResource,
		})
	}
}

func sweepBdsBdsCapacityReservationResource(compartment string) error {
	bdsClient := acctest.GetTestClients(&schema.ResourceData{}).BdsClient()
	bdsCapacityReservationIds, err := getBdsBdsCapacityReservationIds(compartment)
	if err != nil {
		return err
	}
	for _, bdsCapacityReservationId := range bdsCapacityReservationIds {
		if ok := acctest.SweeperDefaultResourceId[bdsCapacityReservationId]; !ok {
			deleteBdsCapacityReservationRequest := oci_bds.DeleteBdsCapacityReservationRequest{}

			deleteBdsCapacityReservationRequest.BdsCapacityReservationId = &bdsCapacityReservationId

			deleteBdsCapacityReservationRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "bds")
			_, error := bdsClient.DeleteBdsCapacityReservation(context.Background(), deleteBdsCapacityReservationRequest)
			if error != nil {
				fmt.Printf("Error deleting BdsCapacityReservation %s %s, It is possible that the resource is already deleted. Please verify manually \n", bdsCapacityReservationId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &bdsCapacityReservationId, BdsBdsCapacityReservationSweepWaitCondition, time.Duration(3*time.Minute),
				BdsBdsCapacityReservationSweepResponseFetchOperation, "bds", true)
		}
	}
	return nil
}

func getBdsBdsCapacityReservationIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "BdsCapacityReservationId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	bdsClient := acctest.GetTestClients(&schema.ResourceData{}).BdsClient()

	listBdsCapacityReservationsRequest := oci_bds.ListBdsCapacityReservationsRequest{}
	listBdsCapacityReservationsRequest.CompartmentId = &compartmentId
	listBdsCapacityReservationsRequest.LifecycleState = oci_bds.BdsCapacityReservationLifecycleStateActive
	listBdsCapacityReservationsResponse, err := bdsClient.ListBdsCapacityReservations(context.Background(), listBdsCapacityReservationsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting BdsCapacityReservation list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, bdsCapacityReservation := range listBdsCapacityReservationsResponse.Items {
		id := *bdsCapacityReservation.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "BdsCapacityReservationId", id)
	}
	return resourceIds, nil
}

func BdsBdsCapacityReservationSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if bdsCapacityReservationResponse, ok := response.Response.(oci_bds.GetBdsCapacityReservationResponse); ok {
		return bdsCapacityReservationResponse.LifecycleState != oci_bds.BdsCapacityReservationLifecycleStateDeleted
	}
	return false
}

func BdsBdsCapacityReservationSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.BdsClient().GetBdsCapacityReservation(context.Background(), oci_bds.GetBdsCapacityReservationRequest{
		BdsCapacityReservationId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
