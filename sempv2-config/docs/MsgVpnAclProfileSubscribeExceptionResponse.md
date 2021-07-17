# MsgVpnAclProfileSubscribeExceptionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**MsgVpnAclProfileSubscribeException**](MsgVpnAclProfileSubscribeException.md) |  | [optional] 
**Links** | Pointer to [**MsgVpnAclProfileSubscribeExceptionLinks**](MsgVpnAclProfileSubscribeExceptionLinks.md) |  | [optional] 
**Meta** | [**SempMeta**](SempMeta.md) |  | 

## Methods

### NewMsgVpnAclProfileSubscribeExceptionResponse

`func NewMsgVpnAclProfileSubscribeExceptionResponse(meta SempMeta, ) *MsgVpnAclProfileSubscribeExceptionResponse`

NewMsgVpnAclProfileSubscribeExceptionResponse instantiates a new MsgVpnAclProfileSubscribeExceptionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMsgVpnAclProfileSubscribeExceptionResponseWithDefaults

`func NewMsgVpnAclProfileSubscribeExceptionResponseWithDefaults() *MsgVpnAclProfileSubscribeExceptionResponse`

NewMsgVpnAclProfileSubscribeExceptionResponseWithDefaults instantiates a new MsgVpnAclProfileSubscribeExceptionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *MsgVpnAclProfileSubscribeExceptionResponse) GetData() MsgVpnAclProfileSubscribeException`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *MsgVpnAclProfileSubscribeExceptionResponse) GetDataOk() (*MsgVpnAclProfileSubscribeException, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *MsgVpnAclProfileSubscribeExceptionResponse) SetData(v MsgVpnAclProfileSubscribeException)`

SetData sets Data field to given value.

### HasData

`func (o *MsgVpnAclProfileSubscribeExceptionResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetLinks

`func (o *MsgVpnAclProfileSubscribeExceptionResponse) GetLinks() MsgVpnAclProfileSubscribeExceptionLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *MsgVpnAclProfileSubscribeExceptionResponse) GetLinksOk() (*MsgVpnAclProfileSubscribeExceptionLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *MsgVpnAclProfileSubscribeExceptionResponse) SetLinks(v MsgVpnAclProfileSubscribeExceptionLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *MsgVpnAclProfileSubscribeExceptionResponse) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetMeta

`func (o *MsgVpnAclProfileSubscribeExceptionResponse) GetMeta() SempMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *MsgVpnAclProfileSubscribeExceptionResponse) GetMetaOk() (*SempMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *MsgVpnAclProfileSubscribeExceptionResponse) SetMeta(v SempMeta)`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


