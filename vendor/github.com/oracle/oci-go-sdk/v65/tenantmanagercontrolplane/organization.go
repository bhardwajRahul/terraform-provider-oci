// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Organizations API
//
// Use the Organizations API to consolidate multiple OCI tenancies into an organization, and centrally manage your tenancies and organization resources. For more information, see Organization Management Overview (https://docs.oracle.com/iaas/Content/General/Concepts/organization_management_overview.htm).
//

package tenantmanagercontrolplane

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// Organization An organization entity.
type Organization struct {

	// OCID of the organization.
	Id *string `mandatory:"true" json:"id"`

	// OCID of the compartment containing the organization. Always a tenancy OCID.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// OCID of the default Universal Credits Model subscription. Any tenancy joining the organization will automatically get assigned this subscription, if a subscription is not explictly assigned.
	DefaultUcmSubscriptionId *string `mandatory:"true" json:"defaultUcmSubscriptionId"`

	// Lifecycle state of the organization.
	LifecycleState OrganizationLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// Date and time when the organization was created.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// A display name for the organization. Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// The name of the tenancy that is the organization parent.
	ParentName *string `mandatory:"false" json:"parentName"`

	// Date and time when the organization was last updated.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`
}

func (m Organization) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m Organization) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingOrganizationLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetOrganizationLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// OrganizationLifecycleStateEnum Enum with underlying type: string
type OrganizationLifecycleStateEnum string

// Set of constants representing the allowable values for OrganizationLifecycleStateEnum
const (
	OrganizationLifecycleStateCreating OrganizationLifecycleStateEnum = "CREATING"
	OrganizationLifecycleStateActive   OrganizationLifecycleStateEnum = "ACTIVE"
	OrganizationLifecycleStateUpdating OrganizationLifecycleStateEnum = "UPDATING"
	OrganizationLifecycleStateDeleting OrganizationLifecycleStateEnum = "DELETING"
	OrganizationLifecycleStateDeleted  OrganizationLifecycleStateEnum = "DELETED"
	OrganizationLifecycleStateFailed   OrganizationLifecycleStateEnum = "FAILED"
)

var mappingOrganizationLifecycleStateEnum = map[string]OrganizationLifecycleStateEnum{
	"CREATING": OrganizationLifecycleStateCreating,
	"ACTIVE":   OrganizationLifecycleStateActive,
	"UPDATING": OrganizationLifecycleStateUpdating,
	"DELETING": OrganizationLifecycleStateDeleting,
	"DELETED":  OrganizationLifecycleStateDeleted,
	"FAILED":   OrganizationLifecycleStateFailed,
}

var mappingOrganizationLifecycleStateEnumLowerCase = map[string]OrganizationLifecycleStateEnum{
	"creating": OrganizationLifecycleStateCreating,
	"active":   OrganizationLifecycleStateActive,
	"updating": OrganizationLifecycleStateUpdating,
	"deleting": OrganizationLifecycleStateDeleting,
	"deleted":  OrganizationLifecycleStateDeleted,
	"failed":   OrganizationLifecycleStateFailed,
}

// GetOrganizationLifecycleStateEnumValues Enumerates the set of values for OrganizationLifecycleStateEnum
func GetOrganizationLifecycleStateEnumValues() []OrganizationLifecycleStateEnum {
	values := make([]OrganizationLifecycleStateEnum, 0)
	for _, v := range mappingOrganizationLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetOrganizationLifecycleStateEnumStringValues Enumerates the set of values in String for OrganizationLifecycleStateEnum
func GetOrganizationLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"ACTIVE",
		"UPDATING",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingOrganizationLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOrganizationLifecycleStateEnum(val string) (OrganizationLifecycleStateEnum, bool) {
	enum, ok := mappingOrganizationLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
