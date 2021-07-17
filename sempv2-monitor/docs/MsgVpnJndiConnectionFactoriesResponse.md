# MsgVpnJndiConnectionFactoriesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Collections** | Pointer to **[]map[string]interface{}** |  | [optional] 
**Data** | Pointer to [**[]MsgVpnJndiConnectionFactory**](MsgVpnJndiConnectionFactory.md) |  | [optional] 
**Links** | Pointer to [**[]MsgVpnJndiConnectionFactoryLinks**](MsgVpnJndiConnectionFactoryLinks.md) |  | [optional] 
**Meta** | [**SempMeta**](SempMeta.md) |  | 

## Methods

### NewMsgVpnJndiConnectionFactoriesResponse

`func NewMsgVpnJndiConnectionFactoriesResponse(meta SempMeta, ) *MsgVpnJndiConnectionFactoriesResponse`

NewMsgVpnJndiConnectionFactoriesResponse instantiates a new MsgVpnJndiConnectionFactoriesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMsgVpnJndiConnectionFactoriesResponseWithDefaults

`func NewMsgVpnJndiConnectionFactoriesResponseWithDefaults() *MsgVpnJndiConnectionFactoriesResponse`

NewMsgVpnJndiConnectionFactoriesResponseWithDefaults instantiates a new MsgVpnJndiConnectionFactoriesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCollections

`func (o *MsgVpnJndiConnectionFactoriesResponse) GetCollections() []map[string]interface{}`

GetCollections returns the Collections field if non-nil, zero value otherwise.

### GetCollectionsOk

`func (o *MsgVpnJndiConnectionFactoriesResponse) GetCollectionsOk() (*[]map[string]interface{}, bool)`

GetCollectionsOk returns a tuple with the Collections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollections

`func (o *MsgVpnJndiConnectionFactoriesResponse) SetCollections(v []map[string]interface{})`

SetCollections sets Collections field to given value.

### HasCollections

`func (o *MsgVpnJndiConnectionFactoriesResponse) HasCollections() bool`

HasCollections returns a boolean if a field has been set.

### GetData

`func (o *MsgVpnJndiConnectionFactoriesResponse) GetData() []MsgVpnJndiConnectionFactory`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *MsgVpnJndiConnectionFactoriesResponse) GetDataOk() (*[]MsgVpnJndiConnectionFactory, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *MsgVpnJndiConnectionFactoriesResponse) SetData(v []MsgVpnJndiConnectionFactory)`

SetData sets Data field to given value.

### HasData

`func (o *MsgVpnJndiConnectionFactoriesResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetLinks

`func (o *MsgVpnJndiConnectionFactoriesResponse) GetLinks() []MsgVpnJndiConnectionFactoryLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *MsgVpnJndiConnectionFactoriesResponse) GetLinksOk() (*[]MsgVpnJndiConnectionFactoryLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *MsgVpnJndiConnectionFactoriesResponse) SetLinks(v []MsgVpnJndiConnectionFactoryLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *MsgVpnJndiConnectionFactoriesResponse) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetMeta

`func (o *MsgVpnJndiConnectionFactoriesResponse) GetMeta() SempMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *MsgVpnJndiConnectionFactoriesResponse) GetMetaOk() (*SempMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *MsgVpnJndiConnectionFactoriesResponse) SetMeta(v SempMeta)`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


