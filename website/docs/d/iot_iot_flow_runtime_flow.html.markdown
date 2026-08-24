---
subcategory: "Iot"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_iot_iot_flow_runtime_flow"
sidebar_current: "docs-oci-datasource-iot-iot_flow_runtime_flow"
description: |-
  Provides details about a specific Iot Flow Runtime Flow in Oracle Cloud Infrastructure Iot service
---

# Data Source: oci_iot_iot_flow_runtime_flow
This data source provides details about a specific Iot Flow Runtime Flow resource in Oracle Cloud Infrastructure Iot service.

Gets the opaque flows document for the IoT flow runtime identified by the specified OCID.

## Example Usage

```hcl
data "oci_iot_iot_flow_runtime_flow" "test_iot_flow_runtime_flow" {
	#Required
	iot_flow_runtime_id = oci_iot_iot_flow_runtime.test_iot_flow_runtime.id
}
```

## Argument Reference

The following arguments are supported:

* `iot_flow_runtime_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of an IoT flow runtime.


## Attributes Reference

The following attributes are exported:


