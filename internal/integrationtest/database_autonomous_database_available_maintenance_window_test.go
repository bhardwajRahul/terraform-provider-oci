// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"

	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	autonomousDatabaseRepresentationRegular = acctest.RepresentationCopyWithNewProperties(acctest.RepresentationCopyWithRemovedProperties(DatabaseAutonomousDatabaseRepresentation, []string{"cpu_core_count", "db_tools_details", "display_name"}), map[string]interface{}{
		"compute_count":            acctest.Representation{RepType: acctest.Required, Create: `4.0`},
		"compute_model":            acctest.Representation{RepType: acctest.Required, Create: `ECPU`},
		"data_storage_size_in_tbs": acctest.Representation{RepType: acctest.Required, Create: `5`},
	})

	DayOfWeekMonday = map[string]interface{}{
		"name": acctest.Representation{RepType: acctest.Required, Create: `MONDAY`},
	}

	DayOfWeekWednesday = map[string]interface{}{
		"name": acctest.Representation{RepType: acctest.Required, Create: `WEDNESDAY`},
	}

	ScheduledMaintenanceWindowRepresentationInitial = map[string]interface{}{
		"day_of_week":                            acctest.RepresentationGroup{RepType: acctest.Required, Group: DayOfWeekMonday},
		"availability_domain":                    acctest.Representation{RepType: acctest.Required, Create: `kIdk:PHX-AD-1`},
		"maintenance_start_time":                 acctest.Representation{RepType: acctest.Required, Create: `02:00`},
		"maintenance_end_time":                   acctest.Representation{RepType: acctest.Required, Create: `05:00`},
		"is_maintenance_window_change_scheduled": acctest.Representation{RepType: acctest.Required, Create: `true`},
	}

	ScheduledMaintenanceWindowRepresentationUpdated = map[string]interface{}{
		"day_of_week":                            acctest.RepresentationGroup{RepType: acctest.Required, Group: DayOfWeekWednesday},
		"availability_domain":                    acctest.Representation{RepType: acctest.Required, Create: `kIdk:PHX-AD-2`},
		"maintenance_start_time":                 acctest.Representation{RepType: acctest.Required, Create: `06:00`},
		"maintenance_end_time":                   acctest.Representation{RepType: acctest.Required, Create: `09:00`},
		"is_maintenance_window_change_scheduled": acctest.Representation{RepType: acctest.Required, Create: `true`},
	}

	ScheduledMaintenanceWindowRepresentationDisabled = map[string]interface{}{
		"is_maintenance_window_change_scheduled": acctest.Representation{RepType: acctest.Required, Create: `false`},
	}

	timeScheduled = time.Now().UTC().AddDate(0, 0, 5).Format(time.RFC3339)

	autonomousDatabaseRepresentationWithMWInitial = acctest.RepresentationCopyWithNewProperties(acctest.RepresentationCopyWithRemovedProperties(autonomousDatabaseRepresentationRegular, []string{"cpu_core_count", "db_tools_details", "display_name"}), map[string]interface{}{
		"scheduled_maintenance_window":             acctest.RepresentationGroup{RepType: acctest.Required, Group: ScheduledMaintenanceWindowRepresentationInitial},
		"time_scheduled_maintenance_window_update": acctest.Representation{RepType: acctest.Required, Create: timeScheduled},
	})

	autonomousDatabaseRepresentationWithMWUpdated = acctest.RepresentationCopyWithNewProperties(acctest.RepresentationCopyWithRemovedProperties(autonomousDatabaseRepresentationRegular, []string{"cpu_core_count", "db_tools_details", "display_name"}), map[string]interface{}{
		"scheduled_maintenance_window":             acctest.RepresentationGroup{RepType: acctest.Required, Group: ScheduledMaintenanceWindowRepresentationUpdated},
		"time_scheduled_maintenance_window_update": acctest.Representation{RepType: acctest.Required, Create: timeScheduled},
		"data_storage_size_in_tbs":                 acctest.Representation{RepType: acctest.Required, Create: `10`},
	})

	autonomousDatabaseRepresentationWithMWDisabled = acctest.RepresentationCopyWithNewProperties(acctest.RepresentationCopyWithRemovedProperties(autonomousDatabaseRepresentationRegular, []string{"cpu_core_count", "db_tools_details", "display_name"}), map[string]interface{}{
		"scheduled_maintenance_window":             acctest.RepresentationGroup{RepType: acctest.Required, Group: ScheduledMaintenanceWindowRepresentationDisabled},
		"time_scheduled_maintenance_window_update": acctest.Representation{RepType: acctest.Required, Create: timeScheduled},
		"data_storage_size_in_tbs":                 acctest.Representation{RepType: acctest.Required, Create: `10`},
	})

	DatabaseAutonomousDatabaseAvailableMaintenanceWindowDataSourceRepresentation = map[string]interface{}{
		"autonomous_database_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_database_autonomous_database.test_autonomous_database.id}`},
	}
)

