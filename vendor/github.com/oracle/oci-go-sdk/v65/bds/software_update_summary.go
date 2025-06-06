// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Big Data Service API
//
// REST API for Oracle Big Data Service. Use this API to build, deploy, and manage fully elastic Big Data Service clusters. Build on Hadoop, Spark and Data Science distributions, which can be fully integrated with existing enterprise data in Oracle Database and Oracle applications.
//

package bds

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// SoftwareUpdateSummary The software update that is currently available for the cluster.
type SoftwareUpdateSummary struct {

	// Unique identifier of a given software update
	SoftwareUpdateKey *string `mandatory:"true" json:"softwareUpdateKey"`

	// The version of the software update.
	SoftwareUpdateVersion *string `mandatory:"true" json:"softwareUpdateVersion"`

	// The time when the software update was released.
	TimeReleased *common.SDKTime `mandatory:"true" json:"timeReleased"`

	// Type of current software update.
	SoftwareUpdateType SoftwareUpdateSoftwareUpdateTypeEnum `mandatory:"true" json:"softwareUpdateType"`

	// The lifecycle state of the software update.
	LifecycleState SoftwareUpdateLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`
}

func (m SoftwareUpdateSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m SoftwareUpdateSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingSoftwareUpdateSoftwareUpdateTypeEnum(string(m.SoftwareUpdateType)); !ok && m.SoftwareUpdateType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SoftwareUpdateType: %s. Supported values are: %s.", m.SoftwareUpdateType, strings.Join(GetSoftwareUpdateSoftwareUpdateTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingSoftwareUpdateLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetSoftwareUpdateLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
