---
subcategory: "Iot"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_iot_iot_flow_runtime"
sidebar_current: "docs-oci-resource-iot-iot_flow_runtime"
description: |-
  Provides the Iot Flow Runtime resource in Oracle Cloud Infrastructure Iot service
---

# oci_iot_iot_flow_runtime
This resource provides the Iot Flow Runtime resource in Oracle Cloud Infrastructure Iot service.
Api doc link for the resource: https://docs.oracle.com/iaas/api/#/en/iot/latest/IotFlowRuntime

Example terraform configs related to the resource : https://github.com/oracle/terraform-provider-oci/tree/master/examples/iot

Creates an IoT flow runtime.


## Example Usage

```hcl
resource "oci_iot_iot_flow_runtime" "test_iot_flow_runtime" {
	#Required
	compartment_id = var.compartment_id
	iot_domain_id = oci_iot_iot_domain.test_iot_domain.id

	#Optional
	defined_tags = {"Operations.CostCenter"= "42"}
	description = var.iot_flow_runtime_description
	display_name = var.iot_flow_runtime_display_name
	freeform_tags = {"Department"= "Finance"}
	log_config {
		#Required
		log_group_id = oci_logging_log_group.test_log_group.id

		#Optional
		log_id = oci_apm_traces_log.test_log.id
	}
	network_config {
		#Required
		subnet_id = oci_core_subnet.test_subnet.id

		#Optional
		file_storage_mounts {
			#Required
			export_id = oci_file_storage_export.test_export.id
			mount_path = var.iot_flow_runtime_network_config_file_storage_mounts_mount_path
			mount_target_id = oci_file_storage_mount_target.test_mount_target.id
		}
		network_security_group_ids = var.iot_flow_runtime_network_config_network_security_group_ids
	}
	scale = var.iot_flow_runtime_scale
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment corresponding to the resource.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `description` - (Optional) (Updatable) A short description of the resource. 
* `display_name` - (Optional) (Updatable) A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `iot_domain_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the IoT domain.
* `log_config` - (Optional) (Updatable) Logging configuration for an IoT flow runtime.
	* `log_group_id` - (Required) (Updatable) The OCID of the log group associated with the IoT flow runtime.
	* `log_id` - (Optional) (Updatable) The OCID of the log that receives IoT flow runtime container logs.
* `network_config` - (Optional) (Updatable) Network access configuration for an IoT flow runtime.
	* `file_storage_mounts` - (Optional) (Updatable) The File Storage mounts to attach to the IoT flow runtime. If specified, this list must contain from 1 to 5 items.
		* `export_id` - (Required) (Updatable) The OCID of the File Storage export.
		* `mount_path` - (Required) (Updatable) The path relative to `/mnt` where the File Storage export is mounted. Do not include a leading slash or the `/mnt` prefix.
		* `mount_target_id` - (Required) (Updatable) The OCID of the File Storage mount target.
	* `network_security_group_ids` - (Optional) (Updatable) The OCIDs of the network security groups associated with the IoT flow runtime.
	* `subnet_id` - (Required) (Updatable) The OCID of the subnet used by the IoT flow runtime.
* `scale` - (Optional) (Updatable) The scale of the IoT flow runtime. Larger values allocate more CPU and memory for higher throughput and operational headroom. MEDIUM is the default value. 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment corresponding to the resource.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `description` - A short description of the resource. 
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.
* `flow_runtime_host` - Hostname of the IoT flow runtime.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the resource.
* `iot_domain_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the IoT domain.
* `log_config` - Logging configuration for an IoT flow runtime.
	* `log_group_id` - The OCID of the log group associated with the IoT flow runtime.
	* `log_id` - The OCID of the log that receives IoT flow runtime container logs.
* `network_config` - Network access configuration for an IoT flow runtime.
	* `file_storage_mounts` - The File Storage mounts to attach to the IoT flow runtime. If specified, this list must contain from 1 to 5 items.
		* `export_id` - The OCID of the File Storage export.
		* `mount_path` - The path relative to `/mnt` where the File Storage export is mounted. Do not include a leading slash or the `/mnt` prefix.
		* `mount_target_id` - The OCID of the File Storage mount target.
	* `network_security_group_ids` - The OCIDs of the network security groups associated with the IoT flow runtime.
	* `subnet_id` - The OCID of the subnet used by the IoT flow runtime.
* `scale` - The scale of the IoT flow runtime. Larger values allocate more CPU and memory for higher throughput and operational headroom. MEDIUM is the default value. 
* `state` - The current state of the IoT flow runtime.
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace.  Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The date and time when the resource was created, in the format defined by [RFC 3339](https://tools.ietf.org/html/rfc3339). Example: `2016-08-25T21:10:29.600Z` 
* `time_updated` - The date and time when the resource was last updated, in the format defined by [RFC 3339](https://tools.ietf.org/html/rfc3339). Example: `2016-08-25T21:10:29.600Z` 

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Iot Flow Runtime
	* `update` - (Defaults to 20 minutes), when updating the Iot Flow Runtime
	* `delete` - (Defaults to 20 minutes), when destroying the Iot Flow Runtime


## Import

IotFlowRuntimes can be imported using the `id`, e.g.

```
$ terraform import oci_iot_iot_flow_runtime.test_iot_flow_runtime "id"
```

