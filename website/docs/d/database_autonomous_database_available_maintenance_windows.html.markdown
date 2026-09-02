---
subcategory: "Database"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_autonomous_database_available_maintenance_windows"
sidebar_current: "docs-oci-datasource-database-autonomous_database_available_maintenance_windows"
description: |-
  Provides the list of Autonomous Database Available Maintenance Windows in Oracle Cloud Infrastructure Database service
---

# Data Source: oci_database_autonomous_database_available_maintenance_windows
This data source provides the list of Autonomous Database Available Maintenance Windows in Oracle Cloud Infrastructure Database service.

Retrieves the list of available maintenance window options for the specified Autonomous AI Database.

## Example Usage

```hcl
data "oci_database_autonomous_database_available_maintenance_windows" "test_autonomous_database_available_maintenance_windows" {
	#Required
	autonomous_database_id = oci_database_autonomous_database.test_autonomous_database.id
}
```

## Argument Reference

The following arguments are supported:

* `autonomous_database_id` - (Required) The database [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).


## Attributes Reference

The following attributes are exported:

* `autonomous_database_maintenance_window_collection` - The list of autonomous_database_maintenance_window_collection.

### AutonomousDatabaseAvailableMaintenanceWindow Reference

The following attributes are exported:

* `items` - List of Autonomous AI Database maintenance windows.
	* `availability_domain` - The AD in which the maintenance will occur.
	* `day_of_week` - Day of the week.
		* `name` - Name of the day of the week.
	* `is_maintenance_window_change_scheduled` - Indicates if the maintenance window change is scheduled or not for the Autonomous AI Database.
	* `maintenance_end_time` - The maintenance end time. The value must use the ISO-8601 format "hh:mm".
	* `maintenance_start_time` - The maintenance start time. The value must use the ISO-8601 format "hh:mm".

