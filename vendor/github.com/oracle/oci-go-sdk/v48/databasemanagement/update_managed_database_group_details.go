// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Management API
//
// Use the Database Management API to perform tasks such as obtaining performance and resource usage metrics
// for a fleet of Managed Databases or a specific Managed Database, creating Managed Database Groups, and
// running a SQL job on a Managed Database or Managed Database Group.
//

package databasemanagement

import (
	"github.com/oracle/oci-go-sdk/v48/common"
)

// UpdateManagedDatabaseGroupDetails The details required to update a Managed Database Group.
type UpdateManagedDatabaseGroupDetails struct {

	// The information specified by the user about the Managed Database Group.
	Description *string `mandatory:"false" json:"description"`
}

func (m UpdateManagedDatabaseGroupDetails) String() string {
	return common.PointerString(m)
}