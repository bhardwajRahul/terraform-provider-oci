---
subcategory: "Fleet Apps Management"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_fleet_apps_management_runbooks"
sidebar_current: "docs-oci-datasource-fleet_apps_management-runbooks"
description: |-
  Provides the list of Runbooks in Oracle Cloud Infrastructure Fleet Apps Management service
---

# Data Source: oci_fleet_apps_management_runbooks
This data source provides the list of Runbooks in Oracle Cloud Infrastructure Fleet Apps Management service.

Returns a list of all the runbooks in the specified compartment.
The query parameter `compartmentId` is required unless the query parameter `id` is specified.


## Example Usage

```hcl
data "oci_fleet_apps_management_runbooks" "test_runbooks" {

	#Optional
	compartment_id = var.compartment_id
	display_name = var.runbook_display_name
	id = var.runbook_id
	operation = var.runbook_operation
	platform = var.runbook_platform
	state = var.runbook_state
	type = var.runbook_type
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Optional) The ID of the compartment in which to list resources. Empty only if the resource OCID query param is not specified. 
* `display_name` - (Optional) A filter to return only resources that match the entire display name given.
* `id` - (Optional) Unique identifier or OCID for listing a single Runbook by id. Either compartmentId or id must be provided. 
* `operation` - (Optional) A filter to return runbooks whose operation matches the given lifecycle operation.
* `platform` - (Optional) A filter to return runbooks whose platform matches the given platform.
* `state` - (Optional) A filter to return only resources whose lifecycleState matches the given lifecycleState.
* `type` - (Optional) A filter to return runbooks whose type matches the given type.


## Attributes Reference

The following attributes are exported:

* `runbook_collection` - The list of runbook_collection.

### Runbook Reference

The following attributes are exported:

* `compartment_id` - 
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `description` - A user-friendly description. To provide some insight about the resource. Avoid entering confidential information. 
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.  Example: `My new resource` 
* `estimated_time` - Estimated time to successfully complete the runbook execution.
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `has_draft_version` - Does this runbook has draft versions?
* `id` - The OCID of the resource.
* `is_default` - Is the runbook default? Sets this runbook as the default for the chosen product/product stack for the specified lifecycle operation. 
* `is_sudo_access_needed` - Does this runbook need SUDO access to execute?
* `latest_version` - Latest runbook version
* `lifecycle_details` - A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
* `operation` - The lifecycle operation performed by the runbook.
* `os_type` - The OS type for the runbook.
* `platform` - The platform of the runbook.
* `resource_region` - Associated region
* `runbook_version` - Version for the runbook.
	* `execution_workflow_details` - Execution Workflow details.
		* `workflow` - Execution Workflow for the runbook.
			* `group_name` - Name of the group.
			* `steps` - Steps within the Group.
				* `group_name` - Name of the group.
				* `step_name` - Provide StepName for the Task.
				* `steps` - Tasks within the Group. Provide the stepName for all applicable tasks. 
				* `type` - Content Source Details. 
			* `type` - Workflow Group  Details. 
	* `groups` - The groups of the runbook. 
		* `name` - The name of the group.
		* `properties` - The properties of the component.
			* `action_on_failure` - The action to be taken in case of a failure.
			* `notification_preferences` - Preferences to send notifications on the task activities.
				* `should_notify_on_pause` - Enables notification on pause.
				* `should_notify_on_task_failure` - Enables or disables notification on Task Failures.
				* `should_notify_on_task_success` - Enables or disables notification on Task Success.
			* `pause_details` - Pause Details
				* `duration_in_minutes` - Time in minutes to apply Pause.
				* `kind` - Pause based On. 
			* `pre_condition` - Build control flow conditions that determine the relevance of the task execution. 
			* `run_on` - The runon conditions
				* `condition` - Build control flow conditions that determine the relevance of the task execution. 
				* `host` - OCID of the self hosted instance.
				* `kind` - Run on based On. 
				* `previous_task_instance_details` - Previous Task Instance Details 
					* `output_variable_details` - The details of the output variable that will be used for mapping.
						* `output_variable_name` - The name of the output variable whose value has to be mapped.
						* `step_name` - The name of the task step the output variable belongs to.
					* `resource_id` - Resource Ocid.
					* `resource_type` - Resource Type.
		* `type` - The type of the group. PARALLEL_TASK_GROUP : Helps to execute tasks parallelly inside a resource. PARALLEL_RESOURCE_GROUP : Executes tasks across resources parallelly. ROLLING_RESOURCE_GROUP : Executes tasks across resources in a rolling order. 
	* `is_latest` - Is this version the latest? 
	* `rollback_workflow_details` - Rollback Workflow details.
		* `scope` - rollback Scope 
		* `workflow` - Rollback Workflow for the runbook.
			* `group_name` - Name of the group.
			* `steps` - Steps within the Group.
				* `group_name` - Name of the group.
				* `step_name` - Provide StepName for the Task.
				* `steps` - Tasks within the Group. Provide the stepName for all applicable tasks. 
				* `type` - Content Source Details. 
			* `type` - Workflow Group  Details. 
	* `tasks` - A set of tasks to execute in the runbook.
		* `output_variable_mappings` - Mapping output variables of previous tasks to the input variables of the current task.
			* `name` - The name of the input variable.
			* `output_variable_details` - The details of the output variable that will be used for mapping.
				* `output_variable_name` - The name of the output variable whose value has to be mapped.
				* `step_name` - The name of the task step the output variable belongs to.
		* `step_name` - The name of the task step.
		* `step_properties` - The properties of the component.
			* `action_on_failure` - The action to be taken in case of a failure.
			* `notification_preferences` - Preferences to send notifications on the task activities.
				* `should_notify_on_pause` - Enables notification on pause.
				* `should_notify_on_task_failure` - Enables or disables notification on Task Failures.
				* `should_notify_on_task_success` - Enables or disables notification on Task Success.
			* `pause_details` - Pause Details
				* `duration_in_minutes` - Time in minutes to apply Pause.
				* `kind` - Pause based On. 
			* `pre_condition` - Build control flow conditions that determine the relevance of the task execution. 
			* `run_on` - The runon conditions
				* `condition` - Build control flow conditions that determine the relevance of the task execution. 
				* `host` - OCID of the self hosted instance.
				* `kind` - Run on based On. 
				* `previous_task_instance_details` - Previous Task Instance Details 
					* `output_variable_details` - The details of the output variable that will be used for mapping.
						* `output_variable_name` - The name of the output variable whose value has to be mapped.
						* `step_name` - The name of the task step the output variable belongs to.
					* `resource_id` - Resource Ocid.
					* `resource_type` - Resource Type.
		* `task_record_details` - The details of the task.
			* `description` - The description of the task.
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
			* `is_apply_subject_task` - Is this an Apply Subject Task? Ex. Patch Execution Task
			* `is_copy_to_library_enabled` - Make a copy of this task in Library
			* `is_discovery_output_task` - Is this a discovery output task?
			* `name` - The name of the task
			* `os_type` - The OS for the task.
			* `platform` - The platform of the runbook.
			* `properties` - The properties of the task.
				* `num_retries` - The number of retries allowed.
				* `timeout_in_seconds` - The timeout in seconds for the task.
			* `scope` - The scope of the task.
			* `task_record_id` - The ID of taskRecord.
	* `version` - The version of the runbook.
* `state` - The current state of the runbook.
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The time this resource was created. An RFC3339 formatted datetime string.
* `time_updated` - The time this resource was last updated. An RFC3339 formatted datetime string.
* `type` - The type of the runbook.

