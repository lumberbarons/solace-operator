# ClientCertAuthorityLinks

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OcspTlsTrustedCommonNamesUri** | Pointer to **string** | The URI of this Client Certificate Authority&#39;s collection of OCSP Responder Trusted Common Name objects. | [optional] 
**Uri** | Pointer to **string** | The URI of this Client Certificate Authority object. | [optional] 

## Methods

### NewClientCertAuthorityLinks

`func NewClientCertAuthorityLinks() *ClientCertAuthorityLinks`

NewClientCertAuthorityLinks instantiates a new ClientCertAuthorityLinks object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClientCertAuthorityLinksWithDefaults

`func NewClientCertAuthorityLinksWithDefaults() *ClientCertAuthorityLinks`

NewClientCertAuthorityLinksWithDefaults instantiates a new ClientCertAuthorityLinks object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOcspTlsTrustedCommonNamesUri

`func (o *ClientCertAuthorityLinks) GetOcspTlsTrustedCommonNamesUri() string`

GetOcspTlsTrustedCommonNamesUri returns the OcspTlsTrustedCommonNamesUri field if non-nil, zero value otherwise.

### GetOcspTlsTrustedCommonNamesUriOk

`func (o *ClientCertAuthorityLinks) GetOcspTlsTrustedCommonNamesUriOk() (*string, bool)`

GetOcspTlsTrustedCommonNamesUriOk returns a tuple with the OcspTlsTrustedCommonNamesUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOcspTlsTrustedCommonNamesUri

`func (o *ClientCertAuthorityLinks) SetOcspTlsTrustedCommonNamesUri(v string)`

SetOcspTlsTrustedCommonNamesUri sets OcspTlsTrustedCommonNamesUri field to given value.

### HasOcspTlsTrustedCommonNamesUri

`func (o *ClientCertAuthorityLinks) HasOcspTlsTrustedCommonNamesUri() bool`

HasOcspTlsTrustedCommonNamesUri returns a boolean if a field has been set.

### GetUri

`func (o *ClientCertAuthorityLinks) GetUri() string`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *ClientCertAuthorityLinks) GetUriOk() (*string, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *ClientCertAuthorityLinks) SetUri(v string)`

SetUri sets Uri field to given value.

### HasUri

`func (o *ClientCertAuthorityLinks) HasUri() bool`

HasUri returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


