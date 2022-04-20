---
subcategory: "Data Safe"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_data_safe_report_definition"
sidebar_current: "docs-oci-datasource-data_safe-report_definition"
description: |-
  Provides details about a specific Report Definition in Oracle Cloud Infrastructure Data Safe service
---

# Data Source: oci_data_safe_report_definition
This data source provides details about a specific Report Definition resource in Oracle Cloud Infrastructure Data Safe service.

Gets the details of report definition specified by the identifier

## Example Usage

```hcl
data "oci_data_safe_report_definition" "test_report_definition" {
	#Required
	report_definition_id = oci_data_safe_report_definition.test_report_definition.id
}
```

## Argument Reference

The following arguments are supported:

* `report_definition_id` - (Required) Unique report definition identifier


## Attributes Reference

The following attributes are exported:

* `category` - Specifies the name of the category that this report belongs to.
* `column_filters` - An array of column filter objects. A column Filter object stores all information about a column filter including field name, an operator, one or more expressions, if the filter is enabled, or if the filter is hidden.
	* `expressions` - An array of expressions based on the operator type. A filter may have one or more expressions.
	* `field_name` - Name of the column on which the filter must be applied.
	* `is_enabled` - Indicates if the filter is enabled. Values can either be 'true' or 'false'.
	* `is_hidden` - Indicates if the filter is hidden. Values can either be 'true' or 'false'.
	* `operator` - Specifies the type of operator that must be applied for example in, eq etc.
* `column_info` - An array of column objects in the order (left to right) displayed in the report. A column object stores all information about a column, including the name displayed on the UI, corresponding field name in the data source, data type of the column, and column visibility (if the column is visible to the user).
	* `data_type` - Specifies the data type of the column.
	* `display_name` - Name of the column displayed on UI.
	* `display_order` - Specifies the display order of the column.
	* `field_name` - Specifies the corresponding field name in the data source.
	* `is_hidden` - Indicates if the column is hidden. Values can either be 'true' or 'false'.
* `column_sortings` - An array of column sorting objects. Each column sorting object stores the column name to be sorted and if the sorting is in ascending order; sorting is done by the first column in the array, then by the second column in the array, etc.
	* `field_name` - Name of the column that must be sorted.
	* `is_ascending` - Indicates if the column must be sorted in ascending order. Values can either be 'true' or 'false'.
	* `sorting_order` - Indicates the order at which column must be sorted.
* `compartment_id` - The OCID of the compartment containing the report definition.
* `data_source` - Specifies the name of a resource that provides data for the report. For example alerts, events.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)  Example: `{"Operations.CostCenter": "42"}` 
* `description` - A description of the report definition.
* `display_name` - Name of the report definition.
* `display_order` - Specifies how the report definitions are ordered in the display.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)  Example: `{"Department": "Finance"}` 
* `id` - The OCID of the report definition.
* `is_seeded` - Signifies whether the definition is seeded or user defined. Values can either be 'true' or 'false'.
* `parent_id` - The OCID of the parent report definition. In the case of seeded report definition, this is same as definition OCID.
* `scim_filter` - Additional scim filters used to specialize the report.
* `state` - The current state of the report.
* `summary` - An array of report summary objects in the order (left to right)  displayed in the report.  A  report summary object stores all information about summary of report to be displayed, including the name displayed on UI, the display order, corresponding group by and count of values, summary visibility (if the summary is visible to user).
	* `count_of` - Name of the key or count of object.
	* `display_order` - Specifies the order in which the summary must be displayed.
	* `group_by_field_name` - A comma-delimited string that specifies the names of the fields by which the records must be aggregated to get the summary.
	* `is_hidden` - Indicates if the summary is hidden. Values can either be 'true' or 'false'.
	* `name` - Name of the report summary.
	* `scim_filter` - Additional scim filters used to get the specific summary.
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace. For more information, see Resource Tags. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - Specifies the time at which the report definition was created.
* `time_updated` - The date and time of the report definition update in Data Safe.
