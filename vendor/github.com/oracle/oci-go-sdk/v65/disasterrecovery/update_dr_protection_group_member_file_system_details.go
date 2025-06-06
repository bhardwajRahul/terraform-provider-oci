// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Full Stack Disaster Recovery API
//
// Use the Full Stack Disaster Recovery (DR) API to manage disaster recovery for business applications.
// Full Stack DR is an OCI disaster recovery orchestration and management service that provides comprehensive disaster
// recovery capabilities for all layers of an application stack, including infrastructure, middleware, database,
// and application.
//

package disasterrecovery

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// UpdateDrProtectionGroupMemberFileSystemDetails Update properties for a file system member.
type UpdateDrProtectionGroupMemberFileSystemDetails struct {

	// The OCID of the member.
	// Example: `ocid1.database.oc1..uniqueID`
	MemberId *string `mandatory:"true" json:"memberId"`

	// The availability domain of the destination mount target.
	// Example: `BBTh:region-AD`
	DestinationAvailabilityDomain *string `mandatory:"false" json:"destinationAvailabilityDomain"`

	// A list of mappings between file system exports in the primary region and mount targets in the standby region.
	ExportMappings []FileSystemExportMappingDetails `mandatory:"false" json:"exportMappings"`

	DestinationEncryptionKey *UpdateVaultAndEncryptionKeyDetails `mandatory:"false" json:"destinationEncryptionKey"`

	// The OCID of the snapshot policy to use in the destination region. This policy will be attached to the file system after it moves to the destination region.
	// Example: `ocid1.filesystemsnapshotpolicy.oc1..uniqueID`
	DestinationSnapshotPolicyId *string `mandatory:"false" json:"destinationSnapshotPolicyId"`
}

// GetMemberId returns MemberId
func (m UpdateDrProtectionGroupMemberFileSystemDetails) GetMemberId() *string {
	return m.MemberId
}

func (m UpdateDrProtectionGroupMemberFileSystemDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateDrProtectionGroupMemberFileSystemDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m UpdateDrProtectionGroupMemberFileSystemDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeUpdateDrProtectionGroupMemberFileSystemDetails UpdateDrProtectionGroupMemberFileSystemDetails
	s := struct {
		DiscriminatorParam string `json:"memberType"`
		MarshalTypeUpdateDrProtectionGroupMemberFileSystemDetails
	}{
		"FILE_SYSTEM",
		(MarshalTypeUpdateDrProtectionGroupMemberFileSystemDetails)(m),
	}

	return json.Marshal(&s)
}
