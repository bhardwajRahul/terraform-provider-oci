// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package dataintegration

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// GetFunctionLibraryRequest wrapper for the GetFunctionLibrary operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dataintegration/GetFunctionLibrary.go.html to see an example of how to use GetFunctionLibraryRequest.
type GetFunctionLibraryRequest struct {

	// The workspace ID.
	WorkspaceId *string `mandatory:"true" contributesTo:"path" name:"workspaceId"`

	// The functionLibrary key.
	FunctionLibraryKey *string `mandatory:"true" contributesTo:"path" name:"functionLibraryKey"`

	// Unique Oracle-assigned identifier for the request. If
	// you need to contact Oracle about a particular request,
	// please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// This parameter allows users to specify which view of the object to return. CHILD_COUNT_STATISTICS - This option is used to get statistics on immediate children of the object by their type.
	Projection []GetFunctionLibraryProjectionEnum `contributesTo:"query" name:"projection" omitEmpty:"true" collectionFormat:"multi"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request GetFunctionLibraryRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request GetFunctionLibraryRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request GetFunctionLibraryRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request GetFunctionLibraryRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request GetFunctionLibraryRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	for _, val := range request.Projection {
		if _, ok := GetMappingGetFunctionLibraryProjectionEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Projection: %s. Supported values are: %s.", val, strings.Join(GetGetFunctionLibraryProjectionEnumStringValues(), ",")))
		}
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// GetFunctionLibraryResponse wrapper for the GetFunctionLibrary operation
type GetFunctionLibraryResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The FunctionLibrary instance
	FunctionLibrary `presentIn:"body"`

	// For optimistic concurrency control. See ETags for Optimistic Concurrency Control (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#eleven).
	Etag *string `presentIn:"header" name:"etag"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response GetFunctionLibraryResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response GetFunctionLibraryResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// GetFunctionLibraryProjectionEnum Enum with underlying type: string
type GetFunctionLibraryProjectionEnum string

// Set of constants representing the allowable values for GetFunctionLibraryProjectionEnum
const (
	GetFunctionLibraryProjectionChildCountStatistics GetFunctionLibraryProjectionEnum = "CHILD_COUNT_STATISTICS"
)

var mappingGetFunctionLibraryProjectionEnum = map[string]GetFunctionLibraryProjectionEnum{
	"CHILD_COUNT_STATISTICS": GetFunctionLibraryProjectionChildCountStatistics,
}

var mappingGetFunctionLibraryProjectionEnumLowerCase = map[string]GetFunctionLibraryProjectionEnum{
	"child_count_statistics": GetFunctionLibraryProjectionChildCountStatistics,
}

// GetGetFunctionLibraryProjectionEnumValues Enumerates the set of values for GetFunctionLibraryProjectionEnum
func GetGetFunctionLibraryProjectionEnumValues() []GetFunctionLibraryProjectionEnum {
	values := make([]GetFunctionLibraryProjectionEnum, 0)
	for _, v := range mappingGetFunctionLibraryProjectionEnum {
		values = append(values, v)
	}
	return values
}

// GetGetFunctionLibraryProjectionEnumStringValues Enumerates the set of values in String for GetFunctionLibraryProjectionEnum
func GetGetFunctionLibraryProjectionEnumStringValues() []string {
	return []string{
		"CHILD_COUNT_STATISTICS",
	}
}

// GetMappingGetFunctionLibraryProjectionEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingGetFunctionLibraryProjectionEnum(val string) (GetFunctionLibraryProjectionEnum, bool) {
	enum, ok := mappingGetFunctionLibraryProjectionEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
