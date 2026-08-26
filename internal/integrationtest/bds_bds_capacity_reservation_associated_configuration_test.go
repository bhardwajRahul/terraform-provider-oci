// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"

	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	BdsBdsCapacityReservationAssociatedConfigurationDataSourceRepresentation = map[string]interface{}{
		"bds_capacity_reservation_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_bds_bds_capacity_reservation.test_bds_capacity_reservation.id}`},
		"compartment_id":              acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":                acctest.Representation{RepType: acctest.Optional, Create: `displayName`},
		"state":                       acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
	}

	BdsBdsCapacityReservationAssociatedConfigurationResourceConfig = BdsBdsCapacityReservationResourceDependencies + acctest.GenerateResourceFromRepresentationMap("oci_bds_bds_capacity_reservation", "test_bds_capacity_reservation", acctest.Required, acctest.Create, BdsBdsCapacityReservationRepresentation)
)

// issue-routing-tag: bds/default
func TestBdsBdsCapacityReservationAssociatedConfigurationResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestBdsBdsCapacityReservationAssociatedConfigurationResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_bds_bds_capacity_reservation_associated_configurations.test_bds_capacity_reservation_associated_configurations"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_bds_bds_capacity_reservation_associated_configurations", "test_bds_capacity_reservation_associated_configurations", acctest.Required, acctest.Create, BdsBdsCapacityReservationAssociatedConfigurationDataSourceRepresentation) +
				compartmentIdVariableStr + BdsBdsCapacityReservationAssociatedConfigurationResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "bds_capacity_reservation_id"),
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),
				resource.TestCheckResourceAttrSet(datasourceName, "bds_capacity_reservation_associated_configuration_collection.#"),
				resource.TestCheckResourceAttr(datasourceName, "bds_capacity_reservation_associated_configuration_collection.0.items.#", "1"),
			),
		},
	})
}
