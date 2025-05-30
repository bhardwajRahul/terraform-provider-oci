// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Visual Builder API
//
// Oracle Visual Builder enables developers to quickly build web and mobile applications. With a visual development environment that makes it easy to connect to Oracle data and third-party REST services, developers can build modern, consumer-grade applications in a fraction of the time it would take in other tools.
// The Visual Builder Instance Management API allows users to create and manage a Visual Builder instance.
//

package visualbuilder

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// UpdatePublicEndpointDetails Public endpoint configuration details (update).
type UpdatePublicEndpointDetails struct {

	// Source IP addresses or IP address ranges ingress rules. (ex: "168.122.59.5/32", "10.20.30.0/26")
	// An invalid IP or CIDR block will result in a 400 response.
	AllowlistedHttpIps []string `mandatory:"false" json:"allowlistedHttpIps"`

	// Virtual Cloud Networks allowed to access this network endpoint.
	AllowlistedHttpVcns []VirtualCloudNetwork `mandatory:"false" json:"allowlistedHttpVcns"`
}

func (m UpdatePublicEndpointDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdatePublicEndpointDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m UpdatePublicEndpointDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeUpdatePublicEndpointDetails UpdatePublicEndpointDetails
	s := struct {
		DiscriminatorParam string `json:"networkEndpointType"`
		MarshalTypeUpdatePublicEndpointDetails
	}{
		"PUBLIC",
		(MarshalTypeUpdatePublicEndpointDetails)(m),
	}

	return json.Marshal(&s)
}