// issue-routing-tag: database/dbaas-adb
func TestDatabaseAutonomousDatabaseAvailableMaintenanceWindowResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseAutonomousDatabaseAvailableMaintenanceWindowResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceAdbName := "oci_database_autonomous_database.test_autonomous_database"
	datasourceName := "data.oci_database_autonomous_database_available_maintenance_windows.test_autonomous_database_available_maintenance_windows"

	var resId, resId2 string

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		//0. Regular ADB create
		{
			Config: config + compartmentIdVariableStr + DatabaseAutonomousDatabaseResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", acctest.Required, acctest.Create, autonomousDatabaseRepresentationRegular),

			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceAdbName, "compute_count", "4"),
				resource.TestCheckResourceAttr(resourceAdbName, "data_storage_size_in_tbs", "5"),
			),
		},
		//1. Datasource: oci_database_autonomous_database_available_maintenance_windows
		{
			Config: config + compartmentIdVariableStr + DatabaseAutonomousDatabaseResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", acctest.Required, acctest.Create, autonomousDatabaseRepresentationRegular) +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_autonomous_database_available_maintenance_windows", "test_autonomous_database_available_maintenance_windows", acctest.Required, acctest.Create, DatabaseAutonomousDatabaseAvailableMaintenanceWindowDataSourceRepresentation),

			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "autonomous_database_id"),

				resource.TestCheckResourceAttrSet(datasourceName, "autonomous_database_maintenance_window_collection.#"),
				resource.TestCheckResourceAttrSet(datasourceName, "autonomous_database_maintenance_window_collection.0.items.#"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceAdbName, "id")
					return err
				},
			),
		},
		//2. Update initial Scheduled Maintenance Window
		{
			Config: config + compartmentIdVariableStr + DatabaseAutonomousDatabaseResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", acctest.Required, acctest.Create, autonomousDatabaseRepresentationWithMWInitial) +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_autonomous_database_available_maintenance_windows", "test_autonomous_database_available_maintenance_windows", acctest.Required, acctest.Create, DatabaseAutonomousDatabaseAvailableMaintenanceWindowDataSourceRepresentation),

			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceAdbName, "scheduled_maintenance_window.#", "1"),
				resource.TestCheckResourceAttr(resourceAdbName, "scheduled_maintenance_window.0.availability_domain", "kIdk:PHX-AD-1"),
				resource.TestCheckResourceAttr(resourceAdbName, "scheduled_maintenance_window.0.day_of_week.#", "1"),
				resource.TestCheckResourceAttr(resourceAdbName, "scheduled_maintenance_window.0.day_of_week.0.name", "MONDAY"),
				resource.TestCheckResourceAttr(resourceAdbName, "scheduled_maintenance_window.0.maintenance_start_time", "02:00"),
				resource.TestCheckResourceAttr(resourceAdbName, "scheduled_maintenance_window.0.maintenance_end_time", "05:00"),
				resource.TestCheckResourceAttr(resourceAdbName, "scheduled_maintenance_window.0.is_maintenance_window_change_scheduled", "true"),
				resource.TestCheckResourceAttr(resourceAdbName, "time_scheduled_maintenance_window_update", timeScheduled),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceAdbName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},
		//3. Update updated Scheduled Maintenance Window
		{
			Config: config + compartmentIdVariableStr + DatabaseAutonomousDatabaseResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", acctest.Required, acctest.Create, autonomousDatabaseRepresentationWithMWUpdated) +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_autonomous_database_available_maintenance_windows", "test_autonomous_database_available_maintenance_windows", acctest.Required, acctest.Create, DatabaseAutonomousDatabaseAvailableMaintenanceWindowDataSourceRepresentation),

			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceAdbName, "data_storage_size_in_tbs", "10"),
				resource.TestCheckResourceAttr(resourceAdbName, "scheduled_maintenance_window.#", "1"),
				resource.TestCheckResourceAttr(resourceAdbName, "scheduled_maintenance_window.0.availability_domain", "kIdk:PHX-AD-2"),
				resource.TestCheckResourceAttr(resourceAdbName, "scheduled_maintenance_window.0.day_of_week.#", "1"),
				resource.TestCheckResourceAttr(resourceAdbName, "scheduled_maintenance_window.0.day_of_week.0.name", "WEDNESDAY"),
				resource.TestCheckResourceAttr(resourceAdbName, "scheduled_maintenance_window.0.maintenance_start_time", "06:00"),
				resource.TestCheckResourceAttr(resourceAdbName, "scheduled_maintenance_window.0.maintenance_end_time", "09:00"),
				resource.TestCheckResourceAttr(resourceAdbName, "scheduled_maintenance_window.0.is_maintenance_window_change_scheduled", "true"),
				resource.TestCheckResourceAttr(resourceAdbName, "time_scheduled_maintenance_window_update", timeScheduled),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceAdbName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},
		//4. Disable Scheduled Maintenance Window
		{
			Config: config + compartmentIdVariableStr + DatabaseAutonomousDatabaseResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", acctest.Required, acctest.Create, autonomousDatabaseRepresentationWithMWDisabled) +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_autonomous_database_available_maintenance_windows", "test_autonomous_database_available_maintenance_windows", acctest.Required, acctest.Create, DatabaseAutonomousDatabaseAvailableMaintenanceWindowDataSourceRepresentation),

			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceAdbName, "scheduled_maintenance_window.0.is_maintenance_window_change_scheduled", "false"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceAdbName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},
	})
}
