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

// UpdateLoadBalancerShapeDetails The representation of UpdateLoadBalancerShapeDetails
type UpdateLoadBalancerShapeDetails struct {

	// The new shape name for the load balancer.
	// Allowed values are :
	//   *  10Mbps
	//   *  100Mbps
	//   *  400Mbps
	//   *  8000Mbps
	//   *  Flexible
	//   Example: `flexible`
	//   * NOTE: Fixed shapes 10Mbps, 100Mbps, 400Mbps, 8000Mbps will be deprecated after May 2023. This api
	//   * will only support `Flexible` shape after that date.
	ShapeName *string `mandatory:"true" json:"shapeName"`

	// The configuration details to update load balancer to a different profile.
	ShapeDetails *ShapeDetails `mandatory:"false" json:"shapeDetails"`
}

func (m UpdateLoadBalancerShapeDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateLoadBalancerShapeDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
