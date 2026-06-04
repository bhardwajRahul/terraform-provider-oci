---
subcategory: "Cluster Health"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_cluster_health_diagnosis_stores"
sidebar_current: "docs-oci-datasource-cluster_health-diagnosis_stores"
description: |-
  Provides the list of Diagnosis Stores in Oracle Cloud Infrastructure Cluster Health service
---

# Data Source: oci_cluster_health_diagnosis_stores
This data source provides the list of Diagnosis Stores in Oracle Cloud Infrastructure Cluster Health service.

Gets a list of Diagnosis Stores.


## Example Usage

```hcl
data "oci_cluster_health_diagnosis_stores" "test_diagnosis_stores" {

	#Optional
	compartment_id = var.compartment_id
	diagnosis_store_id = oci_cluster_health_diagnosis_store.test_diagnosis_store.id
	display_name = var.diagnosis_store_display_name
	state = var.diagnosis_store_state
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment in which to list resources.
* `diagnosis_store_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Diagnosis.
* `display_name` - (Optional) A filter to return only resources that match the given display name exactly.
* `state` - (Optional) A filter to return only resources that match the given lifecycle state. The state value is case-insensitive.


## Attributes Reference

The following attributes are exported:

* `diagnosis_store_collection` - The list of diagnosis_store_collection.

### DiagnosisStore Reference

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
