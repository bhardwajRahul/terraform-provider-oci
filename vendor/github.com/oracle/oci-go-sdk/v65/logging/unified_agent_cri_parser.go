// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Logging Management API
//
// Use the Logging Management API to create, read, list, update, move and delete
// log groups, log objects, log saved searches, and agent configurations.
// For more information, see Logging Overview (https://docs.oracle.com/iaas/Content/Logging/Concepts/loggingoverview.htm).
//

package logging

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// UnifiedAgentCriParser CRI parser.
type UnifiedAgentCriParser struct {

	// Specifies the time field for the event time. If the event doesn't have this field, the current time is used.
	FieldTimeKey *string `mandatory:"false" json:"fieldTimeKey"`

	// Specify types for converting a field into another type.
	// For example,
	//   With this configuration:
	//       <parse>
	//         @type csv
	//         keys time,host,req_id,user
	//         time_key time
	//       </parse>
	//   This incoming event:
	//     "2013/02/28 12:00:00,192.168.0.1,111,-"
	//   is parsed as:
	//     1362020400 (2013/02/28/ 12:00:00)
	//     record:
	//     {
	//       "host"   : "192.168.0.1",
	//       "req_id" : "111",
	//       "user"   : "-"
	//     }
	Types map[string]string `mandatory:"false" json:"types"`

	// Specify the null value pattern.
	NullValuePattern *string `mandatory:"false" json:"nullValuePattern"`

	// If true, an empty string field is replaced with a null value.
	IsNullEmptyString *bool `mandatory:"false" json:"isNullEmptyString"`

	// If true, use Fluent::EventTime.now(current time) as a timestamp when the time_key is specified.
	IsEstimateCurrentEvent *bool `mandatory:"false" json:"isEstimateCurrentEvent"`

	// If true, keep the time field in the record.
	IsKeepTimeKey *bool `mandatory:"false" json:"isKeepTimeKey"`

	// Specify the timeout for parse processing. This is mainly for detecting an incorrect regexp pattern.
	TimeoutInMilliseconds *int `mandatory:"false" json:"timeoutInMilliseconds"`

	// If you don't need stream or logtag fields, set this to false.
	IsMergeCriFields *bool `mandatory:"false" json:"isMergeCriFields"`

	// Optional nested JSON Parser for CRI. Supported fields are fieldTimeKey, timeFormat, and isKeepTimeKey.
	NestedParser *UnifiedJsonParser `mandatory:"false" json:"nestedParser"`
}

// GetFieldTimeKey returns FieldTimeKey
func (m UnifiedAgentCriParser) GetFieldTimeKey() *string {
	return m.FieldTimeKey
}

// GetTypes returns Types
func (m UnifiedAgentCriParser) GetTypes() map[string]string {
	return m.Types
}

// GetNullValuePattern returns NullValuePattern
func (m UnifiedAgentCriParser) GetNullValuePattern() *string {
	return m.NullValuePattern
}

// GetIsNullEmptyString returns IsNullEmptyString
func (m UnifiedAgentCriParser) GetIsNullEmptyString() *bool {
	return m.IsNullEmptyString
}

// GetIsEstimateCurrentEvent returns IsEstimateCurrentEvent
func (m UnifiedAgentCriParser) GetIsEstimateCurrentEvent() *bool {
	return m.IsEstimateCurrentEvent
}

// GetIsKeepTimeKey returns IsKeepTimeKey
func (m UnifiedAgentCriParser) GetIsKeepTimeKey() *bool {
	return m.IsKeepTimeKey
}

// GetTimeoutInMilliseconds returns TimeoutInMilliseconds
func (m UnifiedAgentCriParser) GetTimeoutInMilliseconds() *int {
	return m.TimeoutInMilliseconds
}

func (m UnifiedAgentCriParser) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UnifiedAgentCriParser) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m UnifiedAgentCriParser) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeUnifiedAgentCriParser UnifiedAgentCriParser
	s := struct {
		DiscriminatorParam string `json:"parserType"`
		MarshalTypeUnifiedAgentCriParser
	}{
		"CRI",
		(MarshalTypeUnifiedAgentCriParser)(m),
	}

	return json.Marshal(&s)
}
