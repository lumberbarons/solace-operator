# MsgVpnAclProfilesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]MsgVpnAclProfile**](MsgVpnAclProfile.md) |  | [optional] 
**Links** | Pointer to [**[]MsgVpnAclProfileLinks**](MsgVpnAclProfileLinks.md) |  | [optional] 
**Meta** | [**SempMeta**](SempMeta.md) |  | 

## Methods

### NewMsgVpnAclProfilesResponse

`func NewMsgVpnAclProfilesResponse(meta SempMeta, ) *MsgVpnAclProfilesResponse`

NewMsgVpnAclProfilesResponse instantiates a new MsgVpnAclProfilesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMsgVpnAclProfilesResponseWithDefaults

`func NewMsgVpnAclProfilesResponseWithDefaults() *MsgVpnAclProfilesResponse`

NewMsgVpnAclProfilesResponseWithDefaults instantiates a new MsgVpnAclProfilesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *MsgVpnAclProfilesResponse) GetData() []MsgVpnAclProfile`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *MsgVpnAclProfilesResponse) GetDataOk() (*[]MsgVpnAclProfile, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *MsgVpnAclProfilesResponse) SetData(v []MsgVpnAclProfile)`

SetData sets Data field to given value.

### HasData

`func (o *MsgVpnAclProfilesResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetLinks

`func (o *MsgVpnAclProfilesResponse) GetLinks() []MsgVpnAclProfileLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *MsgVpnAclProfilesResponse) GetLinksOk() (*[]MsgVpnAclProfileLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *MsgVpnAclProfilesResponse) SetLinks(v []MsgVpnAclProfileLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *MsgVpnAclProfilesResponse) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetMeta

`func (o *MsgVpnAclProfilesResponse) GetMeta() SempMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *MsgVpnAclProfilesResponse) GetMetaOk() (*SempMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *MsgVpnAclProfilesResponse) SetMeta(v SempMeta)`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


