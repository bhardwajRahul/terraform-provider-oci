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
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// BlockVolumeAttachmentDetails Deprecated. Use the 'ComputeInstanceNonMovableBlockVolumeAttachOperationDetails' definition instead of this.
// The details for attaching or detaching a block volume.
type BlockVolumeAttachmentDetails struct {

	// The OCID of the reference compute instance needed to obtain the volume attachment details.
	// This reference compute instance belongs to the peer DR protection group.
	// Example: `ocid1.instance.oc1..uniqueID`
	VolumeAttachmentReferenceInstanceId *string `mandatory:"true" json:"volumeAttachmentReferenceInstanceId"`
}

func (m BlockVolumeAttachmentDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m BlockVolumeAttachmentDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
