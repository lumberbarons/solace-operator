# MsgVpnAuthenticationOauthProvidersResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Collections** | Pointer to **[]map[string]interface{}** |  | [optional] 
**Data** | Pointer to [**[]MsgVpnAuthenticationOauthProvider**](MsgVpnAuthenticationOauthProvider.md) |  | [optional] 
**Links** | Pointer to [**[]MsgVpnAuthenticationOauthProviderLinks**](MsgVpnAuthenticationOauthProviderLinks.md) |  | [optional] 
**Meta** | [**SempMeta**](SempMeta.md) |  | 

## Methods

### NewMsgVpnAuthenticationOauthProvidersResponse

`func NewMsgVpnAuthenticationOauthProvidersResponse(meta SempMeta, ) *MsgVpnAuthenticationOauthProvidersResponse`

NewMsgVpnAuthenticationOauthProvidersResponse instantiates a new MsgVpnAuthenticationOauthProvidersResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMsgVpnAuthenticationOauthProvidersResponseWithDefaults

`func NewMsgVpnAuthenticationOauthProvidersResponseWithDefaults() *MsgVpnAuthenticationOauthProvidersResponse`

NewMsgVpnAuthenticationOauthProvidersResponseWithDefaults instantiates a new MsgVpnAuthenticationOauthProvidersResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCollections

`func (o *MsgVpnAuthenticationOauthProvidersResponse) GetCollections() []map[string]interface{}`

GetCollections returns the Collections field if non-nil, zero value otherwise.

### GetCollectionsOk

`func (o *MsgVpnAuthenticationOauthProvidersResponse) GetCollectionsOk() (*[]map[string]interface{}, bool)`

GetCollectionsOk returns a tuple with the Collections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollections

`func (o *MsgVpnAuthenticationOauthProvidersResponse) SetCollections(v []map[string]interface{})`

SetCollections sets Collections field to given value.

### HasCollections

`func (o *MsgVpnAuthenticationOauthProvidersResponse) HasCollections() bool`

HasCollections returns a boolean if a field has been set.

### GetData

`func (o *MsgVpnAuthenticationOauthProvidersResponse) GetData() []MsgVpnAuthenticationOauthProvider`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *MsgVpnAuthenticationOauthProvidersResponse) GetDataOk() (*[]MsgVpnAuthenticationOauthProvider, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *MsgVpnAuthenticationOauthProvidersResponse) SetData(v []MsgVpnAuthenticationOauthProvider)`

SetData sets Data field to given value.

### HasData

`func (o *MsgVpnAuthenticationOauthProvidersResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetLinks

`func (o *MsgVpnAuthenticationOauthProvidersResponse) GetLinks() []MsgVpnAuthenticationOauthProviderLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *MsgVpnAuthenticationOauthProvidersResponse) GetLinksOk() (*[]MsgVpnAuthenticationOauthProviderLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *MsgVpnAuthenticationOauthProvidersResponse) SetLinks(v []MsgVpnAuthenticationOauthProviderLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *MsgVpnAuthenticationOauthProvidersResponse) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetMeta

`func (o *MsgVpnAuthenticationOauthProvidersResponse) GetMeta() SempMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *MsgVpnAuthenticationOauthProvidersResponse) GetMetaOk() (*SempMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *MsgVpnAuthenticationOauthProvidersResponse) SetMeta(v SempMeta)`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


