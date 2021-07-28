// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"testing"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

// issue-routing-tag: load_balancer/default
func TestAccDatasourceLoadBalancerPolicies_basic(t *testing.T) {
	httpreplay.SetScenario("TestAccDatasourceLoadBalancerPolicies_basic")
	defer httpreplay.SaveScenario()
	providers := testAccProviders
	config := legacyTestProviderConfig() + `
	data "oci_load_balancer_policies" "t" {
		compartment_id = "${var.compartment_id}"
	}`

	resourceName := "data.oci_load_balancer_policies.t"

	resource.Test(t, resource.TestCase{
		PreCheck:                  func() { testAccPreCheck(t) },
		PreventPostDestroyRefresh: true,
		Providers:                 providers,
		Steps: []resource.TestStep{
			{
				Config: config,
				Check: ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(resourceName, "compartment_id"),
					resource.TestCheckResourceAttrSet(resourceName, "policies.#"),
					resource.TestCheckResourceAttrSet(resourceName, "policies.0.name"),
				),
			},
		},
	})
}
