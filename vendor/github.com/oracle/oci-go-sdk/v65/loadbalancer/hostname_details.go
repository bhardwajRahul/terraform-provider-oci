// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Load Balancing API
//
// API for the Load Balancing service. Use this API to manage load balancers, backend sets, and related items. For more
// information, see Overview of Load Balancing (https://docs.oracle.com/iaas/Content/Balance/Concepts/balanceoverview.htm).
//

package loadbalancer

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// HostnameDetails The details of a hostname resource associated with a load balancer.
type HostnameDetails struct {

	// The name of the hostname resource.
	// Example: `example_hostname_001`
	Name *string `mandatory:"true" json:"name"`

	// A virtual hostname. For more information about virtual hostname string construction, see
	// Managing Request Routing (https://docs.oracle.com/iaas/Content/Balance/Tasks/managingrequest.htm#routing).
	// Example: `app.example.com`
	Hostname *string `mandatory:"true" json:"hostname"`
}

func (m HostnameDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m HostnameDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
