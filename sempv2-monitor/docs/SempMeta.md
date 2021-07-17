# SempMeta

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **int64** | The total number of objects requested, irrespective of page size. This may be a count of all objects in a collection or a filtered subset. It represents a snapshot in time and may change when paging through results. | [optional] 
**Error** | Pointer to [**SempError**](SempError.md) |  | [optional] 
**Paging** | Pointer to [**SempPaging**](SempPaging.md) |  | [optional] 
**Request** | [**SempRequest**](SempRequest.md) |  | 
**ResponseCode** | **int32** | The HTTP response code, one of 200 (success), 4xx (client error), or 5xx (server error). | 

## Methods

### NewSempMeta

`func NewSempMeta(request SempRequest, responseCode int32, ) *SempMeta`

NewSempMeta instantiates a new SempMeta object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSempMetaWithDefaults

`func NewSempMetaWithDefaults() *SempMeta`

NewSempMetaWithDefaults instantiates a new SempMeta object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *SempMeta) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *SempMeta) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *SempMeta) SetCount(v int64)`

SetCount sets Count field to given value.

### HasCount

`func (o *SempMeta) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetError

`func (o *SempMeta) GetError() SempError`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *SempMeta) GetErrorOk() (*SempError, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *SempMeta) SetError(v SempError)`

SetError sets Error field to given value.

### HasError

`func (o *SempMeta) HasError() bool`

HasError returns a boolean if a field has been set.

### GetPaging

`func (o *SempMeta) GetPaging() SempPaging`

GetPaging returns the Paging field if non-nil, zero value otherwise.

### GetPagingOk

`func (o *SempMeta) GetPagingOk() (*SempPaging, bool)`

GetPagingOk returns a tuple with the Paging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaging

`func (o *SempMeta) SetPaging(v SempPaging)`

SetPaging sets Paging field to given value.

### HasPaging

`func (o *SempMeta) HasPaging() bool`

HasPaging returns a boolean if a field has been set.

### GetRequest

`func (o *SempMeta) GetRequest() SempRequest`

GetRequest returns the Request field if non-nil, zero value otherwise.

### GetRequestOk

`func (o *SempMeta) GetRequestOk() (*SempRequest, bool)`

GetRequestOk returns a tuple with the Request field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequest

`func (o *SempMeta) SetRequest(v SempRequest)`

SetRequest sets Request field to given value.


### GetResponseCode

`func (o *SempMeta) GetResponseCode() int32`

GetResponseCode returns the ResponseCode field if non-nil, zero value otherwise.

### GetResponseCodeOk

`func (o *SempMeta) GetResponseCodeOk() (*int32, bool)`

GetResponseCodeOk returns a tuple with the ResponseCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseCode

`func (o *SempMeta) SetResponseCode(v int32)`

SetResponseCode sets ResponseCode field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


