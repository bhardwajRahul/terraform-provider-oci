// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "iot_flow_runtime_id" {}

variable "iot_flow_runtime_flow_flows_document" {
  default = {"flows":[{"id":"simple-flow-tab","type":"tab","label":"Simple Flow"}]}
}

provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

resource "oci_iot_iot_flow_runtime_flow" "test_iot_flow_runtime_flow" {
  #Required
  flows_document      = jsonencode(var.iot_flow_runtime_flow_flows_document)
  iot_flow_runtime_id = var.iot_flow_runtime_id
}

data "oci_iot_iot_flow_runtime_flow" "test_iot_flow_runtime_flow" {
  #Required
  iot_flow_runtime_id = var.iot_flow_runtime_id
}

