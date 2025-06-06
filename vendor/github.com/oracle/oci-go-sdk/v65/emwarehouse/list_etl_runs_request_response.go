// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package emwarehouse

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListEtlRunsRequest wrapper for the ListEtlRuns operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/emwarehouse/ListEtlRuns.go.html to see an example of how to use ListEtlRunsRequest.
type ListEtlRunsRequest struct {

	// unique EmWarehouse identifier
	EmWarehouseId *string `mandatory:"true" contributesTo:"path" name:"emWarehouseId"`

	// The ID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"false" contributesTo:"query" name:"compartmentId"`

	// A filter to return only resources that match the entire display name given.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// A token representing the position at which to start retrieving results. This must come from the `opc-next-page` header field of a previous response.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'ASC' or 'DESC'.
	SortOrder ListEtlRunsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. Default order for displayName is ascending.
	SortBy ListEtlRunsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListEtlRunsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListEtlRunsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListEtlRunsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListEtlRunsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListEtlRunsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListEtlRunsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListEtlRunsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListEtlRunsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListEtlRunsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListEtlRunsResponse wrapper for the ListEtlRuns operation
type ListEtlRunsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of EtlRunCollection instances
	EtlRunCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListEtlRunsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListEtlRunsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListEtlRunsSortOrderEnum Enum with underlying type: string
type ListEtlRunsSortOrderEnum string

// Set of constants representing the allowable values for ListEtlRunsSortOrderEnum
const (
	ListEtlRunsSortOrderAsc  ListEtlRunsSortOrderEnum = "ASC"
	ListEtlRunsSortOrderDesc ListEtlRunsSortOrderEnum = "DESC"
)

var mappingListEtlRunsSortOrderEnum = map[string]ListEtlRunsSortOrderEnum{
	"ASC":  ListEtlRunsSortOrderAsc,
	"DESC": ListEtlRunsSortOrderDesc,
}

var mappingListEtlRunsSortOrderEnumLowerCase = map[string]ListEtlRunsSortOrderEnum{
	"asc":  ListEtlRunsSortOrderAsc,
	"desc": ListEtlRunsSortOrderDesc,
}

// GetListEtlRunsSortOrderEnumValues Enumerates the set of values for ListEtlRunsSortOrderEnum
func GetListEtlRunsSortOrderEnumValues() []ListEtlRunsSortOrderEnum {
	values := make([]ListEtlRunsSortOrderEnum, 0)
	for _, v := range mappingListEtlRunsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListEtlRunsSortOrderEnumStringValues Enumerates the set of values in String for ListEtlRunsSortOrderEnum
func GetListEtlRunsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListEtlRunsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListEtlRunsSortOrderEnum(val string) (ListEtlRunsSortOrderEnum, bool) {
	enum, ok := mappingListEtlRunsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListEtlRunsSortByEnum Enum with underlying type: string
type ListEtlRunsSortByEnum string

// Set of constants representing the allowable values for ListEtlRunsSortByEnum
const (
	ListEtlRunsSortByTimecreated ListEtlRunsSortByEnum = "timeCreated"
	ListEtlRunsSortByDisplayname ListEtlRunsSortByEnum = "displayName"
)

var mappingListEtlRunsSortByEnum = map[string]ListEtlRunsSortByEnum{
	"timeCreated": ListEtlRunsSortByTimecreated,
	"displayName": ListEtlRunsSortByDisplayname,
}

var mappingListEtlRunsSortByEnumLowerCase = map[string]ListEtlRunsSortByEnum{
	"timecreated": ListEtlRunsSortByTimecreated,
	"displayname": ListEtlRunsSortByDisplayname,
}

// GetListEtlRunsSortByEnumValues Enumerates the set of values for ListEtlRunsSortByEnum
func GetListEtlRunsSortByEnumValues() []ListEtlRunsSortByEnum {
	values := make([]ListEtlRunsSortByEnum, 0)
	for _, v := range mappingListEtlRunsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListEtlRunsSortByEnumStringValues Enumerates the set of values in String for ListEtlRunsSortByEnum
func GetListEtlRunsSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListEtlRunsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListEtlRunsSortByEnum(val string) (ListEtlRunsSortByEnum, bool) {
	enum, ok := mappingListEtlRunsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
