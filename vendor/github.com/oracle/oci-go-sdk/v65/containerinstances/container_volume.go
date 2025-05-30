// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Container Instance API
//
// A description of the Container Instance API
//

package containerinstances

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ContainerVolume A volume represents a directory with data that is accessible across multiple containers in a
// container instance.
type ContainerVolume interface {

	// The name of the volume. This must be unique within a single container instance.
	GetName() *string
}

type containervolume struct {
	JsonData   []byte
	Name       *string `mandatory:"true" json:"name"`
	VolumeType string  `json:"volumeType"`
}

// UnmarshalJSON unmarshals json
func (m *containervolume) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalercontainervolume containervolume
	s := struct {
		Model Unmarshalercontainervolume
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Name = s.Model.Name
	m.VolumeType = s.Model.VolumeType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *containervolume) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.VolumeType {
	case "EMPTYDIR":
		mm := ContainerEmptyDirVolume{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "CONFIGFILE":
		mm := ContainerConfigFileVolume{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for ContainerVolume: %s.", m.VolumeType)
		return *m, nil
	}
}

// GetName returns Name
func (m containervolume) GetName() *string {
	return m.Name
}

func (m containervolume) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m containervolume) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
