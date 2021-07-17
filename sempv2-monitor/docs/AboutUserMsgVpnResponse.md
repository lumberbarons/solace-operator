# AboutUserMsgVpnResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Collections** | Pointer to **map[string]interface{}** |  | [optional] 
**Data** | Pointer to [**AboutUserMsgVpn**](AboutUserMsgVpn.md) |  | [optional] 
**Links** | Pointer to [**AboutUserMsgVpnLinks**](AboutUserMsgVpnLinks.md) |  | [optional] 
**Meta** | [**SempMeta**](SempMeta.md) |  | 

## Methods

### NewAboutUserMsgVpnResponse

`func NewAboutUserMsgVpnResponse(meta SempMeta, ) *AboutUserMsgVpnResponse`

NewAboutUserMsgVpnResponse instantiates a new AboutUserMsgVpnResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAboutUserMsgVpnResponseWithDefaults

`func NewAboutUserMsgVpnResponseWithDefaults() *AboutUserMsgVpnResponse`

NewAboutUserMsgVpnResponseWithDefaults instantiates a new AboutUserMsgVpnResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCollections

`func (o *AboutUserMsgVpnResponse) GetCollections() map[string]interface{}`

GetCollections returns the Collections field if non-nil, zero value otherwise.

### GetCollectionsOk

`func (o *AboutUserMsgVpnResponse) GetCollectionsOk() (*map[string]interface{}, bool)`

GetCollectionsOk returns a tuple with the Collections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollections

`func (o *AboutUserMsgVpnResponse) SetCollections(v map[string]interface{})`

SetCollections sets Collections field to given value.

### HasCollections

`func (o *AboutUserMsgVpnResponse) HasCollections() bool`

HasCollections returns a boolean if a field has been set.

### GetData

`func (o *AboutUserMsgVpnResponse) GetData() AboutUserMsgVpn`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *AboutUserMsgVpnResponse) GetDataOk() (*AboutUserMsgVpn, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *AboutUserMsgVpnResponse) SetData(v AboutUserMsgVpn)`

SetData sets Data field to given value.

### HasData

`func (o *AboutUserMsgVpnResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetLinks

`func (o *AboutUserMsgVpnResponse) GetLinks() AboutUserMsgVpnLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *AboutUserMsgVpnResponse) GetLinksOk() (*AboutUserMsgVpnLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *AboutUserMsgVpnResponse) SetLinks(v AboutUserMsgVpnLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *AboutUserMsgVpnResponse) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetMeta

`func (o *AboutUserMsgVpnResponse) GetMeta() SempMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *AboutUserMsgVpnResponse) GetMetaOk() (*SempMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *AboutUserMsgVpnResponse) SetMeta(v SempMeta)`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


