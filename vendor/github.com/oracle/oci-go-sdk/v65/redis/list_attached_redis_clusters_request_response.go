// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package redis

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListAttachedRedisClustersRequest wrapper for the ListAttachedRedisClusters operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/redis/ListAttachedRedisClusters.go.html to see an example of how to use ListAttachedRedisClustersRequest.
type ListAttachedRedisClustersRequest struct {

	// A filter to return only resources, that match with the given OCI cache user ID (OCID).
	OciCacheUserId *string `mandatory:"true" contributesTo:"path" name:"ociCacheUserId"`

	// The ID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"false" contributesTo:"query" name:"compartmentId"`

	// A filter to return only resources that match the entire display name given.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// A token representing the position at which to start retrieving results. This must come from the `opc-next-page` header field of a previous response.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'ASC' or 'DESC'.
	SortOrder ListAttachedRedisClustersSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. Default order for displayName is ascending.
	SortBy ListAttachedRedisClustersSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListAttachedRedisClustersRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListAttachedRedisClustersRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListAttachedRedisClustersRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListAttachedRedisClustersRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListAttachedRedisClustersRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListAttachedRedisClustersSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListAttachedRedisClustersSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListAttachedRedisClustersSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListAttachedRedisClustersSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListAttachedRedisClustersResponse wrapper for the ListAttachedRedisClusters operation
type ListAttachedRedisClustersResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []AttachedOciCacheCluster instances
	Items []AttachedOciCacheCluster `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListAttachedRedisClustersResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListAttachedRedisClustersResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListAttachedRedisClustersSortOrderEnum Enum with underlying type: string
type ListAttachedRedisClustersSortOrderEnum string

// Set of constants representing the allowable values for ListAttachedRedisClustersSortOrderEnum
const (
	ListAttachedRedisClustersSortOrderAsc  ListAttachedRedisClustersSortOrderEnum = "ASC"
	ListAttachedRedisClustersSortOrderDesc ListAttachedRedisClustersSortOrderEnum = "DESC"
)

var mappingListAttachedRedisClustersSortOrderEnum = map[string]ListAttachedRedisClustersSortOrderEnum{
	"ASC":  ListAttachedRedisClustersSortOrderAsc,
	"DESC": ListAttachedRedisClustersSortOrderDesc,
}

var mappingListAttachedRedisClustersSortOrderEnumLowerCase = map[string]ListAttachedRedisClustersSortOrderEnum{
	"asc":  ListAttachedRedisClustersSortOrderAsc,
	"desc": ListAttachedRedisClustersSortOrderDesc,
}

// GetListAttachedRedisClustersSortOrderEnumValues Enumerates the set of values for ListAttachedRedisClustersSortOrderEnum
func GetListAttachedRedisClustersSortOrderEnumValues() []ListAttachedRedisClustersSortOrderEnum {
	values := make([]ListAttachedRedisClustersSortOrderEnum, 0)
	for _, v := range mappingListAttachedRedisClustersSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListAttachedRedisClustersSortOrderEnumStringValues Enumerates the set of values in String for ListAttachedRedisClustersSortOrderEnum
func GetListAttachedRedisClustersSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListAttachedRedisClustersSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListAttachedRedisClustersSortOrderEnum(val string) (ListAttachedRedisClustersSortOrderEnum, bool) {
	enum, ok := mappingListAttachedRedisClustersSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListAttachedRedisClustersSortByEnum Enum with underlying type: string
type ListAttachedRedisClustersSortByEnum string

// Set of constants representing the allowable values for ListAttachedRedisClustersSortByEnum
const (
	ListAttachedRedisClustersSortByTimecreated ListAttachedRedisClustersSortByEnum = "timeCreated"
	ListAttachedRedisClustersSortByDisplayname ListAttachedRedisClustersSortByEnum = "displayName"
)

var mappingListAttachedRedisClustersSortByEnum = map[string]ListAttachedRedisClustersSortByEnum{
	"timeCreated": ListAttachedRedisClustersSortByTimecreated,
	"displayName": ListAttachedRedisClustersSortByDisplayname,
}

var mappingListAttachedRedisClustersSortByEnumLowerCase = map[string]ListAttachedRedisClustersSortByEnum{
	"timecreated": ListAttachedRedisClustersSortByTimecreated,
	"displayname": ListAttachedRedisClustersSortByDisplayname,
}

// GetListAttachedRedisClustersSortByEnumValues Enumerates the set of values for ListAttachedRedisClustersSortByEnum
func GetListAttachedRedisClustersSortByEnumValues() []ListAttachedRedisClustersSortByEnum {
	values := make([]ListAttachedRedisClustersSortByEnum, 0)
	for _, v := range mappingListAttachedRedisClustersSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListAttachedRedisClustersSortByEnumStringValues Enumerates the set of values in String for ListAttachedRedisClustersSortByEnum
func GetListAttachedRedisClustersSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListAttachedRedisClustersSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListAttachedRedisClustersSortByEnum(val string) (ListAttachedRedisClustersSortByEnum, bool) {
	enum, ok := mappingListAttachedRedisClustersSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
