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

// CreateDeploymentDetails The information about new deployment.
type CreateDeploymentDetails interface {

	// The OCID of a pipeline.
	GetDeployPipelineId() *string

	// Deployment display name. Avoid entering confidential information.
	GetDisplayName() *string

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.  See Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"bar-key": "value"}`
	GetFreeformTags() map[string]string

	// Defined tags for this resource. Each key is predefined and scoped to a namespace. See Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"foo-namespace": {"bar-key": "value"}}`
	GetDefinedTags() map[string]map[string]interface{}
}

type createdeploymentdetails struct {
	JsonData         []byte
	DisplayName      *string                           `mandatory:"false" json:"displayName"`
	FreeformTags     map[string]string                 `mandatory:"false" json:"freeformTags"`
	DefinedTags      map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
	DeployPipelineId *string                           `mandatory:"true" json:"deployPipelineId"`
	DeploymentType   string                            `json:"deploymentType"`
}

// UnmarshalJSON unmarshals json
func (m *createdeploymentdetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalercreatedeploymentdetails createdeploymentdetails
	s := struct {
		Model Unmarshalercreatedeploymentdetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.DeployPipelineId = s.Model.DeployPipelineId
	m.DisplayName = s.Model.DisplayName
	m.FreeformTags = s.Model.FreeformTags
	m.DefinedTags = s.Model.DefinedTags
	m.DeploymentType = s.Model.DeploymentType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *createdeploymentdetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.DeploymentType {
	case "PIPELINE_REDEPLOYMENT":
		mm := CreateDeployPipelineRedeploymentDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "PIPELINE_DEPLOYMENT":
		mm := CreateDeployPipelineDeploymentDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "SINGLE_STAGE_DEPLOYMENT":
		mm := CreateSingleDeployStageDeploymentDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "SINGLE_STAGE_REDEPLOYMENT":
		mm := CreateSingleDeployStageRedeploymentDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for CreateDeploymentDetails: %s.", m.DeploymentType)
		return *m, nil
	}
}

// GetDisplayName returns DisplayName
func (m createdeploymentdetails) GetDisplayName() *string {
	return m.DisplayName
}

// GetFreeformTags returns FreeformTags
func (m createdeploymentdetails) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m createdeploymentdetails) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

// GetDeployPipelineId returns DeployPipelineId
func (m createdeploymentdetails) GetDeployPipelineId() *string {
	return m.DeployPipelineId
}

func (m createdeploymentdetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m createdeploymentdetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
