# MsgVpnClientUsernamesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Collections** | Pointer to **[]map[string]interface{}** |  | [optional] 
**Data** | Pointer to [**[]MsgVpnClientUsername**](MsgVpnClientUsername.md) |  | [optional] 
**Links** | Pointer to [**[]MsgVpnClientUsernameLinks**](MsgVpnClientUsernameLinks.md) |  | [optional] 
**Meta** | [**SempMeta**](SempMeta.md) |  | 

## Methods

### NewMsgVpnClientUsernamesResponse

`func NewMsgVpnClientUsernamesResponse(meta SempMeta, ) *MsgVpnClientUsernamesResponse`

NewMsgVpnClientUsernamesResponse instantiates a new MsgVpnClientUsernamesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMsgVpnClientUsernamesResponseWithDefaults

`func NewMsgVpnClientUsernamesResponseWithDefaults() *MsgVpnClientUsernamesResponse`

NewMsgVpnClientUsernamesResponseWithDefaults instantiates a new MsgVpnClientUsernamesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCollections

`func (o *MsgVpnClientUsernamesResponse) GetCollections() []map[string]interface{}`

GetCollections returns the Collections field if non-nil, zero value otherwise.

### GetCollectionsOk

`func (o *MsgVpnClientUsernamesResponse) GetCollectionsOk() (*[]map[string]interface{}, bool)`

GetCollectionsOk returns a tuple with the Collections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollections

`func (o *MsgVpnClientUsernamesResponse) SetCollections(v []map[string]interface{})`

SetCollections sets Collections field to given value.

### HasCollections

`func (o *MsgVpnClientUsernamesResponse) HasCollections() bool`

HasCollections returns a boolean if a field has been set.

### GetData

`func (o *MsgVpnClientUsernamesResponse) GetData() []MsgVpnClientUsername`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *MsgVpnClientUsernamesResponse) GetDataOk() (*[]MsgVpnClientUsername, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *MsgVpnClientUsernamesResponse) SetData(v []MsgVpnClientUsername)`

SetData sets Data field to given value.

### HasData

`func (o *MsgVpnClientUsernamesResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetLinks

`func (o *MsgVpnClientUsernamesResponse) GetLinks() []MsgVpnClientUsernameLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *MsgVpnClientUsernamesResponse) GetLinksOk() (*[]MsgVpnClientUsernameLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *MsgVpnClientUsernamesResponse) SetLinks(v []MsgVpnClientUsernameLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *MsgVpnClientUsernamesResponse) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetMeta

`func (o *MsgVpnClientUsernamesResponse) GetMeta() SempMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *MsgVpnClientUsernamesResponse) GetMetaOk() (*SempMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *MsgVpnClientUsernamesResponse) SetMeta(v SempMeta)`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


