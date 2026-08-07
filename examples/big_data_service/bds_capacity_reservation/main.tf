// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0
variable "compartment_id" {}
variable "bds_capacity_reservation_display_name" {}
variable "domain1reservation_id" {}
variable "domain2reservation_id" {}
variable "domain3reservation_id" {}
variable "bds_capacity_reservation_state" {}
# Optional
# variable "bds_capacity_reservation_freeform_tags" {
#   default = {
#     purpose = "bds-capacity-reservation"
#   }
# }

resource "oci_bds_bds_capacity_reservation" "test_bds_capacity_reservation" {
  compartment_id = var.compartment_id
  display_name   = var.bds_capacity_reservation_display_name
  # freeform_tags  = var.bds_capacity_reservation_freeform_tags

  compute_capacity_reservations {
    domain1reservation_id = var.domain1reservation_id
    domain2reservation_id = var.domain2reservation_id
    domain3reservation_id = var.domain3reservation_id
  }
}

data "oci_bds_bds_capacity_reservation" "test_bds_capacity_reservation" {
  bds_capacity_reservation_id = oci_bds_bds_capacity_reservation.test_bds_capacity_reservation.id
}

data "oci_bds_bds_capacity_reservations" "test_bds_capacity_reservations" {
  compartment_id = var.compartment_id
  display_name   = var.bds_capacity_reservation_display_name
  state          = var.bds_capacity_reservation_state

  depends_on = [oci_bds_bds_capacity_reservation.test_bds_capacity_reservation]
}

output "bds_capacity_reservation_id" {
  value = oci_bds_bds_capacity_reservation.test_bds_capacity_reservation.id
}

output "bds_capacity_reservation" {
  value = data.oci_bds_bds_capacity_reservation.test_bds_capacity_reservation
}

output "bds_capacity_reservations" {
  value = data.oci_bds_bds_capacity_reservations.test_bds_capacity_reservations
}
