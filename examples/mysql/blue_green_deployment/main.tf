// Copyright (c) 2017, 2026, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {
  type = string
}

variable "user_ocid" {
  type = string
}

variable "fingerprint" {
  type = string
}

variable "private_key_path" {
  type = string
}

variable "region" {
  type = string
}

variable "compartment_ocid" {
  type = string
}

variable "source_mysql_version" {
  description = "MySQL version for the source DB system."
  type        = string
  default     = "8.4.10"
}

variable "target_mysql_version" {
  description = "MySQL version for the target DB system."
  type        = string
  default     = "8.4.10"
}

variable "mysql_admin_username" {
  description = "Administrator username for the source DB system."
  type        = string
  default     = "adminUser"
}

variable "mysql_admin_password" {
  description = "Administrator password for the source DB system and Blue/Green replication channel. Override with TF_VAR_mysql_admin_password."
  type        = string
  default     = "BEstrO0ng_#11"
}

variable "mysql_shape_name" {
  description = "Shape used by the source DB system. The target inherits this shape."
  type        = string
  default     = "MySQL.2"
}

variable "mysql_data_storage_size_in_gb" {
  description = "Storage size used by the source DB system. The target inherits this size."
  type        = number
  default     = 50
}

variable "switchover_trigger" {
  description = "Leave null during creation. Set to 1 after the deployment reaches READY_FOR_SWITCHOVER."
  type        = number
  default     = null
}

provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

resource "oci_core_vcn" "test_vcn" {
  cidr_block     = "10.0.0.0/16"
  compartment_id = var.compartment_ocid
  display_name   = "blue-green-example-vcn"
}

resource "oci_core_security_list" "mysql_security_list" {
  compartment_id = var.compartment_ocid
  display_name   = "blue-green-example-mysql-security-list"
  vcn_id         = oci_core_vcn.test_vcn.id

  egress_security_rules {
    destination = "0.0.0.0/0"
    protocol    = "all"
    stateless   = false
  }

  ingress_security_rules {
    description = "Allow MySQL replication traffic between Blue/Green DB systems"
    protocol    = "6"
    source      = oci_core_vcn.test_vcn.cidr_block
    stateless   = false

    tcp_options {
      max = 3306
      min = 3306
    }
  }
}

resource "oci_core_subnet" "test_subnet" {
  cidr_block        = "10.0.0.0/24"
  compartment_id    = var.compartment_ocid
  display_name      = "blue-green-example-subnet"
  security_list_ids = [oci_core_security_list.mysql_security_list.id]
  vcn_id            = oci_core_vcn.test_vcn.id
}

data "oci_identity_availability_domains" "test_availability_domains" {
  compartment_id = var.tenancy_ocid
}

data "oci_mysql_mysql_configurations" "test_mysql_configurations" {
  compartment_id = var.compartment_ocid
  shape_name     = var.mysql_shape_name
  state          = "ACTIVE"
}

resource "oci_mysql_mysql_db_system" "test_mysql_db_system" {
  admin_password          = var.mysql_admin_password
  admin_username          = var.mysql_admin_username
  availability_domain     = data.oci_identity_availability_domains.test_availability_domains.availability_domains[0].name
  compartment_id          = var.compartment_ocid
  configuration_id        = data.oci_mysql_mysql_configurations.test_mysql_configurations.configurations[0].id
  data_storage_size_in_gb = var.mysql_data_storage_size_in_gb
  display_name            = "blue-green-example-source"
  mysql_version           = var.source_mysql_version
  shape_name              = var.mysql_shape_name
  subnet_id               = oci_core_subnet.test_subnet.id

  backup_policy {
    is_enabled = false
  }
}

resource "oci_mysql_blue_green_deployment" "test_blue_green_deployment" {
  compartment_id                    = var.compartment_ocid
  delete_target_db_system_on_delete = true
  display_name                      = "blue-green-example"
  source_db_system_id               = oci_mysql_mysql_db_system.test_mysql_db_system.id

  channel_details {
    source_password = var.mysql_admin_password
    source_username = var.mysql_admin_username
    ssl_mode        = "REQUIRED"
  }

  target_db_system_details {
    mysql_version = var.target_mysql_version
  }

  # Keep this null on the initial apply. After switchover_status reaches
  # READY_FOR_SWITCHOVER, set it to 1 and apply again to initiate switchover.
  switchover_trigger = var.switchover_trigger

  freeform_tags = {
    Example = "blue-green-deployment"
  }

  timeouts {
    create = "2h"
    update = "2h"
    delete = "2h"
  }
}

data "oci_mysql_blue_green_deployment" "test_blue_green_deployment" {
  blue_green_deployment_id = oci_mysql_blue_green_deployment.test_blue_green_deployment.id
}

output "blue_green_deployment_id" {
  value = oci_mysql_blue_green_deployment.test_blue_green_deployment.id
}

output "source_db_system_id" {
  value = oci_mysql_blue_green_deployment.test_blue_green_deployment.source_db_system_id
}

output "target_db_system_id" {
  value = oci_mysql_blue_green_deployment.test_blue_green_deployment.target_db_system_id
}

output "active_db_system_id" {
  value = oci_mysql_blue_green_deployment.test_blue_green_deployment.active_db_system_id
}

output "switchover_status" {
  value = data.oci_mysql_blue_green_deployment.test_blue_green_deployment.switchover_status
}
