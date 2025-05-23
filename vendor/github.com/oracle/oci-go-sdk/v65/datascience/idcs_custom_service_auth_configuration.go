// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Science API
//
// Use the Data Science API to organize your data science work, access data and computing resources, and build, train, deploy and manage models and model deployments. For more information, see Data Science (https://docs.oracle.com/iaas/data-science/using/data-science.htm).
//

package datascience

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// IdcsCustomServiceAuthConfiguration Using existing Oracle Cloud Service for online prediction AuthN/Z
type IdcsCustomServiceAuthConfiguration struct {

	// URI of IDCS Stripe
	IdcsStripeUri *string `mandatory:"true" json:"idcsStripeUri"`

	// Audience of the IDCS application
	Audience *string `mandatory:"true" json:"audience"`

	// Scope of the IDCS application
	Scope *string `mandatory:"true" json:"scope"`

	// Name of the IDCS application
	ApplicationName *string `mandatory:"true" json:"applicationName"`

	// Name of the IDCS application role
	RoleName *string `mandatory:"true" json:"roleName"`
}

func (m IdcsCustomServiceAuthConfiguration) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m IdcsCustomServiceAuthConfiguration) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m IdcsCustomServiceAuthConfiguration) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeIdcsCustomServiceAuthConfiguration IdcsCustomServiceAuthConfiguration
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeIdcsCustomServiceAuthConfiguration
	}{
		"IDCS_CUSTOM_SERVICE",
		(MarshalTypeIdcsCustomServiceAuthConfiguration)(m),
	}

	return json.Marshal(&s)
}
