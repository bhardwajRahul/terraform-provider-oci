---
subcategory: "Big Data Service"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_bds_bds_capacity_reservation"
sidebar_current: "docs-oci-resource-bds-bds_capacity_reservation"
description: |-
  Provides the Bds Capacity Reservation resource in Oracle Cloud Infrastructure Big Data Service service
---

# oci_bds_bds_capacity_reservation
This resource provides the Bds Capacity Reservation resource in Oracle Cloud Infrastructure Big Data Service service.
Api doc link for the resource: https://docs.oracle.com/iaas/api/#/en/bigdata/latest/BdsCapacityReservation

Example terraform configs related to the resource : https://github.com/oracle/terraform-provider-oci/tree/master/examples/big_data_service

Creates a reusable BDS capacity reservation resource.


## Example Usage

```hcl
resource "oci_bds_bds_capacity_reservation" "test_bds_capacity_reservation" {
	#Required
	compartment_id = var.compartment_id
	compute_capacity_reservations {

		#Optional
		domain1reservation_id = var.domain1reservation_id
		domain2reservation_id = var.domain2reservation_id
		domain3reservation_id = var.domain3reservation_id
	}
	display_name = var.bds_capacity_reservation_display_name

	#Optional
	defined_tags = var.bds_capacity_reservation_defined_tags
	freeform_tags = var.bds_capacity_reservation_freeform_tags
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) (Updatable) The OCID of the compartment in which to create the BDS capacity reservation.
* `compute_capacity_reservations` - (Required) (Updatable) Compute capacity reservation ID mappings by domain. For a multi-AD region, domain1, domain2, and domain3 correspond to AD1, AD2, and AD3 respectively. For a single-AD region, domain1, domain2, and domain3 correspond to FD1, FD2, and FD3 respectively. 
	* `domain1reservation_id` - (Optional) (Updatable) Compute capacity reservation OCID corresponding to AD1 for a multi-AD region or FD1 for a single-AD region.
	* `domain2reservation_id` - (Optional) (Updatable) Compute capacity reservation OCID corresponding to AD2 for a multi-AD region or FD2 for a single-AD region.
	* `domain3reservation_id` - (Optional) (Updatable) Compute capacity reservation OCID corresponding to AD3 for a multi-AD region or FD3 for a single-AD region.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For example, `{"foo-namespace": {"bar-key": "value"}}` 
* `display_name` - (Required) (Updatable) A user-friendly name for the BDS capacity reservation.
* `freeform_tags` - (Optional) (Updatable) Simple key-value pair that is applied without any predefined name, type, or scope. Exists for cross-compatibility only. For example, `{"bar-key": "value"}` 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `compartment_id` - The OCID of the compartment that contains the BDS capacity reservation.
* `compute_capacity_reservations` - Compute capacity reservation ID mappings by domain. For a multi-AD region, domain1, domain2, and domain3 correspond to AD1, AD2, and AD3 respectively. For a single-AD region, domain1, domain2, and domain3 correspond to FD1, FD2, and FD3 respectively. 
	* `domain1reservation_id` - Capacity reservation OCID corresponding to AD1 for a multi-AD region or FD1 for a single-AD region.
	* `domain2reservation_id` - Capacity reservation OCID corresponding to AD2 for a multi-AD region or FD2 for a single-AD region.
	* `domain3reservation_id` - Capacity reservation OCID corresponding to AD3 for a multi-AD region or FD3 for a single-AD region.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For example, `{"foo-namespace": {"bar-key": "value"}}` 
* `display_name` - A user-friendly name for the BDS capacity reservation.
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type, or scope. Exists for cross-compatibility only. For example, `{"bar-key": "value"}` 
* `id` - The OCID of the BDS capacity reservation.
* `state` - The lifecycle state of the BDS capacity reservation.
* `system_tags` - Usage of system tag keys. These predefined keys are scoped to namespaces.
* `time_created` - The time the BDS capacity reservation was created, shown as an RFC 3339 formatted datetime string.
* `time_updated` - The time the BDS capacity reservation was updated, shown as an RFC 3339 formatted datetime string.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Bds Capacity Reservation
	* `update` - (Defaults to 20 minutes), when updating the Bds Capacity Reservation
	* `delete` - (Defaults to 20 minutes), when destroying the Bds Capacity Reservation


## Import

BdsCapacityReservations can be imported using the `id`, e.g.

```
$ terraform import oci_bds_bds_capacity_reservation.test_bds_capacity_reservation "id"
```
