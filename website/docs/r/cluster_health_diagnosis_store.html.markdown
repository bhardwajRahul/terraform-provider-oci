---
subcategory: "Cluster Health"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_cluster_health_diagnosis_store"
sidebar_current: "docs-oci-resource-cluster_health-diagnosis_store"
description: |-
  Provides the Diagnosis Store resource in Oracle Cloud Infrastructure Cluster Health service
---

# oci_cluster_health_diagnosis_store
This resource provides the Diagnosis Store resource in Oracle Cloud Infrastructure Cluster Health service.
Api doc link for the resource: https://docs.oracle.com/iaas/api/#/en/

Example terraform configs related to the resource : https://github.com/oracle/terraform-provider-oci/tree/master/examples/cluster_health

Creates a Diagnosis Store.


## Example Usage

```hcl
resource "oci_cluster_health_diagnosis_store" "test_diagnosis_store" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	defined_tags = {"Operations.CostCenter"= "42"}
	display_name = var.diagnosis_store_display_name
	freeform_tags = {"Department"= "Finance"}
	object_store_bucket = var.diagnosis_store_object_store_bucket
	object_store_namespace = var.diagnosis_store_object_store_namespace
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment to create the Diagnosis in.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}`
* `display_name` - (Optional) (Updatable) A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}`
* `object_store_bucket` - (Optional) (Updatable) A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.
* `object_store_namespace` - (Optional) (Updatable) A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}`
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}`
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the DiagnosisStore.
* `last_accepted_request_token` - The last accepted request token.
* `last_completed_request_token` - The last completed request token.
* `lifecycle_details` - A message that describes the current state of the Diagnosis in more detail. For example, can be used to provide actionable information for a resource in the Failed state.
* `object_store_bucket` - The name of the object store bucket.
* `object_store_namespace` - The namespace of the object store.
* `state` - The current state of the Diagnosis.
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace.  Example: `{"orcl-cloud.free-tier-retained": "true"}`
* `time_created` - The date and time the Diagnosis Store was created, in the format defined by [RFC 3339](https://tools.ietf.org/html/rfc3339).  Example: `2016-08-25T21:10:29.600Z`
* `time_updated` - The date and time the Diagnosis Store was updated, in the format defined by [RFC 3339](https://tools.ietf.org/html/rfc3339).  Example: `2016-08-25T21:10:29.600Z`

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Diagnosis Store
	* `update` - (Defaults to 20 minutes), when updating the Diagnosis Store
	* `delete` - (Defaults to 20 minutes), when destroying the Diagnosis Store


## Import

DiagnosisStores can be imported using the `id`, e.g.

```
$ terraform import oci_cluster_health_diagnosis_store.test_diagnosis_store "id"
```
