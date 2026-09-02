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
	oci_cluster_health "github.com/oracle/oci-go-sdk/v65/clusterhealth"
	"github.com/oracle/oci-go-sdk/v65/common"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	clusterHealthDiagnosisStoreObjectStoreBucket    = utils.GetEnvSettingWithDefault("cluster_health_object_store_bucket", "objectStoreBucket")
	clusterHealthDiagnosisStoreObjectStoreNamespace = utils.GetEnvSettingWithDefault("cluster_health_object_store_namespace", "objectStoreNamespace")

	ClusterHealthDiagnosisStoreRequiredOnlyResource = ClusterHealthDiagnosisStoreResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_cluster_health_diagnosis_store", "test_diagnosis_store", acctest.Required, acctest.Create, ClusterHealthDiagnosisStoreRepresentation)

	ClusterHealthDiagnosisStoreResourceConfig = ClusterHealthDiagnosisStoreResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_cluster_health_diagnosis_store", "test_diagnosis_store", acctest.Optional, acctest.Update, ClusterHealthDiagnosisStoreRepresentation)

	ClusterHealthDiagnosisStoreSingularDataSourceRepresentation = map[string]interface{}{
		"diagnosis_store_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_cluster_health_diagnosis_store.test_diagnosis_store.id}`},
	}

	ClusterHealthDiagnosisStoreDataSourceRepresentation = map[string]interface{}{
		"compartment_id":     acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
		"diagnosis_store_id": acctest.Representation{RepType: acctest.Optional, Create: `${oci_cluster_health_diagnosis_store.test_diagnosis_store.id}`},
		"display_name":       acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"state":              acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"filter":             acctest.RepresentationGroup{RepType: acctest.Required, Group: ClusterHealthDiagnosisStoreDataSourceFilterRepresentation}}
	ClusterHealthDiagnosisStoreDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_cluster_health_diagnosis_store.test_diagnosis_store.id}`}},
	}

	ClusterHealthDiagnosisStoreRepresentation = map[string]interface{}{
		"compartment_id":         acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":           acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"freeform_tags":          acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Finance"}},
		"object_store_bucket":    acctest.Representation{RepType: acctest.Required, Create: clusterHealthDiagnosisStoreObjectStoreBucket, Update: clusterHealthDiagnosisStoreObjectStoreBucket},
		"object_store_namespace": acctest.Representation{RepType: acctest.Required, Create: clusterHealthDiagnosisStoreObjectStoreNamespace, Update: clusterHealthDiagnosisStoreObjectStoreNamespace},
	}

	ClusterHealthDiagnosisStoreResourceDependencies = ""
)

// issue-routing-tag: cluster_health/default
func TestClusterHealthDiagnosisStoreResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestClusterHealthDiagnosisStoreResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_cluster_health_diagnosis_store.test_diagnosis_store"
	datasourceName := "data.oci_cluster_health_diagnosis_stores.test_diagnosis_stores"
	singularDatasourceName := "data.oci_cluster_health_diagnosis_store.test_diagnosis_store"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+ClusterHealthDiagnosisStoreResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_cluster_health_diagnosis_store", "test_diagnosis_store", acctest.Optional, acctest.Create, ClusterHealthDiagnosisStoreRepresentation), "clusterhealth", "diagnosisStore", t)

	acctest.ResourceTest(t, testAccCheckClusterHealthDiagnosisStoreDestroy, []resource.TestStep{
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + ClusterHealthDiagnosisStoreResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_cluster_health_diagnosis_store", "test_diagnosis_store", acctest.Optional, acctest.Create, ClusterHealthDiagnosisStoreRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "object_store_bucket", clusterHealthDiagnosisStoreObjectStoreBucket),
				resource.TestCheckResourceAttr(resourceName, "object_store_namespace", clusterHealthDiagnosisStoreObjectStoreNamespace),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),

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
			Config: config + compartmentIdVariableStr + ClusterHealthDiagnosisStoreResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_cluster_health_diagnosis_store", "test_diagnosis_store", acctest.Optional, acctest.Update, ClusterHealthDiagnosisStoreRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "object_store_bucket", clusterHealthDiagnosisStoreObjectStoreBucket),
				resource.TestCheckResourceAttr(resourceName, "object_store_namespace", clusterHealthDiagnosisStoreObjectStoreNamespace),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_cluster_health_diagnosis_stores", "test_diagnosis_stores", acctest.Optional, acctest.Update, ClusterHealthDiagnosisStoreDataSourceRepresentation) +
				compartmentIdVariableStr + ClusterHealthDiagnosisStoreResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_cluster_health_diagnosis_store", "test_diagnosis_store", acctest.Optional, acctest.Update, ClusterHealthDiagnosisStoreRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(datasourceName, "diagnosis_store_id"),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),

				resource.TestCheckResourceAttr(datasourceName, "diagnosis_store_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "diagnosis_store_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_cluster_health_diagnosis_store", "test_diagnosis_store", acctest.Required, acctest.Create, ClusterHealthDiagnosisStoreSingularDataSourceRepresentation) +
				compartmentIdVariableStr + ClusterHealthDiagnosisStoreResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "diagnosis_store_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "object_store_bucket", clusterHealthDiagnosisStoreObjectStoreBucket),
				resource.TestCheckResourceAttr(singularDatasourceName, "object_store_namespace", clusterHealthDiagnosisStoreObjectStoreNamespace),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// verify resource import
		{
			Config:                  config + ClusterHealthDiagnosisStoreRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckClusterHealthDiagnosisStoreDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).DiagnosisClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_cluster_health_diagnosis_store" {
			noResourceFound = false
			request := oci_cluster_health.GetDiagnosisStoreRequest{}

			tmp := rs.Primary.ID
			request.DiagnosisStoreId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "cluster_health")

			response, err := client.GetDiagnosisStore(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_cluster_health.DiagnosisStoreLifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("ClusterHealthDiagnosisStore") {
		resource.AddTestSweepers("ClusterHealthDiagnosisStore", &resource.Sweeper{
			Name:         "ClusterHealthDiagnosisStore",
			Dependencies: acctest.DependencyGraph["diagnosisStore"],
			F:            sweepClusterHealthDiagnosisStoreResource,
		})
	}
}

func sweepClusterHealthDiagnosisStoreResource(compartment string) error {
	diagnosisClient := acctest.GetTestClients(&schema.ResourceData{}).DiagnosisClient()
	diagnosisStoreIds, err := getClusterHealthDiagnosisStoreIds(compartment)
	if err != nil {
		return err
	}
	for _, diagnosisStoreId := range diagnosisStoreIds {
		if ok := acctest.SweeperDefaultResourceId[diagnosisStoreId]; !ok {
			deleteDiagnosisStoreRequest := oci_cluster_health.DeleteDiagnosisStoreRequest{}

			deleteDiagnosisStoreRequest.DiagnosisStoreId = &diagnosisStoreId

			deleteDiagnosisStoreRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "cluster_health")
			_, error := diagnosisClient.DeleteDiagnosisStore(context.Background(), deleteDiagnosisStoreRequest)
			if error != nil {
				fmt.Printf("Error deleting DiagnosisStore %s %s, It is possible that the resource is already deleted. Please verify manually \n", diagnosisStoreId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &diagnosisStoreId, ClusterHealthDiagnosisStoreSweepWaitCondition, time.Duration(3*time.Minute),
				ClusterHealthDiagnosisStoreSweepResponseFetchOperation, "cluster_health", true)
		}
	}
	return nil
}

func getClusterHealthDiagnosisStoreIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "DiagnosisStoreId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	diagnosisClient := acctest.GetTestClients(&schema.ResourceData{}).DiagnosisClient()

	listDiagnosisStoresRequest := oci_cluster_health.ListDiagnosisStoresRequest{}
	listDiagnosisStoresRequest.CompartmentId = &compartmentId
	listDiagnosisStoresRequest.LifecycleState = oci_cluster_health.DiagnosisStoreLifecycleStateActive
	listDiagnosisStoresResponse, err := diagnosisClient.ListDiagnosisStores(context.Background(), listDiagnosisStoresRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting DiagnosisStore list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, diagnosisStore := range listDiagnosisStoresResponse.Items {
		id := *diagnosisStore.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "DiagnosisStoreId", id)
	}
	return resourceIds, nil
}

func ClusterHealthDiagnosisStoreSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if diagnosisStoreResponse, ok := response.Response.(oci_cluster_health.GetDiagnosisStoreResponse); ok {
		return diagnosisStoreResponse.LifecycleState != oci_cluster_health.DiagnosisStoreLifecycleStateDeleted
	}
	return false
}

func ClusterHealthDiagnosisStoreSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.DiagnosisClient().GetDiagnosisStore(context.Background(), oci_cluster_health.GetDiagnosisStoreRequest{
		DiagnosisStoreId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
