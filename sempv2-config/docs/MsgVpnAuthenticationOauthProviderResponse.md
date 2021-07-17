# MsgVpnAuthenticationOauthProviderResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**MsgVpnAuthenticationOauthProvider**](MsgVpnAuthenticationOauthProvider.md) |  | [optional] 
**Links** | Pointer to [**MsgVpnAuthenticationOauthProviderLinks**](MsgVpnAuthenticationOauthProviderLinks.md) |  | [optional] 
**Meta** | [**SempMeta**](SempMeta.md) |  | 

## Methods

### NewMsgVpnAuthenticationOauthProviderResponse

`func NewMsgVpnAuthenticationOauthProviderResponse(meta SempMeta, ) *MsgVpnAuthenticationOauthProviderResponse`

NewMsgVpnAuthenticationOauthProviderResponse instantiates a new MsgVpnAuthenticationOauthProviderResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMsgVpnAuthenticationOauthProviderResponseWithDefaults

`func NewMsgVpnAuthenticationOauthProviderResponseWithDefaults() *MsgVpnAuthenticationOauthProviderResponse`

NewMsgVpnAuthenticationOauthProviderResponseWithDefaults instantiates a new MsgVpnAuthenticationOauthProviderResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *MsgVpnAuthenticationOauthProviderResponse) GetData() MsgVpnAuthenticationOauthProvider`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *MsgVpnAuthenticationOauthProviderResponse) GetDataOk() (*MsgVpnAuthenticationOauthProvider, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *MsgVpnAuthenticationOauthProviderResponse) SetData(v MsgVpnAuthenticationOauthProvider)`

SetData sets Data field to given value.

### HasData

`func (o *MsgVpnAuthenticationOauthProviderResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetLinks

`func (o *MsgVpnAuthenticationOauthProviderResponse) GetLinks() MsgVpnAuthenticationOauthProviderLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *MsgVpnAuthenticationOauthProviderResponse) GetLinksOk() (*MsgVpnAuthenticationOauthProviderLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *MsgVpnAuthenticationOauthProviderResponse) SetLinks(v MsgVpnAuthenticationOauthProviderLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *MsgVpnAuthenticationOauthProviderResponse) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetMeta

`func (o *MsgVpnAuthenticationOauthProviderResponse) GetMeta() SempMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *MsgVpnAuthenticationOauthProviderResponse) GetMetaOk() (*SempMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *MsgVpnAuthenticationOauthProviderResponse) SetMeta(v SempMeta)`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


