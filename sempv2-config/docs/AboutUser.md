# AboutUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GlobalAccessLevel** | Pointer to **string** | The global access level of the User. The allowed values and their meaning are:  &lt;pre&gt; \&quot;admin\&quot; - Full administrative access. \&quot;none\&quot; - No access. \&quot;read-only\&quot; - Read only access. \&quot;read-write\&quot; - Read and write access. &lt;/pre&gt;  | [optional] 
**SessionCreateTime** | Pointer to **int32** | The timestamp of when the session was created. This value represents the number of seconds since 1970-01-01 00:00:00 UTC (Unix time). Available since 2.21. | [optional] 
**SessionCurrentTime** | Pointer to **int32** | The current server timestamp. This is provided as a reference point for the other timestamps provided. This value represents the number of seconds since 1970-01-01 00:00:00 UTC (Unix time). Available since 2.21. | [optional] 
**SessionHardExpiryTime** | Pointer to **int32** | The hard expiry time for the session. After this time the session will be invalid, regardless of activity. This value represents the number of seconds since 1970-01-01 00:00:00 UTC (Unix time). Available since 2.21. | [optional] 
**SessionId** | Pointer to **string** | An identifier for the session to differentiate this session from other sessions for the same user. This value is not guaranteed to be unique between active sessions for different users. Available since 2.21. | [optional] 
**SessionIdleExpiryTime** | Pointer to **int32** | The session idle expiry time. After this time the session will be invalid if there has been no activity. This value represents the number of seconds since 1970-01-01 00:00:00 UTC (Unix time). Available since 2.21. | [optional] 
**Username** | Pointer to **string** | The username of the User. Available since 2.21. | [optional] 

## Methods

### NewAboutUser

`func NewAboutUser() *AboutUser`

NewAboutUser instantiates a new AboutUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAboutUserWithDefaults

`func NewAboutUserWithDefaults() *AboutUser`

NewAboutUserWithDefaults instantiates a new AboutUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGlobalAccessLevel

`func (o *AboutUser) GetGlobalAccessLevel() string`

GetGlobalAccessLevel returns the GlobalAccessLevel field if non-nil, zero value otherwise.

### GetGlobalAccessLevelOk

`func (o *AboutUser) GetGlobalAccessLevelOk() (*string, bool)`

GetGlobalAccessLevelOk returns a tuple with the GlobalAccessLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalAccessLevel

`func (o *AboutUser) SetGlobalAccessLevel(v string)`

SetGlobalAccessLevel sets GlobalAccessLevel field to given value.

### HasGlobalAccessLevel

`func (o *AboutUser) HasGlobalAccessLevel() bool`

HasGlobalAccessLevel returns a boolean if a field has been set.

### GetSessionCreateTime

`func (o *AboutUser) GetSessionCreateTime() int32`

GetSessionCreateTime returns the SessionCreateTime field if non-nil, zero value otherwise.

### GetSessionCreateTimeOk

`func (o *AboutUser) GetSessionCreateTimeOk() (*int32, bool)`

GetSessionCreateTimeOk returns a tuple with the SessionCreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionCreateTime

`func (o *AboutUser) SetSessionCreateTime(v int32)`

SetSessionCreateTime sets SessionCreateTime field to given value.

### HasSessionCreateTime

`func (o *AboutUser) HasSessionCreateTime() bool`

HasSessionCreateTime returns a boolean if a field has been set.

### GetSessionCurrentTime

`func (o *AboutUser) GetSessionCurrentTime() int32`

GetSessionCurrentTime returns the SessionCurrentTime field if non-nil, zero value otherwise.

### GetSessionCurrentTimeOk

`func (o *AboutUser) GetSessionCurrentTimeOk() (*int32, bool)`

GetSessionCurrentTimeOk returns a tuple with the SessionCurrentTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionCurrentTime

`func (o *AboutUser) SetSessionCurrentTime(v int32)`

SetSessionCurrentTime sets SessionCurrentTime field to given value.

### HasSessionCurrentTime

`func (o *AboutUser) HasSessionCurrentTime() bool`

HasSessionCurrentTime returns a boolean if a field has been set.

### GetSessionHardExpiryTime

`func (o *AboutUser) GetSessionHardExpiryTime() int32`

GetSessionHardExpiryTime returns the SessionHardExpiryTime field if non-nil, zero value otherwise.

### GetSessionHardExpiryTimeOk

`func (o *AboutUser) GetSessionHardExpiryTimeOk() (*int32, bool)`

GetSessionHardExpiryTimeOk returns a tuple with the SessionHardExpiryTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionHardExpiryTime

`func (o *AboutUser) SetSessionHardExpiryTime(v int32)`

SetSessionHardExpiryTime sets SessionHardExpiryTime field to given value.

### HasSessionHardExpiryTime

`func (o *AboutUser) HasSessionHardExpiryTime() bool`

HasSessionHardExpiryTime returns a boolean if a field has been set.

### GetSessionId

`func (o *AboutUser) GetSessionId() string`

GetSessionId returns the SessionId field if non-nil, zero value otherwise.

### GetSessionIdOk

`func (o *AboutUser) GetSessionIdOk() (*string, bool)`

GetSessionIdOk returns a tuple with the SessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionId

`func (o *AboutUser) SetSessionId(v string)`

SetSessionId sets SessionId field to given value.

### HasSessionId

`func (o *AboutUser) HasSessionId() bool`

HasSessionId returns a boolean if a field has been set.

### GetSessionIdleExpiryTime

`func (o *AboutUser) GetSessionIdleExpiryTime() int32`

GetSessionIdleExpiryTime returns the SessionIdleExpiryTime field if non-nil, zero value otherwise.

### GetSessionIdleExpiryTimeOk

`func (o *AboutUser) GetSessionIdleExpiryTimeOk() (*int32, bool)`

GetSessionIdleExpiryTimeOk returns a tuple with the SessionIdleExpiryTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionIdleExpiryTime

`func (o *AboutUser) SetSessionIdleExpiryTime(v int32)`

SetSessionIdleExpiryTime sets SessionIdleExpiryTime field to given value.

### HasSessionIdleExpiryTime

`func (o *AboutUser) HasSessionIdleExpiryTime() bool`

HasSessionIdleExpiryTime returns a boolean if a field has been set.

### GetUsername

`func (o *AboutUser) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *AboutUser) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *AboutUser) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *AboutUser) HasUsername() bool`

HasUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


