# MsgVpnAclProfileClientConnectExceptionsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]MsgVpnAclProfileClientConnectException**](MsgVpnAclProfileClientConnectException.md) |  | [optional] 
**Links** | Pointer to [**[]MsgVpnAclProfileClientConnectExceptionLinks**](MsgVpnAclProfileClientConnectExceptionLinks.md) |  | [optional] 
**Meta** | [**SempMeta**](SempMeta.md) |  | 

## Methods

### NewMsgVpnAclProfileClientConnectExceptionsResponse

`func NewMsgVpnAclProfileClientConnectExceptionsResponse(meta SempMeta, ) *MsgVpnAclProfileClientConnectExceptionsResponse`

NewMsgVpnAclProfileClientConnectExceptionsResponse instantiates a new MsgVpnAclProfileClientConnectExceptionsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMsgVpnAclProfileClientConnectExceptionsResponseWithDefaults

`func NewMsgVpnAclProfileClientConnectExceptionsResponseWithDefaults() *MsgVpnAclProfileClientConnectExceptionsResponse`

NewMsgVpnAclProfileClientConnectExceptionsResponseWithDefaults instantiates a new MsgVpnAclProfileClientConnectExceptionsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *MsgVpnAclProfileClientConnectExceptionsResponse) GetData() []MsgVpnAclProfileClientConnectException`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *MsgVpnAclProfileClientConnectExceptionsResponse) GetDataOk() (*[]MsgVpnAclProfileClientConnectException, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *MsgVpnAclProfileClientConnectExceptionsResponse) SetData(v []MsgVpnAclProfileClientConnectException)`

SetData sets Data field to given value.

### HasData

`func (o *MsgVpnAclProfileClientConnectExceptionsResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetLinks

`func (o *MsgVpnAclProfileClientConnectExceptionsResponse) GetLinks() []MsgVpnAclProfileClientConnectExceptionLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *MsgVpnAclProfileClientConnectExceptionsResponse) GetLinksOk() (*[]MsgVpnAclProfileClientConnectExceptionLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *MsgVpnAclProfileClientConnectExceptionsResponse) SetLinks(v []MsgVpnAclProfileClientConnectExceptionLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *MsgVpnAclProfileClientConnectExceptionsResponse) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetMeta

`func (o *MsgVpnAclProfileClientConnectExceptionsResponse) GetMeta() SempMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *MsgVpnAclProfileClientConnectExceptionsResponse) GetMetaOk() (*SempMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *MsgVpnAclProfileClientConnectExceptionsResponse) SetMeta(v SempMeta)`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


