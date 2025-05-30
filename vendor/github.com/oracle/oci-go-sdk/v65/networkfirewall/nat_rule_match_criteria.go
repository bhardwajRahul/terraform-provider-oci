// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Network Firewall API
//
// Use the Network Firewall API to create network firewalls and configure policies that regulates network traffic in and across VCNs.
//

package networkfirewall

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// NatRuleMatchCriteria Match criteria used in NAT Rule used on the firewall policy.
type NatRuleMatchCriteria struct {

	// An array of IP address list names to be evaluated against the traffic source address.
	SourceAddress []string `mandatory:"false" json:"sourceAddress"`

	// An array of IP address list names to be evaluated against the traffic destination address.
	DestinationAddress []string `mandatory:"false" json:"destinationAddress"`

	// A Service name to be evaluated against the traffic protocol and protocol-specific parameters.
	Service *string `mandatory:"false" json:"service"`
}

func (m NatRuleMatchCriteria) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m NatRuleMatchCriteria) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
