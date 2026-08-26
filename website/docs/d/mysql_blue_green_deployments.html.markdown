---
subcategory: "MySQL Database"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_mysql_blue_green_deployments"
sidebar_current: "docs-oci-datasource-mysql-blue_green_deployments"
description: |-
  Provides the list of Blue Green Deployments in Oracle Cloud Infrastructure MySQL Database service
---

# Data Source: oci_mysql_blue_green_deployments
This data source provides the list of Blue Green Deployments in Oracle Cloud Infrastructure MySQL Database service.

Lists blue/green deployments in a compartment.

## Example Usage

```hcl
data "oci_mysql_blue_green_deployments" "test_blue_green_deployments" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	display_name = var.blue_green_deployment_display_name
	source_db_system_id = oci_mysql_mysql_db_system.test_mysql_db_system.id
	state = var.blue_green_deployment_state
	target_db_system_id = oci_mysql_mysql_db_system.test_mysql_db_system.id
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The compartment [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
* `display_name` - (Optional) Filters deployments by display name.
* `source_db_system_id` - (Optional) Source DB system [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) filter.
* `state` - (Optional) Filters deployments by lifecycle state.
* `target_db_system_id` - (Optional) Target DB system [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) filter.


## Attributes Reference

The following attributes are exported:

* `blue_green_deployment_collection` - The list of blue_green_deployment_collection.

### BlueGreenDeployment Reference

The following attributes are exported:

* `active_db_system_id` - The DB system [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) that currently owns the client-facing VIP and serves traffic.
* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `display_name` - The display name of the blue/green deployment.
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the blue/green deployment.
* `lifecycle_details` - Additional lifecycle details.
* `replication_channel_id` - Replication channel [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
* `source_db_system_id` - Blue/original source DB system [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) for the deployment pair.
* `ssl_mode` - SSL mode used for the replication channel created by the blue/green workflow.
* `state` - The current lifecycle state.
* `switchover_status` - Stage of the most recent switchover workflow. `SWITCHOVER_FAILED` indicates terminal switchover failure. 
* `system_tags` - Usage of system tag keys. These predefined keys are scoped to namespaces. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `target_db_system_details` - Target DB System details for a blue/green deployment.
	* `configuration_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the configuration applied to the target DB System.
	* `data_storage_size_in_gb` - Initial data storage size in GiBs for the target DB System.
	* `mysql_version` - Target MySQL engine version.
	* `shape_name` - The shape of the target DB System. The shape determines resources allocated to the DB System - CPU cores and memory for VM shapes; CPU cores, memory and storage for non-VM (or bare metal) shapes. To get a list of shapes, use the [ListShapes](https://docs.cloud.oracle.com/iaas/api/#/en/mysql/20181021/ShapeSummary/ListShapes) operation. 
* `target_db_system_id` - Green/target DB system [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) for the deployment pair.
* `time_created` - The time the deployment was created.
* `time_updated` - The time the deployment was last updated.

