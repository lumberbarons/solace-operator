# MsgVpnsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]MsgVpn**](MsgVpn.md) |  | [optional] 
**Links** | Pointer to [**[]MsgVpnLinks**](MsgVpnLinks.md) |  | [optional] 
**Meta** | [**SempMeta**](SempMeta.md) |  | 

## Methods

### NewMsgVpnsResponse

`func NewMsgVpnsResponse(meta SempMeta, ) *MsgVpnsResponse`

NewMsgVpnsResponse instantiates a new MsgVpnsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMsgVpnsResponseWithDefaults

`func NewMsgVpnsResponseWithDefaults() *MsgVpnsResponse`

NewMsgVpnsResponseWithDefaults instantiates a new MsgVpnsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *MsgVpnsResponse) GetData() []MsgVpn`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *MsgVpnsResponse) GetDataOk() (*[]MsgVpn, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *MsgVpnsResponse) SetData(v []MsgVpn)`

SetData sets Data field to given value.

### HasData

`func (o *MsgVpnsResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetLinks

`func (o *MsgVpnsResponse) GetLinks() []MsgVpnLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *MsgVpnsResponse) GetLinksOk() (*[]MsgVpnLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *MsgVpnsResponse) SetLinks(v []MsgVpnLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *MsgVpnsResponse) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetMeta

`func (o *MsgVpnsResponse) GetMeta() SempMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *MsgVpnsResponse) GetMetaOk() (*SempMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *MsgVpnsResponse) SetMeta(v SempMeta)`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


