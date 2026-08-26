---
subcategory: "Big Data Service"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_bds_bds_instance_bds_capacity_reservation_configuration"
sidebar_current: "docs-oci-resource-bds-bds_instance_bds_capacity_reservation_configuration"
description: |-
  Provides the Bds Instance Bds Capacity Reservation Configuration resource in Oracle Cloud Infrastructure Big Data Service service
---

# oci_bds_bds_instance_bds_capacity_reservation_configuration
This resource provides the Bds Instance Bds Capacity Reservation Configuration resource in Oracle Cloud Infrastructure Big Data Service service.
Api doc link for the resource: https://docs.oracle.com/iaas/api/#/en/bigdata/latest/BdsCapacityReservationConfiguration

Example terraform configs related to the resource : https://github.com/oracle/terraform-provider-oci/tree/master/examples/big_data_service

Creates a configuration between the specified BDS cluster and a BDS capacity reservation.


## Example Usage

```hcl
resource "oci_bds_bds_instance_bds_capacity_reservation_configuration" "test_bds_instance_bds_capacity_reservation_configuration" {
	#Required
	bds_capacity_reservation_id = oci_bds_bds_capacity_reservation.test_bds_capacity_reservation.id
	bds_instance_id = oci_bds_bds_instance.test_bds_instance.id
	display_name = var.bds_instance_bds_capacity_reservation_configuration_display_name
}
```

## Argument Reference

The following arguments are supported:

* `bds_capacity_reservation_id` - (Required) (Updatable) The OCID of the BDS capacity reservation to associate with the BDS cluster.
* `bds_instance_id` - (Required) The OCID of the cluster.
* `display_name` - (Required) (Updatable) A user-friendly name for the BDS capacity reservation configuration.
* `activate_trigger` - (Optional) (Updatable) An optional property that triggers Activate when its value is increased. If set during creation, Activate is called after the configuration is created.
* `deactivate_trigger` - (Optional) (Updatable) An optional property that triggers Deactivate when its value is increased. If set during creation, Deactivate is called after the configuration is created.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `bds_capacity_reservation_id` - The OCID of the BDS capacity reservation associated with the BDS cluster.
* `bds_instance_id` - The OCID of the BDS cluster associated with the BDS capacity reservation.
* `display_name` - The display name of the BDS capacity reservation configuration.
* `id` - The composite Terraform ID of the BDS capacity reservation configuration in the format `bdsInstances/{bdsInstanceId}/bdsCapacityReservationConfigurations/{bdsCapacityReservationConfigurationId}`.
* `state` - The lifecycle state of the BDS capacity reservation configuration.
* `time_created` - The time the BDS capacity reservation configuration was created, shown as an RFC 3339 formatted datetime string.
* `time_updated` - The time the BDS capacity reservation configuration was updated, shown as an RFC 3339 formatted datetime string.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Bds Instance Bds Capacity Reservation Configuration
	* `update` - (Defaults to 20 minutes), when updating the Bds Instance Bds Capacity Reservation Configuration
	* `delete` - (Defaults to 20 minutes), when destroying the Bds Instance Bds Capacity Reservation Configuration


## Import

BdsInstanceBdsCapacityReservationConfigurations can be imported using the `id`, e.g.

```
$ terraform import oci_bds_bds_instance_bds_capacity_reservation_configuration.test_bds_instance_bds_capacity_reservation_configuration "bdsInstances/{bdsInstanceId}/bdsCapacityReservationConfigurations/{bdsCapacityReservationConfigurationId}" 
```
