// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// APM Availability Monitoring API
//
// Use the APM Availability Monitoring API to query Scripts, Monitors, Dedicated Vantage Points and On-Premise Vantage Points resources. For more information, see Application Performance Monitoring (https://docs.oracle.com/iaas/application-performance-monitoring/index.html).
//

package apmsynthetics

import (
	"strings"
)

// DedicatedVantagePointStatusEnum Enum with underlying type: string
type DedicatedVantagePointStatusEnum string

// Set of constants representing the allowable values for DedicatedVantagePointStatusEnum
const (
	DedicatedVantagePointStatusEnabled  DedicatedVantagePointStatusEnum = "ENABLED"
	DedicatedVantagePointStatusDisabled DedicatedVantagePointStatusEnum = "DISABLED"
)

var mappingDedicatedVantagePointStatusEnum = map[string]DedicatedVantagePointStatusEnum{
	"ENABLED":  DedicatedVantagePointStatusEnabled,
	"DISABLED": DedicatedVantagePointStatusDisabled,
}

var mappingDedicatedVantagePointStatusEnumLowerCase = map[string]DedicatedVantagePointStatusEnum{
	"enabled":  DedicatedVantagePointStatusEnabled,
	"disabled": DedicatedVantagePointStatusDisabled,
}

// GetDedicatedVantagePointStatusEnumValues Enumerates the set of values for DedicatedVantagePointStatusEnum
func GetDedicatedVantagePointStatusEnumValues() []DedicatedVantagePointStatusEnum {
	values := make([]DedicatedVantagePointStatusEnum, 0)
	for _, v := range mappingDedicatedVantagePointStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetDedicatedVantagePointStatusEnumStringValues Enumerates the set of values in String for DedicatedVantagePointStatusEnum
func GetDedicatedVantagePointStatusEnumStringValues() []string {
	return []string{
		"ENABLED",
		"DISABLED",
	}
}

// GetMappingDedicatedVantagePointStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDedicatedVantagePointStatusEnum(val string) (DedicatedVantagePointStatusEnum, bool) {
	enum, ok := mappingDedicatedVantagePointStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
