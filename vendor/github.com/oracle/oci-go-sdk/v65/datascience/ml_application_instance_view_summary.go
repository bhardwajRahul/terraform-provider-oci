// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Science API
//
// Use the Data Science API to organize your data science work, access data and computing resources, and build, train, deploy and manage models and model deployments. For more information, see Data Science (https://docs.oracle.com/iaas/data-science/using/data-science.htm).
//

package datascience

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// MlApplicationInstanceViewSummary Summary of the MlApplicationInstanceView.
type MlApplicationInstanceViewSummary struct {

	// The OCID of the MlApplicationInstanceView. Unique identifier that is immutable after creation
	Id *string `mandatory:"true" json:"id"`

	// The name of the MlApplicationInstance (created by the consumer) which this MlApplicationInstanceView is mirroring.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The OCID of ML Application. This resource is an instance of ML Application referenced by this OCID.
	MlApplicationId *string `mandatory:"true" json:"mlApplicationId"`

	// The name of ML Application (based on mlApplicationId).
	MlApplicationName *string `mandatory:"true" json:"mlApplicationName"`

	// The OCID of the MlApplicationInstance (created by the consumer) which this MlApplicationInstanceView is mirroring.
	MlApplicationInstanceId *string `mandatory:"true" json:"mlApplicationInstanceId"`

	// This field is a copy from MlApplicationInstance created by the consumer. The OCID of ML Application Implementation selected as a certain solution for a given ML problem (ML Application)
	MlApplicationImplementationId *string `mandatory:"true" json:"mlApplicationImplementationId"`

	// This field is a copy from MlApplicationInstance created by the consumer. The name of Ml Application Implemenation (based on mlApplicationImplementationId)
	MlApplicationImplementationName *string `mandatory:"true" json:"mlApplicationImplementationName"`

	// Version of MlApplication package which is currently used by this MlApplicationInstance.
	PackageVersion *string `mandatory:"true" json:"packageVersion"`

	// The OCID of the currently used MlApplicationImplementationVersion
	MlApplicationImplementationVersionId *string `mandatory:"true" json:"mlApplicationImplementationVersionId"`

	// This field is a copy from MlApplicationInstance created by the consumer. States whether the MlApplicationInstance is supposed to be in ACTIVE lifecycle state.
	IsEnabled *bool `mandatory:"true" json:"isEnabled"`

	// The OCID of the compartment where the MlApplicationInstanceView is created.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The time the the MlApplicationInstanceView was created. An RFC3339 formatted datetime string
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The current state of the MlApplicationInstance(View).
	LifecycleState MlApplicationInstanceViewLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The current substate of the MlApplicationInstance. The substate has MlApplicationInstance specific values in comparison with lifecycleState which has standard values common for all OCI resources.
	LifecycleSubstate MlApplicationInstanceViewLifecycleSubstateEnum `mandatory:"true" json:"lifecycleSubstate"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. See Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace. See Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Usage of system tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m MlApplicationInstanceViewSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m MlApplicationInstanceViewSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingMlApplicationInstanceViewLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetMlApplicationInstanceViewLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingMlApplicationInstanceViewLifecycleSubstateEnum(string(m.LifecycleSubstate)); !ok && m.LifecycleSubstate != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleSubstate: %s. Supported values are: %s.", m.LifecycleSubstate, strings.Join(GetMlApplicationInstanceViewLifecycleSubstateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
