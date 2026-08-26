---
subcategory: "Big Data Service"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_bds_bds_instance_bds_capacity_reservation_configuration"
sidebar_current: "docs-oci-datasource-bds-bds_instance_bds_capacity_reservation_configuration"
description: |-
  Provides details about a specific Bds Instance Bds Capacity Reservation Configuration in Oracle Cloud Infrastructure Big Data Service service
---

# Data Source: oci_bds_bds_instance_bds_capacity_reservation_configuration
This data source provides details about a specific Bds Instance Bds Capacity Reservation Configuration resource in Oracle Cloud Infrastructure Big Data Service service.

Returns information about the BDS capacity reservation configuration identified by the given ID.

## Example Usage

```hcl
data "oci_bds_bds_instance_bds_capacity_reservation_configuration" "test_bds_instance_bds_capacity_reservation_configuration" {
	#Required
	bds_capacity_reservation_configuration_id = regex(".*/bdsCapacityReservationConfigurations/([^/]+)$", oci_bds_bds_instance_bds_capacity_reservation_configuration.test_bds_instance_bds_capacity_reservation_configuration.id)[0]
	bds_instance_id = oci_bds_bds_instance.test_bds_instance.id
}
```

## Argument Reference

The following arguments are supported:

* `bds_capacity_reservation_configuration_id` - (Required) The BDS capacity reservation configuration identifier. If using the resource ID from `oci_bds_bds_instance_bds_capacity_reservation_configuration`, extract the trailing `{bdsCapacityReservationConfigurationId}` value from the composite resource ID.
* `bds_instance_id` - (Required) The OCID of the cluster.


## Attributes Reference

The following attributes are exported:

* `bds_capacity_reservation_id` - The OCID of the BDS capacity reservation associated with the BDS cluster.
* `bds_instance_id` - The OCID of the BDS cluster associated with the BDS capacity reservation.
* `display_name` - The display name of the BDS capacity reservation configuration.
* `id` - The OCID of the BDS capacity reservation configuration.
* `state` - The lifecycle state of the BDS capacity reservation configuration.
* `time_created` - The time the BDS capacity reservation configuration was created, shown as an RFC 3339 formatted datetime string.
* `time_updated` - The time the BDS capacity reservation configuration was updated, shown as an RFC 3339 formatted datetime string.
