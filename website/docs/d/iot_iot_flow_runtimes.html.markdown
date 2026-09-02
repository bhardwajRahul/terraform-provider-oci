---
subcategory: "Iot"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_iot_iot_flow_runtimes"
sidebar_current: "docs-oci-datasource-iot-iot_flow_runtimes"
description: |-
  Provides the list of Iot Flow Runtimes in Oracle Cloud Infrastructure Iot service
---

# Data Source: oci_iot_iot_flow_runtimes
This data source provides the list of Iot Flow Runtimes in Oracle Cloud Infrastructure Iot service.

Lists IoT flow runtimes in the specified compartment.


## Example Usage

```hcl
data "oci_iot_iot_flow_runtimes" "test_iot_flow_runtimes" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	display_name = var.iot_flow_runtime_display_name
	id = var.iot_flow_runtime_id
	iot_domain_id = oci_iot_iot_domain.test_iot_domain.id
	state = var.iot_flow_runtime_state
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment in which to list resources.
* `display_name` - (Optional) Filter resources whose display name matches the specified value. 
* `id` - (Optional) Filter resources by [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm). Must be a valid OCID of the resource type. 
* `iot_domain_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the IoT domain in which to list flow runtime. 
* `state` - (Optional) Filter resources whose lifecycleState matches the specified value. 


## Attributes Reference

The following attributes are exported:

* `iot_flow_runtime_collection` - The list of iot_flow_runtime_collection.

### IotFlowRuntime Reference

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

