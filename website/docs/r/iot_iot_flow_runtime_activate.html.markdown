---
subcategory: "Iot"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_iot_iot_flow_runtime_activate"
sidebar_current: "docs-oci-resource-iot-iot_flow_runtime_activate"
description: |-
  Provides the Iot Flow Runtime Activate resource in Oracle Cloud Infrastructure Iot service
---

# oci_iot_iot_flow_runtime_activate
This resource provides the Iot Flow Runtime Activate resource in Oracle Cloud Infrastructure Iot service.
Api doc link for the resource: https://docs.oracle.com/iaas/api/#/en/iot/latest/IotFlowRuntime/Activate

Example terraform configs related to the resource : https://github.com/oracle/terraform-provider-oci/tree/master/examples/iot

Activates the IoT flow runtime identified by the specified OCID.


## Example Usage

```hcl
resource "oci_iot_iot_flow_runtime_activate" "test_iot_flow_runtime_activate" {
	#Required
	iot_flow_runtime_id = oci_iot_iot_flow_runtime.test_iot_flow_runtime.id
}
```

## Argument Reference

The following arguments are supported:

* `iot_flow_runtime_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of an IoT flow runtime.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:


## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Iot Flow Runtime Activate
	* `update` - (Defaults to 20 minutes), when updating the Iot Flow Runtime Activate
	* `delete` - (Defaults to 20 minutes), when destroying the Iot Flow Runtime Activate


## Import

Import is not supported for this resource.

