---
subcategory: "Fleet Apps Management"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_fleet_apps_management_task_records"
sidebar_current: "docs-oci-datasource-fleet_apps_management-task_records"
description: |-
  Provides the list of Task Records in Oracle Cloud Infrastructure Fleet Apps Management service
---

# Data Source: oci_fleet_apps_management_task_records
This data source provides the list of Task Records in Oracle Cloud Infrastructure Fleet Apps Management service.

Returns a list of all the task records in the specified compartment.
The query parameter `compartmentId` is required unless the query parameter `id` is specified.


## Example Usage

```hcl
data "oci_fleet_apps_management_task_records" "test_task_records" {

	#Optional
	compartment_id = var.compartment_id
	display_name = var.task_record_display_name
	id = var.task_record_id
	operation = var.task_record_operation
	platform = var.task_record_platform
	state = var.task_record_state
	type = var.task_record_type
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Optional) The ID of the compartment in which to list resources. Empty only if the resource OCID query param is not specified. 
* `display_name` - (Optional) A filter to return only resources that match the entire display name given.
* `id` - (Optional) Unique identifier or OCID for listing a single task record by id. Either compartmentId or id must be provided. 
* `operation` - (Optional) A filter to return task records whose operation matches the given lifecycle operation.
* `platform` - (Optional) The platform for the task record.
* `state` - (Optional) The current state of the task record.
* `type` - (Optional) The type of the Task.


## Attributes Reference

The following attributes are exported:

* `task_record_collection` - The list of task_record_collection.

### TaskRecord Reference

The following attributes are exported:

* `compartment_id` - 
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `description` - A user-friendly description. To provide some insight about the resource. Avoid entering confidential information. 
* `details` - The details of the task.
	* `execution_details` - Execution details.
		* `catalog_id` - Catalog Id having terraform package.
		* `command` - Optional command to execute the content. You can provide any commands/arguments that can't be part of the script. 
		* `config_file` - Catalog Id having config file.
		* `content` - Content Source details.
			* `bucket` - Bucket Name.
			* `catalog_id` - Catalog Id having terraform package.
			* `checksum` - md5 checksum of the artifact.
			* `namespace` - Namespace.
			* `object` - Object Name.
			* `source_type` - Content Source type details. 
		* `credentials` - Credentials required for executing the task. 
			* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.  Example: `My new resource` 
			* `id` - The OCID of the resource.
		* `endpoint` - Endpoint to be invoked.
		* `execution_type` - The action type of the task
		* `is_executable_content` - Is the Content an executable file?
		* `is_locked` - Is the script locked to prevent changes directly in Object Storage?
		* `is_read_output_variable_enabled` - Is read output variable enabled
		* `target_compartment_id` - OCID of the compartment to which the resource belongs to.
		* `variables` - The variable of the task. At least one of the dynamicArguments or output needs to be provided. 
			* `input_variables` - The input variables for the task.
				* `description` - The description of the argument.
				* `name` - The name of the argument.
				* `type` - Input argument Type. 
			* `output_variables` - The list of output variables.
	* `is_apply_subject_task` - Is this an Apply Subject Task?  Set this to true for a Patch Execution Task which applies patches(subjects) on a target. 
	* `is_discovery_output_task` - Is this a discovery output task?
	* `operation` - The lifecycle operation performed by the runbook.
	* `os_type` - The OS for the task
	* `platform` - The platform of the runbook.
	* `properties` - The properties of the task.
		* `num_retries` - The number of retries allowed.
		* `timeout_in_seconds` - The timeout in seconds for the task.
	* `scope` - The scope of the task
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.  Example: `My new resource` 
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `id` - The OCID of the resource.
* `lifecycle_details` - A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
* `resource_region` - Associated region
* `state` - The current state of the task record.
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The time this resource was created. An RFC3339 formatted datetime string.
* `time_updated` - The time this resource was last updated. An RFC3339 formatted datetime string.
* `type` - Task type.
* `version` - The version of the task record.

