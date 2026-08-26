---
subcategory: "MySQL Database"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_mysql_blue_green_deployment"
sidebar_current: "docs-oci-resource-mysql-blue_green_deployment"
description: |-
  Provides the Blue Green Deployment resource in Oracle Cloud Infrastructure MySQL Database service
---

# oci_mysql_blue_green_deployment
This resource provides the Blue Green Deployment resource in Oracle Cloud Infrastructure MySQL Database service.
Api doc link for the resource: https://docs.oracle.com/iaas/api/#/en/mysql/latest/BlueGreenDeployment

Example terraform configs related to the resource : https://github.com/oracle/terraform-provider-oci/tree/master/examples/mysql

Creates a blue/green deployment resource and its replication channel.

## Example Usage

```hcl
resource "oci_mysql_blue_green_deployment" "test_blue_green_deployment" {
	#Required
	channel_details {
		#Required
		source_password = var.blue_green_deployment_channel_details_source_password
		source_username = var.blue_green_deployment_channel_details_source_username
		ssl_mode = var.blue_green_deployment_channel_details_ssl_mode

		#Optional
		applier_username = var.blue_green_deployment_channel_details_applier_username
		ssl_ca_certificate {
			#Required
			certificate_type = var.blue_green_deployment_channel_details_ssl_ca_certificate_certificate_type
			contents = var.blue_green_deployment_channel_details_ssl_ca_certificate_contents
		}
	}
	compartment_id = var.compartment_id
	display_name = var.blue_green_deployment_display_name
	source_db_system_id = oci_mysql_mysql_db_system.test_mysql_db_system.id
	target_db_system_details {
		#Required
		mysql_version = var.blue_green_deployment_target_db_system_details_mysql_version

		#Optional
		configuration_id = oci_mysql_mysql_configuration.test_mysql_configuration.id
		data_storage_size_in_gb = var.blue_green_deployment_target_db_system_details_data_storage_size_in_gb
		shape_name = oci_mysql_shape.test_shape.name
	}

	#Optional
	defined_tags = {"foo-namespace.bar-key"= "value"}
	freeform_tags = {"bar-key"= "value"}
}
```

## Argument Reference

The following arguments are supported:

* `channel_details` - (Required) Replication channel details for a blue/green deployment.
	* `applier_username` - (Optional) The username for the replication applier of the target MySQL DB System.
	* `source_password` - (Required) The password for the source DB system user used by the blue/green workflow to configure the replication channel. The password must be between 8 and 32 characters long, and must contain at least 1 numeric character, 1 lowercase character, 1 uppercase character, and 1 special (nonalphanumeric) character. 
	* `source_username` - (Required) The username on the source DB system used by the blue/green workflow to configure the replication channel. The username has a maximum length of 96 characters. For more information, please see the [MySQL documentation](https://dev.mysql.com/doc/refman/8.0/en/change-master-to.html) 
	* `ssl_ca_certificate` - (Optional) The CA certificate of the server used for VERIFY_IDENTITY and VERIFY_CA ssl modes.
		* `certificate_type` - (Required) The type of CA certificate.
		* `contents` - (Required) The string containing the CA certificate in PEM format.
	* `ssl_mode` - (Required) The SSL mode of the replication channel created by the blue/green workflow. `VERIFY_CA` and `VERIFY_IDENTITY` require `sslCaCertificate`. `REQUIRED` and `DISABLED` must not include `sslCaCertificate`. 
* `compartment_id` - (Required) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `delete_target_db_system_on_delete` - (Optional) (Updatable) Whether to delete the target DB System when this Blue/Green deployment is destroyed. The service deletes the associated replication channel as part of deleting the target DB System. Defaults to `false`. Run `terraform apply` after changing this value so it is stored in state before running `terraform destroy`.
* `display_name` - (Required) (Updatable) The display name of the blue/green deployment.
* `freeform_tags` - (Optional) (Updatable) Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `source_db_system_id` - (Required) Source DB system [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
* `target_db_system_details` - (Required) Target DB System overrides for a blue/green deployment.
	* `configuration_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the configuration to apply to the target DB System. If omitted, the target DB System inherits the source DB System configuration. 
	* `data_storage_size_in_gb` - (Optional) Initial data storage size in GiBs for the target DB System. If omitted, the target DB System uses the source DB System storage size. 
	* `mysql_version` - (Required) Target MySQL engine version.
	* `shape_name` - (Optional) The shape of the target DB System. The shape determines resources allocated to the DB System - CPU cores and memory for VM shapes; CPU cores, memory and storage for non-VM (or bare metal) shapes. To get a list of shapes, use the [ListShapes](https://docs.cloud.oracle.com/iaas/api/#/en/mysql/20181021/ShapeSummary/ListShapes) operation. 
* `switchover_trigger` - (Optional) (Updatable) An optional property when incremented triggers Switchover. Could be set to any integer value.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

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

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 60 minutes), when creating the Blue Green Deployment
	* `update` - (Defaults to 30 minutes), when updating the Blue Green Deployment
	* `delete` - (Defaults to 20 minutes), when destroying the Blue Green Deployment


## Import

BlueGreenDeployments can be imported using the `id`, e.g.

```
$ terraform import oci_mysql_blue_green_deployment.test_blue_green_deployment "id"
```
