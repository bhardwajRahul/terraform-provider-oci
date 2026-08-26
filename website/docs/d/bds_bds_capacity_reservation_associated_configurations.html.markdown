---
subcategory: "Big Data Service"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_bds_bds_capacity_reservation_associated_configurations"
sidebar_current: "docs-oci-datasource-bds-bds_capacity_reservation_associated_configurations"
description: |-
  Provides the list of Bds Capacity Reservation Associated Configurations in Oracle Cloud Infrastructure Big Data Service service
---

# Data Source: oci_bds_bds_capacity_reservation_associated_configurations
This data source provides the list of Bds Capacity Reservation Associated Configurations in Oracle Cloud Infrastructure Big Data Service service.

Returns a list of BDS capacity reservation configurations associated with the specified BDS capacity reservation.


## Example Usage

```hcl
data "oci_bds_bds_capacity_reservation_associated_configurations" "test_bds_capacity_reservation_associated_configurations" {
	#Required
	bds_capacity_reservation_id = oci_bds_bds_capacity_reservation.test_bds_capacity_reservation.id
	compartment_id = var.compartment_id

	#Optional
	display_name = var.bds_capacity_reservation_associated_configuration_display_name
	state = var.bds_capacity_reservation_associated_configuration_state
}
```

## Argument Reference

The following arguments are supported:

* `bds_capacity_reservation_id` - (Required) The OCID of the BDS capacity reservation.
* `compartment_id` - (Required) The OCID of the compartment.
* `display_name` - (Optional) A filter to return only resources that match the entire display name given.
* `state` - (Optional) The lifecycle state of the BDS capacity reservation configuration.


## Attributes Reference

The following attributes are exported:

* `bds_capacity_reservation_associated_configuration_collection` - The list of bds_capacity_reservation_associated_configuration_collection.

### BdsCapacityReservationAssociatedConfiguration Reference

The following attributes are exported:

* `items` - List of BDS capacity reservation configuration summaries associated with a BDS capacity reservation.
	* `bds_instance_id` - The OCID of the BDS cluster linked through the BDS capacity reservation configuration.
	* `compartment_id` - The OCID of the compartment that contains the BDS cluster.
	* `display_name` - The display name of the BDS capacity reservation configuration.
	* `id` - The OCID of the BDS capacity reservation configuration.
	* `state` - The lifecycle state of the BDS capacity reservation configuration.
	* `time_created` - The time the configuration was created, shown as an RFC 3339 formatted datetime string.
	* `time_updated` - The time the configuration was updated, shown as an RFC 3339 formatted datetime string.

