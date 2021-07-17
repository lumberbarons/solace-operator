# ClientCertAuthorityOcspTlsTrustedCommonNamesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]ClientCertAuthorityOcspTlsTrustedCommonName**](ClientCertAuthorityOcspTlsTrustedCommonName.md) |  | [optional] 
**Links** | Pointer to [**[]ClientCertAuthorityOcspTlsTrustedCommonNameLinks**](ClientCertAuthorityOcspTlsTrustedCommonNameLinks.md) |  | [optional] 
**Meta** | [**SempMeta**](SempMeta.md) |  | 

## Methods

### NewClientCertAuthorityOcspTlsTrustedCommonNamesResponse

`func NewClientCertAuthorityOcspTlsTrustedCommonNamesResponse(meta SempMeta, ) *ClientCertAuthorityOcspTlsTrustedCommonNamesResponse`

NewClientCertAuthorityOcspTlsTrustedCommonNamesResponse instantiates a new ClientCertAuthorityOcspTlsTrustedCommonNamesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClientCertAuthorityOcspTlsTrustedCommonNamesResponseWithDefaults

`func NewClientCertAuthorityOcspTlsTrustedCommonNamesResponseWithDefaults() *ClientCertAuthorityOcspTlsTrustedCommonNamesResponse`

NewClientCertAuthorityOcspTlsTrustedCommonNamesResponseWithDefaults instantiates a new ClientCertAuthorityOcspTlsTrustedCommonNamesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ClientCertAuthorityOcspTlsTrustedCommonNamesResponse) GetData() []ClientCertAuthorityOcspTlsTrustedCommonName`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ClientCertAuthorityOcspTlsTrustedCommonNamesResponse) GetDataOk() (*[]ClientCertAuthorityOcspTlsTrustedCommonName, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ClientCertAuthorityOcspTlsTrustedCommonNamesResponse) SetData(v []ClientCertAuthorityOcspTlsTrustedCommonName)`

SetData sets Data field to given value.

### HasData

`func (o *ClientCertAuthorityOcspTlsTrustedCommonNamesResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetLinks

`func (o *ClientCertAuthorityOcspTlsTrustedCommonNamesResponse) GetLinks() []ClientCertAuthorityOcspTlsTrustedCommonNameLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ClientCertAuthorityOcspTlsTrustedCommonNamesResponse) GetLinksOk() (*[]ClientCertAuthorityOcspTlsTrustedCommonNameLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ClientCertAuthorityOcspTlsTrustedCommonNamesResponse) SetLinks(v []ClientCertAuthorityOcspTlsTrustedCommonNameLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ClientCertAuthorityOcspTlsTrustedCommonNamesResponse) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetMeta

`func (o *ClientCertAuthorityOcspTlsTrustedCommonNamesResponse) GetMeta() SempMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ClientCertAuthorityOcspTlsTrustedCommonNamesResponse) GetMetaOk() (*SempMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ClientCertAuthorityOcspTlsTrustedCommonNamesResponse) SetMeta(v SempMeta)`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


