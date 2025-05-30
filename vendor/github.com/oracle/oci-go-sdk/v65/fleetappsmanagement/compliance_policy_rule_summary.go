// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Fleet Application Management Service API
//
// Fleet Application Management provides a centralized platform to help you automate resource management tasks, validate patch compliance, and enhance operational efficiency across an enterprise.
//

package fleetappsmanagement

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CompliancePolicyRuleSummary Summary information about a CompliancePolicyRule.
type CompliancePolicyRuleSummary struct {

	// Unique OCID of the CompliancePolicyRule.
	Id *string `mandatory:"true" json:"id"`

	// A user-friendly name. Does not have to be unique, and it's changeable.
	// Avoid entering confidential information.
	// Example: `My new resource`
	DisplayName *string `mandatory:"true" json:"displayName"`

	// Unique OCID of the CompliancePolicy.
	CompliancePolicyId *string `mandatory:"true" json:"compliancePolicyId"`

	ProductVersion *ProductVersionDetails `mandatory:"true" json:"productVersion"`

	// PlatformConfiguration OCID for the patch type to which this CompliancePolicyRule applies.
	PatchTypeId []string `mandatory:"true" json:"patchTypeId"`

	PatchSelection PatchSelectionDetails `mandatory:"true" json:"patchSelection"`

	// The OCID of the compartment the CompliancePolicyRule belongs to.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The date and time the CompliancePolicyRule was created, in the format defined by RFC 3339 (https://tools.ietf.org/html/rfc3339).
	// Example: `2016-08-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The current state of the CompliancePolicyRule.
	LifecycleState CompliancePolicyRuleLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// Severity to which this CompliancePolicyRule applies.
	Severity []ComplianceRuleSeverityEnum `mandatory:"false" json:"severity,omitempty"`

	// Grace period in days,weeks,months or years the exemption is applicable for the rule.
	// This enables a grace period when Fleet Application Management doesn't report the product as noncompliant when patch is not applied.
	GracePeriod *string `mandatory:"false" json:"gracePeriod"`

	// The date and time the CompliancePolicyRule was updated, in the format defined by RFC 3339 (https://tools.ietf.org/html/rfc3339).
	// Example: `2016-08-25T21:10:29.600Z`
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// A message that describes the current state of the CompliancePolicyRule in more detail. For example,
	// can be used to provide actionable information for a resource in the Failed state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// System tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m CompliancePolicyRuleSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CompliancePolicyRuleSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingCompliancePolicyRuleLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetCompliancePolicyRuleLifecycleStateEnumStringValues(), ",")))
	}

	for _, val := range m.Severity {
		if _, ok := GetMappingComplianceRuleSeverityEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Severity: %s. Supported values are: %s.", val, strings.Join(GetComplianceRuleSeverityEnumStringValues(), ",")))
		}
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *CompliancePolicyRuleSummary) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Severity           []ComplianceRuleSeverityEnum           `json:"severity"`
		GracePeriod        *string                                `json:"gracePeriod"`
		TimeUpdated        *common.SDKTime                        `json:"timeUpdated"`
		LifecycleDetails   *string                                `json:"lifecycleDetails"`
		FreeformTags       map[string]string                      `json:"freeformTags"`
		DefinedTags        map[string]map[string]interface{}      `json:"definedTags"`
		SystemTags         map[string]map[string]interface{}      `json:"systemTags"`
		Id                 *string                                `json:"id"`
		DisplayName        *string                                `json:"displayName"`
		CompliancePolicyId *string                                `json:"compliancePolicyId"`
		ProductVersion     *ProductVersionDetails                 `json:"productVersion"`
		PatchTypeId        []string                               `json:"patchTypeId"`
		PatchSelection     patchselectiondetails                  `json:"patchSelection"`
		CompartmentId      *string                                `json:"compartmentId"`
		TimeCreated        *common.SDKTime                        `json:"timeCreated"`
		LifecycleState     CompliancePolicyRuleLifecycleStateEnum `json:"lifecycleState"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.Severity = make([]ComplianceRuleSeverityEnum, len(model.Severity))
	copy(m.Severity, model.Severity)
	m.GracePeriod = model.GracePeriod

	m.TimeUpdated = model.TimeUpdated

	m.LifecycleDetails = model.LifecycleDetails

	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	m.SystemTags = model.SystemTags

	m.Id = model.Id

	m.DisplayName = model.DisplayName

	m.CompliancePolicyId = model.CompliancePolicyId

	m.ProductVersion = model.ProductVersion

	m.PatchTypeId = make([]string, len(model.PatchTypeId))
	copy(m.PatchTypeId, model.PatchTypeId)
	nn, e = model.PatchSelection.UnmarshalPolymorphicJSON(model.PatchSelection.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.PatchSelection = nn.(PatchSelectionDetails)
	} else {
		m.PatchSelection = nil
	}

	m.CompartmentId = model.CompartmentId

	m.TimeCreated = model.TimeCreated

	m.LifecycleState = model.LifecycleState

	return
}
