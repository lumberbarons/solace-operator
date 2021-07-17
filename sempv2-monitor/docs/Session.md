# Session

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreateTime** | Pointer to **int32** | The timestamp of when the session was created. This value represents the number of seconds since 1970-01-01 00:00:00 UTC (Unix time). | [optional] 
**LastActivityTime** | Pointer to **int32** | The timestamp of when the last activity on the session occurred. This value represents the number of seconds since 1970-01-01 00:00:00 UTC (Unix time). | [optional] 
**SessionId** | Pointer to **string** | The unique identifier for the session. | [optional] 
**SessionUsername** | Pointer to **string** | The username used for authorization. | [optional] 

## Methods

### NewSession

`func NewSession() *Session`

NewSession instantiates a new Session object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSessionWithDefaults

`func NewSessionWithDefaults() *Session`

NewSessionWithDefaults instantiates a new Session object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreateTime

`func (o *Session) GetCreateTime() int32`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *Session) GetCreateTimeOk() (*int32, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *Session) SetCreateTime(v int32)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *Session) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetLastActivityTime

`func (o *Session) GetLastActivityTime() int32`

GetLastActivityTime returns the LastActivityTime field if non-nil, zero value otherwise.

### GetLastActivityTimeOk

`func (o *Session) GetLastActivityTimeOk() (*int32, bool)`

GetLastActivityTimeOk returns a tuple with the LastActivityTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastActivityTime

`func (o *Session) SetLastActivityTime(v int32)`

SetLastActivityTime sets LastActivityTime field to given value.

### HasLastActivityTime

`func (o *Session) HasLastActivityTime() bool`

HasLastActivityTime returns a boolean if a field has been set.

### GetSessionId

`func (o *Session) GetSessionId() string`

GetSessionId returns the SessionId field if non-nil, zero value otherwise.

### GetSessionIdOk

`func (o *Session) GetSessionIdOk() (*string, bool)`

GetSessionIdOk returns a tuple with the SessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionId

`func (o *Session) SetSessionId(v string)`

SetSessionId sets SessionId field to given value.

### HasSessionId

`func (o *Session) HasSessionId() bool`

HasSessionId returns a boolean if a field has been set.

### GetSessionUsername

`func (o *Session) GetSessionUsername() string`

GetSessionUsername returns the SessionUsername field if non-nil, zero value otherwise.

### GetSessionUsernameOk

`func (o *Session) GetSessionUsernameOk() (*string, bool)`

GetSessionUsernameOk returns a tuple with the SessionUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionUsername

`func (o *Session) SetSessionUsername(v string)`

SetSessionUsername sets SessionUsername field to given value.

### HasSessionUsername

`func (o *Session) HasSessionUsername() bool`

HasSessionUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


