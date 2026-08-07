---
subcategory: "Big Data Service"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_bds_bds_capacity_reservation"
sidebar_current: "docs-oci-datasource-bds-bds_capacity_reservation"
description: |-
  Provides details about a specific Bds Capacity Reservation in Oracle Cloud Infrastructure Big Data Service service
---

# Data Source: oci_bds_bds_capacity_reservation
This data source provides details about a specific Bds Capacity Reservation resource in Oracle Cloud Infrastructure Big Data Service service.

Returns information about the BDS capacity reservation identified by the given ID.

## Example Usage

```hcl
data "oci_bds_bds_capacity_reservation" "test_bds_capacity_reservation" {
	#Required
	bds_capacity_reservation_id = oci_bds_bds_capacity_reservation.test_bds_capacity_reservation.id
}
```

## Argument Reference

The following arguments are supported:

* `bds_capacity_reservation_id` - (Required) The OCID of the BDS capacity reservation.


## Attributes Reference

The following attributes are exported:

* `compartment_id` - The OCID of the compartment that contains the BDS capacity reservation.
* `compute_capacity_reservations` - Compute capacity reservation ID mappings by domain. For a multi-AD region, domain1, domain2, and domain3 correspond to AD1, AD2, and AD3 respectively. For a single-AD region, domain1, domain2, and domain3 correspond to FD1, FD2, and FD3 respectively. 
	* `domain1reservation_id` - Compute capacity reservation OCID corresponding to AD1 for a multi-AD region or FD1 for a single-AD region.
	* `domain2reservation_id` - Compute capacity reservation OCID corresponding to AD2 for a multi-AD region or FD2 for a single-AD region.
	* `domain3reservation_id` - Compute capacity reservation OCID corresponding to AD3 for a multi-AD region or FD3 for a single-AD region.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For example, `{"foo-namespace": {"bar-key": "value"}}` 
* `display_name` - A user-friendly name for the BDS capacity reservation.
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type, or scope. Exists for cross-compatibility only. For example, `{"bar-key": "value"}` 
* `id` - The OCID of the BDS capacity reservation.
* `state` - The lifecycle state of the BDS capacity reservation.
* `system_tags` - Usage of system tag keys. These predefined keys are scoped to namespaces.
* `time_created` - The time the BDS capacity reservation was created, shown as an RFC 3339 formatted datetime string.
* `time_updated` - The time the BDS capacity reservation was updated, shown as an RFC 3339 formatted datetime string.
