// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Network Monitoring API
//
// Use the Network Monitoring API to troubleshoot routing and security issues for resources such as virtual cloud networks (VCNs) and compute instances. For more information, see the console
// documentation for the Network Path Analyzer (https://docs.oracle.com/iaas/Content/Network/Concepts/path_analyzer.htm) tool.
//

package vnmonitoring

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// StatefulNsgConfiguration Defines the stateful network security group configuration that allowed the traffic.
type StatefulNsgConfiguration struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the network
	// security group that allowed the traffic.
	NsgId *string `mandatory:"true" json:"nsgId"`

	SecurityRule *SecurityRule `mandatory:"true" json:"securityRule"`
}

func (m StatefulNsgConfiguration) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m StatefulNsgConfiguration) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m StatefulNsgConfiguration) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeStatefulNsgConfiguration StatefulNsgConfiguration
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeStatefulNsgConfiguration
	}{
		"STATEFUL_NSG",
		(MarshalTypeStatefulNsgConfiguration)(m),
	}

	return json.Marshal(&s)
}
