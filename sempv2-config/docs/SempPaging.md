# SempPaging

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CursorQuery** | **string** | The cursor, or position, for the next page of objects. Use this as the &#x60;cursor&#x60; query parameter of the next request. | 
**NextPageUri** | **string** | The URI of the next page of objects. &#x60;cursorQuery&#x60; is already embedded within this URI. | 

## Methods

### NewSempPaging

`func NewSempPaging(cursorQuery string, nextPageUri string, ) *SempPaging`

NewSempPaging instantiates a new SempPaging object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSempPagingWithDefaults

`func NewSempPagingWithDefaults() *SempPaging`

NewSempPagingWithDefaults instantiates a new SempPaging object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCursorQuery

`func (o *SempPaging) GetCursorQuery() string`

GetCursorQuery returns the CursorQuery field if non-nil, zero value otherwise.

### GetCursorQueryOk

`func (o *SempPaging) GetCursorQueryOk() (*string, bool)`

GetCursorQueryOk returns a tuple with the CursorQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCursorQuery

`func (o *SempPaging) SetCursorQuery(v string)`

SetCursorQuery sets CursorQuery field to given value.


### GetNextPageUri

`func (o *SempPaging) GetNextPageUri() string`

GetNextPageUri returns the NextPageUri field if non-nil, zero value otherwise.

### GetNextPageUriOk

`func (o *SempPaging) GetNextPageUriOk() (*string, bool)`

GetNextPageUriOk returns a tuple with the NextPageUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageUri

`func (o *SempPaging) SetNextPageUri(v string)`

SetNextPageUri sets NextPageUri field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


