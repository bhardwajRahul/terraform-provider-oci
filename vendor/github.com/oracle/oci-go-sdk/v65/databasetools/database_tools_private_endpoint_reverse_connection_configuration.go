// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Tools
//
// Use the Database Tools API to manage connections, private endpoints, and work requests in the Database Tools service.
//

package databasetools

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// DatabaseToolsPrivateEndpointReverseConnectionConfiguration Reverse connection configuration details of the private endpoint.
type DatabaseToolsPrivateEndpointReverseConnectionConfiguration struct {

	// A list of IP addresses in the customer VCN to be used as the source IPs for reverse connection packets
	// traveling from the service's VCN to the customer's VCN.
	ReverseConnectionsSourceIps []DatabaseToolsPrivateEndpointReverseConnectionsSourceIp `mandatory:"false" json:"reverseConnectionsSourceIps"`
}

func (m DatabaseToolsPrivateEndpointReverseConnectionConfiguration) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DatabaseToolsPrivateEndpointReverseConnectionConfiguration) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
