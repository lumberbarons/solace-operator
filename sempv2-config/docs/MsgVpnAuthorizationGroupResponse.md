# MsgVpnAuthorizationGroupResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**MsgVpnAuthorizationGroup**](MsgVpnAuthorizationGroup.md) |  | [optional] 
**Links** | Pointer to [**MsgVpnAuthorizationGroupLinks**](MsgVpnAuthorizationGroupLinks.md) |  | [optional] 
**Meta** | [**SempMeta**](SempMeta.md) |  | 

## Methods

### NewMsgVpnAuthorizationGroupResponse

`func NewMsgVpnAuthorizationGroupResponse(meta SempMeta, ) *MsgVpnAuthorizationGroupResponse`

NewMsgVpnAuthorizationGroupResponse instantiates a new MsgVpnAuthorizationGroupResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMsgVpnAuthorizationGroupResponseWithDefaults

`func NewMsgVpnAuthorizationGroupResponseWithDefaults() *MsgVpnAuthorizationGroupResponse`

NewMsgVpnAuthorizationGroupResponseWithDefaults instantiates a new MsgVpnAuthorizationGroupResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *MsgVpnAuthorizationGroupResponse) GetData() MsgVpnAuthorizationGroup`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *MsgVpnAuthorizationGroupResponse) GetDataOk() (*MsgVpnAuthorizationGroup, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *MsgVpnAuthorizationGroupResponse) SetData(v MsgVpnAuthorizationGroup)`

SetData sets Data field to given value.

### HasData

`func (o *MsgVpnAuthorizationGroupResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetLinks

`func (o *MsgVpnAuthorizationGroupResponse) GetLinks() MsgVpnAuthorizationGroupLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *MsgVpnAuthorizationGroupResponse) GetLinksOk() (*MsgVpnAuthorizationGroupLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *MsgVpnAuthorizationGroupResponse) SetLinks(v MsgVpnAuthorizationGroupLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *MsgVpnAuthorizationGroupResponse) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetMeta

`func (o *MsgVpnAuthorizationGroupResponse) GetMeta() SempMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *MsgVpnAuthorizationGroupResponse) GetMetaOk() (*SempMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *MsgVpnAuthorizationGroupResponse) SetMeta(v SempMeta)`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


