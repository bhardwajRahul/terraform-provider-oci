// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Migration API
//
// Use the Oracle Cloud Infrastructure Database Migration APIs to perform database migration operations.
//

package databasemigration

import (
	"strings"
)

// SortOrdersEnum Enum with underlying type: string
type SortOrdersEnum string

// Set of constants representing the allowable values for SortOrdersEnum
const (
	SortOrdersAsc  SortOrdersEnum = "ASC"
	SortOrdersDesc SortOrdersEnum = "DESC"
)

var mappingSortOrdersEnum = map[string]SortOrdersEnum{
	"ASC":  SortOrdersAsc,
	"DESC": SortOrdersDesc,
}

var mappingSortOrdersEnumLowerCase = map[string]SortOrdersEnum{
	"asc":  SortOrdersAsc,
	"desc": SortOrdersDesc,
}

// GetSortOrdersEnumValues Enumerates the set of values for SortOrdersEnum
func GetSortOrdersEnumValues() []SortOrdersEnum {
	values := make([]SortOrdersEnum, 0)
	for _, v := range mappingSortOrdersEnum {
		values = append(values, v)
	}
	return values
}

// GetSortOrdersEnumStringValues Enumerates the set of values in String for SortOrdersEnum
func GetSortOrdersEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingSortOrdersEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSortOrdersEnum(val string) (SortOrdersEnum, bool) {
	enum, ok := mappingSortOrdersEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
