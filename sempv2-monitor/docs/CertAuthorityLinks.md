# CertAuthorityLinks

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OcspTlsTrustedCommonNamesUri** | Pointer to **string** | The URI of this Certificate Authority&#39;s collection of OCSP Responder Trusted Common Name objects. Deprecated since 2.19. Replaced by clientCertAuthorities. | [optional] 
**Uri** | Pointer to **string** | The URI of this Certificate Authority object. | [optional] 

## Methods

### NewCertAuthorityLinks

`func NewCertAuthorityLinks() *CertAuthorityLinks`

NewCertAuthorityLinks instantiates a new CertAuthorityLinks object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCertAuthorityLinksWithDefaults

`func NewCertAuthorityLinksWithDefaults() *CertAuthorityLinks`

NewCertAuthorityLinksWithDefaults instantiates a new CertAuthorityLinks object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOcspTlsTrustedCommonNamesUri

`func (o *CertAuthorityLinks) GetOcspTlsTrustedCommonNamesUri() string`

GetOcspTlsTrustedCommonNamesUri returns the OcspTlsTrustedCommonNamesUri field if non-nil, zero value otherwise.

### GetOcspTlsTrustedCommonNamesUriOk

`func (o *CertAuthorityLinks) GetOcspTlsTrustedCommonNamesUriOk() (*string, bool)`

GetOcspTlsTrustedCommonNamesUriOk returns a tuple with the OcspTlsTrustedCommonNamesUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOcspTlsTrustedCommonNamesUri

`func (o *CertAuthorityLinks) SetOcspTlsTrustedCommonNamesUri(v string)`

SetOcspTlsTrustedCommonNamesUri sets OcspTlsTrustedCommonNamesUri field to given value.

### HasOcspTlsTrustedCommonNamesUri

`func (o *CertAuthorityLinks) HasOcspTlsTrustedCommonNamesUri() bool`

HasOcspTlsTrustedCommonNamesUri returns a boolean if a field has been set.

### GetUri

`func (o *CertAuthorityLinks) GetUri() string`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *CertAuthorityLinks) GetUriOk() (*string, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *CertAuthorityLinks) SetUri(v string)`

SetUri sets Uri field to given value.

### HasUri

`func (o *CertAuthorityLinks) HasUri() bool`

HasUri returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


