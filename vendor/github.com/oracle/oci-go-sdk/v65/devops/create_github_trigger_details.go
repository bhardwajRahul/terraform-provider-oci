// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// DevOps API
//
// Use the DevOps API to create DevOps projects, configure code repositories,  add artifacts to deploy, build and test software applications, configure  target deployment environments, and deploy software applications.  For more information, see DevOps (https://docs.oracle.com/iaas/Content/devops/using/home.htm).
//

package devops

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CreateGithubTriggerDetails The trigger for GitHub as the caller.
type CreateGithubTriggerDetails struct {

	// The OCID of the DevOps project to which the trigger belongs to.
	ProjectId *string `mandatory:"true" json:"projectId"`

	// The list of actions that are to be performed for this trigger.
	Actions []TriggerAction `mandatory:"true" json:"actions"`

	// Trigger display name. Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Optional description about the trigger.
	Description *string `mandatory:"false" json:"description"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.  See Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace. See Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// The OCID of the connection resource used to get details for triggered events.
	ConnectionId *string `mandatory:"false" json:"connectionId"`
}

// GetDisplayName returns DisplayName
func (m CreateGithubTriggerDetails) GetDisplayName() *string {
	return m.DisplayName
}

// GetDescription returns Description
func (m CreateGithubTriggerDetails) GetDescription() *string {
	return m.Description
}

// GetProjectId returns ProjectId
func (m CreateGithubTriggerDetails) GetProjectId() *string {
	return m.ProjectId
}

// GetActions returns Actions
func (m CreateGithubTriggerDetails) GetActions() []TriggerAction {
	return m.Actions
}

// GetFreeformTags returns FreeformTags
func (m CreateGithubTriggerDetails) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m CreateGithubTriggerDetails) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

func (m CreateGithubTriggerDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateGithubTriggerDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m CreateGithubTriggerDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCreateGithubTriggerDetails CreateGithubTriggerDetails
	s := struct {
		DiscriminatorParam string `json:"triggerSource"`
		MarshalTypeCreateGithubTriggerDetails
	}{
		"GITHUB",
		(MarshalTypeCreateGithubTriggerDetails)(m),
	}

	return json.Marshal(&s)
}

// UnmarshalJSON unmarshals from json
func (m *CreateGithubTriggerDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		DisplayName  *string                           `json:"displayName"`
		Description  *string                           `json:"description"`
		FreeformTags map[string]string                 `json:"freeformTags"`
		DefinedTags  map[string]map[string]interface{} `json:"definedTags"`
		ConnectionId *string                           `json:"connectionId"`
		ProjectId    *string                           `json:"projectId"`
		Actions      []triggeraction                   `json:"actions"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.DisplayName = model.DisplayName

	m.Description = model.Description

	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	m.ConnectionId = model.ConnectionId

	m.ProjectId = model.ProjectId

	m.Actions = make([]TriggerAction, len(model.Actions))
	for i, n := range model.Actions {
		nn, e = n.UnmarshalPolymorphicJSON(n.JsonData)
		if e != nil {
			return e
		}
		if nn != nil {
			m.Actions[i] = nn.(TriggerAction)
		} else {
			m.Actions[i] = nil
		}
	}
	return
}
