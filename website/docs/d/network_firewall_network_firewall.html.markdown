---
subcategory: "Network Firewall"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_network_firewall_network_firewall"
sidebar_current: "docs-oci-datasource-network_firewall-network_firewall"
description: |-
  Provides details about a specific Network Firewall in Oracle Cloud Infrastructure Network Firewall service
---

# Data Source: oci_network_firewall_network_firewall
This data source provides details about a specific Network Firewall resource in Oracle Cloud Infrastructure Network Firewall service.

Gets a NetworkFirewall by identifier

## Example Usage

```hcl
data "oci_network_firewall_network_firewall" "test_network_firewall" {
	#Required
	network_firewall_id = oci_network_firewall_network_firewall.test_network_firewall.id
}
```

## Argument Reference

The following arguments are supported:

* `network_firewall_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Network Firewall resource.


## Attributes Reference

The following attributes are exported:

* `availability_domain` - Availability Domain where Network Firewall instance is created. To get a list of availability domains for a tenancy, use the [ListAvailabilityDomains](https://docs.cloud.oracle.com/iaas/api/#/en/identity/20160918/AvailabilityDomain/ListAvailabilityDomains) operation. Example: `kIdk:PHX-AD-1` 
* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment containing the Network Firewall.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - A user-friendly name for the Network Firewall. Does not have to be unique, and it's changeable. Avoid entering confidential information.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Network Firewall resource.
* `ipv4address` - IPv4 address for the Network Firewall.
* `ipv6address` - IPv6 address for the Network Firewall.
* `lifecycle_details` - A message describing the current state in more detail. For example, it can be used to provide actionable information for a resource in 'FAILED' state.
* `nat_configuration` - Nat Configuration response.
	* `must_enable_private_nat` - To allocate private NAT IPs to the firewall. The attached network firewall policy must also have NAT rules to enable NAT on any traffic passing through the firewall.
	* `nat_ip_address_list` - An array of NAT IP addresses that are associated with the Network Firewall. These IPs are reserved for NAT and shouldn't be used for any other purpose in the subnet.
* `network_firewall_policy_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Network Firewall Policy.
* `network_security_group_ids` - An array of network security groups [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) associated with the Network Firewall.
* `state` - The current state of the Network Firewall.
* `subnet_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the subnet associated with the Network Firewall.
* `system_tags` - Usage of system tag keys. These predefined keys are scoped to namespaces. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The time at which the Network Firewall was created in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339). Example: `2016-08-25T21:10:29.600Z` 
* `time_updated` - The time at which the Network Firewall was updated in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339). Example: `2016-08-25T21:10:29.600Z` 

