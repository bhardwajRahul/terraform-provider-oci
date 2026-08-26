---
subcategory: "Big Data Service"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_bds_bds_instance_bds_capacity_reservation_configurations"
sidebar_current: "docs-oci-datasource-bds-bds_instance_bds_capacity_reservation_configurations"
description: |-
  Provides the list of Bds Instance Bds Capacity Reservation Configurations in Oracle Cloud Infrastructure Big Data Service service
---

# Data Source: oci_bds_bds_instance_bds_capacity_reservation_configurations
This data source provides the list of Bds Instance Bds Capacity Reservation Configurations in Oracle Cloud Infrastructure Big Data Service service.

Returns a list of BDS capacity reservation configurations for the specified BDS cluster.


## Example Usage

```hcl
data "oci_bds_bds_instance_bds_capacity_reservation_configurations" "test_bds_instance_bds_capacity_reservation_configurations" {
	#Required
	bds_instance_id = oci_bds_bds_instance.test_bds_instance.id

	#Optional
	display_name = var.bds_instance_bds_capacity_reservation_configuration_display_name
	state = var.bds_instance_bds_capacity_reservation_configuration_state
}
```

## Argument Reference

The following arguments are supported:

* `bds_instance_id` - (Required) The OCID of the cluster.
* `display_name` - (Optional) A filter to return only resources that match the entire display name given.
* `state` - (Optional) The lifecycle state of the BDS capacity reservation configuration.


## Attributes Reference

The following attributes are exported:

* `bds_capacity_reservation_configuration_collection` - The list of bds_capacity_reservation_configuration_collection.

### BdsInstanceBdsCapacityReservationConfiguration Reference

The following attributes are exported:

* `items` - List of BDS capacity reservation configuration summaries for the specified BDS cluster.
	* `bds_capacity_reservation_id` - The OCID of the BDS capacity reservation associated with the BDS cluster.
	* `bds_instance_id` - The OCID of the BDS cluster associated with the BDS capacity reservation.
	* `display_name` - The display name of the BDS capacity reservation configuration.
	* `id` - The OCID of the BDS capacity reservation configuration.
	* `state` - The lifecycle state of the BDS capacity reservation configuration.
	* `time_created` - The time the BDS capacity reservation configuration was created, shown as an RFC 3339 formatted datetime string.
	* `time_updated` - The time the BDS capacity reservation configuration was updated, shown as an RFC 3339 formatted datetime string.
