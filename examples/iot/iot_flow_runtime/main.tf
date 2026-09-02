// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_id" {}
variable "iot_domain_ocid" {}
variable "log_group_id" {}
variable "log_id" {}
variable "subnet_id" {}
variable "export_id" {}
variable "mount_target_id" {}

variable "iot_flow_runtime_defined_tags_value" {
  default = "value"
}

variable "iot_flow_runtime_description" {
  default = "description"
}

variable "iot_flow_runtime_display_name" {
  default = "displayName"
}

variable "iot_flow_runtime_freeform_tags" {
  default = { "Department" = "Finance" }
}

variable "iot_flow_runtime_id" {
  default = "id"
}

variable "iot_flow_runtime_network_config_file_storage_mounts_mount_path" {
  default = "mountPath"
}

variable "iot_flow_runtime_network_config_network_security_group_ids" {
  default = []
}

variable "iot_flow_runtime_scale" {
  default = "MEDIUM"
}

variable "iot_flow_runtime_state" {
  default = "ACTIVE"
}



provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

resource "oci_iot_iot_flow_runtime" "test_iot_flow_runtime" {
  #Required
  compartment_id = var.compartment_id
  iot_domain_id  = var.iot_domain_ocid

  #Optional
  #defined_tags  = tomap({ oci_identity_tag_namespace.tag-namespace1.name.oci_identity_tag.tag1.name = var.iot_flow_runtime_defined_tags_value })
  description   = var.iot_flow_runtime_description
  display_name  = var.iot_flow_runtime_display_name
  freeform_tags = var.iot_flow_runtime_freeform_tags
  log_config {
    #Required
    log_group_id = var.log_group_id

    #Optional
    log_id = var.log_id
  }
  network_config {
    #Required
    subnet_id = var.subnet_id

    #Optional
    file_storage_mounts {
      #Required
      export_id       = var.export_id
      mount_path      = var.iot_flow_runtime_network_config_file_storage_mounts_mount_path
      mount_target_id = var.mount_target_id
    }
    network_security_group_ids = var.iot_flow_runtime_network_config_network_security_group_ids
  }
  scale = var.iot_flow_runtime_scale
}

data "oci_iot_iot_flow_runtimes" "test_iot_flow_runtimes" {
  #Required
  compartment_id = var.compartment_id

  #Optional
  display_name  = var.iot_flow_runtime_display_name
  id            = oci_iot_iot_flow_runtime.test_iot_flow_runtime.id
  iot_domain_id = var.iot_domain_ocid
  state         = var.iot_flow_runtime_state
}

