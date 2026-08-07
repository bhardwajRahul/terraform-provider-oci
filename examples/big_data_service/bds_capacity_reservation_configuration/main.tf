// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "bds_instance_id" {}
variable "bds_capacity_reservation_id" {}
variable "bds_capacity_reservation_configuration_display_name" {}


data "oci_bds_bds_instance" "test_bds_instance" {
  bds_instance_id = var.bds_instance_id
}

resource "oci_bds_bds_instance_bds_capacity_reservation_configuration" "test_bds_capacity_reservation_configuration" {

  bds_instance_id             = var.bds_instance_id
  bds_capacity_reservation_id = var.bds_capacity_reservation_id
  display_name                = var.bds_capacity_reservation_configuration_display_name
}

data "oci_bds_bds_instance_bds_capacity_reservation_configurations" "test_bds_capacity_reservation_configurations" {
  bds_instance_id = var.bds_instance_id

  depends_on = [oci_bds_bds_instance_bds_capacity_reservation_configuration.test_bds_capacity_reservation_configuration]
}

output "bds_instance" {
  value     = data.oci_bds_bds_instance.test_bds_instance
  sensitive = true
}

output "bds_capacity_reservation_configuration_id" {
  value = oci_bds_bds_instance_bds_capacity_reservation_configuration.test_bds_capacity_reservation_configuration.id
}

output "bds_capacity_reservation_configurations" {
  value     = data.oci_bds_bds_instance_bds_capacity_reservation_configurations.test_bds_capacity_reservation_configurations
  sensitive = true
}

