---
subcategory: "Database"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_cloud_exadata_infrastructure"
sidebar_current: "docs-oci-resource-database-cloud_exadata_infrastructure"
description: |-
  Provides the Cloud Exadata Infrastructure resource in Oracle Cloud Infrastructure Database service
---

# oci_database_cloud_exadata_infrastructure
This resource provides the Cloud Exadata Infrastructure resource in Oracle Cloud Infrastructure Database service.

Creates a cloud Exadata infrastructure resource. This resource is used to create either an [Exadata Cloud Service](https://docs.cloud.oracle.com/iaas/Content/Database/Concepts/exaoverview.htm) instance or an Autonomous Database on dedicated Exadata infrastructure.


## Example Usage

```hcl
resource "oci_database_cloud_exadata_infrastructure" "test_cloud_exadata_infrastructure" {
	#Required
	availability_domain = var.cloud_exadata_infrastructure_availability_domain
	compartment_id = var.compartment_id
	display_name = var.cloud_exadata_infrastructure_display_name
	shape = var.cloud_exadata_infrastructure_shape

	#Optional
    cluster_placement_group_id = var.cloud_exadata_infrastructure_cluster_placement_group_id
	compute_count = var.cloud_exadata_infrastructure_compute_count
	customer_contacts {

		#Optional
		email = var.cloud_exadata_infrastructure_customer_contacts_email
	}
	database_server_type = var.cloud_exadata_infrastructure_database_server_type
	defined_tags = var.cloud_exadata_infrastructure_defined_tags
	freeform_tags = {"Department"= "Finance"}
	maintenance_window {

		#Optional
		custom_action_timeout_in_mins = var.cloud_exadata_infrastructure_maintenance_window_custom_action_timeout_in_mins
		days_of_week {
			#Required
			name = var.cloud_exadata_infrastructure_maintenance_window_days_of_week_name
		}
		hours_of_day = var.cloud_exadata_infrastructure_maintenance_window_hours_of_day
		is_custom_action_timeout_enabled = var.cloud_exadata_infrastructure_maintenance_window_is_custom_action_timeout_enabled
		is_monthly_patching_enabled = var.cloud_exadata_infrastructure_maintenance_window_is_monthly_patching_enabled
		lead_time_in_weeks = var.cloud_exadata_infrastructure_maintenance_window_lead_time_in_weeks
		months {
			#Required
			name = var.cloud_exadata_infrastructure_maintenance_window_months_name
		}
		patching_mode = var.cloud_exadata_infrastructure_maintenance_window_patching_mode
		preference = var.cloud_exadata_infrastructure_maintenance_window_preference
		weeks_of_month = var.cloud_exadata_infrastructure_maintenance_window_weeks_of_month
	}
	storage_count = var.cloud_exadata_infrastructure_storage_count
	storage_server_type = var.cloud_exadata_infrastructure_storage_server_type
	subscription_id = var.tenant_subscription_id
}
```

## Argument Reference

The following arguments are supported:

* `availability_domain` - (Required) The availability domain where the cloud Exadata infrastructure is located.
* `cluster_placement_group_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the cluster placement group of the Exadata Infrastructure.
* `compartment_id` - (Required) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `compute_count` - (Optional) (Updatable) The number of compute servers for the cloud Exadata infrastructure.
* `customer_contacts` - (Optional) (Updatable) Customer contacts.
	* `email` - (Optional) (Updatable) The email address used by Oracle to send notifications regarding databases and infrastructure.
* `database_server_type` - (Optional) The database server type of the Exadata infrastructure.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). 
* `display_name` - (Required) (Updatable) The user-friendly name for the cloud Exadata infrastructure resource. The name does not need to be unique. 
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `maintenance_window` - (Optional) (Updatable) The scheduling details for the quarterly maintenance window. Patching and system updates take place during the maintenance window. 
	* `custom_action_timeout_in_mins` - (Optional) (Updatable) Determines the amount of time the system will wait before the start of each database server patching operation. Custom action timeout is in minutes and valid value is between 15 to 120 (inclusive). 
	* `days_of_week` - (Optional) (Updatable) Days during the week when maintenance should be performed.
		* `name` - (Required) (Updatable) Name of the day of the week.
	* `hours_of_day` - (Optional) (Updatable) The window of hours during the day when maintenance should be performed. The window is a 4 hour slot. Valid values are
		* 0 - represents time slot 0:00 - 3:59 UTC - 4 - represents time slot 4:00 - 7:59 UTC - 8 - represents time slot 8:00 - 11:59 UTC - 12 - represents time slot 12:00 - 15:59 UTC - 16 - represents time slot 16:00 - 19:59 UTC - 20 - represents time slot 20:00 - 23:59 UTC
	* `is_custom_action_timeout_enabled` - (Optional) (Updatable) If true, enables the configuration of a custom action timeout (waiting period) between database server patching operations.
	* `is_monthly_patching_enabled` - (Optional) (Updatable) If true, enables the monthly patching option.
	* `lead_time_in_weeks` - (Optional) (Updatable) Lead time window allows user to set a lead time to prepare for a down time. The lead time is in weeks and valid value is between 1 to 4. 
	* `months` - (Optional) (Updatable) Months during the year when maintenance should be performed.
		* `name` - (Required) (Updatable) Name of the month of the year.
	* `patching_mode` - (Optional) (Updatable) Cloud Exadata infrastructure node patching method, either "ROLLING" or "NONROLLING". Default value is ROLLING.

		*IMPORTANT*: Non-rolling infrastructure patching involves system down time. See [Oracle-Managed Infrastructure Maintenance Updates](https://docs.cloud.oracle.com/iaas/Content/Database/Concepts/examaintenance.htm#Oracle) for more information. 
	* `preference` - (Optional) (Updatable) The maintenance window scheduling preference.
	* `weeks_of_month` - (Optional) (Updatable) Weeks during the month when maintenance should be performed. Weeks start on the 1st, 8th, 15th, and 22nd days of the month, and have a duration of 7 days. Weeks start and end based on calendar dates, not days of the week. For example, to allow maintenance during the 2nd week of the month (from the 8th day to the 14th day of the month), use the value 2. Maintenance cannot be scheduled for the fifth week of months that contain more than 28 days. Note that this parameter works in conjunction with the  daysOfWeek and hoursOfDay parameters to allow you to specify specific days of the week and hours that maintenance will be performed. 
* `shape` - (Required) The shape of the cloud Exadata infrastructure resource. 
* `storage_count` - (Optional) (Updatable) The number of storage servers for the cloud Exadata infrastructure.
* `storage_server_type` - (Optional) The storage server type of the Exadata infrastructure.
* `subscription_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the subscription with which resource needs to be associated with.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `activated_storage_count` - The requested number of additional storage servers activated for the Exadata infrastructure.
* `additional_storage_count` - The requested number of additional storage servers for the Exadata infrastructure.
* `availability_domain` - The name of the availability domain that the cloud Exadata infrastructure resource is located in.
* `available_storage_size_in_gbs` - The available storage can be allocated to the cloud Exadata infrastructure resource, in gigabytes (GB).
* `cluster_placement_group_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the cluster placement group of the Exadata Infrastructure.
* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `compute_count` - The number of compute servers for the cloud Exadata infrastructure.
* `compute_model` - The compute model of the Exadata infrastructure.
* `cpu_count` - The total number of CPU cores allocated.
* `customer_contacts` - The list of customer email addresses that receive information from Oracle about the specified Oracle Cloud Infrastructure Database service resource. Oracle uses these email addresses to send notifications about planned and unplanned software maintenance updates, information about system hardware, and other information needed by administrators. Up to 10 email addresses can be added to the customer contacts for a cloud Exadata infrastructure instance. 
	* `email` - The email address used by Oracle to send notifications regarding databases and infrastructure.
* `data_storage_size_in_tbs` - Size, in terabytes, of the DATA disk group. 
* `database_server_type` - The database server type of the Exadata infrastructure.
* `db_node_storage_size_in_gbs` - The local node storage allocated in GBs.
* `db_server_version` - The software version of the database servers (dom0) in the cloud Exadata infrastructure. Example: 20.1.15 
* `defined_file_system_configurations` - Details of the file system configuration of the Exadata infrastructure.
	* `is_backup_partition` - If true, the file system is used to create a backup prior to Exadata VM OS update.
	* `is_resizable` - If true, the file system resize is allowed for the Exadata Infrastructure cluster. If false, the file system resize is not allowed.
	* `min_size_gb` - The minimum size of file system.
	* `mount_point` - The mount point of file system.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). 
* `display_name` - The user-friendly name for the cloud Exadata infrastructure resource. The name does not need to be unique.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the cloud Exadata infrastructure resource.
* `is_scheduling_policy_associated` - If true, the infrastructure is using granular maintenance scheduling preference.
* `last_maintenance_run_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the last maintenance run.
* `lifecycle_details` - Additional information about the current lifecycle state.
* `maintenance_window` - The scheduling details for the quarterly maintenance window. Patching and system updates take place during the maintenance window. 
	* `custom_action_timeout_in_mins` - Determines the amount of time the system will wait before the start of each database server patching operation. Custom action timeout is in minutes and valid value is between 15 to 120 (inclusive). 
	* `days_of_week` - Days during the week when maintenance should be performed.
		* `name` - Name of the day of the week.
	* `hours_of_day` - The window of hours during the day when maintenance should be performed. The window is a 4 hour slot. Valid values are
		* 0 - represents time slot 0:00 - 3:59 UTC - 4 - represents time slot 4:00 - 7:59 UTC - 8 - represents time slot 8:00 - 11:59 UTC - 12 - represents time slot 12:00 - 15:59 UTC - 16 - represents time slot 16:00 - 19:59 UTC - 20 - represents time slot 20:00 - 23:59 UTC
	* `is_custom_action_timeout_enabled` - If true, enables the configuration of a custom action timeout (waiting period) between database server patching operations.
	* `is_monthly_patching_enabled` - If true, enables the monthly patching option.
	* `lead_time_in_weeks` - Lead time window allows user to set a lead time to prepare for a down time. The lead time is in weeks and valid value is between 1 to 4. 
	* `months` - Months during the year when maintenance should be performed.
		* `name` - Name of the month of the year.
	* `patching_mode` - Cloud Exadata infrastructure node patching method, either "ROLLING" or "NONROLLING". Default value is ROLLING.

		*IMPORTANT*: Non-rolling infrastructure patching involves system down time. See [Oracle-Managed Infrastructure Maintenance Updates](https://docs.cloud.oracle.com/iaas/Content/Database/Concepts/examaintenance.htm#Oracle) for more information. 
	* `preference` - The maintenance window scheduling preference.
	* `weeks_of_month` - Weeks during the month when maintenance should be performed. Weeks start on the 1st, 8th, 15th, and 22nd days of the month, and have a duration of 7 days. Weeks start and end based on calendar dates, not days of the week. For example, to allow maintenance during the 2nd week of the month (from the 8th day to the 14th day of the month), use the value 2. Maintenance cannot be scheduled for the fifth week of months that contain more than 28 days. Note that this parameter works in conjunction with the  daysOfWeek and hoursOfDay parameters to allow you to specify specific days of the week and hours that maintenance will be performed. 
* `max_cpu_count` - The total number of CPU cores available.
* `max_data_storage_in_tbs` - The total available DATA disk group size.
* `max_db_node_storage_in_gbs` - The total local node storage available in GBs.
* `max_memory_in_gbs` - The total memory available in GBs.
* `memory_size_in_gbs` - The memory allocated in GBs.
* `monthly_db_server_version` - The monthly software version of the database servers (dom0) in the cloud Exadata infrastructure. Example: 20.1.15 
* `monthly_storage_server_version` - The monthly software version of the storage servers (cells) in the cloud Exadata infrastructure. Example: 20.1.15 
* `next_maintenance_run_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the next maintenance run.
* `shape` - The model name of the cloud Exadata infrastructure resource. 
* `state` - The current lifecycle state of the cloud Exadata infrastructure resource.
* `storage_count` - The number of storage servers for the cloud Exadata infrastructure.
* `storage_server_type` - The storage server type of the Exadata infrastructure.
* `storage_server_version` - The software version of the storage servers (cells) in the cloud Exadata infrastructure. Example: 20.1.15 
* `subscription_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the subscription with which resource needs to be associated with.
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). 
* `time_created` - The date and time the cloud Exadata infrastructure resource was created.
* `total_storage_size_in_gbs` - The total storage allocated to the cloud Exadata infrastructure resource, in gigabytes (GB).

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 12 hours), when creating the Cloud Exadata Infrastructure
	* `update` - (Defaults to 12 hours), when updating the Cloud Exadata Infrastructure
	* `delete` - (Defaults to 12 hours), when destroying the Cloud Exadata Infrastructure


## Import

CloudExadataInfrastructures can be imported using the `id`, e.g.

```
$ terraform import oci_database_cloud_exadata_infrastructure.test_cloud_exadata_infrastructure "id"
```

